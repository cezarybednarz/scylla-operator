package v1alpha1

import (
	"github.com/blang/semver"
	"github.com/pkg/errors"
	"github.com/scylladb/go-log"
	"github.com/scylladb/scylla-operator/pkg/auth"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"net/http"
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	"context"
	"github.com/scylladb/scylla-operator/pkg/scyllaclient"
	k8sClient "k8s.io/client-go/kubernetes"

	prometheusApi "github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

var (
	ctx    context.Context
	atom   zap.AtomicLevel
	logger log.Logger
	_      error
)

func init() {
	ctx = log.WithNewTraceID(context.Background())
	atom = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	logger, _ = log.NewProduction(log.Config{
		Level: atom,
	})
}

func checkValues(c *Cluster) error {

	// creating client for Prometheus API
	prometheusClient, err := prometheusApi.NewClient(prometheusApi.Config{
		Address: "http://scylla-prom-prometheus-server.monitoring.svc.cluster.local",
	})

	if err != nil {
		logger.Error(ctx, "Error creating client", "error", err)
		os.Exit(1)
	}

	v1Api := v1.NewAPI(prometheusClient)

	result, warnings, err := v1Api.Query(ctx, "up{app=\"scylla\"}", time.Now())
	if err != nil {
		logger.Error(ctx, "Error querying Prometheus", "error", err)
		return nil
	}
	if len(warnings) > 0 {
		logger.Error(ctx, "Warnings: ", "warnings", warnings)
	}
	logger.Info(ctx, "Result: ", "result", result)

	// $ k describe ScyllaCluster
	// Spec:
	//  ...
	//  Datacenter:
	//    ...
	//    Racks:
	//      ...
	//      Scylla Agent Config:  scylla-agent-config    <<-- THIS
	scheme := runtime.NewScheme()
	_ = AddToScheme(scheme)

	clientGo, err := client.New(config.GetConfigOrDie(), client.Options{Scheme: scheme})
	if err != nil {
		logger.Error(ctx, "failed to create a client")
	}
	scyllaCluster := &Cluster{}
	err = clientGo.Get(ctx, client.ObjectKey{
		Namespace: "scylla",
		Name:      "simple-cluster"}, scyllaCluster)
	if err != nil {
		logger.Error(ctx, "failed to get cluster 'simple-cluster' in namespace 'scylla'", "error", err)
	}
	logger.Info(ctx, "got", "scyllaCluster", scyllaCluster)
	scyllaAgentConfig := scyllaCluster.Spec.Datacenter.Racks[0].ScyllaAgentConfig

	// $ k describe secrets <Scylla Agent Config>
	// Name:         default-token-72fwl
	//Namespace:    scylla
	//Labels:       <none>
	//Annotations:  kubernetes.io/service-account.name: default
	//              kubernetes.io/service-account.uid: 8cbd9d85-2ac9-4f9b-99a9-146af28ac28b
	//Type:  kubernetes.io/service-account-token
	//
	//Data
	//ca.crt:     1111 bytes
	//namespace:  6 bytes
	//token:      eyJhbGc..abcd							 <<-- THIS

	//config, err := rest.InClusterConfig()
	config := config.GetConfigOrDie()
	clientset, err := k8sClient.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	secrets := clientset.CoreV1().Secrets("scylla")
	s, err := secrets.Get(context.TODO(), scyllaAgentConfig, metav1.GetOptions{})
	var bearerToken string
	if err != nil {
		bearerToken = ""
	} else {
		bearerToken = string(s.Data["token"])
	}
	logger.Info(ctx, "secrets got", "secret", bearerToken)

	// polaczyc sie z scylla-manager-agent na localhost:10001 z bearer token
	// tam składać zapytania
	// /storage_service/load (to jest w B, KB?)

	sc, err := scyllaclient.NewClient(scyllaclient.DefaultConfig(), logger)
	if err != nil {
		panic(err.Error())
	}

	transport := config.Transport
	transport = auth.AddToken(transport, bearerToken)
	httpClient := &http.Client{Transport: transport}
	// to debug
	logger.Info(ctx, "ping", "sc", sc, "err", err, "httpClient", httpClient)

	sc.AddBearerToken(bearerToken)
	load, err := sc.StorageServiceLoadGetWithParamsFromHTTPClient(ctx)
	logger.Info(ctx, "getLoad", "load", load, "error", err)

	// pojemność dysku
	// $ kubectl describe sts
	// Volume Claims:
	// ...
	//  Capacity:      5Gi              				   <<-- THIS

	/*for _, rack := range c.Spec.Datacenter.Racks {

		pods := clientset.CoreV1().Pods("scylla")
		p, err := pods.Get(context.TODO(), rack.Name, metav1.GetOptions{})
		capacity := p.Spec.Volumes
		logger.Info(ctx, "secrets got", "secret", bearerToken)
	}*/

	rackNames := sets.NewString()

	if len(c.Spec.ScyllaArgs) > 0 {
		version, err := semver.Parse(c.Spec.Version)
		if err == nil && version.LT(ScyllaVersionThatSupportsArgs) {
			return errors.Errorf("ScyllaArgs is only supported starting from %s", ScyllaVersionThatSupportsArgsText)
		}
	}

	for _, rack := range c.Spec.Datacenter.Racks {
		// Check that rack is not scaled to 0 members
		if rack.Members == 0 {
			return errors.Errorf("cannot scale rack '%s' to 0 members", rack.Name)
		}

		// Check that no two racks have the same name
		if rackNames.Has(rack.Name) {
			return errors.Errorf("two racks have the same name: '%s'", rack.Name)
		}
		rackNames.Insert(rack.Name)

		// Check that limits are defined
		limits := rack.Resources.Limits
		if limits == nil || limits.Cpu().Value() == 0 || limits.Memory().Value() == 0 {
			return errors.Errorf("set cpu, memory resource limits for rack %s", rack.Name)
		}

		// Check that requested values are not 0
		requests := rack.Resources.Requests
		if requests != nil {
			if requests.Cpu() != nil && requests.Cpu().MilliValue() == 0 {
				return errors.Errorf("requesting 0 cpus is invalid for rack %s", rack.Name)
			}
			if requests.Memory() != nil && requests.Memory().MilliValue() == 0 {
				return errors.Errorf("requesting 0 memory is invalid for rack %s", rack.Name)
			}
			if requests.Storage() != nil && requests.Storage().MilliValue() == 0 {
				return errors.Errorf("requesting 0 storage is invalid for rack %s", rack.Name)
			}
		}

		// If the cluster has cpuset
		if c.Spec.CpuSet {
			cores := limits.Cpu().MilliValue()

			// CPU limits must be whole cores
			if cores%1000 != 0 {
				return errors.Errorf("when using cpuset, you must use whole cpu cores, but rack %s has %dm", rack.Name, cores)
			}

			// Requests == Limits and Requests must be set and equal for QOS class guaranteed
			if requests != nil {
				if requests.Cpu().MilliValue() != limits.Cpu().MilliValue() {
					return errors.Errorf("when using cpuset, cpu requests must be the same as cpu limits in rack %s", rack.Name)
				}
				if requests.Memory().MilliValue() != limits.Memory().MilliValue() {
					return errors.Errorf("when using cpuset, memory requests must be the same as memory limits in rack %s", rack.Name)
				}
			} else {
				// Copy the limits
				rack.Resources.Requests = limits.DeepCopy()
			}
		}
	}

	return nil
}

func checkTransitions(old, new *Cluster) error {
	oldVersion, err := semver.Parse(old.Spec.Version)
	if err != nil {
		return errors.Errorf("invalid old semantic version, err=%s", err)
	}
	newVersion, err := semver.Parse(new.Spec.Version)
	if err != nil {
		return errors.Errorf("invalid new semantic version, err=%s", err)
	}
	// Check that version remained the same
	if newVersion.Major != oldVersion.Major || newVersion.Minor != oldVersion.Minor {
		return errors.Errorf("only upgrading of patch versions are supported")
	}

	// Check that repository remained the same
	if !reflect.DeepEqual(old.Spec.Repository, new.Spec.Repository) {
		return errors.Errorf("repository change is currently not supported, old=%v, new=%v", *old.Spec.Repository, *new.Spec.Repository)
	}

	// Check that sidecarImage remained the same
	if !reflect.DeepEqual(old.Spec.SidecarImage, new.Spec.SidecarImage) {
		return errors.Errorf("change of sidecarImage is currently not supported")
	}

	// Check that the datacenter name didn't change
	if old.Spec.Datacenter.Name != new.Spec.Datacenter.Name {
		return errors.Errorf("change of datacenter name is currently not supported")
	}

	// Check that all rack names are the same as before
	oldRackNames, newRackNames := sets.NewString(), sets.NewString()
	for _, rack := range old.Spec.Datacenter.Racks {
		oldRackNames.Insert(rack.Name)
	}
	for _, rack := range new.Spec.Datacenter.Racks {
		newRackNames.Insert(rack.Name)
	}
	diff := oldRackNames.Difference(newRackNames)
	if diff.Len() != 0 {
		return errors.Errorf("racks %v not found, you cannot remove racks from the spec", diff.List())
	}

	rackMap := make(map[string]RackSpec)
	for _, oldRack := range old.Spec.Datacenter.Racks {
		rackMap[oldRack.Name] = oldRack
	}
	for _, newRack := range new.Spec.Datacenter.Racks {
		oldRack, exists := rackMap[newRack.Name]
		if !exists {
			continue
		}

		// Check that placement is the same as before
		if !reflect.DeepEqual(oldRack.Placement, newRack.Placement) {
			return errors.Errorf("rack %s: changes in placement are not currently supported", oldRack.Name)
		}

		// Check that storage is the same as before
		if !reflect.DeepEqual(oldRack.Storage, newRack.Storage) {
			return errors.Errorf("rack %s: changes in storage are not currently supported", oldRack.Name)
		}

		// Check that resources are the same as before
		if !reflect.DeepEqual(oldRack.Resources, newRack.Resources) {
			return errors.Errorf("rack %s: changes in resources are not currently supported", oldRack.Name)
		}
	}

	return nil
}
