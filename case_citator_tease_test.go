package canlii_test

import (
	"encoding/json"
	"github.com/aybabtme/canlii"
	"reflect"
	"strings"
	"testing"
)

func TestCaseCitatorTeaseCitingCases(t *testing.T) {
	resp := `{
  "citingCases": [
    {
      "databaseId": "abqb",
      "caseId": {
        "en": "2012abqb559"
      },
      "title": "Kent v. Postmedia Network Inc.",
      "citation": "2012 ABQB 559 (CanLII)"
    },
    {
      "databaseId": "qccs",
      "caseId": {
        "fr": "2012qccs3078"
      },
      "title": "Laforest c. Collins",
      "citation": "2012 QCCS 3078 (CanLII)"
    },
    {
      "databaseId": "qcca",
      "caseId": {
        "fr": "2013qcca1482"
      },
      "title": "Lavoie c. Vailles",
      "citation": "2013 QCCA 1482 (CanLII)"
    },
    {
      "databaseId": "qccs",
      "caseId": {
        "fr": "2012qccs1464"
      },
      "title": "9080-5128 Québec inc. c. Morin-Ogilvy",
      "citation": "2012 QCCS 1464 (CanLII)"
    },
    {
      "databaseId": "qccq",
      "caseId": {
        "fr": "2012qccq1354"
      },
      "title": "Chayer c. Vaillancourt",
      "citation": "2012 QCCQ 1354 (CanLII)"
    },
    {
      "databaseId": "onca",
      "caseId": {
        "en": "2012onca892"
      },
      "title": "R. v. Smith",
      "citation": "2012 ONCA 892 (CanLII)"
    },
    {
      "databaseId": "bcsc",
      "caseId": {
        "en": "2012bcsc205"
      },
      "title": "Nazerali v. Mitchell",
      "citation": "2012 BCSC 205 (CanLII)"
    },
    {
      "databaseId": "qccq",
      "caseId": {
        "fr": "2013qccq12382"
      },
      "title": "Maltais c. Saunders Gordon",
      "citation": "2013 QCCQ 12382 (CanLII)"
    },
    {
      "databaseId": "onsc",
      "caseId": {
        "en": "2011onsc6784"
      },
      "title": "Elfarnawani v. International Olympic Committee",
      "citation": "2011 ONSC 6784 (CanLII)"
    },
    {
      "databaseId": "onsc",
      "caseId": {
        "en": "2012onsc7189"
      },
      "title": "Foulidis v. Ford",
      "citation": "2012 ONSC 7189 (CanLII)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "2012scc18"
      },
      "title": "Éditions Écosociété Inc. v. Banro Corp.",
      "citation": "[2012] 1 SCR 636, 2012 SCC 18 (CanLII)"
    },
    {
      "databaseId": "abqb",
      "caseId": {
        "en": "2012abqb640"
      },
      "title": "Court v. Debaie",
      "citation": "2012 ABQB 640 (CanLII)"
    },
    {
      "databaseId": "onsc",
      "caseId": {
        "en": "2013onsc1716"
      },
      "title": "Asselin v. McDougall",
      "citation": "2013 ONSC 1716 (CanLII)"
    },
    {
      "databaseId": "bcca",
      "caseId": {
        "en": "2014bcca448"
      },
      "title": "Equustek Solutions Inc. v. Google Inc.",
      "citation": "2014 BCCA 448 (CanLII)"
    },
    {
      "databaseId": "qccs",
      "caseId": {
        "fr": "2013qccs6286"
      },
      "title": "Rosenberg c. Lacerte",
      "citation": "2013 QCCS 6286 (CanLII)"
    },
    {
      "databaseId": "mbpc",
      "caseId": {
        "en": "2014mbpc12"
      },
      "title": "R v Beaudry",
      "citation": "2014 MBPC 12 (CanLII)"
    },
    {
      "databaseId": "bcsc",
      "caseId": {
        "en": "2014bcsc1601"
      },
      "title": "Popat v. MacLennan",
      "citation": "2014 BCSC 1601 (CanLII)"
    },
    {
      "databaseId": "bcca",
      "caseId": {
        "en": "2013bcca341"
      },
      "title": "Mainstream Canada v. Staniford",
      "citation": "2013 BCCA 341 (CanLII)"
    },
    {
      "databaseId": "abqb",
      "caseId": {
        "en": "2012abqb507"
      },
      "title": "Kent v. Martin",
      "citation": "2012 ABQB 507 (CanLII)"
    },
    {
      "databaseId": "nlsctd",
      "caseId": {
        "en": "2013canlii15005"
      },
      "title": "Harvey v. Memorial University of Newfoundland",
      "citation": "2013 CanLII 15005 (NL SCTD)"
    },
    {
      "databaseId": "bcsc",
      "caseId": {
        "en": "2012bcsc299"
      },
      "title": "Tjelta v. Wang",
      "citation": "2012 BCSC 299 (CanLII)"
    },
    {
      "databaseId": "abqb",
      "caseId": {
        "en": "2013abqb332"
      },
      "title": "Gouin v. White",
      "citation": "2013 ABQB 332 (CanLII)"
    },
    {
      "databaseId": "bcsc",
      "caseId": {
        "en": "2012bcsc693"
      },
      "title": "Gichuru v. Pallai",
      "citation": "2012 BCSC 693 (CanLII)"
    },
    {
      "databaseId": "bcsc",
      "caseId": {
        "en": "2015bcsc165"
      },
      "title": "Weaver v. Corcoran",
      "citation": "2015 BCSC 165 (CanLII)"
    },
    {
      "databaseId": "bcsc",
      "caseId": {
        "en": "2013bcsc2111"
      },
      "title": "Demenuk v. Dhadwal",
      "citation": "2013 BCSC 2111 (CanLII)"
    },
    {
      "databaseId": "onsc",
      "caseId": {
        "en": "2015onsc155"
      },
      "title": "Bernstein v. Poon",
      "citation": "2015 ONSC 155 (CanLII)"
    }
  ]
}`
	var v struct {
		DBs []canlii.CaseListItem `json:"citingCases"`
	}
	err := json.NewDecoder(strings.NewReader(resp)).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	wantDB := v.DBs
	api := "derp"
	path := "/v1/caseCitator/en/csc-scc/2008scc9/citingCases"

	client, close := testClient(t, path, resp, api)
	defer close()

	gotDBs, _, err := client.CaseCitator.CitingCases("csc-scc", "2008scc9")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantDB, gotDBs) {
		t.Logf("want=%+v", wantDB)
		t.Logf(" got=%+v", gotDBs)
		t.Error("mismatch yo")
	}
}

