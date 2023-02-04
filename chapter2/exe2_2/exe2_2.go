// 通用的单位转换程序
// 从命令行或者标准输入获取数字
// 将每一个数字转换为摄氏度和华氏温度，英寸和米表示的长度，磅和千克表示的重量
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Celsius float64
type Fahrenheit float64
type Inch float64
type Meter float64
type Pound float64
type Kilogram float64

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g℉", f) }
func (i Inch) String() string       { return fmt.Sprintf("%gin", i) }
func (m Meter) String() string      { return fmt.Sprintf("%gm", m) }
func (p Pound) String() string      { return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string   { return fmt.Sprintf("%gkg", k) }

// 华氏度与摄氏度转换
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 9 / 5) }

// 英寸和米转换
func InToM(i Inch) Meter { return Meter(Decimal(float64(i * 0.0254))) }
func MToIn(m Meter) Inch { return Inch(Decimal(float64(m / 0.0254))) }

// 磅和千克转换
func LbToKg(p Pound) Kilogram { return Kilogram(Decimal(float64(p * 0.4535924))) }
func KgToLb(k Kilogram) Pound { return Pound(Decimal(float64(k * 2.2046226))) }

// 保留四位小数
func Decimal(num float64) float64 {
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", num), 64)
	return num
}

func main() {
	para := os.Args[1:]
	if len(para) == 0 {
		//如果命令行参数列表为空，从标准输入获取数字
		fmt.Println("请输入要转换的数字，多个数字以空格隔开，按回车结束输入")
		reader := bufio.NewReader(os.Stdin)
		s, err := reader.ReadString('\r') //以回车符结束输入
		if err != nil {
			fmt.Println("reader error", err)
		}
		num := strings.SplitAfter(s, " ")
		for _, n := range num {
			t, err := strconv.ParseFloat(strings.TrimSpace(n), 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "kgtolb: %v\n", err)
				os.Exit(2)
			}
			convert(t)
		}
	} else {
		for _, arg := range para {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "kgtolb: %v\n", err)
				os.Exit(3)
			}
			convert(t)
		}
	}
}

func convert(t float64) {
	c, f, i, m, p, k := Celsius(t), Fahrenheit(t), Inch(t), Meter(t), Pound(t), Kilogram(t)
	fmt.Printf("%s = %s, %s = %s\n", f, FToC(f), c, CToF(c))
	fmt.Printf("%s = %s, %s = %s\n", i, InToM(i), m, MToIn(m))
	fmt.Printf("%s = %s, %s = %s\n", p, LbToKg(p), k, KgToLb(k))
	fmt.Println()
}
