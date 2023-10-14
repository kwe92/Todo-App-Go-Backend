package responseheader

import "net/http"

// ResponseHeader: middleware that sets the response header.
type ResponseHeader struct {
	Handler     http.Handler
	HeaderName  string
	HeaderValue string
}

// ServeHTTP: handlers the request by adding response headers.
func (rh *ResponseHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(rh.HeaderName, rh.HeaderValue)
	rh.Handler.ServeHTTP(w, r)
}

// NewResponseHeader: returns a new ResponseHeader object wrapped around the handler.
func NewResponseHeader(handler http.Handler, headerName string, headerValue string) *ResponseHeader {
	return &ResponseHeader{
		Handler:     handler,
		HeaderName:  headerName,
		HeaderValue: headerValue,
	}
}
