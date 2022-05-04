package checker

import (
	"fmt"
	"log"
	"net/http"
)

func CheckWebServer(urlStr string) {
	resp, err := http.Get(urlStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The status code we got is:", resp.StatusCode)
}
