// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package v1

import (
	proto "github.com/gogo/protobuf/proto"
)

// DeepCopyInto for the FilterDeployment.Spec
func (in *FilterDeploymentSpec) DeepCopyInto(out *FilterDeploymentSpec) {
	p := proto.Clone(in).(*FilterDeploymentSpec)
	*out = *p
}

// DeepCopyInto for the FilterDeployment.Status
func (in *FilterDeploymentStatus) DeepCopyInto(out *FilterDeploymentStatus) {
	p := proto.Clone(in).(*FilterDeploymentStatus)
	*out = *p
}
