/*
**
// Controller Logic
**
*/
package controller

import (
	"context"

	"github.com/go-logr/logr"
	v1alpha1 "github.com/orenr2301/KubeTag/pkg/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	// "sigs.k8s.io/controller-runtime/pkg/log"
)


type NsLabelSetReconciler struct {
	client.Client
	Log logr.Logger
}


func (r *NsLabelSetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	var nslabelset v1alpha1.NsLabelSet
	if err := r.Get(ctx, req.NamespacedName, &nslabelset); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Get the target namespace
	var ns corev1.Namespace
	if err := r.Get(ctx, types.NamespacedName{Name: nslabelset.Spec.Namespace}, &ns); err != nil {
		nslabelset.Status.Applied = false
		nslabelset.Status.Message = "Namespace not found"
		_ = r.Status().Update(ctx, &nslabelset)
		return ctrl.Result{}, err
	}

	// Patch labels
	if ns.Labels == nil {
		ns.Labels = map[string]string{}
	}
	for k, v := range nslabelset.Spec.Labels {
		ns.Labels[k] = v
	}
	if err := r.Update(ctx, &ns); err != nil {
		nslabelset.Status.Applied = false
		nslabelset.Status.Message = "Failed to update labels"
		_ = r.Status().Update(ctx, &nslabelset)
		return ctrl.Result{}, err
	}

	// Update status to success
	nslabelset.Status.Applied = true
	nslabelset.Status.Message = "Labels applied successfully"
	_ = r.Status().Update(ctx, &nslabelset)
	return ctrl.Result{}, nil
}

///Continue tomorrow
