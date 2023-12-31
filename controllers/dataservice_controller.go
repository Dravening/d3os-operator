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
	"time"

	d3osoperatorv1 "d3os-operator/api/v1"
	"d3os-operator/resource/dataservice"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// DataServiceReconciler reconciles a DataService object
type DataServiceReconciler struct {
	client.Client
	DsBackendMap map[string]*d3osoperatorv1.DataServiceBackend
	Scheme       *runtime.Scheme
}

//+kubebuilder:rbac:groups=d3os-product.com.d3os,resources=dataservices,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=d3os-product.com.d3os,resources=dataservices/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=d3os-product.com.d3os,resources=dataservices/finalizers,verbs=update
//+kubebuilder:rbac:groups=apps,resources=deployments;statefulsets;daemonsets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=services;configmaps,verbs=get;list;watch;create;update;patch;delete

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
	dsInstance := &d3osoperatorv1.DataService{}

	// 通过客户端工具查询，查询条件是否有dsInstance存在
	err := r.Get(ctx, req.NamespacedName, dsInstance)
	if err != nil {
		// 如果没有实例，就返回空结果，这样外部就不再立即调用Reconcile方法了
		if errors.IsNotFound(err) {
			rLog.Info("dataservice instance not found, maybe removed")
			return reconcile.Result{}, nil
		}
		rLog.Error(err, "error getting dataservice instance")
		// 返回错误信息给外部
		return ctrl.Result{}, err
	}

	// 这是debug级别
	rLog.V(1).Info("got dsInstance, go on")

	// 调谐开始，首先要生成新的dataServiceBackend
	dsBackend := dsInstance.Spec.NewDSBackend(req)

	// 获取旧的dsBackend
	oldDsBackend, ok := r.DsBackendMap[req.String()]
	if !ok {
		oldDsBackend = dsBackend
	}
	r.DsBackendMap[req.String()] = dsBackend

	// 1.中间件 Mysql Uuc Eureka
	// Mysql
	if err = CheckExistsOrCreateMidBackend(ctx, r, oldDsBackend.Mysql, dsBackend.Mysql, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = CheckMidBackendStatus(ctx, r, dsBackend.Mysql, 5*time.Second); err != nil {
		return ctrl.Result{RequeueAfter: time.Second * 10}, err
	}
	if err = updateDsStatus(ctx, r, req.NamespacedName, dataservice.None, dataservice.Mysql); err != nil {
		return ctrl.Result{}, err
	}
	// Uuc
	if err = CheckExistsOrCreateMidBackend(ctx, r, oldDsBackend.Uuc, dsBackend.Uuc, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = CheckMidBackendStatus(ctx, r, dsBackend.Uuc, 5*time.Second); err != nil {
		return ctrl.Result{RequeueAfter: time.Second * 10}, err
	}
	if err = updateDsStatus(ctx, r, req.NamespacedName, dataservice.Mysql, dataservice.UUC); err != nil {
		return ctrl.Result{}, err
	}
	// Eureka
	if err = CheckExistsOrCreateMidBackend(ctx, r, oldDsBackend.Eureka, dsBackend.Eureka, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = CheckMidBackendStatus(ctx, r, dsBackend.Eureka, 5*time.Second); err != nil {
		return ctrl.Result{RequeueAfter: time.Second * 10}, err
	}
	if err = updateDsStatus(ctx, r, req.NamespacedName, dataservice.UUC, dataservice.Eureka); err != nil {
		return ctrl.Result{}, err
	}

	// 2.调谐服务 ApiManager Auth DsAdapter EsAdapter Eureka TrdAdapter GatewayMaster GatewayWeb Proxy
	// ApiManager
	if err = CheckExistsOrCreateAppBackend(ctx, r, oldDsBackend.ApiManager, dsBackend.ApiManager, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = CheckAppBackendStatus(ctx, r, dsBackend.ApiManager, 5*time.Second); err != nil {
		return ctrl.Result{RequeueAfter: time.Second * 10}, err
	}
	if err = updateDsStatus(ctx, r, req.NamespacedName, dataservice.Eureka, dataservice.ApiManager); err != nil {
		return ctrl.Result{}, err
	}
	// Auth
	if err = CheckExistsOrCreateAppBackend(ctx, r, oldDsBackend.Auth, dsBackend.Auth, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = CheckAppBackendStatus(ctx, r, dsBackend.Auth, 5*time.Second); err != nil {
		return ctrl.Result{RequeueAfter: time.Second * 10}, err
	}
	if err = updateDsStatus(ctx, r, req.NamespacedName, dataservice.ApiManager, dataservice.Auth); err != nil {
		return ctrl.Result{}, err
	}
	// DsAdapter
	if err = CheckExistsOrCreateAppBackend(ctx, r, oldDsBackend.DsAdapter, dsBackend.DsAdapter, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = CheckAppBackendStatus(ctx, r, dsBackend.DsAdapter, 5*time.Second); err != nil {
		return ctrl.Result{RequeueAfter: time.Second * 10}, err
	}
	if err = updateDsStatus(ctx, r, req.NamespacedName, dataservice.Auth, dataservice.DsAdapter); err != nil {
		return ctrl.Result{}, err
	}
	// EsAdapter
	if err = CheckExistsOrCreateAppBackend(ctx, r, oldDsBackend.EsAdapter, dsBackend.EsAdapter, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = CheckAppBackendStatus(ctx, r, dsBackend.EsAdapter, 5*time.Second); err != nil {
		return ctrl.Result{RequeueAfter: time.Second * 10}, err
	}
	if err = updateDsStatus(ctx, r, req.NamespacedName, dataservice.DsAdapter, dataservice.EsAdapter); err != nil {
		return ctrl.Result{}, err
	}
	// TrdAdapter
	if err = CheckExistsOrCreateAppBackend(ctx, r, oldDsBackend.TrdAdapter, dsBackend.TrdAdapter, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = CheckAppBackendStatus(ctx, r, dsBackend.TrdAdapter, 5*time.Second); err != nil {
		return ctrl.Result{RequeueAfter: time.Second * 10}, err
	}
	if err = updateDsStatus(ctx, r, req.NamespacedName, dataservice.EsAdapter, dataservice.TrdAdapter); err != nil {
		return ctrl.Result{}, err
	}
	// GatewayMaster
	if err = CheckExistsOrCreateAppBackend(ctx, r, oldDsBackend.GatewayMaster, dsBackend.GatewayMaster, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = CheckAppBackendStatus(ctx, r, dsBackend.GatewayMaster, 5*time.Second); err != nil {
		return ctrl.Result{RequeueAfter: time.Second * 10}, err
	}
	if err = updateDsStatus(ctx, r, req.NamespacedName, dataservice.TrdAdapter, dataservice.GatewayMaster); err != nil {
		return ctrl.Result{}, err
	}
	// GatewayWeb
	if err = CheckExistsOrCreateAppBackend(ctx, r, oldDsBackend.GatewayWeb, dsBackend.GatewayWeb, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = CheckAppBackendStatus(ctx, r, dsBackend.GatewayWeb, 5*time.Second); err != nil {
		return ctrl.Result{RequeueAfter: time.Second * 10}, err
	}
	if err = updateDsStatus(ctx, r, req.NamespacedName, dataservice.GatewayMaster, dataservice.GatewayWeb); err != nil {
		return ctrl.Result{}, err
	}
	// Proxy
	if err = CheckExistsOrCreateAppBackend(ctx, r, oldDsBackend.Proxy, dsBackend.Proxy, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = CheckAppBackendStatus(ctx, r, dsBackend.Proxy, 5*time.Second); err != nil {
		return ctrl.Result{RequeueAfter: time.Second * 10}, err
	}
	if err = updateDsStatus(ctx, r, req.NamespacedName, dataservice.GatewayWeb, dataservice.Proxy); err != nil {
		return ctrl.Result{}, err
	}

	// 3.查找并部署web服务
	if err = CheckExistsOrCreateAppBackend(ctx, r, oldDsBackend.Web, dsBackend.Web, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = CheckAppBackendStatus(ctx, r, dsBackend.Web, 5*time.Second); err != nil {
		return ctrl.Result{RequeueAfter: time.Second * 10}, err
	}
	if err = updateDsStatus(ctx, r, req.NamespacedName, dataservice.Proxy, dataservice.Web); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DataServiceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&d3osoperatorv1.DataService{}).
		Complete(r)
}
