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

func (c *legislationBrowse) ListDatabases() ([]LegislationBrowseDatabase, *http.Response, error) {
	var v struct {
		DBs []LegislationBrowseDatabase `json:"legislationDatabases"`
	}
	resp, err := c.client.Get("legislationBrowse", "", nil, &v)
	return v.DBs, resp, err
}

type Legislation struct {
	DatabaseID    string `json:"databaseID"`
	LegislationID string `json:"legislationID"`
	Title         string `json:"title"`
	Citation      string `json:"citation"`
	Type          string `json:"type"`
}

func (c *legislationBrowse) ListLegislations(dbID string) ([]Legislation, *http.Response, error) {
	var v struct {
		Legislations []Legislation `json:"legislations"`
	}
	resp, err := c.client.Get("legislationBrowse", dbID, nil, &v)
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

func (c *legislationBrowse) LegislationMetadata(dbID, legislationID string) (Legislation, *http.Response, error) {
	legis := Legislation{}
	resp, err := c.client.Get("legislationBrowse", dbID+"/"+legislationID, nil, &legis)
	return legis, resp, err
}
