package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ConfigMapGeneratorSpec defines the desired state of ConfigMapGenerator
type ConfigMapGeneratorSpec struct {
	// Key to be added in the ConfigMap
	Key string `json:"key"`
	// Value associated with the key in the ConfigMap
	Value string `json:"value"`
}

// ConfigMapGeneratorStatus defines the observed state of ConfigMapGenerator
type ConfigMapGeneratorStatus struct {
	// Status of the resource
	Status string `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ConfigMapGenerator is the Schema for the configmapgenerators API
type ConfigMapGenerator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigMapGeneratorSpec   `json:"spec,omitempty"`
	Status ConfigMapGeneratorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConfigMapGeneratorList contains a list of ConfigMapGenerator
type ConfigMapGeneratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigMapGenerator `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfigMapGenerator{}, &ConfigMapGeneratorList{})
}
