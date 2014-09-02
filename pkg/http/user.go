package http

import (
	"net/http"

	tt "github.com/czertbytes/go-tigertonic"
	c "github.com/czertbytes/pocket/pkg/clients"
	u "github.com/czertbytes/pocket/pkg/users"
)

type UserHandler struct {
	handler http.Handler
}

func UserHandled(handler http.Handler) *UserHandler {
	return &UserHandler{
		handler: handler,
	}
}

func (self *UserHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	appEngineContext := tt.Context(request).(*RequestContext).AppEngineContext
	clientId := tt.Context(request).(*RequestContext).ClientId
	token := tt.Context(request).(*RequestContext).Token

	client, err := c.NewClients(appEngineContext).FindByClientIdAndToken(clientId, token)
	if err != nil {
		UnauthorizedResponsePayload(appEngineContext, responseWriter, err)
		return
	}

	user, err := u.NewUsers(appEngineContext).Find(client.UserId)
	if err != nil {
		UnauthorizedResponsePayload(appEngineContext, responseWriter, err)
		return
	}

	tt.Context(request).(*RequestContext).Client = &client
	tt.Context(request).(*RequestContext).User = &user

	self.handler.ServeHTTP(responseWriter, request)
}
