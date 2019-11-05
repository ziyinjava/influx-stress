package main

import (
	"flag"
	"github.com/spf13/cobra/cobra/cmd"
	_ "net/http/pprof"
)

var (
	path string
)

func init() {
	flag.StringVar(&path, "config", "./config.conf", "--config=$path")
	flag.Parse()
}

func main() {

	cmd.Execute()
}
