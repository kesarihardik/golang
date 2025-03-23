package main

import (
	"fmt"
)

// type error interface {
//     Error() string
// }

type ErrNegativeSqrt float64

// fmt package looks for error interface while printing values
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e) //formats as well as return string
}

func Square(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return x * x, nil
}

func main() {
	fmt.Println(Square(2))
	fmt.Println(Square(-2))
}
