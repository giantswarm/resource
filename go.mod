module github.com/giantswarm/resource/v3

go 1.14

require (
	github.com/giantswarm/apiextensions/v3 v3.38.0
	github.com/giantswarm/app/v5 v5.4.0
	github.com/giantswarm/microerror v0.3.0
	github.com/giantswarm/micrologger v0.5.0
	github.com/giantswarm/operatorkit/v5 v5.0.0
	k8s.io/apimachinery v0.18.19
)

// keep in sync with giantswarm/apiextensions
replace (
	github.com/coreos/etcd v3.3.10+incompatible => github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/etcd v3.3.13+incompatible => github.com/coreos/etcd v3.3.25+incompatible
	github.com/dgrijalva/jwt-go => github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	sigs.k8s.io/cluster-api => github.com/giantswarm/cluster-api v0.3.10-gs
)
