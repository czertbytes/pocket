package types

type AuthOriginService int8

func ParseAuthOriginService(value int) AuthOriginService {
	switch value {
	case 1:
		return AuthOriginServiceGooglePlus
	case 2:
		return AuthOriginServiceFacebook
	default:
		return AuthOriginServiceUnknown
	}
}

func (self AuthOriginService) String() string {
	switch self {
	case AuthOriginServiceGooglePlus:
		return "google_plus"
	case AuthOriginServiceFacebook:
		return "facebook"
	default:
		return "unknown"
	}
}

var (
	AuthOriginServiceUnknown    AuthOriginService = 0
	AuthOriginServiceGooglePlus AuthOriginService = 1
	AuthOriginServiceFacebook   AuthOriginService = 2
)

type AuthOrigin struct {
	Service          AuthOriginService `json:"auth_origin_service" datastore:"-"`
	ServiceFormatted string            `json:"auth_origin_service_formatted" datastore:"-"`
	EntityId         string            `json:"auth_origin_entity_id" datastore:"-"`
	Token            string            `json:"auth_origin_token" datastore:"-"`
}

func (self *AuthOrigin) SetFormattedValues() {
	self.SetServiceFormatted()
}

func (self *AuthOrigin) SetServiceFormatted() {
	self.ServiceFormatted = self.Service.String()
}
