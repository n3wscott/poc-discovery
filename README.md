# Discovery API

This is Work in Progress.

## Install

```shell script
ko apply -f ./config
```

## DuckType.discovery.knative.dev

The goal is to have a custom type that is use installable to help a developer, cluster admin, or tooling to better 
understand the duck types that are installed in the cluster. This information could be used to understand which Kinds
could fulfill a role for another resource. 

### example.yaml

```bash
kubectl apply --filename - << END
apiVersion: discovery.knative.dev/v1alpha1
kind: DuckType
metadata:
  name: addressable.duck.knative.dev
spec:
  # selectors is a list of CRD label selectors to find CRDs that have been
  # labeled as the given duck type.
  selectors:
    - selector: "duck.knative.dev/addressable=true"

  # refs allows for adding native types, or crds directly as the ducks via
  # Group/Version/Kind/Resource
  refs:
    - version: v1
      resource: services
      kind: Service

  # Names allows us to give a short name to the duck type.
  names:
    plural: addressables
    singular: addressable

  # additionalPrinterColumns is intended to understand what printer columns
  # should be used for the custom objects.
  additionalPrinterColumns:
    - name: Ready
      type: string
      JSONPath: ".status.conditions[?(@.type=='Ready')].status"
    - name: Reason
      type: string
      JSONPath: ".status.conditions[?(@.type=='Ready')].reason"
    - name: URL
      type: string
      JSONPath: .status.address.url
    - name: Age
      type: date
      JSONPath: .metadata.creationTimestamp

  # schema is the partial schema of the duck type.
  schema:
    openAPIV3Schema:
      properties:
        status:
          type: object
          properties:
            address:
              type: object
              properties:
                url:
                  type: string
END
```

After applying this, you can fetch it:

```shell
$  kubectl get ducktypes addressable.duck.knative.dev
NAME                           SHORT NAME    DUCKS   READY   REASON
addressable.duck.knative.dev   addressable   11      True
```

And get the full DuckType `addressable.duck.knative.dev` resource: 

```shell
$ kubectl get ducktypes addressable.duck.knative.dev -oyaml
apiVersion: discovery.knative.dev/v1alpha1
kind: DuckType
metadata:
  generation: 1
  name: addressable.duck.knative.dev
spec:
  additionalPrinterColumns:
  - JSONPath: .status.conditions[?(@.type=='Ready')].status
    name: Ready
    type: string
  - JSONPath: .status.conditions[?(@.type=='Ready')].reason
    name: Reason
    type: string
  - JSONPath: .status.address.url
    name: URL
    type: string
  - JSONPath: .metadata.creationTimestamp
    name: Age
    type: date
  names:
    plural: addressables
    singular: addressable
  refs:
  - kind: Service
    resource: services
    version: v1
  schema:
    openAPIV3Schema:
      properties:
        status:
          properties:
            address:
              properties:
                url:
                  type: string
              type: object
          type: object
  selectors:
  - selector: duck.knative.dev/addressable=true
status:
  conditions:
  - lastTransitionTime: "2020-01-16T22:13:24Z"
    status: "True"
    type: Ready
  duckCount: 11
  ducks:
  - group: eventing.knative.dev
    kind: Broker
    resource: brokers
    version: v1alpha1
  - group: flows.knative.dev
    kind: Parallel
    resource: parallels
    version: v1alpha1
  - group: flows.knative.dev
    kind: Sequence
    resource: sequences
    version: v1alpha1
  - group: messaging.knative.dev
    kind: Channel
    resource: channels
    version: v1alpha1
  - group: messaging.knative.dev
    kind: InMemoryChannel
    resource: inmemorychannels
    version: v1alpha1
  - group: messaging.knative.dev
    kind: Parallel
    resource: parallels
    version: v1alpha1
  - group: messaging.knative.dev
    kind: Sequence
    resource: sequences
    version: v1alpha1
  - group: n3wscott.com
    kind: Task
    resource: tasks
    version: v1alpha1
  - kind: Service
    resource: services
    version: v1
  - group: serving.knative.dev
    kind: Route
    resource: routes
    version: v1alpha1
  - group: serving.knative.dev
    kind: Service
    resource: services
    version: v1alpha1
  observedGeneration: 1
```

## Knative Duck Types

If the `./config/knative` directory is applied:

```shell
ko apply -f ./config/knative/
```

then a quick view of the duck types that are on the cluster becomes easier to get:

```shell
$ kubectl get ducktypes
NAME                                 SHORT NAME     DUCKS   READY   REASON
addressable.duck.knative.dev         addressable    11      True
binding.duck.knative.dev             binding        2       True
podspecable.duck.knative.dev         podspecable    7       True
source.duck.knative.dev              source         7       True
subscribable.messaging.knative.dev   subscribable   2       True
```
