package http

import (
	"net/http"

	"appengine"

	tt "github.com/czertbytes/go-tigertonic"
)

type AppEngineHandler struct {
	handler http.Handler
}

func AppEngineHandled(handler http.Handler) *AppEngineHandler {
	return &AppEngineHandler{
		handler: handler,
	}
}

func (self *AppEngineHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	appEngineContext := appengine.NewContext(request)
	tt.Context(request).(*RequestContext).AppEngineContext = appEngineContext

	self.handler.ServeHTTP(responseWriter, request)
}
