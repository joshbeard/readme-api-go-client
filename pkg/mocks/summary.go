package mocks

import (
	"github.com/liveoaklabs/readme-api-go-client/pkg/mocks/mockdata"
	"github.com/liveoaklabs/readme-api-go-client/readme"
)

type MockSummaryService interface {
	Get() (string, *readme.APIResponse, error)
}

type MockSummaryClient struct {
	client *MockClient
}

func (c *MockSummaryClient) Get() (readme.Summary, []error) {
	return mockdata.Summary, nil
}
