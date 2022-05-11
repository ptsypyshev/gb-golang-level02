package main

import (
	"flag"
	"fmt"
	"time"
)

// CustomError with datetime information
type CustomError struct {
	msg      string
	datetime time.Time
}

func (e CustomError) Error() string {
	return fmt.Sprintf("%s: %s", e.datetime.Format("02-01-2006 15:04:05"), e.msg)
}

// getFlag returns flag value or CustomError if something went wrong
func getFlag(name string) (res string, err error) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("recovered: ", v)
			err = CustomError{
				msg:      fmt.Sprintf("Used broken %v", v),
				datetime: time.Now(),
			}
		}
	}()
	brokenFlag := flag.String(name, "default value", "some default usage")
	flag.Parse()
	return *brokenFlag, nil
}

func main() {
	if result, err := getFlag("broken=flag"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
