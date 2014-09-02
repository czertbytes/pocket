package http

import (
	"net/http"

	tt "github.com/czertbytes/go-tigertonic"
)

type EntityIdHandler struct {
	handler http.Handler
}

func EntityIdHandled(handler http.Handler) *EntityIdHandler {
	return &EntityIdHandler{
		handler: handler,
	}
}

func (self *EntityIdHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	appEngineContext := tt.Context(request).(*RequestContext).AppEngineContext

	if _, hasEntityIdValue := request.URL.Query()["id"]; hasEntityIdValue {
		entityId, err := RequestParamInt64(request.URL, "id")
		if err != nil {
			EntityIdNotValidResponsePayload(appEngineContext, responseWriter, err)
			return
		}

		tt.Context(request).(*RequestContext).EntityId = entityId
	}

	self.handler.ServeHTTP(responseWriter, request)
}
