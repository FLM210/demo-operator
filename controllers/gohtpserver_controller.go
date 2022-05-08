/*
Copyright 2022 dlbppx.

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

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1alpha1 "github.com/FLM210/demo-operator/api/v1alpha1"
)

// GohtpserverReconciler reconciles a Gohtpserver object
type GohtpserverReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=apps.dlbppx.online,resources=gohtpservers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps.dlbppx.online,resources=gohtpservers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apps.dlbppx.online,resources=gohtpservers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Gohtpserver object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *GohtpserverReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	var Gohttpserver appsv1alpha1.Gohtpserver
	if err := r.Get(ctx, req.NamespacedName, &Gohttpserver); err != nil {
		// EtcdCluster was gone，skip
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// CreateOrUpdate Service
	var svc corev1.Service
	svc.Name = Gohttpserver.Name
	svc.Namespace = Gohttpserver.Namespace
	or, err := ctrl.CreateOrUpdate(ctx, r.Client, &svc, func() error {
		MutateSvc(&Gohttpserver, &svc)
		return controllerutil.SetControllerReference(&Gohttpserver, &svc, r.Scheme)
	})
	if err != nil {
		return ctrl.Result{}, err
	}

	log.Log.Info("Create Or Update gohtpserver", "Service", or)

	// CreateOrUpdate Deployment
	var dep appsv1.Deployment
	dep.Name = Gohttpserver.Spec.Name
	dep.Namespace = Gohttpserver.Spec.Namespace
	or, err = ctrl.CreateOrUpdate(ctx, r.Client, &dep, func() error {
		MutateDeployment(&Gohttpserver, &dep)
		return controllerutil.SetControllerReference(&Gohttpserver, &dep, r.Scheme)
	})
	if err != nil {
		return ctrl.Result{}, err
	}
	log.Log.Info("Create Or Update gohtpserver", "deployment", or)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *GohtpserverReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1alpha1.Gohtpserver{}).
		Complete(r)
}
