package main

import (
	"runtime"

	"github.com/yanndr/imgresize/cmd"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	cmd.Execute()
}
