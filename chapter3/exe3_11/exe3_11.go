// comma向表示十进制数字的字符串中插入千分位分隔符
// 可处理浮点数，以及带有可选正负号的数字
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer
	var f string
	//处理正负号
	if s[0] == '-' || s[0] == '+' {
		fmt.Fprintf(&buf, "%s", string(s[0]))
		s = s[1:]
	}
	//处理小数部分
	if strings.Contains(s, ".") {
		i := strings.Index(s, ".")
		f = s[i:]
		s = s[:i]
	}

	n := len(s)
	//整数部分不超过三个数字的，不做处理
	if n <= 3 {
		buf.WriteString(s + f)
		return buf.String()
	}
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%s", string(s[i]))
	}
	buf.WriteString(f)
	return buf.String()
}
func main() {
	fmt.Println(comma("+2343564561.289347234"))
}
