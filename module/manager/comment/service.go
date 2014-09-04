package comment

import (
	"fmt"

	"appengine"

	c "github.com/czertbytes/pocket/pkg/comments"
	h "github.com/czertbytes/pocket/pkg/http"
	t "github.com/czertbytes/pocket/pkg/types"
)

var (
	ErrCommentUnauthorized error = fmt.Errorf("comment: Unauthorized access")
)

type Service struct {
	AppEngineContext appengine.Context
	RequestContext   *h.RequestContext
	notificator      *Notificator
	Comments         *c.Comments
}

func NewService(RequestContext *h.RequestContext) *Service {
	return &Service{
		AppEngineContext: RequestContext.AppEngineContext,
		RequestContext:   RequestContext,
		notificator:      NewNotificator(RequestContext),
		Comments:         c.NewComments(RequestContext.AppEngineContext),
	}
}

func (self *Service) Find(id t.CommentId, user *t.User) (t.Comment, error) {
	return self.find(id, user)
}

func (self *Service) Delete(id t.CommentId, user *t.User) error {
	if _, err := self.find(id, user); err != nil {
		return err
	}

	return self.Comments.Delete(id)
}

func (self *Service) find(id t.CommentId, user *t.User) (t.Comment, error) {
	comment, err := self.Comments.Find(id)
	if err != nil {
		return t.Comment{}, err
	}

	if comment.UserId != user.Id {
		return t.Comment{}, ErrCommentUnauthorized
	}

	return comment, nil
}
