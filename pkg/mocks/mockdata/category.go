package mockdata

import "github.com/liveoaklabs/readme-api-go-client/readme"

// Category represents a category in a project.
// This is the response from the API when calling Get() and Update() and is
// also referenced by other types of responses throughout the API.
var Category = readme.Category{
	Title:     "Documentation",
	Slug:      "documentation",
	Order:     1,
	Reference: false,
	ID:        "63a77777f52b9f006b6bf212",
	Version:   "63a77777f52b9f006b6bf215",
	Project:   "638cf4cedea3ff0096d1a955",
	CreatedAt: "2022-12-04T19:28:15.240Z",
	Type:      "guide",
}

var Categories = []readme.Category{
	Category,
	{
		Title:     "API",
		Slug:      "api",
		Order:     2,
		Reference: false,
		ID:        "63a77777f52b9f006b6bf213",
		Version:   "63a77777f52b9f006b6bf215",
		Project:   "638cf4cedea3ff0096d1a955",
		CreatedAt: "2022-12-04T19:28:15.240Z",
		Type:      "guide",
	},
	{
		Title:     "SDK",
		Slug:      "sdk",
		Order:     3,
		Reference: false,
		ID:        "63a77777f52b9f006b6bf214",
		Version:   "63a77777f52b9f006b6bf215",
		Project:   "638cf4cedea3ff0096d1a955",
		CreatedAt: "2022-12-04T19:28:15.240Z",
		Type:      "guide",
	},
}

// CategoryVersionCreateResponse represents the response from the API when
// calling Create().
var CategoryVersionCreateResponse = readme.CategoryVersionSaved{
	Title:     "Test Category",
	Slug:      "test-category",
	Order:     9999,
	Reference: false,
	ID:        "63b463a1d692480062747ef6",
	Project:   "638cf4cedea3ff0096d1a955",
	Version: readme.CategoryVersion{
		Version:      "1.1.0",
		VersionClean: "1.1.0",
		Codename:     "Some test",
		IsStable:     true,
		IsBeta:       true,
		IsHidden:     false,
		IsDeprecated: false,
		Categories:   Categories,
		ID:           "63a77777f52b9f006b6bf215",
		Project:      "638cf4cedea3ff0096d1a955",
		ReleaseDate:  "2022-12-04T19:28:15.190Z",
		CreatedAt:    "2022-12-04T19:28:15.190Z",
		ForkedFrom: readme.CategoryVersionForkedFrom{
			Version: Version.VersionClean,
			ID:      Version.Version,
		},
	},
	CreatedAt: "2023-01-03T17:19:29.388Z",
	Type:      "guide",
}
