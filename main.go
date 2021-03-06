package main

import (
	"runtime"

	"github.com/secoura/eventgen/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	commands.Execute()
}
