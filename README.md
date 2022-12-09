# Eathar

This is a program designed to quickly pull some interesting security related information from Kubernetes clusters. There are a couple of categories of checks that have been implemented so far.

## PSS

Eathar can check containers running in the cluster for various things that are on the [Pod Security Standards](https://kubernetes.io/docs/concepts/security/pod-security-standards/) List

- hostpid - Provides a list of pods in the cluster configured to use Host PID.
- hostnet - Provides a list of pods in the cluster configured to use Host Networking.
- hostipc - Provides a list of pods in the cluster configured to use Host IPC.
- hostports - Provides a list of containers in the cluster configured to use Host Ports.
- hostpath - Provides a list of pods that mount host path volumes.
- hostprocess - Provides a list of Windows pods and containers that run with hostprocess rights.
- privileged - Provides a list of containers in the cluster configured to be privileged.
- allowprivesc - Provides a list of containers in the cluster configured to allow privilege escalation.
- capadded - Provides a list of containers which have capabilities added over the default set.
- cadropped - Provides a list of containers which have capabilities dropped from the default set.
- seccomp - Look for containers which have no seccomp profile specified or explicitly set unconfined.
- apparmor - Look for containers where the apparmor profile is explicitly set to unconfined.
- procmount - Look for containers with an unmasked proc filesystem mount.
- sysctl - Look for dangerous sysctls being set
- allPSS - Run all configured checks

## Image List

Eathar can also provide a unique list of images used in the cluster, which could be useful for checking versions or feeding into a vulnerability scanner (e.g. trivy)

- imagelist - Provides a list of images used in the cluster.

## Demo

![Eathar Demo](https://user-images.githubusercontent.com/68317/183242375-5420ce90-26aa-4d36-bae0-1583dfec1dd8.gif)

## Running Eathar

Eathar connects to a Kubernetes cluster, it works based on whatever you have your current context set to.

## Reporting

By Default reporting is to STDOUT in text format. There's a couple of options for changing that

`-j` will output to JSON
`-f <FILENAME>` sends output to a file (`.txt` or `.json` gets appended to the name specified)

## Name

An Eathar is a small boat in Scots Gaelic.