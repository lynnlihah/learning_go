//学习资源：《Go语言编程》
package main

import (
	"fmt"
	"math"
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
//预定义常量：true、false、iota
//iota比较特殊，可以被认为是一个可被编译器修改的常量，在每一个const关键字出现时被
//重置为0，然后在下一个const出现之前，每出现一次iota，其所代表的数字会自动增1。
const (
	ui float64 = iota * 42 //0 如果两个const的赋值语句的表达式是一样的，那么可以省略后一个赋值表达式。
	vi                     //float64 = iota * 42 //42.0
	wi                     //float64 = iota * 42 //84
)

//枚举
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays // 这个常量没有导出 同Go语言的其他符号（symbol）一样，以大写字母开头的常量在包外可见。
)

//数值运算：+ - * / %
//比较运算：>、<、==、>=、<=和!= ，两个不同类型的整型数不能直接比较
//位运算：<<　>> ^ & | 特殊：^x 取反

//浮点数比较
func IsEqual(f1, f2, p float64) bool {
	return math.Dim(f1, f2) < p
}

//复数类型
var cplex complex64

func main() {
	fmt.Println("Hello world!")
	//匿名变量
	_, _, nickName := GetName() // := 不需声明nickName
	fmt.Println(nickName)       //output chibi Maruko

	//类型
	//布尔类型
	var v_bool bool
	//v_bool = true
	v_b := (1 == 2)
	//整型 - int8 uint8 int16 uint16 int32 uint32 int64 uint64 int uint uintptr-长度同指针
	//int和int32 在Go语言里被认为是两种不同的类型，编译器不会自动做类型转换，需要自己强转
	//比较严格
	var v_int1 int32
	v_int2 := 64 // 自动推导为int型
	//v_int1 = v_int2 // 会报错  use v_int2 (type int) as type int32 in assignment
	v_int1 = int32(v_int2)

	i, j := 1, 2
	if i == j {
		fmt.Println("i and j are equal.")
	}
	f1, f2 := 1.0003, 1.00003
	v_bool = IsEqual(f1, f2, 0.000001)
	fmt.Println(v_bool, v_b, v_int1)

	//复数
	cplex = 3.2 + 12i
	fmt.Println(real(cplex)) //获取实部
	fmt.Println(imag(cplex)) //获取虚部

}
