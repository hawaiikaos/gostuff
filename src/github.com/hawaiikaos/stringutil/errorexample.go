package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	y := 6
	i := 0
	z := 1.0
	if x < 0 {
		//return x, e
		return 0, ErrNegativeSqrt(x)
	} else {
		for i < y {
			z = z - (((z*z) - x) / (2*z))
			i++
		}
	//return z
		return z, nil
	}
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

/*func run() error {
	return ErrNegativeSqrt
}*/

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}