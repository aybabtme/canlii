package canlii_test

import (
	"encoding/json"
	"github.com/aybabtme/canlii"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func testClient(t *testing.T, urlPath, resp, api string) (*canlii.Client, func()) {
	srv := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.URL.Path == urlPath {
			rw.WriteHeader(200)
			rw.Write([]byte(resp))
		} else {
			t.Fatalf("\n got=%q\nwant=%q", req.URL.Path, urlPath)
		}
	}))

	client, err := canlii.NewClient(nil, srv.URL, api)
	if err != nil {
		t.Fatal(err)
	}

	return client, func() { srv.Close() }
}

func TestCaseBrowseDatabaseList(t *testing.T) {
	resp := `{
  "caseDatabases": [
    {
      "databaseId": "abwcac",
      "jurisdiction": "ab",
      "name": "Appeals Commission for Alberta Workers\u0027 Compensation"
    },
    {
      "databaseId": "nssirt",
      "jurisdiction": "ns",
      "name": "Nova Scotia Serious Incident Response Team"
    },
    {
      "databaseId": "qcoarq",
      "jurisdiction": "qc",
      "name": "Comité de discipline de l\u0027Ordre des architectes du Québec"
    },
    {
      "databaseId": "abgaa",
      "jurisdiction": "ab",
      "name": "Alberta Grievance Arbitration Awards"
    },
    {
      "databaseId": "ondr",
      "jurisdiction": "on",
      "name": "Ontario Court of the Drainage Referee"
    },
    {
      "databaseId": "qccptaq",
      "jurisdiction": "qc",
      "name": "Commission de protection du territoire agricole du Quebec"
    },
    {
      "databaseId": "pesctd",
      "jurisdiction": "pe",
      "name": "Supreme Court of Prince Edward Island"
    },
    {
      "databaseId": "oncece",
      "jurisdiction": "on",
      "name": "Ontario College of Early Childhood Educators"
    },
    {
      "databaseId": "sst-tss",
      "jurisdiction": "ca",
      "name": "Social Security Tribunal of Canada"
    },
    {
      "databaseId": "onpprb",
      "jurisdiction": "on",
      "name": "Ontario Physician Payment Review Board"
    },
    {
      "databaseId": "onpsgb",
      "jurisdiction": "on",
      "name": "Ontario Public Service Grievance Board"
    },
    {
      "databaseId": "nttc",
      "jurisdiction": "nt",
      "name": "Territorial Court of the Northwest Territories"
    },
    {
      "databaseId": "bcrec",
      "jurisdiction": "bc",
      "name": "Real Estate Council of British Columbia"
    },
    {
      "databaseId": "csc-scc",
      "jurisdiction": "ca",
      "name": "Supreme Court of Canada"
    },
    {
      "databaseId": "yksm",
      "jurisdiction": "yk",
      "name": "Small Claims Court of the Yukon"
    },
    {
      "databaseId": "pelrb",
      "jurisdiction": "pe",
      "name": "Prince Edward Island Labour Relations Board"
    },
    {
      "databaseId": "yksc",
      "jurisdiction": "yk",
      "name": "Supreme Court  of Yukon"
    },
    {
      "databaseId": "qccdccoq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline des Conseillers et conseillères d\u0027orientation du Québec"
    },
    {
      "databaseId": "bcwcat",
      "jurisdiction": "bc",
      "name": "British Columbia Workers\u0027 Compensation Appeal Tribunal"
    },
    {
      "databaseId": "qccdchad",
      "jurisdiction": "qc",
      "name": "Chambre de l\u0027assurance de dommages Discipline Committee"
    },
    {
      "databaseId": "citt-tcce",
      "jurisdiction": "ca",
      "name": "Canadian International Trade Tribunal"
    },
    {
      "databaseId": "oncj",
      "jurisdiction": "on",
      "name": "Ontario Court of Justice"
    },
    {
      "databaseId": "ntro",
      "jurisdiction": "nt",
      "name": "Rental Officer"
    },
    {
      "databaseId": "ntca",
      "jurisdiction": "nt",
      "name": "Court of Appeal for the Northwest Territories"
    },
    {
      "databaseId": "cart-crac",
      "jurisdiction": "ca",
      "name": "Canada Agricultural Review Tribunal"
    },
    {
      "databaseId": "skcp",
      "jurisdiction": "sk",
      "name": "Saskatchewan College of Pharmacists"
    },
    {
      "databaseId": "ntipc",
      "jurisdiction": "nt",
      "name": "Northwest Territories Information and Privacy Commissioner"
    },
    {
      "databaseId": "ntsc",
      "jurisdiction": "nt",
      "name": "Supreme Court of the Northwest Territories"
    },
    {
      "databaseId": "qcopgq",
      "jurisdiction": "qc",
      "name": "Ordre professionnel des géologues du Québec"
    },
    {
      "databaseId": "abtsb",
      "jurisdiction": "ab",
      "name": "Alberta Transportation Safety Board"
    },
    {
      "databaseId": "nlpc",
      "jurisdiction": "nl",
      "name": "Provincial Court of Newfoundland and Labrador"
    },
    {
      "databaseId": "nbla",
      "jurisdiction": "nb",
      "name": "Labour Arbitration Awards"
    },
    {
      "databaseId": "qcohdq",
      "jurisdiction": "qc",
      "name": "Ordre des hygiénistes dentaires du Québec"
    },
    {
      "databaseId": "qcoaciq",
      "jurisdiction": "qc",
      "name": "Comité de discipline de l\u0027organisme d\u0027autoréglementation du courtage immobilier du Québec"
    },
    {
      "databaseId": "qcrde",
      "jurisdiction": "qc",
      "name": "Régie de l\u0027énergie"
    },
    {
      "databaseId": "skdc",
      "jurisdiction": "sk",
      "name": "Saskatchewan District Court"
    },
    {
      "databaseId": "mblb",
      "jurisdiction": "mb",
      "name": "Manitoba Labour Board"
    },
    {
      "databaseId": "qcbdrvm",
      "jurisdiction": "qc",
      "name": "Bureau de décision et de révision en valeurs mobilières"
    },
    {
      "databaseId": "qccdoii",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre des infirmières et infirmiers du Québec"
    },
    {
      "databaseId": "mbla",
      "jurisdiction": "mb",
      "name": "Labour Arbitration Awards"
    },
    {
      "databaseId": "qcoeaq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre professionnel des évaluateurs agréés du Québec"
    },
    {
      "databaseId": "cbsc-ccnr",
      "jurisdiction": "ca",
      "name": "Canadian Broadcast Standards Council"
    },
    {
      "databaseId": "skufc",
      "jurisdiction": "sk",
      "name": "Saskatchewan Unified Family Court"
    },
    {
      "databaseId": "bcca",
      "jurisdiction": "bc",
      "name": "Court of Appeal"
    },
    {
      "databaseId": "qcoeq",
      "jurisdiction": "qc",
      "name": "Comité de discipline de l\u0027Ordre des ergothérapeutes du Québec"
    },
    {
      "databaseId": "onocco",
      "jurisdiction": "on",
      "name": "Office of the Chief Coroner for Ontario"
    },
    {
      "databaseId": "nslb",
      "jurisdiction": "ns",
      "name": "Nova Scotia Labour Board"
    },
    {
      "databaseId": "nsla",
      "jurisdiction": "ns",
      "name": "Labour Arbitration Awards"
    },
    {
      "databaseId": "mbls",
      "jurisdiction": "mb",
      "name": "Manitoba Law Society Discipline Committee"
    },
    {
      "databaseId": "oncfsrb",
      "jurisdiction": "on",
      "name": "Child and Family Services Review Board"
    },
    {
      "databaseId": "skaia",
      "jurisdiction": "sk",
      "name": "Automobile Injury Appeal Commission"
    },
    {
      "databaseId": "skca",
      "jurisdiction": "sk",
      "name": "Court of Appeal for Saskatchewan"
    },
    {
      "databaseId": "qccvm",
      "jurisdiction": "qc",
      "name": "Commission des valeurs mobilières du Québec"
    },
    {
      "databaseId": "qcracj",
      "jurisdiction": "qc",
      "name": "Régie des alcools des courses et des jeux"
    },
    {
      "databaseId": "qctdp",
      "jurisdiction": "qc",
      "name": "Human Rights Tribunal"
    },
    {
      "databaseId": "sklgb",
      "jurisdiction": "sk",
      "name": "Saskatchewan Local Government Board"
    },
    {
      "databaseId": "peipc",
      "jurisdiction": "pe",
      "name": "Information and Privacy Commissioner"
    },
    {
      "databaseId": "absrb",
      "jurisdiction": "ab",
      "name": "Alberta Surface Rights Board"
    },
    {
      "databaseId": "qcooaq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre professionnel des orthophonistes et audiologistes du Québec"
    },
    {
      "databaseId": "bcsc",
      "jurisdiction": "bc",
      "name": "Supreme Court of British Columbia"
    },
    {
      "databaseId": "sksc",
      "jurisdiction": "sk",
      "name": "Supreme Court of Saskatchewan"
    },
    {
      "databaseId": "caiiroc",
      "jurisdiction": "ca",
      "name": "Investment Industry Regulatory Organization of Canada"
    },
    {
      "databaseId": "nswcat",
      "jurisdiction": "ns",
      "name": "Nova Scotia Workers\u0027 Compensation Appeals Tribunal"
    },
    {
      "databaseId": "nlsctd",
      "jurisdiction": "nl",
      "name": "Supreme Court of Newfoundland and Labrador, Trial Division"
    },
    {
      "databaseId": "qccfp",
      "jurisdiction": "qc",
      "name": "Commission de la fonction publique"
    },
    {
      "databaseId": "ntllb",
      "jurisdiction": "nt",
      "name": "Northwest Territories Liquor Licensing Board"
    },
    {
      "databaseId": "qcadmaq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre des administrateurs agréés du Québec"
    },
    {
      "databaseId": "qcopiq",
      "jurisdiction": "qc",
      "name": "Ordre professionnel des inhalothérapeutes du Québec"
    },
    {
      "databaseId": "nlipc",
      "jurisdiction": "nl",
      "name": "Information and Privacy Commissioner"
    },
    {
      "databaseId": "sksu",
      "jurisdiction": "sk",
      "name": "Saskatchewan Surrogate Court"
    },
    {
      "databaseId": "bccp",
      "jurisdiction": "bc",
      "name": "College of Pharmacists of British Columbia"
    },
    {
      "databaseId": "qccdnq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de la Chambre des notaires du Québec"
    },
    {
      "databaseId": "nbleb",
      "jurisdiction": "nb",
      "name": "New Brunswick Labour and Employment Board"
    },
    {
      "databaseId": "onhparb",
      "jurisdiction": "on",
      "name": "Health Professions Appeal and Review Board"
    },
    {
      "databaseId": "ablcb",
      "jurisdiction": "ab",
      "name": "Alberta Land Compensation Board"
    },
    {
      "databaseId": "bclrb",
      "jurisdiction": "bc",
      "name": "Labour Relations Board"
    },
    {
      "databaseId": "qcottiaq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre professionnel des traducteurs, terminologues et interprètes agréés du Québec"
    },
    {
      "databaseId": "bcsec",
      "jurisdiction": "bc",
      "name": "British Columbia Securities Commission"
    },
    {
      "databaseId": "onhsarb",
      "jurisdiction": "on",
      "name": "Health Services Appeal and Review Board"
    },
    {
      "databaseId": "mbcpsdc",
      "jurisdiction": "mb",
      "name": "College of Physicians \u0026 Surgeons of Manitoba Discipline Committee"
    },
    {
      "databaseId": "skpc",
      "jurisdiction": "sk",
      "name": "Provincial Court of Saskatchewan"
    },
    {
      "databaseId": "qccmq",
      "jurisdiction": "qc",
      "name": "Conseil de la magistrature"
    },
    {
      "databaseId": "onwsib",
      "jurisdiction": "on",
      "name": "Ontario Workplace Safety and Insurance Board"
    },
    {
      "databaseId": "nbwhscc",
      "jurisdiction": "nb",
      "name": "Workplace Health, Safety and Compensation Commission Appeal Tribunal"
    },
    {
      "databaseId": "ablerb",
      "jurisdiction": "ab",
      "name": "Alberta Law Enforcement Review Board"
    },
    {
      "databaseId": "pela",
      "jurisdiction": "pe",
      "name": "Prince Edward Island Labour Arbitration Awards"
    },
    {
      "databaseId": "qccmnq",
      "jurisdiction": "qc",
      "name": "Commission municipale du Québec"
    },
    {
      "databaseId": "camfda",
      "jurisdiction": "ca",
      "name": "Mutual Fund Dealers Association of Canada"
    },
    {
      "databaseId": "skac",
      "jurisdiction": "sk",
      "name": "Saskatchewan Assessment Commission"
    },
    {
      "databaseId": "qctacarra",
      "jurisdiction": "qc",
      "name": "Tribunal d\u0027arbitrage de la Commission administrative des régimes de retraite et d\u0027assurances"
    },
    {
      "databaseId": "bccps",
      "jurisdiction": "bc",
      "name": "College of Physicians and Surgeons of British Columbia"
    },
    {
      "databaseId": "nuca",
      "jurisdiction": "nu",
      "name": "Court of Appeal of Nunavut"
    },
    {
      "databaseId": "nbomb",
      "jurisdiction": "nb",
      "name": "New Brunswick Office of the Ombudsman"
    },
    {
      "databaseId": "abca",
      "jurisdiction": "ab",
      "name": "Court of Appeal"
    },
    {
      "databaseId": "nucj",
      "jurisdiction": "nu",
      "name": "Nunavut Court of Justice"
    },
    {
      "databaseId": "qctt",
      "jurisdiction": "qc",
      "name": "Labour Court"
    },
    {
      "databaseId": "onpeht",
      "jurisdiction": "on",
      "name": "Ontario Pay Equity Hearings Tribunal"
    },
    {
      "databaseId": "skqb",
      "jurisdiction": "sk",
      "name": "Court of Queen\u0027s Bench for Saskatchewan"
    },
    {
      "databaseId": "qctaq",
      "jurisdiction": "qc",
      "name": "Administrative Tribunal of Québec"
    },
    {
      "databaseId": "qcclp",
      "jurisdiction": "qc",
      "name": "Commission des lésions professionnelles du Québec"
    },
    {
      "databaseId": "qctp",
      "jurisdiction": "qc",
      "name": "Professions Tribunal"
    },
    {
      "databaseId": "nbpc",
      "jurisdiction": "nb",
      "name": "Provincial Court"
    },
    {
      "databaseId": "qcochq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre professionnel des chimistes du Québec"
    },
    {
      "databaseId": "qccdhj",
      "jurisdiction": "qc",
      "name": "Comité de discipline de la Chambre des Huissiers de Justice du Québec"
    },
    {
      "databaseId": "cm",
      "jurisdiction": "ca",
      "name": "Courts Martial"
    },
    {
      "databaseId": "qctaa",
      "jurisdiction": "qc",
      "name": "Arbitration - performing, recording and film artists"
    },
    {
      "databaseId": "nshrc",
      "jurisdiction": "ns",
      "name": "Nova Scotia Human Rights Commission"
    },
    {
      "databaseId": "absec",
      "jurisdiction": "ab",
      "name": "Alberta Securities Commission"
    },
    {
      "databaseId": "abesu",
      "jurisdiction": "ab",
      "name": "Alberta Employment Standards Umpire"
    },
    {
      "databaseId": "pslreb",
      "jurisdiction": "ca",
      "name": "Public Service Labour Relations and Employment Board"
    },
    {
      "databaseId": "nscps",
      "jurisdiction": "ns",
      "name": "College of Physicians and Surgeons of Nova Scotia"
    },
    {
      "databaseId": "onafraat",
      "jurisdiction": "on",
      "name": "Agriculture, Food \u0026 Rural Affairs Appeal Tribunal"
    },
    {
      "databaseId": "onipc",
      "jurisdiction": "on",
      "name": "Information and Privacy Commissioner Ontario"
    },
    {
      "databaseId": "mbsec",
      "jurisdiction": "mb",
      "name": "Manitoba Securities Commission"
    },
    {
      "databaseId": "bcipc",
      "jurisdiction": "bc",
      "name": "Information and Privacy Commissioner"
    },
    {
      "databaseId": "qcoapq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre des audioprothésistes du Québec"
    },
    {
      "databaseId": "qcotstcfq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre professionnel des travailleurs sociaux et des thérapeutes conjugaux et familiaux du Québec"
    },
    {
      "databaseId": "nbqb",
      "jurisdiction": "nb",
      "name": "Court of Queen\u0027s Bench of New Brunswick"
    },
    {
      "databaseId": "onca",
      "jurisdiction": "on",
      "name": "Court of Appeal for Ontario"
    },
    {
      "databaseId": "csc-scc-al",
      "jurisdiction": "ca",
      "name": "Supreme Court of Canada - Applications for Leave"
    },
    {
      "databaseId": "sklrb",
      "jurisdiction": "sk",
      "name": "Saskatchewan Labour Relations Board"
    },
    {
      "databaseId": "onsbt",
      "jurisdiction": "on",
      "name": "Ontario Social Benefits Tribunal"
    },
    {
      "databaseId": "qcamf",
      "jurisdiction": "qc",
      "name": "Quebec Autorité des marchés financiers"
    },
    {
      "databaseId": "nsohsap",
      "jurisdiction": "ns",
      "name": "Nova Scotia Occupational Health and Safety Appeal Panel"
    },
    {
      "databaseId": "nsfc",
      "jurisdiction": "ns",
      "name": "Nova Scotia Family Court"
    },
    {
      "databaseId": "mbhrc",
      "jurisdiction": "mb",
      "name": "Manitoba Human Rights Commission"
    },
    {
      "databaseId": "qccm",
      "jurisdiction": "qc",
      "name": "Municipal Courts"
    },
    {
      "databaseId": "cact",
      "jurisdiction": "ca",
      "name": "Competition Tribunal"
    },
    {
      "databaseId": "qcca",
      "jurisdiction": "qc",
      "name": "Court of Appeal"
    },
    {
      "databaseId": "bchrt",
      "jurisdiction": "bc",
      "name": "British Columbia Human Rights Tribunal"
    },
    {
      "databaseId": "sklss",
      "jurisdiction": "sk",
      "name": "Law Society of Saskatchewan"
    },
    {
      "databaseId": "fct",
      "jurisdiction": "ca",
      "name": "Federal Court of Canada"
    },
    {
      "databaseId": "skhrc",
      "jurisdiction": "sk",
      "name": "Saskatchewan Human Rights Commission"
    },
    {
      "databaseId": "qccja",
      "jurisdiction": "qc",
      "name": "Conseil de la justice administrative"
    },
    {
      "databaseId": "nlca",
      "jurisdiction": "nl",
      "name": "Supreme Court of Newfoundland and Labrador, Court of Appeal"
    },
    {
      "databaseId": "fca",
      "jurisdiction": "ca",
      "name": "Federal Court of Appeal"
    },
    {
      "databaseId": "ongsb",
      "jurisdiction": "on",
      "name": "Grievance Settlement Board"
    },
    {
      "databaseId": "nllrb",
      "jurisdiction": "nl",
      "name": "Newfoundland and Labrador Labour Relations Board"
    },
    {
      "databaseId": "skmt",
      "jurisdiction": "sk",
      "name": "Saskatchewan Master of Titles"
    },
    {
      "databaseId": "qccmmtq",
      "jurisdiction": "qc",
      "name": "Corporation of Master Pipe-Mechanics of Québec"
    },
    {
      "databaseId": "qcrmaaq",
      "jurisdiction": "qc",
      "name": "Régie des marchés agricoles et alimentaires du Québec"
    },
    {
      "databaseId": "qcoifq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre professionnel des ingénieurs forestiers du Québec"
    },
    {
      "databaseId": "qccpa",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre des comptables professionnels agréés du Québec"
    },
    {
      "databaseId": "onset",
      "jurisdiction": "on",
      "name": "Ontario Special Education  (English) Tribunal"
    },
    {
      "databaseId": "qcopq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre des psychologues du Québec"
    },
    {
      "databaseId": "abcpsdc",
      "jurisdiction": "ab",
      "name": "College of Physicians and Surgeons Discipline Committee"
    },
    {
      "databaseId": "nlhrc",
      "jurisdiction": "nl",
      "name": "Newfoundland and Labrador Human Rights Commission"
    },
    {
      "databaseId": "qcoagbrn",
      "jurisdiction": "qc",
      "name": "Arbitration - The Guarantee Plan For New Residential Buildings"
    },
    {
      "databaseId": "qcdag",
      "jurisdiction": "qc",
      "name": "Labour Arbitration Awards (including Conférence des arbitres)"
    },
    {
      "databaseId": "qccpq",
      "jurisdiction": "qc",
      "name": "Conseil de presse du Québec"
    },
    {
      "databaseId": "qccai",
      "jurisdiction": "qc",
      "name": "Commission d\u0027accès à l\u0027information"
    },
    {
      "databaseId": "nsprb",
      "jurisdiction": "ns",
      "name": "Nova Scotia Police Review Board"
    },
    {
      "databaseId": "skfca",
      "jurisdiction": "sk",
      "name": "Saskatchewan Board of Review under the Farmers\u0027 Creditors Arrangement Act, 1934"
    },
    {
      "databaseId": "nspr",
      "jurisdiction": "ns",
      "name": "Nova Scotia Probate Court"
    },
    {
      "databaseId": "qccq",
      "jurisdiction": "qc",
      "name": "Court of Quebec"
    },
    {
      "databaseId": "aboipc",
      "jurisdiction": "ab",
      "name": "Office of the Information and Privacy Commissioner"
    },
    {
      "databaseId": "qccs",
      "jurisdiction": "qc",
      "name": "Superior Court"
    },
    {
      "databaseId": "onla",
      "jurisdiction": "on",
      "name": "Labour Arbitration Awards"
    },
    {
      "databaseId": "cisr",
      "jurisdiction": "ca",
      "name": "Immigration and Refugee Board of Canada"
    },
    {
      "databaseId": "peihrc",
      "jurisdiction": "pe",
      "name": "Prince Edward Island Human Rights Commission"
    },
    {
      "databaseId": "nslrb",
      "jurisdiction": "ns",
      "name": "Nova Scotia Labour Relations Board"
    },
    {
      "databaseId": "qcodlq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre professionnel des denturologistes du Québec"
    },
    {
      "databaseId": "qccdcm",
      "jurisdiction": "qc",
      "name": "Collège des médecins du Québec Disciplinary Council"
    },
    {
      "databaseId": "oncps",
      "jurisdiction": "on",
      "name": "College of Physicians and Surgeons of Ontario"
    },
    {
      "databaseId": "ntla",
      "jurisdiction": "nt",
      "name": "Labour Arbitration Awards"
    },
    {
      "databaseId": "qccdomv",
      "jurisdiction": "qc",
      "name": "Comité de discipline de l\u0027Ordre des médecins vétérinaires du Québec"
    },
    {
      "databaseId": "mbca",
      "jurisdiction": "mb",
      "name": "Court of Appeal"
    },
    {
      "databaseId": "skla",
      "jurisdiction": "sk",
      "name": "Labour Arbitration Awards"
    },
    {
      "databaseId": "onsec",
      "jurisdiction": "on",
      "name": "Ontario Securities Commission"
    },
    {
      "databaseId": "qcagq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre des agronomes du Québec"
    },
    {
      "databaseId": "nbca",
      "jurisdiction": "nb",
      "name": "Court of Appeal of New Brunswick"
    },
    {
      "databaseId": "qccdbq",
      "jurisdiction": "qc",
      "name": "Barreau du Québec  Disciplinary Council"
    },
    {
      "databaseId": "nssec",
      "jurisdiction": "ns",
      "name": "Nova Scotia Securities Commission"
    },
    {
      "databaseId": "nssm",
      "jurisdiction": "ns",
      "name": "Small Claims Court"
    },
    {
      "databaseId": "nuhrt",
      "jurisdiction": "nu",
      "name": "Nunavut Human Rights Tribunal"
    },
    {
      "databaseId": "nssf",
      "jurisdiction": "ns",
      "name": "Supreme Court of Nova Scotia (Family Division)"
    },
    {
      "databaseId": "nssc",
      "jurisdiction": "ns",
      "name": "Supreme Court of Nova Scotia"
    },
    {
      "databaseId": "abpc",
      "jurisdiction": "ab",
      "name": "Provincial Court"
    },
    {
      "databaseId": "ntls",
      "jurisdiction": "nt",
      "name": "Law Society of the Northwest Territories"
    },
    {
      "databaseId": "nsca",
      "jurisdiction": "ns",
      "name": "Nova Scotia Court of Appeal"
    },
    {
      "databaseId": "onsc",
      "jurisdiction": "on",
      "name": "Superior Court of Justice"
    },
    {
      "databaseId": "onwsiat",
      "jurisdiction": "on",
      "name": "Ontario Workplace Safety and Insurance Appeals Tribunal"
    },
    {
      "databaseId": "onagc",
      "jurisdiction": "on",
      "name": "Alcohol and Gaming Commission of Ontario"
    },
    {
      "databaseId": "cci-tcc",
      "jurisdiction": "ca",
      "name": "Tax Court of Canada"
    },
    {
      "databaseId": "onlst",
      "jurisdiction": "on",
      "name": "Law Society Tribunal"
    },
    {
      "databaseId": "onrc",
      "jurisdiction": "on",
      "name": "Ontario Racing Commission"
    },
    {
      "databaseId": "qcoppq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre professionnel de la physiothérapie du Québec"
    },
    {
      "databaseId": "qcooq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre des optométristes du Québec"
    },
    {
      "databaseId": "oncno",
      "jurisdiction": "on",
      "name": "College of Nurses of Ontario Discipline Committee"
    },
    {
      "databaseId": "nsbs",
      "jurisdiction": "ns",
      "name": "Nova Scotia Barristers\u0027 Society Hearing Panel"
    },
    {
      "databaseId": "qccsst",
      "jurisdiction": "qc",
      "name": "Commission de la santé et de la sécurité du travail"
    },
    {
      "databaseId": "ykyc",
      "jurisdiction": "yk",
      "name": "Territorial Court of Yukon (Youth Court)"
    },
    {
      "databaseId": "onltb",
      "jurisdiction": "on",
      "name": "Landlord and Tenant Board"
    },
    {
      "databaseId": "bcla",
      "jurisdiction": "bc",
      "name": "Labour Arbitration Awards"
    },
    {
      "databaseId": "qccdopq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre des pharmaciens du Québec"
    },
    {
      "databaseId": "cirb",
      "jurisdiction": "ca",
      "name": "Canada Industrial Relations Board"
    },
    {
      "databaseId": "abqb",
      "jurisdiction": "ab",
      "name": "Court of Queen\u0027s Bench"
    },
    {
      "databaseId": "qcotpq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre des technologues professionnels du Québec"
    },
    {
      "databaseId": "qcotimro",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre des technologues en imagerie médicale et en radio-oncologie du Québec"
    },
    {
      "databaseId": "abhrc",
      "jurisdiction": "ab",
      "name": "Alberta Human Rights Commission"
    },
    {
      "databaseId": "nslst",
      "jurisdiction": "ns",
      "name": "Nova Scotia Labour Standards Tribunal"
    },
    {
      "databaseId": "qccdosfq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre des sages femmes du Québec"
    },
    {
      "databaseId": "onscsm",
      "jurisdiction": "on",
      "name": "Small Claims Court"
    },
    {
      "databaseId": "pssrb",
      "jurisdiction": "ca",
      "name": "Public Service Labour Relations Board"
    },
    {
      "databaseId": "onscdc",
      "jurisdiction": "on",
      "name": "Divisional Court"
    },
    {
      "databaseId": "pescad",
      "jurisdiction": "pe",
      "name": "Prince Edward Island Court of Appeal"
    },
    {
      "databaseId": "vrab",
      "jurisdiction": "ca",
      "name": "Veterans Review and Appeal Board of Canada"
    },
    {
      "databaseId": "skatmpa",
      "jurisdiction": "sk",
      "name": "Appeal Tribunal under the Medical Profession Act"
    },
    {
      "databaseId": "pcc-cvpc",
      "jurisdiction": "ca",
      "name": "Privacy Commissioner of Canada"
    },
    {
      "databaseId": "qcodq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre des dentistes du Québec"
    },
    {
      "databaseId": "qcces",
      "jurisdiction": "qc",
      "name": "Commission de l\u0027équité salariale"
    },
    {
      "databaseId": "ykca",
      "jurisdiction": "yk",
      "name": "Court of Appeal"
    },
    {
      "databaseId": "onfst",
      "jurisdiction": "on",
      "name": "Financial Services Tribunal"
    },
    {
      "databaseId": "onhrt",
      "jurisdiction": "on",
      "name": "Human Rights Tribunal of Ontario"
    },
    {
      "databaseId": "qccraaap",
      "jurisdiction": "qc",
      "name": "Commission de reconnaissance des associations d\u0027artistes et des associations de producteurs"
    },
    {
      "databaseId": "qcoagq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre des arpenteurs-géomètres du Québec"
    },
    {
      "databaseId": "qcotmq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre professionnel des technologistes médicaux du Québec"
    },
    {
      "databaseId": "sksec",
      "jurisdiction": "sk",
      "name": "Saskatchewan Financial Services Commission"
    },
    {
      "databaseId": "nbsec",
      "jurisdiction": "nb",
      "name": "Financial and Consumer Services Tribunal"
    },
    {
      "databaseId": "qcouq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre des urbanistes du Québec"
    },
    {
      "databaseId": "onccb",
      "jurisdiction": "on",
      "name": "Consent and Capacity Board"
    },
    {
      "databaseId": "oncpdc",
      "jurisdiction": "on",
      "name": "Ontario College of Pharmacists Discipline Committee"
    },
    {
      "databaseId": "skipc",
      "jurisdiction": "sk",
      "name": "Information and Privacy Commissioner"
    },
    {
      "databaseId": "ablrb",
      "jurisdiction": "ab",
      "name": "Alberta Labour Relations Board"
    },
    {
      "databaseId": "nlls",
      "jurisdiction": "nl",
      "name": "Law Society of Newfoundland and Labrador"
    },
    {
      "databaseId": "qccdp",
      "jurisdiction": "qc",
      "name": "Comité de déontologie policière"
    },
    {
      "databaseId": "qcocq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre des chiropraticiens du Québec"
    },
    {
      "databaseId": "nsuarb",
      "jurisdiction": "ns",
      "name": "Nova Scotia Utility and Review Board"
    },
    {
      "databaseId": "onlrb",
      "jurisdiction": "on",
      "name": "Ontario Labour Relations Board"
    },
    {
      "databaseId": "tmob",
      "jurisdiction": "ca",
      "name": "Trade-marks Opposition Board"
    },
    {
      "databaseId": "qcopodq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre des podiatres du Québec"
    },
    {
      "databaseId": "nlla",
      "jurisdiction": "nl",
      "name": "Labour Arbitration Awards"
    },
    {
      "databaseId": "capsdpt",
      "jurisdiction": "ca",
      "name": "Public Servants Disclosure Protection Tribunal"
    },
    {
      "databaseId": "qccdcsf",
      "jurisdiction": "qc",
      "name": "Comité de discipline de la Chambre de la sécurité financière"
    },
    {
      "databaseId": "nusec",
      "jurisdiction": "nu",
      "name": "Nunavut Registrar of Securities"
    },
    {
      "databaseId": "mbqb",
      "jurisdiction": "mb",
      "name": "Court of Queen\u0027s Bench of Manitoba"
    },
    {
      "databaseId": "oncrb",
      "jurisdiction": "on",
      "name": "Ontario Custody Review Board"
    },
    {
      "databaseId": "qccrt",
      "jurisdiction": "qc",
      "name": "Commission des relations du travail"
    },
    {
      "databaseId": "chrt",
      "jurisdiction": "ca",
      "name": "Canadian Human Rights Tribunal"
    },
    {
      "databaseId": "qccdoiia",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre des infirmières et infirmiers auxiliaires du Québec"
    },
    {
      "databaseId": "abplab",
      "jurisdiction": "ab",
      "name": "Alberta Public Lands Appeal Board"
    },
    {
      "databaseId": "qccsj",
      "jurisdiction": "qc",
      "name": "Commission des services juridiques"
    },
    {
      "databaseId": "nuyc",
      "jurisdiction": "nu",
      "name": "Youth Justice Court of Nunavut"
    },
    {
      "databaseId": "yktc",
      "jurisdiction": "yk",
      "name": "Territorial Court of Yukon"
    },
    {
      "databaseId": "abls",
      "jurisdiction": "ab",
      "name": "Law Society of Alberta - Hearing Committee"
    },
    {
      "databaseId": "qccdrhri",
      "jurisdiction": "qc",
      "name": "Conseil de discipline de l\u0027Ordre professionnel des conseillers en ressources humaines et en relations industrielles agrées du Québec"
    },
    {
      "databaseId": "nbapab",
      "jurisdiction": "nb",
      "name": "New Brunswick Assessment and Planning Appeal Board"
    },
    {
      "databaseId": "qccse",
      "jurisdiction": "qc",
      "name": "Conseil des services essentiels"
    },
    {
      "databaseId": "onlat",
      "jurisdiction": "on",
      "name": "Ontario Licence Appeal Tribunal"
    },
    {
      "databaseId": "qcrbq",
      "jurisdiction": "qc",
      "name": "Régie du Bâtiment - licences d\u0027entrepreneur de construction"
    },
    {
      "databaseId": "psst",
      "jurisdiction": "ca",
      "name": "Public Service Staffing Tribunal"
    },
    {
      "databaseId": "qccdppq",
      "jurisdiction": "qc",
      "name": "Conseil de discipline des psychoéducateurs et psychoéducatrices du Québec"
    },
    {
      "databaseId": "nthrap",
      "jurisdiction": "nt",
      "name": "Human Rights Adjudication Panel"
    },
    {
      "databaseId": "nspc",
      "jurisdiction": "ns",
      "name": "Provincial Court of Nova Scotia"
    },
    {
      "databaseId": "ntlsb",
      "jurisdiction": "nt",
      "name": "Employment Standards Appeals Office"
    },
    {
      "databaseId": "absdab",
      "jurisdiction": "ab",
      "name": "Calgary Subdivision \u0026 Development Appeal Board"
    },
    {
      "databaseId": "qcoaq",
      "jurisdiction": "qc",
      "name": "Comité de discipline de l\u0027Ordre des acupuncteurs du Québec"
    },
    {
      "databaseId": "casa-cala",
      "jurisdiction": "ca",
      "name": "Labour Arbitration Awards"
    },
    {
      "databaseId": "bcpc",
      "jurisdiction": "bc",
      "name": "Provincial Court of British Columbia"
    },
    {
      "databaseId": "lsbc",
      "jurisdiction": "bc",
      "name": "Law Society of British Columbia"
    },
    {
      "databaseId": "nsfoipop",
      "jurisdiction": "ns",
      "name": "Nova Scotia Freedom of Information and Protection of Privacy Review Officer"
    },
    {
      "databaseId": "mbpc",
      "jurisdiction": "mb",
      "name": "Provincial Court of Manitoba"
    }
  ]
}`
	var v struct {
		DBs []canlii.CaseDatabase `json:"caseDatabases"`
	}
	err := json.NewDecoder(strings.NewReader(resp)).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	wantDB := v.DBs
	api := "derp"
	path := "/v1/caseBrowse/en/"

	client, close := testClient(t, path, resp, api)
	defer close()

	gotDBs, _, err := client.CaseBrowse.ListDatabases()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantDB, gotDBs) {
		t.Logf("want=%+v", wantDB)
		t.Logf(" got=%+v", gotDBs)
		t.Error("mismatch yo")
	}
}

