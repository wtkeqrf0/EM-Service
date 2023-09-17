package enricher

import (
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"sync"
)

// Enricher represents a client for interacting with REST service endpoints.
//
// Controllers are generated based on Enricher interface and mock.MockEnricher implementation.
//
//go:generate ifacemaker -f enricher.go -o controller/enricher.go -i Enricher -s Enricher -p controller -y "Controller describes methods, implemented by the enricher package."
//go:generate mockgen -package mock -source controller/enricher.go -destination controller/mock/mock_enricher.go
type Enricher struct {
	HTTPClient
	AgeURL     string
	GenderURL  string
	CountryURL string
}

// New creates a new Enricher.
func New(HTTPClient HTTPClient, ageURL, genderURL, countryURL string) *Enricher {
	return &Enricher{
		HTTPClient: HTTPClient,
		AgeURL:     ageURL,
		GenderURL:  genderURL,
		CountryURL: countryURL,
	}
}

// FIO struct represents the full name.
type FIO struct {
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic *string `json:"patronymic,omitempty"`
}

// EnrichedFIO represents enriched full name.
type EnrichedFIO struct {
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic *string `json:"patronymic,omitempty"`
	Age        int     `json:"age"`
	Gender     string  `json:"gender"`
	CountryID  string  `json:"countryId"`
}

// EnrichFIO sends requests to enrich the FIO structure.
// The method expects the context (ctx) and the full user name (FIO).
//
// Eventually the FIOEnriched structure is returned.
//
// If FIO is incorrect, all new fields will have zero values.
func (e *Enricher) EnrichFIO(
	ctx context.Context,
	fio FIO,
) (resp EnrichedFIO, err error) {
	var wg sync.WaitGroup

	go func() {
		defer wg.Done()
		var pc PossibleCountry
		err = e.executeEnrichRequest(
			ctx,
			e.CountryURL,
			fio.Name,
			&pc,
		)
		if pc.Country != nil && len(pc.Country) > 0 {
			resp.CountryID = pc.Country[0].ID
		}
	}()

	go func() {
		defer wg.Done()
		var pa PossibleAge
		err = e.executeEnrichRequest(
			ctx,
			e.AgeURL,
			fio.Name,
			&pa,
		)
		resp.Age = *pa.Age
	}()

	go func() {
		defer wg.Done()
		var pg PossibleGender
		err = e.executeEnrichRequest(
			ctx,
			e.GenderURL,
			fio.Name,
			&pg,
		)
		resp.Gender = pg.Gender
	}()

	resp.Name, resp.Surname, resp.Patronymic = fio.Name, fio.Surname, fio.Patronymic
	wg.Add(3)
	wg.Wait()
	log.Debug(resp)
	return
}

// executeEnrichRequest executes a REST request to endpoint <urlName>.
//
// The function accepts the context (ctx), the URL of the REST service (urlName),
// the user name (name) and the place to record the response (response).
// It is assumed that response is a structure that can be serialized in JSON.
//
// First, a REST request is created with the url parameter `name`.
// The request is sent to <urlName> using the GET method.
//
// After receiving a response from the API, the function reads the received response body,
// and tries to sterilize it into the response structure.
//
// In case of any errors during request or response processing, the function returns an error.
func (e *Enricher) executeEnrichRequest(
	ctx context.Context,
	urlName,
	name string,
	response any,
) error {
	var req, _ = http.NewRequestWithContext(ctx, http.MethodGet, urlName, nil)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.URL.RawQuery = url.Values{"name": []string{name}}.Encode()

	jsonResp, err := e.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	respBody, err := io.ReadAll(jsonResp.Body)
	if err != nil {
		return err
	}
	defer jsonResp.Body.Close()

	if err = json.Unmarshal(respBody, response); err != nil {
		return err
	}

	return nil
}

// ValidateName on exist. Accepts name
// and returns true, if name can be enriched.
func (e *Enricher) ValidateName(name string) bool {
	var pa PossibleAge
	if err := e.executeEnrichRequest(
		context.Background(), e.AgeURL,
		name, &pa); err != nil {
		return false
	}

	return pa.Age != nil
}
