# Eathar

This is a program designed to quickly pull some interesting security related information from Kubernetes clusters. It's primarily written as a learning project for me to practice with Go, so don't expect good code or brilliant functionality!

## Running Eathar

Eathar connects to a Kubernetes cluster, it'll look for a kubeconfig in `~/.kube/config` and use that if found. If you want to specify a different location use the `--kubeconfig` file. within that file it'll use the `current-context` (allowing others is on the ToDo list)

At the moment here's the things it can check for

- hostpid - Provides a list of pods in the cluster configured to use Host PID.
- hostnet - Provides a list of pods in the cluster configured to use Host Networking.
- hostipc - Provides a list of pods in the cluster configured to use Host IPC.
- hostports - Provides a list of containers in the cluster configured to use Host Ports.
- privileged - Provides a list of containers in the cluster configured to be privileged.
- allowprivesc - Provides a list of containers in the cluster configured to allow privilege escalation.
- capadded - Provides a list of containers which have capabilities added over the default set.
- cadropped - Provides a list of containers which have capabilities dropped from the default set.
- all - Run all configured checks

## Reporting

There's two options for reporting. The first (which is the default) is just to report to STDOUT . The other option specified with `-j` is to output JSON to STDOUT. File output is on the ToDo list :)

## Name

An Eathar is a small boat in Scots Gaelic.