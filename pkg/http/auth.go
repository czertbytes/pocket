package http

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	tt "github.com/czertbytes/go-tigertonic"
	t "github.com/czertbytes/pocket/pkg/types"
)

type AuthHandler struct {
	handler http.Handler
}

func AuthHandled(handler http.Handler) *AuthHandler {
	return &AuthHandler{
		handler: handler,
	}
}

func (self *AuthHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	clientId, token := parseAuthHeader(request)

	tt.Context(request).(*RequestContext).ClientId = clientId
	tt.Context(request).(*RequestContext).Token = token

	self.handler.ServeHTTP(responseWriter, request)
}

func parseAuthHeader(request *http.Request) (t.ClientClientId, string) {
	fullAuthHeader := request.Header.Get("Authorization")
	if len(fullAuthHeader) == 0 {
		return t.ClientClientId(""), ""
	}

	var authHash string
	if _, err := fmt.Sscanf(fullAuthHeader, "Bearer %s", &authHash); err != nil {
		return t.ClientClientId(""), ""
	}

	return parseClientIdAuthToken(authHash)
}

func parseClientIdAuthToken(authHash string) (t.ClientClientId, string) {
	decodedAuthHash, err := base64.StdEncoding.DecodeString(authHash)
	if err != nil {
		return t.ClientClientId(""), ""
	}

	authHashParts := strings.Split(string(decodedAuthHash), " ")
	if len(authHashParts) != 2 {
		return t.ClientClientId(""), ""
	}

	return t.ClientClientId(strings.TrimSpace(authHashParts[0])), strings.TrimSpace(authHashParts[1])
}
