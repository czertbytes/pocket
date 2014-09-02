package social

type Profile struct {
	FullName  string
	FirstName string
	LastName  string
	Email     string
	PhotoURL  string
}

type Fetcher interface {
	Fetch(profileId, authToken string) (Profile, error)
}
