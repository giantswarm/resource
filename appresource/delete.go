package appresource

import (
	"context"
	"fmt"

	"github.com/giantswarm/apiextensions-application/api/v1alpha1"
	"github.com/giantswarm/microerror"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

func (r *Resource) ApplyDeleteChange(ctx context.Context, obj, deleteChange interface{}) error {
	appCRsToDelete, err := toAppCRs(deleteChange)
	if err != nil {
		return microerror.Mask(err)
	}

	for _, appCR := range appCRsToDelete {
		r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("deleting App CR %#q in namespace %#q", appCR.Name, appCR.Namespace))

		err := r.g8sClient.Delete(ctx, appCR)
		if apierrors.IsNotFound(err) {
			r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("already deleted App CR %#q in namespace %#q", appCR.Name, appCR.Namespace))
		} else if err != nil {
			return microerror.Mask(err)
		} else {
			r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("deleted App CR %#q in namespace %#q", appCR.Name, appCR.Namespace))
		}

	}

	return nil
}

func (r *Resource) newDeleteChange(ctx context.Context, obj, currentState, desiredState interface{}) ([]*v1alpha1.App, error) {
	currentAppCRs, err := toAppCRs(currentState)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	desiredAppCRs, err := toAppCRs(desiredState)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var appCRsToDelete []*v1alpha1.App
	{
		r.logger.LogCtx(ctx, "level", "debug", "message", "computing App CRs to delete")

		for _, c := range currentAppCRs {
			if !containsAppCR(desiredAppCRs, c) {
				appCRsToDelete = append(appCRsToDelete, c)
			}
		}

		r.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("computed %d App CRs to delete", len(appCRsToDelete)))
	}

	return appCRsToDelete, nil
}
