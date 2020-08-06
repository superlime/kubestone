title: Kubestone - Ethr: Network bandwidth performance benchmark

# Ethr - Network bandwidth benchmark


    [Ethr](https://github.com/microsoft/ethr) is a cross platform network performance measurement tool written in golang. The goal of this project is to provide a native tool for comprehensive network performance measurements of bandwidth, connections/s, packets/s, latency, loss & jitter, across multiple protocols such as TCP, UDP, HTTP, HTTPS, and across multiple platforms such as Windows, Linux and other Unix systems.




## Mode of operation

As ethr requires a server and a client the controller creates the following objects during benchmark:

- Server Deployment

- Server Service

- Client Pod

  

At the first step, the Server Deployment and Service are created. Once both becomes available, the Client Pod is created to execute the benchmark. Once the benchmark is completed (regardless of it's success), the server deployment and service is deleted from Kubernetes.

In order to avoid measuring loopback performance, it is advised that you set the affinity and anti-affinity scheduling primitives for the benchmark. The provided sample benchmark shows how to avoid executing the client and the server on the same machine. For further documentation please refer to Kubernetes' [respective documentation](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/).

## Sample benchmark
```bash
$ kubectl create --namespace kubestone -f kubestone/master/config/samples/ethr.yaml
```


Please refer to the [quickstart guide](../quickstart.md) for details on generic principles and setup of Kubestone.




## Ethr Configuration

The complete documentation of ethr CR can be found in the [API Docs](../apidocs.md#perf.kubestone.xridge.io/v1alpha1.EthrSpec).



## Docker Image

[Docker Image for Ethr](https://hub.docker.com/repository/docker/bwatada/ethr)
