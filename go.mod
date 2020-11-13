module github.com/giantswarm/resource/v2

go 1.15

require (
	github.com/giantswarm/apiextensions/v3 v3.7.0
	github.com/giantswarm/app/v3 v3.2.0
	github.com/giantswarm/microerror v0.2.1
	github.com/giantswarm/micrologger v0.3.4
	github.com/giantswarm/operatorkit/v2 v2.0.0
	k8s.io/apimachinery v0.18.9
)

replace (
	// Use fork of CAPI with Kubernetes 1.18 support.
	sigs.k8s.io/cluster-api => github.com/giantswarm/cluster-api v0.3.10-gs
)
