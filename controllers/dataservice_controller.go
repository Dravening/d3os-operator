/*
Copyright 2023.

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

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	d3osproductv1 "d3os-operator/api/v1"
)

// DataServiceReconciler reconciles a DataService object
type DataServiceReconciler struct {
	client.Client
	DsBackend *d3osproductv1.DataServiceBackend
	Scheme    *runtime.Scheme
}

//+kubebuilder:rbac:groups=d3os-product.com.d3os,resources=dataservices,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=d3os-product.com.d3os,resources=dataservices/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=d3os-product.com.d3os,resources=dataservices/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments;statefulSets;daemonSets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DataService object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *DataServiceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	rLog := log.FromContext(ctx)
	rLog.Info("start reconcile logic")

	// 实例化数据结构
	dsInstance := &d3osproductv1.DataService{}

	// 通过客户端工具查询，查询条件是否有dsInstance存在
	err := r.Get(ctx, req.NamespacedName, dsInstance)
	if err != nil {
		// 如果没有实例，就返回空结果，这样外部就不再立即调用Reconcile方法了
		if errors.IsNotFound(err) {
			rLog.Info("dataservice instance %s.%s not found, maybe removed", req.Name, req.Namespace)
			return reconcile.Result{}, nil
		}
		rLog.Error(err, "error getting dataservice instance %s.%s", req.Name, req.Namespace)
		// 返回错误信息给外部
		return ctrl.Result{}, err
	}

	// 调谐开始，首先要生成新的dataServiceBackend
	// oldDSBackend := r.DsBackend
	r.DsBackend = dsInstance.Spec.NewDSBackend(req)
	// todo: 比较新DsBackend与oldDSBackend

	// 1.查找中间件实例是否存在	Mysql Uuc
	if CheckExistsOrCreateMidBackend(ctx, r, req, r.DsBackend.Mysql) != nil {
		return ctrl.Result{}, err
	}
	if CheckExistsOrCreateMidBackend(ctx, r, req, r.DsBackend.Uuc) != nil {
		return ctrl.Result{}, err
	}

	// 2.调谐服务 ApiManager Auth DsAdapter EsAdapter Eureka TrdAdapter GatewayMaster GatewayWeb Proxy
	// ApiManager
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.ApiManager) != nil {
		return ctrl.Result{}, err
	}
	// Auth
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.Auth) != nil {
		return ctrl.Result{}, err
	}
	// DsAdapter
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.DsAdapter) != nil {
		return ctrl.Result{}, err
	}
	// EsAdapter
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.EsAdapter) != nil {
		return ctrl.Result{}, err
	}
	// TrdAdapter
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.TrdAdapter) != nil {
		return ctrl.Result{}, err
	}
	// GatewayMaster
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.GatewayMaster) != nil {
		return ctrl.Result{}, err
	}
	// GatewayWeb
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.GatewayWeb) != nil {
		return ctrl.Result{}, err
	}
	// Proxy
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.Proxy) != nil {
		return ctrl.Result{}, err
	}

	// 3.查找并部署web服务
	if CheckExistsOrCreateSvcBackend(ctx, r, req, r.DsBackend.Web) != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DataServiceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&d3osproductv1.DataService{}).
		Complete(r)
}
