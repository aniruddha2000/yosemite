package client

import (
	"fmt"
	"log"
)

const (
	ENVNAME = "TEST_ENV_NAME"
)

func (c *Client) CheckPodEnv(ns string) {
	err := c.check(ns)
	if err != nil {
		log.Fatalf("error checking envvar: %v", err)
	}
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
					log.Println("Env value set. All set to go!")
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
