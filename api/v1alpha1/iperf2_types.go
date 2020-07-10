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

// Iperf2ConfigurationSpec contains configuration parameters
// with scheduling options for the both the iperf2 client
// and server instances.
type Iperf2ConfigurationSpec struct {
	PodConfigurationSpec `json:",inline"`

	// CmdLineArgs are appended to the predefined iperf2 parameters
	// +optional
	CmdLineArgs string `json:"cmdLineArgs,omitempty"`

	// HostNetwork requested for the iperf2 pod, if enabled the
	// hosts network namespace is used. Default to false.
	// +optional
	HostNetwork bool `json:"hostNetwork,omitempty"`
}

// Iperf2Spec defines the Iperf2 Benchmark Stone which
// consist of server deployment with service definition
// and client pod.
type Iperf2Spec struct {
	// Image defines the iperf2 docker image used for the benchmark
	Image ImageSpec `json:"image"`

	// ServerConfiguration contains the configuration of the iperf2 server
	// +optional
	ServerConfiguration Iperf2ConfigurationSpec `json:"serverConfiguration,omitempty"`

	// ClientConfiguration contains the configuration of the iperf2 client
	// +optional
	ClientConfiguration Iperf2ConfigurationSpec `json:"clientConfiguration,omitempty"`

	// UDP to use rather than TCP.
	// If enabled the '--udp' parameter is added to iperf command line args
	// +optional
	UDP bool `json:"udp,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Running",type="boolean",JSONPath=".status.running"
// +kubebuilder:printcolumn:name="Completed",type="boolean",JSONPath=".status.completed"

// Iperf2 is the Schema for the iperf2s API
type Iperf2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Iperf2Spec      `json:"spec,omitempty"`
	Status BenchmarkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Iperf2List contains a list of Iperf2
type Iperf2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Iperf2 `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Iperf2{}, &Iperf2List{})
}
