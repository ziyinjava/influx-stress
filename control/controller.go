package control

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/bemyth/influx-stress/config"
)

// Controller 写入的整体控制
type Controller struct {
	cfg    config.Config
	tasks  []*Task
	client *http.Client
}

func (c *Controller) log() {
	//t := time.Now().Format("2006:01:02:15:04:05")
	var v int64
	var lv int64
	var pps int64
	timer := time.NewTimer(2 * time.Second)
	for {
		timer.Reset(2 * time.Second)
		select {
		case <-timer.C:
		}
		v = 0
		for i := range c.tasks {
			v += atomic.LoadInt64(&c.tasks[i].writeOK)
		}
		pps = (v - lv) / 2
		lv = v
		fmt.Printf("write speed: %d \n", pps)
		//data := newflogData(t, int64(c.cfg.BatchSize), int64(c.cfg.Concurrent), pps, time.Now())
		//fmt.Println(data)
	}
}

// New  初始化一个控制器
func New(cfg config.Config) *Controller {
	c := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
			MaxConnsPerHost:     1000,
			MaxIdleConns:        500,
		},
	}
	var tasks []*Task
	for i := 0; i < len(cfg.Measurements); i++ {
		tasks = append(tasks, NewTask(c, cfg.Measurements[i], cfg))
	}
	return &Controller{
		cfg:    cfg,
		tasks:  tasks,
		client: c,
	}

}

func pprof() {
	http.ListenAndServe(":6060", nil)
}

// Exec 控制器开始执行
func (c *Controller) Exec() {
	go c.log()
	var wg sync.WaitGroup
	for i := range c.tasks {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.tasks[i].Run()
		}(i)
	}
	wg.Wait()
}

// func (c *Controller) send(point []byte) {
// 	//b := point.MarshalPointsThenUpdate(t.points)
// 	buf := bytes.NewBuffer(point)
// 	req, err := http.NewRequest("POST", fmt.Sprintf("%s/write", "http://172.16.22.191:8086"), buf)
// 	para := req.URL.Query()
// 	para.Set("db", "stress")
// 	req.URL.RawQuery = para.Encode()
// 	if err != nil {
// 		panic(err)
// 	}
// 	resp, err := c.client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
// 		fmt.Printf("Bad request status code: %d \n", resp.StatusCode)
// 	}
// 	resp.Body.Close()
// }
