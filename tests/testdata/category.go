package testdata

import (
	"net/http"

	"github.com/liveoaklabs/readme-api-go-client/readme"
)

var Categories = []readme.Category{
	{
		CreatedAt: "2022-12-04T19:28:15.240Z",
		ID:        "63a77777f52b9f006b6bf212",
		Order:     10,
		Project:   "go-testing",
		Reference: false,
		Slug:      "documentation",
		Title:     "Documentation",
		Type:      "guide",
		Version:   Versions[0].Version,
	},
	{
		CreatedAt: "2022-12-04T19:28:15.240Z",
		ID:        "6543c5bf91e232000cbdc4cf",
		Order:     20,
		Project:   "go-testing",
		Slug:      "example-category",
		Title:     "Example Category",
		Type:      "guide",
		Version:   Versions[0].Version,
	},
	{
		CreatedAt: "2022-12-04T19:28:15.240Z",
		ID:        "654ac951d35521000d16d325",
		Order:     30,
		Project:   "go-testing",
		Slug:      "swagger-petstore-openapi",
		Title:     "Swagger Petstore - OpenAPI",
		Type:      "reference",
		Version:   Versions[0].Version,
	},
}

var CategoryResponse = &readme.APIResponse{
	Body: []byte(ToJSON(Categories)),
	HTTPResponse: &http.Response{
		StatusCode: 200,
	},
}

var CategoryDocs = []readme.CategoryDocs{
	{
		Hidden: false,
		ID:     "63a77777f52b9f006b6bf212",
		Order:  10,
		Slug:   "documentation",
		Title:  "Documentation",
		Children: []readme.CategoryDocsChildren{
			{
				Hidden: false,
				ID:     "63a77777f52b9f006b6bf212",
				Order:  10,
				Slug:   "documentation",
				Title:  "Documentation",
			},
		},
	},
}

var CategorySaved = readme.CategorySaved{
	Title:     Categories[0].Title,
	Slug:      Categories[0].Slug,
	Order:     Categories[0].Order,
	ID:        Categories[0].ID,
	CreatedAt: Categories[0].CreatedAt,
	Reference: Categories[0].Reference,
	Type:      Categories[0].Type,
	Version:   Versions[0],
	Project:   Categories[0].Project,
}

var CategoryVersion = readme.CategoryVersion{
	Categories: []readme.Category{Categories[0]},
	Codename:   Versions[0].Codename,
	CreatedAt:  Versions[0].CreatedAt,
	ForkedFrom: readme.CategoryVersionForkedFrom{
		Version: Versions[0].VersionClean,
		ID:      Versions[0].Version,
	},
	ID:           Versions[0].ID,
	IsBeta:       Versions[0].IsBeta,
	IsDeprecated: Versions[0].IsDeprecated,
	IsHidden:     Versions[0].IsHidden,
	IsStable:     Versions[0].IsStable,
	Project:      Versions[0].Project,
	ReleaseDate:  Versions[0].CreatedAt,
	Version:      Versions[0].Version,
	VersionClean: Versions[0].VersionClean,
}

var CategoryVersionSaved = readme.CategoryVersionSaved{
	Title:     Categories[0].Title,
	Slug:      Categories[0].Slug,
	Order:     Categories[0].Order,
	ID:        Categories[0].ID,
	CreatedAt: Categories[0].CreatedAt,
	Reference: Categories[0].Reference,
	Type:      Categories[0].Type,
	Version:   CategoryVersion,
	Project:   Categories[0].Project,
}

// type CategoryVersion struct {
// 	Categories   []Category                `json:"categories"`
// 	Codename     string                    `json:"codename"`
// 	CreatedAt    string                    `json:"createdAt"`
// 	ForkedFrom   CategoryVersionForkedFrom `json:"forked_from"`
// 	ID           string                    `json:"id"`
// 	IsBeta       bool                      `json:"is_beta"`
// 	IsDeprecated bool                      `json:"is_deprecated"`
// 	IsHidden     bool                      `json:"is_hidden"`
// 	IsStable     bool                      `json:"is_stable"`
// 	Project      string                    `json:"project"`
// 	ReleaseDate  string                    `json:"releaseDate"`
// 	Version      string                    `json:"version"`
// 	VersionClean string                    `json:"version_clean"`
// }

// [{"categoryType":"","createdAt":"2023-11-02T15:52:31.795Z","id":"6543c5bf91e232000cbdc4cf","order":9999,"project":"650bcc2ce74f7f0025788168","reference":false,"slug":"example-category","title":"Example Category","type":"guide","version":"650bcc2de74f7f002578816e"},{"categoryType":"","createdAt":"2023-11-07T23:33:37.168Z","id":"654ac951d35521000d16d325","order":9999,"project":"650bcc2ce74f7f0025788168","reference":true,"slug":"swagger-petstore-openapi-30","title":"Swagger Petstore - OpenAPI 3.0","type":"reference","version":"650bcc2de74f7f002578816e"},{"categoryType":"","createdAt":"2023-11-08T00:51:56.502Z","id":"654adbac38410f00432df038","order":9999,"project":"650bcc2ce74f7f0025788168","reference":true,"slug":"swagger-petstore-openapi-30-1","title":"Swagger Petstore - OpenAPI 3.0","type":"reference","version":"650bcc2de74f7f002578816e"},{"categoryType":"","createdAt":"2023-11-17T05:10:20.420Z","id":"6556f5bc9a338001f3043a7b","order":9999,"project":"650bcc2ce74f7f0025788168","reference":false,"slug":"example-category-1","title":"Example Category","type":"guide","version":"650bcc2de74f7f002578816e"}]

// mockCategoryVersionCreateResponse represents the response from the API when
// calling Create().
// var MockCategoryVersionCreateResponse = readme.CategoryVersionSaved{
// 	Title:     "Test Category",
// 	Slug:      "test-category",
// 	Order:     9999,
// 	Reference: false,
// 	ID:        "63b463a1d692480062747ef6",
// 	Project:   "638cf4cedea3ff0096d1a955",
// 	Version: readme.CategoryVersion{
// 		Version:      "1.1.0",
// 		VersionClean: "1.1.0",
// 		Codename:     "Some test",
// 		IsStable:     true,
// 		IsBeta:       true,
// 		IsHidden:     false,
// 		IsDeprecated: false,
// 		Categories:   []readme.Category{mockCategory},
// 		ID:           "63a77777f52b9f006b6bf215",
// 		Project:      "638cf4cedea3ff0096d1a955",
// 		ReleaseDate:  "2022-12-04T19:28:15.190Z",
// 		CreatedAt:    "2022-12-04T19:28:15.190Z",
// 		ForkedFrom: readme.CategoryVersionForkedFrom{
// 			Version: mockVersions[0].VersionClean,
// 			ID:      mockVersions[0].Version,
// 		},
// 	},
// 	CreatedAt: "2023-01-03T17:19:29.388Z",
// 	Type:      "guide",
// }
