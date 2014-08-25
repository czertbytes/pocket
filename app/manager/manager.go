package manager

import (
	"log"
	"net/http"

	tt "github.com/czertbytes/go-tigertonic"
	"github.com/czertbytes/pocket/module/manager/auth"
	"github.com/czertbytes/pocket/module/manager/comment"
	"github.com/czertbytes/pocket/module/manager/document"
	"github.com/czertbytes/pocket/module/manager/overview"
	"github.com/czertbytes/pocket/module/manager/payment"
	"github.com/czertbytes/pocket/module/manager/user"
	shttp "github.com/czertbytes/pocket/pkg/http"
)

func init() {
	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile)

	http.Handle("/", tt.WithContext(Mux(), shttp.RequestContext{}))
}

func Mux() *tt.TrieServeMux {
	mux := tt.NewTrieServeMux()

	// Auth
	mux.Handle("POST", "/manager/auth", ManagerPublicHandler(auth.Post))
	mux.Handle("DELETE", "/manager/auth", ManagerHandler(auth.Delete))

	// Comment
	mux.Handle("GET", "/manager/comments/{id}", ManagerHandler(comment.Get))
	mux.Handle("DELETE", "/manager/comments/{id}", ManagerHandler(comment.Delete))

	// Document
	mux.Handle("GET", "/manager/documents/{id}", ManagerHandler(document.Get))
	mux.Handle("DELETE", "/manager/documents/{id}", ManagerHandler(document.Delete))

	// Overviews
	mux.Handle("POST", "/manager/overviews", ManagerHandler(overview.Post))
	mux.Handle("GET", "/manager/overviews/{id}", ManagerHandler(overview.Get))
	mux.Handle("PUT", "/manager/overviews/{id}", ManagerHandler(overview.Put))
	mux.Handle("PATCH", "/manager/overviews/{id}", ManagerHandler(overview.Patch))
	mux.Handle("DELETE", "/manager/overviews/{id}", ManagerHandler(overview.Delete))
	mux.Handle("GET", "/manager/overviews/{id}/payments", ManagerHandler(overview.GetPayments))
	mux.Handle("POST", "/manager/overviews/{id}/participants", ManagerHandler(overview.PostParticipants))
	mux.Handle("GET", "/manager/overviews/{id}/participants", ManagerHandler(overview.GetParticipants))

	// Payment
	mux.Handle("POST", "/manager/payments", ManagerHandler(payment.Post))
	mux.Handle("GET", "/manager/payments/{id}", ManagerHandler(payment.Get))
	mux.Handle("PUT", "/manager/payments/{id}", ManagerHandler(payment.Put))
	mux.Handle("PATCH", "/manager/payments/{id}", ManagerHandler(payment.Patch))
	mux.Handle("DELETE", "/manager/payments/{id}", ManagerHandler(payment.Delete))
	mux.Handle("POST", "/manager/payments/{id}/documents", ManagerUploadHandler(payment.PostDocuments))
	mux.Handle("GET", "/manager/payments/{id}/documents", ManagerHandler(payment.GetDocuments))
	mux.Handle("POST", "/manager/payments/{id}/comments", ManagerHandler(payment.PostComments))
	mux.Handle("GET", "/manager/payments/{id}/comments", ManagerHandler(payment.GetComments))

	// User
	mux.Handle("GET", "/manager/users/{id}", ManagerHandler(user.Get))
	mux.Handle("DELETE", "/manager/users{id}", ManagerHandler(user.Delete))
	mux.Handle("GET", "/manager/users/{id}/overviews", ManagerHandler(user.GetOverviews))
	mux.Handle("GET", "/manager/users/{id}/payments", ManagerHandler(user.GetPayments))
	mux.Handle("GET", "/manager/users/{id}/comments", ManagerHandler(user.GetComments))
	mux.Handle("GET", "/manager/users/me", ManagerMeHandler(user.GetMe))
	mux.Handle("GET", "/manager/users/me/overviews", ManagerMeHandler(user.GetMeOverviews))
	mux.Handle("GET", "/manager/users/me/payments", ManagerMeHandler(user.GetMePayments))
	mux.Handle("GET", "/manager/users/me/comments", ManagerMeHandler(user.GetMeComments))

	return mux
}

func ManagerHandler(handler interface{}) http.Handler {
	return shttp.AppEngineHandled(tt.Marshaled(handler))
}

func ManagerMeHandler(handler interface{}) http.Handler {
	return shttp.AppEngineHandled(tt.Marshaled(handler))
}

func ShiftManagerUploadHandler(handler interface{}) http.Handler {
	return shttp.AppEngineHandled(tt.Marshaled(handler))
}

func ManagerPublicHandler(handler interface{}) http.Handler {
	return shttp.AppEngineHandled(tt.Marshaled(handler))
}
