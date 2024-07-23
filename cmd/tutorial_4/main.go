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

	// To insert multiple values we can use spread operator ...
	var slice2 []int16 = []int16{1, 2}
	slice2 = append(arr1, slice2...)
	fmt.Println(slice2)

	var slice3 []int16 = make([]int16, 3, 8) // param2 is length , param 3 is capacity
	fmt.Println(slice3)

	//Maps

	//Declaration - way 1
	var map1 map[string]uint8 = make(map[string]uint8)
	fmt.Println(map1)

	var map2 map[string]uint8 = map[string]uint8{"Adam": 8, "Roza": 10}
	fmt.Println(map2["Roza"])

	//Interesting - By default even if key doesnt exists
	// Go will return default value for the data type eg - 0 for ints
	// 2 values are returned from map[key] , first is value, second is boolean
	// which denotes if key exists or not
	delete(map2, "Adam")
	var age, doesKeyExists = map2["Adam"]
	fmt.Println(age, doesKeyExists)
	if doesKeyExists {
		fmt.Println(age, doesKeyExists)
	} else {
		fmt.Println("Key does not exists")
	}

	// Iterate - Maps Loops
	for name, age := range map2 {
		fmt.Printf("\n Name: %v and Age is %v", name, age)
	}

	// Iterate - Arrays
	for index, value := range slice2 {
		fmt.Printf("\n Index is %v and value is %v", index, value)
	}

	for i := 0; i < 10; i++ { // shorthands op - i*=1 , i/=10 etc
		fmt.Println(i)
	}

}
