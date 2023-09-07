package controllers

import (
	"context"
	"fmt"
	"time"

	d3osoperatorv1 "d3os-operator/api/v1"
)

// CheckExistsOrCreateMidBackend 验证或新建MiddlewareBackend
func CheckExistsOrCreateMidBackend(ctx context.Context, r *DataServiceReconciler, oldMiddlewareBackend, newMiddlewareBackend *d3osoperatorv1.MiddlewareBackend, dsInstance *d3osoperatorv1.DataService) error {
	var err error
	switch newMiddlewareBackend.Kind {
	case d3osoperatorv1.Deployment:
		err = createDeploymentIfNotExists(ctx, r, oldMiddlewareBackend.Deployment, newMiddlewareBackend.Deployment, dsInstance)
	case d3osoperatorv1.StatefulSet:
		err = createStatefulSetIfNotExists(ctx, r, oldMiddlewareBackend.StatefulSet, newMiddlewareBackend.StatefulSet, dsInstance)
	case d3osoperatorv1.DaemonSet:
		err = createDaemonSetIfNotExists(ctx, r, oldMiddlewareBackend.DaemonSet, newMiddlewareBackend.DaemonSet, dsInstance)
	case d3osoperatorv1.ThirdParty:
		// 要查询checkCMD
		if err = cmdCall(newMiddlewareBackend.ThirdParty.CheckCMD); err != nil {
			return fmt.Errorf("unable to connect thirdparty middleware %s, err: %s", newMiddlewareBackend.ThirdParty.Url, err)
		}
	}
	if err != nil {
		return err
	}
	if newMiddlewareBackend.Service != nil {
		err = createServiceIfNotExists(ctx, r, oldMiddlewareBackend.Service, newMiddlewareBackend.Service, dsInstance)
		if err != nil {
			return err
		}
	}
	if newMiddlewareBackend.ConfigMap != nil {
		for cmName, cm := range newMiddlewareBackend.ConfigMap {
			oldCm := oldMiddlewareBackend.ConfigMap[cmName]
			err = createConfigMapIfNotExists(ctx, r, oldCm, cm, dsInstance)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// CheckExistsOrCreateAppBackend 验证或新建DeployWithSvcCM, deployment和service都是必填项, 如果nil直接报错
func CheckExistsOrCreateAppBackend(ctx context.Context, r *DataServiceReconciler, oldSvcBackend, newSvcBackend *d3osoperatorv1.ServiceBackend, dsInstance *d3osoperatorv1.DataService) error {
	var err error
	err = createDeploymentIfNotExists(ctx, r, oldSvcBackend.Deployment, newSvcBackend.Deployment, dsInstance)
	if err != nil {
		return err
	}
	err = createServiceIfNotExists(ctx, r, oldSvcBackend.Service, newSvcBackend.Service, dsInstance)
	if err != nil {
		return err
	}
	if newSvcBackend.ConfigMap != nil {
		for cmName, cm := range newSvcBackend.ConfigMap {
			oldCm := newSvcBackend.ConfigMap[cmName]
			err = createConfigMapIfNotExists(ctx, r, oldCm, cm, dsInstance)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func CheckMidBackendStatus(ctx context.Context, r *DataServiceReconciler, middlewareBackend *d3osoperatorv1.MiddlewareBackend, waitSec time.Duration) error {
	var err error
	time.Sleep(waitSec)

	switch middlewareBackend.Kind {
	case d3osoperatorv1.Deployment:
		err = checkDeploymentStatus(ctx, r, middlewareBackend.Deployment)
	case d3osoperatorv1.StatefulSet:
		err = checkStatefulSetStatus(ctx, r, middlewareBackend.StatefulSet)
	case d3osoperatorv1.DaemonSet:
		err = checkDaemonSetStatus(ctx, r, middlewareBackend.DaemonSet)
	case d3osoperatorv1.ThirdParty:
		// 要查询checkCMD
		if err = cmdCall(middlewareBackend.ThirdParty.CheckCMD); err != nil {
			return fmt.Errorf("unable to connect thirdparty middleware %s, err: %s", middlewareBackend.ThirdParty.Url, err)
		}
	}
	return nil
}

func CheckAppBackendStatus(ctx context.Context, r *DataServiceReconciler, svcBackend *d3osoperatorv1.ServiceBackend, waitSec time.Duration) error {
	time.Sleep(waitSec)

	if err := checkDeploymentStatus(ctx, r, svcBackend.Deployment); err != nil {
		return err
	}
	return nil
}
