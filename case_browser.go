package canlii

import (
	"net/http"
	"net/url"
	"time"
)

type caseBrowser struct {
	client *Client
}

type CaseBrowserDatabase struct {
	ID           string `json:"databaseId"`
	Jurisdiction string `json:"jurisdiction"`
	Name         string `json:"name"`
}

func (c *caseBrowser) ListDatabases() ([]CaseBrowserDatabase, *http.Response, error) {
	var v struct {
		DBs []CaseBrowserDatabase `json:"caseDatabases"`
	}
	resp, err := c.client.Get("caseBrowse", "", nil, &v)
	return v.DBs, resp, err
}

type ListCaseOptions struct {
	Offset          int
	ResultCount     int
	PublishedBefore time.Time
	PublishedAfter  time.Time
	DecisionBefore  time.Time
	DecisionAfter   time.Time
}

type CaseListItem struct {
	ID struct {
		EN string `json:"en"`
		FR string `json:"fr"`
	} `json:"caseID"`
	DatabaseID string `json:"databaseId"`
	Title      string `json:"title"`
	Citation   string `json:"citation"`
}

func (c *caseBrowser) ListCases(dbID string, opts *ListCaseOptions) ([]CaseListItem, *http.Response, error) {
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
		Cases []CaseListItem `json:"cases"`
	}
	resp, err := c.client.Get("caseBrowse", dbID, q, &v)
	return v.Cases, resp, err
}

type Case struct {
	DatabaseID   string    `json:"databaseID"`
	CaseID       string    `json:"caseID"`
	URL          string    `json:"url"`
	Title        string    `json:"title"`
	Citation     string    `json:"citation"`
	DocketNumber string    `json:"docketNumber"`
	DecisionDate time.Time `json:"decisionDate"`
}

func (c *caseBrowser) CaseMetadata(dbID, caseID string) ([]Case, *http.Response, error) {
	var v struct {
		Cases []Case `json:"cases"`
	}
	resp, err := c.client.Get("caseBrowse", dbID+"/"+caseID, nil, &v)
	return v.Cases, resp, err
}
