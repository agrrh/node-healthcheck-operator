// +build !ignore_autogenerated

/*
Copyright 2021.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backoff) DeepCopyInto(out *Backoff) {
	*out = *in
	out.Limit = in.Limit
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backoff.
func (in *Backoff) DeepCopy() *Backoff {
	if in == nil {
		return nil
	}
	out := new(Backoff)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeHealthCheck) DeepCopyInto(out *NodeHealthCheck) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeHealthCheck.
func (in *NodeHealthCheck) DeepCopy() *NodeHealthCheck {
	if in == nil {
		return nil
	}
	out := new(NodeHealthCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeHealthCheck) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeHealthCheckList) DeepCopyInto(out *NodeHealthCheckList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeHealthCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeHealthCheckList.
func (in *NodeHealthCheckList) DeepCopy() *NodeHealthCheckList {
	if in == nil {
		return nil
	}
	out := new(NodeHealthCheckList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeHealthCheckList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeHealthCheckSpec) DeepCopyInto(out *NodeHealthCheckSpec) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
	if in.UnhealthyConditions != nil {
		in, out := &in.UnhealthyConditions, &out.UnhealthyConditions
		*out = make([]UnhealthyCondition, len(*in))
		copy(*out, *in)
	}
	if in.MaxUnhealthy != nil {
		in, out := &in.MaxUnhealthy, &out.MaxUnhealthy
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.RemediationTemplate != nil {
		in, out := &in.RemediationTemplate, &out.RemediationTemplate
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.PauseRequests != nil {
		in, out := &in.PauseRequests, &out.PauseRequests
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Backoff != nil {
		in, out := &in.Backoff, &out.Backoff
		*out = new(Backoff)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeHealthCheckSpec.
func (in *NodeHealthCheckSpec) DeepCopy() *NodeHealthCheckSpec {
	if in == nil {
		return nil
	}
	out := new(NodeHealthCheckSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeHealthCheckStatus) DeepCopyInto(out *NodeHealthCheckStatus) {
	*out = *in
	if in.InFlightRemediations != nil {
		in, out := &in.InFlightRemediations, &out.InFlightRemediations
		*out = make(map[string]metav1.Time, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeHealthCheckStatus.
func (in *NodeHealthCheckStatus) DeepCopy() *NodeHealthCheckStatus {
	if in == nil {
		return nil
	}
	out := new(NodeHealthCheckStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnhealthyCondition) DeepCopyInto(out *UnhealthyCondition) {
	*out = *in
	out.Duration = in.Duration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnhealthyCondition.
func (in *UnhealthyCondition) DeepCopy() *UnhealthyCondition {
	if in == nil {
		return nil
	}
	out := new(UnhealthyCondition)
	in.DeepCopyInto(out)
	return out
}
