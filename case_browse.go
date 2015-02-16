package canlii

import (
	"net/http"
	"net/url"
	"time"
)

type caseBrowse struct {
	client *Client
}

type CaseDatabase struct {
	ID           string `json:"databaseId"`
	Jurisdiction string `json:"jurisdiction"`
	Name         string `json:"name"`
}

// ListDatabase lists all the databases supported by the API.
func (c *caseBrowse) ListDatabases() ([]CaseDatabase, *http.Response, error) {
	var v struct {
		DBs []CaseDatabase `json:"caseDatabases"`
	}
	resp, err := c.client.get("caseBrowse", "", nil, &v)
	return v.DBs, resp, err
}

type ListCaseOptions struct {
	Offset          int       // Starting record number for the list. First value is 0.
	ResultCount     int       // Number of cases listed in each response. Maximum value is 10 000.
	PublishedBefore time.Time // Cases published before a given date. The publication date is the first publication on CanLII. The date is inclusive.
	PublishedAfter  time.Time // Cases published after a given date. The publication date is the first publication on CanLII. The date is inclusive.
	DecisionBefore  time.Time // Cases with a decision date before a given date. The decision date is usually the date of the registraction by the docket.
	DecisionAfter   time.Time // Cases with a decision date after a given date. The decision date is usually the date of the registraction by the docket.
}

type Case struct {
	ID struct {
		EN string `json:"en"`
		FR string `json:"fr"`
	} `json:"caseID"`
	DatabaseID string `json:"databaseId"`
	Title      string `json:"title"`
	Citation   string `json:"citation"`
}

// ListCases lists all the cases within a database. dbID is a
// unique identifier of a database as provided in the database list.
func (c *caseBrowse) ListCases(dbID string, opts *ListCaseOptions) ([]Case, *http.Response, error) {
	const maxResultCount = 10000

	if opts == nil {
		opts = &ListCaseOptions{
			ResultCount: maxResultCount,
		}
	}
	if opts.ResultCount > maxResultCount {
		opts.ResultCount = maxResultCount
	}

	q := make(url.Values)
	setIntParam(q, "offset", opts.Offset)
	setIntParam(q, "resultCount", opts.ResultCount)
	setTimeParam(q, "publishedBefore", opts.PublishedBefore)
	setTimeParam(q, "publishedAfter", opts.PublishedAfter)
	setTimeParam(q, "decisionBefore", opts.DecisionBefore)
	setTimeParam(q, "decisionAfter", opts.DecisionAfter)

	var v struct {
		Cases []Case `json:"cases"`
	}
	resp, err := c.client.get("caseBrowse", dbID, q, &v)
	return v.Cases, resp, err
}

type CaseMetadata struct {
	DatabaseID   string    `json:"databaseID"`
	CaseID       string    `json:"caseID"`
	URL          string    `json:"url"`
	Title        string    `json:"title"`
	Citation     string    `json:"citation"`
	DocketNumber string    `json:"docketNumber"`
	DecisionDate time.Time `json:"decisionDate"`
}

// CaseMetadata gets the metadata for a given case. dbID is a
// unique identifier of a database as provided in the database list.
// caseID is a unique identifier of a case as provided in the case list.
func (c *caseBrowse) CaseMetadata(dbID, caseID string) ([]CaseMetadata, *http.Response, error) {
	var v struct {
		Cases []CaseMetadata `json:"cases"`
	}
	resp, err := c.client.get("caseBrowse", dbID+"/"+caseID, nil, &v)
	return v.Cases, resp, err
}
