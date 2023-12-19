package appresource

import (
	"context"
	"fmt"

	"github.com/giantswarm/app/v7/pkg/validation"
	"github.com/giantswarm/microerror"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/apiextensions-application/api/v1alpha1"
)

func (r *Resource) ApplyUpdateChange(ctx context.Context, obj, updateChange interface{}) error {
	appCRs, err := toAppCRs(updateChange)
	if err != nil {
		return microerror.Mask(err)
	}

	for _, appCR := range appCRs {
		r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("updating App CR %#q in namespace %#q", appCR.Name, appCR.Namespace))

		// Get app CR again to ensure the resource version is correct.
		var currentCR v1alpha1.App
		err := r.g8sClient.Get(ctx, client.ObjectKey{
			Namespace: appCR.Namespace,
			Name:      appCR.Name,
		}, &currentCR)
		if err != nil {
			return microerror.Mask(err)
		}

		appCR.ResourceVersion = currentCR.ResourceVersion

		err = r.g8sClient.Update(ctx, appCR)
		if validation.IsAppConfigMapNotFound(err) {
			// Don't return error as there can be a delay for the cluster configmap being created on cluster creation.
			r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("configmap for App CR %#q in namespace %#q does not exist yet", appCR.Name, appCR.Namespace))
			continue
		} else if validation.IsKubeConfigNotFound(err) {
			// Don't return error as there can be a delay for the cluster kubeconfig being created on cluster creation.
			r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("kubeconfig secret for App CR %#q in namespace %#q does not exist yet", appCR.Name, appCR.Namespace))
			continue
		} else if err != nil {
			return microerror.Mask(err)
		}

		r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("updated App CR %#q in namespace %#q", appCR.Name, appCR.Namespace))
	}

	return nil
}

func (r *Resource) newUpdateChange(ctx context.Context, obj, currentState, desiredState interface{}) (interface{}, error) {
	currentAppCRs, err := toAppCRs(currentState)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	desiredAppCRs, err := toAppCRs(desiredState)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var appCRsToUpdate []*v1alpha1.App
	{
		r.logger.LogCtx(ctx, "level", "debug", "message", "computing App CRs to update")

		for _, c := range currentAppCRs {
			for _, d := range desiredAppCRs {
				m := newAppCRToUpdate(c, d, r.allowedAnnotations)
				if m != nil {
					appCRsToUpdate = append(appCRsToUpdate, m)
				}
			}
		}

		r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("computed %d App CRs to update", len(appCRsToUpdate)))
	}

	return appCRsToUpdate, nil
}
