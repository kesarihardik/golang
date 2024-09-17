package main

import (
	"fmt"
)

// Struct is a collection of fields
type GasEngine struct { //type name struct
	KMPL   uint8
	Litres uint8
}

// method
func (e *GasEngine) totalRange() uint32 { //it can receive both struct and pointer to struct
	return uint32(e.KMPL) * uint32(e.Litres)
}

func main() {
	//struct literal
	var eng GasEngine = GasEngine{65, 24} // var eng2 gasEngine = gasEngine{kmpl: 65, litres: 24}

	//go automatically derefences struct fields
	fmt.Println(eng.KMPL, eng.Litres) //eng.KMPL = *(&eng.KMPL)

	fmt.Printf("Range of bike is %v km.", eng.totalRange())
}
