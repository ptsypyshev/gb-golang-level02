# ДЗ 07

Все решения лежат в отдельных папках с номером задания.

## 1. Написать функцию, которая принимает на вход структуру in (struct или кастомную struct) и values map[string]interface{} (key - название поля структуры, которому нужно присвоить value этой мапы)
Необходимо по значениям из мапы изменить входящую структуру in с помощью
пакета reflect. Функция может возвращать только ошибку error. 
Написать к данной функции тесты (чем больше, тем лучше - зачтется в плюс).  

Объявляем структуру, в которую будем писать:  
```golang
type Student struct {
    FirstName string
    LastName  string
    Age       int
    IsMale    bool
    Grade     float64
}
```

Далее предполагаем вариант, что ключи в мапе записаны в стиле snake_case. 
И объявляем функцию, которая будет преобразовывать PascalCase (camelCase) в snake_case:  
```golang
var matchCamelCase = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
snake := matchCamelCase.ReplaceAllString(str, "${1}_${2}")
return strings.ToLower(snake)
}
}
```
Для парсинга исходной строки используем регулярные выражения.  

Далее, объявляем функцию `ParseToStudent`, в которой с помощью рефлексии обращаемся к исходной структуре
и проходим по её полям. Сопоставляем ключ мапы с полем структуры и пытаемся привести значение к нужному типу.
Если привидение успешно, то пишем значение в поле, если нет - возвращаем ошибку.

```golang
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
```

Дополнительно к этой задаче написаны и успешно проходят тесты.
```shell
$ go test -v .
=== RUN   TestParseToStudent
--- PASS: TestParseToStudent (0.00s)
=== RUN   TestToSnakeCase
--- PASS: TestToSnakeCase (0.00s)
=== RUN   TestToSnakeCase2
--- PASS: TestToSnakeCase2 (0.00s)
PASS
ok      github.com/ptsypyshev/gb-golang-level02/lesson07/01-parseMapToStruct    0.331s
```

## 2. Написать функцию, которая принимает на вход имя файла и название функции.
Необходимо подсчитать в этой функции количество вызовов асинхронных функций.
Результат работы должен возвращать количество вызовов int и ошибку error.
Разрешается использовать только `go/parser`, `go/ast` и `go/token`.  

Объявляем функцию `countGoroutinesInFunc(filePath string, functionName string) (runCounter int)`, 
которая принимает на вход путь до файла и имя искомой функции, а возвращает количество вызовов горутин.

```golang
func countGoroutinesInFunc(filePath string, functionName string) (runCounter int) {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, filePath, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("cannot parse file %s", filePath)
	}

	ast.Inspect(file, func(n ast.Node) bool {
		funcDeclaration, ok := n.(*ast.FuncDecl)
		if ok && funcDeclaration.Name.String() == functionName {
			runCounter = goFinder(funcDeclaration, runCounter, WithoutFor)
			runCounter = goInForFinder(funcDeclaration, runCounter)
		}
		return true
	})
	return
}
```

Пытаемся распарсить файл в AST и инспектируем его на предмет поиска функций.
Если имя функции совпадает с параметром `functionName`, то запускаем функции, 
считающие количество вхождений строк с запуском горутин в данной функции, а также в циклах.

```golang
func goFinder(currentNode ast.Node, counter int, forIterCounter int) int {
	ast.Inspect(currentNode, func(n ast.Node) bool {
		if _, ok := n.(*ast.GoStmt); ok {
			counter = goCounter(counter, forIterCounter)
		}
		return true
	})
	return counter
}
```

Функция возвращает количество найденных в элементе AST горутин.
Т.к. аналогичный кусок кода используется для поиска внутри циклов, 
то вынес это в отдельную функцию с доп. параметром `forIterCounter`, меняющим 
логику подсчета количества вызовов горутин в цикле (функция `goCounter`).

```golang
func goCounter(counter int, forIterCounter int) int {
	if forIterCounter == 0 {
		counter++
	} else {
		counter--
		counter += forIterCounter
	}
	return counter
}
```

Аналогично `goFinder` ищем циклы в выбранной функции, получаем количество итераций цикла,
а также количество вызовов горутин внутри цикла.

```golang
func goInForFinder(currentNode ast.Node, counter int) int {
	ast.Inspect(currentNode, func(n ast.Node) bool {
		forStatement, ok := n.(*ast.ForStmt)
		if ok {
			iterCountField := forStatement.Cond.(*ast.BinaryExpr).Y
			iterCount, err := strconv.Atoi(types.ExprString(iterCountField)) // I haven't find another way to get count of iterations in for loop
			if err != nil {
				fmt.Printf("cannot convert 'for condition expression operand' to int: %v", err)
				iterCount = 0
			}
			counter = goFinder(forStatement, counter, iterCount)
		}
		return true
	})
	return counter
}
```

Я не нашел в документации на пакет `go/ast`, как вывести количество итераций цикла.
Поэтому пришлось "доставать" операнд из условия for и приводить его к int.


## 3. (*не обязательное*). Написать кодогенератор под какую-нибудь задачу.

За основу берем задачу из п. 1 по конвертации данных из мапы в структуру.
Кодогенератор будет создавать пакет `student`, в которым из мапы будет создана структура с полями 
и метод FromMap. За основу файла пакета возьмем go-шаблон:
```gotemplate
package {{.Package}}

type Student struct {
    {{- range .Fields}}
    {{.Name}} {{.Type}}
    {{- end}}
}

func (s *{{.TypeName}}) FromMap(m map[string]interface{}) {
    {{- range .Fields}}
    if v, ok := m["{{.JsonName}}"].({{.Type}}); ok {
        s.{{.Name}} = v
    }
    {{- end}}
}
```

Далее в файле генератора считываем этот шаблон: 
```golang
tpl, _ := f.ReadFile("gen.tpl")
```

И структуру которая задает, поля и типы данных, а также сопоставление полей и ключей мапы.
После чего создаем новый файл и пишем в него сгенерированный из шаблона результат:
```golang
curDir, err := os.Getwd()
genFile, err := os.Create(filepath.Join(curDir, "student", "student.go"))
t := template.Must(template.New("student").Parse(string(tpl)))
if err := t.Execute(genFile, templateData); err != nil {
    log.Fatal(err)
}
```

Запускаем генерацию файла с командой:  
```shell
$ go generate ./...
```
