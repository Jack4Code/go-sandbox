package main

import (
	"errors"
	"fmt"

	"github.com/Jack4Code/go-internal-tools/throw_away"
)

func main() {
	err := throw_away.FuncThatMaybeReturnsErr()
	switch {
	case err == nil:
		fmt.Println("All good, no error")
	case errors.Is(err, throw_away.ErrRIs2):
		fmt.Println("ErrRIs2 occurred, static kind")
	case errors.Is(err, throw_away.ErrOtherKind):
		fmt.Println("ErrOtherKind occurred, dynamic kind")
	default:
		fmt.Println("An unexpected error occurred:", err)
	}
}
