// Package canlii offers a client for the HTTP API of CanLII, a Canadian legal database.
//
// About CanLII
//
// CanLII stands for Canadian Legal Information Institute. They provide
// an API to a cases and legislations database.
//
// You can register an account at: http://developer.canlii.org/member/register.
//
// You can obtain an API at: http://developer.canlii.org/apps/register.
//
// API Docs: http://developer.canlii.org/io-docs.
//
// http://www.canlii.org/en/index.html
package canlii

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	dateFormat = "2006-01-02"

	defaultApiHost  = "http://api.canlii.org"
	defaultVersion  = "v1"
	defaultLanguage = "en"
)

// Client to the CanLII API.
type Client struct {
	client *http.Client

	apiKey string

	// CanliiHost
	CanliiHost string
	APIVersion string
	Language   string

	BaseURL *url.URL

	CaseBrowse interface {
		ListDatabases() ([]CaseDatabase, *http.Response, error)
		ListCases(dbID string, opts *ListCaseOptions) ([]Case, *http.Response, error)
		CaseMetadata(dbID, caseID string) ([]CaseMetadata, *http.Response, error)
	}
	CaseCitator interface {
		CitingCases(dbID, caseID string) ([]Case, *http.Response, error)
		CitedCases(dbID, caseID string) ([]Case, *http.Response, error)
		CitedLegislation(dbID, caseID string) ([]Legislation, *http.Response, error)
	}
	CaseCitatorTease interface {
		CitingCases(dbID, caseID string) ([]Case, *http.Response, error)
		CitedCases(dbID, caseID string) ([]Case, *http.Response, error)
		CitedLegislation(dbID, caseID string) ([]Legislation, *http.Response, error)
	}
	LegislationBrowse interface {
		ListDatabases() ([]LegislationBrowseDatabase, *http.Response, error)
		ListLegislations(dbID string) ([]Legislation, *http.Response, error)
		LegislationMetadata(dbID, legislationID string) (Legislation, *http.Response, error)
	}
	Search interface {
		Search(fulltext string, opts *SearchOptions) (SearchResult, *http.Response, error)
	}
}

// NewClient to the API. `host` is optional and if empty, the default URL for
// the API is used. `apiKey` is mandatory and can be obtained at http://developer.canlii.org.
func NewClient(client *http.Client, host, apiKey string) (*Client, error) {
	if host == "" {
		host = defaultApiHost
	}
	u, err := url.Parse(host)
	if err != nil {
		return nil, fmt.Errorf("bad url for API host: %v", err)
	}

	if client == nil {
		client = http.DefaultClient
	}

	c := &Client{
		client: client,

		CanliiHost: host,
		APIVersion: defaultVersion,
		Language:   defaultLanguage,

		apiKey:  apiKey,
		BaseURL: u,
	}

	c.CaseBrowse = &caseBrowse{client: c}
	c.CaseCitator = &caseCitator{client: c}
	c.CaseCitatorTease = &caseCitatorTease{client: c}
	c.LegislationBrowse = &legislationBrowse{client: c}
	c.Search = &search{client: c}

	return c, nil
}

func (c *Client) get(res, urlStr string, params url.Values, v interface{}) (*http.Response, error) {
	rel, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.APIVersion, res, c.Language, urlStr))
	if err != nil {
		return nil, err
	}
	u := c.BaseURL.ResolveReference(rel)
	q := u.Query()
	for k, vals := range params {
		for _, v := range vals {
			q.Add(k, v)
		}
	}
	q.Set("api_key", c.apiKey)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := checkResponse(urlStr, resp); err != nil {
		return nil, err
	}

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

// checkResponse checks if the HTTP code was not in the 200 range.
func checkResponse(path string, r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorData := &Error{
		Status:   r.Status,
		Code:     r.StatusCode,
		Resource: path,
	}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		jerr := json.Unmarshal(data, &errorData.API)
		if jerr != nil {
			errorData.Cause = string(data)
		}
	}
	return errorData
}

type Error struct {
	Status   string
	Code     int
	Resource string
	Cause    string

	API []struct {
		ErrType string `json:"error"`
		Message string `json:"message"`
	}
}

func (e *Error) Error() string {
	buf := bytes.NewBuffer(nil)
	fmt.Fprintf(buf, "%s (%d) for resource %q:", e.Status, e.Code, e.Resource)
	if len(e.API) != 0 {
		for _, apiErr := range e.API {
			fmt.Fprintf(buf, " %s-%s,", apiErr.ErrType, apiErr.Message)
		}
	} else {
		fmt.Fprintf(buf, " %s", e.Cause)
	}
	return buf.String()
}

func setTimeParam(q url.Values, key string, t time.Time) {
	if t.IsZero() {
		return
	}
	q.Set(key, t.Format(dateFormat))
}

func setIntParam(q url.Values, key string, v int) {
	q.Set(key, strconv.Itoa(v))
}
