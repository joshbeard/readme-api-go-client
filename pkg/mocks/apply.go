package mocks

import "github.com/liveoaklabs/readme-api-go-client/readme"

type MockApplyService interface {
	Apply(application readme.Application) (readme.ApplyResponse, *readme.APIResponse, error)
}

type MockApplyClient struct {
	client *MockClient
}

func (c *MockApplyClient) Apply(application readme.Application) (readme.ApplyResponse, *readme.APIResponse, error) {
	response := readme.ApplyResponse{
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

	return response, nil, nil
}
