# yosemite
----------

[package](https://pkg.go.dev/github.com/aniruddha2000/yosemite)

It's a package that allows deleting the deployment if it doesn't have the desired environment variable in the specified namespace.

## Setup

```bash
$ kubectl apply -f templates/rbac/

$ kubectl apply -f templates/deployment/test_deployment.yaml

$ kubectl apply -f templates/deployment/client.yaml
```
