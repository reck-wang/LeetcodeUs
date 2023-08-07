package main

// method one: use recursion
//func myPow(x float64, n int) float64 {
//	if n == 0 {
//		return 1
//	} else if n < 0 {
//		return 1 / myPow(x, -n)
//	} else if n&1 == 1 {
//		return x * myPow(x, n-1)
//	}
//
//	return myPow(x*x, n/2)
//}

// method two: use iterative
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	pow := 1.0
	for n > 0 {
		if n&1 == 1 {
			pow *= x
		}
		x *= x
		n >>= 1
	}

	return pow
}
