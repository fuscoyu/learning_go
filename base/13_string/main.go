package main

import "fmt"

func main(){
  s := "中国欢迎你"
  rs := []rune(s)
  sl := []byte(s)

  for i,v := range rs{
    var uft8Bytes []byte
    for j:=i*3; j<(i+1)*3; j++{
      uft8Bytes = append(uft8Bytes, sl[j])
    }
    fmt.Printf("%s => %X => %X\n", string(v), v, uft8Bytes)
  }


}
