package user

import (
	"fmt"

	"appengine"

	c "github.com/czertbytes/pocket/pkg/comments"
	shttp "github.com/czertbytes/pocket/pkg/http"
	o "github.com/czertbytes/pocket/pkg/overviews"
	p "github.com/czertbytes/pocket/pkg/payments"
	t "github.com/czertbytes/pocket/pkg/types"
	u "github.com/czertbytes/pocket/pkg/users"
)

var (
	ErrUserUnauthorized error = fmt.Errorf("user: Unauthorized access")
)

type Service struct {
	AppEngineContext appengine.Context
	RequestContext   *shttp.RequestContext
	notificator      *Notificator
	Overviews        *o.Overviews
	Payments         *p.Payments
	Comments         *c.Comments
	Users            *u.Users
}

func NewService(RequestContext *shttp.RequestContext) *Service {
	return &Service{
		AppEngineContext: RequestContext.AppEngineContext,
		RequestContext:   RequestContext,
		notificator:      NewNotificator(RequestContext),
		Overviews:        o.NewOverviews(RequestContext.AppEngineContext),
		Payments:         p.NewPayments(RequestContext.AppEngineContext),
		Comments:         c.NewComments(RequestContext.AppEngineContext),
		Users:            u.NewUsers(RequestContext.AppEngineContext),
	}
}

func (self *Service) Find(id t.UserId, user *t.User) (t.User, error) {
	return self.getUser(id, user)
}

func (self *Service) Delete(id t.UserId, user *t.User) error {
	if _, err := self.getUser(id, user); err != nil {
		return err
	}

	if err := self.Users.Delete(id); err != nil {
		return err
	}

	return nil
}

func (self *Service) FindAllOverviews(id t.UserId, user *t.User) (t.Overviews, error) {
	if _, err := self.getUser(id, user); err != nil {
		return nil, err
	}

	userOverviews, err := self.Overviews.FindAllByUserId(id)
	if err != nil {
		return nil, err
	}

	return userOverviews, nil
}

func (self *Service) FindAllPayments(id t.UserId, user *t.User) (t.Payments, error) {
	if _, err := self.getUser(id, user); err != nil {
		return nil, err
	}

	userPayments, err := self.Payments.FindAllByUserId(id)
	if err != nil {
		return nil, err
	}

	return userPayments, nil
}

func (self *Service) FindAllComments(id t.UserId, user *t.User) (t.Comments, error) {
	if _, err := self.getUser(id, user); err != nil {
		return nil, err
	}

	userComments, err := self.Comments.FindAllByUserId(id)
	if err != nil {
		return nil, err
	}

	return userComments, nil
}

func (self *Service) getUser(id t.UserId, userAuth *t.User) (t.User, error) {
	user, err := self.Users.Find(id)
	if err != nil {
		return t.User{}, err
	}

	if user.Id == userAuth.Id {
		return user, nil
	}

	return t.User{}, ErrUserUnauthorized
}
