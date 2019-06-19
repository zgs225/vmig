Vmig
===

## Usage

A wrapper for golang-migrate that let it support version managed migrations and
multiple environments.

Usage:
  vmig [command]

Available Commands:
  create-version Create a new version
  help           Help about any command
  init           Init vmig environment
  new            Create a migration file in default version directory.

Flags:
      --config string   config file (default is $PWD/.vmig.yaml)
  -h, --help            help for vmig
  -v, --verbose         Output debug message.

Use "vmig [command] --help" for more information about a command.
