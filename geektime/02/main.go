// hello golang
// package main
// import "fmt"
//
// func main(){
//   fmt.Println("hello golang!")
// }

// command line

package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)

func init() {
	// flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine = flag.NewFlagSet("", flag.PanicOnError) // input argument error 
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	cmdLine.StringVar(&name, "name", "everyone", "The greeting object.")

}

func main() {
	// flag.Usage = func(){
	//   fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//   flag.PrintDefaults()
	// }
  cmdLine.Parse(os.Args[1:])
	fmt.Printf("hello, %s!\n", name)
}
