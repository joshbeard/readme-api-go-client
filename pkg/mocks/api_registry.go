package mocks

import (
	"net/http"

	"github.com/liveoaklabs/readme-api-go-client/readme"
)

type MockAPIRegistryService interface {
	Create(definition string, version ...string) (readme.APIRegistrySaved, *readme.APIResponse, error)
	Get(uuid string) (string, *readme.APIResponse, error)

	CreateWith400(definition string, version ...string) (readme.APIRegistrySaved, *readme.APIResponse, error)
}

type MockAPIRegistryClient struct {
	client *MockClient
}

func (c *MockAPIRegistryClient) Create(definition string, version ...string) (readme.APIRegistrySaved, *readme.APIResponse, error) {
	definitionMap := map[string]interface{}{}
	definitionMap["openapi"] = "3.0.0"
	definitionMap["info"] = map[string]interface{}{
		"title":   "ReadMe API",
		"version": "1.0.0",
	}

	response := readme.APIRegistrySaved{
		RegistryUUID: "43ebc7c3-4653-4037-841e-075ad428a68d",
		Definition:   definitionMap,
	}

	return response, nil, nil
}

func (c *MockAPIRegistryClient) CreateWith400(definition string, version ...string) (readme.APIRegistrySaved, *readme.APIResponse, error) {
	definitionMap := map[string]interface{}{}
	definitionMap["openapi"] = "3.0.0"
	definitionMap["info"] = map[string]interface{}{
		"title":   "ReadMe API",
		"version": "1.0.0",
	}

	response := readme.APIRegistrySaved{
		RegistryUUID: "43ebc7c3-4653-4037-841e-075ad428a68d",
		Definition:   definitionMap,
	}

	apiResponse := readme.APIResponse{
		APIErrorResponse: readme.APIErrorResponse{
			Message: "Invalid OpenAPI specification",
			Error:   "Invalid OpenAPI specification",
		},
		HTTPResponse: &http.Response{
			StatusCode: 400,
		},
		Request: &readme.APIRequest{
			Endpoint: "/registry",
			Headers: []readme.RequestHeader{
				{"Content-Type": "application/json"},
			},
			Method: "POST",
			URL:    "http://readme-test.local/api/v1/registry",
		},
	}

	return response, &apiResponse, nil
}

func (c *MockAPIRegistryClient) Get(uuid string) (string, *readme.APIResponse, error) {
	// Return sample OpenAPI spec
	response := ``

	apiResponse := readme.APIResponse{}

	return response, &apiResponse, nil
}
