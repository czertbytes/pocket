package payment

import (
	"net/http"
	"net/url"

	shttp "github.com/czertbytes/pocket/pkg/http"
	t "github.com/czertbytes/pocket/pkg/types"
)

const (
	LocationPath = "http://api.tripmoneymgmt.com/manager/overviews"
)

func Post(url *url.URL, header http.Header, overview *t.Overview, requestContext *shttp.RequestContext) (int, http.Header, *t.Overview, error) {
	if err := NewController(RequestContext).Post(overview); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusCreated, nil, overview, nil
}

func Get(url *url.URL, header http.Header, _ interface{}, RequestContext *shttp.RequestContext) (int, http.Header, *t.Overview, error) {
	overview, err := NewController(RequestContext).Get(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, overview, nil
}

func Put(url *url.URL, header http.Header, overview *t.Overview, RequestContext *shttp.RequestContext) (int, http.Header, *t.Overview, error) {
	if err := NewController(RequestContext).Put(url, overview); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, overview, nil
}

func Patch(url *url.URL, header http.Header, overview *t.Overview, RequestContext *shttp.RequestContext) (int, http.Header, *t.Overview, error) {
	if err := NewController(RequestContext).Patch(url, overview); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, overview, nil
}

func Delete(url *url.URL, header http.Header, _ interface{}, RequestContext *shttp.RequestContext) (int, http.Header, interface{}, error) {
	if err := NewController(RequestContext).Delete(url); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusNoContent, nil, nil, nil
}

func PostDocuments(url *url.URL, header http.Header, overview *t.Overview, requestContext *shttp.RequestContext) (int, http.Header, *t.Overview, error) {
	if err := NewController(RequestContext).PostDocuments(overview); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusCreated, nil, overview, nil
}

func GetDocuments(url *url.URL, header http.Header, _ interface{}, RequestContext *shttp.RequestContext) (int, http.Header, *t.Overview, error) {
	overview, err := NewController(RequestContext).Get(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, overview, nil
}

func PostComments(url *url.URL, header http.Header, overview *t.Overview, requestContext *shttp.RequestContext) (int, http.Header, *t.Overview, error) {
	if err := NewController(RequestContext).Post(overview); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusCreated, nil, overview, nil
}

func GetComments(url *url.URL, header http.Header, _ interface{}, RequestContext *shttp.RequestContext) (int, http.Header, *t.Overview, error) {
	overview, err := NewController(RequestContext).Get(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, overview, nil
}
