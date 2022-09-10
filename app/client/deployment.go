package client

import (
	"fmt"
	"log"

	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListDeploymentWithNamespace(ns string, clientset *kubernetes.Clientset) (*v1.DeploymentList, error) {
	deployment, err := clientset.AppsV1().Deployments(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return deployment, nil
}

func DeleteDeploymentWithNamespce(ns, name string, clientset *kubernetes.Clientset) error {
	err := clientset.AppsV1().Deployments(ns).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			log.Printf("Deployment don't exists with name %s\n", name)
			return nil
		} else {
			return fmt.Errorf("delete Deployment: %v", err)
		}
	}
	log.Printf("Deployment deleted with name: %v\n", name)

	return nil
}
