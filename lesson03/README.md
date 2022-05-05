# ДЗ 03

Для тестирования версионирования пришлось создать отдельный репозиторий "github.com/ptsypyshev/golang_semver_learning", т.к. импорт работал только при изменениях в ветке master(main).

## п. 1 Создание нового проекта с помощью go mod (create new project)
Создаем папку и выполняем инициализацию модуля:  
`$ mkdir golang_semver_learning`  
`$ cd golang_semver_learning`  
`$ go mod init github.com/ptsypyshev/golang_semver_learning`

Создаем отдельный пакет, который будем использовать в составе модуля:  
`$ mkdir checker`  
`$ cd checker`  
Пишем собственный пакет в файле checker.go

## п. 2 Публикация проекта в репозитории (publish a project to repository)
Коммитим и пушим код в репозиторий, добавляем тег:  
`$ git commit -a -m 'Initial commit'`  
`$ git push`  
`$ git tag -a v0.0.1 -m 'Initial commit'`  
`$ git push origin --tags`

В отдельный проект импортируем пакет из нашего модуля:  
`$ cat main.go`  
`package main`  
`import "github.com/ptsypyshev/golang_semver_learning/checker"`  
`...`

Подгружаем зависимости:  
`$ go get github.com/ptsypyshev/golang_semver_learning/checker`  
`go: downloading github.com/ptsypyshev/golang_semver_learning v0.0.1`

Проверяем работоспособность:  
`$ go run main.go`


## п. 3 Обновление проекта в репозитории (update a project to repository)
Вносим изменения в функциональность пакета (github.com/ptsypyshev/golang_semver_learning/checker).
Коммитим изменения и добавляем новый тег:  
`$ git commit -a -m 'Add CheckBulk feature'`  
`$ git push`  
`$ git tag -a v1.0.1 -m 'Add CheckBulk feature'`  
`$ git push origin --tags`

Обновляем зависимости (для testmain):  
`$ go get github.com/ptsypyshev/golang_semver_learning/checker`  
`go get: upgraded github.com/ptsypyshev/golang_semver_learning v0.0.1 => v1.0.1`

Проверяем работоспособность:  
`$ go run main.go`  
Работает некорректно, т.к. при обновлении пакета не учитывалась обратная совместимость.
Функция CheckWebServer v0.0.1 выводила информацию через fmt.Println, а v1.0.1 возвращает значения и ошибку.

## п. 4 Вносим изменения и пушим мажорную версию (update a project and push new major version to repository)
### п. 4.1 Откат изменений с повышением версии тега (revert changes and increment tag version)
Создаем папку для новой верcии, копируем необходимые файлы, инициализируем новую версию:  
`$ mkdir v2`  
`$ cp * v2/`  
`go mod edit -module github.com/ptsypyshev/golang_semver_learning/v2 v2/go.mod`

Откатываем изменения, сломавшие совместимость:  
`$ git log --oneline`  
`1dcc8ae (HEAD -> master, tag: v1.0.1, origin/master) Add CheckBulk feature`  
`8ae2ad7 (tag: v0.0.1) Initial commit`  

`$ git revert HEAD`  
`[master 22a7a2f] Revert "Add CheckBulk feature"`  
`1 file changed, 38 insertions(+), 34 deletions(-)`  
`rewrite checker/checker.go (80%)`  

`$ git log --oneline`  
`22a7a2f (HEAD -> master) Revert "Add CheckBulk feature"`  
`1dcc8ae (tag: v1.0.1, origin/master) Add CheckBulk feature`    
`8ae2ad7 (tag: v0.0.1) Initial commit`  

Пушим изменения после реверта, инкрементируем версию тега, чтобы зависимые проекты использовали совместимую версию:  
`$ git push`  
`$ git tag -a v1.0.2 -m 'Revert CheckBulk feature'`    
`$ git push origin --tags`  

Обновляем зависимости (для testmain):  
`$ go get github.com/ptsypyshev/golang_semver_learning/checker`  
`go get: upgraded github.com/ptsypyshev/golang_semver_learning v1.0.1 => v1.0.2`

Проверяем работоспособность:  
`$ go run main.go`  
Работает корректно.

### п. 4.2 Публикуем новую версию проекта (publish new project version)
Возвращаемся к checker и обновляем версию:  
`$ git add .`  
`$ git commit -a -m 'Add CheckBulk feature (new version)'`  
`$ git push`  
`$ git tag -a v2.0.0 -m 'Add CheckBulk feature (new version)'`  
`$ git push origin --tags`  

Затем в проект testmain добавляем новую зависимость:  
`package main`  
`import (`  
`"fmt"`  
`"github.com/ptsypyshev/golang_semver_learning/checker"`  
`checkerV2 "github.com/ptsypyshev/golang_semver_learning/checker/v2"`  
`)`  
`...`

И получаем эту зависимость из репозитория:  
`$ go get github.com/ptsypyshev/golang_semver_learning/v2/checker`  
`go: downloading github.com/ptsypyshev/golang_semver_learning/v2 v2.0.0`  
`go get: added github.com/ptsypyshev/golang_semver_learning/v2 v2.0.0`  

Проверяем работоспособность  
`$ go run main.go`  

Работает корректно.

## п. 5 Очистка неиспользуемых библиотек (clean unused modules)
Очищаем с помощью tidy:  
`$ go mod tidy`  

Выводим список используемых модулей:  
`$ go list -m all`  
`testmain`  
`github.com/ptsypyshev/golang_semver_learning v1.0.2`  
`github.com/ptsypyshev/golang_semver_learning/v2 v2.0.0`  

Также из go.sum исчезли неиспользуемые версии:  
`$ cat go.sum`  
`github.com/ptsypyshev/golang_semver_learning v1.0.2 h1:Y9llVpCCeqUnl98TrjoR17u5404dHaMkcMs9eX5FraE=`  
`github.com/ptsypyshev/golang_semver_learning v1.0.2/go.mod h1:u+MzkJIvSlFTaKNA3zXfSkh4czu8Ctrb0jOCpS1TkjQ=`  
`github.com/ptsypyshev/golang_semver_learning/v2 v2.0.0 h1:4ZuRXV3+tmlQXE4INjQFCYVuS0x4Ad1s/JNjxgJJGhU=`  
`github.com/ptsypyshev/golang_semver_learning/v2 v2.0.0/go.mod h1:bbFjgY7hrPG4ikhqpi8RcCw+0ku61GhVisBhaTFf9mA=`  