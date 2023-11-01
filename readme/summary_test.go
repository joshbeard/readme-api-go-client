package readme_test

import (
	"errors"
	"net/http"
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

// func Test_Summary_Get(t *testing.T) {
// 	t.Run("when API responds with 200", func(t *testing.T) {
// 		// Arrange
// 		expect := readme.Summary{
// 			Counts: readme.SummaryCounts{
// 				APISpecifications: 1,
// 				Categories:        1,
// 				Docs:              1,
// 				Changelogs:        1,
// 				CustomPages:       1,
// 				Versions:          1,
// 			},
// 			APISpecifications: []readme.SummaryAPISpecification{
// 				{
// 					ID:      "1234567890abcdef",
// 					Title:   "API Specification 1",
// 					Version: "1.0.0",
// 				},
// 				{
// 					ID:      "1234567890abcdef",
// 					Title:   "API Specification 2",
// 					Version: "1.0.0",
// 				},
// 			},
// 			Categories: []readme.SummaryCategory{
// 				{
// 					ID:      "1234567890abcdef",
// 					Slug:    "category-1",
// 					Version: "abcdef1234567890",
// 				},
// 			},
// 			Docs: []readme.SummaryDoc{
// 				{
// 					ID:      "1234567890abcdef",
// 					Slug:    "doc-1",
// 					Version: "abcdef1234567890",
// 				},
// 			},
// 			Changelogs: []readme.SummaryChangelog{
// 				{
// 					ID:   "1234567890abcdef",
// 					Slug: "changelog-1",
// 				},
// 			},
// 			CustomPages: []readme.SummaryCustomPage{
// 				{
// 					ID:   "1234567890abcdef",
// 					Slug: "custom-page-1",
// 				},
// 			},
// 			Versions: []readme.SummaryVersion{
// 				{
// 					ID:   "abcdef1234567890",
// 					Name: "1.0.0",
// 				},
// 			},
// 		}
// 		// mockResponse := testutil.APITestResponse{
// 		// 	URL:    projectEndpoint,
// 		// 	Status: 200,
// 		// 	Body: `
// 		// 		{
// 		// 			"name": "Go Testing",
// 		// 			"subdomain": "foobar",
// 		// 			"jwtSecret": "123456789abcdef",
// 		// 			"baseUrl": "https://developer.example.com",
// 		// 			"plan": "enterprise"
// 		// 		}
// 		// 	`,
// 		// }
// 		// testutil.JsonToStruct(t, mockResponse.Body, &expect)
// 		// api := mockResponse.New(t)
//
// 		// Act
// 		got, err := api.Summary()
//
// 		// Assert
// 		assert.Nil(t, err, "it does not return an error")
// 		assert.Equal(t, expect, got, "it returns expected Project struct")
// 	})
// }
