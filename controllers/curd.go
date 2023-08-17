package controllers

import (
	"context"
	"fmt"
	"os/exec"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
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

func createDeploymentIfNotExists(ctx context.Context, r *DataServiceReconciler, req ctrl.Request, deploy *appsv1.Deployment) error {
	if deploy == nil {
		return fmt.Errorf("spec Deployment info doesn't exist, please check crd config")
	}
	rLog := log.FromContext(ctx)
	deployTemp := &appsv1.Deployment{}
	err := r.Get(ctx, req.NamespacedName, deployTemp)
	if err != nil {
		if errors.IsNotFound(err) {
			rLog.Info("deployment %s not found", req.String())
			return nil
		}
		rLog.Error(err, "getting deployment %s error: ", req.String(), err)
		return err
	}
	// 创建deployment
	rLog.Info("creating deployment %s", req.String())
	if err = r.Create(ctx, deploy); err != nil {
		rLog.Error(err, "creating deployment %s error: ", req.String(), err)
		return err
	}
	return nil
}

func createStatefulSetIfNotExists(ctx context.Context, r *DataServiceReconciler, req ctrl.Request, statefulSet *appsv1.StatefulSet) error {
	if statefulSet == nil {
		return fmt.Errorf("spec StatefulSet info doesn't exist, please check crd config")
	}
	rLog := log.FromContext(ctx)
	stateTemp := &appsv1.StatefulSet{}
	err := r.Get(ctx, req.NamespacedName, stateTemp)
	if err != nil {
		if errors.IsNotFound(err) {
			rLog.Info("statefulSet %s not found", req.String())
			return nil
		}
		rLog.Error(err, "getting statefulSet %s error: ", req.String(), err)
		return err
	}
	// 创建statefulSet
	rLog.Info("creating statefulSet %s", req.String())
	if err = r.Create(ctx, statefulSet); err != nil {
		rLog.Error(err, "creating statefulSet %s error: ", req.String(), err)
		return err
	}
	return nil
}

func createDaemonSetIfNotExists(ctx context.Context, r *DataServiceReconciler, req ctrl.Request, daemonSet *appsv1.DaemonSet) error {
	if daemonSet == nil {
		return fmt.Errorf("spec DaemonSet info doesn't exist, please check crd config")
	}
	rLog := log.FromContext(ctx)
	daemonTemp := &appsv1.DaemonSet{}
	err := r.Get(ctx, req.NamespacedName, daemonTemp)
	if err != nil {
		if errors.IsNotFound(err) {
			rLog.Info("daemonSet %s not found", req.String())
			return nil
		}
		rLog.Error(err, "getting daemonSet %s error: ", req.String(), err)
		return err
	}
	// 创建daemonSet
	rLog.Info("creating daemonSet %s", req.String())
	if err = r.Create(ctx, daemonSet); err != nil {
		rLog.Error(err, "creating daemonSet %s error: ", req.String(), err)
		return err
	}
	return nil
}

func createServiceIfNotExists(ctx context.Context, r *DataServiceReconciler, req ctrl.Request, service *corev1.Service) error {
	if service == nil {
		return fmt.Errorf("spec Service info doesn't exist, please check crd config")
	}
	rLog := log.FromContext(ctx)
	serviceTemp := &corev1.Service{}
	err := r.Get(ctx, req.NamespacedName, serviceTemp)
	if err != nil {
		if errors.IsNotFound(err) {
			rLog.Info("service %s not found", req.String())
			return nil
		}
		rLog.Error(err, "getting service %s error: ", req.String(), err)
		return err
	}
	// 创建service
	rLog.Info("creating service %s", req.String())
	if err = r.Create(ctx, service); err != nil {
		rLog.Error(err, "creating service %s error: ", req.String(), err)
		return err
	}
	return nil
}

func createConfigMapIfNotExists(ctx context.Context, r *DataServiceReconciler, req ctrl.Request, configMap *corev1.ConfigMap) error {
	if configMap == nil {
		return fmt.Errorf("spec ConfigMap info doesn't exist, please check crd config")
	}
	rLog := log.FromContext(ctx)
	configMapTemp := &corev1.ConfigMap{}
	err := r.Get(ctx, req.NamespacedName, configMapTemp)
	if err != nil {
		if errors.IsNotFound(err) {
			rLog.Info("configMap %s not found", req.String())
			return nil
		}
		rLog.Error(err, "getting configMap %s error: %s", req.String(), err)
		return err
	}
	// 创建configMap
	rLog.Info("creating configMap %s", req.String())
	if err = r.Create(ctx, configMap); err != nil {
		rLog.Error(err, "creating configMap %s error: ", req.String(), err)
		return err
	}
	return nil
}
