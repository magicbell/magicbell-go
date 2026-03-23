package integrations

import "encoding/json"

type SlackInstallation struct {
	// Bot token returned from the Slack OAuth exchange.
	AccessToken *string `json:"access_token,omitempty" required:"true"`
	// Slack app identifier for the installed app.
	AppId      *string     `json:"app_id,omitempty" required:"true"`
	AuthedUser *AuthedUser `json:"authed_user,omitempty" required:"true"`
	// Slack user ID of the installed bot.
	BotUserId  *string     `json:"bot_user_id,omitempty"`
	Enterprise *Enterprise `json:"enterprise,omitempty"`
	// Seconds until the bot access token expires.
	ExpiresIn *int64 `json:"expires_in,omitempty"`
	// Unique identifier MagicBell assigns to the Slack installation.
	Id              *string          `json:"id,omitempty" pattern:"^[A-Z0-9]+-.*$"`
	IncomingWebhook *IncomingWebhook `json:"incoming_webhook,omitempty"`
	// Indicates whether the installation occurred on an enterprise grid.
	IsEnterpriseInstall *bool `json:"is_enterprise_install,omitempty"`
	// Refresh token for regenerating the bot access token.
	RefreshToken *string `json:"refresh_token,omitempty"`
	// Space-delimited OAuth scopes granted to the bot token.
	Scope *string `json:"scope,omitempty"`
	Team  *Team   `json:"team,omitempty" required:"true"`
	// Type of bot token returned by Slack.
	TokenType *string `json:"token_type,omitempty"`
}

func (s *SlackInstallation) GetAccessToken() *string {
	if s == nil {
		return nil
	}
	return s.AccessToken
}

func (s *SlackInstallation) SetAccessToken(accessToken string) {
	s.AccessToken = &accessToken
}

func (s *SlackInstallation) GetAppId() *string {
	if s == nil {
		return nil
	}
	return s.AppId
}

func (s *SlackInstallation) SetAppId(appId string) {
	s.AppId = &appId
}

func (s *SlackInstallation) GetAuthedUser() *AuthedUser {
	if s == nil {
		return nil
	}
	return s.AuthedUser
}

func (s *SlackInstallation) SetAuthedUser(authedUser AuthedUser) {
	s.AuthedUser = &authedUser
}

func (s *SlackInstallation) GetBotUserId() *string {
	if s == nil {
		return nil
	}
	return s.BotUserId
}

func (s *SlackInstallation) SetBotUserId(botUserId string) {
	s.BotUserId = &botUserId
}

func (s *SlackInstallation) GetEnterprise() *Enterprise {
	if s == nil {
		return nil
	}
	return s.Enterprise
}

func (s *SlackInstallation) SetEnterprise(enterprise Enterprise) {
	s.Enterprise = &enterprise
}

func (s *SlackInstallation) GetExpiresIn() *int64 {
	if s == nil {
		return nil
	}
	return s.ExpiresIn
}

func (s *SlackInstallation) SetExpiresIn(expiresIn int64) {
	s.ExpiresIn = &expiresIn
}

func (s *SlackInstallation) GetId() *string {
	if s == nil {
		return nil
	}
	return s.Id
}

func (s *SlackInstallation) SetId(id string) {
	s.Id = &id
}

func (s *SlackInstallation) GetIncomingWebhook() *IncomingWebhook {
	if s == nil {
		return nil
	}
	return s.IncomingWebhook
}

func (s *SlackInstallation) SetIncomingWebhook(incomingWebhook IncomingWebhook) {
	s.IncomingWebhook = &incomingWebhook
}

func (s *SlackInstallation) GetIsEnterpriseInstall() *bool {
	if s == nil {
		return nil
	}
	return s.IsEnterpriseInstall
}

func (s *SlackInstallation) SetIsEnterpriseInstall(isEnterpriseInstall bool) {
	s.IsEnterpriseInstall = &isEnterpriseInstall
}

func (s *SlackInstallation) GetRefreshToken() *string {
	if s == nil {
		return nil
	}
	return s.RefreshToken
}

func (s *SlackInstallation) SetRefreshToken(refreshToken string) {
	s.RefreshToken = &refreshToken
}

func (s *SlackInstallation) GetScope() *string {
	if s == nil {
		return nil
	}
	return s.Scope
}

func (s *SlackInstallation) SetScope(scope string) {
	s.Scope = &scope
}

func (s *SlackInstallation) GetTeam() *Team {
	if s == nil {
		return nil
	}
	return s.Team
}

func (s *SlackInstallation) SetTeam(team Team) {
	s.Team = &team
}

func (s *SlackInstallation) GetTokenType() *string {
	if s == nil {
		return nil
	}
	return s.TokenType
}

func (s *SlackInstallation) SetTokenType(tokenType string) {
	s.TokenType = &tokenType
}

func (s SlackInstallation) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackInstallation to string"
	}
	return string(jsonData)
}
