package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return nil
}
