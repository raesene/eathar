# Eather Architecture

Eathar's goal is to connect to a Kubernetes cluster and pull information of "security interest" which can then be analyzed by a human. The information is pulled from the Kubernetes API and then analyzed by the program. The program is designed to be extensible so that new checks can be added easily.

The definition of "security interest" is fairly broad but can include things like the following:

- What images are running in the cluster?
- what pods have security settings as specified in the [Pod Security Standards](https://kubernetes.io/docs/concepts/security/pod-security-standards/)?
- What users/groups/service accounts have access to privilege escalation permissions as defined in the Kubernetes [RBAC Good Practice](https://kubernetes.io/docs/concepts/security/rbac-good-practices/#privilege-escalation-risks) document?

This list is not exhaustive and is meant to be a starting point for discussion.

## Architecture

We use [cobra](https://github.com/spf13/cobra) for the CLI, [kubernetes/client-go](https://github.com/kubernetes/client-go) for the Kubernetes API and [zerolog](https://github.com/rs/zerolog) for logging.

The cobra CLI files are in their default location of `cmd` and the main program is in `pkg/eathar`. Within the `eathar` package we split the functionality into different files for ease of use.

At the moment we have

- `connection.go` - Handles connection to the Kubernetes API.
- `container.go` - Handles checks related to container images containers generally (but not the PSS ones :) )
- `pss.go` - Handles checks related to the Pod Security Standards
- `rbac.go` - Handles checks related to RBAC
- `reporting.go` - Handles reporting of the results of the checks
