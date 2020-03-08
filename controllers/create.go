package controllers

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	ref "k8s.io/client-go/tools/reference"

	codiv1alpha1 "github.com/giantswarm/codimd-operator/api/v1alpha1"
)

func (r *CodiMDReconciler) create(ctx context.Context, cr codiv1alpha1.CodiMD, deployment *appsv1.Deployment) error {
	// Create or update the k8s deployment.
	err := r.MgrClient.Create(ctx, deployment)
	if apierrors.IsAlreadyExists(err) {
		err = r.MgrClient.Update(ctx, deployment)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// Update the status with a reference to the deployment.
	if cr.Status.Target.ResourceVersion == "" {
		deploymentRef, err := ref.GetReference(r.Scheme, deployment)
		if err != nil {
			return err
		}
		cr.Status.Target = *deploymentRef

		// Execute the update of the status against the kubernetes API.
		err = r.MgrClient.Status().Update(ctx, &cr)
		if err != nil {
			return err
		}
	}

	return nil
}
