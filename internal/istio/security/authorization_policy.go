/**
 * Copyright (c) 2020-present, The kubequery authors
 *
 * This source code is licensed as defined by the LICENSE file found in the
 * root directory of this source tree.
 *
 * SPDX-License-Identifier: (Apache-2.0 OR GPL-2.0-only)
 */

package security

import (
	"context"

	"github.com/Uptycs/basequery-go/plugin/table"
	istio "github.com/Uptycs/kubequery/internal/istio"
	k8s "github.com/Uptycs/kubequery/internal/k8s"
	v1alpha1 "istio.io/api/analysis/v1alpha1"
	metav1alpha1 "istio.io/api/meta/v1alpha1"
	secv1beta1 "istio.io/api/security/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type authorizationPolicy struct {
	k8s.CommonNamespacedFields
	MatchLabels        *map[string]string
	Rules              []*secv1beta1.Rule
	Action             secv1beta1.AuthorizationPolicy_Action
	Conditions         []*metav1alpha1.IstioCondition
	ValidationMessages []*v1alpha1.AnalysisMessageBase
	ObservedGeneration int64
}

// AuthorizationPolicyColumns TODO.
func AuthorizationPolicyColumns() []table.ColumnDefinition {
	return k8s.GetSchema(&authorizationPolicy{})
}

// AuthorizationPoliciesGenerate TODO.
func AuthorizationPoliciesGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	options := metav1.ListOptions{}
	results := make([]map[string]string, 0)

	for {
		aps, err := istio.GetClient().SecurityV1beta1().AuthorizationPolicies(metav1.NamespaceAll).List(ctx, options)
		if err != nil {
			return nil, err
		}

		for _, ap := range aps.Items {
			item := &authorizationPolicy{
				CommonNamespacedFields: k8s.GetCommonNamespacedFields(ap.ObjectMeta),
				Rules:                  ap.Spec.Rules,
				Action:                 ap.Spec.Action,
				Conditions:             ap.Status.Conditions,
				ValidationMessages:     ap.Status.ValidationMessages,
				ObservedGeneration:     ap.Status.ObservedGeneration,
			}
			if ap.Spec.Selector != nil {
				item.MatchLabels = &ap.Spec.Selector.MatchLabels
			}
			results = append(results, k8s.ToMap(item))
		}

		if aps.Continue == "" {
			break
		}
		options.Continue = aps.Continue
	}

	return results, nil
}
