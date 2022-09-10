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

func (c *Client) CheckDeploymentEnv(ns string) {
	informerFactory := informers.NewSharedInformerFactory(c.C, 30*time.Second)

	deploymentInformer := informerFactory.Apps().V1().Deployments()
	deploymentInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			log.Println("Deployment added. Let's start checking!")
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
	deployments, err := ListDeploymentWithNamespace(namespace, c.C)
	if err != nil {
		return fmt.Errorf("list deployment: %s", err.Error())
	}

	for _, deployment := range deployments.Items {
		var envSet bool
		for _, cntr := range deployment.Spec.Template.Spec.Containers {
			for _, env := range cntr.Env {
				if env.Name == ENVNAME {
					log.Printf("Deployment name: %s has envvar. All set to go!", deployment.Name)
					envSet = true
				}
			}
		}
		if !envSet {
			log.Printf("No envvar name %s - Deleting deployment with name %s\n", ENVNAME, deployment.Name)
			err = DeleteDeploymentWithNamespce(namespace, deployment.Name, c.C)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
