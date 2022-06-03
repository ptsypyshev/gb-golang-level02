package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"strings"
	"testing"
	"time"
	"unicode"
)

const (
	firstLowerCharCode   = 97
	alphabetLenght       = 26
	shiftToUpperCharCode = 32
)

var snakeCases = map[string]string{
	"LoremIpsumDolor":             "lorem_ipsum_dolor",
	"SitAmet":                     "sit_amet",
	"ConsecteturAdipisicingElit":  "consectetur_adipisicing_elit",
	"CupiditateDucimusExcepturi":  "cupiditate_ducimus_excepturi",
	"ImpeditMaximeNecessitatibus": "impedit_maxime_necessitatibus",
	"PorroVoluptatibus":           "porro_voluptatibus",
	"AmetAnimiQui":                "amet_animi_qui",
	"RationeSaepeVeniamVoluptas":  "ratione_saepe_veniam_voluptas",
	"ConsequaturCorporisDelectus": "consequatur_corporis_delectus",
	"EsseIpsamA":                  "esse_ipsam_a",
	"aspernatur":                  "aspernatur",
	"DIGNISSIMOS":                 "dignissimos",
	"___":                         "___",
	" ":                           " ",
	"":                            "",
}

type testCase struct {
	result error
	input  *Student
	values map[string]interface{}
}

var testCases = []testCase{
	{
		result: nil,
		input:  new(Student),
		values: map[string]interface{}{
			"first_name": "Ivan",
			"last_name":  "Ivanov",
			"age":        "21",
			"is_male":    "true",
			"grade":      "4.9",
		},
	},
	{
		result: fmt.Errorf("cannot parse field value to float64"),
		input:  new(Student),
		values: map[string]interface{}{
			"first_name": "Ivan",
			"last_name":  "Ivanov",
			"age":        "21",
			"is_male":    "true",
			"grade":      "bool",
		},
	},
	{
		result: fmt.Errorf("cannot parse field value to int"),
		input:  new(Student),
		values: map[string]interface{}{
			"first_name": "Ivan",
			"last_name":  "Ivanov",
			"age":        "seven",
			"is_male":    "true",
			"grade":      "4.9",
		},
	},
	{
		result: fmt.Errorf("cannot parse field value to bool"),
		input:  new(Student),
		values: map[string]interface{}{
			"first_name": "Ivan",
			"last_name":  "Ivanov",
			"age":        "21",
			"is_male":    "NaN",
			"grade":      "4.9",
		},
	},
}

func generateRandomCase() (string, string) {
	wordsCount := rand.Intn(10) + 1
	maxWordLenght := rand.Intn(10) + 1
	in := make([]string, 0, wordsCount*maxWordLenght)
	out := make([]string, 0, wordsCount*maxWordLenght)
	var prevCharCode rune

	for i := 0; i < wordsCount; i++ {
		wordLenght := rand.Intn(maxWordLenght)
		for j := 0; j < wordLenght; j++ {
			charCode := rune(firstLowerCharCode + rand.Intn(alphabetLenght))
			if j == 0 {
				charCode -= shiftToUpperCharCode
			}
			in = append(in, string(charCode))
			if unicode.IsUpper(charCode) && unicode.IsLower(prevCharCode) && i != 0 {
				out = append(out, "_", strings.ToLower(string(charCode)))
			} else {
				out = append(out, strings.ToLower(string(charCode)))
			}
			prevCharCode = charCode
		}
	}

	return strings.Join(in, ""), strings.Join(out, "")
}

func generateRandomCases(length int) ([]string, []string) {
	rand.Seed(time.Now().UnixNano())
	PascalCase := make([]string, 0, length)
	SnakeCase := make([]string, 0, length)

	for i := 0; i < length; i++ {
		p, s := generateRandomCase()
		PascalCase = append(PascalCase, p)
		SnakeCase = append(SnakeCase, s)
	}

	return PascalCase, SnakeCase
}

func TestParseToStudent(t *testing.T) {
	for _, elem := range testCases {
		assert.Equal(t, elem.result, ParseToStudent(elem.input, elem.values))
	}
}

func TestToSnakeCase(t *testing.T) {
	for key, value := range snakeCases {
		assert.Equal(t, value, ToSnakeCase(key))
	}
}

func TestToSnakeCase2(t *testing.T) {
	input, result := generateRandomCases(30)
	for i := 0; i < len(input); i++ {
		assert.Equal(t, result[i], ToSnakeCase(input[i]))
	}
}

func ExampleToSnakeCase() {
	in := "FirstName"
	out := ToSnakeCase(in)
	fmt.Println(out)

	//Output: first_name
}

func ExampleParseToStudent() {
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

	//Output: &{FirstName:Ivan LastName:Ivanov Age:21 IsMale:true Grade:4.9}
}
