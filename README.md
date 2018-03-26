# img

## About
This is a Image CLI written in Go. 
My wife needed a simple tool to resize images quickly, so I decided to write it myself :). This is a learning excercise; the goal is to learn how to build a CLI. So far, the client only has one command: Resize. I may add more features later.

I used [Cobra](https://github.com/spf13/cobra) for the CLI and [resize lib](https://github.com/nfnt/resize) for the image resize functiom.

## Instalation
```
go get -v github.com/yanndr/img
```

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

### resize usage
img resize [files ...] [flags]

Flags:
  -f, --format string   Force the format of the output: png or jpg. if empty it will keep the input image format.
  
  -h, --help            help for resize
  
  -o, --out string      Output directory for the images defautl: out. (default "out")
  
  -s, --size int        Size in % of the original image. (default 50)
