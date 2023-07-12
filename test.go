package actions

// import (
// 	"fmt"
// 	"runtime"

// 	"rsc.io/quote"
// )

func Demo() {
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("GOOS: %s\n", runtime.GOOS)

	fmt.Println(quote.Go())
}
