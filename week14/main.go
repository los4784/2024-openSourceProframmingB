package main

import "fmt"

func main() {
	var student struct {
		id   int
		name string
		gpa  float32
	}

	student.id = 20241234
	student.name = "Son Heungmin"
	student.gpa = 4.5
	fmt.Println(student.gpa)
	var student2 struct {
		id   int
		name string
		gpa  float32
	}
	student.id = 20244321
	student.name = "Seo Jihye"
	student.gpa = 4.43
	fmt.Println(student2.id)
}
