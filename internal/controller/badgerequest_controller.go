package controller

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/exp/rand"
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
	_ = log.FromContext(ctx)

	// resty := resty.New()
	// resp, err := resty.R().Get("")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return ctrl.Result{}, err
	// }
	// fmt.Println("Response:", resp.String())

	var badgeRequest badgesv1.BadgeRequest
	if err := r.Get(ctx, req.NamespacedName, &badgeRequest); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	rand.Seed(uint64(time.Now().UnixNano()))
	num1 := rand.Intn(301) + 200
	num2 := rand.Intn(301) + 200

	badgeRequest.Status.OutputURL = fmt.Sprintf("https://picsum.photos/%d/%d", num1, num2)
	badgeRequest.Status.Status = "Completed"

	if err := r.Status().Update(ctx, &badgeRequest); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *BadgeRequestReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&badgesv1.BadgeRequest{}).
		Named("badgerequest").
		Complete(r)
}
