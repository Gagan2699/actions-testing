package main

import (
	"fmt"
	"runtime"
)

func main() {
        fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("GOOS: %s\n", runtime.GOOS)
}
// func Demo() {
// 	fmt.Printf("Go version: %s\n", runtime.Version())
// 	fmt.Printf("GOOS: %s\n", runtime.GOOS)

// 	fmt.Println(quote.Go())
// }
