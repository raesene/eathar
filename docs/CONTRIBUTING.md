# Contributing to Eathar

Eathar is very much a work in progress, initially started as a tool to help me learn Go :)

That said, contributions are welcome!

## Issues

If you notice a bug, please fill out an [issue on Github](https://github.com/raesene/eathar/issues)

## Pull Requests

There's a couple of areas where contributions would be welcome:

- New checks. This can be anything that you think would be useful to check for security in a Kubernetes cluster. Please see the [architecture](docs/architecture.md) document for more information on how checks are structured. The general principal at the moment is to pull the information and report it, rather than trying to assess "bad/good" or "pass/fail" as that is quite situational.
- Improved/changed reporting. At the moment reporting is pretty basic and utilitarian. Things like HTML reports or better text output would be nice.