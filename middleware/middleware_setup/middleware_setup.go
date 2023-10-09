package middleware_setup

import (
	"logger"
	"net/http"
	responseheader "response_header"
)

// SetupMiddleware: wraps the handler passed in with middleware handlers.
func SetupMiddleware(handler http.Handler, headers ...map[string]string) http.Handler {

	if len(headers) > 0 {
		var key string
		var value string
		for k, v := range headers[0] {
			key = k
			value = v
		}
		return logger.NewLogger(responseheader.NewResponseHeader(handler, key, value))
	}
	return logger.NewLogger(responseheader.NewResponseHeader(handler, "Content-Type", "application/json"))
}
