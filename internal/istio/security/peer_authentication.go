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

type peerAuthentication struct {
	k8s.CommonNamespacedFields
	MatchLabels        *map[string]string
	Mtls               *v1beta1.PeerAuthentication_MutualTLS
	PortLevelMtls      map[uint32]*v1beta1.PeerAuthentication_MutualTLS
	Conditions         []*metav1alpha1.IstioCondition
	ValidationMessages []*v1alpha1.AnalysisMessageBase
	ObservedGeneration int64
}

// PeerAuthenticationColumns TODO.
func PeerAuthenticationColumns() []table.ColumnDefinition {
	return k8s.GetSchema(&peerAuthentication{})
}

// PeerAuthenticationsGenerate TODO.
func PeerAuthenticationsGenerate(ctx context.Context, queryContext table.QueryContext) ([]map[string]string, error) {
	options := metav1.ListOptions{}
	results := make([]map[string]string, 0)

	for {
		pas, err := istio.GetClient().SecurityV1beta1().PeerAuthentications(metav1.NamespaceAll).List(ctx, options)
		if err != nil {
			return nil, err
		}

		for _, pa := range pas.Items {
			item := &peerAuthentication{
				CommonNamespacedFields: k8s.GetCommonNamespacedFields(pa.ObjectMeta),
				Mtls:                   pa.Spec.Mtls,
				PortLevelMtls:          pa.Spec.PortLevelMtls,
				Conditions:             pa.Status.Conditions,
				ValidationMessages:     pa.Status.ValidationMessages,
				ObservedGeneration:     pa.Status.ObservedGeneration,
			}
			if pa.Spec.Selector != nil {
				item.MatchLabels = &pa.Spec.Selector.MatchLabels
			}
			results = append(results, k8s.ToMap(item))
		}

		if pas.Continue == "" {
			break
		}
		options.Continue = pas.Continue
	}

	return results, nil
}
