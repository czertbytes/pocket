package auth

import (
	"net/http"
	"net/url"

	h "github.com/czertbytes/pocket/pkg/http"
	t "github.com/czertbytes/pocket/pkg/types"
)

const (
	LocationPath = "http://api.tripmoneymgmt.com/manager/clients"
)

func Post(url *url.URL, header http.Header, client *t.Client, requestContext *h.RequestContext) (int, http.Header, *t.Client, error) {
	newClient, err := NewController(requestContext).Post(client)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusCreated, nil, &newClient, nil
}

func Delete(url *url.URL, header http.Header, _ interface{}, requestContext *h.RequestContext) (int, http.Header, interface{}, error) {
	if err := NewController(requestContext).Delete(url); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusNoContent, nil, nil, nil
}
