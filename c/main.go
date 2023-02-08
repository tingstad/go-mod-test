package main

import "fmt"
import "github.com/tingstad/go-mod-test/c/pkg"

func main() {
	fmt.Printf("Hello %s\n", pkg.Version())
}
