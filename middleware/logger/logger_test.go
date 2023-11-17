package logger

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNewLogger(t *testing.T) {
	newLogger := NewLogger(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))

	expected := &Logger{
		Handler: http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}),
		Prefix:  "",
	}

	fmt.Printf("type from NewLogger: %T | type from expected %T\n", newLogger, expected)

	fmt.Println(*expected == *newLogger)

}
