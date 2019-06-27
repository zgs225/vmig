Vmig
===

## Installation

``` bash
go get -u github.com/zgs225/vmig
```

## Example

``` bash
# Initialize
vmig init

# Create a version and set to default
vmig create-version v1.0.0 -d

# Create a migration
vmig new create_table_users

# Apply migration
vmig up

# Rollback
vmig down
```

## Usage

```
A wrapper for golang-migrate that let it support version managed migrations and
multiple environments.

Usage:
  vmig [command]

Available Commands:
  config         Config settings
  create-version Create a new version
  down           Rollback migrations
  help           Help about any command
  init           Init vmig environment
  new            Create a migration file in default version directory.
  up             Apply all or given N up migration files.
  version        Show vmig version

Flags:
      --config string   config file (default is $PWD/.vmig.yaml)
  -h, --help            help for vmig
      --verbose         Output debug message.

Use "vmig [command] --help" for more information about a command.
```
