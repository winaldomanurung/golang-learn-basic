package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	var s1 student

	s1.name = "John"
	s1.grade = 2

	fmt.Println("name  :", s1.name)
	fmt.Println("grade :", s1.grade)

	var s2 = student{"ethan", 2}

	var s3 = student{name: "jason"}

	fmt.Println("student 2 name :", s2.name)
	fmt.Println("student 2 grade :", s2.grade)
	fmt.Println("student 3 name :", s3.name)
	fmt.Println("student 3 grade:", s3.grade)

	// object made with type struct can be referenced with pointer and saved in a variable with type struct pointer
	var s4 *student = &s1
	fmt.Println("student 1, name :", s1.name) // John
	fmt.Println("student 4, name :", s4.name) // John

	// changin s4 (pointer reference to s1) will change s1
	s4.name = "Ethan"
	fmt.Println("student 1, name :", s1.name) // Ethan
	fmt.Println("student 4, name :", s4.name) // Ethan
}
