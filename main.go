package main

import (
	"runtime"

	"github.com/yanndr/img/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
