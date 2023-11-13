package readme_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/liveoaklabs/readme-api-go-client/internal/testutil"
	"github.com/liveoaklabs/readme-api-go-client/readme"
	"github.com/liveoaklabs/readme-api-go-client/tests/mocks"
	"github.com/stretchr/testify/assert"
)

const applyTestEndpoint = "http://readme-test.local/api/v1/apply"

func Test_Apply_Get(t *testing.T) {
	// Arrange
	expect := []readme.OpenRole{
		{
			Slug:        "front-end-engineer",
			Title:       "Front End Engineer",
			Description: "foo bar",
			Pullquote:   "Collaborative: Your code reviews and pull requests are detailed, you’re always happy to share knowledge, and you love a good pairing session.",
			Location:    "Remote, US",
			Department:  "Engineering",
			URL:         "https://jobs.ashbyhq.com/readme/43ebc7c3-4653-4037-841e-075ad428a68d/application",
		},
	}

	expectResponse := &readme.APIResponse{
		APIErrorResponse: readme.APIErrorResponse{},
		Body:             []byte(testutil.StructToJson(t, expect)),
		HTTPResponse: &http.Response{
			StatusCode: 200,
		},
	}

	api := mocks.NewTestClient(t)
	api.Apply.EXPECT().Get().Return(expect, expectResponse, nil)

	// Act
	got, _, err := api.Apply.Get()

	// Assert
	assert.NoError(t, err, "it does not return an error")
	assert.Equal(t, expect, got, "it returns open roles")
}

func Test_Apply_Apply(t *testing.T) {
	t.Run("when called with valid params", func(t *testing.T) {
		// Arrange
		expect := readme.ApplyResponse{
			Message:   "Thanks for applying, Gordon! We'll reach out to you soon!",
			Keyvalues: "https://www.keyvalues.com/readme",
			Careers:   "https://readme.com/careers",
			Questions: "greg@readme.io",
			Poem: []string{
				"Thanks for applying to work at ReadMe!",
				"Your application is lookin' spiffy",
				"We're going to review it ASAP",
				"And we'll get back to you in a jiffy!",
			},
		}

		expectResponse := &readme.APIResponse{
			APIErrorResponse: readme.APIErrorResponse{},
			Body:             []byte(testutil.StructToJson(t, expect)),
		}
		application := readme.Application{
			Name:  "Gordon Ramsay",
			Email: "gordon@example.com",
			Job:   "Front End Engineer",
		}

		api := mocks.NewTestClient(t)
		api.Apply.EXPECT().Apply(application).Return(expect, expectResponse, nil)

		// Act
		got, _, err := api.Apply.Apply(application)

		// Assert
		assert.NoError(t, err, "it does not return an error")
		assert.Equal(t, expect, got, "it returns the expected API response message")

		api.Apply.AssertExpectations(t)
	})

	t.Run("when called with invalid params", func(t *testing.T) {
		// Arrange
		expect := readme.APIErrorResponse{
			Error:      "APPLY_INVALID_NAME",
			Message:    "You need to provide a name.",
			Suggestion: "To apply for a job, you need to include your name as a body parameter.",
			Docs:       "https://docs.readme.com/main/logs/ddfb22d8-dfc6-43e8-b15f-4f4d9092f27d",
			Help:       "https://docs.readme.com/main/reference and include the following link to your API log: 'https://docs.readme.com/main/logs/ddfb22d8-dfc6-43e8-b15f-4f4d9092f27d'.",
			Poem: []string{
				"We normally support online anonymity",
				"Except when you're trying to apply,",
				"Because if we're gonna work together",
				"We need to know what you go by!",
			},
		}

		expectResponse := &readme.APIResponse{
			APIErrorResponse: expect,
			Body:             []byte(testutil.StructToJson(t, expect)),
			HTTPResponse: &http.Response{
				StatusCode: 400,
			},
		}

		expectError := errors.New("API responded with a non-OK status: 400")

		application := readme.Application{
			Name:  "",
			Email: "",
			Job:   "Front End Engineer",
		}

		api := mocks.NewTestClient(t)
		api.Apply.EXPECT().Apply(application).Return(readme.ApplyResponse{}, expectResponse, expectError)

		// Act
		_, apiResponse, err := api.Apply.Apply(application)

		// Assert
		assert.Error(t, err, "it returns an error")
		assert.Equal(t, expect, apiResponse.APIErrorResponse, "it returns the expected API response message")

		api.Apply.AssertExpectations(t)
	})
}
