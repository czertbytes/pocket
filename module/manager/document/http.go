package document

import (
	"net/http"
	"net/url"

	shttp "github.com/czertbytes/pocket/pkg/http"
	t "github.com/czertbytes/pocket/pkg/types"
)

func Get(url *url.URL, header http.Header, _ interface{}, RequestContext *shttp.RequestContext) (int, http.Header, *t.Document, error) {
	document, err := NewController(RequestContext).Get(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, document, nil
}

func Delete(url *url.URL, header http.Header, _ interface{}, RequestContext *shttp.RequestContext) (int, http.Header, interface{}, error) {
	if err := NewController(RequestContext).Delete(url); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusNoContent, nil, nil, nil
}
