package middleware_setup

import (
	"logger"
	"net/http"
	responseheader "response_header"
)

// TODO: add coments

// TODO: change implementation to set headers dynamically

func SetupMiddleware(handler http.Handler) http.Handler {

	return logger.NewLogger(responseheader.NewResponseHeader(handler, "Content-Type", "application/json"))
}
