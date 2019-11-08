package main

import (
	"github.com/bemyth/influx-stress/cmd/influx-stress/cmd"
	"github.com/bemyth/influx-stress/pprof"
)

func main() {
	go pprof.ServerHttp()
	cmd.Execute()
}
