package readme

import (
	"reflect"
	"testing"
)

// Create a mockClient that implements the same methods as the actual Client.
type mockClient struct{}

func (mc *mockClient) Summary() (Summary, []error) {
	// Simulate a successful summary response.
	specs := []SummaryAPISpecification{{ID: "1", Title: "API Spec", Version: "1.0"}}
	categories := []SummaryCategorie{{ID: "1", Slug: "cat", Version: "1.0"}}
	docs := []SummaryDoc{{ID: "1", Slug: "doc", Version: "1.0"}}
	changelogs := []SummaryChangelog{{ID: "1", Slug: "change"}}
	customPages := []SummaryCustomPage{{ID: "1", Slug: "page"}}
	versions := []SummaryVersion{{ID: "1", Name: "v1.0"}}
	return Summary{
		Counts: SummaryCounts{
			APISpecifications: 1,
			Categories:        1,
			Changelogs:        1,
			CustomPages:       1,
			Docs:              1,
			Versions:          1,
		},
		APISpecifications: specs,
		Categories:        categories,
		Changelogs:        changelogs,
		CustomPages:       customPages,
		Docs:              docs,
		Versions:          versions,
	}, nil
}

func TestClient_Summary(t *testing.T) {
	// Replace the real client with the mock client for testing.
	client := &Client{}

	// Test successful summary.
	expected := Summary{
		Counts: SummaryCounts{
			APISpecifications: 1,
			Categories:        1,
			Changelogs:        1,
			CustomPages:       1,
			Docs:              1,
			Versions:          1,
		},
		APISpecifications: []SummaryAPISpecification{{ID: "1", Title: "API Spec", Version: "1.0"}},
		Categories:        []SummaryCategorie{{ID: "1", Slug: "cat", Version: "1.0"}},
		Changelogs:        []SummaryChangelog{{ID: "1", Slug: "change"}},
		CustomPages:       []SummaryCustomPage{{ID: "1", Slug: "page"}},
		Docs:              []SummaryDoc{{ID: "1", Slug: "doc", Version: "1.0"}},
		Versions:          []SummaryVersion{{ID: "1", Name: "v1.0"}},
	}
	actual, err := client.Summary()
	if err != nil {
		t.Errorf("Expected no errors, but got %d errors", len(err))
	}

	t.Logf("Expected %+v", expected)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %+v, but got %+v", expected, actual)
	}

	// Test error handling.
	// mockErr := errors.New("mock error")
	// client = mockClient{}
	// client.(*Client).mockSummaryError = mockErr
	// _, errs := client.Summary()
	// if len(errs) != 1 {
	// 	t.Errorf("Expected 1 error, but got %d errors", len(errs))
	// }
}
