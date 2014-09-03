package comment

import (
	"net/http"
	"net/url"

	h "github.com/czertbytes/pocket/pkg/http"
	t "github.com/czertbytes/pocket/pkg/types"
)

func Get(url *url.URL, header http.Header, _ interface{}, RequestContext *h.RequestContext) (int, http.Header, *t.Comment, error) {
	comment, err := NewController(RequestContext).Get(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, &comment, nil
}

func Delete(url *url.URL, header http.Header, _ interface{}, RequestContext *h.RequestContext) (int, http.Header, interface{}, error) {
	if err := NewController(RequestContext).Delete(url); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusNoContent, nil, nil, nil
}
