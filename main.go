package main

import (
	"runtime"

	"github.com/Secoura/eventgen/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	commands.Execute()
}
