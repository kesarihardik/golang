package main

import (
	"fmt"
)

// type error interface {
//     Error() string
// }

type ErrNegativeSqrt float64 //type alias

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

type InvalidAgeErr struct {
	Age int
}

func (x *InvalidAgeErr) Error() string {
	return fmt.Sprintf("You are not of valid age - %d", x.Age)
}

func validate(age int) (string, error) {
	if age >= 18 {
		return fmt.Sprintf("You age - %d is valid.", age), nil
	}
	return "", &InvalidAgeErr{Age: age}
}

func main() {
	fmt.Println(Square(2))
	fmt.Println(Square(-2))

	age := 16
	s, err := validate(age)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
