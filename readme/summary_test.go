package readme_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/liveoaklabs/readme-api-go-client/internal/testutil"
	"github.com/liveoaklabs/readme-api-go-client/readme"
)

type MockHTTPClient struct {
	Responses []*http.Response
	Errors    []error
	Index     int
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if m.Index >= len(m.Responses) {
		return nil, errors.New("Exceeded expected number of requests")
	}

	resp := m.Responses[m.Index]
	err := m.Errors[m.Index]
	m.Index++

	return resp, err
}

func Test_Summary(t *testing.T) {
	t.Run("when API responds with 200", func(t *testing.T) {
		// Arrange
		expect := readme.Summary{
			Counts: readme.SummaryCounts{
				APISpecifications: 1,
				Categories:        1,
				Docs:              1,
				Changelogs:        1,
				CustomPages:       1,
				Versions:          1,
			},
			APISpecifications: []readme.SummaryAPISpecification{
				{
					ID:      "1234567890abcdef",
					Title:   "API Specification 1",
					Version: "1.0.0",
				},
				{
					ID:      "1234567890abcdef",
					Title:   "API Specification 2",
					Version: "1.0.0",
				},
			},
			Categories: []readme.SummaryCategory{
				{
					ID:      "1234567890abcdef",
					Slug:    "category-1",
					Version: "abcdef1234567890",
				},
			},
			Docs: []readme.SummaryDoc{
				{
					ID:      "1234567890abcdef",
					Slug:    "doc-1",
					Version: "abcdef1234567890",
				},
			},
			Changelogs: []readme.SummaryChangelog{
				{
					ID:   "1234567890abcdef",
					Slug: "changelog-1",
				},
			},
			CustomPages: []readme.SummaryCustomPage{
				{
					ID:   "1234567890abcdef",
					Slug: "custom-page-1",
				},
			},
			Versions: []readme.SummaryVersion{
				{
					ID:   "1234567890abcdef",
					Name: "1.0.0",
				},
			},
		}

		mockClient := &MockHTTPClient{
			Responses: []*http.Response{
				{
					StatusCode: 200,
					Body: ioutil.NopCloser(bytes.NewBufferString(
						testutil.StructToJson(t, mockVersionSummary))),
				},
			},
			Errors: []error{nil, nil},
		}

		client := readme.NewClient(mockClient, "test-api-key")

		// Act
		summary, errs := client.Summary()

		// Assert
		if len(errs) > 0 {
			t.Fatalf("Expected no errors, got %d", len(errs))
		}

		if summary.Counts != expect.Counts {
			t.Fatalf("Expected Counts to be %v, got %v", expect.Counts, summary.Counts)
		}

		if len(summary.APISpecifications) != len(expect.APISpecifications) {
			t.Fatalf("Expected %d API Specifications, got %d",
				len(expect.APISpecifications), len(summary.APISpecifications))
		}

	})
}
