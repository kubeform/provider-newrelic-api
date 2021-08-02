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

type Monitor struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorSpec   `json:"spec,omitempty"`
	Status            MonitorStatus `json:"status,omitempty"`
}

type MonitorSpec struct {
	State *MonitorSpecResource `json:"state,omitempty" tf:"-"`

	Resource MonitorSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type MonitorSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// Bypass HEAD request.
	// +optional
	BypassHeadRequest *bool `json:"bypassHeadRequest,omitempty" tf:"bypass_head_request"`
	// The interval (in minutes) at which this monitor should run. Valid values are 1, 5, 10, 15, 30, 60, 360, 720, or 1440.
	Frequency *int64 `json:"frequency" tf:"frequency"`
	// The locations in which this monitor should be run.
	// +kubebuilder:validation:MinItems=1
	Locations []string `json:"locations" tf:"locations"`
	// The title of this monitor.
	Name *string `json:"name" tf:"name"`
	// The base threshold for the SLA report.
	// +optional
	SlaThreshold *float64 `json:"slaThreshold,omitempty" tf:"sla_threshold"`
	// The monitor status (i.e. ENABLED, MUTED, DISABLED).
	Status *string `json:"status" tf:"status"`
	// Fail the monitor check if redirected.
	// +optional
	TreatRedirectAsFailure *bool `json:"treatRedirectAsFailure,omitempty" tf:"treat_redirect_as_failure"`
	// The monitor type. Valid values are SIMPLE, BROWSER, SCRIPT_BROWSER, and SCRIPT_API.
	Type *string `json:"type" tf:"type"`
	// The URI for the monitor to hit.
	// +optional
	Uri *string `json:"uri,omitempty" tf:"uri"`
	// The string to validate against in the response.
	// +optional
	ValidationString *string `json:"validationString,omitempty" tf:"validation_string"`
	// Verify SSL.
	// +optional
	VerifySsl *bool `json:"verifySsl,omitempty" tf:"verify_ssl"`
}

type MonitorStatus struct {
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

// MonitorList is a list of Monitors
type MonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Monitor CRD objects
	Items []Monitor `json:"items,omitempty"`
}