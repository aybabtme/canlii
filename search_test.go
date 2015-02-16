package canlii_test

import (
	"encoding/json"
	"github.com/aybabtme/canlii"
	"reflect"
	"strings"
	"testing"
)

func TestSearchLegislation(t *testing.T) {
	resp := `{
  "resultCount": 609455,
  "results": [
    {
      "legislation": {
        "databaseId": "pes",
        "legislationId": "rspei-1988-c-l-7",
        "title": "Legislative Assembly Act",
        "citation": "RSPEI 1988, c L-7",
        "type": "STATUTE"
      }
    },
    {
      "case": {
        "databaseId": "bclrb",
        "caseId": {
          "en": "2003canlii62606"
        },
        "title": "British Columbia (Legislative Assembly) v. British Columbia",
        "citation": "2003 CanLII 62606 (BC LRB)"
      }
    },
    {
      "legislation": {
        "databaseId": "nlr",
        "legislationId": "cnlr-995-96",
        "title": "Subordinate Legislation Regulations",
        "citation": "CNLR 995/96",
        "type": "REGULATION"
      }
    },
    {
      "case": {
        "databaseId": "bcipc",
        "caseId": {
          "en": "2010bcipc6"
        },
        "title": "British Columbia (Attorney General) (Re)",
        "citation": "2010 BCIPC 6 (CanLII)"
      }
    },
    {
      "legislation": {
        "databaseId": "nls",
        "legislationId": "rsnl-1990-c-s-27",
        "title": "Statutes and Subordinate Legislation Act",
        "citation": "RSNL 1990, c S-27",
        "type": "STATUTE"
      }
    },
    {
      "case": {
        "databaseId": "nbqb",
        "caseId": {
          "en": "1995canlii10150"
        },
        "title": "New Brunswick  (Attorney General) v Morgentaler",
        "citation": "1995 CanLII 10150 (NB QB)"
      }
    },
    {
      "legislation": {
        "databaseId": "nls",
        "legislationId": "rsnl-1990-c-s-26",
        "title": "Statutes Act",
        "citation": "RSNL 1990, c S-26",
        "type": "STATUTE"
      }
    },
    {
      "legislation": {
        "databaseId": "qcr",
        "legislationId": "cqlr-c-r-9-r-30",
        "title": "Regulation respecting the implementation of the Agreement on Social Security between the Gouvernement du Québec and the Government of the Kingdom of Norway",
        "citation": "CQLR c R-9, r 30",
        "type": "REGULATION"
      }
    },
    {
      "legislation": {
        "databaseId": "nbs",
        "legislationId": "rsnb-1973-c-l-3",
        "title": "Legislative Assembly Act",
        "citation": "RSNB 1973, c L-3",
        "type": "STATUTE"
      }
    },
    {
      "legislation": {
        "databaseId": "qcr",
        "legislationId": "cqlr-c-r-9-r-26",
        "title": "Regulation respecting the implementation of the Agreement on Social Security between the Gouvernement du Québec and the Government of the Grand Duchy of Luxembourg",
        "citation": "CQLR c R-9, r 26",
        "type": "REGULATION"
      }
    }
  ]
}`
	var v struct {
		ResultCount int `json:"resultCount"`
		Results     []struct {
			Case        canlii.CaseListItem `json:"case"`
			Legislation canlii.Legislation  `json:"legislation"`
		} `json:"results"`
	}
	err := json.NewDecoder(strings.NewReader(resp)).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	wantLegis := canlii.SearchResult{
		TotalResults: v.ResultCount,
	}
	for _, r := range v.Results {
		switch {
		case r.Case.DatabaseID != "":
			wantLegis.Cases = append(wantLegis.Cases, r.Case)
		case r.Legislation.DatabaseID != "":
			wantLegis.Legislations = append(wantLegis.Legislations, r.Legislation)

		}
	}

	api := "derp"
	path := "/v1/search/en/"

	client, close := testClient(t, path, resp, api)
	defer close()

	gotLegis, _, err := client.Search.Search("legislation cases", &canlii.SearchOptions{
		ResultCount: 10,
	})
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantLegis, gotLegis) {
		t.Logf("want=%+v", wantLegis)
		t.Logf(" got=%+v", gotLegis)
		t.Error("mismatch yo")
	}
}
