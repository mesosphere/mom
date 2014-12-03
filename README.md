Mesos on Mesos
===

## Way to short build instructions

Assuming you have GOPATH setup and GOPATH/bin in your PATH.

```
$ git clone git@github.com:mesosphere/mom.git
$ cd mom
$ go get ./...
$ mom
NAME:
   mom - Mesos on Mesos cluster testing

USAGE:
   mom [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   cluster      Cluster utilies
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h           show help
   --version, -v        print the version
```

### Set up environment

Copy mom.json to ~/.mom.json and change variables to match your setup.

```
$ cat ~/.mom.json
{
  "marathon_url": "http://foobar:8080",
  "docker_hub": "barbaz:5000",
  "zookeeper": "zk://localhost:2181/mom/",
  "app_prefix": "/mom",
  "quorum_size": 1,
  "master_count": 3,
  "master_cpu": 0.1,
  "master_mem": 32,
  "master_flags": "",
  "slave_count": 10,
  "slave_cpu": 0.1,
  "slave_mem": 32,
  "slave_flags": ""
}
```

### Launch a cluster

Assuming that you have $GOPATH/bin in your $PATH

```
$ mom cluster launch
Launching cluster id:  81b6ab3a-1d05-4f89-6f30-64db7714c3c4

masters:
        instances:      3
        cpu:            0.100000
        mem:            32

slaves:
        instances:      10
        cpu:            0.100000
        mem:            32
```

### Get details on your new cluster

You can get details about the new cluster by:

```
$ mom cluster status 81b6ab3a-1d05-4f89-6f30-64db7714c3c4
masters:
        srv3.hw.ca1.mesosphere.com:31006
        srv2.hw.ca1.mesosphere.com:31324
        srv4.hw.ca1.mesosphere.com:31959
```

### Tear down cluster

You tear down the cluster by

```
$ mom cluster destroy 81b6ab3a-1d05-4f89-6f30-64db7714c3c4
Cluster session 81b6ab3a-1d05-4f89-6f30-64db7714c3c4 destroyed
```
