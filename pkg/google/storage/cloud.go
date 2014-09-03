package storage

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/urlfetch"

	ga "github.com/czertbytes/pocket/pkg/google/auth"
)

type Cloud struct {
	AppEngineContext appengine.Context
	auth             *ga.Auth
}

func NewCloud(appEngineContext appengine.Context) *Cloud {
	config := ga.NewConfig(urlfetch.Client(appEngineContext)).
		WithProjectId(ga.ProjectId).
		WithProjectClientEmail(ga.ProjectClientEmail).
		WithAuthScope(ga.AuthScope).
		WithAuthTokenURI(ga.AuthTokenURI).
		WithAuthPrivateKey(ga.AuthPrivateKey)

	return &Cloud{
		AppEngineContext: appEngineContext,
		auth:             ga.NewAuth(appEngineContext, config),
	}
}

func (self *Cloud) Save(cloudObject *CloudObject) error {
	response, err := self.do("POST", cloudObject)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Object storage response code is not OK. Return code %d", response.StatusCode)
	}

	return nil
}

func (self *Cloud) do(method string, cloudObject *CloudObject) (*http.Response, error) {
	request, accessToken, err := self.auth.NewOAuth2Request(method, cloudObject.CreateURLPath(), cloudObject.Content)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Content-Type", cloudObject.ContentType)

	return self.auth.Do(request)
}
