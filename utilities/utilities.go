package utilities

import (
	"math/rand"
	"strconv"
	"time"
)

func GetDate(format ...string) string {

	if len(format) > 0 {
		return time.Now().Format(format[0])

	}
	return time.Now().Format("01-02-2006")
}

func GetId() string {
	return strconv.Itoa(rand.Intn(1000))
}
