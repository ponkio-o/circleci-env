# ccienv

A management tool for CircleCI Project's environment variables.  
(Currently, only available on GitHub)

## Installation

```
go install github.com/threepipes/circleci-env/cmd/ccienv@1.2.1
```

### Uninstallation

```
rm $(which ccienv)
```

## Setup

```
ccienv config init
```

Set these variables.

- CircleCI API Token
    - A personal API token of CircleCI
- GitHub organization
    - GitHub organization name or GitHub username of your repository

## Run

```
ccienv -r <your_repo_name> <cmd> [<args>]
```

If `-r <your_repo_name>` is omitted, the current directory name is used as the target repository name.

### Example

```
ccienv -r circleci-env ls
```

## Help

```
ccienv -h
```
