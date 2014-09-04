package proof

import (
	"fmt"
	"time"

	"appengine"
	"appengine/delay"
	"appengine/taskqueue"

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

	newClient, err := self.updateClient(user.Id)
	if err != nil {
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

func (self *Proof) updateClient(userId t.UserId) (t.Client, error) {
	client, err := self.Clients.FindByUserId(userId)
	if err != nil {
		return t.Client{}, nil
	}

	client.RegenerateToken()

	if _, err := self.Clients.Update(client); err != nil {
		return t.Client{}, err
	}

	if err := self.expireTokenTask(client.Id); err != nil {
		return t.Client{}, err
	}

	return client, nil
}

func (self *Proof) expireTokenTask(id t.ClientId) error {
	checkExpiredClientTokenTask, err := checkExpiredClientTokenFunc.Task(id)
	if err != nil {
		return err
	}

	checkExpiredClientTokenTask.Delay = t.ClientTokenExpirationTime
	if _, err := taskqueue.Add(self.AppEngineContext, checkExpiredClientTokenTask, "check-expired-tokens"); err != nil {
		return err
	}

	return nil
}

var checkExpiredClientTokenFunc = delay.Func("expired-client-token", func(appEngineContext appengine.Context, id t.ClientId) {
	Clients := c.NewClients(appEngineContext)

	client, err := Clients.Find(id)
	if err != nil {
		appEngineContext.Errorf("proof: Finding Client for expiration failed with error %s!", err)
		return
	}

	location, _ := time.LoadLocation(t.DefaultLocation)
	now := time.Now().In(location)

	if client.ClientToken.ExpireAt.Before(now) {
		client.RegenerateToken()

		if _, err := Clients.Update(client); err != nil {
			appEngineContext.Errorf("proof: Regenerating expired Client token failed with error %s!", err)
			return
		}
	}
})
