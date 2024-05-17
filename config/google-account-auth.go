package config

import (
	"fmt"
	"os"
)

func newGoogleAccountAuth() *GoogleAccountAuth {
	return &GoogleAccountAuth{
		GoogleAuthAfterRedirectUI: os.Getenv("GOOGLE_AUTH_AFTER_REDIRECT_UI"),
	}
}

type GoogleAccountAuth struct {
	GoogleAuthAfterRedirectUI string
}

func (c *GoogleAccountAuth) CreateGoogleAuthAfterRedirectURL(
	token string,
) string {
	return fmt.Sprintf("%s?token=%s", c.GoogleAuthAfterRedirectUI, token)
}
