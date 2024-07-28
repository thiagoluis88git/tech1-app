# FastFood API - Kubernetes infrastructure

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Description](#description)
- [Kubernetes Cluster](#kubernetes-cluster)
- [Minikube](#minikube)


## Description

The Fast food project was designed to work with **Docker compose** as well as `Kubernetes cluster`. Here it will be detailed all the **k8s** files to build and run this project:

## Kubernetes Cluster

The project is structutred in specific components within the cluster to make the project scale when it needs. To do so, it requires some Kubernetes components installed in the Cluster. These are:

 - ConfigMaps: in `k8s/configmaps` is stored all the environment variables used in **deployments**
 - Deployments: in `k8s/deployments` is the main App, devided in **Application** and **Database**. Inside the *.yml* files are also implemented some **Services** components to make the comunication works between the **Application** and **Database**
 - HPA: in `k8s/hpa` has the **HorizontalPodAutoscaler** specification. With *HorizontalPodAutoscaler* we can enable an auto scaling of application PODs
 - Metrics: in `k8s/metrics` has some components to enable the Cluster metrics access. To enable *HPA* work as expected, this component is required because its activate some deployments to consume the Cluster metrics. With these metrics, the *HPA* component can scale the PODs via some metric types. 
 
 To apply the metrics in the `AWS EKS`, run the command:

 ```
 kubectl apply -f infra/k8s/metrics/components.yaml
 ```

 To apply the metrics in the `Minikube`, run the command:

 ```
 kubectl apply -f infra/k8s/metrics/components-insecure.yaml
 ```
 
 After applied, wait some minutes. To show the metrics just run the command:

 ```
kubectl top nodes

NAME                              CPU(cores)   CPU%   MEMORY(bytes)   MEMORY%   
ip-192-168-128-206.ec2.internal   34m          1%     584Mi           17% 
 ```

 - Secret: in `k8s/secret` has the Kustomization component. **This folder is not tracked in Github**. This will set in the Kubernetes cluster some **environment secrets** to be consumed by the application. 
 To use tha same environment varibles, use the `kind: Kustomization` like this:

```
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
generatorOptions:
  disableNameSuffixHash: true
secretGenerator:
- name: app-env-secret
  literals:
  - QR_CODE_GATEWAY_ROOT_URL=<MERCADO_LIVRE_API>
  - QR_CODE_GATEWAY_TOKEN=<MERCADO_LIVRE_TOKEN>
  - WEBHOOK_MERCADO_LIVRE_PAYMENT=<FASTFOOD_API_WEBHOOK_HANDLER>
  - POSTGRES_PASSWORD=<RDS_POSTGRES_PASSWORD>
```

 - Volumes: in `k8s/volumes` have the **PV** and **PVC***. With those we can create a persistent storage to work with the Database application

## Minikube

To make easy to deploy and test this project we can use **Minikube** test cluster environment.
Just follow the steps on: [Minikube](https://minikube.sigs.k8s.io/docs/)

