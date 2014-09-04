package proof

import (
	"fmt"

	"appengine"

	c "github.com/czertbytes/pocket/pkg/clients"
	s "github.com/czertbytes/pocket/pkg/social"
	t "github.com/czertbytes/pocket/pkg/types"
	u "github.com/czertbytes/pocket/pkg/users"
)

var (
	ErrProofAuthOriginServiceNotValid error = fmt.Errorf("proof: AuthOriginService is not valid!")
)

type Proof struct {
	AppEngineContext appengine.Context
	Clients          *c.Clients
	Users            *u.Users
}

func NewProof(appEngineContext appengine.Context) *Proof {
	return &Proof{
		AppEngineContext: appEngineContext,
		Clients:          c.NewClients(appEngineContext),
		Users:            u.NewUsers(appEngineContext),
	}
}

func (self *Proof) Login(client *t.Client) error {
	profile, err := self.validateAuthToken(client.AuthOrigin)
	if err != nil {
		return err
	}

	user, err := self.Users.FindByEmail(profile.Email)
	if err != nil {
		if err != u.ErrUserNotFound {
			return err
		}

		user := &t.User{
			Status:   t.UserStatusActive,
			FullName: profile.FullName,
			Email:    profile.Email,
		}

		switch client.AuthOrigin.Service {
		case t.AuthOriginServiceGooglePlus:
			user.GooglePlusId = client.AuthOrigin.EntityId
		case t.AuthOriginServiceFacebook:
			user.FacebookId = client.AuthOrigin.EntityId
		}

		if err := self.Users.Create(user); err != nil {
			return err
		}

		client.Status = t.ClientStatusActive
		client.UserId = user.Id
		client.User = *user
		if err := self.Clients.Create(client); err != nil {
		}
	}

	newClient, err := self.Clients.FindByUserId(user.Id)
	if err != nil {
		return nil
	}

	newClient.RegenerateToken()

	if _, err := self.Clients.Update(newClient); err != nil {
		return err
	}

	client = &newClient

	return nil
}

func (self *Proof) Logout(client *t.Client) error {
	newClient, err := self.Clients.FindByClientIdAndToken(client.ClientId, client.ClientToken.Value)
	if err != nil {
		return err
	}

	newClient.RegenerateToken()

	if _, err := self.Clients.Update(newClient); err != nil {
		return err
	}

	return nil
}

func (self *Proof) validateAuthToken(authOrigin t.AuthOrigin) (s.Profile, error) {
	var fetcher s.Fetcher

	switch authOrigin.Service {
	case t.AuthOriginServiceGooglePlus:
		fetcher = s.NewGooglePlus(self.AppEngineContext)
	case t.AuthOriginServiceFacebook:
		fetcher = s.NewFacebook(self.AppEngineContext, FacebookAppSecret)
	default:
		return s.Profile{}, ErrProofAuthOriginServiceNotValid
	}

	return fetcher.Fetch(authOrigin.EntityId, authOrigin.Token)
}
