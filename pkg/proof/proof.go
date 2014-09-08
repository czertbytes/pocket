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

func (self *Proof) Login(client *t.Client) (t.Client, error) {
	profile, err := self.validateAuthToken(client.AuthOrigin)
	if err != nil {
		return t.Client{}, err
	}

	user, err := self.Users.FindByEmail(profile.Email)
	if err == nil {
		// user exists
		newClient, err := self.updateClient(*client, user)
		if err != nil {
			return t.Client{}, err
		}

		return newClient, nil
	}

	if err != u.ErrUserNotFound {
		return t.Client{}, err
	}

	// user does not exist
	newClient, _, err := self.createNewUser(*client, profile)
	if err != nil {
		return t.Client{}, err
	}

	return newClient, nil
}

func (self *Proof) Logout(client *t.Client) error {
	newClient, err := self.Clients.FindByClientIdAndToken(client.ClientId, client.ClientToken.Value)
	if err != nil {
		return err
	}

	newClient.GenerateClientId()
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

func (self *Proof) updateClient(client t.Client, user t.User) (t.Client, error) {
	newClient, err := self.Clients.FindByUserId(user.Id)
	if err != nil {
		return t.Client{}, nil
	}

	newClient.GenerateClientId()
	newClient.RegenerateToken()

	if _, err := self.Clients.Update(newClient); err != nil {
		return t.Client{}, err
	}

	if err := self.expireTokenTask(newClient.Id); err != nil {
		return t.Client{}, err
	}

	newClient.AuthOrigin = client.AuthOrigin
	newClient.User = user
	newClient.SetFormattedValues()

	return newClient, nil
}

func (self *Proof) createNewUser(client t.Client, profile s.Profile) (t.Client, t.User, error) {
	newUser := &t.User{
		Status:   t.UserStatusActive,
		FullName: profile.FullName,
		Email:    profile.Email,
	}

	switch client.AuthOrigin.Service {
	case t.AuthOriginServiceGooglePlus:
		newUser.GooglePlusId = client.AuthOrigin.EntityId
	case t.AuthOriginServiceFacebook:
		newUser.FacebookId = client.AuthOrigin.EntityId
	}

	if err := self.Users.Create(newUser); err != nil {
		return t.Client{}, t.User{}, err
	}

	draftClient := &t.Client{
		Status: t.ClientStatusActive,
		UserId: newUser.Id,
		User:   *newUser,
	}
	if err := self.Clients.Create(draftClient); err != nil {
		return t.Client{}, t.User{}, err
	}

	draftClient.GenerateClientId()
	draftClient.RegenerateToken()

	newClient, err := self.Clients.Update(*draftClient)
	if err != nil {
		return t.Client{}, t.User{}, err
	}

	if err := self.expireTokenTask(newClient.Id); err != nil {
		return t.Client{}, t.User{}, err
	}

	newClient.AuthOrigin = client.AuthOrigin
	newClient.User = *newUser
	newClient.SetFormattedValues()

	return newClient, *newUser, nil
}

func (self *Proof) expireTokenTask(id t.ClientId) error {
	checkExpiredClientTokenTask, err := checkExpiredClientTokenFunc.Task(id)
	if err != nil {
		return err
	}

	hostName, _ := appengine.ModuleHostname(self.AppEngineContext, appengine.ModuleName(self.AppEngineContext), "", "")

	checkExpiredClientTokenTask.Header = make(map[string][]string)
	checkExpiredClientTokenTask.Header.Set("Host", hostName)
	checkExpiredClientTokenTask.Delay = t.ClientTokenExpirationTime
	if _, err := taskqueue.Add(self.AppEngineContext, checkExpiredClientTokenTask, "default"); err != nil {
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

	if client.ClientToken.ExpireAtTime.Before(now) {
		client.RegenerateToken()

		if _, err := Clients.Update(client); err != nil {
			appEngineContext.Errorf("proof: Regenerating expired Client token failed with error %s!", err)
			return
		}
	}
})
