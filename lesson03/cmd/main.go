package main

import (
	"fmt"
	"github.com/ptsypyshev/golang_semver_learning/checker"
	checkerV2 "github.com/ptsypyshev/golang_semver_learning/v2/checker"
)

func main() {
	checker.CheckWebServer("https://www.google.com")

	urls := []string{
		"https://www.google.com",
		"https://yandex.ru/",
		"https://mail.ru/",
		"https://gb.ru/",
	}
	checkResult := checkerV2.CheckBulk(urls)
	for u, code := range checkResult {
		fmt.Printf("The status code for %s is: %d\n", u, code)
	}
}
