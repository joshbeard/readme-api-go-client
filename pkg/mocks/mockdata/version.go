package mockdata

import "github.com/liveoaklabs/readme-api-go-client/readme"

// Version represents a version of a project.
// This is the response from the API when calling Get().
var Version = readme.Version{
	Codename:     "",
	CreatedAt:    "2022-12-04T19:28:15.190Z",
	ID:           "638cf4cfdea3ff0096d1a95a",
	IsBeta:       false,
	IsDeprecated: false,
	IsHidden:     false,
	IsStable:     true,
	Version:      "1.0.0",
	VersionClean: "1.0.0",
}

// VersionSummary represents a list of versions of a project.
// This is the response from the API when calling GetAll().
var VersionSummary = []readme.VersionSummary{
	{
		ForkedFrom:   "",
		Codename:     Version.Codename,
		CreatedAt:    Version.CreatedAt,
		ID:           Version.ID,
		IsBeta:       Version.IsBeta,
		IsDeprecated: Version.IsDeprecated,
		IsHidden:     Version.IsHidden,
		IsStable:     Version.IsStable,
		Version:      Version.Version,
		VersionClean: Version.VersionClean,
	},
	{
		ForkedFrom:   "1.0.0",
		Codename:     "",
		CreatedAt:    "2022-12-04T19:28:15.190Z",
		ID:           "639fe99365417e008daa1f55",
		IsBeta:       false,
		IsDeprecated: false,
		IsHidden:     false,
		IsStable:     true,
		Version:      "1.1.0",
		VersionClean: "1.1.0",
	},
}
