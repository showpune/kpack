//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Binding) DeepCopyInto(out *Binding) {
	*out = *in
	if in.MetadataRef != nil {
		in, out := &in.MetadataRef, &out.MetadataRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Binding.
func (in *Binding) DeepCopy() *Binding {
	if in == nil {
		return nil
	}
	out := new(Binding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Bindings) DeepCopyInto(out *Bindings) {
	{
		in := &in
		*out = make(Bindings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bindings.
func (in Bindings) DeepCopy() Bindings {
	if in == nil {
		return nil
	}
	out := new(Bindings)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Blob) DeepCopyInto(out *Blob) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Blob.
func (in *Blob) DeepCopy() *Blob {
	if in == nil {
		return nil
	}
	out := new(Blob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildBuilderSpec) DeepCopyInto(out *BuildBuilderSpec) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildBuilderSpec.
func (in *BuildBuilderSpec) DeepCopy() *BuildBuilderSpec {
	if in == nil {
		return nil
	}
	out := new(BuildBuilderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildStack) DeepCopyInto(out *BuildStack) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildStack.
func (in *BuildStack) DeepCopy() *BuildStack {
	if in == nil {
		return nil
	}
	out := new(BuildStack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildpackInfo) DeepCopyInto(out *BuildpackInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildpackInfo.
func (in *BuildpackInfo) DeepCopy() *BuildpackInfo {
	if in == nil {
		return nil
	}
	out := new(BuildpackInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildpackMetadata) DeepCopyInto(out *BuildpackMetadata) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildpackMetadata.
func (in *BuildpackMetadata) DeepCopy() *BuildpackMetadata {
	if in == nil {
		return nil
	}
	out := new(BuildpackMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildpackRef) DeepCopyInto(out *BuildpackRef) {
	*out = *in
	out.BuildpackInfo = in.BuildpackInfo
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildpackRef.
func (in *BuildpackRef) DeepCopy() *BuildpackRef {
	if in == nil {
		return nil
	}
	out := new(BuildpackRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildpackStack) DeepCopyInto(out *BuildpackStack) {
	*out = *in
	if in.Mixins != nil {
		in, out := &in.Mixins, &out.Mixins
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildpackStack.
func (in *BuildpackStack) DeepCopy() *BuildpackStack {
	if in == nil {
		return nil
	}
	out := new(BuildpackStack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildpackageInfo) DeepCopyInto(out *BuildpackageInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildpackageInfo.
func (in *BuildpackageInfo) DeepCopy() *BuildpackageInfo {
	if in == nil {
		return nil
	}
	out := new(BuildpackageInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
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
func (in *Git) DeepCopyInto(out *Git) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Git.
func (in *Git) DeepCopy() *Git {
	if in == nil {
		return nil
	}
	out := new(Git)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageBuild) DeepCopyInto(out *ImageBuild) {
	*out = *in
	if in.Bindings != nil {
		in, out := &in.Bindings, &out.Bindings
		*out = make(Bindings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageBuild.
func (in *ImageBuild) DeepCopy() *ImageBuild {
	if in == nil {
		return nil
	}
	out := new(ImageBuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotaryConfig) DeepCopyInto(out *NotaryConfig) {
	*out = *in
	if in.V1 != nil {
		in, out := &in.V1, &out.V1
		*out = new(NotaryV1Config)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotaryConfig.
func (in *NotaryConfig) DeepCopy() *NotaryConfig {
	if in == nil {
		return nil
	}
	out := new(NotaryConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotarySecretRef) DeepCopyInto(out *NotarySecretRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotarySecretRef.
func (in *NotarySecretRef) DeepCopy() *NotarySecretRef {
	if in == nil {
		return nil
	}
	out := new(NotarySecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NotaryV1Config) DeepCopyInto(out *NotaryV1Config) {
	*out = *in
	out.SecretRef = in.SecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NotaryV1Config.
func (in *NotaryV1Config) DeepCopy() *NotaryV1Config {
	if in == nil {
		return nil
	}
	out := new(NotaryV1Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderEntry) DeepCopyInto(out *OrderEntry) {
	*out = *in
	if in.Group != nil {
		in, out := &in.Group, &out.Group
		*out = make([]BuildpackRef, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderEntry.
func (in *OrderEntry) DeepCopy() *OrderEntry {
	if in == nil {
		return nil
	}
	out := new(OrderEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolvedBlobSource) DeepCopyInto(out *ResolvedBlobSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolvedBlobSource.
func (in *ResolvedBlobSource) DeepCopy() *ResolvedBlobSource {
	if in == nil {
		return nil
	}
	out := new(ResolvedBlobSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolvedGitSource) DeepCopyInto(out *ResolvedGitSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolvedGitSource.
func (in *ResolvedGitSource) DeepCopy() *ResolvedGitSource {
	if in == nil {
		return nil
	}
	out := new(ResolvedGitSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolvedRegistrySource) DeepCopyInto(out *ResolvedRegistrySource) {
	*out = *in
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolvedRegistrySource.
func (in *ResolvedRegistrySource) DeepCopy() *ResolvedRegistrySource {
	if in == nil {
		return nil
	}
	out := new(ResolvedRegistrySource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolvedSourceConfig) DeepCopyInto(out *ResolvedSourceConfig) {
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(ResolvedGitSource)
		**out = **in
	}
	if in.Blob != nil {
		in, out := &in.Blob, &out.Blob
		*out = new(ResolvedBlobSource)
		**out = **in
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(ResolvedRegistrySource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolvedSourceConfig.
func (in *ResolvedSourceConfig) DeepCopy() *ResolvedSourceConfig {
	if in == nil {
		return nil
	}
	out := new(ResolvedSourceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceConfig) DeepCopyInto(out *SourceConfig) {
	*out = *in
	if in.Git != nil {
		in, out := &in.Git, &out.Git
		*out = new(Git)
		**out = **in
	}
	if in.Blob != nil {
		in, out := &in.Blob, &out.Blob
		*out = new(Blob)
		**out = **in
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(Registry)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceConfig.
func (in *SourceConfig) DeepCopy() *SourceConfig {
	if in == nil {
		return nil
	}
	out := new(SourceConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Status) DeepCopyInto(out *Status) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Status.
func (in *Status) DeepCopy() *Status {
	if in == nil {
		return nil
	}
	out := new(Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreBuildpack) DeepCopyInto(out *StoreBuildpack) {
	*out = *in
	out.BuildpackInfo = in.BuildpackInfo
	out.Buildpackage = in.Buildpackage
	out.StoreImage = in.StoreImage
	if in.Order != nil {
		in, out := &in.Order, &out.Order
		*out = make([]OrderEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Stacks != nil {
		in, out := &in.Stacks, &out.Stacks
		*out = make([]BuildpackStack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreBuildpack.
func (in *StoreBuildpack) DeepCopy() *StoreBuildpack {
	if in == nil {
		return nil
	}
	out := new(StoreBuildpack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreImage) DeepCopyInto(out *StoreImage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreImage.
func (in *StoreImage) DeepCopy() *StoreImage {
	if in == nil {
		return nil
	}
	out := new(StoreImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolatileTime) DeepCopyInto(out *VolatileTime) {
	*out = *in
	in.Inner.DeepCopyInto(&out.Inner)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolatileTime.
func (in *VolatileTime) DeepCopy() *VolatileTime {
	if in == nil {
		return nil
	}
	out := new(VolatileTime)
	in.DeepCopyInto(out)
	return out
}
