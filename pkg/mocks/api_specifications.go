package mocks

import "github.com/liveoaklabs/readme-api-go-client/readme"

type MockAPISpecificationService interface {
	Create(uuid string, definition string, version ...string) (readme.APISpecificationSaved, *readme.APIResponse, error)
	Get(uuid string) (string, *readme.APIResponse, error)
}

type MockAPISpecificationClient struct {
	client *MockClient
}

func (c *MockAPISpecificationClient) Create(uuid string, definition string, version ...string) (readme.APISpecificationSaved, *readme.APIResponse, error) {
	return c.client.APISpecification.Create(uuid, definition, version...)
}

func (c *MockAPISpecificationClient) Get(uuid string) (string, *readme.APIResponse, error) {
	return c.client.APISpecification.Get(uuid)
}
