package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Rgp struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec RgpSpec
}

type RgpSpec struct {
	Location string
	Name     string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type RgpList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Rgp
}
