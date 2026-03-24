package integrations

import "encoding/json"

type AuthedUser struct {
	// User token returned from the OAuth exchange.
	AccessToken *string `json:"access_token,omitempty"`
	// Seconds until the user token expires.
	ExpiresIn *int64 `json:"expires_in,omitempty"`
	// Slack user ID for the installer.
	Id *string `json:"id,omitempty" required:"true"`
	// Refresh token for the authed user.
	RefreshToken *string `json:"refresh_token,omitempty"`
	// Space-delimited OAuth scopes granted to the user token.
	Scope *string `json:"scope,omitempty"`
	// Token type value provided by Slack.
	TokenType *string `json:"token_type,omitempty"`
}

func (a *AuthedUser) GetAccessToken() *string {
	if a == nil {
		return nil
	}
	return a.AccessToken
}

func (a *AuthedUser) SetAccessToken(accessToken string) {
	a.AccessToken = &accessToken
}

func (a *AuthedUser) GetExpiresIn() *int64 {
	if a == nil {
		return nil
	}
	return a.ExpiresIn
}

func (a *AuthedUser) SetExpiresIn(expiresIn int64) {
	a.ExpiresIn = &expiresIn
}

func (a *AuthedUser) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *AuthedUser) SetId(id string) {
	a.Id = &id
}

func (a *AuthedUser) GetRefreshToken() *string {
	if a == nil {
		return nil
	}
	return a.RefreshToken
}

func (a *AuthedUser) SetRefreshToken(refreshToken string) {
	a.RefreshToken = &refreshToken
}

func (a *AuthedUser) GetScope() *string {
	if a == nil {
		return nil
	}
	return a.Scope
}

func (a *AuthedUser) SetScope(scope string) {
	a.Scope = &scope
}

func (a *AuthedUser) GetTokenType() *string {
	if a == nil {
		return nil
	}
	return a.TokenType
}

func (a *AuthedUser) SetTokenType(tokenType string) {
	a.TokenType = &tokenType
}

func (a AuthedUser) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AuthedUser to string"
	}
	return string(jsonData)
}
