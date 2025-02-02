package model

import (
	"fmt"
	"log"
)

type RawLanguageCycle struct {
	Cycle           string `json:"cycle"`
	ReleaseDate     string `json:"releaseDate"`
	Eol             any    `json:"eol"`
	Latest          string `json:"latest"`
	Link            string `json:"link,omitempty"`
	Lts             any    `json:"lts"`
	Support         any    `json:"support,omitempty"`
	Discontinued    any    `json:"discontinued,omitempty"`
	ExtendedSupport any    `json:"extendedSupport,omitempty"`
}

// Method to parse the JSON response to a struct with string fields
func (r RawLanguageCycle) ParseJSON() LanguageCycle {
	var (
		feol, flts, fsupport, fdiscontinued, fextendedSupport string
	)
	switch r.Eol.(type) {
	case string:
		feol = r.Eol.(string)
	case bool:
		if r.Eol.(bool) {
			feol = "True"
		} else {
			feol = "False"
		}
	default:
		feol = "Unknown"
	}
	switch r.Lts.(type) {
	case string:
		flts = r.Lts.(string)
	case bool:
		if r.Lts.(bool) {
			flts = "True"
		} else {
			flts = "False"
		}
	default:
		flts = "Unknown"
	}
	switch r.Support.(type) {
	case string:
		fsupport = r.Support.(string)
	case bool:
		if r.Support.(bool) {
			fsupport = "True"
		} else {
			fsupport = "False"
		}
	default:
		fsupport = "Unknown"
	}
	switch r.Discontinued.(type) {
	case string:
		fdiscontinued = r.Discontinued.(string)
	case bool:
		if r.Discontinued.(bool) {
			fdiscontinued = "True"
		} else {
			fdiscontinued = "False"
		}
	default:
		fdiscontinued = "Unknown"
	}
	switch r.ExtendedSupport.(type) {
	case string:
		fextendedSupport = r.ExtendedSupport.(string)
	case bool:
		if r.ExtendedSupport.(bool) {
			fextendedSupport = "True"
		} else {
			fextendedSupport = "False"
		}
	default:
		fextendedSupport = "Unknown"
	}

	return LanguageCycle{
		Cycle:           r.Cycle,
		ReleaseDate:     r.ReleaseDate,
		Eol:             feol,
		Latest:          r.Latest,
		Link:            r.Link,
		Lts:             flts,
		Support:         fsupport,
		Discontinued:    fdiscontinued,
		ExtendedSupport: fextendedSupport,
	}
}

// LanguageCycle struct for standardizing the data
type LanguageCycle struct {
	Cycle           string `json:"cycle"`
	ReleaseDate     string `json:"releaseDate"`
	Eol             string `json:"eol"`
	Latest          string `json:"latest"`
	Link            string `json:"link"`
	Lts             string `json:"lts"`
	Support         string `json:"support"`
	Discontinued    string `json:"discontinued"`
	ExtendedSupport string `json:"extendedSupport"`
}

// Response from API
var RawResponse []RawLanguageCycle

// Details after parsing the response
var Details []LanguageCycle

func init() {
	RawResponse = make([]RawLanguageCycle, 0)
	Details = make([]LanguageCycle, 0)
}

func PrintDetails(element LanguageCycle) {
	formattedString := fmt.Sprintf(`
	Cycle: %v
	Release Date: %v
	EOL: %v
	Latest: %v
	Link: %v
	LTS: %v
	Support: %v
	Discontinued: %v
	Extended Support: %v
	`, element.Cycle,
		element.ReleaseDate,
		element.Eol,
		element.Latest,
		element.Link,
		element.Lts,
		element.Support,
		element.Discontinued,
		element.ExtendedSupport)
	log.Printf("Details: %v\n", formattedString)
}
