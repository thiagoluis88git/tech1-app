# FastFood API - AWS Infrastructure

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Description](#description)
- [AWS VPC Description](#aws-vpc-description)
- [AWS EKS](#aws-eks)

## Description

The Fast food project uses `AWS Cloud` to host its software components. 

## AWS VPC Description

To follow some of the best practices of network security in Cloud, the FastFood API uses the Private and Public network strategy to configure the `VPC` on **AWS**. This configuration can be created via `AWS Cloud Formation` using [this YAML file](https://s3.us-west-2.amazonaws.com/amazon-eks/cloudformation/2020-10-29/amazon-eks-vpc-private-subnets.yaml). The **most important** aspect for this *hybrid* solution is:

- Software components are stored in `Private Subnets`
- Infrastructure components, such as `NAT Gateways` are stored in `Public Subnets`
- The `EKS Load Balancers Services` will traffic the internet data between applications inside **Private Subnets**
- All the traffic between **Private Subnets** and **Public Subnets** are done by `Route tables` 

<img width="602" alt="AWS infra" src="https://github.com/user-attachments/assets/178cafb0-7552-41e7-b826-40795f595145">


With this approach the **FastFood API** is secure and safe from possible internet attackers

## AWS EKS

The main AWS product used in FastFood project is `AWS EKS`. To know more about this configuration, read: [Kubernetes README](infra/k8s/README.md)
