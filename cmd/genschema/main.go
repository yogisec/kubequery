/**
 * Copyright (c) 2020-present, The kubequery authors
 *
 * This source code is licensed as defined by the LICENSE file found in the
 * root directory of this source tree.
 *
 * SPDX-License-Identifier: (Apache-2.0 OR GPL-2.0-only)
 */

package main

import (
	"fmt"

	"github.com/Uptycs/kubequery/internal/common"
	istiotables "github.com/Uptycs/kubequery/internal/istio/tables"
	k8stables "github.com/Uptycs/kubequery/internal/k8s/tables"
)

func main() {
	for _, tables := range [...][]common.Table{k8stables.GetTables(), istiotables.GetTables()} {
		for _, t := range tables {
			fmt.Printf("CREATE TABLE %s (\n", t.Name)
			for i, c := range t.Columns {
				fmt.Printf("    `%s` %s", c.Name, c.Type)
				if i < len(t.Columns)-1 {
					fmt.Printf(",")
				}
				fmt.Println()
			}
			fmt.Print(");\n\n")
		}
	}
}
