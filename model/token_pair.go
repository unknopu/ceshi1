package model

// User defines domain model and it is jason and db representations
type TokenPair struct {
	IDToken string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
}
