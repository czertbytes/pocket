package http

import (
	"net/http"
	"strings"

	tt "github.com/czertbytes/go-tigertonic"
)

type RequestSourceHandler struct {
	handler http.Handler
}

func RequestSourceHandled(handler http.Handler) *RequestSourceHandler {
	return &RequestSourceHandler{
		handler: handler,
	}
}

func (self *RequestSourceHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	requestSource := &RequestSource{
		UserAgent:  request.UserAgent(),
		IPAddress:  requestIPAddress(request),
		Country:    request.Header.Get("X-AppEngine-Country"),
		Region:     request.Header.Get("X-AppEngine-Region"),
		City:       request.Header.Get("X-AppEngine-City"),
		CityLatLng: request.Header.Get("X-AppEngine-CityLatLong"),
	}

	tt.Context(request).(*RequestContext).RequestSource = requestSource

	self.handler.ServeHTTP(responseWriter, request)
}

func requestIPAddress(request *http.Request) string {
	ipAddress := request.RemoteAddr
	xff := request.Header.Get("X-Forwarded-For")
	if len(xff) > 0 {
		xffParts := strings.Split(xff, ",")
		ipAddress = xffParts[0]
	}

	return ipAddress
}
