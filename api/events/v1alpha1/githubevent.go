/*
Copyright 2026 The BlanketOps Authors.
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

// +k8s:openapi-gen=true
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

//
// ─────────────────────────────────────────────────────────────
// GitHubEvent Spec (Kubernetes-facing envelope)
// ─────────────────────────────────────────────────────────────
//

// GitHubEventSpec is a Kubernetes-native envelope around the canonical
// BlanketOps GitHubEvent contract.
//
// IMPORTANT:
// - Kubernetes does NOT understand the contents of `Contract`
// - Kubernetes does NOT validate the contents of `Contract`
// - Kubernetes ONLY stores and round-trips this field
//
// Ownership boundaries:
// - API server: envelope + metadata
// - Controller: lifecycle orchestration
// - Contract layer: semantic meaning
type GitHubEventSpec struct {

	// Contract is the canonical BlanketOps GitHubEvent specification.
	//
	// This field is intentionally opaque to Kubernetes and schema generation.
	// It is preserved verbatim by the API server.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:pruning:PreserveUnknownFields
	Contract runtime.RawExtension `json:"contract"`
}

//
// ─────────────────────────────────────────────────────────────
// GitHubEvent Status (Kubernetes-facing envelope)
// ─────────────────────────────────────────────────────────────
//

// GitHubEventStatus represents observed state owned by the controller.
//
// This mirrors the contract status but remains opaque to Kubernetes.
type GitHubEventStatus struct {

	// Contract is the canonical BlanketOps GitHubEvent status.
	//
	// This field is opaque and preserved verbatim.
	//
	// +optional
	// +kubebuilder:pruning:PreserveUnknownFields
	Contract runtime.RawExtension `json:"contract,omitempty"`

	// Conditions follows standard Kubernetes condition conventions.
	// These are intended for kubectl, UIs, and ecosystem tooling.
	//
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//
// ─────────────────────────────────────────────────────────────
// GitHubEvent CRD
// ─────────────────────────────────────────────────────────────
//

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// GitHubEvent represents a declarative request to GitHubEvent an artifact.
//
// This resource is a Kubernetes-native envelope around a
// transport-agnostic BlanketOps GitHubEvent contract.
type GitHubEvent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GitHubEventSpec   `json:"spec,omitempty"`
	Status GitHubEventStatus `json:"status,omitempty"`
}

//
// ─────────────────────────────────────────────────────────────
// GitHubEventList
// ─────────────────────────────────────────────────────────────
//

// +kubebuilder:object:root=true

// GitHubEventList contains a list of GitHubEvent resources.
type GitHubEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitHubEvent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GitHubEvent{}, &GitHubEventList{})
}
