module github.com/giantswarm/resource/v2

go 1.15

require (
	github.com/giantswarm/apiextensions/v3 v3.8.0
	github.com/giantswarm/app/v3 v3.2.1-0.20201113134454-b8c4868e92c6
	github.com/giantswarm/microerror v0.2.1
	github.com/giantswarm/micrologger v0.3.4
	github.com/giantswarm/operatorkit/v2 v2.0.0
	k8s.io/apimachinery v0.18.9
)

replace (
	// Use v0.8.10 of hcsshim to fix nancy alert.
	github.com/Microsoft/hcsshim v0.8.7 => github.com/Microsoft/hcsshim v0.8.10
	// Apply fix for CVE-2020-15114 not yet released in github.com/spf13/viper.
	github.com/bketelsen/crypt => github.com/bketelsen/crypt v0.0.3
	// Use v1.0.0-rc7 of runc to fix nancy alert.
	github.com/opencontainers/runc v0.1.1 => github.com/opencontainers/runc v1.0.0-rc7
	// Use fork of CAPI with Kubernetes 1.18 support.
	sigs.k8s.io/cluster-api => github.com/giantswarm/cluster-api v0.3.10-gs
)
