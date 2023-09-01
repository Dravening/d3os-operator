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
	"fmt"

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
	// DsBackendMap map[string]*d3osoperatorv1.DataServiceBackend
	Scheme *runtime.Scheme
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
	//_, ok := r.DsBackendMap[req.String()]
	//if ok {
	//	// what should I do？
	//} else {
	//	// what should I do？
	//}
	//r.DsBackendMap[req.String()] = dsBackend

	// 目前并不准备比较dsBackend的各项指标；只需要实现kubectl update就好了
	// 1.查找中间件实例是否存在	Mysql Uuc
	if err = CheckExistsOrCreateMidBackend(ctx, r, dsBackend.Mysql, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = updateStatus(ctx, r, dataservice.Mysql, req.NamespacedName); err != nil {
		rLog.Error(err, fmt.Sprintf("error updating dataservice status after create %s", dataservice.Mysql))
		return ctrl.Result{}, err
	}
	if err = CheckExistsOrCreateMidBackend(ctx, r, dsBackend.Uuc, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = updateStatus(ctx, r, dataservice.UUC, req.NamespacedName); err != nil {
		rLog.Error(err, fmt.Sprintf("error updating dataservice status after create %s", dataservice.UUC))
		return ctrl.Result{}, err
	}
	if err = CheckExistsOrCreateMidBackend(ctx, r, dsBackend.Eureka, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = updateStatus(ctx, r, dataservice.Eureka, req.NamespacedName); err != nil {
		rLog.Error(err, fmt.Sprintf("error updating dataservice status after create %s", dataservice.Eureka))
		return ctrl.Result{}, err
	}

	// 2.调谐服务 ApiManager Auth DsAdapter EsAdapter Eureka TrdAdapter GatewayMaster GatewayWeb Proxy
	// ApiManager
	if err = CheckExistsOrCreateSvcBackend(ctx, r, dsBackend.ApiManager, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = updateStatus(ctx, r, dataservice.ApiManager, req.NamespacedName); err != nil {
		rLog.Error(err, fmt.Sprintf("error updating dataservice status after create %s", dataservice.ApiManager))
		return ctrl.Result{}, err
	}
	// Auth
	if err = CheckExistsOrCreateSvcBackend(ctx, r, dsBackend.Auth, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = updateStatus(ctx, r, dataservice.Auth, req.NamespacedName); err != nil {
		rLog.Error(err, fmt.Sprintf("error updating dataservice status after create %s", dataservice.Auth))
		return ctrl.Result{}, err
	}
	// DsAdapter
	if err = CheckExistsOrCreateSvcBackend(ctx, r, dsBackend.DsAdapter, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = updateStatus(ctx, r, dataservice.DsAdapter, req.NamespacedName); err != nil {
		rLog.Error(err, fmt.Sprintf("error updating dataservice status after create %s", dataservice.DsAdapter))
		return ctrl.Result{}, err
	}
	// EsAdapter
	if err = CheckExistsOrCreateSvcBackend(ctx, r, dsBackend.EsAdapter, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = updateStatus(ctx, r, dataservice.EsAdapter, req.NamespacedName); err != nil {
		rLog.Error(err, fmt.Sprintf("error updating dataservice status after create %s", dataservice.EsAdapter))
		return ctrl.Result{}, err
	}
	// TrdAdapter
	if err = CheckExistsOrCreateSvcBackend(ctx, r, dsBackend.TrdAdapter, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = updateStatus(ctx, r, dataservice.TrdAdapter, req.NamespacedName); err != nil {
		rLog.Error(err, fmt.Sprintf("error updating dataservice status after create %s", dataservice.TrdAdapter))
		return ctrl.Result{}, err
	}
	// GatewayMaster
	if err = CheckExistsOrCreateSvcBackend(ctx, r, dsBackend.GatewayMaster, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = updateStatus(ctx, r, dataservice.GatewayMaster, req.NamespacedName); err != nil {
		rLog.Error(err, fmt.Sprintf("error updating dataservice status after create %s", dataservice.GatewayMaster))
		return ctrl.Result{}, err
	}
	// GatewayWeb
	if err = CheckExistsOrCreateSvcBackend(ctx, r, dsBackend.GatewayWeb, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = updateStatus(ctx, r, dataservice.GatewayWeb, req.NamespacedName); err != nil {
		rLog.Error(err, fmt.Sprintf("error updating dataservice status after create %s", dataservice.GatewayWeb))
		return ctrl.Result{}, err
	}
	// Proxy
	if err = CheckExistsOrCreateSvcBackend(ctx, r, dsBackend.Proxy, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = updateStatus(ctx, r, dataservice.Proxy, req.NamespacedName); err != nil {
		rLog.Error(err, fmt.Sprintf("error updating dataservice status after create %s", dataservice.Proxy))
		return ctrl.Result{}, err
	}

	// 3.查找并部署web服务
	if err = CheckExistsOrCreateSvcBackend(ctx, r, dsBackend.Web, dsInstance); err != nil {
		return ctrl.Result{}, err
	}
	if err = updateStatus(ctx, r, dataservice.Web, req.NamespacedName); err != nil {
		rLog.Error(err, fmt.Sprintf("error updating dataservice status after create %s", dataservice.Web))
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
