package overview

import (
	"fmt"

	"appengine"

	h "github.com/czertbytes/pocket/pkg/http"
	o "github.com/czertbytes/pocket/pkg/overviews"
	p "github.com/czertbytes/pocket/pkg/payments"
	t "github.com/czertbytes/pocket/pkg/types"
	u "github.com/czertbytes/pocket/pkg/users"
	uo "github.com/czertbytes/pocket/pkg/users/overviews"
)

var (
	ErrOverviewUnauthorized error = fmt.Errorf("overview: Unauthorized access")
)

type Service struct {
	AppEngineContext appengine.Context
	RequestContext   *h.RequestContext
	notificator      *Notificator
	Overviews        *o.Overviews
	Users            *u.Users
	UserOverviews    *uo.Overviews
	Payments         *p.Payments
}

func NewService(RequestContext *h.RequestContext) *Service {
	return &Service{
		AppEngineContext: RequestContext.AppEngineContext,
		RequestContext:   RequestContext,
		notificator:      NewNotificator(RequestContext),
		Overviews:        o.NewOverviews(RequestContext.AppEngineContext),
		Users:            u.NewUsers(RequestContext.AppEngineContext),
		UserOverviews:    uo.NewOverviews(RequestContext.AppEngineContext),
		Payments:         p.NewPayments(RequestContext.AppEngineContext),
	}
}

func (self *Service) Create(overview *t.Overview, user *t.User) error {
	overview.OwnerId = user.Id
	overview.Owner = *user
	overview.URLToken = t.Hash(
		"8b2d80510b6482f95ae98449e7251c1b",
		overview.Name,
		overview.Description,
		overview.OwnerId,
	)

	if err := self.Overviews.Create(overview); err != nil {
		return err
	}

	if err := self.notificator.Create(overview); err != nil {
		return err
	}

	return nil
}

func (self *Service) Find(id t.OverviewId, user *t.User) (t.Overview, error) {
	return self.getOverview(id, user)
}

func (self *Service) Update(overview *t.Overview, user *t.User) error {
	oldOverview, err := self.getOverview(overview.Id, user)
	if err != nil {
		return err
	}

	oldOverview.Status = overview.Status
	oldOverview.Name = overview.Name
	oldOverview.Description = overview.Description

	if err := self.Update(&oldOverview, user); err != nil {
		return err
	}

	overview = &oldOverview

	if err := self.notificator.Update(overview); err != nil {
		return err
	}

	return nil
}

func (self *Service) Patch(overview *t.Overview, fields []string, user *t.User) error {
	oldOverview, err := self.getOverview(overview.Id, user)
	if err != nil {
		return err
	}

	for _, field := range fields {
		switch field {
		case "status":
			oldOverview.Status = overview.Status
		case "name":
			oldOverview.Name = overview.Name
		case "description":
			oldOverview.Description = overview.Description
		}
	}

	if err := self.Update(&oldOverview, user); err != nil {
		return err
	}

	overview = &oldOverview

	return nil
}

func (self *Service) Delete(id t.OverviewId, user *t.User) error {
	if _, err := self.getOverview(id, user); err != nil {
		return err
	}

	if err := self.Overviews.Delete(id); err != nil {
		return err
	}

	return nil
}

func (self *Service) CreatePayment(payment *t.Payment, id t.OverviewId, user *t.User) error {
	overview, err := self.getOverview(id, user)
	if err != nil {
		return err
	}

	payment.FromId = user.Id
	payment.From = *user
	payment.OverviewId = overview.Id
	if err := self.Payments.Create(payment); err != nil {
		return err
	}

	if err := self.notificator.CreatePayment(payment); err != nil {
		return err
	}

	return nil
}

func (self *Service) FindAllPayments(id t.OverviewId, user *t.User) (t.Payments, error) {
	if _, err := self.getOverview(id, user); err != nil {
		return nil, err
	}

	overviewPayments, err := self.Payments.FindAllByOverviewId(id)
	if err != nil {
		return nil, err
	}

	return overviewPayments, nil
}

func (self *Service) CreateParticipant(participant *t.User, id t.OverviewId, user *t.User) error {
	if _, err := self.getOverview(id, user); err != nil {
		return err
	}

	if len(participant.Email) > 0 {
		user, err := self.Users.FindByEmail(participant.Email)
		if err != nil {
			if err != u.ErrUserNotFound {
				return err
			}

			if err := self.Users.Create(participant); err != nil {
				return err
			}

			user = *participant
		}

		userOverview := &t.UserOverview{
			Status:     t.UserOverviewStatusActive,
			OverviewId: id,
			UserId:     user.Id,
		}
		if err := self.UserOverviews.Create(userOverview); err != nil {
			return err
		}

		if err := self.notificator.CreateParticipant(&user); err != nil {
			return err
		}
	}

	return nil
}

func (self *Service) FindAllParticipants(id t.OverviewId, user *t.User) (t.Users, error) {
	if _, err := self.getOverview(id, user); err != nil {
		return nil, err
	}

	userOverviews, err := self.UserOverviews.FindAllByOverviewId(id)
	if err != nil {
		return nil, err
	}

	participantIds := make(t.UserIds, len(userOverviews))
	for i, userOverview := range userOverviews {
		participantIds[i] = userOverview.UserId
	}

	overviewParticipants, err := self.Users.FindMulti(participantIds)
	if err != nil {
		return nil, err
	}

	overviewParticipants = append(overviewParticipants, *user)

	return overviewParticipants, nil
}

func (self *Service) getOverview(id t.OverviewId, user *t.User) (t.Overview, error) {
	overview, err := self.Overviews.Find(id)
	if err != nil {
		return t.Overview{}, err
	}

	if overview.OwnerId == user.Id {
		return overview, nil
	}

	userOverviews, err := self.UserOverviews.FindAllByOverviewId(id)
	if err != nil {
		return t.Overview{}, err
	}

	participantIds := make(t.UserIds, len(userOverviews))
	for i, userOverview := range userOverviews {
		participantIds[i] = userOverview.UserId
	}

	overviewParticipants, err := self.Users.FindMulti(participantIds)
	if err != nil {
		return t.Overview{}, err
	}

	for _, participant := range overviewParticipants {
		if participant.Id == user.Id {
			return overview, nil
		}
	}

	return t.Overview{}, ErrOverviewUnauthorized
}
