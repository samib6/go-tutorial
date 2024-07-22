package main

import "fmt"

func main() {
	// Arrays - fixed length, same data type, stored in contigous
	// memory and indexable.

	//int16 -  2bytes
	//int32 - 4 bytes
	//it will look for 12 bytes of contigous memory
	var intArr [3]int32
	intArr[1] = 90
	fmt.Println(intArr[0])
	fmt.Println(intArr)

	// Print address location
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])

	// Intialize array directly
	//Way 1
	// var intArr1 [3]int16 = [3]int16{1, 2, 3}
	// fmt.Println(intArr1)

	// Way 2
	// intArr1 := [3]int16{1, 2, 3}
	// fmt.Println(intArr1)

	// Way 3 - dynamically calculate size of array
	intArr1 := [...]int16{1, 2, 3}
	fmt.Println(intArr1)

	// Slicing - Array
	// When we eliminate specifying the size that means slicing
	var arr1 []int16 = []int16{1, 2, 3}
	fmt.Printf("The lenght is %v and capacity is %v", len(arr1), cap(arr1))
	arr1 = append(arr1, 7)
	fmt.Printf("\nThe lenght is %v and capacity is %v", len(arr1), cap(arr1))
	// Note - length is actual number of elements in array
	// Cap is capacity of array - intially capacity was 3
	// when append happened  then go checks if in current array memory
	// is sufficient to accomodate more ele if not then
	// searches for another location with double(not sure of incrementing factor)
	// size of current array
	// and copies all elements from original array to this new location

}
