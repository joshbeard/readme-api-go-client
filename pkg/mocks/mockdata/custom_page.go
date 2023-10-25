package mockdata

import "github.com/liveoaklabs/readme-api-go-client/readme"

var (
	CustomPage = readme.CustomPage{
		Metadata: readme.DocMetadata{
			Image:       []any{},
			Title:       "",
			Description: "",
		},
		Title:      "Some Test Page",
		Slug:       "some-test-page",
		Body:       "# A test page.",
		HTML:       "",
		HTMLMode:   false,
		Fullscreen: false,
		Hidden:     true,
		Revision:   2,
		ID:         "63ae563ec4d05f018b26cf18",
		CreatedAt:  "2022-12-30T03:08:46.695Z",
		UpdatedAt:  "2022-12-30T03:09:46.695Z",
	}
)

// mockCustomPageBody represents the raw JSON response body for a custom page from the ReadMe.com API.
// var CustomPageString = `
// 	{
// 		"metadata": {
// 		"image": [],
// 		"title": "",
// 		"description": ""
// 		},
// 		"title": "Some Test Page",
// 		"slug": "some-test-page",
// 		"body": "# A test page.",
// 		"html": "",
// 		"htmlmode": false,
// 		"fullscreen": false,
// 		"hidden": true,
// 		"pendingAlgoliaPublish": false,
// 		"revision": 2,
// 		"_id": "63ae563ec4d05f018b26cf18",
// 		"createdAt": "2022-12-30T03:08:46.695Z",
// 		"updatedAt": "2022-12-30T03:09:46.695Z",
// 		"__v": 0
// 	}
// 	`
