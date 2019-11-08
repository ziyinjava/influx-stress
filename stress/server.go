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
	NumberOfTagValue = 1000
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
	nowTime := time.Now().UnixNano()
	for {
		pts := s.genPoints(uint64(nowTime))
		nowTime += 1
		for i := range pts {
			s.c.Write(pts[i])
		}
	}
}

func (s *Server) genPoints(timestamp uint64) [NumberOfTagValue]string {
	var rst [NumberOfTagValue]string
	for i := 0; i < NumberOfTagValue; i++ {
		rst[i] = fmt.Sprintf("stress,tag1=%s,tag2=%s value=%.6f %d\n", s.tags[i], s.tags[i], rand.Float64()*100, timestamp)
	}
	return rst
}

func genTagValue() [NumberOfTagValue]string {
	var rst [NumberOfTagValue]string
	for i := 0; i < NumberOfTagValue; i++ {
		rst[i] = fmt.Sprintf("%s%03d", PrefixOfTagValue, i)
	}
	return rst
}
