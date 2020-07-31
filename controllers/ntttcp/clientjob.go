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

package ntttcp

import (
	"time"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
	
	"github.com/firepear/qsplit"
	perfv1alpha1 "github.com/xridge/kubestone/api/v1alpha1"
	"github.com/xridge/kubestone/pkg/k8s"
)

// +kubebuilder:rbac:groups="",resources=pods,verbs=get;list;create;delete

func clientJobName(cr *perfv1alpha1.Ntttcp) string {
	// Should not match with service name as the pod's
	// hostname is set to it's name. If the two matches
	// the destination ip will resolve to 127.0.0.1 and
	// the server will be unreachable.
	return serverServiceName(cr) + "-client"
}

// NewClientJob creates an Ntttcp Client Job (targeting the
// Server Deployment via the Server Service) from the provided
// Ntttcp Benchmark Definition.
// ntttcp -r -m 1,*,127.0.0.1 -u -ns -v -sp -p 50001 -wu 2 -cd 2 -t 10 -l 32000 -sb -1
func NewClientJob(cr *perfv1alpha1.Ntttcp, serverAddress string) *batchv1.Job {
	objectMeta := metav1.ObjectMeta{
		Name:      clientJobName(cr),
		Namespace: cr.Namespace,
	}

	ntttcpCmdLineArgs := []string{
		"-s", "-m",
		"1,*,"+serverAddress,
	}
	//port, processor, address ()
	ntttcpCmdLineArgs = append(ntttcpCmdLineArgs,
		qsplit.ToStrings([]byte(cr.Spec.ClientConfiguration.CmdLineArgs))...)

	job := k8s.NewPerfJob(objectMeta, "ntttcp-client", cr.Spec.Image,
		cr.Spec.ClientConfiguration.PodConfigurationSpec)

	if cr.Spec.Log.Enabled {

		ntttcpCmdLineArgs = append(ntttcpCmdLineArgs, "-xml")
		
		ntttcpCmdLineArgs = append(ntttcpCmdLineArgs, cr.Spec.Log.VolumeMount.Path + cr.Spec.Log.FileName + time.Unix(1573142098, 0).Format(time.UnixDate) + cr.Spec.Log.Extension)

		volumes := []corev1.Volume{
			corev1.Volume{
				Name: cr.Spec.Log.Volume.Name,
				VolumeSource: corev1.VolumeSource{
					HostPath: &corev1.HostPathVolumeSource{
						Path: cr.Spec.Log.Volume.Path,
					},
				},
			},
		}
		volumeMounts := []corev1.VolumeMount{
			corev1.VolumeMount{
				Name:      cr.Spec.Log.VolumeMount.Name,
				MountPath: cr.Spec.Log.VolumeMount.Path,
			},
		}
		completions := int32(cr.Spec.Completions)
		job.Spec.Completions = &completions
		job.Spec.Template.Spec.Volumes = volumes
		job.Spec.Template.Spec.Containers[0].VolumeMounts = volumeMounts
	}

	backoffLimit := int32(6)
	job.Spec.BackoffLimit = &backoffLimit
	job.Spec.Template.Spec.Containers[0].Args = ntttcpCmdLineArgs
	job.Spec.Template.Spec.HostNetwork = cr.Spec.ClientConfiguration.HostNetwork

	return job
}
