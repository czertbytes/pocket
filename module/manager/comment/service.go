package comment

import (
	"fmt"

	"appengine"

	c "github.com/czertbytes/pocket/pkg/comments"
	shttp "github.com/czertbytes/pocket/pkg/http"
	t "github.com/czertbytes/pocket/pkg/types"
)

type Service struct {
	AppEngineContext appengine.Context
	RequestContext   *shttp.RequestContext
	notificator      *Notificator
	Comments         *c.Comments
}

func NewService(RequestContext *shttp.RequestContext) *Service {
	return &Service{
		AppEngineContext: RequestContext.AppEngineContext,
		RequestContext:   RequestContext,
		notificator:      NewNotificator(RequestContext),
		Comments:         c.NewComments(RequestContext.AppEngineContext),
	}
}

func (self *Service) Find(id t.CommentId, user *t.User) (*t.Comment, error) {
	return self.find(id, user)
}

func (self *Service) Delete(id t.CommentId, user *t.User) error {
	comment, err := self.find(id, user)
	if err != nil {
		return nil, err
	}

	return self.Comments.Delete(id)
}

func (self *Service) find(id t.CommentId, user *t.User) (*t.Comment, error) {
	comment, err := self.Comments.Find(id)
	if err != nil {
		return nil, err
	}

	if comment.UserId != user.Id {
		return nil, fmt.Errorf("Comment is not your!")
	}

	return comment, nil
}
