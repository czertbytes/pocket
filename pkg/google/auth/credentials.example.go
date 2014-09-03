package auth

const (
	ProjectId          string = ""
	ProjectClientId    string = ""
	ProjectClientEmail string = ""
)

const (
	AuthScope    string = "https://www.googleapis.com/auth/devstorage.read_write"
	AuthTokenURI string = "https://accounts.google.com/o/oauth2/token"
)

var (
	AuthPrivateKey = []byte(``)
)
