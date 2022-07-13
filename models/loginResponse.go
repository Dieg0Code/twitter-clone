package models

// LoginResponse is the response for the login endpoint
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
