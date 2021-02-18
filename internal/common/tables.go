/**
 * Copyright (c) 2020-present, The kubequery authors
 *
 * This source code is licensed as defined by the LICENSE file found in the
 * root directory of this source tree.
 *
 * SPDX-License-Identifier: (Apache-2.0 OR GPL-2.0-only)
 */

package common

import (
	"github.com/Uptycs/basequery-go/plugin/table"
)

// Table structure holds Osquery extension table definition.
type Table struct {
	Name    string
	Columns []table.ColumnDefinition
	GenFunc table.GenerateFunc
}
