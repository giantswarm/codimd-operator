package controllers

import (
	"context"

	"github.com/giantswarm/microerror"
	appsv1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"

	codiv1alpha1 "github.com/giantswarm/codimd-operator/api/v1alpha1"
)

func (r *CodiMDReconciler) delete(ctx context.Context, cr codiv1alpha1.CodiMD) error {
	// Get the deployment reference we want to delete from our CR status.
	key := types.NamespacedName{
		Name:      cr.Status.Target.Name,
		Namespace: cr.Status.Target.Namespace,
	}
	var deployment appsv1.Deployment

	// Get the actual kubernetes deployment with information from our status.
	err := r.MgrClient.Get(ctx, key, &deployment)
	if apierrors.IsNotFound(err) {
		return nil
	} else if err != nil {
		return microerror.Mask(err)
	}

	// Delete the kubernetes deployment.
	err = r.MgrClient.Delete(ctx, &deployment)
	if apierrors.IsNotFound(err) {
		return nil
	} else if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
