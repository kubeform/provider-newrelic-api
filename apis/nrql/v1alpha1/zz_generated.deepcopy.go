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
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdDuration != nil {
		in, out := &in.ThresholdDuration, &out.ThresholdDuration
		*out = new(int64)
		**out = **in
	}
	if in.ThresholdOccurrences != nil {
		in, out := &in.ThresholdOccurrences, &out.ThresholdOccurrences
		*out = new(string)
		**out = **in
	}
	if in.TimeFunction != nil {
		in, out := &in.TimeFunction, &out.TimeFunction
		*out = new(string)
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
func (in *AlertConditionSpecNrql) DeepCopyInto(out *AlertConditionSpecNrql) {
	*out = *in
	if in.EvaluationOffset != nil {
		in, out := &in.EvaluationOffset, &out.EvaluationOffset
		*out = new(int64)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	if in.SinceValue != nil {
		in, out := &in.SinceValue, &out.SinceValue
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertConditionSpecNrql.
func (in *AlertConditionSpecNrql) DeepCopy() *AlertConditionSpecNrql {
	if in == nil {
		return nil
	}
	out := new(AlertConditionSpecNrql)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertConditionSpecResource) DeepCopyInto(out *AlertConditionSpecResource) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(int64)
		**out = **in
	}
	if in.AggregationWindow != nil {
		in, out := &in.AggregationWindow, &out.AggregationWindow
		*out = new(int64)
		**out = **in
	}
	if in.BaselineDirection != nil {
		in, out := &in.BaselineDirection, &out.BaselineDirection
		*out = new(string)
		**out = **in
	}
	if in.CloseViolationsOnExpiration != nil {
		in, out := &in.CloseViolationsOnExpiration, &out.CloseViolationsOnExpiration
		*out = new(bool)
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
	if in.ExpectedGroups != nil {
		in, out := &in.ExpectedGroups, &out.ExpectedGroups
		*out = new(int64)
		**out = **in
	}
	if in.ExpirationDuration != nil {
		in, out := &in.ExpirationDuration, &out.ExpirationDuration
		*out = new(int64)
		**out = **in
	}
	if in.FillOption != nil {
		in, out := &in.FillOption, &out.FillOption
		*out = new(string)
		**out = **in
	}
	if in.FillValue != nil {
		in, out := &in.FillValue, &out.FillValue
		*out = new(float64)
		**out = **in
	}
	if in.IgnoreOverlap != nil {
		in, out := &in.IgnoreOverlap, &out.IgnoreOverlap
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Nrql != nil {
		in, out := &in.Nrql, &out.Nrql
		*out = new(AlertConditionSpecNrql)
		(*in).DeepCopyInto(*out)
	}
	if in.OpenViolationOnExpiration != nil {
		in, out := &in.OpenViolationOnExpiration, &out.OpenViolationOnExpiration
		*out = new(bool)
		**out = **in
	}
	if in.OpenViolationOnGroupOverlap != nil {
		in, out := &in.OpenViolationOnGroupOverlap, &out.OpenViolationOnGroupOverlap
		*out = new(bool)
		**out = **in
	}
	if in.PolicyID != nil {
		in, out := &in.PolicyID, &out.PolicyID
		*out = new(int64)
		**out = **in
	}
	if in.RunbookURL != nil {
		in, out := &in.RunbookURL, &out.RunbookURL
		*out = new(string)
		**out = **in
	}
	if in.Term != nil {
		in, out := &in.Term, &out.Term
		*out = make([]AlertConditionSpecTerm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.ValueFunction != nil {
		in, out := &in.ValueFunction, &out.ValueFunction
		*out = new(string)
		**out = **in
	}
	if in.ViolationTimeLimit != nil {
		in, out := &in.ViolationTimeLimit, &out.ViolationTimeLimit
		*out = new(string)
		**out = **in
	}
	if in.ViolationTimeLimitSeconds != nil {
		in, out := &in.ViolationTimeLimitSeconds, &out.ViolationTimeLimitSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Warning != nil {
		in, out := &in.Warning, &out.Warning
		*out = new(AlertConditionSpecWarning)
		(*in).DeepCopyInto(*out)
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
func (in *AlertConditionSpecTerm) DeepCopyInto(out *AlertConditionSpecTerm) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(int64)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(string)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdDuration != nil {
		in, out := &in.ThresholdDuration, &out.ThresholdDuration
		*out = new(int64)
		**out = **in
	}
	if in.ThresholdOccurrences != nil {
		in, out := &in.ThresholdOccurrences, &out.ThresholdOccurrences
		*out = new(string)
		**out = **in
	}
	if in.TimeFunction != nil {
		in, out := &in.TimeFunction, &out.TimeFunction
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertConditionSpecTerm.
func (in *AlertConditionSpecTerm) DeepCopy() *AlertConditionSpecTerm {
	if in == nil {
		return nil
	}
	out := new(AlertConditionSpecTerm)
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
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Threshold != nil {
		in, out := &in.Threshold, &out.Threshold
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdDuration != nil {
		in, out := &in.ThresholdDuration, &out.ThresholdDuration
		*out = new(int64)
		**out = **in
	}
	if in.ThresholdOccurrences != nil {
		in, out := &in.ThresholdOccurrences, &out.ThresholdOccurrences
		*out = new(string)
		**out = **in
	}
	if in.TimeFunction != nil {
		in, out := &in.TimeFunction, &out.TimeFunction
		*out = new(string)
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DropRule) DeepCopyInto(out *DropRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DropRule.
func (in *DropRule) DeepCopy() *DropRule {
	if in == nil {
		return nil
	}
	out := new(DropRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DropRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DropRuleList) DeepCopyInto(out *DropRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DropRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DropRuleList.
func (in *DropRuleList) DeepCopy() *DropRuleList {
	if in == nil {
		return nil
	}
	out := new(DropRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DropRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DropRuleSpec) DeepCopyInto(out *DropRuleSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(DropRuleSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DropRuleSpec.
func (in *DropRuleSpec) DeepCopy() *DropRuleSpec {
	if in == nil {
		return nil
	}
	out := new(DropRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DropRuleSpecResource) DeepCopyInto(out *DropRuleSpecResource) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(int64)
		**out = **in
	}
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Nrql != nil {
		in, out := &in.Nrql, &out.Nrql
		*out = new(string)
		**out = **in
	}
	if in.RuleID != nil {
		in, out := &in.RuleID, &out.RuleID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DropRuleSpecResource.
func (in *DropRuleSpecResource) DeepCopy() *DropRuleSpecResource {
	if in == nil {
		return nil
	}
	out := new(DropRuleSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DropRuleStatus) DeepCopyInto(out *DropRuleStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DropRuleStatus.
func (in *DropRuleStatus) DeepCopy() *DropRuleStatus {
	if in == nil {
		return nil
	}
	out := new(DropRuleStatus)
	in.DeepCopyInto(out)
	return out
}
