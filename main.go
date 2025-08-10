package main

import (
	"fmt"
	"strings"

	"github.com/Jack4Code/go-internal-tools/throw_away"
)

func main() {
	err := throw_away.FuncThatMaybeReturnsErr()
	if err != nil {
		if strings.HasPrefix(err.Error(), "an error occurred because r is") {
			fmt.Println("Error that has dynamic message: ", err)
		} else {
			fmt.Println("Error that has static message: ", err)
		}
	}
}
