// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	version "github.com/percona/percona-server-mongodb-operator/version"
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Arbiter) DeepCopyInto(out *Arbiter) {
	*out = *in
	in.MultiAZ.DeepCopyInto(&out.MultiAZ)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Arbiter.
func (in *Arbiter) DeepCopy() *Arbiter {
	if in == nil {
		return nil
	}
	out := new(Arbiter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupCoordinatorSpec) DeepCopyInto(out *BackupCoordinatorSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	in.MultiAZ.DeepCopyInto(&out.MultiAZ)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupCoordinatorSpec.
func (in *BackupCoordinatorSpec) DeepCopy() *BackupCoordinatorSpec {
	if in == nil {
		return nil
	}
	out := new(BackupCoordinatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSpec) DeepCopyInto(out *BackupSpec) {
	*out = *in
	if in.RestartOnFailure != nil {
		in, out := &in.RestartOnFailure, &out.RestartOnFailure
		*out = new(bool)
		**out = **in
	}
	in.Coordinator.DeepCopyInto(&out.Coordinator)
	if in.Storages != nil {
		in, out := &in.Storages, &out.Storages
		*out = make(map[string]BackupStorageSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]BackupTaskSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSpec.
func (in *BackupSpec) DeepCopy() *BackupSpec {
	if in == nil {
		return nil
	}
	out := new(BackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStorageS3Spec) DeepCopyInto(out *BackupStorageS3Spec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStorageS3Spec.
func (in *BackupStorageS3Spec) DeepCopy() *BackupStorageS3Spec {
	if in == nil {
		return nil
	}
	out := new(BackupStorageS3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStorageSpec) DeepCopyInto(out *BackupStorageSpec) {
	*out = *in
	out.S3 = in.S3
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStorageSpec.
func (in *BackupStorageSpec) DeepCopy() *BackupStorageSpec {
	if in == nil {
		return nil
	}
	out := new(BackupStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupTaskSpec) DeepCopyInto(out *BackupTaskSpec) {
	*out = *in
	return
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
func (in *Expose) DeepCopyInto(out *Expose) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Expose.
func (in *Expose) DeepCopy() *Expose {
	if in == nil {
		return nil
	}
	out := new(Expose)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpec) DeepCopyInto(out *MongodSpec) {
	*out = *in
	if in.Net != nil {
		in, out := &in.Net, &out.Net
		*out = new(MongodSpecNet)
		**out = **in
	}
	if in.AuditLog != nil {
		in, out := &in.AuditLog, &out.AuditLog
		*out = new(MongodSpecAuditLog)
		**out = **in
	}
	if in.OperationProfiling != nil {
		in, out := &in.OperationProfiling, &out.OperationProfiling
		*out = new(MongodSpecOperationProfiling)
		**out = **in
	}
	if in.Replication != nil {
		in, out := &in.Replication, &out.Replication
		*out = new(MongodSpecReplication)
		**out = **in
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = new(MongodSpecSecurity)
		**out = **in
	}
	if in.SetParameter != nil {
		in, out := &in.SetParameter, &out.SetParameter
		*out = new(MongodSpecSetParameter)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(MongodSpecStorage)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpec.
func (in *MongodSpec) DeepCopy() *MongodSpec {
	if in == nil {
		return nil
	}
	out := new(MongodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecAuditLog) DeepCopyInto(out *MongodSpecAuditLog) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecAuditLog.
func (in *MongodSpecAuditLog) DeepCopy() *MongodSpecAuditLog {
	if in == nil {
		return nil
	}
	out := new(MongodSpecAuditLog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecInMemory) DeepCopyInto(out *MongodSpecInMemory) {
	*out = *in
	if in.EngineConfig != nil {
		in, out := &in.EngineConfig, &out.EngineConfig
		*out = new(MongodSpecInMemoryEngineConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecInMemory.
func (in *MongodSpecInMemory) DeepCopy() *MongodSpecInMemory {
	if in == nil {
		return nil
	}
	out := new(MongodSpecInMemory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecInMemoryEngineConfig) DeepCopyInto(out *MongodSpecInMemoryEngineConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecInMemoryEngineConfig.
func (in *MongodSpecInMemoryEngineConfig) DeepCopy() *MongodSpecInMemoryEngineConfig {
	if in == nil {
		return nil
	}
	out := new(MongodSpecInMemoryEngineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecMMAPv1) DeepCopyInto(out *MongodSpecMMAPv1) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecMMAPv1.
func (in *MongodSpecMMAPv1) DeepCopy() *MongodSpecMMAPv1 {
	if in == nil {
		return nil
	}
	out := new(MongodSpecMMAPv1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecNet) DeepCopyInto(out *MongodSpecNet) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecNet.
func (in *MongodSpecNet) DeepCopy() *MongodSpecNet {
	if in == nil {
		return nil
	}
	out := new(MongodSpecNet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecOperationProfiling) DeepCopyInto(out *MongodSpecOperationProfiling) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecOperationProfiling.
func (in *MongodSpecOperationProfiling) DeepCopy() *MongodSpecOperationProfiling {
	if in == nil {
		return nil
	}
	out := new(MongodSpecOperationProfiling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecReplication) DeepCopyInto(out *MongodSpecReplication) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecReplication.
func (in *MongodSpecReplication) DeepCopy() *MongodSpecReplication {
	if in == nil {
		return nil
	}
	out := new(MongodSpecReplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecSecurity) DeepCopyInto(out *MongodSpecSecurity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecSecurity.
func (in *MongodSpecSecurity) DeepCopy() *MongodSpecSecurity {
	if in == nil {
		return nil
	}
	out := new(MongodSpecSecurity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecSetParameter) DeepCopyInto(out *MongodSpecSetParameter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecSetParameter.
func (in *MongodSpecSetParameter) DeepCopy() *MongodSpecSetParameter {
	if in == nil {
		return nil
	}
	out := new(MongodSpecSetParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecStorage) DeepCopyInto(out *MongodSpecStorage) {
	*out = *in
	if in.InMemory != nil {
		in, out := &in.InMemory, &out.InMemory
		*out = new(MongodSpecInMemory)
		(*in).DeepCopyInto(*out)
	}
	if in.MMAPv1 != nil {
		in, out := &in.MMAPv1, &out.MMAPv1
		*out = new(MongodSpecMMAPv1)
		**out = **in
	}
	if in.WiredTiger != nil {
		in, out := &in.WiredTiger, &out.WiredTiger
		*out = new(MongodSpecWiredTiger)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecStorage.
func (in *MongodSpecStorage) DeepCopy() *MongodSpecStorage {
	if in == nil {
		return nil
	}
	out := new(MongodSpecStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecWiredTiger) DeepCopyInto(out *MongodSpecWiredTiger) {
	*out = *in
	if in.CollectionConfig != nil {
		in, out := &in.CollectionConfig, &out.CollectionConfig
		*out = new(MongodSpecWiredTigerCollectionConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.EngineConfig != nil {
		in, out := &in.EngineConfig, &out.EngineConfig
		*out = new(MongodSpecWiredTigerEngineConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.IndexConfig != nil {
		in, out := &in.IndexConfig, &out.IndexConfig
		*out = new(MongodSpecWiredTigerIndexConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecWiredTiger.
func (in *MongodSpecWiredTiger) DeepCopy() *MongodSpecWiredTiger {
	if in == nil {
		return nil
	}
	out := new(MongodSpecWiredTiger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecWiredTigerCollectionConfig) DeepCopyInto(out *MongodSpecWiredTigerCollectionConfig) {
	*out = *in
	if in.BlockCompressor != nil {
		in, out := &in.BlockCompressor, &out.BlockCompressor
		*out = new(WiredTigerCompressor)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecWiredTigerCollectionConfig.
func (in *MongodSpecWiredTigerCollectionConfig) DeepCopy() *MongodSpecWiredTigerCollectionConfig {
	if in == nil {
		return nil
	}
	out := new(MongodSpecWiredTigerCollectionConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecWiredTigerEngineConfig) DeepCopyInto(out *MongodSpecWiredTigerEngineConfig) {
	*out = *in
	if in.JournalCompressor != nil {
		in, out := &in.JournalCompressor, &out.JournalCompressor
		*out = new(WiredTigerCompressor)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecWiredTigerEngineConfig.
func (in *MongodSpecWiredTigerEngineConfig) DeepCopy() *MongodSpecWiredTigerEngineConfig {
	if in == nil {
		return nil
	}
	out := new(MongodSpecWiredTigerEngineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongodSpecWiredTigerIndexConfig) DeepCopyInto(out *MongodSpecWiredTigerIndexConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongodSpecWiredTigerIndexConfig.
func (in *MongodSpecWiredTigerIndexConfig) DeepCopy() *MongodSpecWiredTigerIndexConfig {
	if in == nil {
		return nil
	}
	out := new(MongodSpecWiredTigerIndexConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongosSpec) DeepCopyInto(out *MongosSpec) {
	*out = *in
	if in.ResourcesSpec != nil {
		in, out := &in.ResourcesSpec, &out.ResourcesSpec
		*out = new(ResourcesSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongosSpec.
func (in *MongosSpec) DeepCopy() *MongosSpec {
	if in == nil {
		return nil
	}
	out := new(MongosSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiAZ) DeepCopyInto(out *MultiAZ) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(PodAffinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiAZ.
func (in *MultiAZ) DeepCopy() *MultiAZ {
	if in == nil {
		return nil
	}
	out := new(MultiAZ)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PMMSpec) DeepCopyInto(out *PMMSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PMMSpec.
func (in *PMMSpec) DeepCopy() *PMMSpec {
	if in == nil {
		return nil
	}
	out := new(PMMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDB) DeepCopyInto(out *PerconaServerMongoDB) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDB.
func (in *PerconaServerMongoDB) DeepCopy() *PerconaServerMongoDB {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDB)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMongoDB) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBBackup) DeepCopyInto(out *PerconaServerMongoDBBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBBackup.
func (in *PerconaServerMongoDBBackup) DeepCopy() *PerconaServerMongoDBBackup {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMongoDBBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBBackupList) DeepCopyInto(out *PerconaServerMongoDBBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaServerMongoDBBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBBackupList.
func (in *PerconaServerMongoDBBackupList) DeepCopy() *PerconaServerMongoDBBackupList {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMongoDBBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBBackupSpec) DeepCopyInto(out *PerconaServerMongoDBBackupSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBBackupSpec.
func (in *PerconaServerMongoDBBackupSpec) DeepCopy() *PerconaServerMongoDBBackupSpec {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBBackupStatus) DeepCopyInto(out *PerconaServerMongoDBBackupStatus) {
	*out = *in
	if in.StartAt != nil {
		in, out := &in.StartAt, &out.StartAt
		*out = (*in).DeepCopy()
	}
	if in.CompletedAt != nil {
		in, out := &in.CompletedAt, &out.CompletedAt
		*out = (*in).DeepCopy()
	}
	if in.LastScheduled != nil {
		in, out := &in.LastScheduled, &out.LastScheduled
		*out = (*in).DeepCopy()
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(BackupStorageS3Spec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBBackupStatus.
func (in *PerconaServerMongoDBBackupStatus) DeepCopy() *PerconaServerMongoDBBackupStatus {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBList) DeepCopyInto(out *PerconaServerMongoDBList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaServerMongoDB, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBList.
func (in *PerconaServerMongoDBList) DeepCopy() *PerconaServerMongoDBList {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMongoDBList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBRestore) DeepCopyInto(out *PerconaServerMongoDBRestore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBRestore.
func (in *PerconaServerMongoDBRestore) DeepCopy() *PerconaServerMongoDBRestore {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMongoDBRestore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBRestoreList) DeepCopyInto(out *PerconaServerMongoDBRestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaServerMongoDBRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBRestoreList.
func (in *PerconaServerMongoDBRestoreList) DeepCopy() *PerconaServerMongoDBRestoreList {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBRestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaServerMongoDBRestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBRestoreSpec) DeepCopyInto(out *PerconaServerMongoDBRestoreSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBRestoreSpec.
func (in *PerconaServerMongoDBRestoreSpec) DeepCopy() *PerconaServerMongoDBRestoreSpec {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBRestoreStatus) DeepCopyInto(out *PerconaServerMongoDBRestoreStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBRestoreStatus.
func (in *PerconaServerMongoDBRestoreStatus) DeepCopy() *PerconaServerMongoDBRestoreStatus {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBRestoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBSpec) DeepCopyInto(out *PerconaServerMongoDBSpec) {
	*out = *in
	if in.Platform != nil {
		in, out := &in.Platform, &out.Platform
		*out = new(version.Platform)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Mongod != nil {
		in, out := &in.Mongod, &out.Mongod
		*out = new(MongodSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Replsets != nil {
		in, out := &in.Replsets, &out.Replsets
		*out = make([]*ReplsetSpec, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ReplsetSpec)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = new(SecretsSpec)
		**out = **in
	}
	in.Backup.DeepCopyInto(&out.Backup)
	out.PMM = in.PMM
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBSpec.
func (in *PerconaServerMongoDBSpec) DeepCopy() *PerconaServerMongoDBSpec {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaServerMongoDBStatus) DeepCopyInto(out *PerconaServerMongoDBStatus) {
	*out = *in
	if in.Replsets != nil {
		in, out := &in.Replsets, &out.Replsets
		*out = make(map[string]*ReplsetStatus, len(*in))
		for key, val := range *in {
			var outVal *ReplsetStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ReplsetStatus)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaServerMongoDBStatus.
func (in *PerconaServerMongoDBStatus) DeepCopy() *PerconaServerMongoDBStatus {
	if in == nil {
		return nil
	}
	out := new(PerconaServerMongoDBStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodAffinity) DeepCopyInto(out *PodAffinity) {
	*out = *in
	if in.TopologyKey != nil {
		in, out := &in.TopologyKey, &out.TopologyKey
		*out = new(string)
		**out = **in
	}
	if in.Advanced != nil {
		in, out := &in.Advanced, &out.Advanced
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodAffinity.
func (in *PodAffinity) DeepCopy() *PodAffinity {
	if in == nil {
		return nil
	}
	out := new(PodAffinity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetSpec) DeepCopyInto(out *PodDisruptionBudgetSpec) {
	*out = *in
	if in.MinAvailable != nil {
		in, out := &in.MinAvailable, &out.MinAvailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetSpec.
func (in *PodDisruptionBudgetSpec) DeepCopy() *PodDisruptionBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplsetMemberStatus) DeepCopyInto(out *ReplsetMemberStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplsetMemberStatus.
func (in *ReplsetMemberStatus) DeepCopy() *ReplsetMemberStatus {
	if in == nil {
		return nil
	}
	out := new(ReplsetMemberStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplsetSpec) DeepCopyInto(out *ReplsetSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourcesSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Arbiter.DeepCopyInto(&out.Arbiter)
	out.Expose = in.Expose
	if in.VolumeSpec != nil {
		in, out := &in.VolumeSpec, &out.VolumeSpec
		*out = new(VolumeSpec)
		(*in).DeepCopyInto(*out)
	}
	in.MultiAZ.DeepCopyInto(&out.MultiAZ)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplsetSpec.
func (in *ReplsetSpec) DeepCopy() *ReplsetSpec {
	if in == nil {
		return nil
	}
	out := new(ReplsetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplsetStatus) DeepCopyInto(out *ReplsetStatus) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]*ReplsetMemberStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ReplsetMemberStatus)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplsetStatus.
func (in *ReplsetStatus) DeepCopy() *ReplsetStatus {
	if in == nil {
		return nil
	}
	out := new(ReplsetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpecRequirements) DeepCopyInto(out *ResourceSpecRequirements) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpecRequirements.
func (in *ResourceSpecRequirements) DeepCopy() *ResourceSpecRequirements {
	if in == nil {
		return nil
	}
	out := new(ResourceSpecRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesSpec) DeepCopyInto(out *ResourcesSpec) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = new(ResourceSpecRequirements)
		**out = **in
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = new(ResourceSpecRequirements)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesSpec.
func (in *ResourcesSpec) DeepCopy() *ResourcesSpec {
	if in == nil {
		return nil
	}
	out := new(ResourcesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretsSpec) DeepCopyInto(out *SecretsSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretsSpec.
func (in *SecretsSpec) DeepCopy() *SecretsSpec {
	if in == nil {
		return nil
	}
	out := new(SecretsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerVersion) DeepCopyInto(out *ServerVersion) {
	*out = *in
	out.Info = in.Info
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerVersion.
func (in *ServerVersion) DeepCopy() *ServerVersion {
	if in == nil {
		return nil
	}
	out := new(ServerVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSpec) DeepCopyInto(out *VolumeSpec) {
	*out = *in
	if in.EmptyDir != nil {
		in, out := &in.EmptyDir, &out.EmptyDir
		*out = new(v1.EmptyDirVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.HostPath != nil {
		in, out := &in.HostPath, &out.HostPath
		*out = new(v1.HostPathVolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSpec.
func (in *VolumeSpec) DeepCopy() *VolumeSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeSpec)
	in.DeepCopyInto(out)
	return out
}
