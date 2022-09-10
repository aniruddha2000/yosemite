package client

import (
	"context"

	"k8s.io/client-go/kubernetes"
)

var (
	ctx = context.TODO()
)

type Client struct {
	C *kubernetes.Clientset
}

// Return a new Client
func NewClient() *Client {
	return &Client{}
}
