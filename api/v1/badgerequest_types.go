package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type BadgeRequestSpec struct {
	Title              string   `json:"title"`            // Badge title
	Theme              string   `json:"theme"`            // Badge theme
	Icons              []string `json:"icons,omitempty"`  // List of icons
	Colors             []string `json:"colors,omitempty"` // List of colors
	AdditionalElements string   `json:"additionalElements"`
}

type BadgeRequestStatus struct {
	Status    string `json:"status"`
	OutputURL string `json:"outputURL"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// BadgeRequest is the Schema for the badgerequests API.
type BadgeRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BadgeRequestSpec   `json:"spec,omitempty"`
	Status BadgeRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BadgeRequestList contains a list of BadgeRequest.
type BadgeRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BadgeRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BadgeRequest{}, &BadgeRequestList{})
}
