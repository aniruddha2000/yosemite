package client

import (
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListPodWithNamespace(namspace string, clientset *kubernetes.Clientset) (*v1.PodList, error) {
	pods, err := clientset.CoreV1().Pods(namspace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return pods, nil
}

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
			log.Printf("Pod already exists with name %s\n", name)
			return nil
		} else {
			return fmt.Errorf("create pod: %v", err)
		}
	}
	log.Printf("Pod object created with name %s\n", pod.ObjectMeta.Name)

	return nil
}

func DeletePodWithNamespce(namespace, name string, clientset *kubernetes.Clientset) error {
	err := clientset.CoreV1().Pods(namespace).Delete(ctx, name, metav1.DeleteOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			log.Printf("Pod don't exists with name %s\n", name)
			return nil
		} else {
			return fmt.Errorf("delete pod: %v", err)
		}
	}
	log.Printf("pod deleted with name: %v\n", name)

	return nil
}
