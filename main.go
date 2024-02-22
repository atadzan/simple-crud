package main

import (
	"flag"
	"github.com/atadzan/simple-crud/app"
	"github.com/atadzan/simple-crud/third_party/errorx"
	"runtime"
)

var configPath *string

func init() {
	// app configuration path flag
	configPath = flag.String("config", "./configs/debug/dashboard-service.yaml", "Default config path")
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
