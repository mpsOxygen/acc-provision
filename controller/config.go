// Copyright 2017 Cisco Systems, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
)

type OpflexGroup struct {
	PolicySpace string `json:"policy-space,omitempty"`
	Name        string `json:"name,omitempty"`
}

type IPRange struct {
	start string `json:"start,omitempty"`
	end   string `json:"end,omitempty"`
}

// Configuration for the controller
type ControllerConfig struct {
	// Log level
	LogLevel string `json:"log-level,omitempty"`

	// Absolute path to a kubeconfig file
	KubeConfig string `json:"kubeconfig,omitempty"`

	// Default endpoint group annotation value
	DefaultEg OpflexGroup `json:"default-endpoint-group,omitempty"`

	// Default security group annotation value
	DefaultSg []OpflexGroup `json:"default-security-group,omitempty"`

	// IP addresses used for externally exposed load balanced services
	ServiceIPPool []IPRange `json:"service-ip-pool,omitempty"`

	// IP addresses that can be requested as static service IPs in
	// service spec
	StaticServiceIPPool []IPRange `json:"static-service-ip-pool,omitempty"`

	// IP addresses used for pod network
	PodIPPool []IPRange `json:"pod-ip-pool,omitempty"`
}

func initFlags() {
	flag.StringVar(&config.LogLevel, "log-level", "info", "Log level")

	flag.StringVar(&config.KubeConfig, "kubeconfig", "", "Absolute path to a kubeconfig file")
}
