package auth

import (
	"net/http"
	"net/url"

	shttp "github.com/czertbytes/pocket/pkg/http"
	t "github.com/czertbytes/pocket/pkg/types"
)

const (
	LocationPath = "http://api.tripmoneymgmt.com/manager/clients"
)

func Post(url *url.URL, header http.Header, client *t.Client, requestContext *shttp.RequestContext) (int, http.Header, *t.Client, error) {
	if err := NewController(RequestContext).Post(client); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusCreated, nil, client, nil
}

func Delete(url *url.URL, header http.Header, _ interface{}, RequestContext *shttp.RequestContext) (int, http.Header, interface{}, error) {
	if err := NewController(RequestContext).Delete(url); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusNoContent, nil, nil, nil
}
