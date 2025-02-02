package apicall

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rcsolis/endoflife_server/internal/model"
)

// Constant for the base URL of the API
const baseURL = "https://endoflife.date/api/"

/**
 * GetAll fetches all the languages from the API
 */
func GetAll() ([]string, error) {
	url := fmt.Sprintf("%s%s.json", baseURL, "all")
	resp, err := http.Get(url)
	if err != nil {
		nerr := Throw(
			ServiceUnavailableErrorType,
			&APIError{msg: fmt.Sprintf("Error fetching data: %v", err)},
		)
		log.Println(nerr)
		return nil, nerr
	}
	// Close body
	defer resp.Body.Close()
	// Decode the response
	var allTechs []string
	err = json.NewDecoder(resp.Body).Decode(&allTechs)
	if err != nil {
		nerr := Throw(
			InternalServerErrorType,
			&APIError{msg: fmt.Sprintf("Error decoding response: %v", err)},
		)
		log.Println(nerr)
		return nil, nerr
	}

	return allTechs, nil
}

/**
 * GetAllDetails fetches all the cycle data for a specific language from the API
 */
func GetAllDetails(language string) ([]model.LanguageCycle, error) {
	// Create the url
	url := fmt.Sprintf("%s%s.json", baseURL, language)
	// Make the request
	resp, err := http.Get(url)
	if err != nil {
		nerr := Throw(
			ServiceUnavailableErrorType,
			&APIError{msg: fmt.Sprintf("Error fetching data: %v", err)},
		)
		log.Println(nerr)
		return nil, nerr
	}
	defer resp.Body.Close()
	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&model.RawResponse)
	if err != nil {
		nerr := Throw(
			InternalServerErrorType,
			&APIError{msg: fmt.Sprintf("Error decoding response: %v", err)},
		)
		log.Println(nerr)
		return nil, nerr
	}
	// Parse the response
	for _, value := range model.RawResponse {
		model.Details = append(model.Details, value.ParseJSON())
	}
	return model.Details, nil
}

/**
 * GetCycleDetails fetches the cycle data for a specific language and version
 * from the API
 */
func GetCycleDetails(language string, version string) (model.LanguageCycle, error) {
	var rawResponse model.RawLanguageCycle = model.RawLanguageCycle{}
	// Create the url
	url := fmt.Sprintf("%s%s/%s.json", baseURL, language, version)
	// Make the request
	resp, err := http.Get(url)
	if err != nil {
		nerr := Throw(
			ServiceUnavailableErrorType,
			&APIError{msg: fmt.Sprintf("Error fetching data: %v", err)},
		)
		log.Println(nerr)
		return model.LanguageCycle{}, nerr
	}
	defer resp.Body.Close()
	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&rawResponse)
	if err != nil {
		nerr := Throw(
			InternalServerErrorType,
			&APIError{msg: fmt.Sprintf("Error decoding response: %v", err)},
		)
		log.Println(nerr)
		return model.LanguageCycle{}, nerr
	}
	// Parse the response
	return rawResponse.ParseJSON(), nil
}
