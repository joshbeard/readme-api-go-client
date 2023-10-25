package readme

import (
	"fmt"
)

// NukeService provides functions for resetting a ReadMe project to its default
// nearly empty state. This is not a functionality provided by the ReadMe API.
// Instead, it is convenience functions that use the ReadMe API to delete
// resources.

// ReadMe doesn't allow the default 1.0.0 version to be deleted.
const (
	ExcludedVersion = "1.0.0"
)

type NukeService interface {
	// Destroy all user-created data in a ReadMe project.
	DestroyAll() []error

	// Destroy all categories in a ReadMe project.
	DestroyCategories() []error

	// Destroy all changelogs in a ReadMe project.
	DestroyChangelogs() []error

	// Destroy all custom pages in a ReadMe project.
	DestroyCustomPages() []error

	// Destroy all docs in a ReadMe project.
	DestroyDocs() []error

	// Destroy all API specifications in a ReadMe project.
	DestroySpecs() []error

	// Destroy all versions in a ReadMe project.
	DestroyVersions() []error
}

// NukeClient handles resetting a ReadMe project to its default nearly empty
// state.
type NukeClient struct {
	client *Client
}

// Ensure the implementation satisfies the expected interfaces.
var _ NukeService = &NukeClient{}

// Destroy all user-created data in a ReadMe project.
// Returns a slice of errors encountered.
func (s NukeClient) DestroyAll() []error {
	var errs []error

	err := s.DestroyChangelogs()
	if err != nil {
		errs = append(errs, err...)
	}

	err = s.DestroyCustomPages()
	if err != nil {
		errs = append(errs, err...)
	}

	err = s.DestroyDocs()
	if err != nil {
		errs = append(errs, err...)
	}

	err = s.DestroySpecs()
	if err != nil {
		errs = append(errs, err...)
	}

	err = s.DestroyCategories()
	if err != nil {
		errs = append(errs, err...)
	}

	err = s.DestroyVersions()
	if err != nil {
		errs = append(errs, err...)
	}

	return errs
}

// DestroySpecs destroys all API specifications in a ReadMe project.
// A slice of errors is returned.
func (s NukeClient) DestroySpecs() []error {
	var errs []error

	versions, _, err := s.client.Version.GetAll()
	if err != nil {
		return []error{fmt.Errorf("cannot destroy specs: %w", err)}
	}

	for _, version := range versions {
		requestOptions := RequestOptions{
			Version: version.VersionClean,
		}

		specs, _, err := s.client.APISpecification.GetAll(requestOptions)
		if err != nil {
			errs = append(errs, fmt.Errorf("cannot destroy specs: %w", err))
		}

		for _, spec := range specs {
			_, _, err := s.client.APISpecification.Delete(spec.ID)
			if err != nil {
				errs = append(errs, fmt.Errorf("cannot destroy spec %s: %w", spec.ID, err))
			}
		}
	}

	return errs
}

// DestroyDocs destroys all docs in a ReadMe project.
// A slice of errors is returned.
func (s NukeClient) DestroyDocs() []error {
	var errs []error

	versions, _, err := s.client.Version.GetAll()
	if err != nil {
		return []error{fmt.Errorf("cannot destroy docs: %w", err)}
	}

	for _, version := range versions {
		requestOptions := RequestOptions{
			Version: version.VersionClean,
		}

		docs, _, err := s.client.Doc.Search("*", requestOptions)
		if err != nil {
			errs = append(errs, fmt.Errorf("cannot destroy docs: %w", err))
		}

		for _, doc := range docs {
			_, _, err := s.client.Doc.Delete(doc.Slug, requestOptions)
			if err != nil {
				errs = append(errs, fmt.Errorf("cannot destroy doc %s: %w", doc.Slug, err))
			}
		}
	}

	return errs
}

// DestroyCategories destroys all categories in a ReadMe project.
// A slice of errors is returned.
func (s NukeClient) DestroyCategories() []error {
	var errs []error

	versions, _, err := s.client.Version.GetAll()
	if err != nil {
		return []error{fmt.Errorf("cannot destroy categories: %w", err)}
	}

	for _, version := range versions {
		requestOptions := RequestOptions{
			Version: version.VersionClean,
		}

		categories, _, err := s.client.Category.GetAll(requestOptions)
		if err != nil {
			errs = append(errs, fmt.Errorf("cannot destroy categories: %w", err))
		}

		for _, category := range categories {
			_, _, err := s.client.Category.Delete(category.Slug, requestOptions)
			if err != nil {
				errs = append(errs, fmt.Errorf("cannot destroy category %s: %w",
					category.ID, err))
			}
		}
	}

	return errs
}

// DestroyChangelogs destroys all changelogs in a ReadMe project.
// A slice of errors is returned.
func (s NukeClient) DestroyChangelogs() []error {
	var errs []error

	changelogs, _, err := s.client.Changelog.GetAll()
	if err != nil {
		return []error{fmt.Errorf("cannot destroy changelogs: %w", err)}
	}

	for _, changelog := range changelogs {
		_, _, err := s.client.Changelog.Delete(changelog.Slug)
		if err != nil {
			errs = append(errs, fmt.Errorf("cannot destroy changelog %s (%s): %w",
				changelog.Slug, changelog.ID, err))
		}
	}

	return errs
}

// DestroyCustomPages destroys all custom pages in a ReadMe project.
// A slice of errors is returned.
func (s NukeClient) DestroyCustomPages() []error {
	var errs []error

	pages, _, err := s.client.CustomPage.GetAll()
	if err != nil {
		return []error{fmt.Errorf("cannot destroy custom pages: %w", err)}
	}

	for _, page := range pages {
		_, _, err := s.client.CustomPage.Delete(page.Slug)
		if err != nil {
			errs = append(errs, fmt.Errorf("cannot destroy custom page %s: %w",
				page.ID, err))
		}
	}

	return errs
}

// DestroyVersions destroys all versions in a ReadMe project.
// A slice of errors is returned.
func (s NukeClient) DestroyVersions() []error {
	var errs []error

	versions, _, err := s.client.Version.GetAll()
	if err != nil {
		return []error{fmt.Errorf("cannot destroy versions: %w", err)}
	}

	for _, version := range versions {
		if version.VersionClean == ExcludedVersion {
			continue
		}

		_, _, err := s.client.Version.Delete(version.VersionClean)
		if err != nil {
			errs = append(errs, fmt.Errorf("cannot destroy version %s: %w",
				version.VersionClean, err))
		}
	}

	return errs
}
