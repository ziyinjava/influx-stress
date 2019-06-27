package control

import (
	"net/http"
	"sync"

	"github.com/bemyth/influx-stress/config"
)

type Controller struct {
	cfg   config.Config
	tasks []*Task
}

func New(cfg config.Config) *Controller {
	c := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
			MaxConnsPerHost:     1000,
			MaxIdleConns:        500,
		},
	}
	var tasks []*Task
	for i := 0; i < len(cfg.Prods); i++ {
		tasks = append(tasks, NewTask(c, cfg.Prods[i], cfg))
	}
	return &Controller{
		cfg:   cfg,
		tasks: tasks,
	}

}
func (c *Controller) Run() {
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
