/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	codiv1alpha1 "github.com/giantswarm/codimd-operator/api/v1alpha1"
	"github.com/giantswarm/codimd-operator/hack"
)

const (
	finalizerName = "codimd.workshop.giantswarm.io"
)

// CodiMDReconciler reconciles a CodiMD object
type CodiMDReconciler struct {
	Log              logr.Logger
	MgrEventRecorder record.EventRecorder
	MgrClient        client.Client
	Scheme           *runtime.Scheme
}

// +kubebuilder:rbac:groups=deploy.workshop.giantswarm.io,resources=codimds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=deploy.workshop.giantswarm.io,resources=codimds/status,verbs=get;update;patch

func (r *CodiMDReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	logger := r.Log.WithValues("codimd", req.NamespacedName)

	var cr codiv1alpha1.CodiMD
	if err := r.MgrClient.Get(ctx, req.NamespacedName, &cr); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if cr.ObjectMeta.DeletionTimestamp.IsZero() {
		// Add our finalizer if it is not already present.
		if !containsString(cr.ObjectMeta.Finalizers, finalizerName) {
			cr.ObjectMeta.Finalizers = append(cr.ObjectMeta.Finalizers, finalizerName)
			if err := r.MgrClient.Update(ctx, &cr); err != nil {
				return ctrl.Result{}, microerror.Mask(err)
			}
		}

		// Try to get a k8s deployment from the remote URL.
		deployment, err := hack.GetDeployment(logger, cr.Spec.URL)
		if err != nil {
			return ctrl.Result{}, microerror.Mask(err)
		}

		// Call the create func to handle creation.
		err = r.create(ctx, cr, deployment)
		if err != nil {
			return ctrl.Result{}, microerror.Mask(err)
		}
	} else {
		// Call the delete func to handle deletion.
		err := r.delete(ctx, cr)
		if err != nil {
			return ctrl.Result{}, microerror.Mask(err)
		}

		// Remove our finalizer if deletion completed successfully.
		cr.ObjectMeta.Finalizers = removeString(cr.ObjectMeta.Finalizers, finalizerName)
		if err := r.MgrClient.Update(ctx, &cr); err != nil {
			return ctrl.Result{}, microerror.Mask(err)
		}
	}

	return ctrl.Result{}, nil
}

func (r *CodiMDReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&codiv1alpha1.CodiMD{}).
		Complete(r)
}

func containsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func removeString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}
