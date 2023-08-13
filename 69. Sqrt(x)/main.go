package main

// method one: use binary search
//func mySqrt(x int) int {
//	if x == 0 || x == 1 {
//		return x
//	}
//
//	l, r := 0, x
//	var res int
//
//	for l <= r {
//		mid := (r-l)/2 + l
//		if mid == x/mid {
//			return mid
//		} else if mid < x/mid {
//			l = mid + 1
//			res = mid
//		} else {
//			r = mid - 1
//		}
//	}
//
//	return res
//}

// method two: use Newton's method
func mySqrt(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r
}
