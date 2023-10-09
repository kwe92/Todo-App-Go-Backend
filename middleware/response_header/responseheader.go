package responseheader

import "net/http"

// TODO: add coments

type ResponseHeader struct {
	Handler     http.Handler
	HeaderName  string
	HeaderValue string
}

func (rh *ResponseHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(rh.HeaderName, rh.HeaderValue)
	rh.Handler.ServeHTTP(w, r)
}

func NewResponseHeader(handler http.Handler, headerName string, headerValue string) *ResponseHeader {
	return &ResponseHeader{
		Handler:     handler,
		HeaderName:  headerName,
		HeaderValue: headerValue,
	}
}
