package canlii_test

import (
	"encoding/json"
	"github.com/aybabtme/canlii"
	"reflect"
	"strings"
	"testing"
)

func TestCaseCitatorCitingCases(t *testing.T) {
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
		DBs []canlii.Case `json:"citingCases"`
	}
	err := json.NewDecoder(strings.NewReader(resp)).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	wantDB := v.DBs
	api := "derp"
	path := "/v1/caseCitator/en/csc-scc/2011scc47/citingCases"

	client, close := testClient(t, path, resp, api)
	defer close()

	gotDBs, _, err := client.CaseCitator.CitingCases("csc-scc", "2011scc47")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantDB, gotDBs) {
		t.Logf("want=%+v", wantDB)
		t.Logf(" got=%+v", gotDBs)
		t.Error("mismatch yo")
	}
}

func TestCaseCitatorCitedCases(t *testing.T) {
	resp := `{
  "citedCases": [
    {
      "databaseId": "onca",
      "caseId": {
        "en": "2004canlii12938"
      },
      "title": "Barrick Gold Corp. v. Lopehandia",
      "citation": "2004 CanLII 12938 (ON CA)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "1995canlii60"
      },
      "title": "Botiuk v. Toronto Free Press Publications Ltd.",
      "citation": "[1995] 3 SCR 3, 1995 CanLII 60 (SCC)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "2011scc9"
      },
      "title": "Bou Malhab v. Diffusion Métromédia CMR inc.",
      "citation": "[2011] 1 SCR 214, 2011 SCC 9 (CanLII)"
    },
    {
      "databaseId": "nsca",
      "caseId": {
        "en": "2001nsca121"
      },
      "title": "Butler v. Southam Inc.",
      "citation": "2001 NSCA 121 (CanLII)"
    },
    {
      "databaseId": "bcca",
      "caseId": {
        "en": "2005bcca398"
      },
      "title": "Carter v. B.C. Federation of Foster Parents Assn.",
      "citation": "2005 BCCA 398 (CanLII)"
    },
    {
      "databaseId": "bcca",
      "caseId": {
        "en": "2009bcca392"
      },
      "title": "Crookes v. Newton",
      "citation": "2009 BCCA 392 (CanLII)"
    },
    {
      "databaseId": "bcsc",
      "caseId": {
        "en": "2008bcsc1424"
      },
      "title": "Crookes v. Wikimedia Foundation Inc.",
      "citation": "2008 BCSC 1424 (CanLII)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "2007scc34"
      },
      "title": "Dell Computer Corp. v. Union des consommateurs",
      "citation": "[2007] 2 SCR 801, 2007 SCC 34 (CanLII)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "1965canlii8"
      },
      "title": "Gaskin v. Retail Credit Co. et al",
      "citation": "[1965] SCR 297, 1965 CanLII 8 (SCC)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "2009scc61"
      },
      "title": "Grant v. Torstar Corp.",
      "citation": "[2009] 3 SCR 640, 2009 SCC 61 (CanLII)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "1995canlii59"
      },
      "title": "Hill v. Church of Scientology of Toronto",
      "citation": "[1995] 2 SCR 1130, 1995 CanLII 59 (SCC)"
    },
    {
      "databaseId": "nssc",
      "caseId": {
        "en": "1997canlii542"
      },
      "title": "Hiltz and Seamone Co. Ltd. v. Nova Scotia (Attorney General)",
      "citation": "1997 CanLII 542 (NS SC)"
    },
    {
      "databaseId": "nsca",
      "caseId": {
        "en": "1999canlii13144"
      },
      "title": "Hiltz and Seamone Co. Ltd. v. Nova Scotia (Attorney General) et al.",
      "citation": "1999 CanLII 13144 (NS CA)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "1931canlii99"
      },
      "title": "McNichol v. Grandy",
      "citation": "[1931] SCR 696, 1931 CanLII 99 (SCC)"
    },
    {
      "databaseId": "bcsc",
      "caseId": {
        "en": "1986canlii1117"
      },
      "title": "Smith v. Matsqui (Dist.)",
      "citation": "1986 CanLII 1117 (BC SC)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "2004scc45"
      },
      "title": "Society of Composers, Authors and Music Publishers of Canada v. Canadian Assn. of Internet Providers",
      "citation": "[2004] 2 SCR 427, 2004 SCC 45 (CanLII)"
    },
    {
      "databaseId": "bcca",
      "caseId": {
        "en": "2006bcca467"
      },
      "title": "Stanley v. Shaw and Tracey",
      "citation": "2006 BCCA 467 (CanLII)"
    },
    {
      "databaseId": "onca",
      "caseId": {
        "en": "1937canlii14"
      },
      "title": "Thomson v. Lambert",
      "citation": "1937 CanLII 14 (ON CA)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "2008scc40"
      },
      "title": "WIC Radio Ltd. v. Simpson",
      "citation": "[2008] 2 SCR 420, 2008 SCC 40 (CanLII)"
    }
  ]
}`
	var v struct {
		Cases []canlii.Case `json:"citedCases"`
	}
	err := json.NewDecoder(strings.NewReader(resp)).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	wantCases := v.Cases
	api := "derp"
	path := "/v1/caseCitator/en/csc-scc/2011scc47/citedCases"

	client, close := testClient(t, path, resp, api)
	defer close()

	gotCases, _, err := client.CaseCitator.CitedCases("csc-scc", "2011scc47")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantCases, gotCases) {
		t.Logf("want=%+v", wantCases)
		t.Logf(" got=%+v", gotCases)
		t.Error("mismatch yo")
	}
}

func TestCaseCitatorCitedLegislation(t *testing.T) {
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
      "databaseId": "bcs",
      "legislationId": "rsbc-1996-c-263",
      "title": "Libel and Slander Act",
      "citation": "RSBC 1996, c 263",
      "type": "STATUTE"
    },
    {
      "databaseId": "bcr",
      "legislationId": "bc-reg-221-90",
      "title": "Supreme Court Rules",
      "citation": "BC Reg 221/90",
      "type": "REGULATION"
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
	path := "/v1/caseCitator/en/csc-scc/2011scc47/citedLegislations"

	client, close := testClient(t, path, resp, api)
	defer close()

	gotLegis, _, err := client.CaseCitator.CitedLegislation("csc-scc", "2011scc47")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantLegis, gotLegis) {
		t.Logf("want=%+v", wantLegis)
		t.Logf(" got=%+v", gotLegis)
		t.Error("mismatch yo")
	}
}
