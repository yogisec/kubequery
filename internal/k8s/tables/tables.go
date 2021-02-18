/**
 * Copyright (c) 2020-present, The kubequery authors
 *
 * This source code is licensed as defined by the LICENSE file found in the
 * root directory of this source tree.
 *
 * SPDX-License-Identifier: (Apache-2.0 OR GPL-2.0-only)
 */

package tables

import (
	"github.com/Uptycs/kubequery/internal/common"
	"github.com/Uptycs/kubequery/internal/k8s/admissionregistration"
	"github.com/Uptycs/kubequery/internal/k8s/apps"
	"github.com/Uptycs/kubequery/internal/k8s/autoscaling"
	"github.com/Uptycs/kubequery/internal/k8s/batch"
	"github.com/Uptycs/kubequery/internal/k8s/core"
	"github.com/Uptycs/kubequery/internal/k8s/discovery"
	"github.com/Uptycs/kubequery/internal/k8s/event"
	"github.com/Uptycs/kubequery/internal/k8s/networking"
	"github.com/Uptycs/kubequery/internal/k8s/policy"
	"github.com/Uptycs/kubequery/internal/k8s/rbac"
	"github.com/Uptycs/kubequery/internal/k8s/storage"
)

// GetTables returns the definition of all core Kubernetes tables supported by this extension.
func GetTables() []common.Table {
	return []common.Table{
		// Admission Registration
		{Name: "kubernetes_mutating_webhooks", Columns: admissionregistration.MutatingWebhookColumns(), GenFunc: admissionregistration.MutatingWebhooksGenerate},
		{Name: "kubernetes_validating_webhooks", Columns: admissionregistration.ValidatingWebhookColumns(), GenFunc: admissionregistration.ValidatingWebhooksGenerate},

		// Apps
		{Name: "kubernetes_daemon_sets", Columns: apps.DaemonSetColumns(), GenFunc: apps.DaemonSetsGenerate},
		{Name: "kubernetes_daemon_set_containers", Columns: apps.DaemonSetContainerColumns(), GenFunc: apps.DaemonSetContainersGenerate},
		{Name: "kubernetes_daemon_set_volumes", Columns: apps.DaemonSetVolumeColumns(), GenFunc: apps.DaemonSetVolumesGenerate},
		{Name: "kubernetes_deployments", Columns: apps.DeploymentColumns(), GenFunc: apps.DeploymentsGenerate},
		{Name: "kubernetes_deployments_containers", Columns: apps.DeploymentContainerColumns(), GenFunc: apps.DeploymentContainersGenerate},
		{Name: "kubernetes_deployments_volumes", Columns: apps.DeploymentVolumeColumns(), GenFunc: apps.DeploymentVolumesGenerate},
		{Name: "kubernetes_replica_sets", Columns: apps.ReplicaSetColumns(), GenFunc: apps.ReplicaSetsGenerate},
		{Name: "kubernetes_replica_set_containers", Columns: apps.ReplicaSetContainerColumns(), GenFunc: apps.ReplicaSetContainersGenerate},
		{Name: "kubernetes_replica_set_volumes", Columns: apps.ReplicaSetVolumeColumns(), GenFunc: apps.ReplicaSetVolumesGenerate},
		{Name: "kubernetes_stateful_sets", Columns: apps.StatefulSetColumns(), GenFunc: apps.StatefulSetsGenerate},
		{Name: "kubernetes_stateful_set_containers", Columns: apps.StatefulSetContainerColumns(), GenFunc: apps.StatefulSetContainersGenerate},
		{Name: "kubernetes_stateful_set_volumes", Columns: apps.StatefulSetVolumeColumns(), GenFunc: apps.StatefulSetVolumesGenerate},

		// Autoscaling
		{Name: "kubernetes_horizontal_pod_autoscalers", Columns: autoscaling.HorizontalPodAutoscalersColumns(), GenFunc: autoscaling.HorizontalPodAutoscalerGenerate},

		// Batch
		{Name: "kubernetes_cron_jobs", Columns: batch.CronJobColumns(), GenFunc: batch.CronJobsGenerate},
		{Name: "kubernetes_jobs", Columns: batch.JobColumns(), GenFunc: batch.JobsGenerate},

		// Core
		{Name: "kubernetes_config_maps", Columns: core.ConfigMapColumns(), GenFunc: core.ConfigMapsGenerate},
		{Name: "kubernetes_endpoint_subsets", Columns: core.EndpointSubsetColumns(), GenFunc: core.EndpointSubsetsGenerate},
		{Name: "kubernetes_limit_ranges", Columns: core.LimitRangeColumns(), GenFunc: core.LimitRangesGenerate},
		{Name: "kubernetes_namespaces", Columns: core.NamespaceColumns(), GenFunc: core.NamespacesGenerate},
		{Name: "kubernetes_nodes", Columns: core.NodeColumns(), GenFunc: core.NodesGenerate},
		{Name: "kubernetes_persistent_volume_claims", Columns: core.PersistentVolumeClaimColumns(), GenFunc: core.PersistentVolumeClaimsGenerate},
		{Name: "kubernetes_persistent_volumes", Columns: core.PersistentVolumeColumns(), GenFunc: core.PersistentVolumesGenerate},
		{Name: "kubernetes_pod_templates", Columns: core.PodTemplateColumns(), GenFunc: core.PodTemplatesGenerate},
		{Name: "kubernetes_pod_template_containers", Columns: core.PodTemplateContainerColumns(), GenFunc: core.PodTemplateContainersGenerate},
		{Name: "kubernetes_pod_templates_volumes", Columns: core.PodTemplateVolumeColumns(), GenFunc: core.PodTemplateVolumesGenerate},
		{Name: "kubernetes_pods", Columns: core.PodColumns(), GenFunc: core.PodsGenerate},
		{Name: "kubernetes_pod_containers", Columns: core.PodContainerColumns(), GenFunc: core.PodContainersGenerate},
		{Name: "kubernetes_pod_volumes", Columns: core.PodVolumeColumns(), GenFunc: core.PodVolumesGenerate},
		{Name: "kubernetes_resource_quotas", Columns: core.ResourceQuotaColumns(), GenFunc: core.ResourceQuotasGenerate},
		{Name: "kubernetes_secrets", Columns: core.SecretColumns(), GenFunc: core.SecretsGenerate},
		{Name: "kubernetes_service_accounts", Columns: core.ServiceAccountColumns(), GenFunc: core.ServiceAccountsGenerate},
		{Name: "kubernetes_services", Columns: core.ServiceColumns(), GenFunc: core.ServicesGenerate},

		// Discovery
		{Name: "kubernetes_api_resources", Columns: discovery.APIResourceColumns(), GenFunc: discovery.APIResourcesGenerate},
		{Name: "kubernetes_info", Columns: discovery.InfoColumns(), GenFunc: discovery.InfoGenerate},

		// Event
		{Name: "kubernetes_events", Columns: event.Columns(), GenFunc: event.Generate},

		// Networking
		{Name: "kubernetes_ingress_classes", Columns: networking.IngressClassColumns(), GenFunc: networking.IngressClassesGenerate},
		{Name: "kubernetes_ingresses", Columns: networking.IngressColumns(), GenFunc: networking.IngressesGenerate},
		{Name: "kubernetes_network_policies", Columns: networking.NetworkPolicyColumns(), GenFunc: networking.NetworkPoliciesGenerate},

		// Policy
		{Name: "kubernetes_pod_disruption_budgets", Columns: policy.PodDisruptionBudgetColumns(), GenFunc: policy.PodDisruptionBudgetsGenerate},
		{Name: "kubernetes_pod_security_policies", Columns: policy.PodSecurityPolicyColumns(), GenFunc: policy.PodSecurityPoliciesGenerate},

		// RBAC
		{Name: "kubernetes_cluster_role_binding_subjects", Columns: rbac.ClusterRoleBindingSubjectColumns(), GenFunc: rbac.ClusterRoleBindingSubjectsGenerate},
		{Name: "kubernetes_cluster_role_policy_rules", Columns: rbac.ClusterRolePolicyRuleColumns(), GenFunc: rbac.ClusterRolePolicyRulesGenerate},
		{Name: "kubernetes_role_binding_subjects", Columns: rbac.RoleBindingSubjectColumns(), GenFunc: rbac.RoleBindingSubjectsGenerate},
		{Name: "kubernetes_role_policy_rules", Columns: rbac.RolePolicyRuleColumns(), GenFunc: rbac.RolePolicyRulesGenerate},

		// Storage
		{Name: "kubernetes_csi_drivers", Columns: storage.CSIDriverColumns(), GenFunc: storage.CSIDriversGenerate},
		{Name: "kubernetes_csi_node_drivers", Columns: storage.CSINodeDriverColumns(), GenFunc: storage.CSINodeDriversGenerate},
		{Name: "kubernetes_storage_capacities", Columns: storage.CSIStorageCapacityColumns(), GenFunc: storage.CSIStorageCapacitiesGenerate},
		{Name: "kubernetes_storage_classes", Columns: storage.SGClassColumns(), GenFunc: storage.SGClassesGenerate},
		{Name: "kubernetes_volume_attachments", Columns: storage.VolumeAttachmentColumns(), GenFunc: storage.VolumeAttachmentsGenerate},
	}
}
