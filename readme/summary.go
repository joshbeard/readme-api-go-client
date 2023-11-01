package readme

// Summary provides a high-level summary of a ReadMe project.
type Summary struct {
	Counts            SummaryCounts
	APISpecifications []SummaryAPISpecification
	Categories        []SummaryCategory
	Docs              []SummaryDoc
	Changelogs        []SummaryChangelog
	CustomPages       []SummaryCustomPage
	Versions          []SummaryVersion
}

// SummaryCounts represents the count of each resource type in a ReadMe project.
type SummaryCounts struct {
	APISpecifications int
	Categories        int
	Changelogs        int
	CustomPages       int
	Docs              int
	Versions          int
}

// SummaryAPISpecification represents a brief summary of an API specification.
type SummaryAPISpecification struct {
	ID      string
	Title   string
	Version string
}

// SummaryCategory represents a brief summary of a category.
type SummaryCategory struct {
	ID      string
	Slug    string
	Version string
}

// SummaryDoc represents a brief summary of a doc.
type SummaryDoc struct {
	ID      string
	Slug    string
	Version string
}

// SummaryChangelog represents a brief summary of a changelog.
type SummaryChangelog struct {
	ID   string
	Slug string
}

// SummaryCustomPage represents a brief summary of a custom page.
type SummaryCustomPage struct {
	ID   string
	Slug string
}

// SummaryVersion represents a brief summary of a version.
type SummaryVersion struct {
	ID   string
	Name string
}

// Summary returns a count of each resource type in a ReadMe project.
func (c *Client) Summary() (Summary, []error) {
	specs := make([]SummaryAPISpecification, 0)
	categories := make([]SummaryCategory, 0)
	docs := make([]SummaryDoc, 0)
	changelogs := make([]SummaryChangelog, 0)
	customPages := make([]SummaryCustomPage, 0)
	versionSummary := make([]SummaryVersion, 0)

	// Get all versions.
	versions, _, err := c.Version.GetAll()
	if err != nil {
		return Summary{}, []error{err}
	}

	// Get all versioned resources.
	for _, version := range versions {
		versionSummary = append(versionSummary, SummaryVersion{
			ID:   version.ID,
			Name: version.VersionClean,
		})

		requestOptions := RequestOptions{
			Version: version.VersionClean,
		}

		// Get all specs.
		s, _, err := c.APISpecification.GetAll(requestOptions)
		if err != nil {
			return Summary{}, []error{err}
		}
		for _, spec := range s {
			specs = append(specs, SummaryAPISpecification{
				ID:      spec.ID,
				Title:   spec.Title,
				Version: spec.Version,
			})
		}

		// Get all categories.
		cats, _, err := c.Category.GetAll(requestOptions)
		if err != nil {
			return Summary{}, []error{err}
		}
		for _, cat := range cats {
			categories = append(categories, SummaryCategory{
				ID:      cat.ID,
				Slug:    cat.Slug,
				Version: cat.Version,
			})
		}

		// Get all docs.
		d, _, err := c.Doc.Search("*", requestOptions)
		if err != nil {
			return Summary{}, []error{err}
		}
		for _, doc := range d {
			docs = append(docs, SummaryDoc{
				ID:      doc.ObjectID,
				Slug:    doc.Slug,
				Version: doc.Version,
			})
		}
	}

	// Get all custom pages.
	pages, _, err := c.CustomPage.GetAll()
	if err != nil {
		return Summary{}, []error{err}
	}
	for _, page := range pages {
		customPages = append(customPages, SummaryCustomPage{
			ID:   page.ID,
			Slug: page.Slug,
		})
	}

	// Get all changelogs.
	ch, _, err := c.Changelog.GetAll()
	if err != nil {
		return Summary{}, []error{err}
	}
	for _, changelog := range ch {
		changelogs = append(changelogs, SummaryChangelog{
			ID:   changelog.ID,
			Slug: changelog.Slug,
		})
	}

	return Summary{
		Counts: SummaryCounts{
			APISpecifications: len(specs),
			Categories:        len(categories),
			Changelogs:        len(changelogs),
			CustomPages:       len(customPages),
			Docs:              len(docs),
			Versions:          len(versions),
		},
		APISpecifications: specs,
		Categories:        categories,
		Changelogs:        changelogs,
		CustomPages:       customPages,
		Docs:              docs,
		Versions:          versionSummary,
	}, nil
}
