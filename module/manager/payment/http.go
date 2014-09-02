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

func Post(url *url.URL, header http.Header, payment *t.Payment, requestContext *shttp.RequestContext) (int, http.Header, *t.Payment, error) {
	if err := NewController(requestContext).Post(payment); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusCreated, nil, payment, nil
}

func Get(url *url.URL, header http.Header, _ interface{}, requestContext *shttp.RequestContext) (int, http.Header, *t.Payment, error) {
	payment, err := NewController(requestContext).Get(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, payment, nil
}

func Put(url *url.URL, header http.Header, payment *t.Payment, requestContext *shttp.RequestContext) (int, http.Header, *t.Payment, error) {
	updatedPayment, err := NewController(requestContext).Put(url, payment)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, updatedPayment, nil
}

func Patch(url *url.URL, header http.Header, payment *t.Payment, requestContext *shttp.RequestContext) (int, http.Header, *t.Payment, error) {
	patchedPayment, err := NewController(requestContext).Patch(url, payment)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, patchedPayment, nil
}

func Delete(url *url.URL, header http.Header, _ interface{}, requestContext *shttp.RequestContext) (int, http.Header, interface{}, error) {
	if err := NewController(requestContext).Delete(url); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusNoContent, nil, nil, nil
}

func PostDocuments(url *url.URL, header http.Header, requestContext *shttp.RequestContext) (int, http.Header, interface{}, error) {
	if err := NewController(requestContext).PostDocuments(); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusCreated, nil, nil, nil
}

func GetDocuments(url *url.URL, header http.Header, _ interface{}, requestContext *shttp.RequestContext) (int, http.Header, t.Documents, error) {
	documents, err := NewController(requestContext).GetDocuments(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, documents, nil
}

func PostComments(url *url.URL, header http.Header, comment *t.Comment, requestContext *shttp.RequestContext) (int, http.Header, *t.Comment, error) {
	if err := NewController(requestContext).PostComment(comment); err != nil {
		return 0, nil, nil, err
	}

	return http.StatusCreated, nil, comment, nil
}

func GetComments(url *url.URL, header http.Header, _ interface{}, requestContext *shttp.RequestContext) (int, http.Header, t.Comments, error) {
	comments, err := NewController(requestContext).GetComments(url)
	if err != nil {
		return 0, nil, nil, err
	}

	return http.StatusOK, nil, comments, nil
}
