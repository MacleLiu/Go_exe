// 添加开氏温度与摄氏度的转换
package main

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilongC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g℃", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g℉", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }

//华氏度与摄氏度转换
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 9 / 5) }

//开尔文与摄氏度转换
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

func main() {
	fmt.Println(CToF(BoilongC))
	fmt.Println(CToK(BoilongC))
}
