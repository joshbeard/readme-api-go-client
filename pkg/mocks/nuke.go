package mocks

type MockNukeService interface {
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

type MockNukeClient struct {
	client *MockClient
}

func (c *MockNukeClient) DestroyAll() []error {
	return nil
}

func (c *MockNukeClient) DestroyCategories() []error {
	return nil
}

func (c *MockNukeClient) DestroyChangelogs() []error {
	return nil
}

func (c *MockNukeClient) DestroyCustomPages() []error {
	return nil
}

func (c *MockNukeClient) DestroyDocs() []error {
	return nil
}

func (c *MockNukeClient) DestroySpecs() []error {
	return nil
}

func (c *MockNukeClient) DestroyVersions() []error {
	return nil
}
