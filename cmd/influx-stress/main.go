package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bemyth/influx-stress/config"

	"github.com/BurntSushi/toml"
	"github.com/bemyth/influx-stress/control"
)

var (
	path string
)

func init() {
	flag.StringVar(&path, "config", "./config.conf", "--config=$path")
	flag.Parse()
}

func main() {
	var cfg config.Config
	_, err := toml.DecodeFile(path, &cfg)
	if err != nil {
		fmt.Println("decode file: ", err.Error())
		os.Exit(1)
	}
	c := control.New(cfg)
	c.Run()
}
