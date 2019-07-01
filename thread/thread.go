package thread

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/bemyth/influx-stress/config"

	"github.com/bemyth/influx-stress/point"
)

// Thread 任务线程
type Thread struct {
	client   *http.Client
	url      string
	database string
	points   []point.Point
}

// New 初始化一个线程
func New(client *http.Client, cfg config.Config, points []point.Point) *Thread {
	return &Thread{
		client:   client,
		url:      cfg.URL,
		database: cfg.DataBase,
		points:   points,
	}
}

// Send 发送一次数据
func (t *Thread) Send() {
	b := point.MarshalPointsThenUpdate(t.points)
	buf := bytes.NewBuffer(b)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/write", t.url), buf)
	para := req.URL.Query()
	para.Set("db", t.database)
	req.URL.RawQuery = para.Encode()
	if err != nil {
		panic(err)
	}
	resp, err := t.client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		fmt.Printf("Bad request status code: %d \n", resp.StatusCode)
	}
	resp.Body.Close()
}
