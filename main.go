package main

import (
	"flag"
	"runtime"

	"github.com/atadzan/simple-crud/app"
	"github.com/atadzan/simple-crud/third_party/errorx"
)

var configPath *string

func init() {
	// app configuration path flag
	configPath = flag.String("config", "./configs/debug.yaml", "Default config path")
	flag.Parse()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// app initialization
	if err := app.Init(*configPath); err != nil {
		errorx.PrintDetailedError(err)
		return
	}
}
