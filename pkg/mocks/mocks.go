package mocks

import "net/http"

type MockClient struct {
	APIURL     string
	HTTPClient *http.Client
	Token      string

	APIRegistry      MockAPIRegistryService
	APISpecification MockAPISpecificationService
	Apply            MockApplyService
	// Category         MockCategoryService
	// Changelog        MockChangelogService
	// CustomPage       MockCustomPageService
	// Doc              MockDocService
	// Image            MockImageService
	// Project          MockProjectService
	// Version          MockVersionService
	// Nuke             MockNukeService
}

func NewMockClient() *MockClient {
	return &MockClient{
		APIRegistry:      &MockAPIRegistryClient{},
		APISpecification: &MockAPISpecificationClient{},
		Apply:            &MockApplyClient{},
		// Category:         &MockCategoryClient{},
		// Changelog:        &MockChangelogClient{},
		// CustomPage:       &MockCustomPageClient{},
		// Doc:              &MockDocClient{},
		// Image:            &MockImageClient{},
		// Project:          &MockProjectClient{},
		// Version:          &MockVersionClient{},
		// Nuke:             &MockNukeClient{},
	}
}
