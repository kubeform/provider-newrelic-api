//go:build !ignore_autogenerated
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
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dashboard) DeepCopyInto(out *Dashboard) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dashboard.
func (in *Dashboard) DeepCopy() *Dashboard {
	if in == nil {
		return nil
	}
	out := new(Dashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Dashboard) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardList) DeepCopyInto(out *DashboardList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Dashboard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardList.
func (in *DashboardList) DeepCopy() *DashboardList {
	if in == nil {
		return nil
	}
	out := new(DashboardList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DashboardList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardSpec) DeepCopyInto(out *DashboardSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(DashboardSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardSpec.
func (in *DashboardSpec) DeepCopy() *DashboardSpec {
	if in == nil {
		return nil
	}
	out := new(DashboardSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardSpecFilter) DeepCopyInto(out *DashboardSpecFilter) {
	*out = *in
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EventTypes != nil {
		in, out := &in.EventTypes, &out.EventTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardSpecFilter.
func (in *DashboardSpecFilter) DeepCopy() *DashboardSpecFilter {
	if in == nil {
		return nil
	}
	out := new(DashboardSpecFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardSpecResource) DeepCopyInto(out *DashboardSpecResource) {
	*out = *in
	if in.DashboardURL != nil {
		in, out := &in.DashboardURL, &out.DashboardURL
		*out = new(string)
		**out = **in
	}
	if in.Editable != nil {
		in, out := &in.Editable, &out.Editable
		*out = new(string)
		**out = **in
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(DashboardSpecFilter)
		(*in).DeepCopyInto(*out)
	}
	if in.GridColumnCount != nil {
		in, out := &in.GridColumnCount, &out.GridColumnCount
		*out = new(int64)
		**out = **in
	}
	if in.Icon != nil {
		in, out := &in.Icon, &out.Icon
		*out = new(string)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(string)
		**out = **in
	}
	if in.Widget != nil {
		in, out := &in.Widget, &out.Widget
		*out = make([]DashboardSpecWidget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardSpecResource.
func (in *DashboardSpecResource) DeepCopy() *DashboardSpecResource {
	if in == nil {
		return nil
	}
	out := new(DashboardSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardSpecWidget) DeepCopyInto(out *DashboardSpecWidget) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(int64)
		**out = **in
	}
	if in.Column != nil {
		in, out := &in.Column, &out.Column
		*out = new(int64)
		**out = **in
	}
	if in.CompareWith != nil {
		in, out := &in.CompareWith, &out.CompareWith
		*out = make([]DashboardSpecWidgetCompareWith, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DrilldownDashboardID != nil {
		in, out := &in.DrilldownDashboardID, &out.DrilldownDashboardID
		*out = new(int64)
		**out = **in
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(int64)
		**out = **in
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = new(int64)
		**out = **in
	}
	if in.EntityIDS != nil {
		in, out := &in.EntityIDS, &out.EntityIDS
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.Facet != nil {
		in, out := &in.Facet, &out.Facet
		*out = new(string)
		**out = **in
	}
	if in.Height != nil {
		in, out := &in.Height, &out.Height
		*out = new(int64)
		**out = **in
	}
	if in.Limit != nil {
		in, out := &in.Limit, &out.Limit
		*out = new(int64)
		**out = **in
	}
	if in.Metric != nil {
		in, out := &in.Metric, &out.Metric
		*out = make([]DashboardSpecWidgetMetric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Notes != nil {
		in, out := &in.Notes, &out.Notes
		*out = new(string)
		**out = **in
	}
	if in.Nrql != nil {
		in, out := &in.Nrql, &out.Nrql
		*out = new(string)
		**out = **in
	}
	if in.OrderBy != nil {
		in, out := &in.OrderBy, &out.OrderBy
		*out = new(string)
		**out = **in
	}
	if in.RawMetricName != nil {
		in, out := &in.RawMetricName, &out.RawMetricName
		*out = new(string)
		**out = **in
	}
	if in.Row != nil {
		in, out := &in.Row, &out.Row
		*out = new(int64)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(string)
		**out = **in
	}
	if in.ThresholdRed != nil {
		in, out := &in.ThresholdRed, &out.ThresholdRed
		*out = new(float64)
		**out = **in
	}
	if in.ThresholdYellow != nil {
		in, out := &in.ThresholdYellow, &out.ThresholdYellow
		*out = new(float64)
		**out = **in
	}
	if in.Title != nil {
		in, out := &in.Title, &out.Title
		*out = new(string)
		**out = **in
	}
	if in.Visualization != nil {
		in, out := &in.Visualization, &out.Visualization
		*out = new(string)
		**out = **in
	}
	if in.WidgetID != nil {
		in, out := &in.WidgetID, &out.WidgetID
		*out = new(int64)
		**out = **in
	}
	if in.Width != nil {
		in, out := &in.Width, &out.Width
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardSpecWidget.
func (in *DashboardSpecWidget) DeepCopy() *DashboardSpecWidget {
	if in == nil {
		return nil
	}
	out := new(DashboardSpecWidget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardSpecWidgetCompareWith) DeepCopyInto(out *DashboardSpecWidgetCompareWith) {
	*out = *in
	if in.OffsetDuration != nil {
		in, out := &in.OffsetDuration, &out.OffsetDuration
		*out = new(string)
		**out = **in
	}
	if in.Presentation != nil {
		in, out := &in.Presentation, &out.Presentation
		*out = new(DashboardSpecWidgetCompareWithPresentation)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardSpecWidgetCompareWith.
func (in *DashboardSpecWidgetCompareWith) DeepCopy() *DashboardSpecWidgetCompareWith {
	if in == nil {
		return nil
	}
	out := new(DashboardSpecWidgetCompareWith)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardSpecWidgetCompareWithPresentation) DeepCopyInto(out *DashboardSpecWidgetCompareWithPresentation) {
	*out = *in
	if in.Color != nil {
		in, out := &in.Color, &out.Color
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardSpecWidgetCompareWithPresentation.
func (in *DashboardSpecWidgetCompareWithPresentation) DeepCopy() *DashboardSpecWidgetCompareWithPresentation {
	if in == nil {
		return nil
	}
	out := new(DashboardSpecWidgetCompareWithPresentation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardSpecWidgetMetric) DeepCopyInto(out *DashboardSpecWidgetMetric) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.Units != nil {
		in, out := &in.Units, &out.Units
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardSpecWidgetMetric.
func (in *DashboardSpecWidgetMetric) DeepCopy() *DashboardSpecWidgetMetric {
	if in == nil {
		return nil
	}
	out := new(DashboardSpecWidgetMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardStatus) DeepCopyInto(out *DashboardStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardStatus.
func (in *DashboardStatus) DeepCopy() *DashboardStatus {
	if in == nil {
		return nil
	}
	out := new(DashboardStatus)
	in.DeepCopyInto(out)
	return out
}
