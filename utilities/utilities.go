package utilities

import (
	"constants"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
	"time"
)

// metaData represents meta data constants in a struct.
var metaData = constants.HeaderData()

func GetDate(format ...string) string {

	if len(format) > 0 {
		return time.Now().Format(format[0])

	}
	return time.Now().Format("01-02-2006")
}

func GetId() string {
	return strconv.Itoa(rand.Intn(1000))
}

// jsonEncode converts [data] to a JSON string and sends it over the stream as a response.
func JsonEncode[T any](w io.Writer, data T) error {
	return json.NewEncoder(w).Encode(data)
}

// JsonDecode reads the next JSON-encoded value from its input and stores it in the value pointed to by v.
func JsonDecode[T any](rc io.ReadCloser, ptr *T) error {
	return json.NewDecoder(rc).Decode(&ptr)
}

// setHeader sets up HTTP Header meta data.
func SetHeader(w http.ResponseWriter) {

	w.Header().Set(metaData.ContentTypeHeader, metaData.MediaTypeJson)

}

// ParseURL returns the request host concatenated to the path.
func ParseURL(r *http.Request) string {
	return fmt.Sprintln(r.Host + r.URL.Path)
}

func CheckError(err error) {
	if err != nil {
		log.Fatalf("\nfailed to create request: %s", err.Error())
	}
}

func MatchStatusCode(t *testing.T, statusCode int, expectedStatusCode ...int) {

	var expectedCode int

	if len(expectedStatusCode) > 0 {
		expectedCode = expectedStatusCode[0]
	} else {
		expectedCode = http.StatusOK
	}

	if statusCode != expectedCode {
		t.Fatalf("the status code should be [%d] but received [%d]",
			expectedCode,
			statusCode,
		)
		return
	}
}

func MatchContent[T any](t *testing.T, expected T, received T) {
	if fmt.Sprint(received) != fmt.Sprint(expected) {

		t.Fatalf("the response body should be [%s] but received [%s]",
			fmt.Sprint(expected),
			fmt.Sprint(received),
		)

	}

}
