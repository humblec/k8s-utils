This program or binary is capable of reading dependencies.yaml
which has been given as input and list out the names and versions
of each components listed over there.

Example outputs:

```
Options:
  -input string
    	Input YAML file path (default "dependencies.yaml")
chumble2TR91:dependency-parser chumble$ 

chumble2TR91:dependency-parser chumble$ go run main.go 
Name: zeitgeist
Version: v0.2.0
----------
Name: cni
Version: 1.4.0
----------
Name: coredns-kube-up
Version: 1.11.1
----------
Name: coredns-kubeadm
Version: 1.11.1
----------
Name: crictl
Version: 1.29.0
----------
Name: protoc
Version: 23.4
----------
Name: etcd
Version: 3.5.12
----------
Name: etcd-image
Version: 3.5.12
----------
Name: node-problem-detector
Version: 0.8.13
----------
Name: golang: etcd release version
Version: 1.20.13
----------
Name: golang: upstream version
Version: 1.22rc2
----------
Name: registry.k8s.io/kube-cross: dependents
Version: v1.30.0-go1.22rc2-bullseye.0
----------
Name: registry.k8s.io/debian-base: dependents
Version: bookworm-v1.0.1
----------
Name: registry.k8s.io/distroless-iptables: dependents
Version: v0.5.0
----------
Name: registry.k8s.io/go-runner: dependents
Version: v2.3.1-go1.22rc2-bookworm.0
----------
Name: registry.k8s.io/pause
Version: 3.9
----------
Name: registry.k8s.io/pause: dependents
Version: 3.9
----------
Name: registry.k8s.io/build-image/setcap: dependents
Version: bookworm-v1.0.1
----------
Name: gcr.io/cadvisor/cadvisor: dependents
Version: v0.47.2
----------
Name: gcb-docker-gcloud: dependents
Version: v20230623-56e06d7c18
----------

