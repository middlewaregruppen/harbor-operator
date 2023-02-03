[![Go](https://github.com/middlewaregruppen/harbor-operator/actions/workflows/go.yaml/badge.svg?branch=master)](https://github.com/middlewaregruppen/harbor-operator/actions/workflows/go.yaml) [![Release](https://github.com/middlewaregruppen/harbor-operator/actions/workflows/release.yaml/badge.svg)](https://github.com/middlewaregruppen/harbor-operator/actions/workflows/release.yaml)

# harbor-operator
Manage [Harbor](https://goharbor.io/) registries with Kubernetes

## Description
harbor-operator lets admins and developers manage resources in Harbor using Kubernetes through CRD's, such as Projects & Registries.

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
1. Install CRD's and the controller onto the clusters
```sh
kubectl apply -k config/crd
kubectl apply -k config/default
```

2. Create secret holding credentials to the Harbor server
```sh
kubectl create secret generic harbor-credentials --from-literal="username=admin" --from-literal="password=Harbor12345"
```

2. Create a `HarborService`
```sh
cat <<EOF | kubectl apply -f -
apiVersion: harbor.mdlwr.com/v1alpha1
kind: HarborService
metadata:
  name: harbor-default
spec:
  insecure: true
  scheme: https
  externalBackend:
    host: harbor.my.company.com
    port: 443
  secretRef:
    name: harbor-credentials
EOF
```

3. Start creating resources, such as `Project` with a reference to your `HarborService`
```sh
cat <<EOF | kubectl apply -f -
apiVersion: harbor.mdlwr.com/v1alpha1
kind: HarborProject
metadata:
  name: project-01
spec:
  harbor: harbor-default
EOF
```

### Running locally (for contributors)
1. Install CRD's
```sh
make install
```
2. Run the controller
```sh
make run
```

### Uninstall CRDs
To delete the CRDs from the cluster:

```sh
make uninstall
```

### Undeploy controller
UnDeploy the controller to the cluster:

```sh
make undeploy
```

## Contributing
You are welcome to contribute to this project by opening PR's. Create an Issue if you have feedback