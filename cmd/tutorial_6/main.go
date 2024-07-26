package main

import (
	"fmt"
	"strings"
)

func main() {
	// Strings stored in binary format.
	// characters that fit or take 7 bits - ASCII letters start from
	// a - 97 till 127.
	// Some characters like chinese or emoji representation takes up more
	// than just a single byte.

	// UTF-32 helps to represent same but in here space gets wasted.
	// UTF-8  variable length -determine how much bytes are required
	// and stores accordingly.

	// WHEN DEALING WITH STRINGS IN GO YOU ARE DEALING
	// WITH UNDERLYING REPRESENTATION OF BYTES.
	// Therefore len(str) in go is length of bytes and not length of characters :)

	var mystr = "résumé"

	var indx = mystr[1] // this indexes underlying byte array

	fmt.Printf("%v, %T\n", indx, indx)

	for i, v := range mystr {
		fmt.Println(i, v)
	}

	// The value of é is 233
	// value of r is 114
	// value of a  is 97
	// Output -
	/*
		195, uint8  -- Note the value at index 1 is 195
		0 114
		1 233		-- Note it skipped index 2 because letter é occupies 2 bytes
		3 115
		4 117
		5 109
		6 233
	*/

	fmt.Println("RUNES ........")
	// To get array of unicode numbers for each characters
	// RUNES
	var mystr1 = []rune("résumé") // rune is just alias of int32

	var indx1 = mystr1[1] // this indexes underlying byte array

	fmt.Printf("%v, %T\n", indx1, indx1)

	for i, v := range mystr1 {
		fmt.Println(i, v)
	}
	/*
		RUNES ........
		233, int32
		0 114
		1 233
		2 115
		3 117
		4 109
		5 233
	*/

	// Create strings
	// Dont use concat operation inefficient
	// Use String Builders instead

	var strSlice = []string{"t", "u", "t", "o", "r", "i", "a", "l"}
	var strBuilder strings.Builder // INternally created an array and appends each character
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catString = strBuilder.String() // only in this step is a new string created from appended values
	fmt.Printf("\n%v", catString)

}
