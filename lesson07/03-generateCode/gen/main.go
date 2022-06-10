// go generate ./...
package main

import (
	"embed"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type TplData struct {
	Package  string
	TypeName string
	Fields   []TplField
}

type TplField struct {
	Name     string
	JsonName string
	Type     string
}

//go:embed gen.tpl
var f embed.FS

func main() {
	tpl, _ := f.ReadFile("gen.tpl")

	templateData := TplData{
		Package:  "student",
		TypeName: "Student",
		Fields: []TplField{
			{
				Name:     "FirstName",
				JsonName: "first_name",
				Type:     "string",
			},
			{
				Name:     "LastName",
				JsonName: "last_name",
				Type:     "string",
			},
			{
				Name:     "Age",
				JsonName: "age",
				Type:     "int",
			},
			{
				Name:     "IsMale",
				JsonName: "is_male",
				Type:     "bool",
			},
			{
				Name:     "Grade",
				JsonName: "grade",
				Type:     "float64",
			},
		},
	}
	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	genFile, err := os.Create(filepath.Join(curDir, "student", "student.go"))
	if err != nil {
		log.Fatal(err)
	}

	t := template.Must(template.New("student").Parse(string(tpl)))
	if err := t.Execute(genFile, templateData); err != nil {
		log.Fatal(err)
	}
}
