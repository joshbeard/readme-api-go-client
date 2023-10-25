package mockdata

import "github.com/liveoaklabs/readme-api-go-client/readme"

var (
	APISpecification = readme.APISpecification{
		Category: readme.CategorySummary{
			ID:    "650bcc2de74f7f00257881ab",
			Order: 1,
			Slug:  "test-category",
			Title: "Test Category",
			Type:  "doc",
		},
		ID:         "3e0bcc2de74f7f00257881cd",
		LastSynced: "2021-01-01T00:00:00.000Z",
		Source:     "api",
		Title:      "Test API Specification",
		Type:       "guide",
		Version:    "8c0bcc2de74f7f00257881ef",
	}
)
