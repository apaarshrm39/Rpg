package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Rgp struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec RgpSpec
}

type RgpSpec struct {
	Location string
	Name     string
}

type RgpList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Rgp
}
