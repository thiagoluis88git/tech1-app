# Tech Challenge 1 - Kubernetes infrastructure

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Description](#description)
- [Kubernetes Cluster](#kubernetes-cluster)
- [Minikube](#minikube)


## Description

The Fast food project was designed to work with Docker compose as well as `Kubernetes cluster`. Here it will be detailed all the **k8s** files to build and run this project:

## Kubernetes Cluster

The project is structutred in separate peaces within the cluster:

 - ConfigMaps: in `k8s/configmaps` is stored all the environment variables used in **deployments**
 - Deployments: in `k8s/deployments` is the main App, devided in **Application** and **Database**. Inside the *.yaml* files are also implemented the **Services** to make the comunication works between the **Application** and **Database**

## Minikube

To make easy to deploy and test this project we can use **Minikube** test cluster environment.
Just follow the steps on: [Minikube](https://minikube.sigs.k8s.io/docs/)

![ddd_image](https://github.com/thiagoluis88git/tech1/assets/166969350/2016bfff-3c19-4172-837f-8d5d428525f7)
