package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
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
	fmt.Printf("%+v\n", cfg)
	c := control.New(cfg)

	go pprof()
	c.Exec()
}

func pprof() {
	http.ListenAndServe(":6060", nil)
}
