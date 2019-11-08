package pprof

import (
	"net/http"
	_ "net/http/pprof"
)

func ServerHttp() {
	http.ListenAndServe("localhost:6060", nil)
}