func TestCaseBrowseCaseList(t *testing.T) {
	resp := `{
  "cases": [
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "2015scc7"
      },
      "title": "Canada (Attorney General) v. Federation of Law Societies of Canada",
      "citation": "2015 SCC 7 (CanLII)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "2015scc6"
      },
      "title": "R. v. Goleski",
      "citation": "2015 SCC 6 (CanLII)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "2015scc5"
      },
      "title": "Carter v. Canada (Attorney General)",
      "citation": "2015 SCC 5 (CanLII)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "2015scc4"
      },
      "title": "Saskatchewan Federation of Labour v. Saskatchewan",
      "citation": "2015 SCC 4 (CanLII)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "2015scc3"
      },
      "title": "Tervita Corp. v. Canada (Commissioner of Competition)",
      "citation": "2015 SCC 3 (CanLII)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "2015scc2"
      },
      "title": "Meredith v. Canada (Attorney General)",
      "citation": "2015 SCC 2 (CanLII)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "2015scc1"
      },
      "title": "Mounted Police Association of Ontario v. Canada (Attorney General)",
      "citation": "2015 SCC 1 (CanLII)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "2014scc76"
      },
      "title": "R. v. MacLeod",
      "citation": "2014 SCC 76 (CanLII)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "2014scc77"
      },
      "title": "R. v. Fearon",
      "citation": "2014 SCC 77 (CanLII)"
    },
    {
      "databaseId": "csc-scc",
      "caseId": {
        "en": "2014scc75"
      },
      "title": "R. v. Wilcox",
      "citation": "2014 SCC 75 (CanLII)"
    }
  ]
}`
	var v struct {
		Cases []canlii.Case `json:"cases"`
	}
	err := json.NewDecoder(strings.NewReader(resp)).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	wantCases := v.Cases
	api := "derp"
	path := "/v1/caseBrowse/en/csc-scc"

	client, close := testClient(t, path, resp, api)
	defer close()

	gotCases, _, err := client.CaseBrowse.ListCases("csc-scc", nil)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantCases, gotCases) {
		t.Logf("want=%+v", wantCases)
		t.Logf(" got=%+v", gotCases)
		t.Error("mismatch yo")
	}
}

func TestCaseBrowseCaseMetadata(t *testing.T) {
	resp := `{
  "databaseId": "abwcac",
  "caseId": "2013canlii10946",
  "url": "http://canlii.ca/t/fwfv0",
  "title": "Decision No: 2013-0180",
  "citation": "2013 CanLII 10946 (AB WCAC)",
  "language": "en",
  "docketNumber": "2013-0180; AC0068-12-81",
  "decisionDate": "2013-03-07"
}`
	var v struct {
		Cases []canlii.CaseMetadata `json:"cases"`
	}
	err := json.NewDecoder(strings.NewReader(resp)).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	wantCases := v.Cases
	api := "derp"
	path := "/v1/caseBrowse/en/abwcac/2013canlii10946"

	client, close := testClient(t, path, resp, api)
	defer close()

	gotCases, _, err := client.CaseBrowse.CaseMetadata("abwcac", "2013canlii10946")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantCases, gotCases) {
		t.Logf("want=%+v", wantCases)
		t.Logf(" got=%+v", gotCases)
		t.Error("mismatch yo")
	}
}
