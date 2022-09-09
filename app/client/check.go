package client

import (
	"fmt"
	"log"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

const (
	ENVNAME = "TEST_ENV_NAME"
)

func (c *Client) CheckPodEnv(ns string) {
	informerFactory := informers.NewSharedInformerFactory(c.C, 30*time.Second)

	podInformer := informerFactory.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			log.Println("Pod added. Let's start checking!")
			err := c.check(ns)
			if err != nil {
				log.Fatalf("error checking envvar: %v", err)
			}
		},
	})

	informerFactory.Start(wait.NeverStop)
	informerFactory.WaitForCacheSync(wait.NeverStop)

}

func (c *Client) check(namespace string) error {
	pods, err := ListPodWithNamespace(namespace, c.C)
	if err != nil {
		return fmt.Errorf("list pod: %s", err.Error())
	}

	for _, pod := range pods.Items {
		var envSet bool
		for _, cntr := range pod.Spec.Containers {
			for _, env := range cntr.Env {
				if env.Name == ENVNAME {
					log.Printf("Pod name: %s has envvar. All set to go!", pod.Name)
					envSet = true
				}
			}
		}
		if !envSet {
			log.Printf("No envvar name %s - Deleting pod with name %s\n", ENVNAME, pod.Name)
			err = DeletePodWithNamespce(namespace, pod.Name, c.C)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
