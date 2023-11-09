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
// Errors are returned in a slice.
func (c *Client) Summary() (Summary, []error) {
	var retErrs []error

	// Get all versions.
	versions, _, err := c.Version.GetAll()
	if err != nil {
		return Summary{}, []error{err}
	}

	specs := make([]SummaryAPISpecification, 0)
	categories := make([]SummaryCategory, 0)
	docs := make([]SummaryDoc, 0)
	versionSummary := make([]SummaryVersion, 0)

	// Get all versioned resources.
	for _, version := range versions {
		versionSummary = append(versionSummary, SummaryVersion{
			ID:   version.ID,
			Name: version.VersionClean,
		})

		requestOptions := RequestOptions{
			Version: version.VersionClean,
		}

		// Get API specifications for this version.
		s, errs := c.SummaryAPISpecifications(requestOptions)
		if errs != nil {
			retErrs = append(retErrs, errs...)
		}
		specs = append(specs, s...)

		// Get all categories for this version.
		cats, errs := c.SummaryCategories(requestOptions)
		if errs != nil {
			retErrs = append(retErrs, errs...)
		}
		categories = append(categories, cats...)

		// Get all docs for this version.
		d, errs := c.SummaryDocs(requestOptions)
		if errs != nil {
			retErrs = append(retErrs, errs...)
		}
		docs = append(docs, d...)
	}

	// Get all custom pages.
	customPages, errs := c.SummaryCustomPages()
	if errs != nil {
		retErrs = append(retErrs, errs...)
	}

	// Get all changelogs.
	changelogs, errs := c.SummaryChangelogs()
	if errs != nil {
		retErrs = append(retErrs, errs...)
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
	}, retErrs
}

// SummaryAPISpecifications returns a summary of all API specifications in a
// ReadMe project. The requestOptions argument is used to specify a version.
func (c Client) SummaryAPISpecifications(requestOptions RequestOptions) ([]SummaryAPISpecification, []error) {
	specs, _, err := c.APISpecification.GetAll(requestOptions)
	if err != nil {
		return nil, []error{err}
	}

	summary := make([]SummaryAPISpecification, 0)
	for _, spec := range specs {
		summary = append(summary, SummaryAPISpecification{
			ID:      spec.ID,
			Title:   spec.Title,
			Version: spec.Version,
		})
	}

	return summary, nil
}

// SummaryCategories returns a summary of all categories in a ReadMe project.
// The requestOptions argument is used to specify a version.
func (c Client) SummaryCategories(requestOptions RequestOptions) ([]SummaryCategory, []error) {
	categories, _, err := c.Category.GetAll(requestOptions)
	if err != nil {
		return nil, []error{err}
	}

	summary := make([]SummaryCategory, 0)
	for _, category := range categories {
		summary = append(summary, SummaryCategory{
			ID:      category.ID,
			Slug:    category.Slug,
			Version: category.Version,
		})
	}

	return summary, nil
}

// SummaryChangelogs returns a summary of all changelogs in a ReadMe project.
func (c Client) SummaryChangelogs() ([]SummaryChangelog, []error) {
	changelogs, _, err := c.Changelog.GetAll()
	if err != nil {
		return nil, []error{err}
	}

	summary := make([]SummaryChangelog, 0)
	for _, changelog := range changelogs {
		summary = append(summary, SummaryChangelog{
			ID:   changelog.ID,
			Slug: changelog.Slug,
		})
	}

	return summary, nil
}

// SummaryCustomPages returns a summary of all custom pages in a ReadMe project.
func (c Client) SummaryCustomPages() ([]SummaryCustomPage, []error) {
	pages, _, err := c.CustomPage.GetAll()
	if err != nil {
		return nil, []error{err}
	}

	summary := make([]SummaryCustomPage, 0)
	for _, page := range pages {
		summary = append(summary, SummaryCustomPage{
			ID:   page.ID,
			Slug: page.Slug,
		})
	}

	return summary, nil
}

// SummaryDocs returns a summary of all docs in a ReadMe project. The
// requestOptions argument is used to specify a version.
func (c Client) SummaryDocs(requestOptions RequestOptions) ([]SummaryDoc, []error) {
	docs, _, err := c.Doc.Search("*", requestOptions)
	if err != nil {
		return nil, []error{err}
	}

	summary := make([]SummaryDoc, 0)
	for _, doc := range docs {
		summary = append(summary, SummaryDoc{
			ID:      doc.ObjectID,
			Slug:    doc.Slug,
			Version: doc.Version,
		})
	}

	return summary, nil
}
