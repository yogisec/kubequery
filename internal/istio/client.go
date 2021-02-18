/**
 * Copyright (c) 2020-present, The kubequery authors
 *
 * This source code is licensed as defined by the LICENSE file found in the
 * root directory of this source tree.
 *
 * SPDX-License-Identifier: (Apache-2.0 OR GPL-2.0-only)
 */

package istio

import (
	"sync"

	"istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/rest"
)

var (
	lock      sync.Mutex
	clientset versioned.Interface
)

func initClientset(config *rest.Config) error {
	if config == nil {
		// Get in-cluster configuration if one is not provided
		conf, err := rest.InClusterConfig()
		if err != nil {
			return err
		}
		config = conf
	}

	var err error
	clientset, err = versioned.NewForConfig(config)
	if err != nil {
		return err
	}
	return nil
}

// Init creates in-cluster kubernetes configuration and a client set using the configuration.
// This returns error if KUBERNETES_SERVICE_HOST or KUBERNETES_SERVICE_PORT environment variables are not set.
func Init() error {
	lock.Lock()
	defer lock.Unlock()

	err := initClientset(nil)
	if err != nil {
		return err
	}

	return nil
}

// GetClient returns kubernetes interface that can be used to communicate with API server.
func GetClient() versioned.Interface {
	return clientset
}

// SetClient is helper function to override the kubernetes interface with fake one for testing.
func SetClient(client versioned.Interface) {
	lock.Lock()
	defer lock.Unlock()
	clientset = client
}
