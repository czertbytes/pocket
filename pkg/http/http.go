package http

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"

	"appengine"

	tt "github.com/czertbytes/go-tigertonic"
	t "github.com/czertbytes/pocket/pkg/types"
)

var (
	CORS *tt.CORSBuilder
)

func init() {
	CORS = tt.NewCORSBuilder().AddAllowedOrigins("*")
	for _, header := range []string{"Content-Type", "Accept", "Accept-Language", "Authorization"} {
		CORS.AddAllowedHeaders(header)
	}
}

type RequestContext struct {
	AppEngineContext appengine.Context // add by AppEngineHandler
	Locale           string            // add by LocaleHandler
	EntityId         int64             // add by EntityIdHandler
	RequestSource    *RequestSource    // add by RequestSourceHandler
	ClientId         t.ClientClientId  // add by AuthHandler
	Token            string            // add by AuthHandler
	Client           *t.Client         // add by AuthHandler
	User             *t.User           // add by UserHandler
	MultipartReader  *multipart.Reader // add by UploadHandler
}

type RequestSource struct {
	UserAgent  string
	IPAddress  string
	Country    string
	Region     string
	City       string
	CityLatLng string
}

func RequestParamInt64(url *url.URL, key string) (int64, error) {
	valueStr := url.Query().Get(key)
	value, err := strconv.Atoi(valueStr)

	return int64(value), err
}

func RequestParamFloat64(url *url.URL, key string) (float64, error) {
	valueStr := url.Query().Get(key)
	value, err := strconv.ParseFloat(valueStr, 64)

	return float64(value), err
}

func RequestParamString(url *url.URL, key string) string {
	return url.Query().Get(key)
}

func PocketHandler(handler http.Handler) http.Handler {
	return CORS.Build(
		AppEngineHandled(
			RequestSourceHandled(
				LoggerHandled(
					LocaleHandled(
						EntityIdHandled(
							handler))))))
}

func UnauthorizedResponsePayload(appEngineContext appengine.Context, responseWriter http.ResponseWriter, internalError error) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusUnauthorized)

	if err := json.NewEncoder(responseWriter).Encode(map[string]string{
		"error":          "Request authentication is not valid!",
		"internal_error": internalError.Error(),
	}); err != nil {
		appEngineContext.Errorf(err.Error())
	}
}

func EntityIdNotValidResponsePayload(appEngineContext appengine.Context, responseWriter http.ResponseWriter, internalError error) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusBadRequest)

	if err := json.NewEncoder(responseWriter).Encode(map[string]string{
		"error":          "Request entityId is not valid!",
		"internal_error": internalError.Error(),
	}); err != nil {
		appEngineContext.Errorf(err.Error())
	}
}
