package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"

	d3osoperatorv1 "d3os-operator/api/v1"
	"d3os-operator/resource/dataservice"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

var statusMap = map[dataservice.DSName]d3osoperatorv1.DsStatus{
	dataservice.Mysql:  {Phase: d3osoperatorv1.MysqlDone, Message: d3osoperatorv1.MessageMysql},
	dataservice.UUC:    {Phase: d3osoperatorv1.UucDone, Message: d3osoperatorv1.MessageUuc},
	dataservice.Eureka: {Phase: d3osoperatorv1.EurekaDone, Message: d3osoperatorv1.MessageEureka},

	dataservice.ApiManager:    {Phase: d3osoperatorv1.ApiManagerDone, Message: d3osoperatorv1.MessageApiManager},
	dataservice.Auth:          {Phase: d3osoperatorv1.AuthDone, Message: d3osoperatorv1.MessageAuth},
	dataservice.DsAdapter:     {Phase: d3osoperatorv1.DsAdapterDone, Message: d3osoperatorv1.MessageDsAdapter},
	dataservice.EsAdapter:     {Phase: d3osoperatorv1.EsAdapterDone, Message: d3osoperatorv1.MessageEsAdapter},
	dataservice.GatewayMaster: {Phase: d3osoperatorv1.GatewayMasterDone, Message: d3osoperatorv1.MessageGatewayMaster},
	dataservice.GatewayWeb:    {Phase: d3osoperatorv1.GatewayWebDone, Message: d3osoperatorv1.MessageGatewayWeb},
	dataservice.Proxy:         {Phase: d3osoperatorv1.ProxyDone, Message: d3osoperatorv1.MessageProxy},
	dataservice.TrdAdapter:    {Phase: d3osoperatorv1.TrdAdapterDone, Message: d3osoperatorv1.MessageTrdAdapter},
	dataservice.Web:           {Phase: d3osoperatorv1.WebDone, Message: d3osoperatorv1.MessageWeb},
}

func cmdCall(command string) error {
	cmd := exec.Command("/bin/sh", "-c", command)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("/bin/sh执行[%s]命令时出错: %s\n", command, err)
	}
	return nil
}

func createOrUpdate(r *DataServiceReconciler, ctx context.Context, objKey types.NamespacedName, objType, obj client.Object, dsInstance *d3osoperatorv1.DataService) error {
	rLog := log.FromContext(ctx)
	objName := obj.GetName()
	err := r.Get(ctx, objKey, objType)
	if err != nil {
		if !errors.IsNotFound(err) {
			rLog.Error(err, fmt.Sprintf("getting obj %s error", objName))
			return err
		}
		rLog.Info(fmt.Sprintf("obj %s not found, creating...", objName))

		// 建立关联后，删除dataservice资源时就会将相应的obj也删除掉; 这一步会在obj上增加controllerRef标签
		if err = controllerutil.SetControllerReference(dsInstance, obj, r.Scheme); err != nil {
			rLog.Error(err, fmt.Sprintf("set controller reference with %s error", objName))
			return err
		}

		// 创建obj
		if err = r.Create(ctx, obj); err != nil {
			rLog.Error(err, fmt.Sprintf("creating obj %s error", objName))
			return err
		}
		return nil
	}
	// obj exists, patch
	rLog.Info(fmt.Sprintf("即将update obj %s", objName))
	patchBytes, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	if err = r.Patch(ctx, objType, client.RawPatch(types.MergePatchType, patchBytes)); err != nil {
		rLog.Error(err, fmt.Sprintf("updating obj %s error", obj.GetName()))
		return err
	}
	return nil
}

func createDeploymentIfNotExists(ctx context.Context, r *DataServiceReconciler, deploy *appsv1.Deployment, dsInstance *d3osoperatorv1.DataService) error {
	if deploy == nil {
		return fmt.Errorf("spec Deployment info doesn't exist, please check crd config")
	}
	deployTemp := &appsv1.Deployment{}
	objKey := types.NamespacedName{
		Name:      deploy.Name,
		Namespace: deploy.Namespace,
	}
	if err := createOrUpdate(r, ctx, objKey, deployTemp, deploy, dsInstance); err != nil {
		return err
	}
	return nil
}

func checkDeploymentStatus(ctx context.Context, r *DataServiceReconciler, deploy *appsv1.Deployment) error {
	rLog := log.FromContext(ctx)
	if deploy == nil {
		return fmt.Errorf("spec Deployment info doesn't exist, please check crd config")
	}
	deployNew := &appsv1.Deployment{}
	objKey := types.NamespacedName{
		Name:      deploy.Name,
		Namespace: deploy.Namespace,
	}
	objName := objKey.String()
	if err := r.Get(ctx, objKey, deployNew); err != nil {
		if !errors.IsNotFound(err) {
			rLog.Error(err, fmt.Sprintf("getting obj %s error", objName))
			return err
		}
		rLog.Error(err, fmt.Sprintf("obj %s not found, try reconcile", objName))
		return err
	}
	if deployNew.Status.AvailableReplicas < 1 {
		err := fmt.Errorf("deployment %s is not Ready yet, try reconcile", objName)
		return err
	}
	return nil
}

