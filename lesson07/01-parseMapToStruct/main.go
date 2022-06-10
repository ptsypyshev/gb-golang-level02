package main

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Student structure represents a real student person
type Student struct {
	FirstName string
	LastName  string
	Age       int
	IsMale    bool
	Grade     float64
}

var matchCamelCase = regexp.MustCompile("([a-z0-9])([A-Z])")

// ToSnakeCase function converts PascalCase (camelCase) string to snake_case style.
func ToSnakeCase(str string) string {
	snake := matchCamelCase.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(snake)
}

// ParseToStudent function converts a map collection to a Student structure
func ParseToStudent(in *Student, values map[string]interface{}) error {
	reflectValue := reflect.Indirect(reflect.ValueOf(in))
	for i := 0; i < reflectValue.Type().NumField(); i++ {
		field := reflectValue.Type().Field(i)
		fieldName := ToSnakeCase(field.Name)
		fieldValue := values[fieldName]
		reflectFieldName := reflectValue.FieldByName(field.Name)

		if fieldValueString, ok := fieldValue.(string); ok &&
			reflectFieldName.Type().AssignableTo(reflect.TypeOf(float64(0))) {
			floatVal, err := strconv.ParseFloat(fieldValueString, 64)
			if err != nil {
				return fmt.Errorf("cannot parse field value to float64")
			}
			reflectFieldName.Set(reflect.ValueOf(floatVal))
		} else if fieldValueString, ok := fieldValue.(string); ok &&
			reflectFieldName.Type().AssignableTo(reflect.TypeOf(0)) {
			intVal, err := strconv.Atoi(fieldValueString)
			if err != nil {
				return fmt.Errorf("cannot parse field value to int")
			}
			reflectFieldName.Set(reflect.ValueOf(intVal))
		} else if fieldValueString, ok := fieldValue.(string); ok &&
			reflectFieldName.Type().AssignableTo(reflect.TypeOf(true)) {
			boolVal, err := strconv.ParseBool(fieldValueString)
			if err != nil {
				return fmt.Errorf("cannot parse field value to bool")
			}
			reflectFieldName.Set(reflect.ValueOf(boolVal))
		} else if _, ok := fieldValue.(string); ok {
			reflectFieldName.Set(reflect.ValueOf(fieldValue))
		} else {
			return fmt.Errorf("cannot parse values")
		}
	}
	return nil
}

func main() {
	newValues := map[string]interface{}{
		"first_name": "Ivan",
		"last_name":  "Ivanov",
		"age":        "21",
		"is_male":    "true",
		"grade":      "4.9",
		"additional": "test",
	}

	student := new(Student)
	if err := ParseToStudent(student, newValues); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", student)
}
