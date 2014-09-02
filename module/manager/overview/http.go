package overview

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
	if err := NewController(requestContext).Post(overview); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusCreated, nil, overview, nil
}

func Get(url *url.URL, header http.Header, _ interface{}, requestContext *shttp.RequestContext) (int, http.Header, *t.Overview, error) {
	overview, err := NewController(requestContext).Get(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, overview, nil
}

func Put(url *url.URL, header http.Header, overview *t.Overview, requestContext *shttp.RequestContext) (int, http.Header, *t.Overview, error) {
	if err := NewController(requestContext).Put(url, overview); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, overview, nil
}

func Patch(url *url.URL, header http.Header, overview *t.Overview, requestContext *shttp.RequestContext) (int, http.Header, *t.Overview, error) {
	if err := NewController(requestContext).Patch(url, overview); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, overview, nil
}

func Delete(url *url.URL, header http.Header, _ interface{}, requestContext *shttp.RequestContext) (int, http.Header, interface{}, error) {
	if err := NewController(requestContext).Delete(url); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusNoContent, nil, nil, nil
}

func GetPayments(url *url.URL, header http.Header, _ interface{}, requestContext *shttp.RequestContext) (int, http.Header, t.Payments, error) {
	payments, err := NewController(requestContext).GetOverviewPayments(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, payments, nil
}

func PostParticipants(url *url.URL, header http.Header, user *t.User, requestContext *shttp.RequestContext) (int, http.Header, *t.User, error) {
	if err := NewController(requestContext).PostOverviewParticipant(user); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusCreated, nil, user, nil
}

func GetParticipants(url *url.URL, header http.Header, _ interface{}, requestContext *shttp.RequestContext) (int, http.Header, t.Users, error) {
	participants, err := NewController(requestContext).GetOverviewParticipants(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, participants, nil
}
