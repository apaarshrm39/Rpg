package main

import (
	"flag"
	"path/filepath"
	"time"

	rgp "github.com/apaarshrm39/rgp/pkg/client/clientset/versioned"
	rgp_fac "github.com/apaarshrm39/rgp/pkg/client/informers/externalversions"
	rgp_controller "github.com/apaarshrm39/rgp/pkg/controller"
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

	informers := rgp_fac.NewSharedInformerFactory(rgpset, 10*time.Minute)
	c := rgp_controller.NewController(*rgpset, informers.Apaar().V1alpha1().Rgps())
	ch := make(<-chan struct{})
	informers.Start(ch)
	c.Run(ch)

}
