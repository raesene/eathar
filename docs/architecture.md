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


## Check Structure

Typically a check will connect to a Kubernetes cluster, pull some kind of resource(s) and then extract the information of interest. The information is then reported to the user. This is then connected to an external command by using `cobra-cli`. 

Creating a new check would go through the following rough process

1. `cobra-cli add capdropped -p pssCmd` - This creates a new file in `cmd` which is a sub-command of the `pss` command. All checks should be grouped under a top-level command note that the `-p` parameter has the name of the top level command with `Cmd` at the end. At the moment we have three
  - `pss` - Pod Security Standards
  - `info` - General information checks
  - `rbac` - RBAC checks
2. in the cmd file (`cmd/capdropped.go`) add short and long documentation for the command, and in the `run` function add the code to run the check. for example:
```go
		capdropped := eathar.DroppedCapabilities(options)
		eathar.ReportPSS(capdropped, options, "Dropped Capabilities")
```
3. Create a function in the `eathar` package to run the check. The function should be placed in the file that corresponds to the top-level command.

4. Add the commands from the `cmd` file to whichever top-level command it belongs to (e.g. `cmd/pss.go` for the above). The top level commands should run all their sub-commands. We can't automate this process with cobra at the moment as auto-execution is only supported from `root` (per [this issue](https://github.com/spf13/cobra/issues/1526))