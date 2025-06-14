package v1alpha1

import (
	v1alpha1 "github.com/orenr2301/KubeTag/pkg/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/***
 API Type structures
***/

// desired state definitions of Namespace labels
type NsLabelSetSpec struct {
	Namespace string              `json:"namespace"`
	Labels    map[string]string   `json:"labels"`
}

// observed state definitions of Namespace labels
type NsLabelSetStatus struct {
	Applied bool    `json:"applied,omitempty"`
	Message string  `json:"message,omitempty"`

}

// API Schema for the NsLabelSet resource
type NsLabelSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NsLabelSetSpec     `json:"spec,omitempty"`
	Status NsLabelSetStatus   `json:"status,omitempty"`

}


type NsLabelSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items            []NsLabelSet `json:"items"`
}

