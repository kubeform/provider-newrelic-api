// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

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
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertCondition) DeepCopyInto(out *AlertCondition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertCondition.
func (in *AlertCondition) DeepCopy() *AlertCondition {
	if in == nil {
		return nil
	}
	out := new(AlertCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertCondition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertConditionList) DeepCopyInto(out *AlertConditionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlertCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertConditionList.
func (in *AlertConditionList) DeepCopy() *AlertConditionList {
	if in == nil {
		return nil
	}
	out := new(AlertConditionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertConditionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertConditionSpec) DeepCopyInto(out *AlertConditionSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AlertConditionSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertConditionSpec.
func (in *AlertConditionSpec) DeepCopy() *AlertConditionSpec {
	if in == nil {
		return nil
	}
	out := new(AlertConditionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertConditionSpecCritical) DeepCopyInto(out *AlertConditionSpecCritical) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(int64)
		**out = **in
	}
	if in.TimeFunction != nil {
		in, out := &in.TimeFunction, &out.TimeFunction
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertConditionSpecCritical.
func (in *AlertConditionSpecCritical) DeepCopy() *AlertConditionSpecCritical {
	if in == nil {
		return nil
	}
	out := new(AlertConditionSpecCritical)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertConditionSpecResource) DeepCopyInto(out *AlertConditionSpecResource) {
	*out = *in
	if in.Comparison != nil {
		in, out := &in.Comparison, &out.Comparison
		*out = new(string)
		**out = **in
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(int64)
		**out = **in
	}
	if in.Critical != nil {
		in, out := &in.Critical, &out.Critical
		*out = new(AlertConditionSpecCritical)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Event != nil {
		in, out := &in.Event, &out.Event
		*out = new(string)
		**out = **in
	}
	if in.IntegrationProvider != nil {
		in, out := &in.IntegrationProvider, &out.IntegrationProvider
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PolicyID != nil {
		in, out := &in.PolicyID, &out.PolicyID
		*out = new(int64)
		**out = **in
	}
	if in.ProcessWhere != nil {
		in, out := &in.ProcessWhere, &out.ProcessWhere
		*out = new(string)
		**out = **in
	}
	if in.RunbookURL != nil {
		in, out := &in.RunbookURL, &out.RunbookURL
		*out = new(string)
		**out = **in
	}
	if in.Select != nil {
		in, out := &in.Select, &out.Select
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UpdatedAt != nil {
		in, out := &in.UpdatedAt, &out.UpdatedAt
		*out = new(int64)
		**out = **in
	}
	if in.ViolationCloseTimer != nil {
		in, out := &in.ViolationCloseTimer, &out.ViolationCloseTimer
		*out = new(int64)
		**out = **in
	}
	if in.Warning != nil {
		in, out := &in.Warning, &out.Warning
		*out = new(AlertConditionSpecWarning)
		(*in).DeepCopyInto(*out)
	}
	if in.Where != nil {
		in, out := &in.Where, &out.Where
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertConditionSpecResource.
func (in *AlertConditionSpecResource) DeepCopy() *AlertConditionSpecResource {
	if in == nil {
		return nil
	}
	out := new(AlertConditionSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertConditionSpecWarning) DeepCopyInto(out *AlertConditionSpecWarning) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(int64)
		**out = **in
	}
	if in.TimeFunction != nil {
		in, out := &in.TimeFunction, &out.TimeFunction
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertConditionSpecWarning.
func (in *AlertConditionSpecWarning) DeepCopy() *AlertConditionSpecWarning {
	if in == nil {
		return nil
	}
	out := new(AlertConditionSpecWarning)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertConditionStatus) DeepCopyInto(out *AlertConditionStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertConditionStatus.
func (in *AlertConditionStatus) DeepCopy() *AlertConditionStatus {
	if in == nil {
		return nil
	}
	out := new(AlertConditionStatus)
	in.DeepCopyInto(out)
	return out
}
