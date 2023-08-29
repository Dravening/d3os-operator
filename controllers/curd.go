package controllers

import (
	"context"
	"fmt"
	"os/exec"

	d3osoperatorv1 "d3os-operator/api/v1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

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
	err := r.Get(ctx, objKey, objType)
	if err != nil {
		if !errors.IsNotFound(err) {
			rLog.Error(err, fmt.Sprintf("getting obj %s error", obj.GetName()))
			return err
		}
		rLog.Info(fmt.Sprintf("obj %s not found, creating...", obj.GetName()))

		// 建立关联后，删除dataservice资源时就会将相应的obj也删除掉; 这一步会在obj上增加controllerRef标签
		if err = controllerutil.SetControllerReference(dsInstance, obj, r.Scheme); err != nil {
			rLog.Error(err, fmt.Sprintf("set controller reference with %s error", obj.GetName()))
			return err
		}

		// 创建obj
		if err = r.Create(ctx, obj); err != nil {
			rLog.Error(err, fmt.Sprintf("creating obj %s error", obj.GetName()))
			return err
		}
		return nil
	}
	// obj exists, update
	rLog.Info(fmt.Sprintf("即将update obj,其内容%v", obj))
	if err = r.Update(ctx, obj); err != nil {
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
