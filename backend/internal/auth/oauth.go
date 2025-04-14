package auth

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

var GoogleOAuthConfig *oauth2.Config

func InitOAuthConfig() {
	GoogleOAuthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}
