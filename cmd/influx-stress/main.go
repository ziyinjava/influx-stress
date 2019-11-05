package main

import (
	"github.com/bemyth/influx-stress/cmd/influx-stress/cmd"
	_ "net/http/pprof"
)

func main() {
	cmd.Execute()
}
