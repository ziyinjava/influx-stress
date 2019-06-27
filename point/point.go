package point

import (
	"strconv"
	"time"
)

type Point struct {
	M         []byte
	Tag       []byte
	Fields    []byte
	TimeStamp time.Time
}

func MarshalPoints(points []Point) []byte {
	var rst []byte
	for i := range points {
		rst = append(rst, points[i].M...)
		rst = append(rst, ',')

		rst = append(rst, points[i].Tag...)
		rst = append(rst, ' ')

		rst = append(rst, points[i].Fields...)
		rst = append(rst, ' ')

		ts := strconv.FormatInt(int64(points[i].TimeStamp.UnixNano()), 10)

		rst = append(rst, []byte(ts)...)
		rst = append(rst, '\n')

		// 顺道更新时间
		points[i].TimeStamp = points[i].TimeStamp.Add(1 * time.Second)
	}
	//fmt.Printf(string(rst))
	return rst
}
