package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os/exec"

	d3osoperatorv1 "d3os-operator/api/v1"
	"d3os-operator/resource/dataservice"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

var statusMap = map[dataservice.DSName]d3osoperatorv1.DsStatus{
	dataservice.None:   {Phase: d3osoperatorv1.DoneNone, Message: d3osoperatorv1.MessageNone},
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

func createDeploymentIfNotExists(ctx context.Context, r *DataServiceReconciler, newDeploy *appsv1.Deployment, dsInstance *d3osoperatorv1.DataService) error {
	if newDeploy == nil {
		return fmt.Errorf("spec deployment info doesn't exist, please check crd config")
	}
	objKey := types.NamespacedName{
		Name:      newDeploy.Name,
		Namespace: newDeploy.Namespace,
	}
	rLog := log.FromContext(ctx)
	objName := newDeploy.GetName()
	oldDeploy := &appsv1.Deployment{}
	err := r.Get(ctx, objKey, oldDeploy)
	if err != nil {
		if !errors.IsNotFound(err) {
			rLog.Error(err, fmt.Sprintf("getting deployment %s error", objName))
			return err
		}
		rLog.Info(fmt.Sprintf("deployment %s not found, creating...", objName))
		// 建立关联后，删除dataservice资源时就会将相应的obj也删除掉; 这一步会在obj上增加controllerRef标签
		if err = controllerutil.SetControllerReference(dsInstance, newDeploy, r.Scheme); err != nil {
			rLog.Error(err, fmt.Sprintf("set controller reference with deployment %s error", objName))
			return err
		}
		// 创建obj
		if err = r.Create(ctx, newDeploy); err != nil {
			rLog.Error(err, fmt.Sprintf("creating deployment %s error", objName))
			return err
		}
		return nil
	}
	// obj exists
	newDeploy.Spec.Template.ObjectMeta = metav1.ObjectMeta{}
	if equality.Semantic.DeepEqual(oldDeploy.Spec, newDeploy.Spec) {
		return nil
	}
	// need patch
	rLog.Info(fmt.Sprintf("即将update deployment %s", objName))
	patchBytes, err := json.Marshal(newDeploy)
	if err != nil {
		return err
	}
	if err = r.Patch(ctx, oldDeploy, client.RawPatch(types.MergePatchType, patchBytes)); err != nil {
		rLog.Error(err, fmt.Sprintf("updating deployment %s error", objName))
		return err
	}
	return nil
}

func checkDeploymentStatus(ctx context.Context, r *DataServiceReconciler, deploy *appsv1.Deployment) error {
	rLog := log.FromContext(ctx)
	if deploy == nil {
		return fmt.Errorf("spec deployment info doesn't exist, please check crd config")
	}
	deployNew := &appsv1.Deployment{}
	objKey := types.NamespacedName{
		Name:      deploy.Name,
		Namespace: deploy.Namespace,
	}
	objName := objKey.String()
	if err := r.Get(ctx, objKey, deployNew); err != nil {
		if !errors.IsNotFound(err) {
			rLog.Error(err, fmt.Sprintf("getting deployment %s error", objName))
			return err
		}
		rLog.Error(err, fmt.Sprintf("deployment %s not found, try reconcile", objName))
		return err
	}
	if deployNew.Status.AvailableReplicas < 1 {
		err := fmt.Errorf("deployment %s is not Ready yet, try reconcile", objName)
		return err
	}
	return nil
}

func createStatefulSetIfNotExists(ctx context.Context, r *DataServiceReconciler, newStatefulSet *appsv1.StatefulSet, dsInstance *d3osoperatorv1.DataService) error {
	if newStatefulSet == nil {
		return fmt.Errorf("spec statefulSet info doesn't exist, please check crd config")
	}
	objKey := types.NamespacedName{
		Name:      newStatefulSet.Name,
		Namespace: newStatefulSet.Namespace,
	}
	rLog := log.FromContext(ctx)
	objName := newStatefulSet.GetName()
	oldStatefulSet := &appsv1.StatefulSet{}
	err := r.Get(ctx, objKey, oldStatefulSet)
	if err != nil {
		if !errors.IsNotFound(err) {
			rLog.Error(err, fmt.Sprintf("getting statefulSet %s error", objName))
			return err
		}
		rLog.Info(fmt.Sprintf("statefulSet %s not found, creating...", objName))
		// 建立关联后，删除dataservice资源时就会将相应的obj也删除掉; 这一步会在obj上增加controllerRef标签
		if err = controllerutil.SetControllerReference(dsInstance, newStatefulSet, r.Scheme); err != nil {
			rLog.Error(err, fmt.Sprintf("set controller reference with statefulSet %s error", objName))
			return err
		}
		// 创建obj
		if err = r.Create(ctx, newStatefulSet); err != nil {
			rLog.Error(err, fmt.Sprintf("creating statefulSet %s error", objName))
			return err
		}
		return nil
	}
	// obj exists
	newStatefulSet.Spec.Template.ObjectMeta = metav1.ObjectMeta{}
	if equality.Semantic.DeepEqual(oldStatefulSet.Spec, newStatefulSet.Spec) {
		return nil
	}
	// need patch
	rLog.Info(fmt.Sprintf("即将update statefulSet %s", objName))
	patchBytes, err := json.Marshal(newStatefulSet)
	if err != nil {
		return err
	}
	if err = r.Patch(ctx, oldStatefulSet, client.RawPatch(types.MergePatchType, patchBytes)); err != nil {
		rLog.Error(err, fmt.Sprintf("updating statefulSet %s error", objName))
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

func createDaemonSetIfNotExists(ctx context.Context, r *DataServiceReconciler, newDaemonSet *appsv1.DaemonSet, dsInstance *d3osoperatorv1.DataService) error {
	if newDaemonSet == nil {
		return fmt.Errorf("spec DaemonSet info doesn't exist, please check crd config")
	}
	objKey := types.NamespacedName{
		Name:      newDaemonSet.Name,
		Namespace: newDaemonSet.Namespace,
	}
	rLog := log.FromContext(ctx)
	objName := newDaemonSet.GetName()
	oldDaemonSet := &appsv1.DaemonSet{}
	err := r.Get(ctx, objKey, oldDaemonSet)
	if err != nil {
		if !errors.IsNotFound(err) {
			rLog.Error(err, fmt.Sprintf("getting daemonSet %s error", objName))
			return err
		}
		rLog.Info(fmt.Sprintf("daemonSet %s not found, creating...", objName))
		// 建立关联后，删除dataservice资源时就会将相应的obj也删除掉; 这一步会在obj上增加controllerRef标签
		if err = controllerutil.SetControllerReference(dsInstance, newDaemonSet, r.Scheme); err != nil {
			rLog.Error(err, fmt.Sprintf("set controller reference with daemonSet %s error", objName))
			return err
		}
		// 创建obj
		if err = r.Create(ctx, newDaemonSet); err != nil {
			rLog.Error(err, fmt.Sprintf("creating daemonSet %s error", objName))
			return err
		}
		return nil
	}
	// obj exists
	newDaemonSet.Spec.Template.ObjectMeta = metav1.ObjectMeta{}
	if equality.Semantic.DeepEqual(oldDaemonSet.Spec, newDaemonSet.Spec) {
		return nil
	}
	// need patch
	rLog.Info(fmt.Sprintf("即将update daemonSet %s", objName))
	patchBytes, err := json.Marshal(newDaemonSet)
	if err != nil {
		return err
	}
	if err = r.Patch(ctx, oldDaemonSet, client.RawPatch(types.MergePatchType, patchBytes)); err != nil {
		rLog.Error(err, fmt.Sprintf("updating daemonSet %s error", objName))
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

func createServiceIfNotExists(ctx context.Context, r *DataServiceReconciler, newService *corev1.Service, dsInstance *d3osoperatorv1.DataService) error {
	if newService == nil {
		return fmt.Errorf("spec Service info doesn't exist, please check crd config")
	}
	objKey := types.NamespacedName{
		Name:      newService.Name,
		Namespace: newService.Namespace,
	}
	rLog := log.FromContext(ctx)
	objName := newService.GetName()
	oldService := &corev1.Service{}
	err := r.Get(ctx, objKey, oldService)
	if err != nil {
		if !errors.IsNotFound(err) {
			rLog.Error(err, fmt.Sprintf("getting service %s error", objName))
			return err
		}
		rLog.Info(fmt.Sprintf("service %s not found, creating...", objName))
		// 建立关联后，删除dataservice资源时就会将相应的obj也删除掉; 这一步会在obj上增加controllerRef标签
		if err = controllerutil.SetControllerReference(dsInstance, newService, r.Scheme); err != nil {
			rLog.Error(err, fmt.Sprintf("set controller reference with service %s error", objName))
			return err
		}
		// 创建obj
		if err = r.Create(ctx, newService); err != nil {
			rLog.Error(err, fmt.Sprintf("creating service %s error", objName))
			return err
		}
		return nil
	}
	// obj exists
	if equality.Semantic.DeepEqual(oldService.Spec, newService.Spec) {
		return nil
	}
	// need patch
	rLog.Info(fmt.Sprintf("即将update service %s", objName))
	patchBytes, err := json.Marshal(newService)
	if err != nil {
		return err
	}
	if err = r.Patch(ctx, oldService, client.RawPatch(types.MergePatchType, patchBytes)); err != nil {
		rLog.Error(err, fmt.Sprintf("updating service %s error", objName))
		return err
	}
	return nil
}

func createConfigMapIfNotExists(ctx context.Context, r *DataServiceReconciler, newConfigMap *corev1.ConfigMap, dsInstance *d3osoperatorv1.DataService) error {
	if newConfigMap == nil {
		return fmt.Errorf("spec ConfigMap info doesn't exist, please check crd config")
	}
	objKey := types.NamespacedName{
		Name:      newConfigMap.Name,
		Namespace: newConfigMap.Namespace,
	}
	rLog := log.FromContext(ctx)
	objName := newConfigMap.GetName()
	oldConfigMap := &corev1.ConfigMap{}
	err := r.Get(ctx, objKey, oldConfigMap)
	if err != nil {
		if !errors.IsNotFound(err) {
			rLog.Error(err, fmt.Sprintf("getting configMap %s error", objName))
			return err
		}
		rLog.Info(fmt.Sprintf("configMap %s not found, creating...", objName))
		// 建立关联后，删除dataservice资源时就会将相应的obj也删除掉; 这一步会在obj上增加controllerRef标签
		if err = controllerutil.SetControllerReference(dsInstance, newConfigMap, r.Scheme); err != nil {
			rLog.Error(err, fmt.Sprintf("set controller reference with configMap %s error", objName))
			return err
		}
		// 创建obj
		if err = r.Create(ctx, newConfigMap); err != nil {
			rLog.Error(err, fmt.Sprintf("creating configMap %s error", objName))
			return err
		}
		return nil
	}
	// obj exists
	if equality.Semantic.DeepEqual(oldConfigMap.Data, newConfigMap.Data) {
		return nil
	}
	// need patch
	rLog.Info(fmt.Sprintf("即将update configMap %s", objName))
	patchBytes, err := json.Marshal(newConfigMap)
	if err != nil {
		return err
	}
	if err = r.Patch(ctx, oldConfigMap, client.RawPatch(types.MergePatchType, patchBytes)); err != nil {
		rLog.Error(err, fmt.Sprintf("updating configMap %s error", objName))
		return err
	}
	return nil
}

func updateDsStatus(ctx context.Context, r *DataServiceReconciler, key client.ObjectKey, oldDsName, dsName dataservice.DSName) error {
	rLog := log.FromContext(ctx)
	// 先查询当前dataservice实例的状态
	dsInstance := &d3osoperatorv1.DataService{}
	err := r.Get(ctx, key, dsInstance)
	if err != nil {
		if errors.IsNotFound(err) {
			return err
		}
		return err
	}
	if dsInstance.Status.Phase == statusMap[oldDsName].Phase || dsInstance.Status.Phase == "" {
		// 如果状态是前置状态，更新到当前状态
		dsInstance.Status.Phase = statusMap[dsName].Phase
		dsInstance.Status.Message = statusMap[dsName].Message
		rLog.Info(fmt.Sprintf("即将update dataservice %s的状态至 %s", dsInstance.GetName(), statusMap[dsName].Phase))
		if err = r.Status().Update(ctx, dsInstance); err != nil {
			return err
		}
	}
	return nil
}
