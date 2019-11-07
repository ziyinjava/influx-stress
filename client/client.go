package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	BatchPoints                 = 1 << 11
	BatchDuration time.Duration = 200 * time.Millisecond
	DataBase                    = "influxstress"
)

type Client interface {
	Write(pt string) error
}

type client struct {
	c     *http.Client
	timer *time.Timer
	pts   chan string

	ip, port, username, password string

	counts uint64
}

func (c *client) logSpeed() {
	start := time.Now()
	for {
		time.Sleep(2 * time.Second)

	}
}

/*
 写入策略，每200ms 写一次，或者每 1 << 11 (2048) 条写一次
*/

func (c *client) Write(pt string) error {
	c.pts <- pt
	return nil
}
func (c *client) run() {
	for {
		var count = 0
		c.timer.Reset(BatchDuration)
		var pts []string
		for count < BatchPoints {
			select {
			case <-c.timer.C:
				goto WRITEPOINTS
			case pt := <-c.pts:
				pts = append(pts, pt)
			}
		}
	WRITEPOINTS:
		if len(pts) != 0 {
			c.writePoints(pts)
		}
	}
}
func (c *client) writePoints(pts []string) {
	var bpts []byte
	for i := range pts {
		bpts = append(bpts, []byte(pts[i])...)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s:%s/write", c.ip, c.port), bytes.NewBuffer(bpts))
	if err != nil {
		panic(err)
	}
	params := url.Values{}
	params.Set("db", DataBase)
	params.Set("username", c.username)
	params.Set("password", c.password)
	req.URL.RawQuery = params.Encode()

	resp, err := c.c.Do(req)
	if err != nil {
		panic(err)
	}
	// TODO 返回校验
	switch resp.StatusCode {
	case http.StatusNoContent, http.StatusOK:
	default:
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(b))
		panic(fmt.Sprintf("http response status code is [%d]\n", resp.StatusCode))

	}
}

func NewClient(ip, port, username, password string) *client {
	c := &client{
		c: &http.Client{
			Transport: &http.Transport{
				DisableKeepAlives:   false,
				MaxIdleConns:        10,
				MaxIdleConnsPerHost: 10,
				MaxConnsPerHost:     10,
			},
		},
		pts:      make(chan string, BatchPoints),
		timer:    time.NewTimer(BatchDuration),
		ip:       ip,
		port:     port,
		username: username,
		password: password,
	}
	go c.run()
	return c
}
