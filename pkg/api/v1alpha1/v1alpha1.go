package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

/***
 API Type structures
***/

// desired state definitions of Namespace labels
type NsLabelSetSpec struct {
	Namespace string            `json:"namespace"`
	Labels    map[string]string `json:"labels"`
}

// observed state definitions of Namespace labels
type NsLabelSetStatus struct {
	Applied bool   `json:"applied,omitempty"`
	Message string `json:"message,omitempty"`
}

// API Schema for the NsLabelSet resource
type NsLabelSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NsLabelSetSpec   `json:"spec,omitempty"`
	Status NsLabelSetStatus `json:"status,omitempty"`
}

type NsLabelSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NsLabelSet `json:"items"`
}

/***
DeepCopy for controlller runtime - Controller need the ability to
copy CR object without affecting the other CR struct located in memory
***/

// The Single Card Item with fields
func (in *NsLabelSet) DeepCopy() *NsLabelSet {
	if in == nil {
		return nil
	}
	out := new(NsLabelSet)
	*out = *in
	if in.Spec.Labels != nil {
		out.Spec.Labels = make(map[string]string, len(in.Spec.Labels))
		for k, v := range in.Spec.Labels {
			out.Spec.Labels[k] = v
		}
	}
	return out
}

// The Single Card Item Structure - the CR Single Resource 
func (in *NsLabelSet) DeepCopyObject() runtime.Object {

	if in == nil {
		return nil
	}

	out := new(NsLabelSet)
	*out = *in
	if in.Spec.Labels != nil {
		out.Spec.Labels = make(map[string]string, len(in.Spec.Labels))
		for k, v := range in.Spec.Labels {
			out.Spec.Labels[k] = v
		}
	}
	return out
}

// The The list of cards the actual binder - the list of CRs of the the same type object - List of Resource 
func (in *NsLabelSetList) DeepCopyObject() runtime.Object {

	if in == nil {
		return nil
	}

	out := new(NsLabelSetList)
	*out = *in
	if in.Items != nil {
		out.Items = make([]NsLabelSet, len(in.Items))
		for i := range in.Items {
			out.Items[i] = *in.Items[i].DeepCopy()
		}
	}
	return out
}
