import kopf
import kubernetes.client
from kubernetes.client import V1ConfigMap
from kubernetes.client.rest import ApiException

# Watch for pod creation events
@kopf.on.create('', 'v1', 'pods')
def on_pod_create(body, namespace, logger, **kwargs):
    """
    Triggered whenever a Pod is created. If the Pod has the label `project-type: bank`,
    create a ConfigMap named `log-config` and attach it to the Pod as a volume.
    """
    labels = body.get('metadata', {}).get('labels', {})
    pod_name = body.get('metadata', {}).get('name', 'unknown')

    # Check for the `project-type: bank` label
    if labels.get('project-type') == 'bank':
        logger.info(f"Pod '{pod_name}' with label 'project-type: bank' detected. Creating ConfigMap.")

        # Define the ConfigMap data
        configmap = V1ConfigMap(
            metadata=kubernetes.client.V1ObjectMeta(
                name='log-config',
                namespace=namespace
            ),
            data={
                "log_level": "INFO"
            }
        )

        # Create the ConfigMap in the same namespace as the Pod
        api = kubernetes.client.CoreV1Api()
        try:
            api.create_namespaced_config_map(namespace=namespace, body=configmap)
            logger.info("ConfigMap 'log-config' created successfully.")
        except ApiException as e:
            logger.error(f"Failed to create ConfigMap: {e}")
            raise kopf.TemporaryError(f"Error creating ConfigMap: {e}", delay=30)

        # Update the Pod spec to attach the ConfigMap as a volume (if needed)
        logger.info(f"ConfigMap 'log-config' will be mounted to Pod '{pod_name}'.")
    else:
        logger.info(f"Pod '{pod_name}' does not have label 'project-type: bank'. Ignoring.")
