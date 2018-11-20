# Kube-importer for Containerum Platform
Kube-importer is a service that imports pre-existing Kubernetes objects to [Containerum Platform](https://github.com/containerum/containerum).

## Prerequisites
* Kubernetes

## Installation

### Using Helm

```
  helm repo add containerum https://charts.containerum.io
  helm repo update
  helm install containerum/kube-importer
```

## Contributions
Please submit all contributions concerning Kube-importer component to this repository. Contributing guidelines are available [here](https://github.com/containerum/containerum/blob/master/CONTRIBUTING.md).

## License
Kube-importer project is licensed under the terms of the Apache License Version 2.0. Please see LICENSE in this repository for more details.

