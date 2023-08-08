[![CircleCI](https://dl.circleci.com/status-badge/img/gh/giantswarm/resource/tree/master.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/giantswarm/resource/tree/master)

# Resource

Package resource implements a generic [operatorkit] resources for managing
objects defined in [apiextensions] repository. Currently only App resources are supported.

If you look for resources for built-in Kubernetes objects then operatorkit
[resource/k8s] package is your friend.

## App resource

This is a simplified implementation of OperatorKit's CRUD interface for App CRs.
Currently only used [cluster-operator](https://github.com/giantswarm/cluster-operator).

[apiextensions]: https://github.com/giantswarm/apiextensions/
[operatorkit]: https://github.com/giantswarm/operatorkit/
[pkg/resource/k8s]: https://github.com/giantswarm/operatorkit/tree/master/pkg/resource/k8s/
