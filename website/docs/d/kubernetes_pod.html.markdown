---
subcategory: "kubernetes"
layout: "intersight"
page_title: "Intersight: intersight_kubernetes_pod"
description: |-
  Pod represents a Kubernetes Pod. In Intersight, it is a read only model that represent Kubernetes Pod information. In Kubernetes, A Pod is the basic execution unit of a Kubernetes application the smallest and simplest unit in the Kubernetes object model that you create or deploy. A Pod represents processes running on your Cluster.
---

# Data Source: intersight_kubernetes_pod
Pod represents a Kubernetes Pod. In Intersight, it is a read only model that represent Kubernetes Pod information. In Kubernetes, A Pod is the basic execution unit of a Kubernetes application the smallest and simplest unit in the Kubernetes object model that you create or deploy. A Pod represents processes running on your Cluster.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the referenced kubernetes resource. 
* `uuid`:(string) UUID of the referenced kubernetes resource. It is generated by the kubernetes cluster itself. 
