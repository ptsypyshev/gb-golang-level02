package main

import (
	"fmt"
	"github.com/ptsypyshev/gb-golang-level02/lesson07/03-generateCode/student"
)

//go:generate go run ./gen/main.go
//go:generate goimports -w ./student/student.go

func main() {
	s := &student.Student{}
	newValues := map[string]interface{}{
		"first_name": "Ivan",
		"last_name":  "Ivanov",
		"age":        21,
		"is_male":    true,
		"grade":      4.9,
	}
	s.FromMap(newValues)
	fmt.Println(s)
}
