package canlii

import (
	"fmt"
	"net/http"
)

type caseCitator struct {
	client *Client
}

// CitingCase gets a list of all cases citing the selected case.
// dbID is a unique identifier of a database as provided in the database list.
// caseID is a unique identifier of a case as provided in the case list.
func (c *caseCitator) CitingCases(dbID, caseID string) ([]Case, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/citingCases", dbID, caseID)
	var v struct {
		Cases []Case `json:"citingCases"`
	}
	resp, err := c.client.get("caseCitator", u, nil, &v)
	return v.Cases, resp, err
}

// CitedCase gets a list of all cases cited by the selected case.
// dbID is a unique identifier of a database as provided in the database list.
// caseID is a unique identifier of a case as provided in the case list.
func (c *caseCitator) CitedCases(dbID, caseID string) ([]Case, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/citedCases", dbID, caseID)
	var v struct {
		Cases []Case `json:"citedCases"`
	}
	resp, err := c.client.get("caseCitator", u, nil, &v)
	return v.Cases, resp, err
}

// CitedLegislation gets a list legislation cited by the selected case.
// dbID is a unique identifier of a database as provided in the database list.
// caseID is a unique identifier of a case as provided in the case list.
func (c *caseCitator) CitedLegislation(dbID, caseID string) ([]Legislation, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/citedLegislations", dbID, caseID)
	var v struct {
		Legislations []Legislation `json:"citedLegislations"`
	}
	resp, err := c.client.get("caseCitator", u, nil, &v)
	return v.Legislations, resp, err
}
