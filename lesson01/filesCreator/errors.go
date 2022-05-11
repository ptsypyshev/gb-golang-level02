package filesCreator

import (
	"fmt"
	"time"
)

type CreateFileError struct {
	path       string
	innerError string
	errorTime  time.Time
}

func (e CreateFileError) Error() string {
	return fmt.Sprintf("%s: Cannot create file : %s\nReason: %s",
		e.errorTime.Format("02-01-2006 15:04:05"),
		e.path,
		e.innerError,
	)
}
