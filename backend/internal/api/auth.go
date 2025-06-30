package api

import (
	"net/http"

	"BeatHub-Backend/internal/schemas"
)

// Authenticate the client to make calls to osu api v2
func GetClientAuth(clientID, clientSecret string) (schemas.ClientAuth, error) {
	return Fetch[schemas.ClientAuth](FetchOptions{
		Url:    "https://osu.ppy.sh/oauth/token",
		Method: http.MethodPost,
		Body: map[string]any{
			"client_id":     clientID,
			"client_secret": clientSecret,
			"grant_type":    "client_credentials",
			"scope":         "public",
		},
	})
}
