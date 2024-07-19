package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	////////////////////////
	// Integers
	var intvar int
	intvar = 3383010
	fmt.Println(intvar)

	var int8Num int8 = 127
	fmt.Println(int8Num)

	var int16Num int16
	fmt.Println(int16Num)

	var int32Num int32
	int32Num = 837
	fmt.Println(int32Num)

	var rgbValsRed uint16 = 255
	fmt.Println("RGB VALUE is", rgbValsRed)

	// Eg - When storing RGB values it be useful to declare them
	// as uint16 instead of int64.

	//////////////////////
	// Strings //////////
	a := "er"
	fmt.Println(a)

	str_b := "single line string, use double quotes"
	str_c := `
	multiline string showuld use back ticks 
	`
	str_d := "A" //ascii characters set
	fmt.Println(str_b, " .... ", str_c)
	fmt.Println(len(str_d)) // Length function returns number of bytes

	// Note - Go uses UTF-8 encoding hence characters outside vanilla
	// ASCII character set are stores with more than a single byte.
	// eg - Gamma character etc.

	// To Find length of string
	// Then import unicode/uttf-8 package
	fmt.Println(utf8.RuneCountInString("Γ"))

	//////////////////////
	// Floats
	// Floats doesnt has only float data type - only float32,float64
	var floatNum float32 = 34.0938
	fmt.Println(floatNum, 1, "printing int + string together")

	// Arithmetic operations
	// Note - Cant perform operations on mixed data types

	var i int32 = 78
	var j float32 = 3467.987
	var res float32 = j + float32(i)

	fmt.Println(res)

	// Boolean
	var boolVar = true // another way of declaring variables
	fmt.Println(boolVar)

}
