package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	badgesv1 "github.com/BarelyAnXer/badge-generator/api/v1"
)

type BadgeRequestReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *BadgeRequestReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var badgeRequest badgesv1.BadgeRequest
	if err := r.Get(ctx, req.NamespacedName, &badgeRequest); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	generatedURL := fmt.Sprintf("https://example.com/badges/%s", badgeRequest.Spec.Title)
	badgeRequest.Status.OutputURL = generatedURL
	badgeRequest.Status.Status = "Completed"

	if err := r.Status().Update(ctx, &badgeRequest); err != nil {
		return ctrl.Result{}, err
	}

	log.Info("Badge generated", "Title", badgeRequest.Spec.Title, "URL", generatedURL)
	return ctrl.Result{}, nil
}

func (r *BadgeRequestReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&badgesv1.BadgeRequest{}).
		Named("badgerequest").
		Complete(r)
}
