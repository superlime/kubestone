title: Kubestone - Ntttcp: Network bandwidth performance benchmark

# Ntttcp - Network bandwidth benchmark


    [Ntttcp](https://gallery.technet.microsoft.com/NTttcp-Version-528-Now-f8b12769) Used to profile and measure Windows networking performance, NTttcp is one of the primary tools Microsoft engineering teams leverage to validate network function and utility.




## Mode of operation

As ntttcp requires a server and a client the controller creates the following objects during benchmark:

- Server Deployment

- Server Service

- Client Pod

  

At the first step, the Server Deployment and Service are created. Once both becomes available, the Client Pod is created to execute the benchmark. Once the benchmark is completed (regardless of it's success), the server deployment and service is deleted from Kubernetes.

In order to avoid measuring loopback performance, it is advised that you set the affinity and anti-affinity scheduling primitives for the benchmark. The provided sample benchmark shows how to avoid executing the client and the server on the same machine. For further documentation please refer to Kubernetes' [respective documentation](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/).

## Sample benchmark
```bash
$ kubectl create --namespace kubestone -f kubestone/master/config/samples/ntttcp.yaml
```


Please refer to the [quickstart guide](../quickstart.md) for details on generic principles and setup of Kubestone.




## Ntttcp Configuration

The complete documentation of ntttcp CR can be found in the [API Docs](../apidocs.md#perf.kubestone.xridge.io/v1alpha1.Ntttcp.spec).



## Docker Image

[Docker Image for Ntttcp](https://hub.docker.com/repository/docker/bwatada/ntttcp)
