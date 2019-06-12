package types

import "bytes"

var multipleLineString = `line1
line2
lin3`

/**
高性能的拼接字符串
 */
func multipleString(str string, times int) string {
	var buffer bytes.Buffer
	for i := 0; i < times; i++ {
		buffer.WriteString(str)
	}
	return buffer.String()
}