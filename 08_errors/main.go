// package main

// import (
// 	"errors"
// 	"fmt"
// 	"math"
// )

// func Sqrt(value float64) (float64, error) {
// 	if value < 0 {
// 		return 0, errors.New("value is below 0")
// 	}
// 	return math.Sqrt(value), nil
// }

// func main() {
// 	result, err := Sqrt(-1)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result)
// 	}

// 	result, err = Sqrt(4)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result)
// 	}
// }

// get func that count sqrt and respond with an error if needed

package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("value is below zero")
	}
	return math.Sqrt(value), nil
}

func main() {
	result, err := Sqrt(-1)
	fmt.Println(result)
	fmt.Println(err)

	result, err = Sqrt(5)
	fmt.Println(result)
	fmt.Println(err)
}
