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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Event struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventSpec   `json:"spec,omitempty"`
	Status            EventStatus `json:"status,omitempty"`
}

type EventSpecEventAttribute struct {
	// The name of the attribute.
	Key *string `json:"key" tf:"key"`
	// Specify the type for the attribute value. This is useful when passing integer or float values to Insights. Allowed values are string, int, or float. Defaults to string.
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
	// The value of the attribute.
	Value *string `json:"value" tf:"value"`
}

type EventSpecEvent struct {
	// An attribute to include in your event payload. Multiple attribute blocks can be defined for an event.
	// +kubebuilder:validation:MaxItems=255
	// +kubebuilder:validation:MinItems=1
	Attribute []EventSpecEventAttribute `json:"attribute" tf:"attribute"`
	// Must be a Unix epoch timestamp. You can define timestamps either in seconds or in milliseconds.
	// +optional
	Timestamp *int64 `json:"timestamp,omitempty" tf:"timestamp"`
	// The event's name. Can be a combination of alphanumeric characters, underscores, and colons.
	Type *string `json:"type" tf:"type"`
}

type EventSpec struct {
	State *EventSpecResource `json:"state,omitempty" tf:"-"`

	Resource EventSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type EventSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:MinItems=1
	Event []EventSpecEvent `json:"event" tf:"event"`
}

type EventStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EventList is a list of Events
type EventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Event CRD objects
	Items []Event `json:"items,omitempty"`
}
