package canlii

import (
	"fmt"
	"net/http"
)

type caseCitatorTease struct {
	client *Client
}

func (c *caseCitatorTease) CitingCases(dbID, caseID string) ([]CaseListItem, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/citingCases", dbID, caseID)
	var v struct {
		Cases []CaseListItem `json:"citingCases"`
	}
	resp, err := c.client.Get("caseCitatorTease", u, nil, &v)
	return v.Cases, resp, err
}

func (c *caseCitatorTease) CitedCases(dbID, caseID string) ([]CaseListItem, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/citedCases", dbID, caseID)
	var v struct {
		Cases []CaseListItem `json:"citedCases"`
	}
	resp, err := c.client.Get("caseCitatorTease", u, nil, &v)
	return v.Cases, resp, err
}

func (c *caseCitatorTease) CitedLegislation(dbID, caseID string) ([]Legislation, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/citedLegislations", dbID, caseID)
	var v struct {
		Legislations []Legislation `json:"citedLegislations"`
	}
	resp, err := c.client.Get("caseCitatorTease", u, nil, &v)
	return v.Legislations, resp, err
}
