package controllers

import (
	"context"
	"fmt"
	"os/exec"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func cmdCall(command string) error {
	cmd := exec.Command("/bin/bash", "-c", command)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("/bin/bash执行[%s]命令时出错: %s\n", command, err)
	}
	return nil
}

func createDeploymentIfNotExists(ctx context.Context, r *DataServiceReconciler, deploy *appsv1.Deployment) error {
	if deploy == nil {
		return fmt.Errorf("spec Deployment info doesn't exist, please check crd config")
	}
	rLog := log.FromContext(ctx)
	deployTemp := &appsv1.Deployment{}
	objKey := types.NamespacedName{
		Name:      deploy.Name,
		Namespace: deploy.Namespace,
	}
	err := r.Get(ctx, objKey, deployTemp)
	if err != nil {
		if errors.IsNotFound(err) {
			rLog.Info(fmt.Sprintf("deployment %s not found", objKey.String()))
			return nil
		}
		rLog.Error(err, fmt.Sprintf("getting deployment %s error", objKey.String()))
		return err
	}
	// 创建deployment
	rLog.Info(fmt.Sprintf("creating deployment %s error", objKey.String()))
	if err = r.Create(ctx, deploy); err != nil {
		rLog.Error(err, fmt.Sprintf("creating deployment %s", objKey.String()))
		return err
	}
	return nil
}

func createStatefulSetIfNotExists(ctx context.Context, r *DataServiceReconciler, statefulSet *appsv1.StatefulSet) error {
	if statefulSet == nil {
		return fmt.Errorf("spec StatefulSet info doesn't exist, please check crd config")
	}
	rLog := log.FromContext(ctx)
	stateTemp := &appsv1.StatefulSet{}
	objKey := types.NamespacedName{
		Name:      statefulSet.Name,
		Namespace: statefulSet.Namespace,
	}
	err := r.Get(ctx, objKey, stateTemp)
	if err != nil {
		if errors.IsNotFound(err) {
			rLog.Info(fmt.Sprintf("statefulSet %s not found", objKey.String()))
			return nil
		}
		rLog.Error(err, fmt.Sprintf("getting statefulSet %s error", objKey.String()))
		return err
	}
	// 创建statefulSet
	rLog.Info(fmt.Sprintf("creating statefulSet %s error", objKey.String()))
	if err = r.Create(ctx, statefulSet); err != nil {
		rLog.Error(err, fmt.Sprintf("creating statefulSet %s", objKey.String()))
		return err
	}
	return nil
}

func createDaemonSetIfNotExists(ctx context.Context, r *DataServiceReconciler, daemonSet *appsv1.DaemonSet) error {
	if daemonSet == nil {
		return fmt.Errorf("spec DaemonSet info doesn't exist, please check crd config")
	}
	rLog := log.FromContext(ctx)
	daemonTemp := &appsv1.DaemonSet{}
	objKey := types.NamespacedName{
		Name:      daemonSet.Name,
		Namespace: daemonSet.Namespace,
	}
	err := r.Get(ctx, objKey, daemonTemp)
	if err != nil {
		if errors.IsNotFound(err) {
			rLog.Info(fmt.Sprintf("daemonSet %s not found", objKey.String()))
			return nil
		}
		rLog.Error(err, fmt.Sprintf("getting daemonSet %s error", objKey.String()))
		return err
	}
	// 创建daemonSet
	rLog.Info(fmt.Sprintf("creating daemonSet %s error", objKey.String()))
	if err = r.Create(ctx, daemonSet); err != nil {
		rLog.Error(err, fmt.Sprintf("creating daemonSet %s", objKey.String()))
		return err
	}
	return nil
}

func createServiceIfNotExists(ctx context.Context, r *DataServiceReconciler, service *corev1.Service) error {
	if service == nil {
		return fmt.Errorf("spec Service info doesn't exist, please check crd config")
	}
	rLog := log.FromContext(ctx)
	serviceTemp := &corev1.Service{}
	objKey := types.NamespacedName{
		Name:      service.Name,
		Namespace: service.Namespace,
	}
	err := r.Get(ctx, objKey, serviceTemp)
	if err != nil {
		if errors.IsNotFound(err) {
			rLog.Info(fmt.Sprintf("service %s not found", objKey.String()))
			return nil
		}
		rLog.Error(err, fmt.Sprintf("getting service %s error", objKey.String()))
		return err
	}
	// 创建service
	rLog.Info(fmt.Sprintf("creating service %s error", objKey.String()))
	if err = r.Create(ctx, service); err != nil {
		rLog.Error(err, fmt.Sprintf("creating service %s", objKey.String()))
		return err
	}
	return nil
}

func createConfigMapIfNotExists(ctx context.Context, r *DataServiceReconciler, configMap *corev1.ConfigMap) error {
	if configMap == nil {
		return fmt.Errorf("spec ConfigMap info doesn't exist, please check crd config")
	}
	rLog := log.FromContext(ctx)
	configMapTemp := &corev1.ConfigMap{}
	objKey := types.NamespacedName{
		Name:      configMap.Name,
		Namespace: configMap.Namespace,
	}
	err := r.Get(ctx, objKey, configMapTemp)
	if err != nil {
		if errors.IsNotFound(err) {
			rLog.Info(fmt.Sprintf("configMap %s not found", objKey.String()))
			return nil
		}
		rLog.Error(err, fmt.Sprintf("getting configMap %s error", objKey.String()))
		return err
	}
	// 创建configMap
	rLog.Info(fmt.Sprintf("creating configMap %s error", objKey.String()))
	if err = r.Create(ctx, configMap); err != nil {
		rLog.Error(err, fmt.Sprintf("creating configMap %s", objKey.String()))
		return err
	}
	return nil
}
