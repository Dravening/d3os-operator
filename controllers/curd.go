package controllers

import (
	"context"
	"fmt"
	"os/exec"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
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

func createIfNotExist(r *DataServiceReconciler, ctx context.Context, objKey types.NamespacedName, objType, obj client.Object) error {
	rLog := log.FromContext(ctx)
	err := r.Get(ctx, objKey, objType)
	if err != nil {
		if !errors.IsNotFound(err) {
			rLog.Error(err, fmt.Sprintf("getting obj %s error", objKey.String()))
			return err
		}
		rLog.Info(fmt.Sprintf("obj %s not found, creating...", objKey.String()))
		// 创建service
		if err = r.Create(ctx, obj); err != nil {
			rLog.Error(err, fmt.Sprintf("creating obj %s error", objKey.String()))
			return err
		}
	}
	// obj exists, do nothing
	return nil
}

func createDeploymentIfNotExists(ctx context.Context, r *DataServiceReconciler, deploy *appsv1.Deployment) error {
	if deploy == nil {
		return fmt.Errorf("spec Deployment info doesn't exist, please check crd config")
	}
	deployTemp := &appsv1.Deployment{}
	objKey := types.NamespacedName{
		Name:      deploy.Name,
		Namespace: deploy.Namespace,
	}
	if err := createIfNotExist(r, ctx, objKey, deployTemp, deploy); err != nil {
		return err
	}
	return nil
}

func createStatefulSetIfNotExists(ctx context.Context, r *DataServiceReconciler, statefulSet *appsv1.StatefulSet) error {
	if statefulSet == nil {
		return fmt.Errorf("spec StatefulSet info doesn't exist, please check crd config")
	}
	stateTemp := &appsv1.StatefulSet{}
	objKey := types.NamespacedName{
		Name:      statefulSet.Name,
		Namespace: statefulSet.Namespace,
	}
	if err := createIfNotExist(r, ctx, objKey, stateTemp, statefulSet); err != nil {
		return err
	}
	return nil
}

func createDaemonSetIfNotExists(ctx context.Context, r *DataServiceReconciler, daemonSet *appsv1.DaemonSet) error {
	if daemonSet == nil {
		return fmt.Errorf("spec DaemonSet info doesn't exist, please check crd config")
	}
	daemonTemp := &appsv1.DaemonSet{}
	objKey := types.NamespacedName{
		Name:      daemonSet.Name,
		Namespace: daemonSet.Namespace,
	}
	if err := createIfNotExist(r, ctx, objKey, daemonTemp, daemonSet); err != nil {
		return err
	}
	return nil
}

func createServiceIfNotExists(ctx context.Context, r *DataServiceReconciler, service *corev1.Service) error {
	if service == nil {
		return fmt.Errorf("spec Service info doesn't exist, please check crd config")
	}
	serviceTemp := &corev1.Service{}
	objKey := types.NamespacedName{
		Name:      service.Name,
		Namespace: service.Namespace,
	}
	if err := createIfNotExist(r, ctx, objKey, serviceTemp, service); err != nil {
		return err
	}
	return nil
}

func createConfigMapIfNotExists(ctx context.Context, r *DataServiceReconciler, configMap *corev1.ConfigMap) error {
	if configMap == nil {
		return fmt.Errorf("spec ConfigMap info doesn't exist, please check crd config")
	}
	configMapTemp := &corev1.ConfigMap{}
	objKey := types.NamespacedName{
		Name:      configMap.Name,
		Namespace: configMap.Namespace,
	}
	if err := createIfNotExist(r, ctx, objKey, configMapTemp, configMap); err != nil {
		return err
	}
	return nil
}
