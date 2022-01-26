package main

import "fmt"

func testArray(){
	var arrayInt64 [3]int64
	arrayInt64[0],arrayInt64[1],arrayInt64[2] = 0,1,2
	fmt.Println(arrayInt64)

	arrayString := []string{"hello","go"}
	fmt.Println(arrayString)

	arrayFloat := [...]float64{1,2,3}
	fmt.Println(arrayFloat)

	//二维数组
	matrix := [2][2]int64 {{2,2},{2,2}}
	fmt.Println(matrix)
}

func testSlice(){
	// int型的slice
	ints := make([]int,3)
	ints[0],ints[1],ints[2] = 0,1,2 //没赋值默认为0值
	fmt.Println(ints)
	
	// string类型的slice
	names := []string{"zhang","wang","li","zhao"}
	fmt.Println(names, len(names),cap(names))
	
	// 拷贝names 切片
	names1 := names[0:3]
	fmt.Println(names1)

        // 修改names1的第一个元素   names与names1共享底层数组
	names1[0] = "laozhang"
	fmt.Println(names) // [laozhang wang li zhao] 

	// 遍历names slice 
	for _, name := range names{
		fmt.Println(name)
	}

	value1 := make([]int,0)
	for i:=0;i<3;i++{
		value1 = append(value1, i)
	}
	fmt.Println(value1)

	value2 := []int{3,4,5}
	value3 := append(value1, value2...) // ...解包 连接两个slice
	fmt.Println(value3)


	


}
func main(){
	// testArray()
	testSlice()
}
