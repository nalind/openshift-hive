//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package nutanix

import (
	v1 "github.com/openshift/api/machine/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailureDomain) DeepCopyInto(out *FailureDomain) {
	*out = *in
	out.PrismElement = in.PrismElement
	if in.SubnetUUIDs != nil {
		in, out := &in.SubnetUUIDs, &out.SubnetUUIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.StorageContainers != nil {
		in, out := &in.StorageContainers, &out.StorageContainers
		*out = make([]StorageResourceReference, len(*in))
		copy(*out, *in)
	}
	if in.DataSourceImages != nil {
		in, out := &in.DataSourceImages, &out.DataSourceImages
		*out = make([]StorageResourceReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailureDomain.
func (in *FailureDomain) DeepCopy() *FailureDomain {
	if in == nil {
		return nil
	}
	out := new(FailureDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachinePool) DeepCopyInto(out *MachinePool) {
	*out = *in
	out.OSDisk = in.OSDisk
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(v1.NutanixResourceIdentifier)
		(*in).DeepCopyInto(*out)
	}
	if in.Categories != nil {
		in, out := &in.Categories, &out.Categories
		*out = make([]v1.NutanixCategory, len(*in))
		copy(*out, *in)
	}
	if in.GPUs != nil {
		in, out := &in.GPUs, &out.GPUs
		*out = make([]v1.NutanixGPU, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DataDisks != nil {
		in, out := &in.DataDisks, &out.DataDisks
		*out = make([]v1.NutanixVMDisk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachinePool.
func (in *MachinePool) DeepCopy() *MachinePool {
	if in == nil {
		return nil
	}
	out := new(MachinePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OSDisk) DeepCopyInto(out *OSDisk) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OSDisk.
func (in *OSDisk) DeepCopy() *OSDisk {
	if in == nil {
		return nil
	}
	out := new(OSDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Platform) DeepCopyInto(out *Platform) {
	*out = *in
	out.PrismCentral = in.PrismCentral
	out.CredentialsSecretRef = in.CredentialsSecretRef
	out.CertificatesSecretRef = in.CertificatesSecretRef
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make([]FailureDomain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Platform.
func (in *Platform) DeepCopy() *Platform {
	if in == nil {
		return nil
	}
	out := new(Platform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrismElement) DeepCopyInto(out *PrismElement) {
	*out = *in
	out.Endpoint = in.Endpoint
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrismElement.
func (in *PrismElement) DeepCopy() *PrismElement {
	if in == nil {
		return nil
	}
	out := new(PrismElement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrismEndpoint) DeepCopyInto(out *PrismEndpoint) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrismEndpoint.
func (in *PrismEndpoint) DeepCopy() *PrismEndpoint {
	if in == nil {
		return nil
	}
	out := new(PrismEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageResourceReference) DeepCopyInto(out *StorageResourceReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageResourceReference.
func (in *StorageResourceReference) DeepCopy() *StorageResourceReference {
	if in == nil {
		return nil
	}
	out := new(StorageResourceReference)
	in.DeepCopyInto(out)
	return out
}
