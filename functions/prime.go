package functions

import (
	"math"
)

func IsPrime(n int) bool {
	sqrtN := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true

}
