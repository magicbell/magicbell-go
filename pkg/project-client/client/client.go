package client

import (
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest/hooks"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/project-client/broadcasts"
	"github.com/magicbell/magicbell-go/pkg/project-client/channels"
	"github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
	"github.com/magicbell/magicbell-go/pkg/project-client/events"
	"github.com/magicbell/magicbell-go/pkg/project-client/integrations"
	"github.com/magicbell/magicbell-go/pkg/project-client/users"
	"github.com/magicbell/magicbell-go/pkg/project-client/workflows"
	"time"
)

// Client is the main SDK client that provides access to all service endpoints.
// It manages configuration, authentication, and service instances with centralized settings.
type Client struct {
	Broadcasts   *broadcasts.BroadcastsService
	Channels     *channels.ChannelsService
	Events       *events.EventsService
	Integrations *integrations.IntegrationsService
	Users        *users.UsersService
	Workflows    *workflows.WorkflowsService
	manager      *configmanager.ConfigManager
}

func NewClient(config clientconfig.Config) *Client {
	broadcasts := broadcasts.NewBroadcastsService()
	channels := channels.NewChannelsService()
	events := events.NewEventsService()
	integrations := integrations.NewIntegrationsService()
	users := users.NewUsersService()
	workflows := workflows.NewWorkflowsService()

	manager := configmanager.NewConfigManager(config)
	hook := hooks.NewDefaultHook()
	broadcasts.WithConfigManager(manager)
	channels.WithConfigManager(manager)
	events.WithConfigManager(manager)
	integrations.WithConfigManager(manager)
	users.WithConfigManager(manager)
	workflows.WithConfigManager(manager)
	broadcasts.WithHook(hook)
	channels.WithHook(hook)
	events.WithHook(hook)
	integrations.WithHook(hook)
	users.WithHook(hook)
	workflows.WithHook(hook)

	return &Client{
		Broadcasts:   broadcasts,
		Channels:     channels,
		Events:       events,
		Integrations: integrations,
		Users:        users,
		Workflows:    workflows,
		manager:      manager,
	}
}

func (c *Client) SetBaseUrl(baseUrl string) {
	c.manager.SetBaseUrl(baseUrl)
}

func (c *Client) SetTimeout(timeout time.Duration) {
	c.manager.SetTimeout(timeout)
}

func (c *Client) SetAccessToken(accessToken string) {
	c.manager.SetAccessToken(accessToken)
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
