package http

import (
	"mime/multipart"
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

type ShiftRequestContext struct {
	AppEngineContext appengine.Context // add by AppEngineHandler
	Locale           string            // add by LocaleHandler
	EntityId         int64             // add by EntityIdHandler
	RequestSource    *RequestSource    // add by RequestSourceHandler
	Client           *t.Client         // add by AuthHandler
	ClientId         string            // add by AuthHandler
	Token            string            // add by AuthHandler
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
