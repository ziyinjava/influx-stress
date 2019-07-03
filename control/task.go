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

// Task 任务，每个 表 是一个任务
type Task struct {
	client      *http.Client
	measurement string
	cfg         config.Config
	threads     chan *thread.Thread
	//writeReq    int64
	writeOK int64
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

// NewTask 初始化一个任务
func NewTask(client *http.Client, measurement string, cfg config.Config) *Task {
	threads := make(chan *thread.Thread, cfg.Concurrent)
	for i := 0; i < cfg.Concurrent; i++ {
		threads <- thread.New(client, cfg, initPoint(measurement, cfg.BatchSize, i))
	}

	return &Task{
		client:      client,
		measurement: measurement,
		cfg:         cfg,
		threads:     threads,
	}
}

// func (t *Task) log() {
// 	//t := time.Now().Format("2006:01:02:15:04:05")
// 	// var v int64
// 	// //var lv int64
// 	// //var pps int64
// 	// start := time.Now()
// 	timer := time.NewTimer(2 * time.Second)
// 	for {
// 		timer.Reset(2 * time.Second)
// 		select {
// 		case <-timer.C:
// 		}
// 		v = atomic.LoadInt64(&t.writeOK)
// 		//pps = (v - lv) / 2
// 		//lv = v
// 		//data := newflogData(t.cfg.LogName, t.measurement, int64(t.cfg.BatchSize), int64(t.cfg.Concurrent), pps, time.Now())
// 		//fmt.Println(data)
// 		// avg := float64(v) / time.Since(start).Seconds()
// 		// fmt.Printf("[%s]\t%f\n", t.measurement, avg)
// 		//go t.send([]byte(data))
// 	}
// }

// Run 任务运行
func (t *Task) Run() {
	var th *thread.Thread
	timer1 := time.NewTimer(time.Duration(t.cfg.Time) * time.Second)

	// go t.log()
	for {
		select {
		case <-timer1.C:
			return
		default:
			th = <-t.threads
			go func(th *thread.Thread) {
				//atomic.AddInt64(&t.writeReq, int64(t.cfg.BatchSize))
				th.Send()
				t.threads <- th
				atomic.AddInt64(&t.writeOK, int64(t.cfg.BatchSize))
			}(th)
		}
	}
}

// func (t *Task) send(point []byte) {
// 	//b := point.MarshalPointsThenUpdate(t.points)
// 	buf := bytes.NewBuffer(point)
// 	req, err := http.NewRequest("POST", fmt.Sprintf("%s/write", "http://172.16.22.191:8086"), buf)
// 	para := req.URL.Query()
// 	para.Set("db", "stress")
// 	req.URL.RawQuery = para.Encode()
// 	if err != nil {
// 		panic(err)
// 	}
// 	resp, err := t.client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
// 		fmt.Printf("Bad request status code: %d \n", resp.StatusCode)
// 	}
// 	resp.Body.Close()
// }
