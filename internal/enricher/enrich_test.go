package enricher

import (
	"context"
	"crypto/tls"
	"net/http"
	"testing"
	"time"
)

func TestEnricher_EnrichFIO(t *testing.T) {
	if !realTest {
		t.Skip()
	}

	enricher := New(&http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}, Timeout: time.Second * 5},
		ageBaseURL,
		genderBaseURL,
		countryBaseURL)

	reg, err := enricher.EnrichFIO(
		context.Background(),
		FIO{
			Name:       "Matvey",
			Surname:    "Sizov",
			Patronymic: new(string),
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", reg)
}

func TestEnricher_ValidateName(t *testing.T) {
	if !realTest {
		t.Skip()
	}

	enricher := New(&http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}, Timeout: time.Second * 5},
		ageBaseURL,
		genderBaseURL,
		countryBaseURL)

	if !enricher.ValidateName("Matvey") {
		t.Fatal("name must be valid")
	}
}
