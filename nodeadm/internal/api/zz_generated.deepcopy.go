//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package api

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDetails) DeepCopyInto(out *ClusterDetails) {
	*out = *in
	if in.CertificateAuthority != nil {
		in, out := &in.CertificateAuthority, &out.CertificateAuthority
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.EnableOutpost != nil {
		in, out := &in.EnableOutpost, &out.EnableOutpost
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDetails.
func (in *ClusterDetails) DeepCopy() *ClusterDetails {
	if in == nil {
		return nil
	}
	out := new(ClusterDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceDetails) DeepCopyInto(out *InstanceDetails) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceDetails.
func (in *InstanceDetails) DeepCopy() *InstanceDetails {
	if in == nil {
		return nil
	}
	out := new(InstanceDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfig) DeepCopyInto(out *NodeConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfig.
func (in *NodeConfig) DeepCopy() *NodeConfig {
	if in == nil {
		return nil
	}
	out := new(NodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfigList) DeepCopyInto(out *NodeConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfigList.
func (in *NodeConfigList) DeepCopy() *NodeConfigList {
	if in == nil {
		return nil
	}
	out := new(NodeConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfigSpec) DeepCopyInto(out *NodeConfigSpec) {
	*out = *in
	in.Cluster.DeepCopyInto(&out.Cluster)
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfigSpec.
func (in *NodeConfigSpec) DeepCopy() *NodeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(NodeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfigStatus) DeepCopyInto(out *NodeConfigStatus) {
	*out = *in
	out.Instance = in.Instance
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfigStatus.
func (in *NodeConfigStatus) DeepCopy() *NodeConfigStatus {
	if in == nil {
		return nil
	}
	out := new(NodeConfigStatus)
	in.DeepCopyInto(out)
	return out
}
