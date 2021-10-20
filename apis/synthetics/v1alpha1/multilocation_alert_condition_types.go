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

type MultilocationAlertCondition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MultilocationAlertConditionSpec   `json:"spec,omitempty"`
	Status            MultilocationAlertConditionStatus `json:"status,omitempty"`
}

type MultilocationAlertConditionSpecCritical struct {
	// The minimum number of monitor locations that must be concurrently failing before a violation is opened.
	Threshold *int64 `json:"threshold" tf:"threshold"`
}

type MultilocationAlertConditionSpecWarning struct {
	// The minimum number of monitor locations that must be concurrently failing before a violation is opened.
	Threshold *int64 `json:"threshold" tf:"threshold"`
}

type MultilocationAlertConditionSpec struct {
	State *MultilocationAlertConditionSpecResource `json:"state,omitempty" tf:"-"`

	Resource MultilocationAlertConditionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type MultilocationAlertConditionSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// A condition term with priority set to critical.
	Critical *MultilocationAlertConditionSpecCritical `json:"critical" tf:"critical"`
	// Set whether to enable the alert condition. Defaults to true.
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// The GUIDs of the Synthetics monitors to alert on.
	// +kubebuilder:validation:MinItems=1
	Entities []string `json:"entities" tf:"entities"`
	// The title of this condition.
	Name *string `json:"name" tf:"name"`
	// The ID of the policy where this condition will be used.
	PolicyID *int64 `json:"policyID" tf:"policy_id"`
	// Runbook URL to display in notifications.
	// +optional
	RunbookURL *string `json:"runbookURL,omitempty" tf:"runbook_url"`
	// The maximum number of seconds a violation can remain open before being closed by the system.  Must be one of: 0, 3600, 7200, 14400, 28800, 43200, 86400
	ViolationTimeLimitSeconds *int64 `json:"violationTimeLimitSeconds" tf:"violation_time_limit_seconds"`
	// A condition term with priority set to warning.
	// +optional
	Warning *MultilocationAlertConditionSpecWarning `json:"warning,omitempty" tf:"warning"`
}

type MultilocationAlertConditionStatus struct {
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

// MultilocationAlertConditionList is a list of MultilocationAlertConditions
type MultilocationAlertConditionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MultilocationAlertCondition CRD objects
	Items []MultilocationAlertCondition `json:"items,omitempty"`
}
