package utils

import (
	"errors"
	"os"

	"github.com/hbagdi/go-unsplash/unsplash"
	"golang.org/x/oauth2"
)

var unsplashApp *unsplash.Unsplash

func InitPhotos() {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: os.Getenv("UNSPASH_APP_ID") +
				"" +
				os.Getenv("UNSPLASH_ACCESS_KEY"),
		},
	)

	client := oauth2.NewClient(oauth2.NoContext, ts)

	unsplashApp = unsplash.New(client)
}

func GetRandomPhoto() (string, error) {
	photos, _, err := unsplashApp.Photos.Random(nil)
	if err != nil {
		return "", err
	}

	if len(*photos) == 0 {
		return "", errors.New("no photos found")
	}

	return (*photos)[0].Urls.Regular.String(), nil
}
