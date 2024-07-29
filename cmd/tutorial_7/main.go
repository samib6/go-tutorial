package main

import (
	"fmt"
)

// Defining Struct type
type stud struct {
	id   uint16
	name string
}

func main() {

	// Initiallizing Structs
	//var reema stud = stud{id: 101, name: "Reema"} // Way 1
	var reema stud = stud{101, "Reema"} // Way 2

	// Way 3
	reema.id = 102
	fmt.Println(reema.id, "Name is :", reema.name)

}
