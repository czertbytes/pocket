package http

import (
	"net/http"
	"time"

	tt "github.com/czertbytes/go-tigertonic"
)

type Event struct {
	Created      time.Time
	URL          string
	ResponseCode int
	Duration     int64
	UserAgent    string
	IPAddress    string
	Country      string
	Region       string
	City         string
	CityLatLng   string
	Locale       string
	ClientId     int64
}

type LoggedHandler struct {
	handler http.Handler
}

func LoggerHandled(handler http.Handler) *LoggedHandler {
	return &LoggedHandler{
		handler: handler,
	}
}

func (self *LoggedHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	start := time.Now()
	tee := tt.NewTeeResponseWriter(responseWriter)

	self.handler.ServeHTTP(tee, request)

	appEngineContext := tt.Context(request).(*RequestContext).AppEngineContext
	if appEngineContext == nil {
		panic("AppEngineContext is not set!")
	}

	event := createEvent(request)
	event.Created = start
	event.Duration = durationInMs(start)
	event.ResponseCode = tee.StatusCode

	appEngineContext.Infof("%+v", event)
}

func createEvent(request *http.Request) *Event {
	requestContext := tt.Context(request).(*RequestContext)

	var clientId int64
	if requestContext.Client != nil {
		clientId = int64(requestContext.Client.Id)
	}

	requestSource := requestContext.RequestSource

	event := &Event{
		URL:        request.RequestURI,
		UserAgent:  requestSource.UserAgent,
		IPAddress:  requestSource.IPAddress,
		Country:    requestSource.Country,
		Region:     requestSource.Region,
		City:       requestSource.City,
		CityLatLng: requestSource.CityLatLng,
		Locale:     requestContext.Locale,
		ClientId:   clientId,
	}

	return event
}

func durationInMs(start time.Time) int64 {
	return time.Now().Sub(start).Nanoseconds() / 1000000
}
