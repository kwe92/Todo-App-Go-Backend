package utilities

import (
	"constants"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
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

// jsonEncode converts [object] to a JSON string and sends it over the stream as a response.
func JsonEncode(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

// TODO: fix jsonDecode
func JsonDecode(r *http.Request, ptr *interface{}) {
	json.NewDecoder(r.Body).Decode(&ptr)
}

// setHeader sets up HTTP Header meta data.
func SetHeader(w http.ResponseWriter) {

	w.Header().Set(metaData.ContentTypeHeader, metaData.MediaTypeJson)

}
