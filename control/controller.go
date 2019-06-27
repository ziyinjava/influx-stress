package control

import (
	"net/http"
	"sync"

	"github.com/bemyth/influx-stress/config"
)

// Controller 写入的整体控制
type Controller struct {
	cfg    config.Config
	tasks  []*Task
	client *http.Client
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

// Exec 控制器开始执行
func (c *Controller) Exec() {
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
