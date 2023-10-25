package mockdata

import "github.com/liveoaklabs/readme-api-go-client/readme"

var Doc = readme.Doc{
	API: readme.DocAPI{
		APISetting: "63c77e0bf76dee008e0c5197",
		Auth:       "required",
		Examples: readme.DocAPIExamples{
			Codes: []readme.DocAPIExamplesCodes{
				{
					Code:     "",
					Language: "c",
				},
			},
		},
		Method: "get",
		Params: []readme.DocAPIParams{
			{
				Default:    "foo",
				Desc:       "this is an example path param",
				EnumValues: "",
				In:         "path",
				Name:       "example_path_param",
				Ref:        "",
				Required:   false,
				Type:       "string",
			},
		},
		Results: readme.DocAPIResults{
			Codes: []readme.DocAPIResultsCodes{
				{
					Code:     "{\"message\": \"ok\"}",
					Language: "json",
					Name:     "",
					Status:   200,
				},
				{
					Code:     "{}",
					Language: "json",
					Name:     "",
					Status:   400,
				},
			},
		},
		URL: "/{something}",
	},
	Algolia: readme.DocAlgolia{
		PublishPending: false,
		RecordCount:    0,
		UpdatedAt:      "2023-01-18T05:22:34.529Z",
	},
	Body:         "A turtle has been here.",
	BodyHTML:     "<div class=\"magic-block-textarea\"><p>A turtle has been here.</p></div>",
	Category:     "63a77777f52b9f006b6bf212",
	CreatedAt:    "2023-01-03T01:02:03.731Z",
	Deprecated:   false,
	Excerpt:      "",
	Icon:         "",
	ID:           "63b37e8b65fd5b0057af23f1",
	IsAPI:        false,
	IsReference:  false,
	LinkExternal: false,
	LinkURL:      "",
	Metadata: readme.DocMetadata{
		Description: "This is an example description.",
		Image: []interface{}{
			"https://files.readme.io/53c37f9-live-oak-logo-full-color.svg",
			"live-oak-logo-full-color.svg",
			950,
			135,
			"#1c0e52",
		},
		Title: "Example Doc",
	},
	Next: readme.DocNext{
		Description: "",
		Pages: []readme.DocNextPages{
			{
				Category:   "Documentation",
				Deprecated: false,
				Icon:       "file-text-o",
				Name:       "The Next Doc",
				Slug:       "the-next-doc",
				Type:       "doc",
			},
		},
	},
	Order:         999,
	Project:       "638cf4cedea3ff0096d1a955",
	PreviousSlug:  "",
	Revision:      2,
	Slug:          "example-doc",
	SlugUpdatedAt: "2023-01-02T01:44:37.530Z",
	SyncUnique:    "",
	Title:         "Example Doc",
	Type:          "basic",
	UpdatedAt:     "2023-01-03T01:02:03.731Z",
	User:          "633c5a54187d2c008e2e074c",
	Version:       "63a77777f52b9f006b6bf215",
}

var mockDocResponseBody string = `{
		"metadata": {
			"image": [
			  "https://files.readme.io/53c37f9-live-oak-logo-full-color.svg",
			  "live-oak-logo-full-color.svg",
			  950,
			  135,
			  "#1c0e52"
			],
			"title": "Example Doc",
			"description": "This is an example description."
		},
		"algolia": {
			"recordCount": 0,
			"publishPending": false,
			"updatedAt": "2023-01-18T05:22:34.529Z"
		},
		"api": {
			"method": "get",
			"url": "/{something}",
			"auth": "required",
			"params": [
			  {
				"name": "example_path_param",
				"type": "string",
				"enumValues": "",
				"default": "foo",
				"desc": "this is an example path param",
				"required": false,
				"in": "path",
				"ref": "",
				"_id": "63c77f40e790b60061a9a2ed"
			  }
		    ],
			"apiSetting": "63c77e0bf76dee008e0c5197",
			"examples": {
			  "codes": [
				{
				  "code": "",
				  "language": "c"
				}
			  ]
			},
			"results": {
				"codes": [
				  {
					"name": "",
					"code": "{\"message\": \"ok\"}",
					"language": "json",
					"status": 200
				  },
				  {
					"name": "",
					"code": "{}",
					"language": "json",
					"status": 400
				  }
				]
  		    }
		},
		"next": {
			"description": "",
			"pages": [
			  {
				"type": "doc",
				"icon": "file-text-o",
				"name": "The Next Doc",
				"slug": "the-next-doc",
				"deprecated": false,
				"category": "Documentation"
			  }
			]
		},
		"title": "Example Doc",
		"icon": "",
		"updates": [],
		"type": "basic",
		"slug": "example-doc",
		"excerpt": "",
		"body": "A turtle has been here.",
		"order": 999,
		"isReference": false,
		"deprecated": false,
		"hidden": true,
		"sync_unique": "",
		"link_url": "",
		"link_external": false,
		"pendingAlgoliaPublish": false,
		"previousSlug": "",
		"slugUpdatedAt": "2023-01-02T01:44:37.530Z",
		"revision": 2,
		"_id": "63b37e8b65fd5b0057af23f1",
		"category": "63a77777f52b9f006b6bf212",
		"project": "638cf4cedea3ff0096d1a955",
		"createdAt": "2023-01-03T01:02:03.731Z",
		"updatedAt": "2023-01-03T01:02:03.731Z",
		"user": "633c5a54187d2c008e2e074c",
		"version": "63a77777f52b9f006b6bf215",
		"__v": 0,
		"isApi": false,
		"id": "63b37e8b65fd5b0057af23f1",
		"body_html": "<div class=\"magic-block-textarea\"><p>A turtle has been here.</p></div>"
	}`

var DocSearchResult = []readme.DocSearchResult{{
	IndexName:    "Page",
	Title:        "Example of manual page",
	Type:         "basic",
	Slug:         "example-of-manual-page",
	IsReference:  false,
	Method:       "get",
	LinkURL:      "",
	Version:      "63b891d3ee384600680ce9f9",
	Project:      "63b891d3ee384600680ce9eb",
	ReferenceID:  "63c7828f56279c006fbd4178",
	Subdomain:    "example",
	InternalLink: "docs/example-of-manual-page",
	ObjectID:     "63c7828f56279c006fbd4178-1",
	SnippetResult: readme.DocSearchResultSnippet{
		Title: readme.DocSearchResultSnippetValue{
			Value:      "Example of manual page",
			MatchLevel: "none",
		},
		Excerpt: readme.DocSearchResultSnippetValue{
			Value:      "",
			MatchLevel: "none",
		},
		Body: readme.DocSearchResultSnippetValue{
			Value:      "This is a test page created manually.",
			MatchLevel: "none",
		},
	},
	HighlightResult: readme.DocSearchResultHighlight{
		Title: readme.DocSearchResultHighlightValue{
			Value:      "Example of manual page",
			MatchLevel: "none",
			MatchedWords: []string{
				"manual",
			},
		},
		Excerpt: readme.DocSearchResultHighlightValue{
			Value:        "",
			MatchLevel:   "none",
			MatchedWords: []string{},
		},
		Body: readme.DocSearchResultHighlightValue{
			Value:      "This is a test page created manually.",
			MatchLevel: "none",
			MatchedWords: []string{
				"manual",
			},
		},
	},
	URL: "https://example.readme.io/docs/example-of-manual-page",
}}
