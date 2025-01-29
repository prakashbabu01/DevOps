package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodConfigMapSpec defines the desired state of PodConfigMap
type PodConfigMapSpec struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// PodConfigMapStatus defines the observed state of PodConfigMap
type PodConfigMapStatus struct {
	ConfigMapName string `json:"configMapName,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// PodConfigMap is the Schema for the podconfigmaps API
type PodConfigMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodConfigMapSpec   `json:"spec,omitempty"`
	Status PodConfigMapStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PodConfigMapList contains a list of PodConfigMap
type PodConfigMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodConfigMap `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PodConfigMap{}, &PodConfigMapList{})
}
