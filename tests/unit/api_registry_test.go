package readme_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/liveoaklabs/readme-api-go-client/readme"
	mocks "github.com/liveoaklabs/readme-api-go-client/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_APIRegistry_Get(t *testing.T) {
	t.Run("when called with an existing uuid", func(t *testing.T) {
		// Arrange
		expectResponse := &readme.APIResponse{
			APIErrorResponse: readme.APIErrorResponse{},
			Body: []byte(`{
				"openapi": "3.0.2",
				"info": {
					"description": "OpenAPI Specification for Testing.",
					"version": "2.0.0",
					"title": "API Endpoints",
					"contact": {
						"name": "API Support",
						"url": "https://docs.example.com/docs/contact-support",
						"email": "support@example.com"
					}
				}
			}
		`),
		}
		mockClient := mocks.NewTestClient(t)
		mockClient.APIRegistry.EXPECT().Get("3bbeunznlboryu0o").Return(string(expectResponse.Body), expectResponse, nil)

		// Act
		got, response, err := mockClient.APIRegistry.Get("3bbeunznlboryu0o")

		// Assert
		assert.NoError(t, err, "it does not return an error")
		assert.Equal(t, expectResponse, response, "it returns the expected response")
		assert.Equal(t, string(expectResponse.Body), got, "it returns the expected body")

		mockClient.APIRegistry.AssertExpectations(t)
	})

	t.Run("when API responds with 404", func(t *testing.T) {
		// Arrange
		expectResponse := &readme.APIResponse{
			HTTPResponse: &http.Response{
				StatusCode: 404,
			},
		}
		expectError := errors.New("API responded with a non-OK status: 404")

		mockClient := mocks.NewTestClient(t)
		mockClient.APIRegistry.EXPECT().Get("invalid").Return(string(expectResponse.Body), expectResponse, expectError)

		// Act
		got, gotResponse, err := mockClient.APIRegistry.Get("invalid")

		// Assert
		assert.Error(t, err, "it returns an error")
		assert.ErrorContains(t, err, "API responded with a non-OK status: 404", "it returns the expected error")
		assert.Equal(t, expectResponse, gotResponse, "it returns the expected response")
		assert.Equal(t, string(expectResponse.Body), got, "it returns the expected body")

		mockClient.APIRegistry.AssertExpectations(t)
	})
}

func Test_APIRegistry_Create(t *testing.T) {
	// Arrange
	testDefinition := `
	{
		"openapi": "3.0.0",
		"info": {
			"version": "1.0.0",
			"title": "Test Pet Store Test again",
			"license": {
				"name": "MIT"
			}
		}
	}
	`
	expect := readme.APIRegistrySaved{
		Definition: map[string]interface{}{
			"openapi": "3.0.0",
			"info": map[string]interface{}{
				"version": "1.0.0",
				"title":   "Test Pet Store Test again",
				"license": map[string]interface{}{
					"name": "MIT",
				},
			},
		},
		RegistryUUID: "abcdefghijklmno",
	}

	expectResponse := readme.APIResponse{
		APIErrorResponse: readme.APIErrorResponse{},
		Body:             []byte(`{"registryUUID": "abcdefghijklmno", "definition": ` + testDefinition + `}`),
	}

	t.Run("when called with a definition and no version specified", func(t *testing.T) {
		// Act
		api := mocks.NewTestClient(t)
		api.APIRegistry.EXPECT().Create(testDefinition).Return(expect, &expectResponse, nil)
		got, _, err := api.APIRegistry.Create(testDefinition)

		// Assert
		assert.NoError(t, err, "it does not return an error")
		assert.Equal(t, expect, got, "it returns the expected definition")

		api.APIRegistry.AssertExpectations(t)
	})

	t.Run("when called with a version specified", func(t *testing.T) {
		// Act
		api := mocks.NewTestClient(t)
		expectResponse.Request = &readme.APIRequest{
			RequestOptions: readme.RequestOptions{
				Version: "1.2.3",
			},
		}
		api.APIRegistry.EXPECT().Create(testDefinition, "1.2.3").Return(expect, &expectResponse, nil)
		got, apiResponse, err := api.APIRegistry.Create(testDefinition, "1.2.3")

		// Assert
		assert.NoError(t, err, "it does not return an error")
		assert.Equal(t, "1.2.3", apiResponse.Request.Version, "it returns the expected response")
		assert.Equal(t, "abcdefghijklmno", got.RegistryUUID, "it returns the expected uuid")
		assert.Equal(t, expect, got, "it returns the expected definition")

		api.APIRegistry.AssertExpectations(t)
	})

	t.Run("when API responds with 400", func(t *testing.T) {
		// Arrange
		expectResponse := readme.APIResponse{
			HTTPResponse: &http.Response{
				StatusCode: 400,
			},
			APIErrorResponse: readme.APIErrorResponse{
				Error: "ERROR_SPEC_INVALID",
			},
			Body: []byte(`{"error": "ERROR_SPEC_INVALID"}`),
		}
		expectError := errors.New("API responded with a non-OK status: 400")
		api := mocks.NewTestClient(t)
		api.APIRegistry.EXPECT().Create("").Return(expect, &expectResponse, expectError)

		// Act
		_, apiResponse, err := api.APIRegistry.Create("")

		// Assert
		assert.Equal(t, "ERROR_SPEC_INVALID", apiResponse.APIErrorResponse.Error, "it returns the API error")
		assert.Error(t, err, "it returns an error")
		assert.ErrorContains(t, err, "API responded with a non-OK status: 400", "it returns expected application error")

		api.APIRegistry.AssertExpectations(t)
	})
}

func Test_APIRegistry_validateRegistryUUID(t *testing.T) {
	// Arrange
	testCases := []struct {
		value  string
		expect bool
		msg    string
	}{
		{"uuid:3bbeunznlboryu0o", true, "true when uuid matches pattern and is 16 chars"},
		{"uuid:4xypvmzowqax8hqiix0pqtjx", true, "true when uuid matches pattern and is 24 chars"},
		{"uuid:4xypvmzowqax8hqiix0pxtjxt", false, "false when uuid is >= 25 chars"},
		{"uuid:3bbe0nz", false, "false when uuid is too short"},
		{"notauuid", false, "false when uuid does not match pattern"},
		{`{"openapi": "3.0.0"}`, false, "false when api definition provided instead of uuid"},
	}
	for _, tc := range testCases {
		// Act
		// NOTE: `readme.ParseUUID()` refers to a private `parseUUID()` function that has been exported for tests.
		got, _ := readme.ParseUUID(tc.value)

		// Assert
		assert.Equal(t, tc.expect, got)
	}
}