func TestCaseCitatorTeaseCitedCases(t *testing.T) {
	resp := `{
  "citedCases": [
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "1974canlii12"
      },
      "title": "McLeod v. Egan",
      "citation": "[1975] 1 SCR 517, 1974 CanLII 12 (SCC)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "1978canlii24"
      },
      "title": "Nicholson v. Haldimand-Norfolk Regional Police Commissioners",
      "citation": "[1979] 1 SCR 311, 1978 CanLII 24 (SCC)"
    },
    {
      "databaseId": "bcsc",
      "caseId": {
        "en": "2004bcsc790"
      },
      "title": "Reglin v Creston (Town) et al.",
      "citation": "2004 BCSC 790 (CanLII)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "1959canlii50"
      },
      "title": "Roncarelli v. Duplessis",
      "citation": "[1959] SCR 121, 1959 CanLII 50 (SCC)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "1972canlii139"
      },
      "title": "Woodward Estate v. Minister of Finance",
      "citation": "[1973] SCR 120, 1972 CanLII 139 (SCC)"
    }
  ]
}`
	var v struct {
		Cases []canlii.CaseListItem `json:"citedCases"`
	}
	err := json.NewDecoder(strings.NewReader(resp)).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	wantCases := v.Cases
	api := "derp"
	path := "/v1/caseCitator/en/csc-scc/2008scc9/citedCases"

	client, close := testClient(t, path, resp, api)
	defer close()

	gotCases, _, err := client.CaseCitator.CitedCases("csc-scc", "2008scc9")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantCases, gotCases) {
		t.Logf("want=%+v", wantCases)
		t.Logf(" got=%+v", gotCases)
		t.Error("mismatch yo")
	}
}

func TestCaseCitatorTeaseCitedLegislation(t *testing.T) {
	resp := `{
  "citedLegislations": [
    {
      "databaseId": "qcs",
      "legislationId": "lrq-c-c-1991",
      "title": "Civil Code of Québec",
      "citation": "LRQ, c C-1991",
      "type": "STATUTE"
    },
    {
      "databaseId": "cas",
      "legislationId": "30---31-vict-c-3",
      "title": "The Constitution Act, 1867",
      "citation": "30 \u0026 31 Vict, c 3",
      "type": "STATUTE"
    },
    {
      "databaseId": "nbs",
      "legislationId": "snb-1982-c-e-7.2",
      "title": "Employment Standards Act",
      "citation": "SNB 1982, c E-7.2",
      "type": "STATUTE"
    },
    {
      "databaseId": "nbs",
      "legislationId": "rsnb-1973-c-h-11",
      "title": "Human Rights Act",
      "citation": "RSNB 1973, c H-11",
      "type": "STATUTE"
    },
    {
      "databaseId": "nbs",
      "legislationId": "rsnb-1973-c-p-25",
      "title": "Public Service Labour Relations Act",
      "citation": "RSNB 1973, c P-25",
      "type": "STATUTE"
    }
  ]
}`
	var v struct {
		Legis []canlii.Legislation `json:"citedLegislations"`
	}
	err := json.NewDecoder(strings.NewReader(resp)).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	wantLegis := v.Legis
	api := "derp"
	path := "/v1/caseCitator/en/csc-scc/2008scc9/citedLegislations"

	client, close := testClient(t, path, resp, api)
	defer close()

	gotLegis, _, err := client.CaseCitator.CitedLegislation("csc-scc", "2008scc9")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantLegis, gotLegis) {
		t.Logf("want=%+v", wantLegis)
		t.Logf(" got=%+v", gotLegis)
		t.Error("mismatch yo")
	}
}
