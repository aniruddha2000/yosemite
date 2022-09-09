package client

import (
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CrateNameSpace(namespeceName string, clientset *kubernetes.Clientset) error {
	nameSpaceObj := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespeceName,
		},
	}

	_, err := clientset.CoreV1().Namespaces().Create(ctx, nameSpaceObj, metav1.CreateOptions{})
	if err != nil {
		if errors.IsAlreadyExists(err) {
			log.Printf("Namespace already exists with name %s\n", namespeceName)
			return nil
		} else {
			return fmt.Errorf("create namespace: %s", err.Error())
		}
	}
	log.Printf("namespace created %v\n", namespeceName)

	return nil
}
