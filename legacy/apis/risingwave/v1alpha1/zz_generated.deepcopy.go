//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
 * Copyright 2022 Singularity Data
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseOptions) DeepCopyInto(out *BaseOptions) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = make(map[Arch]ImageOptions, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseOptions.
func (in *BaseOptions) DeepCopy() *BaseOptions {
	if in == nil {
		return nil
	}
	out := new(BaseOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudService) DeepCopyInto(out *CloudService) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudService.
func (in *CloudService) DeepCopy() *CloudService {
	if in == nil {
		return nil
	}
	out := new(CloudService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompactorNodeSpec) DeepCopyInto(out *CompactorNodeSpec) {
	*out = *in
	in.DeployDescriptor.DeepCopyInto(&out.DeployDescriptor)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompactorNodeSpec.
func (in *CompactorNodeSpec) DeepCopy() *CompactorNodeSpec {
	if in == nil {
		return nil
	}
	out := new(CompactorNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompactorNodeStatus) DeepCopyInto(out *CompactorNodeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompactorNodeStatus.
func (in *CompactorNodeStatus) DeepCopy() *CompactorNodeStatus {
	if in == nil {
		return nil
	}
	out := new(CompactorNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeNodeSpec) DeepCopyInto(out *ComputeNodeSpec) {
	*out = *in
	in.DeployDescriptor.DeepCopyInto(&out.DeployDescriptor)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeNodeSpec.
func (in *ComputeNodeSpec) DeepCopy() *ComputeNodeSpec {
	if in == nil {
		return nil
	}
	out := new(ComputeNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComputeNodeStatus) DeepCopyInto(out *ComputeNodeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComputeNodeStatus.
func (in *ComputeNodeStatus) DeepCopy() *ComputeNodeStatus {
	if in == nil {
		return nil
	}
	out := new(ComputeNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployDescriptor) DeepCopyInto(out *DeployDescriptor) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageDescriptor)
		(*in).DeepCopyInto(*out)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]v1.ContainerPort, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CMD != nil {
		in, out := &in.CMD, &out.CMD
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployDescriptor.
func (in *DeployDescriptor) DeepCopy() *DeployDescriptor {
	if in == nil {
		return nil
	}
	out := new(DeployDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendSpec) DeepCopyInto(out *FrontendSpec) {
	*out = *in
	if in.ServiceSpec != nil {
		in, out := &in.ServiceSpec, &out.ServiceSpec
		*out = new(v1.ServiceSpec)
		(*in).DeepCopyInto(*out)
	}
	in.DeployDescriptor.DeepCopyInto(&out.DeployDescriptor)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendSpec.
func (in *FrontendSpec) DeepCopy() *FrontendSpec {
	if in == nil {
		return nil
	}
	out := new(FrontendSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FrontendSpecStatus) DeepCopyInto(out *FrontendSpecStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FrontendSpecStatus.
func (in *FrontendSpecStatus) DeepCopy() *FrontendSpecStatus {
	if in == nil {
		return nil
	}
	out := new(FrontendSpecStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageDescriptor) DeepCopyInto(out *ImageDescriptor) {
	*out = *in
	if in.Repository != nil {
		in, out := &in.Repository, &out.Repository
		*out = new(string)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
	if in.PullPolicy != nil {
		in, out := &in.PullPolicy, &out.PullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageDescriptor.
func (in *ImageDescriptor) DeepCopy() *ImageDescriptor {
	if in == nil {
		return nil
	}
	out := new(ImageDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageOptions) DeepCopyInto(out *ImageOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageOptions.
func (in *ImageOptions) DeepCopy() *ImageOptions {
	if in == nil {
		return nil
	}
	out := new(ImageOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaNodeSpec) DeepCopyInto(out *MetaNodeSpec) {
	*out = *in
	in.DeployDescriptor.DeepCopyInto(&out.DeployDescriptor)
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(MetaStorage)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaNodeSpec.
func (in *MetaNodeSpec) DeepCopy() *MetaNodeSpec {
	if in == nil {
		return nil
	}
	out := new(MetaNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaNodeStatus) DeepCopyInto(out *MetaNodeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaNodeStatus.
func (in *MetaNodeStatus) DeepCopy() *MetaNodeStatus {
	if in == nil {
		return nil
	}
	out := new(MetaNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaStorage) DeepCopyInto(out *MetaStorage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaStorage.
func (in *MetaStorage) DeepCopy() *MetaStorage {
	if in == nil {
		return nil
	}
	out := new(MetaStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinIO) DeepCopyInto(out *MinIO) {
	*out = *in
	in.DeployDescriptor.DeepCopyInto(&out.DeployDescriptor)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinIO.
func (in *MinIO) DeepCopy() *MinIO {
	if in == nil {
		return nil
	}
	out := new(MinIO)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinIOStatus) DeepCopyInto(out *MinIOStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinIOStatus.
func (in *MinIOStatus) DeepCopy() *MinIOStatus {
	if in == nil {
		return nil
	}
	out := new(MinIOStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorageSpec) DeepCopyInto(out *ObjectStorageSpec) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3)
		(*in).DeepCopyInto(*out)
	}
	if in.MinIO != nil {
		in, out := &in.MinIO, &out.MinIO
		*out = new(MinIO)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorageSpec.
func (in *ObjectStorageSpec) DeepCopy() *ObjectStorageSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorageStatus) DeepCopyInto(out *ObjectStorageStatus) {
	*out = *in
	if in.MinIOStatus != nil {
		in, out := &in.MinIOStatus, &out.MinIOStatus
		*out = new(MinIOStatus)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorageStatus.
func (in *ObjectStorageStatus) DeepCopy() *ObjectStorageStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectStorageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RisingWave) DeepCopyInto(out *RisingWave) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RisingWave.
func (in *RisingWave) DeepCopy() *RisingWave {
	if in == nil {
		return nil
	}
	out := new(RisingWave)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RisingWave) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RisingWaveCondition) DeepCopyInto(out *RisingWaveCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RisingWaveCondition.
func (in *RisingWaveCondition) DeepCopy() *RisingWaveCondition {
	if in == nil {
		return nil
	}
	out := new(RisingWaveCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RisingWaveList) DeepCopyInto(out *RisingWaveList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RisingWave, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RisingWaveList.
func (in *RisingWaveList) DeepCopy() *RisingWaveList {
	if in == nil {
		return nil
	}
	out := new(RisingWaveList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RisingWaveList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RisingWaveOptions) DeepCopyInto(out *RisingWaveOptions) {
	*out = *in
	in.MetaNode.DeepCopyInto(&out.MetaNode)
	in.ComputeNode.DeepCopyInto(&out.ComputeNode)
	in.CompactorNode.DeepCopyInto(&out.CompactorNode)
	in.MinIO.DeepCopyInto(&out.MinIO)
	in.Frontend.DeepCopyInto(&out.Frontend)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RisingWaveOptions.
func (in *RisingWaveOptions) DeepCopy() *RisingWaveOptions {
	if in == nil {
		return nil
	}
	out := new(RisingWaveOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RisingWaveSpec) DeepCopyInto(out *RisingWaveSpec) {
	*out = *in
	if in.MetaNode != nil {
		in, out := &in.MetaNode, &out.MetaNode
		*out = new(MetaNodeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ObjectStorage != nil {
		in, out := &in.ObjectStorage, &out.ObjectStorage
		*out = new(ObjectStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ComputeNode != nil {
		in, out := &in.ComputeNode, &out.ComputeNode
		*out = new(ComputeNodeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CompactorNode != nil {
		in, out := &in.CompactorNode, &out.CompactorNode
		*out = new(CompactorNodeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Frontend != nil {
		in, out := &in.Frontend, &out.Frontend
		*out = new(FrontendSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RisingWaveSpec.
func (in *RisingWaveSpec) DeepCopy() *RisingWaveSpec {
	if in == nil {
		return nil
	}
	out := new(RisingWaveSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RisingWaveStatus) DeepCopyInto(out *RisingWaveStatus) {
	*out = *in
	out.MetaNode = in.MetaNode
	in.ObjectStorage.DeepCopyInto(&out.ObjectStorage)
	out.ComputeNode = in.ComputeNode
	out.CompactorNode = in.CompactorNode
	out.Frontend = in.Frontend
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]RisingWaveCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RisingWaveStatus.
func (in *RisingWaveStatus) DeepCopy() *RisingWaveStatus {
	if in == nil {
		return nil
	}
	out := new(RisingWaveStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3) DeepCopyInto(out *S3) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3.
func (in *S3) DeepCopy() *S3 {
	if in == nil {
		return nil
	}
	out := new(S3)
	in.DeepCopyInto(out)
	return out
}