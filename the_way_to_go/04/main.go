package main

import "fmt"


func main(){
  // var i1 = 5
  // fmt.Printf("An interger: %d, it's location in memory: %p\n", i1, &i1)
  //
  // var intP *int
  // intP = &i1 
  // fmt.Printf("intP %d, %p & %p\n", *intP, intP, &intP)


  // s := "Good bye"
  // var p *string = &s
  // *p = "ciao"
  // fmt.Printf("Here is the pointer p: %p\n", p)
  // fmt.Printf("Here is the string p: %s\n", *p)
  // fmt.Printf("Here is the string s: %s\n", s)


  const i = 5
  ptr := &i 
  // 常数、映射索引表达式和函数返回的值是不可寻址的，因此不能对他们使用 & 运算符。
  // 如果常数可以寻址的话，我们就可以通过指针修改常数的值，这无疑是破坏了常数的定义。
  // map 如果map中元素不存在，则返回零值。而零值是不可变对象，不能寻址。
  // 如果map中元素存在，go中map实现中元素地址是变化的，这意味着寻址是无意义的。 
  fmt.Printf("%p\n", ptr)


}
