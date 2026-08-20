# MultiGIT

MultiGIT is a GIT helper for executing a subset of git commands on multiple repositories.

$${\color{red}Note! \space This \space tool \space does \space not \space guarantee \space flawless \space functioning \space - \space use \space on \space your \space own \space risk!}$$

## mgit -h
```
mgit is a GIT helper for executing a subset of git commands on multiple repositories.

Usage:
  mgit [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      Manage configuration
  help        Help about any command
  list        Lists GIT reposities with their currently checked out branches in the given directory.
  pull        Executes git pull in all git repositories in the given directory
  purge       Delete local branches in all git repositories in the given directory that don't have a corresponding remote branch

Flags:
  -h, --help   help for mgit

Use "mgit [command] --help" for more information about a command.
```

## mgit config
```
Commands to manage application configuration

Usage:
  mgit config [command]

Available Commands:
  read        Read configuration
  write       Write configuration values

Flags:
  -h, --help   help for config

Use "mgit config [command] --help" for more information about a command.
```

## mgit list
```
Lists GIT reposities with their currently checked out branches in the given directory.

Usage:
  mgit list  [directory] [flags]

Flags:
  -b, --branches   List all local branches.
  -h, --help       help for list
```

## mgit pull
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

## mgit purge
$${\color{red}Careful! \space This \space will \space delete \space local \space branches - \space make \space sure \space you \space really \space want \space that!}$$
```
Delete all local branches in all git repositories in the given directory that don't have a corresponding remote branch

Usage:
  mgit purge [flags]

Flags:
  -x, --exclude strings   Exclude the repositories with the names given here from purging.
  -h, --help              help for purge
```

## configuration
mgit looks for the configuration in the current directory, or under `$HOME/.mgit/`

the file should be called `config` or `config.yaml`

the configuration is only necessary if some kind of authentication with the GIT server is required
```
git:
  - remote:
      name: origin 
      host: github.com

  - remote
      name: origin
      host: gitlab.com
    auth:
      basic:
        username: xyz
        password: xxx

  - remote
      name: custom
      host: enterprise.host
    auth:
      token:
        token: the-token-xxx
```


