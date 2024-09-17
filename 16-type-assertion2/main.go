package main

import (
	"fmt"
)

func main() {
	var i interface{} = 42

	switch v := i.(type) {
	case int:
		fmt.Printf("\nvariable type is %T", v)
	case float32:
		fmt.Printf("\nvariable type is %T", v)
	case float64:
		fmt.Printf("\nvariable type is %T", v)
	case string:
		fmt.Printf("\nvariable type is %T", v)
	}

}
