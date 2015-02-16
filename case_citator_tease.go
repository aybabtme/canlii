package canlii

import (
	"fmt"
	"net/http"
)

type caseCitatorTease struct {
	client *Client
}

// CitingCases gets a maximim of 5 cases citing the selected case.
// dbID is a unique identifier of a database as provided in the database list.
// caseID is a unique identifier of a case as provided in the case list.
func (c *caseCitatorTease) CitingCases(dbID, caseID string) ([]Case, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/citingCases", dbID, caseID)
	var v struct {
		Cases []Case `json:"citingCases"`
	}
	resp, err := c.client.Get("caseCitatorTease", u, nil, &v)
	return v.Cases, resp, err
}

// CitedCases gets a maximim of 5 cases cited by the selected case.
// dbID is a unique identifier of a database as provided in the database list.
// caseID is a unique identifier of a case as provided in the case list.
func (c *caseCitatorTease) CitedCases(dbID, caseID string) ([]Case, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/citedCases", dbID, caseID)
	var v struct {
		Cases []Case `json:"citedCases"`
	}
	resp, err := c.client.Get("caseCitatorTease", u, nil, &v)
	return v.Cases, resp, err
}

// CitedLegislation gets a maximim of 5 legislations cite by the selected case.
// dbID is a unique identifier of a database as provided in the database list.
// caseID is a unique identifier of a case as provided in the case list.
func (c *caseCitatorTease) CitedLegislation(dbID, caseID string) ([]Legislation, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/citedLegislations", dbID, caseID)
	var v struct {
		Legislations []Legislation `json:"citedLegislations"`
	}
	resp, err := c.client.Get("caseCitatorTease", u, nil, &v)
	return v.Legislations, resp, err
}
