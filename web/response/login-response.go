package response

type LoginResponseDTO struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	AccessToken string `json:"access_token"`
}
