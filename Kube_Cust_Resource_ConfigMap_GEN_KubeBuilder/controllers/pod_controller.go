package controllers

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// PodReconciler reconciles a Pod object
type PodReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// Reconcile watches for Pod creation events
func (r *PodReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// Fetch the Pod instance
	pod := &corev1.Pod{}
	err := r.Get(ctx, req.NamespacedName, pod)
	if err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Check if the Pod has the label project-type=bank
	if pod.Labels["project-type"] == "bank" {
		configMapName := fmt.Sprintf("%s-config", pod.Name)
		configMap := &corev1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      configMapName,
				Namespace: pod.Namespace,
			},
			Data: map[string]string{
				"config.yaml": "key: value\nanother_key: another_value",
			},
		}

		// Set Pod as the owner of the ConfigMap
		if err := controllerutil.SetControllerReference(pod, configMap, r.Scheme); err != nil {
			logger.Error(err, "Failed to set owner reference on ConfigMap")
			return ctrl.Result{}, err
		}

		// Check if the ConfigMap already exists
		found := &corev1.ConfigMap{}
		err = r.Get(ctx, client.ObjectKey{Name: configMap.Name, Namespace: configMap.Namespace}, found)
		if err != nil && client.IgnoreNotFound(err) != nil {
			logger.Error(err, "Failed to get ConfigMap")
			return ctrl.Result{}, err
		}

		if err != nil && client.IgnoreNotFound(err) == nil {
			logger.Info("Creating ConfigMap", "ConfigMap.Namespace", configMap.Namespace, "ConfigMap.Name", configMap.Name)
			err = r.Create(ctx, configMap)
			if err != nil {
				logger.Error(err, "Failed to create ConfigMap")
				return ctrl.Result{}, err
			}
		} else {
			logger.Info("ConfigMap already exists", "ConfigMap.Namespace", found.Namespace, "ConfigMap.Name", found.Name)
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager
func (r *PodReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Pod{}).
		Complete(r)
}
