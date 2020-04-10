# Auction Bidder

Refer the Design Document [here](./docs/design.md)

## Prerequisits
1. [GNU Make](https://www.gnu.org/software/make/)
2. [Docker](https://www.docker.com/get-started)
3. [Kubernetes](https://kubernetes.io/docs/setup/)
4. [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)
5. [HelmCharts](https://helm.sh/docs/intro/)

## How To Run

```sh
$ minikube start
$ minikube addons enable ingress

$ eval $(minikube docker-env)
$ make all

```

Refer the postman collection [here](./api/postman_collection.json)
