package client

import (
	"fmt"
	"log"

	"k8s.io/client-go/kubernetes"
)

const (
	ENVNAME = "TEST_ENV_NAME"
)

func CheckPodEnv(namespace string, clientset *kubernetes.Clientset) error {
	pods, err := ListPodWithNamespace(namespace, clientset)
	if err != nil {
		return fmt.Errorf("list pod: %s", err.Error())
	}

	for _, pod := range pods.Items {
		var envSet bool
		for _, cntr := range pod.Spec.Containers {
			for _, env := range cntr.Env {
				if env.Name == ENVNAME {
					log.Println("Env value set. All set to go!")
					envSet = true
				}
			}
		}
		if !envSet {
			log.Printf("No envvar name %s - Deleting pod with name %s\n", ENVNAME, pod.Name)
			err = DeletePodWithNamespce(namespace, pod.Name, clientset)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
