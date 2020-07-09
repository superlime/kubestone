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

const TCPBandwidthPort = 9999
const TCPConnectionsPort = 9998
const TCPLatencyPort = 9996
const HTTPBandwidthPort = 9899
const HTTPSBandwidthPort = 9799
const HTTPLatencyPort = 9896
// EthrConfigurationSpec contains configuration parameters
// with scheduling options for the both the ethr client
// and server instances.
type EthrConfigurationSpec struct {
	PodConfigurationSpec `json:",inline"`

	// CmdLineArgs are appended to the predefined ethr parameters
	// +optional
	CmdLineArgs string `json:"cmdLineArgs,omitempty"`

	// HostNetwork requested for the ethr pod, if enabled the
	// hosts network namespace is used. Default to false.
	// +optional
	HostNetwork bool `json:"hostNetwork,omitempty"`
}

// EthrSpec defines the Ethr Benchmark Stone which
// consist of server deployment with service definition
// and client pod.
type EthrSpec struct {
	// Image defines the ethr docker image used for the benchmark
	Image ImageSpec `json:"image"`

	// ServerConfiguration contains the configuration of the ethr server
	// +optional
	ServerConfiguration EthrConfigurationSpec `json:"serverConfiguration,omitempty"`

	// ClientConfiguration contains the configuration of the ethr client
	// +optional
	ClientConfiguration EthrConfigurationSpec `json:"clientConfiguration,omitempty"`

	// UDP to use rather than TCP.
	// If enabled the '--udp' parameter is added to iperf command line args
	// +optional
	UDP bool `json:"udp,omitempty"`

	// Volume contains the configuration for the volume that the ethr job should
	// run on.
	Volume VolumeSpec `json:"volume"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Running",type="boolean",JSONPath=".status.running"
// +kubebuilder:printcolumn:name="Completed",type="boolean",JSONPath=".status.completed"

// Ethr is the Schema for the ethrs API
type Ethr struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EthrSpec      `json:"spec,omitempty"`
	Status BenchmarkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EthrList contains a list of Ethr
type EthrList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ethr `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Ethr{}, &EthrList{})
}
