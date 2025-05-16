package main

import (
	"fmt"
	// "math"
	"strconv"
)

// var name string

// var name string = "xiaoming"
// var name2 = "xiaoming"




func testStringdemo() {
	Labels := map[string]string{"node": ""}
	migrationTargetNodeName, ok := Labels["node"]

	fmt.Println(migrationTargetNodeName, ok)

}

func testVar() {
	// fmt.Println("hello luyu")
	var str1 string
	str1 = "hello"
	fmt.Println(str1)

	var str2 string = "hello go1"
	fmt.Println(str2)

	var str3, str4 string = "hello", "go2"
	fmt.Println(str3, str4)

	str5 := "hello go3"
	fmt.Println(str5)

	var i int64
	var s string
	fmt.Println("i is ", i)
	fmt.Println("s is ", s)

	ii := 88
	ss := "hello go"
	fmt.Println("ii is ", ii)
	fmt.Println("ss is ", ss)
}

func testBool() {
	var b1 bool //默认是false
	b2 := true
	b3 := false
	fmt.Println(b1, b2, b3)
	if b1 {
		fmt.Println("b1 is true")
	}
	if b2 {
		fmt.Println("b2 is true")
	}
	if b3 {
		fmt.Println("b3 is true")
	}
}

func testNumber() {
	var i int64 // 默认为0值
	var i2 int32
	fmt.Println(i, i2)
	i3 := 10
	i4 := int64(10)
	fmt.Println(int64(i3) + i4)

	// math.Float64frombits(i4)
	// fmt.Println(1.0 / 0)
}
func testString() {
	s1 := "\"你好 go\""
	s2 := `"你好go"`
	fmt.Println(s1)
	fmt.Println(s2)
	// fmt.Println(s1 + s2)
}

func testConvert() {
	// int -> string
	sint := strconv.Itoa(97)
	fmt.Println(sint, sint == "97")

	// byte -> string
	bint := byte(1)
	fmt.Println(bint)

	// int64 -> string
	sint64 := strconv.FormatInt(int64(97), 10)
	fmt.Println(sint64)

	// int64 -> string (hex) 十六进制 base 指的是进制
	sint64hex := strconv.FormatInt(int64(97), 16)
	fmt.Printf("%T, %v\n", sint64hex, sint64hex)

	// string -> int
	_int, _ := strconv.Atoi("97")
	fmt.Println(_int, _int == int(97))

	// string -> int64
	_int64, _ := strconv.ParseInt("97", 10, 64)
	fmt.Println(_int64, _int64 == int64(97))

	// string -> int32
	_int32, _ := strconv.ParseInt("97", 10, 32)
	fmt.Println(_int32, int32(_int32) == int32(97))

	// int -> int64
	// 不会丢失精度
	var n int = 97
	fmt.Println(int64(n) == int64(97))

	// string -> float32/float64
	f := "3.1415926"
	if s, err := strconv.ParseFloat(f, 32); err == nil {
		fmt.Println(s)
	}
	if s, err := strconv.ParseFloat(f, 64); err == nil {
		fmt.Println(s)
	}

}

const (
	Sunday = iota
	Monday
)

func main() {
	testVar()
	testBool()
	testNumber()
	testString()
	fmt.Println(Sunday, Monday)
	testConvert()

	// 为什么有int还需要int32和int64，业务代码里如何选择？
	// uint 和 int 这两种类型是不带大小的，他们会根据编译参数 GOARCH 平台来决定。
	// 1. 如何和第三方库和标准库有交互的  一般使用相同就好。避免类型转换，导致程序性能下降。
	// 2. 自己使用的根据数据大小选择。

	// int类型取值范围
	// fmt.Println("int8:", math.MinInt8, "~", math.MaxInt8)
	// fmt.Println("int16:", math.MinInt16, "~", math.MaxInt16)
	// fmt.Println("int32:", math.MinInt32, "~", math.MaxInt32)
	// fmt.Println("int64:", math.MinInt64, "~", math.MaxInt64)
	//   var i int64
	//   var b string
	//   fmt.Println("i is ", i) // i is 0
	//   fmt.Println("b is ", b) // b is ""

	//   var floatNum float64 = 1.0
	//   var price1, price2  float64 = 8.8, 9.6
	//   fmt.Println(floatNum, price1, price2) // 1.0, 8.8, 9.6

	//   ii := 1
	//   s := "Hello Go!"
	//   fmt.Println("ii is ", ii) // ii is 1
	//   fmt.Println("s is ", s) // s is Hello Go!
	testStringdemo()

}
