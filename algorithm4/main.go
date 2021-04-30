package main

import (
	"fmt"
	"math"
)

func main() {
	var lamps []int
	total := 100
	root := int(math.Sqrt(float64(total)))
	for i := 1; i < root+1; i++ {
		lamps = append(lamps, i*i)
	}
	fmt.Println(lamps)
}
