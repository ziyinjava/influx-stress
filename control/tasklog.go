package control

import (
	"fmt"
	"strconv"
	"time"
)

func newflogData(idx, m string, batchSize int64, concurrent int64, writeSpeed int64, ts time.Time) string {
	return fmt.Sprintf("%s,m=%s,batchSize=%s,concurrent=%s speed=%d %d\n", idx, m, strconv.FormatInt(batchSize, 10), strconv.FormatInt(concurrent, 10), writeSpeed, ts.UnixNano())
}
