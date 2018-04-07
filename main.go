package main

import (
	"runtime"

	"github.com/yeoji/eventgen/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	commands.Execute()
}
