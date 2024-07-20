package main

import (
	"errors"
	"fmt"
)

func main() {

	var num1 int = 11
	var num2 int = 0
	var num, rem, err = intDivision(num1, num2)
	if err != nil {
		fmt.Printf(err.Error())
	} else if rem == 0 {
		fmt.Printf("The result of the integer division is %v", num)
	} else {
		fmt.Printf("The result of the integer division is %v and remainder is %v ", num, rem)
	}

	// Switch with  case statements
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case rem == 0:
		fmt.Printf("The result of the integer division is %v", num)
	default:
		fmt.Printf("The result of the integer division is %v and remainder is %v ", num, rem)
	}

	// Switch with conditional
	switch rem {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("Thie division was close")
	default:
		fmt.Println("The division was not close")

	}

}

// parameter type as well return type is to be speicified.
// Func declaration way 2 - func name(var1, var2 ,var3 int) -- if all have one type
func intDivision(num1 int, num2 int) (int, int, error) {
	//Yes theres a type 'ERROR' that you can define!
	var error error

	if num2 == 0 {
		error = errors.New("Cannot divide by 0")
		return 0, 0, error
	}
	var num int = num1 / num2
	var rem int = num1 % num2
	return num, rem, error
}
