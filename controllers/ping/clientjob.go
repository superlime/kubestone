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

package ping

import (

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/firepear/qsplit"
	perfv1alpha1 "github.com/xridge/kubestone/api/v1alpha1"
	"github.com/xridge/kubestone/pkg/k8s"
)

// +kubebuilder:rbac:groups="",resources=pods,verbs=get;list;create;delete

func clientJobName(cr *perfv1alpha1.Ping) string {
	// Should not match with service name as the pod's
	// hostname is set to it's name. If the two matches
	// the destination ip will resolve to 127.0.0.1 and
	// the server will be unreachable.
	return serverServiceName(cr) + "-client"
}

// NewClientJob creates an Ping Client Job (targeting the
// Server Deployment via the Server Service) from the provided
// Ping Benchmark Definition.
func NewClientJob(cr *perfv1alpha1.Ping, serviceIp string) *batchv1.Job {
	objectMeta := metav1.ObjectMeta{
		Name:      clientJobName(cr),
		Namespace: cr.Namespace,
	}

	pingCmdLineArgs := []string{
		serviceIp,
	}
	pingCmdLineArgs = append(pingCmdLineArgs, qsplit.ToStrings([]byte(cr.Spec.Options))...)
	pingCmdLineArgs = append(pingCmdLineArgs)

	backoffLimit := int32(6)

	job := k8s.NewPerfJob(objectMeta, "ping-client", cr.Spec.Image, cr.Spec.ClientConfiguration.PodConfigurationSpec)
	job.Spec.BackoffLimit = &backoffLimit
	job.Spec.Template.Spec.Containers[0].Args = pingCmdLineArgs
	job.Spec.Template.Spec.HostNetwork = cr.Spec.ClientConfiguration.HostNetwork

	return job
}
