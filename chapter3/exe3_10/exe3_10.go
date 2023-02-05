// comma向表示十进制数字的字符串中插入千分位分隔符
// 非递归实现comma函数，并运用bytes.Buffer
package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%s", string(s[i]))
	}
	return buf.String()
}
func main() {
	fmt.Println(comma("23454776767"))
}
