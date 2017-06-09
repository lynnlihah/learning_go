//学习资源：《Go语言编程》
package main

import (
	"fmt"
)

//变量定义
var v1 int = 10
var v2 string = "str"
var v3 [10]int
var v4 []int
var point struct {
	x float64
	y float64
	z float64
}
var p *int
var v7 map[string]int
var v8 func(a int) int

var (
	j = 10
	i = 1
)

//匿名变量
func GetName() (firstName, lastName, nickName string) {
	return "May", "Chan", "Chibi Maruko"
}

//常量
const Pi float64 = 3.14159265359
const zero = 0.0
const (
	size int64 = 1024
	eof        = -1
)
const u, v float32 = 0, 3 // u = 0.0, v = 3.0，常量的多重赋值
const a, b, c = 3, 4, "foo"
const mask = 1 << 3 //常量定义的右值也可以是一个在编译期运算的常量表达式,但运行期不行

func main() {
	fmt.Println("Hello world!")
	//匿名变量
	_, _, nickName := GetName() // := 不需声明nickName
	fmt.Println(nickName)       //output chibi Maruko

	fmt.Println(mask, u, v, c, Pi)
}
