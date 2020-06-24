# iptables logger

[<img src="ll-logo.png">](https://lablabs.io/)

We help companies build, run, deploy and scale software and infrastructure by embracing the right technologies and principles. Check out our website at https://lablabs.io/

---

## Description

This tool monitors how many TCP connections are being reused when using iptables.
It was originally developed to monitor how many connections are being reused in our Kubernetes cluster due to a bug described [here](https://github.com/kubernetes/kubernetes/issues/81775).

# Build using docker

```
docker build -t iptables-logger .
```

## Usage

```
docker run --network=host -it iptables-logger
```

### Optional args
```
  -interval int
        evaluation interval in seconds (default 5)
  -path string
        Path to ip_vs_conn file (default "/proc/net/nf_conntrack")
  -service-ip string
        IP address of the target service
```

### Sample output

```
2020-06-17 10:52:36.139931577 +0000 UTC m=+16.000442290 - from: 100.119.69.151:58204 - to 100.70.132.168:53, dest: 100.100.233.255:53 28-114
2020-06-17 10:52:36.14002535 +0000 UTC m=+16.000536052 - from: 100.119.69.139:48692 - to 100.70.132.168:53, dest: 100.108.1.232:53 83-116
2020-06-17 10:52:36.140042704 +0000 UTC m=+16.000553405 - from: 100.119.69.137:36264 - to 100.70.132.168:53, dest: 100.123.202.248:53 21-108
2020-06-17 10:52:36.140415101 +0000 UTC m=+16.000925802 - from: 100.119.69.137:37588 - to 100.70.132.168:53, dest: 100.115.9.245:53 24-112
2020-06-17 10:52:36.140952998 +0000 UTC m=+16.001463699 - from: 100.119.69.150:45608 - to 100.70.132.168:53, dest: 100.108.1.232:53 26-114
2020-06-17 10:52:36.141050268 +0000 UTC m=+16.001560969 - from: 100.119.69.151:56314 - to 100.70.132.168:53, dest: 100.123.202.248:53 22-110
2020-06-17 10:52:36.14182698 +0000 UTC m=+16.002337699 - from: 172.21.3.94:53230 - to 172.21.3.178:30953, dest: 100.122.82.114:32000 24-117
2020-06-17 10:52:36.227754534 +0000 UTC m=+16.088265287 - from: 100.119.69.139:44926 - to 100.70.132.168:53, dest: 100.100.233.255:53 23-110
2020-06-17 10:52:36.227955518 +0000 UTC m=+16.088466295 - from: 100.119.69.150:45236 - to 100.70.132.168:53, dest: 100.127.61.137:53 25-113
2020-06-17 10:52:36.228268838 +0000 UTC m=+16.088779553 - from: 100.119.69.137:40860 - to 100.70.132.168:53, dest: 100.123.202.248:53 83-116
2020-06-17 10:52:36.22937938 +0000 UTC m=+16.089890111 - from: 100.119.69.150:42116 - to 100.70.132.168:53, dest: 100.115.9.245:53 73-105
2020-06-17 10:52:36.230529612 +0000 UTC m=+16.091040333 - from: 100.119.69.139:48738 - to 100.70.132.168:53, dest: 100.123.202.248:53 83-116
2020-06-17 10:52:36.230725993 +0000 UTC m=+16.091236695 - from: 100.119.69.150:41974 - to 100.70.132.168:53, dest: 100.108.1.232:53 72-104
2020-06-17 10:52:36.231374395 +0000 UTC m=+16.091885122 - from: 100.119.69.137:38268 - to 100.70.132.168:53, dest: 100.123.202.248:53 78-114
2020-06-17 10:52:36.232685018 +0000 UTC m=+16.093195767 - from: 100.119.69.139:45838 - to 100.70.132.168:53, dest: 100.127.61.137:53 25-113
2020-06-17 10:52:36.233653104 +0000 UTC m=+16.094163844 - from: 100.119.69.150:45150 - to 100.70.132.168:53, dest: 100.115.9.245:53 77-112
2020-06-17 10:52:36.234143446 +0000 UTC m=+16.094654180 - from: 100.119.69.137:37584 - to 100.70.132.168:53, dest: 100.123.202.248:53 76-112
2020-06-17 10:52:36.234372432 +0000 UTC m=+16.094883204 - from: 100.119.69.139:42734 - to 100.70.132.168:53, dest: 100.108.1.232:53 73-105
2020-06-17 10:52:36.234412896 +0000 UTC m=+16.094923597 - from: 100.119.69.139:46588 - to 100.70.132.168:53, dest: 100.115.9.245:53 79-114
2020-06-17 10:52:36.234964455 +0000 UTC m=+16.095475190 - from: 100.119.69.139:44146 - to 100.70.132.168:53, dest: 100.123.202.248:53 21-108
2020-06-17 10:52:36.235163443 +0000 UTC m=+16.095674163 - from: 100.119.69.137:35720 - to 100.70.132.168:53, dest: 100.108.1.232:53 20-107
2020-06-17 10:52:36.235279947 +0000 UTC m=+16.095790660 - from: 100.119.69.151:57020 - to 100.70.132.168:53, dest: 100.108.1.232:53 76-111
2020-06-17 10:52:36.237719251 +0000 UTC m=+16.098229966 - from: 100.119.69.151:56218 - to 100.70.132.168:53, dest: 100.108.1.232:53 22-109
2020-06-17 10:52:36.237817513 +0000 UTC m=+16.098328218 - from: 100.119.69.150:45524 - to 100.70.132.168:53, dest: 100.123.202.248:53 26-113
2020-06-17 10:52:36.23815506 +0000 UTC m=+16.098665776 - from: 100.119.69.150:42388 - to 100.70.132.168:53, dest: 100.100.233.255:53 18-105
2020-06-17 10:52:36.23882545 +0000 UTC m=+16.099336163 - from: 100.119.69.150:41976 - to 100.70.132.168:53, dest: 100.100.233.255:53 72-104
```

# Docker images

```
docker pull lablabs/iptables-logger:0.0.1
```

# Kubernetes DaemonSet

This tool can be deployed as DaemonSet in Kubernetes with the configuration below. The log output of individual pods can be that collected using various methods (fluentd into ElasticSearch, etc.)

```
apiVersion: v1
kind: Namespace
metadata:
  name: iptables-logger
---
apiVersion: apps/v1
kind: DaemonSet
metadata:
  namespace: iptables-logger
  name: iptables-logger
  labels:
    app: iptables-logger
spec:
  updateStrategy:
    type: RollingUpdate
    rollingUpdate:
      maxUnavailable: 100
  selector:
    matchLabels:
      app: iptables-logger
  template:
    metadata:
      labels:
        app: iptables-logger
    spec:
      containers:
      - name: logger
        image: lablabs/iptables-logger:0.0.1
        imagePullPolicy: Always
        args: [ "-interval", "15" ]
        resources:
          limits:
            cpu: 250m
          requests:
            cpu: 100m
        securityContext:
          privileged: true
          allowPrivilegeEscalation: true
      hostNetwork: true
      tolerations:
        - effect: NoSchedule
          operator: Exists
```

## Contributing and reporting issues

Feel free to create an issue in this repository if you have questions, suggestions or feature requests.

## License

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

See [LICENSE](LICENSE) for full details.

    Licensed to the Apache Software Foundation (ASF) under one
    or more contributor license agreements.  See the NOTICE file
    distributed with this work for additional information
    regarding copyright ownership.  The ASF licenses this file
    to you under the Apache License, Version 2.0 (the
    "License"); you may not use this file except in compliance
    with the License.  You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing,
    software distributed under the License is distributed on an
    "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
    KIND, either express or implied.  See the License for the
    specific language governing permissions and limitations
    under the License.
