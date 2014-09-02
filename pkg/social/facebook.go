package social

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"appengine"
	"appengine/urlfetch"
)

const (
	FacebookProfileMeURL = "https://graph.facebook.com/v2.0/me?access_token=%s&appsecret_proof=%s"
	FacebookPictureURL   = "https://graph.facebook.com/%s/picture?redirect=0&height=320&width=320&type=square"
)

type Facebook struct {
	AppEngineContext appengine.Context
	client           *http.Client
	appSecret        string
}

func NewFacebook(appEngineContext appengine.Context, appSecret string) Fetcher {
	return &Facebook{
		AppEngineContext: appEngineContext,
		client:           urlfetch.Client(appEngineContext),
		appSecret:        appSecret,
	}
}

func (self *Facebook) Fetch(profileId, authToken string) (Profile, error) {
	var wg sync.WaitGroup
	errChan := make(chan error)
	quitChan := make(chan struct{})

	wg.Add(2)
	facebookProfileChan := self.fetchProfile(authToken, errChan, &wg)
	facebookProfileImageChan := self.fetchProfileImage(profileId, errChan, &wg)

	go func() {
		wg.Wait()
		close(quitChan)
	}()

	profile := Profile{}
	timeout := time.After(5 * time.Second)
	for {
		select {
		case facebookProfile := <-facebookProfileChan:
			profile.FirstName = facebookProfile.FirstName
			profile.LastName = facebookProfile.LastName
			profile.FullName = facebookProfile.FullName
			profile.Email = facebookProfile.Email
		case facebookProfileImage := <-facebookProfileImageChan:
			profile.PhotoURL = facebookProfileImage
		case <-quitChan:
			return profile, nil
		case err := <-errChan:
			return Profile{}, err
		case <-timeout:
			return Profile{}, fmt.Errorf("Getting FacebookUser took too long!")
		}
	}
}

func (self *Facebook) fetchProfile(token string, errChan chan<- error, wg *sync.WaitGroup) chan FacebookProfile {
	c := make(chan FacebookProfile)

	go func() {
		defer func() {
			wg.Done()
		}()

		url := fmt.Sprintf(FacebookProfileMeURL, token, self.appSecretProof(token))
		response, err := self.client.Get(url)
		if err != nil {
			self.AppEngineContext.Debugf(err.Error())
			errChan <- err
			return
		}

		if response.StatusCode != http.StatusOK {
			body, _ := ioutil.ReadAll(response.Body)
			response.Body.Close()

			self.AppEngineContext.Debugf(string(body))

			errChan <- fmt.Errorf("Facebook auth error!")
			return
		}

		var facebookProfile FacebookProfile
		err = json.NewDecoder(response.Body).Decode(&facebookProfile)
		if err != nil && err != io.EOF {
			errChan <- err
			return
		}

		c <- facebookProfile
	}()

	return c
}

func (self *Facebook) fetchProfileImage(profileId string, errChan chan<- error, wg *sync.WaitGroup) chan string {
	c := make(chan string)

	go func() {
		defer func() {
			wg.Done()
		}()

		url := fmt.Sprintf(FacebookPictureURL, profileId)
		response, err := self.client.Get(url)
		if err != nil {
			self.AppEngineContext.Debugf(err.Error())
			errChan <- err
			return
		}

		if response.StatusCode != http.StatusOK {
			body, _ := ioutil.ReadAll(response.Body)
			response.Body.Close()

			self.AppEngineContext.Debugf(string(body))

			errChan <- fmt.Errorf("Facebook fetch profile image error!")
			return
		}

		var facebookPicture FacebookPicture
		err = json.NewDecoder(response.Body).Decode(&facebookPicture)
		if err != nil && err != io.EOF {
			errChan <- err
			return
		}

		c <- facebookPicture.Data.URL
	}()

	return c
}

func (self *Facebook) appSecretProof(token string) string {
	mac := hmac.New(sha256.New, []byte(self.appSecret))
	mac.Write([]byte(token))

	return fmt.Sprintf("%x", mac.Sum(nil))
}

type FacebookProfile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FullName  string `json:"name"`
	Email     string `json:"email"`
	Verified  bool   `json:"verified"`
	ImageURL  string
}

type FacebookPicture struct {
	Data FacebookPictureData `json:"data"`
}

type FacebookPictureData struct {
	URL string `json:"url"`
}
