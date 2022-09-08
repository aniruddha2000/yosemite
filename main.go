package main

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/aniruddha2000/yosemite/app/client"
)

func main() {
	config, err := rest.InClusterConfig()
	if err != nil {
		kubeconfig := filepath.Join("home", "aniruddha", ".kube", "config")
		if envvar := os.Getenv("KUBECONFIG"); len(envvar) > 0 {
			kubeconfig = envvar
		}

		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			fmt.Printf("kubeconfig can't be loaded: %v\n", err)
			os.Exit(0)
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("error getting config client: %v\n", err)
		os.Exit(0)
	}

	err = client.CrateNameSpace("test-ns", clientset)
	if err != nil {
		fmt.Printf("error creating namespace: %v\n", err)
		os.Exit(0)
	}

	err = client.CheckPodEnv("test-ns", clientset)
	if err != nil {
		fmt.Printf("error checking envvar: %v", err)
	}

	// err = client.CreatePodWithNamespace("test-ns", "example", clientset)
	// if err != nil {
	// 	fmt.Printf("error creating pod with namespace: %v\n", err)
	// 	os.Exit(0)
	// }

	// err = client.DeletePodWithNamespce("book", "example", clientset)
	// if err != nil {
	// 	fmt.Printf("error deleting pod: %v\n", err)
	// 	os.Exit(0)
	// }
}
