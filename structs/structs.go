package structs

import (
    "time"
)

type Mdocs []struct {
	Title   string `json:"title"`
	Type    string `json:"type"`
	Authors []struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"authors"`
	Year        int    `json:"year"`
	Source      string `json:"source"`
	Identifiers struct {
		Doi  string `json:"doi"`
		Issn string `json:"issn"`
	} `json:"identifiers"`
	ID           string    `json:"id"`
	Created      time.Time `json:"created"`
	ProfileID    string    `json:"profile_id"`
	GroupID      string    `json:"group_id"`
	LastModified time.Time `json:"last_modified"`
	Abstract     string    `json:"abstract"`
}

type Mdoc struct {
	ProfileID      string   `json:"profile_id"`
	GroupID        string   `json:"group_id"`
	LastModified   string   `json:"last_modified"`
	Tags           []string `json:"tags"`
	Read           bool     `json:"read"`
	Starred        bool     `json:"starred"`
	Authored       bool     `json:"authored"`
	Confirmed      bool     `json:"confirmed"`
	Hidden         bool     `json:"hidden"`
	CitationKey    string   `json:"citation_key"`
	SourceType     string   `json:"source_type"`
	Language       string   `json:"language"`
	ShortTitle     string   `json:"short_title"`
	ReprintEdition string   `json:"reprint_edition"`
	Genre          string   `json:"genre"`
	Country        string   `json:"country"`
	Translators    []struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"translators"`
	SeriesEditor            string `json:"series_editor"`
	Code                    string `json:"code"`
	Medium                  string `json:"medium"`
	UserContext             string `json:"user_context"`
	PatentOwner             string `json:"patent_owner"`
	PatentApplicationNumber string `json:"patent_application_number"`
	PatentLegalStatus       string `json:"patent_legal_status"`
	Notes                   string `json:"notes"`
	Accessed                string `json:"accessed"`
	FileAttached            bool   `json:"file_attached"`
	Created                 string `json:"created"`
	ID                      string `json:"id"`
	Year                    int    `json:"year"`
	Month                   int    `json:"month"`
	Day                     int    `json:"day"`
	Source                  string `json:"source"`
	Edition                 string `json:"edition"`
	Authors                 []struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"authors"`
	Keywords     []string `json:"keywords"`
	Pages        string   `json:"pages"`
	Volume       string   `json:"volume"`
	Issue        string   `json:"issue"`
	Websites     []string `json:"websites"`
	Publisher    string   `json:"publisher"`
	City         string   `json:"city"`
	Institution  string   `json:"institution"`
	Department   string   `json:"department"`
	Series       string   `json:"series"`
	SeriesNumber string   `json:"series_number"`
	Chapter      string   `json:"chapter"`
	Editors      []struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"editors"`
	Title       string `json:"title"`
	Revision    string `json:"revision"`
	Identifiers string `json:"identifiers"`
	Abstract    string `json:"abstract"`
	Type        string `json:"type"`
}
