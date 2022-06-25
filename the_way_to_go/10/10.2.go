package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(per *Person) {
	per.firstName = strings.ToUpper(per.firstName)
	per.lastName = strings.ToUpper(per.lastName)
}

func main() {
	var per1 Person
	per1.firstName = "Zhang"
	per1.lastName = "San"
	upPerson(&per1)
	fmt.Println(per1)

	per2 := new(Person)
	per2.firstName = "Li"
	per2.lastName = "Si"
	upPerson(per2)
	fmt.Println(per2)

	per3 := &Person{"Wang", "Wu"}
	upPerson(per3)
	fmt.Println(per3)
}