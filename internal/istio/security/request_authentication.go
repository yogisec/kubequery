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
	"istio.io/api/security/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type requestAuthentication struct {
	k8s.CommonNamespacedFields
	MatchLabels        *map[string]string
	JwtRules           []*v1beta1.JWTRule
	Conditions         []*metav1alpha1.IstioCondition
	ValidationMessages []*v1alpha1.AnalysisMessageBase
	ObservedGeneration int64
}

// RequestAuthenticationColumns TODO.
func RequestAuthenticationColumns() []table.ColumnDefinition {
	return k8s.GetSchema(&requestAuthentication{})
}

// RequestAuthenticationsGenerate TODO.
func RequestAuthenticationsGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	options := metav1.ListOptions{}
	results := make([]map[string]string, 0)

	for {
		ras, err := istio.GetClient().SecurityV1beta1().RequestAuthentications(metav1.NamespaceAll).List(ctx, options)
		if err != nil {
			return nil, err
		}

		for _, ra := range ras.Items {
			item := &requestAuthentication{
				CommonNamespacedFields: k8s.GetCommonNamespacedFields(ra.ObjectMeta),
				JwtRules:               ra.Spec.JwtRules,
				Conditions:             ra.Status.Conditions,
				ValidationMessages:     ra.Status.ValidationMessages,
				ObservedGeneration:     ra.Status.ObservedGeneration,
			}
			if ra.Spec.Selector != nil {
				item.MatchLabels = &ra.Spec.Selector.MatchLabels
			}
			results = append(results, k8s.ToMap(item))
		}

		if ras.Continue == "" {
			break
		}
		options.Continue = ras.Continue
	}

	return results, nil
}
