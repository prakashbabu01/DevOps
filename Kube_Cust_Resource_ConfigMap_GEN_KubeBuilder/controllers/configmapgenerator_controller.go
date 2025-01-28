package controllers

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	corev1 "k8s.io/api/core/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	customv1 "github.com/example/kubebuilder-configmap-generator/api/v1"
)

// ConfigMapGeneratorReconciler reconciles a ConfigMapGenerator object
type ConfigMapGeneratorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// Reconcile is part of the main Kubernetes reconciliation loop
func (r *ConfigMapGeneratorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// Fetch the ConfigMapGenerator instance
	configMapGen := &customv1.ConfigMapGenerator{}
	err := r.Get(ctx, req.NamespacedName, configMapGen)
	if err != nil {
		if errors.IsNotFound(err) {
			logger.Info("ConfigMapGenerator resource not found. Ignoring since object must be deleted.")
			return ctrl.Result{}, nil
		}
		logger.Error(err, "Failed to get ConfigMapGenerator")
		return ctrl.Result{}, err
	}

	// Define the desired ConfigMap
	configMap := &corev1.ConfigMap{
		ObjectMeta: ctrl.ObjectMeta{
			Name:      fmt.Sprintf("%s-configmap", configMapGen.Name),
			Namespace: configMapGen.Namespace,
		},
		Data: map[string]string{
			configMapGen.Spec.Key: configMapGen.Spec.Value,
		},
	}

	// Set the owner reference for garbage collection
	if err := controllerutil.SetControllerReference(configMapGen, configMap, r.Scheme); err != nil {
		logger.Error(err, "Failed to set owner reference")
		return ctrl.Result{}, err
	}

	// Check if the ConfigMap already exists
	found := &corev1.ConfigMap{}
	err = r.Get(ctx, client.ObjectKey{Name: configMap.Name, Namespace: configMap.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		logger.Info("Creating a new ConfigMap", "ConfigMap.Namespace", configMap.Namespace, "ConfigMap.Name", configMap.Name)
		err = r.Create(ctx, configMap)
		if err != nil {
			logger.Error(err, "Failed to create ConfigMap")
			return ctrl.Result{}, err
		}
		// Update the status
		configMapGen.Status.Status = "ConfigMap Created"
		err = r.Status().Update(ctx, configMapGen)
		if err != nil {
			logger.Error(err, "Failed to update ConfigMapGenerator status")
			return ctrl.Result{}, err
		}
	} else if err != nil {
		logger.Error(err, "Failed to get ConfigMap")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager
func (r *ConfigMapGeneratorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&customv1.ConfigMapGenerator{}).
		Complete(r)
}
