package workflows

import (
	"context"
	restClient "github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest/hooks"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
	"time"
)

// WorkflowsService provides methods to interact with WorkflowsService-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type WorkflowsService struct {
	manager *configmanager.ConfigManager
	hook    hooks.Hook
}

func NewWorkflowsService() *WorkflowsService {
	return &WorkflowsService{
		manager: configmanager.NewConfigManager(clientconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *WorkflowsService) WithConfigManager(manager *configmanager.ConfigManager) *WorkflowsService {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *WorkflowsService) WithHook(hook hooks.Hook) *WorkflowsService {
	api.hook = hook
	return api
}

func (api *WorkflowsService) getConfig() *clientconfig.Config {
	return api.manager.GetWorkflows()
}

func (api *WorkflowsService) getHook() hooks.Hook {
	return api.hook
}

func (api *WorkflowsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *WorkflowsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *WorkflowsService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Retrieves all workflow definitions for the project
func (api *WorkflowsService) FetchWorkflows(ctx context.Context) (*shared.ClientResponse[WorkflowList], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/workflows").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WorkflowList, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[[]byte](err)
	}

	return shared.NewClientResponse[WorkflowList](resp), nil
}

// Creates or updates a workflow definition for the project
func (api *WorkflowsService) SaveWorkflow(ctx context.Context, workflowDefinition WorkflowDefinition) (*shared.ClientResponse[WorkflowDefinition], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/workflows").
		WithConfig(config).
		WithBody(workflowDefinition).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WorkflowDefinition, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[[]byte](err)
	}

	return shared.NewClientResponse[WorkflowDefinition](resp), nil
}

// Retrieves a workflow definition by key
func (api *WorkflowsService) FetchWorkflow(ctx context.Context) (*shared.ClientResponse[WorkflowDefinition], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/workflows/*").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WorkflowDefinition, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[[]byte](err)
	}

	return shared.NewClientResponse[WorkflowDefinition](resp), nil
}

// Executes a workflow with the provided input parameters
func (api *WorkflowsService) CreateWorkflowRun(ctx context.Context, executeWorkflowRequest ExecuteWorkflowRequest) (*shared.ClientResponse[CreateRunResponse], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/workflows/runs").
		WithConfig(config).
		WithBody(executeWorkflowRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[CreateRunResponse, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[[]byte](err)
	}

	return shared.NewClientResponse[CreateRunResponse](resp), nil
}

// Retrieves the status and details of a workflow run
func (api *WorkflowsService) FetchWorkflowRun(ctx context.Context, runId string) (*shared.ClientResponse[GetRunResponse], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/workflows/runs/{run_id}").
		WithConfig(config).
		AddPathParam("run_id", runId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[GetRunResponse, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[[]byte](err)
	}

	return shared.NewClientResponse[GetRunResponse](resp), nil
}

// Retrieves all runs for a specific workflow
func (api *WorkflowsService) ListWorkflowRuns(ctx context.Context, workflowKey string) (*shared.ClientResponse[WorkflowRunCollection], *shared.ClientError[[]byte]) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/workflows/{workflow_key}/runs").
		WithConfig(config).
		AddPathParam("workflow_key", workflowKey).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WorkflowRunCollection, []byte](config, api.getHook())
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[[]byte](err)
	}

	return shared.NewClientResponse[WorkflowRunCollection](resp), nil
}
