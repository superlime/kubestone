#!/bin/sh

d=$(date +%Y-%m-%d-%H-%M-%S)

e=$(iperf $@)

echo $e >> /logs/iperf_$d.log

echo logs written to /logs/iperf_$d.log