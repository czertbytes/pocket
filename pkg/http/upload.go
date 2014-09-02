package http

import (
	"encoding/json"
	"net/http"
	"net/url"

	tt "github.com/czertbytes/go-tigertonic"
)

type UploadFunc func(url *url.URL, header http.Header, requestContext *RequestContext) (int, http.Header, interface{}, error)

type UploadHandler struct {
	handler UploadFunc
}

func UploadHandled(handler UploadFunc) *UploadHandler {
	return &UploadHandler{
		handler: handler,
	}
}

func (self *UploadHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	multipartReader, err := request.MultipartReader()
	if err != nil {
		writeError(responseWriter)
		return
	}

	tt.Context(request).(*RequestContext).MultipartReader = multipartReader

	status, headers, _, err := self.handler(request.URL, request.Header, tt.Context(request).(*RequestContext))

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Header().Set("Location", headers.Get("Location"))
	responseWriter.WriteHeader(status)
	if err != nil {
		json.NewEncoder(responseWriter).Encode(map[string]string{
			"error": err.Error(),
		})
	}
}

func writeError(responseWriter http.ResponseWriter) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(responseWriter).Encode(map[string]string{
		"error": "Parsing multipart content failed!!",
	})
}
