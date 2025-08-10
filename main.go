package main

import (
	"fmt"

	"github.com/Jack4Code/go-internal-tools/throw_away"
)

func main() {
	err := throw_away.FuncThatMaybeReturnsErr()
	if err != nil {
		fmt.Println("An error occurred:", err)
	}
}
