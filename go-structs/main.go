package main

import (
	"fmt"
	"reflect"
)

// struct declaration - basis for golang's pseudo-OOP
// structs are also VALUE DATA TYPES
type Person struct {
	// can hold any data type
	name         string
	age          byte
	profession   string
	heightMetres float32
}

type Animal struct {
	// note the naming convention, these have module scope
	Name   string `required_max:"100"`
	Origin string
}

type Bird struct {
	// embedding = composition (not inheritance)
	Animal
	SpeedKPH uint8
	CanFly   bool
}

func main() {
	// struct definition
	firstClient := Person{
		name:         "Gicheru",
		age:          24,
		profession:   "Electronics Engineer",
		heightMetres: 1.77,
	}

	// anonymous struct definition
	anonyStruct := struct {
		// embedded struct
		Person
		school string
		GPA    float32
	}{
		Person: Person{
			name:       "Nyokabi",
			age:        24,
			profession: "Embedded Systems Engineer",
			// heightMetres not given, will default to 0
		},
		school: "University of Nairobi",
		GPA:    3.735,
	}

	fmt.Println(firstClient.profession)
	fmt.Println(anonyStruct.profession, ",", anonyStruct.school)

	// using struct tags
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}
