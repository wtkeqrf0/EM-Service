package enricher

import "net/http"

type (
	// PossibleAge struct contains possible age of the name.
	PossibleAge struct {
		Count int    `json:"count"`
		Name  string `json:"name"`
		Age   *int   `json:"age"`
	}

	// PossibleGender struct contains possible gender of the name.
	PossibleGender struct {
		Count       int     `json:"count"`
		Name        string  `json:"name"`
		Gender      string  `json:"gender"`
		Probability float64 `json:"probability"`
	}

	// PossibleCountry struct contains possible place of residence of the name.
	PossibleCountry struct {
		Count   int    `json:"count"`
		Name    string `json:"name"`
		Country []struct {
			ID          string  `json:"country_id"`
			Probability float64 `json:"probability"`
		} `json:"country"`
	}

	// HTTPClient defines the interface required to send HTTP requests.
	// This allows you to use any http.Client, and it is convenient for testing.
	HTTPClient interface {
		Do(*http.Request) (*http.Response, error)
	}
)
