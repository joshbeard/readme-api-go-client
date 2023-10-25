package mockdata

import "github.com/liveoaklabs/readme-api-go-client/readme"

var Changelog = readme.Changelog{
	Algolia: readme.DocAlgolia{
		PublishPending: false,
		RecordCount:    0,
		UpdatedAt:      "2021-01-01T00:00:00.000Z",
	},
	Body:      "A turtle has been here.",
	CreatedAt: "2021-01-01T00:00:00.000Z",
	HTML:      "<div class=\"magic-block-textarea\"><p>A turtle has been here.</p></div>",
	Hidden:    false,
	ID:        "63b37e8b65fd5b0057af23f1",
	Metadata: readme.DocMetadata{
		Description: "",
		Image:       []interface{}{},
		Title:       "",
	},
	Project:   "63a77777f52b9f006b6bf212",
	Revision:  1,
	Slug:      "test-changelog",
	Title:     "Test Changelog",
	Type:      "changelog",
	UpdatedAt: "2021-01-01T00:00:00.000Z",
	User: readme.DocUser{
		ID: "43a77777f52b9f006b6bf1f6",
	},
}
