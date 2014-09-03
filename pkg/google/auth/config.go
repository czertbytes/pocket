package auth

import "net/http"

type Config struct {
	client             *http.Client
	ProjectId          string
	ProjectClientEmail string
	AuthScope          string
	AuthTokenURI       string
	AuthPrivateKey     []byte
}

func NewConfig(client *http.Client) *Config {
	return &Config{
		client: client,
	}
}

func (self *Config) clone() *Config {
	x := *self

	return &x
}

func (self *Config) WithProjectId(projectId string) *Config {
	self = self.clone()
	self.ProjectId = projectId

	return self
}

func (self *Config) WithProjectClientEmail(projectClientEmail string) *Config {
	self = self.clone()
	self.ProjectClientEmail = projectClientEmail

	return self
}

func (self *Config) WithAuthScope(authScope string) *Config {
	self = self.clone()
	self.AuthScope = authScope

	return self
}

func (self *Config) WithAuthTokenURI(authTokenURI string) *Config {
	self = self.clone()
	self.AuthTokenURI = authTokenURI

	return self
}

func (self *Config) WithAuthPrivateKey(authPrivateKey []byte) *Config {
	self = self.clone()
	self.AuthPrivateKey = authPrivateKey

	return self
}
