title: Kubestone - Iperf2: Network bandwidth performance benchmark

# Iperf2 - Network bandwidth benchmark


    [Iperf2](https://sourceforge.net/projects/iperf2/) is a network traffic tool for measuring TCP and UDP performance. The goals include maintaining an active iperf 2 code base (code originated from iperf 2.0.5), preserving interoperability with iperf 2.0.5 clients and servers, preserving the output for scripts (new enhanced output requires -e), adopt known 2.0.x bug fixes, maintain broad platform support, as well as add some essential feature enhancements mostly driven by WiFi testing needs. Also added python code to centralize test control.




## Mode of operation

As iperf2 requires a server and a client the controller creates the following objects during benchmark:

- Server Deployment

- Server Service

- Client Pod

  

At the first step, the Server Deployment and Service are created. Once both becomes available, the Client Pod is created to execute the benchmark. Once the benchmark is completed (regardless of it's success), the server deployment and service is deleted from Kubernetes.

In order to avoid measuring loopback performance, it is advised that you set the affinity and anti-affinity scheduling primitives for the benchmark. The provided sample benchmark shows how to avoid executing the client and the server on the same machine. For further documentation please refer to Kubernetes' [respective documentation](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/).

## Sample benchmark
```bash
$ kubectl create --namespace kubestone -f kubestone/master/config/samples/iperf2.yaml
```


Please refer to the [quickstart guide](../quickstart.md) for details on generic principles and setup of Kubestone.




## Iperf2 Configuration

The complete documentation of iperf2 CR can be found in the [API Docs](../apidocs.md#perf.kubestone.xridge.io/v1alpha1.Iperf2Spec).



## Docker Image

[Docker Image for Iperf2](https://hub.docker.com/repository/docker/bwatada/iperf2)
