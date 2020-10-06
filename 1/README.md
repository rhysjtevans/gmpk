# Deploy three tier application on Kubernetes
This chart deploys a three tier application. [reactapp](https://github.com/yogeshlonkar/reactapp) is used as front-end, [nodeapp](https://github.com/yogeshlonkar/nodeapp) as back-end along with mongodb `StatefulSet`. This chart also exposes Ingress service for `nodeapp`. The chart manages dependency between different pods using [pod-dependency-init-container](https://github.com/yogeshlonkar/pod-dependency-init-container).

Prerequisites

- Kubernetes 1.8+
- helm 3.0+

## Notes:
A SPA-based application, the preferred method would be to package it and offload it to a CDN e.g. Azure CDN or S3-backed AWS CloudFront.

## Deploy the Chart
To install the chart with the release name my-release:
```
helm upgrade -i -n my-namespace my-release ./3tierchart
```
