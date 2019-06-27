package control

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/bemyth/influx-stress/point"

	"github.com/bemyth/influx-stress/config"
	"github.com/bemyth/influx-stress/thread"
)

type Task struct {
	client   *http.Client
	prod     string
	cfg      config.Config
	threads  chan *thread.Thread
	writeReq int64
}

func initPoint(prod string, size int, idx int) (points []point.Point) {
	points = make([]point.Point, size)
	t := time.Now()
	for i := range points {
		t = t.Add(1 * time.Nanosecond)
		points[i].M = []byte(prod)
		points[i].Tag = []byte(fmt.Sprintf("testtag=tag-%d-%d", idx, rand.Intn(100)))
		points[i].Fields = []byte(fmt.Sprintf("testfield=\"field-%d\"", i))
		points[i].TimeStamp = t
	}
	return
}
func NewTask(client *http.Client, prod string, cfg config.Config) *Task {
	threads := make(chan *thread.Thread, cfg.Concurrent)
	for i := 0; i < cfg.Concurrent; i++ {
		threads <- thread.New(client, cfg, initPoint(prod, cfg.BatchSize, i))
	}

	return &Task{
		client:  client,
		prod:    prod,
		cfg:     cfg,
		threads: threads,
	}
}

func (t *Task) log() {
	var v int64
	var pps int
	for i := 0; i < t.cfg.Time; i++ {
		time.Sleep(1 * time.Second)
		v = atomic.LoadInt64(&t.writeReq)
		pps = int(v) / ((i + 1) * 10000)
		fmt.Printf("[%s]\t%d\tw p/s\n", t.prod, pps)
	}
}
func (t *Task) Run() {
	var th *thread.Thread
	timer1 := time.NewTimer(time.Duration(t.cfg.Time) * time.Second)

	go t.log()
	for {
		select {
		case <-timer1.C:
			return
		default:
			th = <-t.threads
			go func(th *thread.Thread) {
				th.Send()
				t.threads <- th
				atomic.AddInt64(&t.writeReq, int64(t.cfg.BatchSize))
			}(th)
		}
	}
}
