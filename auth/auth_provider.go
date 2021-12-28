package auth

// New Auth Client
func NewAuthClient() IAuthClient {
	return &authClient{}
}

// Auth Keys
type AuthKeys struct {
	ApiKey    string
	SecretKey string
}

// Auth Client
type IAuthClient interface {
	New(*string, *string) *AuthKeys
}

// Auth struct
type authClient struct{}

// Provider
func (a *authClient) New(apiKey, secretKey *string) *AuthKeys {
	return &AuthKeys{
		ApiKey:    *apiKey,
		SecretKey: *secretKey,
	}
}