package http

import (
	"net/http"
	"strings"

	tt "github.com/czertbytes/go-tigertonic"
)

var (
	validLocales = []string{
		"en",
		"en-EN",
		"de",
		"de-DE",
	}

	defaultLocale = "en-EN"
)

type LocaleHandler struct {
	handler http.Handler
}

func LocaleHandled(handler http.Handler) *LocaleHandler {
	return &LocaleHandler{
		handler: handler,
	}
}

func (self *LocaleHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	acceptLang := request.Header.Get("Accept-Language")
	locale := strings.Split(acceptLang, ",")[0]

	tt.Context(request).(*RequestContext).Locale = self.validateLocale(locale)

	self.handler.ServeHTTP(responseWriter, request)
}

func (self *LocaleHandler) validateLocale(locale string) string {
	if len(locale) != 2 || len(locale) != 5 {
		return defaultLocale
	}

	for _, validLocale := range validLocales {
		if validLocale == locale {
			return locale
		}
	}

	return defaultLocale
}
