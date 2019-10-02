// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

package config

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupBucketControllerConfiguration) DeepCopyInto(out *BackupBucketControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupBucketControllerConfiguration.
func (in *BackupBucketControllerConfiguration) DeepCopy() *BackupBucketControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(BackupBucketControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupEntryControllerConfiguration) DeepCopyInto(out *BackupEntryControllerConfiguration) {
	*out = *in
	if in.DeletionGracePeriodHours != nil {
		in, out := &in.DeletionGracePeriodHours, &out.DeletionGracePeriodHours
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupEntryControllerConfiguration.
func (in *BackupEntryControllerConfiguration) DeepCopy() *BackupEntryControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(BackupEntryControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudProfileControllerConfiguration) DeepCopyInto(out *CloudProfileControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudProfileControllerConfiguration.
func (in *CloudProfileControllerConfiguration) DeepCopy() *CloudProfileControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(CloudProfileControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConditionThreshold) DeepCopyInto(out *ConditionThreshold) {
	*out = *in
	out.Duration = in.Duration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConditionThreshold.
func (in *ConditionThreshold) DeepCopy() *ConditionThreshold {
	if in == nil {
		return nil
	}
	out := new(ConditionThreshold)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerInstallationControllerConfiguration) DeepCopyInto(out *ControllerInstallationControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerInstallationControllerConfiguration.
func (in *ControllerInstallationControllerConfiguration) DeepCopy() *ControllerInstallationControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerInstallationControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManagerConfiguration) DeepCopyInto(out *ControllerManagerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ClientConnection = in.ClientConnection
	out.SeedClientConnection = in.SeedClientConnection
	out.ShootClientConnection = in.ShootClientConnection
	in.Controllers.DeepCopyInto(&out.Controllers)
	out.LeaderElection = in.LeaderElection
	in.Discovery.DeepCopyInto(&out.Discovery)
	out.Server = in.Server
	if in.ShootBackup != nil {
		in, out := &in.ShootBackup, &out.ShootBackup
		*out = new(ShootBackup)
		**out = **in
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManagerConfiguration.
func (in *ControllerManagerConfiguration) DeepCopy() *ControllerManagerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerManagerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerManagerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManagerControllerConfiguration) DeepCopyInto(out *ControllerManagerControllerConfiguration) {
	*out = *in
	if in.BackupBucket != nil {
		in, out := &in.BackupBucket, &out.BackupBucket
		*out = new(BackupBucketControllerConfiguration)
		**out = **in
	}
	if in.BackupEntry != nil {
		in, out := &in.BackupEntry, &out.BackupEntry
		*out = new(BackupEntryControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.CloudProfile != nil {
		in, out := &in.CloudProfile, &out.CloudProfile
		*out = new(CloudProfileControllerConfiguration)
		**out = **in
	}
	if in.ControllerRegistration != nil {
		in, out := &in.ControllerRegistration, &out.ControllerRegistration
		*out = new(ControllerRegistrationControllerConfiguration)
		**out = **in
	}
	if in.ControllerInstallation != nil {
		in, out := &in.ControllerInstallation, &out.ControllerInstallation
		*out = new(ControllerInstallationControllerConfiguration)
		**out = **in
	}
	if in.Plant != nil {
		in, out := &in.Plant, &out.Plant
		*out = new(PlantConfiguration)
		**out = **in
	}
	if in.SecretBinding != nil {
		in, out := &in.SecretBinding, &out.SecretBinding
		*out = new(SecretBindingControllerConfiguration)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(ProjectControllerConfiguration)
		**out = **in
	}
	if in.Quota != nil {
		in, out := &in.Quota, &out.Quota
		*out = new(QuotaControllerConfiguration)
		**out = **in
	}
	if in.Seed != nil {
		in, out := &in.Seed, &out.Seed
		*out = new(SeedControllerConfiguration)
		(*in).DeepCopyInto(*out)
	}
	in.Shoot.DeepCopyInto(&out.Shoot)
	in.ShootCare.DeepCopyInto(&out.ShootCare)
	out.ShootMaintenance = in.ShootMaintenance
	out.ShootQuota = in.ShootQuota
	out.ShootHibernation = in.ShootHibernation
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManagerControllerConfiguration.
func (in *ControllerManagerControllerConfiguration) DeepCopy() *ControllerManagerControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerManagerControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerRegistrationControllerConfiguration) DeepCopyInto(out *ControllerRegistrationControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerRegistrationControllerConfiguration.
func (in *ControllerRegistrationControllerConfiguration) DeepCopy() *ControllerRegistrationControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerRegistrationControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiscoveryConfiguration) DeepCopyInto(out *DiscoveryConfiguration) {
	*out = *in
	if in.DiscoveryCacheDir != nil {
		in, out := &in.DiscoveryCacheDir, &out.DiscoveryCacheDir
		*out = new(string)
		**out = **in
	}
	if in.HTTPCacheDir != nil {
		in, out := &in.HTTPCacheDir, &out.HTTPCacheDir
		*out = new(string)
		**out = **in
	}
	if in.TTL != nil {
		in, out := &in.TTL, &out.TTL
		*out = new(v1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiscoveryConfiguration.
func (in *DiscoveryConfiguration) DeepCopy() *DiscoveryConfiguration {
	if in == nil {
		return nil
	}
	out := new(DiscoveryConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPSServer) DeepCopyInto(out *HTTPSServer) {
	*out = *in
	out.Server = in.Server
	out.TLS = in.TLS
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPSServer.
func (in *HTTPSServer) DeepCopy() *HTTPSServer {
	if in == nil {
		return nil
	}
	out := new(HTTPSServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderElectionConfiguration) DeepCopyInto(out *LeaderElectionConfiguration) {
	*out = *in
	out.LeaderElectionConfiguration = in.LeaderElectionConfiguration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderElectionConfiguration.
func (in *LeaderElectionConfiguration) DeepCopy() *LeaderElectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(LeaderElectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlantConfiguration) DeepCopyInto(out *PlantConfiguration) {
	*out = *in
	out.SyncPeriod = in.SyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlantConfiguration.
func (in *PlantConfiguration) DeepCopy() *PlantConfiguration {
	if in == nil {
		return nil
	}
	out := new(PlantConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectControllerConfiguration) DeepCopyInto(out *ProjectControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectControllerConfiguration.
func (in *ProjectControllerConfiguration) DeepCopy() *ProjectControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ProjectControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QuotaControllerConfiguration) DeepCopyInto(out *QuotaControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QuotaControllerConfiguration.
func (in *QuotaControllerConfiguration) DeepCopy() *QuotaControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(QuotaControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBindingControllerConfiguration) DeepCopyInto(out *SecretBindingControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBindingControllerConfiguration.
func (in *SecretBindingControllerConfiguration) DeepCopy() *SecretBindingControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(SecretBindingControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SeedControllerConfiguration) DeepCopyInto(out *SeedControllerConfiguration) {
	*out = *in
	if in.ReserveExcessCapacity != nil {
		in, out := &in.ReserveExcessCapacity, &out.ReserveExcessCapacity
		*out = new(bool)
		**out = **in
	}
	out.SyncPeriod = in.SyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SeedControllerConfiguration.
func (in *SeedControllerConfiguration) DeepCopy() *SeedControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(SeedControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerConfiguration) DeepCopyInto(out *ServerConfiguration) {
	*out = *in
	out.HTTP = in.HTTP
	out.HTTPS = in.HTTPS
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerConfiguration.
func (in *ServerConfiguration) DeepCopy() *ServerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ServerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootBackup) DeepCopyInto(out *ShootBackup) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootBackup.
func (in *ShootBackup) DeepCopy() *ShootBackup {
	if in == nil {
		return nil
	}
	out := new(ShootBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootCareControllerConfiguration) DeepCopyInto(out *ShootCareControllerConfiguration) {
	*out = *in
	out.SyncPeriod = in.SyncPeriod
	if in.ConditionThresholds != nil {
		in, out := &in.ConditionThresholds, &out.ConditionThresholds
		*out = make([]ConditionThreshold, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootCareControllerConfiguration.
func (in *ShootCareControllerConfiguration) DeepCopy() *ShootCareControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootCareControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootControllerConfiguration) DeepCopyInto(out *ShootControllerConfiguration) {
	*out = *in
	if in.ReconcileInMaintenanceOnly != nil {
		in, out := &in.ReconcileInMaintenanceOnly, &out.ReconcileInMaintenanceOnly
		*out = new(bool)
		**out = **in
	}
	if in.RespectSyncPeriodOverwrite != nil {
		in, out := &in.RespectSyncPeriodOverwrite, &out.RespectSyncPeriodOverwrite
		*out = new(bool)
		**out = **in
	}
	out.RetryDuration = in.RetryDuration
	if in.RetrySyncPeriod != nil {
		in, out := &in.RetrySyncPeriod, &out.RetrySyncPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	out.SyncPeriod = in.SyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootControllerConfiguration.
func (in *ShootControllerConfiguration) DeepCopy() *ShootControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootHibernationControllerConfiguration) DeepCopyInto(out *ShootHibernationControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootHibernationControllerConfiguration.
func (in *ShootHibernationControllerConfiguration) DeepCopy() *ShootHibernationControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootHibernationControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootMaintenanceControllerConfiguration) DeepCopyInto(out *ShootMaintenanceControllerConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootMaintenanceControllerConfiguration.
func (in *ShootMaintenanceControllerConfiguration) DeepCopy() *ShootMaintenanceControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootMaintenanceControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShootQuotaControllerConfiguration) DeepCopyInto(out *ShootQuotaControllerConfiguration) {
	*out = *in
	out.SyncPeriod = in.SyncPeriod
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShootQuotaControllerConfiguration.
func (in *ShootQuotaControllerConfiguration) DeepCopy() *ShootQuotaControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ShootQuotaControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSServer) DeepCopyInto(out *TLSServer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSServer.
func (in *TLSServer) DeepCopy() *TLSServer {
	if in == nil {
		return nil
	}
	out := new(TLSServer)
	in.DeepCopyInto(out)
	return out
}
