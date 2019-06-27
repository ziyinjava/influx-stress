package point

import (
	"strconv"
	"time"
)

// Point 数据点
type Point struct {
	M         []byte
	Tag       []byte
	Fields    []byte
	TimeStamp time.Time
}

// MarshalPointsThenUpdate 将数据点转换成 line protocal 格式的字节流，同时更新这些数据点，以便下一次使用
func MarshalPointsThenUpdate(points []Point) []byte {
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
		points[i].TimeStamp = points[i].TimeStamp.Add(10 * time.Millisecond)
	}
	//fmt.Printf(string(rst))
	return rst
}
