package canlii

import (
	"net/http"
	"time"
)

type legislationBrowse struct {
	client *Client
}

type LegislationBrowseDatabase struct {
	ID           string `json:"databaseId"`
	Type         string `json:"type"`
	Jurisdiction string `json:"jurisdiction"`
	Name         string `json:"name"`
}

// ListDatabases lists all the databases supported by the API
func (c *legislationBrowse) ListDatabases() ([]LegislationBrowseDatabase, *http.Response, error) {
	var v struct {
		DBs []LegislationBrowseDatabase `json:"legislationDatabases"`
	}
	resp, err := c.client.get("legislationBrowse", "", nil, &v)
	return v.DBs, resp, err
}

type Legislation struct {
	DatabaseID    string `json:"databaseID"`
	LegislationID string `json:"legislationID"`
	Title         string `json:"title"`
	Citation      string `json:"citation"`
	Type          string `json:"type"`
}

// ListLegislations lists all the legislation within a database.
// dbID is a unique identifier of a database as provided in the database list.
func (c *legislationBrowse) ListLegislations(dbID string) ([]Legislation, *http.Response, error) {
	var v struct {
		Legislations []Legislation `json:"legislations"`
	}
	resp, err := c.client.get("legislationBrowse", dbID, nil, &v)
	return v.Legislations, resp, err
}

type LegislationMetadata struct {
	DatabaseID    string    `json:"databaseID"`
	LegislationID string    `json:"legislationID"`
	Title         string    `json:"title"`
	Citation      string    `json:"citation"`
	Type          string    `json:"type"`
	Language      string    `json:"language"`
	DataScheme    string    `json:"dateScheme"`
	StartDate     time.Time `json:"startDate"`
	EndDate       time.Time `json:"endDate"`
	Repealed      string    `json:"repealed"`
	Content       []struct {
		PartID   string `json:"partID"`
		PartName string `json:"partName"`
	} `json:"content"`
}

// LegislationMetadata gets the metadata for a given legislation.
// dbID is a unique identifier of a database as provided in the database list.
// legislationID is a unique identifier of a legislation as provided in the legislation list.
func (c *legislationBrowse) LegislationMetadata(dbID, legislationID string) (Legislation, *http.Response, error) {
	legis := Legislation{}
	resp, err := c.client.get("legislationBrowse", dbID+"/"+legislationID, nil, &legis)
	return legis, resp, err
}
