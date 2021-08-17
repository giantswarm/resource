module github.com/giantswarm/resource/v2

go 1.14

require (
	github.com/giantswarm/apiextensions/v3 v3.30.0
	github.com/giantswarm/app/v5 v5.2.2
	github.com/giantswarm/microerror v0.3.0
	github.com/giantswarm/micrologger v0.5.0
	github.com/giantswarm/operatorkit/v5 v5.0.0
	k8s.io/apimachinery v0.18.19
)

// keep in sync with giantswarm/apiextensions
replace sigs.k8s.io/cluster-api => github.com/giantswarm/cluster-api v0.3.10-gs
