#!/bin/bash
kubectl run fortio --rm -it --image=fortio/fortio --restart=Never -- load -qps 800 -t 220s -c 70 "http://goserver-service:8000/healthz";