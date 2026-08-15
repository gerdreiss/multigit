# MultiGIT

MultiGIT is a GIT helper for executing a subset of git commands on multiple repositories.

```
Usage:
  mgit [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        Lists GIT reposities with their currently checked out branches in the given directory.
  pull        Executes git pull in all git repositories in the given directory

Flags:
  -h, --help   help for mgit
```
```
Lists GIT reposities with their currently checked out branches in the given directory.

Usage:
  mgit list  [directory] [flags]

Flags:
  -h, --help   help for list
```
```
Executes git pull in all git repositories in the given directory

Usage:
  mgit pull [flags]

Flags:
Executes git pull in all git repositories in the given directory

Usage:
  mgit pull [flags]

Flags:
  -d, --default           Check out the default branch (e.g. main) before pulling the latest changes.
  -x, --exclude strings   Exclude the repositories with the names given here from pulling.
  -f, --force             Force pulling or checking out the default branch when current branch is dirty.
  -h, --help              help for pull
```