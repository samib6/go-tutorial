// Test Performance with preallocation of slice vs without

// Test - Output
// Total time without preallocation is 12.800551ms
// Total time with preallocation is 3.255676ms
package main

import (
	"fmt"
	"time"
)

func main() {
	var n int = 1000000
	var slice = []int{}            // slice initalized without a capacity
	var slice2 = make([]int, 0, n) // capacity initialized beforehand

	fmt.Printf("Total time without preallocation is %v", timeloop(slice, n))
	fmt.Printf("\nTotal time with preallocation is %v", timeloop(slice2, n))
}

func timeloop(slice []int, n int) time.Duration {
	var startTime = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(startTime)
}
