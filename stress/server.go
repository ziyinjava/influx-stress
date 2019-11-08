package stress

import (
	"fmt"
	"github.com/bemyth/influx-stress/client"
	"math/rand"
	"time"
)

/*
速率测试，单表，两个tag，一个field
每个tag1000个值 (0 ~ 999)
*/
const (
	PrefixOfTagValue = "ID"
	NumberOfTagValue = 100000
)

type Server struct {
	c    client.Client
	tags [NumberOfTagValue]string
}

func NewServer(ip, port, username, password string) *Server {
	return &Server{
		c:    client.NewClient(ip, port, username, password),
		tags: genTagValue(),
	}
}
func (s *Server) Run() {
	series := s.genSeries()
	con := len(series) / 20
	for i := 0; i < 20; i++ {
		go func(series []string) {
			nowTime := time.Now().UnixNano()
			for {
				for i := range series {
					pt := fmt.Sprintf("%s value=%.5f %d\n", series[i], rand.Float64()*100, nowTime)
					s.c.Write(pt)
				}
				nowTime += 1
			}
		}(series[i*con : (i+1)*con])
	}
}

func (s *Server) genSeries() [NumberOfTagValue]string {
	var rst [NumberOfTagValue]string
	for i := 0; i < NumberOfTagValue; i++ {
		rst[i] = fmt.Sprintf("stress,tag1=%s,tag2=%s", s.tags[i], s.tags[i])
	}
	return rst
}

func genTagValue() [NumberOfTagValue]string {
	var rst [NumberOfTagValue]string
	for i := 0; i < NumberOfTagValue; i++ {
		rst[i] = fmt.Sprintf("%s%010d", PrefixOfTagValue, i)
	}
	return rst
}
