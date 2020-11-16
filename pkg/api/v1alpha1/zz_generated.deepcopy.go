// +build !ignore_autogenerated

/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlternatorSpec) DeepCopyInto(out *AlternatorSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlternatorSpec.
func (in *AlternatorSpec) DeepCopy() *AlternatorSpec {
	if in == nil {
		return nil
	}
	out := new(AlternatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupTaskSpec) DeepCopyInto(out *BackupTaskSpec) {
	*out = *in
	in.SchedulerTaskSpec.DeepCopyInto(&out.SchedulerTaskSpec)
	if in.DC != nil {
		in, out := &in.DC, &out.DC
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Keyspace != nil {
		in, out := &in.Keyspace, &out.Keyspace
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Location != nil {
		in, out := &in.Location, &out.Location
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RateLimit != nil {
		in, out := &in.RateLimit, &out.RateLimit
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Retention != nil {
		in, out := &in.Retention, &out.Retention
		*out = new(int64)
		**out = **in
	}
	if in.SnapshotParallel != nil {
		in, out := &in.SnapshotParallel, &out.SnapshotParallel
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UploadParallel != nil {
		in, out := &in.UploadParallel, &out.UploadParallel
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupTaskSpec.
func (in *BackupTaskSpec) DeepCopy() *BackupTaskSpec {
	if in == nil {
		return nil
	}
	out := new(BackupTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupTaskStatus) DeepCopyInto(out *BackupTaskStatus) {
	*out = *in
	in.BackupTaskSpec.DeepCopyInto(&out.BackupTaskSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupTaskStatus.
func (in *BackupTaskStatus) DeepCopy() *BackupTaskStatus {
	if in == nil {
		return nil
	}
	out := new(BackupTaskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(string)
		**out = **in
	}
	if in.Alternator != nil {
		in, out := &in.Alternator, &out.Alternator
		*out = new(AlternatorSpec)
		**out = **in
	}
	if in.AgentVersion != nil {
		in, out := &in.AgentVersion, &out.AgentVersion
		*out = new(string)
		**out = **in
	}
	if in.AgentRepository != nil {
		in, out := &in.AgentRepository, &out.AgentRepository
		*out = new(string)
		**out = **in
	}
	in.Datacenter.DeepCopyInto(&out.Datacenter)
	if in.SidecarImage != nil {
		in, out := &in.SidecarImage, &out.SidecarImage
		*out = new(ImageSpec)
		**out = **in
	}
	if in.Sysctls != nil {
		in, out := &in.Sysctls, &out.Sysctls
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.Network = in.Network
	if in.Repairs != nil {
		in, out := &in.Repairs, &out.Repairs
		*out = make([]RepairTaskSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = make([]BackupTaskSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.Racks != nil {
		in, out := &in.Racks, &out.Racks
		*out = make(map[string]RackStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ManagerID != nil {
		in, out := &in.ManagerID, &out.ManagerID
		*out = new(string)
		**out = **in
	}
	if in.Repairs != nil {
		in, out := &in.Repairs, &out.Repairs
		*out = make([]RepairTaskStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Backups != nil {
		in, out := &in.Backups, &out.Backups
		*out = make([]BackupTaskStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatacenterSpec) DeepCopyInto(out *DatacenterSpec) {
	*out = *in
	if in.Racks != nil {
		in, out := &in.Racks, &out.Racks
		*out = make([]RackSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatacenterSpec.
func (in *DatacenterSpec) DeepCopy() *DatacenterSpec {
	if in == nil {
		return nil
	}
	out := new(DatacenterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementSpec) DeepCopyInto(out *PlacementSpec) {
	*out = *in
	if in.NodeAffinity != nil {
		in, out := &in.NodeAffinity, &out.NodeAffinity
		*out = new(v1.NodeAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodAffinity != nil {
		in, out := &in.PodAffinity, &out.PodAffinity
		*out = new(v1.PodAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.PodAntiAffinity != nil {
		in, out := &in.PodAntiAffinity, &out.PodAntiAffinity
		*out = new(v1.PodAntiAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementSpec.
func (in *PlacementSpec) DeepCopy() *PlacementSpec {
	if in == nil {
		return nil
	}
	out := new(PlacementSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RackCondition) DeepCopyInto(out *RackCondition) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RackCondition.
func (in *RackCondition) DeepCopy() *RackCondition {
	if in == nil {
		return nil
	}
	out := new(RackCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RackSpec) DeepCopyInto(out *RackSpec) {
	*out = *in
	in.Storage.DeepCopyInto(&out.Storage)
	if in.Placement != nil {
		in, out := &in.Placement, &out.Placement
		*out = new(PlacementSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RackSpec.
func (in *RackSpec) DeepCopy() *RackSpec {
	if in == nil {
		return nil
	}
	out := new(RackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RackStatus) DeepCopyInto(out *RackStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]RackCondition, len(*in))
		copy(*out, *in)
	}
	if in.ReplaceAddresses != nil {
		in, out := &in.ReplaceAddresses, &out.ReplaceAddresses
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RackStatus.
func (in *RackStatus) DeepCopy() *RackStatus {
	if in == nil {
		return nil
	}
	out := new(RackStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepairTaskSpec) DeepCopyInto(out *RepairTaskSpec) {
	*out = *in
	in.SchedulerTaskSpec.DeepCopyInto(&out.SchedulerTaskSpec)
	if in.DC != nil {
		in, out := &in.DC, &out.DC
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FailFast != nil {
		in, out := &in.FailFast, &out.FailFast
		*out = new(bool)
		**out = **in
	}
	if in.Intensity != nil {
		in, out := &in.Intensity, &out.Intensity
		*out = new(int64)
		**out = **in
	}
	if in.Parallel != nil {
		in, out := &in.Parallel, &out.Parallel
		*out = new(int64)
		**out = **in
	}
	if in.Keyspace != nil {
		in, out := &in.Keyspace, &out.Keyspace
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SmallTableThreshold != nil {
		in, out := &in.SmallTableThreshold, &out.SmallTableThreshold
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepairTaskSpec.
func (in *RepairTaskSpec) DeepCopy() *RepairTaskSpec {
	if in == nil {
		return nil
	}
	out := new(RepairTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepairTaskStatus) DeepCopyInto(out *RepairTaskStatus) {
	*out = *in
	in.RepairTaskSpec.DeepCopyInto(&out.RepairTaskSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepairTaskStatus.
func (in *RepairTaskStatus) DeepCopy() *RepairTaskStatus {
	if in == nil {
		return nil
	}
	out := new(RepairTaskStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerTaskSpec) DeepCopyInto(out *SchedulerTaskSpec) {
	*out = *in
	if in.StartDate != nil {
		in, out := &in.StartDate, &out.StartDate
		*out = new(string)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(string)
		**out = **in
	}
	if in.NumRetries != nil {
		in, out := &in.NumRetries, &out.NumRetries
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerTaskSpec.
func (in *SchedulerTaskSpec) DeepCopy() *SchedulerTaskSpec {
	if in == nil {
		return nil
	}
	out := new(SchedulerTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScyllaCluster) DeepCopyInto(out *ScyllaCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScyllaCluster.
func (in *ScyllaCluster) DeepCopy() *ScyllaCluster {
	if in == nil {
		return nil
	}
	out := new(ScyllaCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScyllaCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScyllaClusterList) DeepCopyInto(out *ScyllaClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScyllaCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScyllaClusterList.
func (in *ScyllaClusterList) DeepCopy() *ScyllaClusterList {
	if in == nil {
		return nil
	}
	out := new(ScyllaClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScyllaClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSpec) DeepCopyInto(out *StorageSpec) {
	*out = *in
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSpec.
func (in *StorageSpec) DeepCopy() *StorageSpec {
	if in == nil {
		return nil
	}
	out := new(StorageSpec)
	in.DeepCopyInto(out)
	return out
}
