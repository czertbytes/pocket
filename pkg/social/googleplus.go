package social

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"

	"appengine"
	"appengine/urlfetch"
)

const (
	GPlusProfileURL string = "https://www.googleapis.com/plus/v1/people/%s?fields=id,displayName,name,emails,image,url&access_token=%s"
)

type GooglePlus struct {
	AppEngineContext appengine.Context
	client           *http.Client
}

func NewGooglePlus(appEngineContext appengine.Context) Fetcher {
	return &GooglePlus{
		AppEngineContext: appEngineContext,
		client:           urlfetch.Client(appEngineContext),
	}
}

func (self *GooglePlus) Fetch(profileId, authToken string) (Profile, error) {
	var wg sync.WaitGroup
	errChan := make(chan error)
	quitChan := make(chan struct{})

	wg.Add(1)
	googlePlusProfileChan := self.fetchProfile(profileId, authToken, errChan, &wg)

	go func() {
		wg.Wait()
		close(quitChan)
	}()

	profile := Profile{}
	timeout := time.After(5 * time.Second)
	for {
		select {
		case googlePlusProfile := <-googlePlusProfileChan:
			profile.FirstName = googlePlusProfile.Name.GivenName
			profile.LastName = googlePlusProfile.Name.FamilyName
			profile.FullName = googlePlusProfile.DisplayName
			if len(googlePlusProfile.Emails) > 0 {
				profile.Email = googlePlusProfile.Emails[0].Value
			}
			profile.PhotoURL = googlePlusProfile.Image.Url
		case <-quitChan:
			return profile, nil
		case err := <-errChan:
			return Profile{}, err
		case <-timeout:
			return Profile{}, fmt.Errorf("Getting GPlusUser took too long!")
		}
	}
}

func (self *GooglePlus) fetchProfile(profileId, authToken string, errChan chan<- error, wg *sync.WaitGroup) chan GooglePlusProfile {
	c := make(chan GooglePlusProfile)

	go func() {
		defer func() {
			wg.Done()
		}()

		url := fmt.Sprintf(GPlusProfileURL, profileId, authToken)
		response, err := self.Do("GET", url)

		if err != nil {
			self.AppEngineContext.Debugf(err.Error())
			errChan <- err
			return
		}

		if response.StatusCode != http.StatusOK {
			body, _ := ioutil.ReadAll(response.Body)
			response.Body.Close()

			self.AppEngineContext.Debugf(string(body))

			errChan <- fmt.Errorf("GPlus auth error!")
			return
		}

		var googlePlusProfile GooglePlusProfile
		err = json.NewDecoder(response.Body).Decode(&googlePlusProfile)
		if err != nil && err != io.EOF {
			errChan <- err
			return
		}

		// change the image url to force the size as 144*144
		imageUrl, err := self.changeImageSize(googlePlusProfile, "144")
		if err != nil {
			errChan <- err
			return
		}

		googlePlusProfile.Image.Url = imageUrl

		c <- googlePlusProfile
	}()

	return c
}

func (self *GooglePlus) changeImageSize(googlePlusProfile GooglePlusProfile, newSizeStr string) (string, error) {
	newImageUrl, err := url.Parse(googlePlusProfile.Image.Url)
	if err != nil {
		return "", err
	}

	query := newImageUrl.Query()
	query.Set("sz", newSizeStr)
	newImageUrl.RawQuery = query.Encode()
	newImageUrl.String()

	return newImageUrl.String(), nil
}

func (self *GooglePlus) Do(method, url string) (*http.Response, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	return self.client.Do(request)
}

type GooglePlusProfile struct {
	Id          string                     `json:"id"`
	Url         string                     `json:"url"`
	Name        GooglePlusUserNameData     `json:"name"`
	DisplayName string                     `json:"displayName"`
	Emails      []GooglePlusEmailData      `json:"emails"`
	Image       GooglePlusProfileImageData `json:"image"`
}

type GooglePlusUserNameData struct {
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
}

type GooglePlusEmailData struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type GooglePlusProfileImageData struct {
	Url string `json:"url"`
}
