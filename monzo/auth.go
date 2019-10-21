package monzo

import (
	"encoding/json"
	"log"
)

// AuthService provides an interface to auth-related endpoints.
type AuthService struct {
	Client *Client
}

// WhoAmI represents the possible fields in a /ping/whoami response body.
type WhoAmI struct {
	// Success
	Authenticated bool   `json:"authenticated"`
	ClientID      string `json:"client_id"`
	UserID        string `json:"user_id"`
	// Error
	Code             string `json:"code"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
	Message          string `json:"message"`
}

// GetStatus queries /ping/whoami.
func (s *AuthService) GetStatus() WhoAmI {
	response := s.Client.get("/ping/whoami")

	var data WhoAmI
	err := json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	return data
}
