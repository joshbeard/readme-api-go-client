package mockdata

import "github.com/liveoaklabs/readme-api-go-client/readme"

var (
	SummaryAPISpecifications = []readme.SummaryAPISpecification{
		{
			ID:      APISpecification.ID,
			Title:   APISpecification.Title,
			Version: APISpecification.Version,
		},
	}

	SummaryCategories = []readme.SummaryCategorie{
		{
			ID:      Category.ID,
			Slug:    Category.Slug,
			Version: Category.Version,
		},
	}

	SummaryDocs = []readme.SummaryDoc{
		{
			ID:      Doc.ID,
			Slug:    Doc.Slug,
			Version: Doc.Version,
		},
	}

	SummaryChangelogs = []readme.SummaryChangelog{
		{
			ID:   Changelog.ID,
			Slug: Changelog.Slug,
		},
	}

	SummaryCustomPages = []readme.SummaryCustomPage{
		{
			ID:   CustomPage.ID,
			Slug: CustomPage.Slug,
		},
	}

	SummaryVersions = []readme.SummaryVersion{
		{
			ID:   Version.ID,
			Name: Version.VersionClean,
		},
	}

	Summary = readme.Summary{
		Counts: readme.SummaryCounts{
			APISpecifications: len(SummaryAPISpecifications),
			Categories:        len(SummaryCategories),
			Changelogs:        len(SummaryChangelogs),
			CustomPages:       len(SummaryCustomPages),
			Docs:              len(SummaryDocs),
			Versions:          len(SummaryVersions),
		},
		APISpecifications: SummaryAPISpecifications,
		Categories:        SummaryCategories,
		Changelogs:        SummaryChangelogs,
		CustomPages:       SummaryCustomPages,
		Docs:              SummaryDocs,
		Versions:          SummaryVersions,
	}
)
