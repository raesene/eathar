# Eathar

This is a program designed to quickly pull some interesting security related information from Kubernetes clusters. There are a couple of categories of checks that have been implemented so far.

## PSS

Eathar can check containers running in the cluster for various things that are on the [Pod Security Standards](https://kubernetes.io/docs/concepts/security/pod-security-standards/) List.

To run all checks just use the top-level `pss` command. To run a specific check use the name of the check below as the subcommand to `pss`. For example to run the hostpid command you would run `eathar pss hostpid`.

- `hostpid` - Provides a list of pods in the cluster configured to use Host PID.
- `hostnet` - Provides a list of pods in the cluster configured to use Host Networking.
- `hostipc` - Provides a list of pods in the cluster configured to use Host IPC.
- `hostports` - Provides a list of containers in the cluster configured to use Host Ports.
- `hostpath` - Provides a list of pods that mount host path volumes.
- `hostprocess` - Provides a list of Windows pods and containers that run with hostprocess rights.
- `privileged` - Provides a list of containers in the cluster configured to be privileged.
- `allowprivesc` - Provides a list of containers in the cluster configured to allow privilege escalation.
- `capadded` - Provides a list of containers which have capabilities added over the default set.
- `cadropped` - Provides a list of containers which have capabilities dropped from the default set.
- `seccomp` - Look for containers which have no seccomp profile specified or explicitly set unconfined.
- `apparmor` - Look for containers where the apparmor profile is explicitly set to unconfined.
- `procmount` - Look for containers with an unmasked proc filesystem mount.
- `sysctl` - Look for dangerous sysctls being set
- `allPSS` - Run all configured checks

## Info Checks

Eathar also has some general cluster information checks. You can run all of these using the `info` command, or you can run a specific check using the name of the check below as the subcommand to `info`. For example to run the imagelist command you would run `eathar info imagelist`.

- `imagelist` - Provides a list of images used in the cluster.

## RBAC

Eather can also provide some information about how RBAC is configured in the cluster, which could be useful for checking if there are any roles or clusterroles that are overly permissive. The goal is to cover the privilege escalation permissions from the Kubernetes [RBAC Good Practice](https://kubernetes.io/docs/concepts/security/rbac-good-practices/#privilege-escalation-risks) document.

You can run all of these using the `rbac` command, or you can run a specific check using the name of the check below as the subcommand to `rbac`. For example to run the clusteradminusers command you would run `eathar rbac clusteradminusers`.
 
 - `clusteradminusers` - Provides a list of users/groups/service accounts who have the cluster-admin clusterrole.
 - `getsecretsuser` - Provides a list of users/groups/service accounts who have `GET` or `LIST` access to secrets at the cluster level.
 - `persistentvolumecreationuser` - Provides a list of users/groups/service accounts who have `CREATE` access to persistentvolumes at the cluster level. 
 - `impersonateuser` - Provides a list of users/groups/service accounts who have `impersonate` access to other users/groups/service accounts at the cluster level.
 - `binduser` - Provides a list of users/groups/service accounts who have `bind` access to clusterroles at the cluster level.
 - `escalate` - Provides a list of users/groups/service accounts who have `escalate` access to clusterroles at the cluster level.
 - `validatingwebhookuser` - Provides a list of users/groups/service accounts who have `create`,  `update`, `patch`, or `delete` access to validatingwebhookconfigurations at the cluster level.
 - `mutatingwebhookuser` - Provides a list of users/groups/service accounts who have `create`,  `update`, `patch`, or `delete` access to mutatingwebhookconfigurations at the cluster level.


## Demo

![Eathar Demo](https://user-images.githubusercontent.com/68317/183242375-5420ce90-26aa-4d36-bae0-1583dfec1dd8.gif)

## Running Eathar

Eathar connects to a Kubernetes cluster, it works based on whatever you have your current context set to.

## Reporting

By Default reporting is to STDOUT in text format. There's a couple of options for changing that

`-j` will output to JSON
`--htmlrep` will output to HTML
`-f <FILENAME>` sends output to a file (`.txt`, `.html` or `.json` gets appended to the name specified)

The HTML report outputs basic tables which look like this :-

![htmlreport](https://user-images.githubusercontent.com/68317/216761034-4210f551-baa9-4b55-bc50-5f832de86e53.png)

## Architecture

The `architecture.md` file in the `docs` directory has some notes on structure and design decisions.

## Name

An Eathar is a small boat in Scots Gaelic.