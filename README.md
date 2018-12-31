# tfe-go-cli

A Golang CLI app for interacting with the Terraform Enterprise API.

NOTE: This is not an officially supported project. Why not use https://github.com/hashicorp/tfe-cli which is!

# build

```
$ go get -v .
$ go install -v
```

# Run

```
$ tfe-go-cli help
tfe-go-cli is a command line interface for interacting with
Terraform Enterprise.

Usage:
  tfe-go-cli [command]

Available Commands:
  configure   Configure your TFE credentials
  help        Help about any command
  validate    Validate your TFE credentials

Flags:
  -c, --config string   config file (default is $HOME/.tgc.yaml)
  -h, --help            help for tfe-go-cli
  -v, --verbose         enable verbose output

Use "tfe-go-cli [command] --help" for more information about a command.
```
