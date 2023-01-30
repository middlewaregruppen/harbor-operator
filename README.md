# harbor-operator
Manage [Harbor](https://goharbor.io/) registries with Kubernetes

## Description
harbor-operator lets admins and developers manage resources in Harbor using Kubernetes through CRD's, such as Projects & Registries.

## Getting Started
You’ll need a Kubernetes cluster to run against. You can use [KIND](https://sigs.k8s.io/kind) to get a local cluster for testing, or run against a remote cluster.
**Note:** Your controller will automatically use the current context in your kubeconfig file (i.e. whatever cluster `kubectl cluster-info` shows).

### Running on the cluster
```sh
kubectl apply -k config/crd
kubectl apply -k config/default
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