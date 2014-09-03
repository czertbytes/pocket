package auth

import (
	"io"
	"net/http"

	"code.google.com/p/goauth2/oauth"
	"code.google.com/p/goauth2/oauth/jwt"

	"appengine"
	"appengine/memcache"
)

const (
	TokenCacheKey string = "oauth-token"
)

type Auth struct {
	AppEngineContext appengine.Context
	config           *Config
}

func NewAuth(appEngineContext appengine.Context, config *Config) *Auth {
	newAppEngineContext, err := appengine.Namespace(appEngineContext, "google-auth")
	if err != nil {
		appEngineContext.Errorf("auth: Creating namespace context failed with error: %s", err)
		newAppEngineContext = appEngineContext
	}

	return &Auth{
		AppEngineContext: newAppEngineContext,
		config:           config,
	}
}

func (self *Auth) NewOAuth2Request(method, url string, body io.Reader) (*http.Request, string, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, "", err
	}

	accessToken, err := self.getOAuth2AccessToken()
	if err != nil {
		return nil, "", err
	}

	request.Header.Set("x-goog-api-version", "2")
	request.Header.Set("x-goog-project-id", self.config.ProjectId)

	return request, accessToken, nil
}

func (self *Auth) getOAuth2AccessToken() (string, error) {
	var oauthToken oauth.Token
	var accessToken string
	var err error

	_, err = memcache.JSON.Get(self.AppEngineContext, TokenCacheKey, &oauthToken)
	if err != nil || oauthToken.Expired() {
		accessToken, err = self.newOAuthTokens()
		if err != nil {
			return "", err
		}
	} else {
		accessToken = oauthToken.AccessToken
	}

	return accessToken, nil
}

func (self *Auth) newOAuthTokens() (string, error) {
	jwtToken := jwt.NewToken(self.config.ProjectClientEmail, self.config.AuthScope, self.config.AuthPrivateKey)
	jwtToken.ClaimSet.Aud = self.config.AuthTokenURI

	oauthToken, err := jwtToken.Assert(self.config.client)
	if err != nil {
		return "", err
	}

	err = memcache.JSON.Set(self.AppEngineContext, &memcache.Item{
		Key:    TokenCacheKey,
		Object: oauthToken,
	})

	return oauthToken.AccessToken, err
}

func (self *Auth) Do(request *http.Request) (*http.Response, error) {
	return self.config.client.Do(request)
}
