# img

## About
This is a Image CLI written in GO. This is a learning excercice for me. The goal is to learn how to build a CLI. And my wife needed a simple tool to resize images quickly. So I decided to wrote it myself :). So Far the client only have one command: Resize. I may add more features later.

I used Cobrahttps://github.com/spf13/cobra for the CLI and resize lib: https://github.com/nfnt/resize

## Usage
Usage:
  img [flags]
  img [command]

Available Commands:
  help        Help about any command
  resize      Resize the images to a percentage or a value.
  version     Print the version number of img

Flags:
  -h, --help   help for img

Use "img [command] --help" for more information about a command.
