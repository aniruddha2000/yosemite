package client

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	ctx = context.TODO()
)

func CreatePodWithNamespace(namespace, name string, clientset *kubernetes.Clientset) error {
	podObj := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  fmt.Sprintf("test-pod-%s", name),
					Image: "hello-world",
				},
			},
		},
	}

	pod, err := clientset.CoreV1().Pods(namespace).Create(ctx, podObj, metav1.CreateOptions{})
	if err != nil {
		if errors.IsAlreadyExists(err) {
			fmt.Printf("Pod already exists with name %s\n", name)
			return nil
		} else {
			return err
		}
	}
	fmt.Printf("Pod object created with name %s\n", pod.ObjectMeta.Name)

	return nil
}

func CrateNameSpace(namespeceName string, clientset *kubernetes.Clientset) error {
	nameSpaceObj := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespeceName,
		},
	}

	_, err := clientset.CoreV1().Namespaces().Create(ctx, nameSpaceObj, metav1.CreateOptions{})
	if err != nil {
		if errors.IsAlreadyExists(err) {
			fmt.Printf("Namespace already exists with name %s\n", namespeceName)
			return nil
		} else {
			return err
		}
	}
	fmt.Printf("namespace created %v\n", namespeceName)

	return nil
}

func DeletePodWithNamespce(namespace, name string, clientset *kubernetes.Clientset) error {
	err := clientset.CoreV1().Pods(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			fmt.Printf("Pod don't exists with name %s\n", name)
			return nil
		} else {
			return err
		}
	}
	fmt.Printf("pod deleted with name: %v\n", name)

	return nil
}
