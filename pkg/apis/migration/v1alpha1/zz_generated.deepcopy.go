// +build !ignore_autogenerated

/*
Copyright 2019 Red Hat Inc.

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
// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupStorageConfig) DeepCopyInto(out *BackupStorageConfig) {
	*out = *in
	if in.CredsSecretRef != nil {
		in, out := &in.CredsSecretRef, &out.CredsSecretRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.S3CustomCABundle != nil {
		in, out := &in.S3CustomCABundle, &out.S3CustomCABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupStorageConfig.
func (in *BackupStorageConfig) DeepCopy() *BackupStorageConfig {
	if in == nil {
		return nil
	}
	out := new(BackupStorageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Conditions) DeepCopyInto(out *Conditions) {
	*out = *in
	if in.List != nil {
		in, out := &in.List, &out.List
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conditions.
func (in *Conditions) DeepCopy() *Conditions {
	if in == nil {
		return nil
	}
	out := new(Conditions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Incompatible) DeepCopyInto(out *Incompatible) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]IncompatibleNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Incompatible.
func (in *Incompatible) DeepCopy() *Incompatible {
	if in == nil {
		return nil
	}
	out := new(Incompatible)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncompatibleGVK) DeepCopyInto(out *IncompatibleGVK) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncompatibleGVK.
func (in *IncompatibleGVK) DeepCopy() *IncompatibleGVK {
	if in == nil {
		return nil
	}
	out := new(IncompatibleGVK)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncompatibleNamespace) DeepCopyInto(out *IncompatibleNamespace) {
	*out = *in
	if in.GVKs != nil {
		in, out := &in.GVKs, &out.GVKs
		*out = make([]IncompatibleGVK, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncompatibleNamespace.
func (in *IncompatibleNamespace) DeepCopy() *IncompatibleNamespace {
	if in == nil {
		return nil
	}
	out := new(IncompatibleNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigAnalytic) DeepCopyInto(out *MigAnalytic) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigAnalytic.
func (in *MigAnalytic) DeepCopy() *MigAnalytic {
	if in == nil {
		return nil
	}
	out := new(MigAnalytic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigAnalytic) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigAnalyticList) DeepCopyInto(out *MigAnalyticList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigAnalytic, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigAnalyticList.
func (in *MigAnalyticList) DeepCopy() *MigAnalyticList {
	if in == nil {
		return nil
	}
	out := new(MigAnalyticList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigAnalyticList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigAnalyticNSImage) DeepCopyInto(out *MigAnalyticNSImage) {
	*out = *in
	out.Size = in.Size.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigAnalyticNSImage.
func (in *MigAnalyticNSImage) DeepCopy() *MigAnalyticNSImage {
	if in == nil {
		return nil
	}
	out := new(MigAnalyticNSImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigAnalyticNSResource) DeepCopyInto(out *MigAnalyticNSResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigAnalyticNSResource.
func (in *MigAnalyticNSResource) DeepCopy() *MigAnalyticNSResource {
	if in == nil {
		return nil
	}
	out := new(MigAnalyticNSResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigAnalyticNamespace) DeepCopyInto(out *MigAnalyticNamespace) {
	*out = *in
	out.PVCapacity = in.PVCapacity.DeepCopy()
	out.ImageSizeTotal = in.ImageSizeTotal.DeepCopy()
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]MigAnalyticNSImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.K8SResources != nil {
		in, out := &in.K8SResources, &out.K8SResources
		*out = make([]MigAnalyticNSResource, len(*in))
		copy(*out, *in)
	}
	if in.ExcludedK8SResources != nil {
		in, out := &in.ExcludedK8SResources, &out.ExcludedK8SResources
		*out = make([]MigAnalyticNSResource, len(*in))
		copy(*out, *in)
	}
	if in.IncompatibleK8SResources != nil {
		in, out := &in.IncompatibleK8SResources, &out.IncompatibleK8SResources
		*out = make([]MigAnalyticNSResource, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigAnalyticNamespace.
func (in *MigAnalyticNamespace) DeepCopy() *MigAnalyticNamespace {
	if in == nil {
		return nil
	}
	out := new(MigAnalyticNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigAnalyticPlan) DeepCopyInto(out *MigAnalyticPlan) {
	*out = *in
	out.PVCapacity = in.PVCapacity.DeepCopy()
	out.ImageSizeTotal = in.ImageSizeTotal.DeepCopy()
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]MigAnalyticNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigAnalyticPlan.
func (in *MigAnalyticPlan) DeepCopy() *MigAnalyticPlan {
	if in == nil {
		return nil
	}
	out := new(MigAnalyticPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigAnalyticSpec) DeepCopyInto(out *MigAnalyticSpec) {
	*out = *in
	if in.MigPlanRef != nil {
		in, out := &in.MigPlanRef, &out.MigPlanRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigAnalyticSpec.
func (in *MigAnalyticSpec) DeepCopy() *MigAnalyticSpec {
	if in == nil {
		return nil
	}
	out := new(MigAnalyticSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigAnalyticStatus) DeepCopyInto(out *MigAnalyticStatus) {
	*out = *in
	in.Conditions.DeepCopyInto(&out.Conditions)
	in.Analytics.DeepCopyInto(&out.Analytics)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigAnalyticStatus.
func (in *MigAnalyticStatus) DeepCopy() *MigAnalyticStatus {
	if in == nil {
		return nil
	}
	out := new(MigAnalyticStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigCluster) DeepCopyInto(out *MigCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigCluster.
func (in *MigCluster) DeepCopy() *MigCluster {
	if in == nil {
		return nil
	}
	out := new(MigCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigClusterList) DeepCopyInto(out *MigClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigClusterList.
func (in *MigClusterList) DeepCopy() *MigClusterList {
	if in == nil {
		return nil
	}
	out := new(MigClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigClusterSpec) DeepCopyInto(out *MigClusterSpec) {
	*out = *in
	if in.ServiceAccountSecretRef != nil {
		in, out := &in.ServiceAccountSecretRef, &out.ServiceAccountSecretRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.CABundle != nil {
		in, out := &in.CABundle, &out.CABundle
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.StorageClasses != nil {
		in, out := &in.StorageClasses, &out.StorageClasses
		*out = make([]StorageClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RestartRestic != nil {
		in, out := &in.RestartRestic, &out.RestartRestic
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigClusterSpec.
func (in *MigClusterSpec) DeepCopy() *MigClusterSpec {
	if in == nil {
		return nil
	}
	out := new(MigClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigClusterStatus) DeepCopyInto(out *MigClusterStatus) {
	*out = *in
	in.Conditions.DeepCopyInto(&out.Conditions)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigClusterStatus.
func (in *MigClusterStatus) DeepCopy() *MigClusterStatus {
	if in == nil {
		return nil
	}
	out := new(MigClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigDirect) DeepCopyInto(out *MigDirect) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigDirect.
func (in *MigDirect) DeepCopy() *MigDirect {
	if in == nil {
		return nil
	}
	out := new(MigDirect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigDirect) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigDirectList) DeepCopyInto(out *MigDirectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigDirect, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigDirectList.
func (in *MigDirectList) DeepCopy() *MigDirectList {
	if in == nil {
		return nil
	}
	out := new(MigDirectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigDirectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigDirectSpec) DeepCopyInto(out *MigDirectSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigDirectSpec.
func (in *MigDirectSpec) DeepCopy() *MigDirectSpec {
	if in == nil {
		return nil
	}
	out := new(MigDirectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigDirectStatus) DeepCopyInto(out *MigDirectStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigDirectStatus.
func (in *MigDirectStatus) DeepCopy() *MigDirectStatus {
	if in == nil {
		return nil
	}
	out := new(MigDirectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigHook) DeepCopyInto(out *MigHook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigHook.
func (in *MigHook) DeepCopy() *MigHook {
	if in == nil {
		return nil
	}
	out := new(MigHook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigHook) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigHookList) DeepCopyInto(out *MigHookList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigHook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigHookList.
func (in *MigHookList) DeepCopy() *MigHookList {
	if in == nil {
		return nil
	}
	out := new(MigHookList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigHookList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigHookSpec) DeepCopyInto(out *MigHookSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigHookSpec.
func (in *MigHookSpec) DeepCopy() *MigHookSpec {
	if in == nil {
		return nil
	}
	out := new(MigHookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigHookStatus) DeepCopyInto(out *MigHookStatus) {
	*out = *in
	in.Conditions.DeepCopyInto(&out.Conditions)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigHookStatus.
func (in *MigHookStatus) DeepCopy() *MigHookStatus {
	if in == nil {
		return nil
	}
	out := new(MigHookStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigMigration) DeepCopyInto(out *MigMigration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigMigration.
func (in *MigMigration) DeepCopy() *MigMigration {
	if in == nil {
		return nil
	}
	out := new(MigMigration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigMigration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigMigrationList) DeepCopyInto(out *MigMigrationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigMigration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigMigrationList.
func (in *MigMigrationList) DeepCopy() *MigMigrationList {
	if in == nil {
		return nil
	}
	out := new(MigMigrationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigMigrationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigMigrationSpec) DeepCopyInto(out *MigMigrationSpec) {
	*out = *in
	if in.MigPlanRef != nil {
		in, out := &in.MigPlanRef, &out.MigPlanRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigMigrationSpec.
func (in *MigMigrationSpec) DeepCopy() *MigMigrationSpec {
	if in == nil {
		return nil
	}
	out := new(MigMigrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigMigrationStatus) DeepCopyInto(out *MigMigrationStatus) {
	*out = *in
	in.Conditions.DeepCopyInto(&out.Conditions)
	in.UnhealthyResources.DeepCopyInto(&out.UnhealthyResources)
	if in.StartTimestamp != nil {
		in, out := &in.StartTimestamp, &out.StartTimestamp
		*out = (*in).DeepCopy()
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigMigrationStatus.
func (in *MigMigrationStatus) DeepCopy() *MigMigrationStatus {
	if in == nil {
		return nil
	}
	out := new(MigMigrationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigPlan) DeepCopyInto(out *MigPlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigPlan.
func (in *MigPlan) DeepCopy() *MigPlan {
	if in == nil {
		return nil
	}
	out := new(MigPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigPlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigPlanHook) DeepCopyInto(out *MigPlanHook) {
	*out = *in
	if in.Reference != nil {
		in, out := &in.Reference, &out.Reference
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigPlanHook.
func (in *MigPlanHook) DeepCopy() *MigPlanHook {
	if in == nil {
		return nil
	}
	out := new(MigPlanHook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigPlanList) DeepCopyInto(out *MigPlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigPlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigPlanList.
func (in *MigPlanList) DeepCopy() *MigPlanList {
	if in == nil {
		return nil
	}
	out := new(MigPlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigPlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigPlanSpec) DeepCopyInto(out *MigPlanSpec) {
	*out = *in
	in.PersistentVolumes.DeepCopyInto(&out.PersistentVolumes)
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SrcMigClusterRef != nil {
		in, out := &in.SrcMigClusterRef, &out.SrcMigClusterRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.DestMigClusterRef != nil {
		in, out := &in.DestMigClusterRef, &out.DestMigClusterRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.MigStorageRef != nil {
		in, out := &in.MigStorageRef, &out.MigStorageRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.Hooks != nil {
		in, out := &in.Hooks, &out.Hooks
		*out = make([]MigPlanHook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigPlanSpec.
func (in *MigPlanSpec) DeepCopy() *MigPlanSpec {
	if in == nil {
		return nil
	}
	out := new(MigPlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigPlanStatus) DeepCopyInto(out *MigPlanStatus) {
	*out = *in
	in.UnhealthyResources.DeepCopyInto(&out.UnhealthyResources)
	in.Conditions.DeepCopyInto(&out.Conditions)
	in.Incompatible.DeepCopyInto(&out.Incompatible)
	if in.ExcludedResources != nil {
		in, out := &in.ExcludedResources, &out.ExcludedResources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigPlanStatus.
func (in *MigPlanStatus) DeepCopy() *MigPlanStatus {
	if in == nil {
		return nil
	}
	out := new(MigPlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigStorage) DeepCopyInto(out *MigStorage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigStorage.
func (in *MigStorage) DeepCopy() *MigStorage {
	if in == nil {
		return nil
	}
	out := new(MigStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigStorage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigStorageList) DeepCopyInto(out *MigStorageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MigStorage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigStorageList.
func (in *MigStorageList) DeepCopy() *MigStorageList {
	if in == nil {
		return nil
	}
	out := new(MigStorageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MigStorageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigStorageSpec) DeepCopyInto(out *MigStorageSpec) {
	*out = *in
	in.BackupStorageConfig.DeepCopyInto(&out.BackupStorageConfig)
	in.VolumeSnapshotConfig.DeepCopyInto(&out.VolumeSnapshotConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigStorageSpec.
func (in *MigStorageSpec) DeepCopy() *MigStorageSpec {
	if in == nil {
		return nil
	}
	out := new(MigStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MigStorageStatus) DeepCopyInto(out *MigStorageStatus) {
	*out = *in
	in.Conditions.DeepCopyInto(&out.Conditions)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MigStorageStatus.
func (in *MigStorageStatus) DeepCopy() *MigStorageStatus {
	if in == nil {
		return nil
	}
	out := new(MigStorageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PV) DeepCopyInto(out *PV) {
	*out = *in
	out.Capacity = in.Capacity.DeepCopy()
	in.Supported.DeepCopyInto(&out.Supported)
	out.Selection = in.Selection
	in.PVC.DeepCopyInto(&out.PVC)
	if in.NFS != nil {
		in, out := &in.NFS, &out.NFS
		*out = new(v1.NFSVolumeSource)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PV.
func (in *PV) DeepCopy() *PV {
	if in == nil {
		return nil
	}
	out := new(PV)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PVC) DeepCopyInto(out *PVC) {
	*out = *in
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]v1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PVC.
func (in *PVC) DeepCopy() *PVC {
	if in == nil {
		return nil
	}
	out := new(PVC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumes) DeepCopyInto(out *PersistentVolumes) {
	*out = *in
	if in.List != nil {
		in, out := &in.List, &out.List
		*out = make([]PV, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.index != nil {
		in, out := &in.index, &out.index
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumes.
func (in *PersistentVolumes) DeepCopy() *PersistentVolumes {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanResources) DeepCopyInto(out *PlanResources) {
	*out = *in
	if in.MigPlan != nil {
		in, out := &in.MigPlan, &out.MigPlan
		*out = new(MigPlan)
		(*in).DeepCopyInto(*out)
	}
	if in.MigStorage != nil {
		in, out := &in.MigStorage, &out.MigStorage
		*out = new(MigStorage)
		(*in).DeepCopyInto(*out)
	}
	if in.SrcMigCluster != nil {
		in, out := &in.SrcMigCluster, &out.SrcMigCluster
		*out = new(MigCluster)
		(*in).DeepCopyInto(*out)
	}
	if in.DestMigCluster != nil {
		in, out := &in.DestMigCluster, &out.DestMigCluster
		*out = new(MigCluster)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanResources.
func (in *PlanResources) DeepCopy() *PlanResources {
	if in == nil {
		return nil
	}
	out := new(PlanResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Selection) DeepCopyInto(out *Selection) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Selection.
func (in *Selection) DeepCopy() *Selection {
	if in == nil {
		return nil
	}
	out := new(Selection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClass) DeepCopyInto(out *StorageClass) {
	*out = *in
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]v1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClass.
func (in *StorageClass) DeepCopy() *StorageClass {
	if in == nil {
		return nil
	}
	out := new(StorageClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Supported) DeepCopyInto(out *Supported) {
	*out = *in
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CopyMethods != nil {
		in, out := &in.CopyMethods, &out.CopyMethods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Supported.
func (in *Supported) DeepCopy() *Supported {
	if in == nil {
		return nil
	}
	out := new(Supported)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnhealthyNamespace) DeepCopyInto(out *UnhealthyNamespace) {
	*out = *in
	if in.Workloads != nil {
		in, out := &in.Workloads, &out.Workloads
		*out = make([]Workload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnhealthyNamespace.
func (in *UnhealthyNamespace) DeepCopy() *UnhealthyNamespace {
	if in == nil {
		return nil
	}
	out := new(UnhealthyNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnhealthyResources) DeepCopyInto(out *UnhealthyResources) {
	*out = *in
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make([]UnhealthyNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnhealthyResources.
func (in *UnhealthyResources) DeepCopy() *UnhealthyResources {
	if in == nil {
		return nil
	}
	out := new(UnhealthyResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeSnapshotConfig) DeepCopyInto(out *VolumeSnapshotConfig) {
	*out = *in
	if in.CredsSecretRef != nil {
		in, out := &in.CredsSecretRef, &out.CredsSecretRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeSnapshotConfig.
func (in *VolumeSnapshotConfig) DeepCopy() *VolumeSnapshotConfig {
	if in == nil {
		return nil
	}
	out := new(VolumeSnapshotConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workload) DeepCopyInto(out *Workload) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workload.
func (in *Workload) DeepCopy() *Workload {
	if in == nil {
		return nil
	}
	out := new(Workload)
	in.DeepCopyInto(out)
	return out
}
