package user

import (
	"net/http"
	"net/url"

	shttp "github.com/czertbytes/pocket/pkg/http"
	t "github.com/czertbytes/pocket/pkg/types"
)

func Get(url *url.URL, header http.Header, _ interface{}, requestContext *shttp.RequestContext) (int, http.Header, *t.User, error) {
	user, err := NewController(requestContext).Get(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, &user, nil
}

func Delete(url *url.URL, header http.Header, _ interface{}, requestContext *shttp.RequestContext) (int, http.Header, interface{}, error) {
	if err := NewController(requestContext).Delete(url); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusNoContent, nil, nil, nil
}

func GetOverviews(url *url.URL, header http.Header, _ interface{}, requestContext *shttp.RequestContext) (int, http.Header, t.Overviews, error) {
	overviews, err := NewController(requestContext).GetOverviews(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, overviews, nil
}

func GetPayments(url *url.URL, header http.Header, _ interface{}, requestContext *shttp.RequestContext) (int, http.Header, t.Payments, error) {
	payments, err := NewController(requestContext).GetPayments(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, payments, nil
}

func GetComments(url *url.URL, header http.Header, _ interface{}, requestContext *shttp.RequestContext) (int, http.Header, t.Comments, error) {
	comments, err := NewController(requestContext).GetComments(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, comments, nil
}

func GetMe(url *url.URL, header http.Header, _ interface{}, requestContext *shttp.RequestContext) (int, http.Header, *t.User, error) {
	user, err := NewController(requestContext).Get(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, &user, nil
}

func GetMeOverviews(url *url.URL, header http.Header, _ interface{}, requestContext *shttp.RequestContext) (int, http.Header, t.Overviews, error) {
	overviews, err := NewController(requestContext).GetOverviews(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, overviews, nil
}

func GetMePayments(url *url.URL, header http.Header, _ interface{}, requestContext *shttp.RequestContext) (int, http.Header, t.Payments, error) {
	payments, err := NewController(requestContext).GetPayments(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, payments, nil
}

func GetMeComments(url *url.URL, header http.Header, _ interface{}, requestContext *shttp.RequestContext) (int, http.Header, t.Comments, error) {
	comments, err := NewController(requestContext).GetComments(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, comments, nil
}
