package appresource

import (
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/apiextensions-application/api/v1alpha1"
)

type Config struct {
	G8sClient   client.Client
	Logger      micrologger.Logger
	StateGetter StateGetter

	AllowedAnnotations []string
	Name               string
}

type Resource struct {
	g8sClient   client.Client
	logger      micrologger.Logger
	stateGetter StateGetter

	allowedAnnotations map[string]bool
	name               string
}

func New(config Config) (*Resource, error) {
	if config.G8sClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.G8sClient must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}
	if config.StateGetter == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.StateGetter must not be empty", config)
	}

	if config.Name == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.Name must not be empty", config)
	}

	r := &Resource{
		g8sClient:   config.G8sClient,
		logger:      config.Logger,
		stateGetter: config.StateGetter,

		name: config.Name,
	}

	if config.AllowedAnnotations != nil {
		allowedAnnotation := map[string]bool{}
		{
			for _, annotation := range config.AllowedAnnotations {
				allowedAnnotation[annotation] = true
			}
		}

		r.allowedAnnotations = allowedAnnotation
	}

	return r, nil
}

func (r *Resource) Name() string {
	return r.name
}

func containsAppCR(appCRs []*v1alpha1.App, appCR *v1alpha1.App) bool {
	for _, a := range appCRs {
		if appCR.Name == a.Name && appCR.Namespace == a.Namespace {
			return true
		}
	}

	return false
}

func toAppCRs(v interface{}) ([]*v1alpha1.App, error) {
	x, ok := v.([]*v1alpha1.App)
	if !ok {
		return nil, microerror.Maskf(wrongTypeError, "expected '%T', got '%T'", x, v)
	}

	return x, nil
}
