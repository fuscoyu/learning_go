package main

import (
	"fmt"
	"strings"
)

func main() {
  // str := "VM-2237: add gpu vendor_id and product_id"
  sub_str := "VM-2139_VM-2237"

  if strings.Contains(sub_str, "VM-2237") {
    fmt.Println("yes")
  }
}
