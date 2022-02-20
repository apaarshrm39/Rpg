package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	rgp "github.com/apaarshrm39/rgp/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "Absoluture Path to the Kubeconfig")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	rgpset, err := rgp.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	rgp, err := rgpset.ApaarV1alpha1().Rgps("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println(len(rgp.Items))
}
