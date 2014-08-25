package user

import (
	"net/http"
	"net/url"

	shttp "github.com/czertbytes/pocket/pkg/http"
	t "github.com/czertbytes/pocket/pkg/types"
)

func Get(url *url.URL, header http.Header, _ interface{}, RequestContext *shttp.RequestContext) (int, http.Header, *t.User, error) {
	comment, err := NewController(RequestContext).Get(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, comment, nil
}

func Delete(url *url.URL, header http.Header, _ interface{}, RequestContext *shttp.RequestContext) (int, http.Header, interface{}, error) {
	if err := NewController(RequestContext).Delete(url); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusNoContent, nil, nil, nil
}

func GetOverviews(url *url.URL, header http.Header, _ interface{}, RequestContext *shttp.RequestContext) (int, http.Header, t.Overviews, error) {
	comment, err := NewController(RequestContext).Get(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, comment, nil
}

func GetPayments(url *url.URL, header http.Header, _ interface{}, RequestContext *shttp.RequestContext) (int, http.Header, t.Payments, error) {
	comment, err := NewController(RequestContext).Get(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, comment, nil
}

func GetComments(url *url.URL, header http.Header, _ interface{}, RequestContext *shttp.RequestContext) (int, http.Header, t.Comments, error) {
	comment, err := NewController(RequestContext).Get(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, comment, nil
}

func GetMe(url *url.URL, header http.Header, _ interface{}, RequestContext *shttp.RequestContext) (int, http.Header, *t.User, error) {
	comment, err := NewController(RequestContext).Get(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, comment, nil
}

func GetMeOverviews(url *url.URL, header http.Header, _ interface{}, RequestContext *shttp.RequestContext) (int, http.Header, t.Overviews, error) {
	comment, err := NewController(RequestContext).Get(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, comment, nil
}

func GetMePayments(url *url.URL, header http.Header, _ interface{}, RequestContext *shttp.RequestContext) (int, http.Header, t.Payments, error) {
	comment, err := NewController(RequestContext).Get(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, comment, nil
}

func GetMeComments(url *url.URL, header http.Header, _ interface{}, RequestContext *shttp.RequestContext) (int, http.Header, t.Comments, error) {
	comment, err := NewController(RequestContext).Get(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, comment, nil
}
