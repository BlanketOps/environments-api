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

/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0
*/
// +k8s:openapi-gen=true

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

//
// ─────────────────────────────────────────────────────────────
// GitHubPayload Spec (Kubernetes-facing envelope)
// ─────────────────────────────────────────────────────────────
//

// GitHubPayloadSpec is a Kubernetes-native envelope around the canonical
// BlanketOps GitHubPayload contract.
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
type GitHubPayloadSpec struct {

	// Contract is the canonical BlanketOps GitHubPayload specification.
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
// GitHubPayload Status (Kubernetes-facing envelope)
// ─────────────────────────────────────────────────────────────
//

// GitHubPayloadStatus represents observed state owned by the controller.
//
// This mirrors the contract status but remains opaque to Kubernetes.
type GitHubPayloadStatus struct {

	// Contract is the canonical BlanketOps GitHubPayload status.
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
// GitHubPayload CRD
// ─────────────────────────────────────────────────────────────
//

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// GitHubPayload represents a declarative request to GitHubPayload an artifact.
//
// This resource is a Kubernetes-native envelope around a
// transport-agnostic BlanketOps GitHubPayload contract.
type GitHubPayload struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GitHubPayloadSpec   `json:"spec,omitempty"`
	Status GitHubPayloadStatus `json:"status,omitempty"`
}

//
// ─────────────────────────────────────────────────────────────
// GitHubPayloadList
// ─────────────────────────────────────────────────────────────
//

// +kubebuilder:object:root=true

// GitHubPayloadList contains a list of GitHubPayload resources.
type GitHubPayloadList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitHubPayload `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GitHubPayload{}, &GitHubPayloadList{})
}
