// This provides a mock client for use in tests to provide a mock
// implementation of the readme.Client interface. This instantiates all of the
// mock services as a convenience, rather than requiring the caller to
// instantiate each service individually.
package mocks

import (
	"testing"

	readme "github.com/liveoaklabs/readme-api-go-client/readme"
)

type MockClient struct {
	APIRegistry      *MockAPIRegistryService
	APISpecification *MockAPISpecificationService
	Apply            *MockApplyService
	Category         *MockCategoryService
	Changelog        *MockChangelogService
	CustomPage       *MockCustomPageService
	Doc              *MockDocService
	Image            *MockImageService
	Project          *MockProjectService
	Version          *MockVersionService
}

func NewReadmeClient(t *testing.T) *readme.Client {
	return &readme.Client{
		APIRegistry:      NewMockAPIRegistryService(t),
		APISpecification: NewMockAPISpecificationService(t),
		Apply:            NewMockApplyService(t),
		Category:         NewMockCategoryService(t),
		Changelog:        NewMockChangelogService(t),
		CustomPage:       NewMockCustomPageService(t),
		Doc:              NewMockDocService(t),
		Image:            NewMockImageService(t),
		Project:          NewMockProjectService(t),
		Version:          NewMockVersionService(t),
	}
}

func NewTestClient(t *testing.T) *MockClient {
	mockClient := &MockClient{}

	mockClient.APIRegistry = NewMockAPIRegistryService(t)
	mockClient.APISpecification = NewMockAPISpecificationService(t)
	mockClient.Apply = NewMockApplyService(t)
	mockClient.Category = NewMockCategoryService(t)
	mockClient.Changelog = NewMockChangelogService(t)
	mockClient.CustomPage = NewMockCustomPageService(t)
	mockClient.Doc = NewMockDocService(t)
	mockClient.Image = NewMockImageService(t)
	mockClient.Project = NewMockProjectService(t)
	mockClient.Version = NewMockVersionService(t)

	return mockClient
}