func createStatefulSetIfNotExists(ctx context.Context, r *DataServiceReconciler, statefulSet *appsv1.StatefulSet, dsInstance *d3osoperatorv1.DataService) error {
	if statefulSet == nil {
		return fmt.Errorf("spec StatefulSet info doesn't exist, please check crd config")
	}
	stateTemp := &appsv1.StatefulSet{}
	objKey := types.NamespacedName{
		Name:      statefulSet.Name,
		Namespace: statefulSet.Namespace,
	}
	if err := createOrUpdate(r, ctx, objKey, stateTemp, statefulSet, dsInstance); err != nil {
		return err
	}
	return nil
}

func checkStatefulSetStatus(ctx context.Context, r *DataServiceReconciler, statefulSet *appsv1.StatefulSet) error {
	rLog := log.FromContext(ctx)
	if statefulSet == nil {
		return fmt.Errorf("spec StatefulSet info doesn't exist, please check crd config")
	}
	statefulSetNew := &appsv1.StatefulSet{}
	objKey := types.NamespacedName{
		Name:      statefulSet.Name,
		Namespace: statefulSet.Namespace,
	}
	objName := objKey.String()
	if err := r.Get(ctx, objKey, statefulSetNew); err != nil {
		if !errors.IsNotFound(err) {
			rLog.Error(err, fmt.Sprintf("getting obj %s error", objName))
			return err
		}
		rLog.Error(err, fmt.Sprintf("obj %s not found, try reconcile", objName))
		return err
	}
	if statefulSetNew.Status.ReadyReplicas < 1 {
		err := fmt.Errorf("statefulSet %s is not Ready yet, try reconcile", objName)
		return err
	}
	return nil
}

func createDaemonSetIfNotExists(ctx context.Context, r *DataServiceReconciler, daemonSet *appsv1.DaemonSet, dsInstance *d3osoperatorv1.DataService) error {
	if daemonSet == nil {
		return fmt.Errorf("spec DaemonSet info doesn't exist, please check crd config")
	}
	daemonTemp := &appsv1.DaemonSet{}
	objKey := types.NamespacedName{
		Name:      daemonSet.Name,
		Namespace: daemonSet.Namespace,
	}
	if err := createOrUpdate(r, ctx, objKey, daemonTemp, daemonSet, dsInstance); err != nil {
		return err
	}
	return nil
}

func checkDaemonSetStatus(ctx context.Context, r *DataServiceReconciler, daemonSet *appsv1.DaemonSet) error {
	rLog := log.FromContext(ctx)
	if daemonSet == nil {
		return fmt.Errorf("spec daemonSet info doesn't exist, please check crd config")
	}
	daemonSetNew := &appsv1.DaemonSet{}
	objKey := types.NamespacedName{
		Name:      daemonSet.Name,
		Namespace: daemonSet.Namespace,
	}
	objName := objKey.String()
	if err := r.Get(ctx, objKey, daemonSetNew); err != nil {
		if !errors.IsNotFound(err) {
			rLog.Error(err, fmt.Sprintf("getting obj %s error", objName))
			return err
		}
		rLog.Error(err, fmt.Sprintf("obj %s not found, try reconcile", objName))
		return err
	}
	if daemonSetNew.Status.NumberUnavailable > 0 {
		err := fmt.Errorf("daemonSet %s is not Ready yet, try reconcile", objName)
		return err
	}
	return nil
}

func createServiceIfNotExists(ctx context.Context, r *DataServiceReconciler, service *corev1.Service, dsInstance *d3osoperatorv1.DataService) error {
	if service == nil {
		return fmt.Errorf("spec Service info doesn't exist, please check crd config")
	}
	serviceTemp := &corev1.Service{}
	objKey := types.NamespacedName{
		Name:      service.Name,
		Namespace: service.Namespace,
	}
	if err := createOrUpdate(r, ctx, objKey, serviceTemp, service, dsInstance); err != nil {
		return err
	}
	return nil
}

func createConfigMapIfNotExists(ctx context.Context, r *DataServiceReconciler, configMap *corev1.ConfigMap, dsInstance *d3osoperatorv1.DataService) error {
	if configMap == nil {
		return fmt.Errorf("spec ConfigMap info doesn't exist, please check crd config")
	}
	configMapTemp := &corev1.ConfigMap{}
	objKey := types.NamespacedName{
		Name:      configMap.Name,
		Namespace: configMap.Namespace,
	}
	if err := createOrUpdate(r, ctx, objKey, configMapTemp, configMap, dsInstance); err != nil {
		return err
	}
	return nil
}

func updateDsStatus(ctx context.Context, r *DataServiceReconciler, key client.ObjectKey, dsName dataservice.DSName) error {
	// 先查询当前dataservice实例的状态
	dsInstance := &d3osoperatorv1.DataService{}
	err := r.Get(ctx, key, dsInstance)
	if err != nil {
		if errors.IsNotFound(err) {
			return err
		}
		return err
	}
	// 直接更新dsInstance的状态
	dsInstance.Status.Phase = statusMap[dsName].Phase
	dsInstance.Status.Message = statusMap[dsName].Message
	if err = r.Status().Update(ctx, dsInstance); err != nil {
		return err
	}
	return nil
}
