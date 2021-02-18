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
	"github.com/Uptycs/kubequery/internal/istio/security"
)

// GetTables returns the definition of all Istio tables supported by this extension.
func GetTables() []common.Table {
	return []common.Table{
		// Security
		{Name: "istio_authorization_policies", Columns: security.AuthorizationPolicyColumns(), GenFunc: security.AuthorizationPoliciesGenerate},
		{Name: "istio_peer_authentications", Columns: security.PeerAuthenticationColumns(), GenFunc: security.PeerAuthenticationsGenerate},
		{Name: "istio_request_authentications", Columns: security.RequestAuthenticationColumns(), GenFunc: security.RequestAuthenticationsGenerate},
	}
}
