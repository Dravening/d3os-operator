package controllers

import (
	"context"
	"fmt"

	d3osoperatorv1 "d3os-operator/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

// CheckExistsOrCreateMidBackend 验证或新建MiddlewareBackend
func CheckExistsOrCreateMidBackend(ctx context.Context, r *DataServiceReconciler, req ctrl.Request, middlewareBackend *d3osoperatorv1.MiddlewareBackend) error {
	var err error
	switch middlewareBackend.Kind {
	case d3osoperatorv1.Deployment:
		err = createDeploymentIfNotExists(ctx, r, req, middlewareBackend.Deployment)
	case d3osoperatorv1.StatefulSet:
		err = createStatefulSetIfNotExists(ctx, r, req, middlewareBackend.StatefulSet)
	case d3osoperatorv1.DaemonSet:
		err = createDaemonSetIfNotExists(ctx, r, req, middlewareBackend.DaemonSet)
	case d3osoperatorv1.ThirdParty:
		// 要查询checkCMD
		if cmdCall(middlewareBackend.ThirdParty.CheckCMD) != nil {
			err = fmt.Errorf("unable to connect thirdparty middleware %s, will waiting", middlewareBackend.ThirdParty.Url)
		}
	}
	if err != nil {
		return err
	}
	if middlewareBackend.Service != nil {
		err = createServiceIfNotExists(ctx, r, req, middlewareBackend.Service)
		if err != nil {
			return err
		}
	}
	if middlewareBackend.ConfigMap != nil {
		err = createConfigMapIfNotExists(ctx, r, req, middlewareBackend.ConfigMap)
		if err != nil {
			return err
		}
	}
	return nil
}

// CheckExistsOrCreateSvcBackend 验证或新建DeployWithSvcCM, deployment和service都是必填项, 如果nil直接报错
func CheckExistsOrCreateSvcBackend(ctx context.Context, r *DataServiceReconciler, req ctrl.Request, svcBackend *d3osoperatorv1.ServiceBackend) error {
	var err error
	err = createDeploymentIfNotExists(ctx, r, req, svcBackend.Deployment)
	if err != nil {
		return err
	}
	err = createServiceIfNotExists(ctx, r, req, svcBackend.Service)
	if err != nil {
		return err
	}
	if svcBackend.ConfigMap != nil {
		for _, cm := range svcBackend.ConfigMap {
			err = createConfigMapIfNotExists(ctx, r, req, cm)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
