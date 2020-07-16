/*
Copyright 2019 The xridge kubestone contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NtttcpConfigurationSpec contains configuration parameters
// with scheduling options for the both the ntttcp client
// and server instances.
type NtttcpConfigurationSpec struct {
	PodConfigurationSpec `json:",inline"`

	// CmdLineArgs are appended to the predefined ntttcp parameters
	// +optional
	CmdLineArgs string `json:"cmdLineArgs,omitempty"`

	// HostNetwork requested for the ntttcp pod, if enabled the
	// hosts network namespace is used. Default to false.
	// +optional
	HostNetwork bool `json:"hostNetwork,omitempty"`
}

// NtttcpSpec defines the Ntttcp Benchmark Stone which
// consist of server deployment with service definition
// and client pod.
type NtttcpSpec struct {
	// Image defines the ntttcp docker image used for the benchmark
	Image ImageSpec `json:"image"`

	// ServerConfiguration contains the configuration of the ntttcp server
	// +optional
	ServerConfiguration NtttcpConfigurationSpec `json:"serverConfiguration,omitempty"`

	// ClientConfiguration contains the configuration of the ntttcp client
	// +optional
	ClientConfiguration NtttcpConfigurationSpec `json:"clientConfiguration,omitempty"`

	// If enabled the controller will create a volume and send the log file to the host node.
	// +optional
	Log LogSpec `json:"log,omitempty"`

	// The port used for both the server and client
	Port int32 `json:"port"`

	// The command used to check pod readiness
	ReadinessCmd []string `json:"readinesscmd"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Running",type="boolean",JSONPath=".status.running"
// +kubebuilder:printcolumn:name="Completed",type="boolean",JSONPath=".status.completed"

// Ntttcp is the Schema for the ethrs API
type Ntttcp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NtttcpSpec      `json:"spec,omitempty"`
	Status BenchmarkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NtttcpList contains a list of Ntttcp
type NtttcpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ntttcp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Ntttcp{}, &NtttcpList{})
}
