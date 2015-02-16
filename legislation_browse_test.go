package canlii_test

import (
	"encoding/json"
	"github.com/aybabtme/canlii"
	"reflect"
	"strings"
	"testing"
)

func TestLegislationBrowseDatabaseList(t *testing.T) {
	resp := `{
  "legislationDatabases": [
    {
      "databaseId": "car",
      "type": "REGULATION",
      "jurisdiction": "ca",
      "name": "Consolidated Regulations of Canada"
    },
    {
      "databaseId": "cas",
      "type": "STATUTE",
      "jurisdiction": "ca",
      "name": "Consolidated Statutes of Canada"
    },
    {
      "databaseId": "abr",
      "type": "REGULATION",
      "jurisdiction": "ab",
      "name": "Consolidated Regulations of Alberta"
    },
    {
      "databaseId": "abs",
      "type": "STATUTE",
      "jurisdiction": "ab",
      "name": "Consolidated Statutes of Alberta"
    },
    {
      "databaseId": "onr",
      "type": "REGULATION",
      "jurisdiction": "on",
      "name": "Consolidated Regulations of Ontario"
    },
    {
      "databaseId": "ons",
      "type": "STATUTE",
      "jurisdiction": "on",
      "name": "Consolidated Statutes of Ontario"
    },
    {
      "databaseId": "nts",
      "type": "STATUTE",
      "jurisdiction": "nt",
      "name": "Consolidated Statutes of the Northwest Territories"
    },
    {
      "databaseId": "ntr",
      "type": "REGULATION",
      "jurisdiction": "nt",
      "name": "Regulations of the Northwest Territories"
    },
    {
      "databaseId": "qcs",
      "type": "STATUTE",
      "jurisdiction": "qc",
      "name": "Consolidated Statutes of Quebec"
    },
    {
      "databaseId": "qcr",
      "type": "REGULATION",
      "jurisdiction": "qc",
      "name": "Regulations of Quebec"
    },
    {
      "databaseId": "nlr",
      "type": "REGULATION",
      "jurisdiction": "nl",
      "name": "Regulations of Newfoundland and Labrador"
    },
    {
      "databaseId": "nls",
      "type": "STATUTE",
      "jurisdiction": "nl",
      "name": "Consolidated Statutes of Newfoundland and Labrador"
    },
    {
      "databaseId": "yks",
      "type": "STATUTE",
      "jurisdiction": "yk",
      "name": "Statutes of Yukon"
    },
    {
      "databaseId": "ykr",
      "type": "REGULATION",
      "jurisdiction": "yk",
      "name": "Regulations of Yukon"
    },
    {
      "databaseId": "nbr",
      "type": "REGULATION",
      "jurisdiction": "nb",
      "name": "Consolidated Regulations of New Brunswick"
    },
    {
      "databaseId": "nbs",
      "type": "STATUTE",
      "jurisdiction": "nb",
      "name": "Consolidated Statutes of New Brunswick"
    },
    {
      "databaseId": "sks",
      "type": "STATUTE",
      "jurisdiction": "sk",
      "name": "Consolidated Statutes of Saskatchewan"
    },
    {
      "databaseId": "skr",
      "type": "REGULATION",
      "jurisdiction": "sk",
      "name": "Consolidated Regulations of Saskatchewan"
    },
    {
      "databaseId": "nsr",
      "type": "REGULATION",
      "jurisdiction": "ns",
      "name": "Regulations of Nova Scotia"
    },
    {
      "databaseId": "per",
      "type": "REGULATION",
      "jurisdiction": "pe",
      "name": "Regulations of Prince Edward Island"
    },
    {
      "databaseId": "nss",
      "type": "STATUTE",
      "jurisdiction": "ns",
      "name": "Consolidated Statutes of Nova Scotia"
    },
    {
      "databaseId": "pes",
      "type": "STATUTE",
      "jurisdiction": "pe",
      "name": "Consolidated Statutes of Prince Edward Island"
    },
    {
      "databaseId": "nus",
      "type": "STATUTE",
      "jurisdiction": "nu",
      "name": "Consolidated Statutes of Nunavut"
    },
    {
      "databaseId": "mbs",
      "type": "STATUTE",
      "jurisdiction": "mb",
      "name": "Consolidated Statutes of Manitoba"
    },
    {
      "databaseId": "bcs",
      "type": "STATUTE",
      "jurisdiction": "bc",
      "name": "Consolidated Statutes of British Columbia"
    },
    {
      "databaseId": "bcr",
      "type": "REGULATION",
      "jurisdiction": "bc",
      "name": "Consolidated Regulations of British Columbia"
    },
    {
      "databaseId": "mbr",
      "type": "REGULATION",
      "jurisdiction": "mb",
      "name": "Consolidated Regulations of Manitoba"
    },
    {
      "databaseId": "nur",
      "type": "REGULATION",
      "jurisdiction": "nu",
      "name": "Regulations of Nunavut"
    }
  ]
}`
	var v struct {
		DBs []canlii.LegislationBrowseDatabase `json:"legislationDatabases"`
	}
	err := json.NewDecoder(strings.NewReader(resp)).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	wantDB := v.DBs
	api := "derp"
	path := "/v1/legislationBrowse/en/"

	client, close := testClient(t, path, resp, api)
	defer close()

	gotDBs, _, err := client.LegislationBrowse.ListDatabases()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantDB, gotDBs) {
		t.Logf("want=%+v", wantDB)
		t.Logf(" got=%+v", gotDBs)
		t.Error("mismatch yo")
	}
}

func TestLegislationBrowseCaseList(t *testing.T) {
	resp := `{
  "legislations": [
    {
      "databaseId": "car",
      "legislationId": "dors-2014-270",
      "title": "2015 Pan American and Parapan American Games Remission Order",
      "citation": "DORS/2014-270",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-946",
      "title": "Abatement of Duties Payable Regulations",
      "citation": "SOR/86-946",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-253",
      "title": "Abbotsford Airport Zoning Regulations",
      "citation": "SOR/83-253",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-332",
      "title": "Aboriginal Communal Fishing Licences Regulations",
      "citation": "SOR/93-332",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-205",
      "title": "Aboriginal Peoples of Canada Adaptations Regulations (Firearms)",
      "citation": "SOR/98-205",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-184",
      "title": "Access to Basic Banking Services Regulations",
      "citation": "SOR/2003-184",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-24",
      "title": "Access to Funds Regulations",
      "citation": "SOR/2012-24",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-207",
      "title": "Access to Information Act Extension Order, No. 1",
      "citation": "SOR/89-207",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-113",
      "title": "Access to Information Act Heads of Government Institutions Designation Order",
      "citation": "SI/83-113",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-360",
      "title": "Access to Information Act Officers or Employees of the Atomic Energy Control Board Designation Order",
      "citation": "SOR/86-360",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-507",
      "title": "Access to Information Regulations",
      "citation": "SOR/83-507",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-438",
      "title": "Accountable Advances Regulations",
      "citation": "SOR/86-438",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-669",
      "title": "Accountable Education and Travel Advance Regulations (Dependants of Members of the Canadian Forces)",
      "citation": "CRC, c 669",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-670",
      "title": "Accountable Travel and Moving Advance Regulations (Canadian Forces)",
      "citation": "CRC, c 670",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-671",
      "title": "Accountable Travel and Moving Advance Regulations (Dependants of Members of the Canadian Forces)",
      "citation": "CRC, c 671",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1062",
      "title": "Accounting for Imported Goods and Payment of Duties Regulations",
      "citation": "SOR/86-1062",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-164",
      "title": "Accredited ICAO Missions Remission Order (Part IX of the Excise Tax Act)",
      "citation": "SI/2003-164",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-115",
      "title": "Order Acknowledging Receipt of the Assessment Done Pursuant to Subsection 23(1) of the Act",
      "citation": "SI/2005-115",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-11",
      "title": "Order Acknowledging Receipt of the Assessment Done Pursuant to Subsection 23(1) of the Act",
      "citation": "SI/2011-11",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-71",
      "title": "Order Acknowledging Receipt of the Assessments Done Pursuant to Subsection 23(1) of the Act",
      "citation": "SI/2005-71",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2007-57",
      "title": "Order Acknowledging Receipt of the Assessments Done Pursuant to Subsection 23(1) of the Act",
      "citation": "SI/2007-57",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-138",
      "title": "Order Acknowledging Receipt of the Assessments Done Pursuant to Subsection 23(1) of the Act",
      "citation": "SI/2004-138",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2008-70",
      "title": "Order Acknowledging Receipt of the Assessments Done Pursuant to Subsection 23(1) of the Act",
      "citation": "SI/2008-70",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2010-42",
      "title": "Order Acknowledging Receipt of the Assessments Done Pursuant to Subsection 23(1) of the Act",
      "citation": "SI/2010-42",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2010-76",
      "title": "Order Acknowledging Receipt of the Assessments Done Pursuant to Subsection 23(1) of the Act",
      "citation": "SI/2010-76",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-48",
      "title": "Order Acknowledging Receipt of the Assessments Done Pursuant to Subsection 23(1) of the Species at Risk Act",
      "citation": "SI/2004-48",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-289",
      "title": "ACOA Loan Insurance Regulations",
      "citation": "SOR/90-289",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-398",
      "title": "160900 Canada Inc. Acquisition Exemption Order",
      "citation": "SOR/88-398",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-77",
      "title": "Acquisition of Permanent Resident Status Fee Remission Order",
      "citation": "SI/2006-77",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-25",
      "title": "Regulation Adapting the Canada Elections Act for the Purposes of a Referendum",
      "citation": "SOR/2002-25",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-188",
      "title": "Proclamation Designating July 28 of Every Year as “A Day of Commemoration of the Great Upheaval”, Commencing on July 28, 2005",
      "citation": "SI/2003-188",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1235",
      "title": "Additional Legislative Powers Designation Order",
      "citation": "CRC, c 1235",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-47",
      "title": "Adjudication Division Rules",
      "citation": "SOR/93-47",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-317",
      "title": "Adjustment Assistance Benefit Regulations (Footwear and Tanning Workers)",
      "citation": "CRC, c 317",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-316",
      "title": "Adjustment Assistance Regulations (Textile and Clothing Workers)",
      "citation": "CRC, c 316",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-165",
      "title": "Administration of Labour Market Development Services Divestiture Regulations",
      "citation": "SOR/97-165",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-391",
      "title": "Administrative and Technical Staff of the Embassy of the United States and Families Duty and Tax Relief Privileges Order",
      "citation": "SOR/93-391",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-384",
      "title": "Administrative and Technical Staff of the Embassy of the United States and Families Privileges and Immunities Order",
      "citation": "SOR/93-384",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-101",
      "title": "Administrative Monetary Penalties (Consumer Products) Regulations",
      "citation": "SOR/2013-101",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-267",
      "title": "Administrative Monetary Penalties (OSFI) Regulations",
      "citation": "SOR/2005-267",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-97",
      "title": "Administrative Monetary Penalties Regulations",
      "citation": "SOR/2008-97",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-139",
      "title": "Administrative Monetary Penalties Regulations (Canadian Nuclear Safety Commission)",
      "citation": "SOR/2013-139",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-149",
      "title": "Administrative Monetary Penalties Regulations (International Bridges and Tunnels)",
      "citation": "SOR/2012-149",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-138",
      "title": "Administrative Monetary Penalties Regulations (National Energy Board)",
      "citation": "SOR/2013-138",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-234",
      "title": "Advance Income Tax Ruling Fees Order",
      "citation": "SOR/90-234",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-419",
      "title": "Advance Payments for Crops Collection Regulations",
      "citation": "SOR/86-419",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-672",
      "title": "Advance Payments for Crops Guarantee Assignments Regulations",
      "citation": "CRC, c 672",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-446",
      "title": "Advance Payments for Crops Regulations",
      "citation": "CRC, c 446",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-739",
      "title": "Advertising Material Remission Order",
      "citation": "CRC, c 739",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-447",
      "title": "Regulations Defining Advertising Revenues",
      "citation": "SOR/98-447",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-82-38",
      "title": "Designating the Advisory Council on the Status of Women as a Portion of the Public Service and Authorizing the President to Exercise and Perform all the Powers and Functions of the Treasury Board in Relation to Personnel Management of the Council",
      "citation": "SI/82-38",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-47",
      "title": "AECB Cost Recovery Fees Remission Order, 1997",
      "citation": "SI/99-47",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-403",
      "title": "AECL Tandem Accelerator Superconducting Cyclotron Complex Remission Order",
      "citation": "SOR/91-403",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-325",
      "title": "Affiliated Persons (Banks) Regulations",
      "citation": "SOR/92-325",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-326",
      "title": "Affiliated Persons (Insurance Companies) Regulations",
      "citation": "SOR/92-326",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-327",
      "title": "Affiliated Persons (Trust and Loan Companies) Regulations",
      "citation": "SOR/92-327",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-360",
      "title": "African Development Bank Privileges and Immunities Order",
      "citation": "SOR/84-360",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1304",
      "title": "African Development Fund Privileges and Immunities Order",
      "citation": "CRC, c 1304",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-61-378",
      "title": "Agassiz Correctional Camp, Petawawa Correctional Camp and Gatineau Correctional Camp Lands Proclaimed Penitentiaries",
      "citation": "SOR/61-378",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-78-165",
      "title": "Age Guidelines",
      "citation": "SI/78-165",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-648",
      "title": "Agencies\u0027 Orders and Regulations Approval Order",
      "citation": "CRC, c 648",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-944",
      "title": "Agents\u0027 Accounting for Imported Goods and Payment of Duties Regulations",
      "citation": "SOR/86-944",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-8",
      "title": "Order Binding Certain Agents of Her Majesty for the Purposes of Part 1 of the Personal Information Protection and Electronic Documents Act",
      "citation": "SOR/2001-8",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-363",
      "title": "Aggregate Financial Exposure (Banks) Regulations",
      "citation": "SOR/2001-363",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-364",
      "title": "Aggregate Financial Exposure (Insurance Companies) Regulations",
      "citation": "SOR/2001-364",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-365",
      "title": "Aggregate Financial Exposure (Trust and Loan Companies) Regulations",
      "citation": "SOR/2001-365",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-55",
      "title": "Proclamation Giving Notice that the Protocol Amending the Agreement between Canada and Barbados for the Avoidance of Double Taxation and the Prevention of Fiscal Evasion with respect to Taxes on Income and on Capital, done at Bridgetown on January 22, 1980 Came into Force on December 17, 2013",
      "citation": "SI/2014-55",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-81-28",
      "title": "Proclamation Declaring the Agreement Between Canada and France on Social Security in Force Effective January 13, 1981",
      "citation": "SI/81-28",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-79-7",
      "title": "Proclamation Giving Notice That an Agreement Expressed in Letters Between the High Commissioner of the United Kingdom and the Minister of National Health and Welfare Dated November 10, 1977 is in Force Effective October 13, 1978",
      "citation": "SI/79-7",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-79-40",
      "title": "Proclamation DeclaringThat the Agreement of Social Security Between Canada and Italy is in Force Effective December 20, 1978",
      "citation": "SI/79-40",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-5",
      "title": "Proclamation Declaring the Agreement on Social Security Between Canada and Antigua and Barbuda in Force January 1, 1994",
      "citation": "SI/94-5",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-82-73",
      "title": "Proclamation Declaring the Agreement on Social Security Between Canada and Hellenic Republic in Force December 8, 1981",
      "citation": "SI/82-73",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-53",
      "title": "Proclamation Declaring the Agreement on Social Security Between Canada and Ireland in Force January 1, 1992",
      "citation": "SI/92-53",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-166",
      "title": "Proclamation Declaring the Agreement on Social Security Between Canada and Jamaica in force June 3, 1983",
      "citation": "SI/83-166",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-73",
      "title": "Proclamation Declaring the Agreement on Social Security Between Canada and Luxembourg in Force April 1, 1990",
      "citation": "SI/90-73",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-48",
      "title": "Proclamation Declaring the Agreement on Social Security Between Canada and New Zealand in Force May 1, 1997",
      "citation": "SI/97-48",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-81-125",
      "title": "Proclamation Declaring the Agreement on Social Security Between Canada and Portugal in Force March 12, 1981",
      "citation": "SI/81-125",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-49",
      "title": "Proclamation Amending the Agreement on Social Security Between Canada and Spain in Force May 1, 1997",
      "citation": "SI/97-49",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-147",
      "title": "Proclamation Declaring the Agreement on Social Security Between Canada and the Commonwealth of Dominica in Force January 1, 1989",
      "citation": "SI/89-147",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-127",
      "title": "Proclamation Declaring the Agreement on Social Security Between Canada and the Federal Republic of Germany in Force April 1, 1988",
      "citation": "SI/88-127",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-6",
      "title": "Proclamation Declaring the Agreement on Social Security Between Canada and the Federation of St. Kitts and Nevis in Force January 1, 1994",
      "citation": "SI/94-6",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-134",
      "title": "Proclamation Amending the Agreement on Social Security Between Canada and the Hellenic Republic in Force December 1, 1997",
      "citation": "SI/97-134",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-129",
      "title": "Proclamation Giving Notice that the Agreement on Social Security between Canada and the Kingdom of Norway Comes into Force on January 1, 2014",
      "citation": "SI/2013-129",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-115",
      "title": "Proclamation Declaring the Agreement on Social Security Between Canada and the Netherlands in Force October 1, 1990",
      "citation": "SI/91-115",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-15",
      "title": "Proclamation Giving Notice that the Agreement on Social Security between Canada and the Republic of Bulgaria and the Administrative Agreement between the Government of Canada and the Government of the Republic of Bulgaria for the Implementation of the Agreement on Social Security between Canada and the Republic of Bulgaria are in force as of March 1, 2014",
      "citation": "SI/2014-15",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-124",
      "title": "Proclamation Declaring the Agreement on Social Security between Canada and the Republic of Cyprus in Force May 1, 1991",
      "citation": "SI/91-124",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-120",
      "title": "Proclamation Amending the Agreement on Social Security Between Canada and the Republic of Finland in Force January 1, 1997",
      "citation": "SI/96-120",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2009-87",
      "title": "Proclamation Giving Notice That the Agreement on Social Security Between Canada and the Republic of Poland",
      "citation": "SI/2009-87",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-89",
      "title": "Proclamation Giving Notice that the Agreement on Social Security between Canada and the Republic of Serbia Comes into Force on December 1, 2014",
      "citation": "SI/2014-89",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-32",
      "title": "Proclamation Declaring the Agreement on Social Security Between Canada and the Republic of the Philippines in Force March 1, 1997",
      "citation": "SI/97-32",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-112",
      "title": "Proclamation Declaring the Agreement on Social Security Between Canada and the Swiss Confederation in Force October 1, 1995",
      "citation": "SI/95-112",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-32",
      "title": "Proclamation Declaring the Agreement on Social Security Between Canada and the United Mexican States in Force May 1, 1996",
      "citation": "SI/96-32",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-82-105",
      "title": "Proclamation Declaring the Agreement on Social Security Between Canada and the United States of America in Force February 9, 1982",
      "citation": "SI/82-105",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-4",
      "title": "Proclamation Declaring the Agreement on Social Security Between Jersey, Guernsey and Canada in Force January 1, 1994",
      "citation": "SI/94-4",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-295",
      "title": "Agricultural Marketing Programs Regulations",
      "citation": "SOR/99-295",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-201",
      "title": "Agricultural Product Priority Claim (Banks) Regulations",
      "citation": "SOR/2007-201",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-187",
      "title": "Agriculture and Agri-Food Administrative Monetary Penalties Regulations",
      "citation": "SOR/2000-187",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-132",
      "title": "Agriculture and Agri-Food Administrative Monetary Penalties Regulations Respecting the Pest Control Products Act and Regulations",
      "citation": "SOR/2001-132",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-39",
      "title": "Agriculture and Fishing Property (GST/HST) Regulations",
      "citation": "SOR/91-39",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1403",
      "title": "Aids to Navigation Protection Regulations",
      "citation": "CRC, c 1403",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-235",
      "title": "Airborne Remote Sensing Services Fees Order, 1989",
      "citation": "SOR/89-235",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-361",
      "title": "Air Canada and C.N.R. Debt Rearrangement Order",
      "citation": "SOR/78-361",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-211",
      "title": "Air Canada Pension Plan Funding Regulations, 2009",
      "citation": "SOR/2009-211",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-244",
      "title": "Air Canada Pension Plan Funding Regulations, 2014",
      "citation": "SOR/2013-244",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-174",
      "title": "Air Canada Pension Plan Solvency Deficiency Funding Regulations",
      "citation": "SOR/2004-174",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-504",
      "title": "Order Authorizing Air Canada to Apply for Articles of Amendment to Amend its Articles of Incorporation",
      "citation": "SOR/88-504",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-79-2",
      "title": "Aircraft (International Service) Remission Order",
      "citation": "SI/79-2",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-109",
      "title": "Aircraft Objects Regulations",
      "citation": "SOR/2008-109",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-4",
      "title": "Air Cushion Vehicle Regulations",
      "citation": "CRC, c 4",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-178",
      "title": "Airline Guide Catalogue Remission Order",
      "citation": "SI/83-178",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1404",
      "title": "Air Pollution Regulations",
      "citation": "CRC, c 1404",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1563",
      "title": "Airport Personal Property Disposal Regulations",
      "citation": "CRC, c 1563",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-886",
      "title": "Airport Traffic Regulations",
      "citation": "CRC, c 886",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-518",
      "title": "Airport Transfer Regulations",
      "citation": "SOR/96-518",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-543",
      "title": "Airport Vehicle Parking Charges Regulations",
      "citation": "SOR/87-543",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-414",
      "title": "Air Services Charges Regulations",
      "citation": "SOR/85-414",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-58",
      "title": "Air Transportation Regulations",
      "citation": "SOR/88-58",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-206",
      "title": "Air Transportation Tax Order, 1995",
      "citation": "SOR/95-206",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-294",
      "title": "Air Transportation Tax Regulations, 1992",
      "citation": "SOR/93-294",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-221",
      "title": "Air Transportation Tax Remission Order",
      "citation": "SI/83-221",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-385",
      "title": "Aklavik Airport Zoning Regulations",
      "citation": "SOR/93-385",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-412",
      "title": "Akwesasne Residents Remission Order",
      "citation": "SOR/91-412",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-338",
      "title": "Alaska Marine Lines, Inc. Remission Order",
      "citation": "SOR/98-338",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-251",
      "title": "Albania and the Sultanate of Oman Goods Remission Order",
      "citation": "SOR/2001-251",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-145",
      "title": "Alberta Chicken Order",
      "citation": "SOR/99-145",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-129",
      "title": "Alberta Egg Marketing Levies Order",
      "citation": "CRC, c 129",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-128",
      "title": "Alberta Egg Order",
      "citation": "CRC, c 128",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-752",
      "title": "Alberta Equivalency Order",
      "citation": "SOR/94-752",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-246",
      "title": "Alberta Fishery Regulations, 1998",
      "citation": "SOR/98-246",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-21",
      "title": "Order Designating Alberta for the Purposes of the Criminal Interest Rate Provisions of the Criminal Code",
      "citation": "SOR/2010-21",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-190",
      "title": "Alberta Hog Marketing Levies Order",
      "citation": "SOR/84-190",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-130",
      "title": "Alberta Hog Order",
      "citation": "CRC, c 130",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-719",
      "title": "Alberta Milk Order",
      "citation": "SOR/94-719",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-1117",
      "title": "Alberta Potato Marketing Levies (Interprovincial and Export) Order",
      "citation": "SOR/85-1117",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-133",
      "title": "Alberta Potato Order",
      "citation": "CRC, c 133",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-599",
      "title": "Alberta Rules of Practice Respecting Reduction in the Number of Years of Imprisonment without Eligibility for Parole",
      "citation": "SOR/88-599",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-239",
      "title": "Alberta Sex Offender Information Registration Regulations",
      "citation": "SOR/2004-239",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-135",
      "title": "Alberta Turkey Marketing Levies Order",
      "citation": "CRC, c 135",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-134",
      "title": "Alberta Turkey Order",
      "citation": "CRC, c 134",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-124",
      "title": "Alexander First Nation Treaty Land Entitlement Remission Order",
      "citation": "SI/2003-124",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1375",
      "title": "Algoma Central Railway Traffic Rules and Regulations",
      "citation": "CRC, c 1375",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-305",
      "title": "Allocation Method Order (2008) – Softwood Lumber Products",
      "citation": "SOR/2007-305",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-10",
      "title": "Allocation Method Order (2009) – Softwood Lumber Products",
      "citation": "SOR/2009-10",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-320",
      "title": "Allocation Method Order (2010) — Softwood Lumber Products",
      "citation": "SOR/2009-320",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-278",
      "title": "Allocation Method Order (2011) — Softwood Lumber Products",
      "citation": "SOR/2010-278",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-269",
      "title": "Allocation Method Order (2012) — Softwood Lumber Products",
      "citation": "SOR/2011-269",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-248",
      "title": "Allocation Method Order (2013) — Softwood Lumber Products",
      "citation": "SOR/2012-248",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-236",
      "title": "Allocation Method Order (2014) — Softwood Lumber Products",
      "citation": "SOR/2013-236",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-321",
      "title": "Allocation Method Order (2015) — Softwood Lumber Products",
      "citation": "SOR/2014-321",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-186",
      "title": "Allocation Method Order (Beef and Veal)",
      "citation": "SOR/96-186",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-166",
      "title": "Allocation Method Order – Softwood Lumber Products",
      "citation": "SOR/2007-166",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-389",
      "title": "Allocation Method Order — Turkey and Turkey Products",
      "citation": "SOR/96-389",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-196",
      "title": "Allocation Methods Order — Cheese and Cheese Products, Ice Cream, Yogurt, Powdered Buttermilk and Concentrated Milk",
      "citation": "SOR/95-196",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-197",
      "title": "Allocation Methods Order — Hatching Eggs, Live Broilers, Eggs and Egg Products",
      "citation": "SOR/95-197",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-535",
      "title": "Allowance Exclusion Order",
      "citation": "SOR/92-535",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-345",
      "title": "Alpine Joe Sportswear Remission Order, 1998",
      "citation": "SOR/98-345",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-453",
      "title": "Alternative Fuels Regulations",
      "citation": "SOR/96-453",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-31",
      "title": "Regulations Prescribing Alternative Means of Publication (Foreign Companies)",
      "citation": "SOR/2003-31",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-132",
      "title": "Aluminerie de Bécancour Remission Order",
      "citation": "SI/89-132",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-33",
      "title": "Amalgamations and Windings-Up Continuation (GST/HST) Regulations",
      "citation": "SOR/91-33",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-11",
      "title": "American Bases in Newfoundland Remission Order, 1990",
      "citation": "SI/91-11",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-185",
      "title": "American Consumption of Softwood Lumber Products Regulations",
      "citation": "SOR/2008-185",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-598",
      "title": "Ammonium Nitrate and Fuel Oil Order",
      "citation": "CRC, c 598",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1145",
      "title": "Ammonium Nitrate Storage Facilities Regulations",
      "citation": "CRC, c 1145",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-467",
      "title": "Order Declaring an Amnesty Period",
      "citation": "SOR/98-467",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-95",
      "title": "Order Declaring an Amnesty Period (2006)",
      "citation": "SOR/2006-95",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-56",
      "title": "Order Declaring an Amnesty Period (2014)",
      "citation": "SOR/2014-56",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-101",
      "title": "Anchorage Regulations",
      "citation": "SOR/88-101",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-366",
      "title": "Ancillary Activities (Insurance Companies, Canadian Societies and Insurance Holding Companies) Regulations",
      "citation": "SOR/2001-366",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1146",
      "title": "Anhydrous Ammonia Bulk Storage Regulations",
      "citation": "CRC, c 1146",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-78",
      "title": "Animals of the Sub-family Bovinae and Their Products Importation Prohibition Regulations",
      "citation": "SOR/2005-78",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-795",
      "title": "Annex Code 9075 Sales Tax Remission Order",
      "citation": "CRC, c 795",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-60",
      "title": "Order Specifying Limits on the Annual Aggregate Quantity of Roses of Tariff Item No. 0603.10.11 that are Entitled to the Canada-Israel Agreement Tariff",
      "citation": "SOR/97-60",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-230",
      "title": "Annual Statement (Banks and Bank Holding Companies) Regulations",
      "citation": "SOR/2010-230",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-231",
      "title": "Annual Statement (Cooperative Credit Associations) Regulations",
      "citation": "SOR/2010-231",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-233",
      "title": "Annual Statement (Insurance Companies and Insurance Holding Companies) Regulations",
      "citation": "SOR/2010-233",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-232",
      "title": "Annual Statement (Trust and Loan Companies) Regulations",
      "citation": "SOR/2010-232",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-319",
      "title": "Annuities Agents Pension Regulations",
      "citation": "CRC, c 319",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-985",
      "title": "Annuities Payable to Survivors and Children of Judges Regulations",
      "citation": "CRC, c 985",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-363",
      "title": "Antarctic Environmental Protection Regulations",
      "citation": "SOR/2003-363",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-324",
      "title": "Regulations Respecting Anti-Competitive Acts of Persons Operating a Domestic Service",
      "citation": "SOR/2000-324",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-464",
      "title": "Regulations Prescribing Antique Firearms",
      "citation": "SOR/98-464",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-143",
      "title": "Apple Stabilization 1977 Regulations",
      "citation": "SOR/79-143",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-1034",
      "title": "Apple Stabilization Regulations, 1980",
      "citation": "SOR/81-1034",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-287",
      "title": "Apple Stabilization Regulations, 1983 and 1984",
      "citation": "SOR/87-287",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-21",
      "title": "Apple Stabilization Regulations, 1982-83",
      "citation": "SOR/84-21",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-106",
      "title": "Order Suspending the Application of Concessions on Imports of Certain Products Originating in the United States",
      "citation": "SOR/2005-106",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-51-84",
      "title": "Application of Defence Services Pension Act to Special Force",
      "citation": "SOR/51-84",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-312",
      "title": "Application of Provincial Laws Regulations",
      "citation": "SOR/96-312",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-253",
      "title": "Application of Regulations made under paragraph 33(1)(m) or (n) of the Northwest Territories Waters Act in Nunavut Order",
      "citation": "SOR/2002-253",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-191",
      "title": "Applications for Authorization under Paragraph 35(2)(b) of the Fisheries Act Regulations",
      "citation": "SOR/2013-191",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-416",
      "title": "Regulations Respecting Applications for Ministerial Review — Miscarriages of Justice",
      "citation": "SOR/2002-416",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-276",
      "title": "Regulations Respecting Applications for Permits for Disposal at Sea",
      "citation": "SOR/2001-276",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1333",
      "title": "Application to Canadian Penitentiary Service Regulations",
      "citation": "CRC, c 1333",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-91",
      "title": "Proclamation Announcing the Appointment of the Governor General",
      "citation": "SI/2005-91",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-22",
      "title": "Proclamation Announcing the Appointment of the Governor General",
      "citation": "SI/95-22",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-334",
      "title": "Appointment of Women in the Department of Indian Affairs and Northern Development (Employment Equity Program) Regulations",
      "citation": "SOR/94-334",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-27",
      "title": "Appointment or Deployment of Alternates Exclusion Approval Order",
      "citation": "SI/2012-27",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-83",
      "title": "Appointment or Deployment of Alternates Regulations",
      "citation": "SOR/2012-83",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-255",
      "title": "Apprentice Loans Regulations",
      "citation": "SOR/2014-255",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-79",
      "title": "Order Designating the Appropriate Authority for a Province With Respect to the Act",
      "citation": "SI/95-79",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-320",
      "title": "Appropriation Act No. 1, 1977, Leasing Regulations",
      "citation": "CRC, c 320",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-201",
      "title": "Approved Breath Analysis Instruments Order",
      "citation": "SI/85-201",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-200",
      "title": "Approved Screening Devices Order",
      "citation": "SI/85-200",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-168",
      "title": "Apricot Stabilization Regulations, 1977",
      "citation": "SOR/78-168",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-607",
      "title": "Apricot Stabilization Regulations, 1982",
      "citation": "SOR/83-607",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-353",
      "title": "Arctic Shipping Pollution Prevention Regulations",
      "citation": "CRC, c 353",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-30",
      "title": "Arctic Star Order",
      "citation": "SI/2014-30",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-417",
      "title": "Arctic Waters Experimental Pollution Regulations, 1978",
      "citation": "SOR/78-417",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-9",
      "title": "Arctic Waters Experimental Pollution Regulations, 1979",
      "citation": "SOR/80-9",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-276",
      "title": "Arctic Waters Experimental Pollution Regulations, 1982",
      "citation": "SOR/82-276",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-832",
      "title": "Arctic Waters Experimental Pollution Regulations, 1982 (Dome Petroleum)",
      "citation": "SOR/82-832",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-354",
      "title": "Arctic Waters Pollution Prevention Regulations",
      "citation": "CRC, c 354",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-543",
      "title": "Area Control List",
      "citation": "SOR/81-543",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1274",
      "title": "Armed Forces Postal Regulations",
      "citation": "CRC, c 1274",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-152",
      "title": "Article 5 North Atlantic Treaty Organisation (NATO) Medal for Operation \"Active Endeavour\" Order",
      "citation": "SI/2003-152",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-153",
      "title": "Article 5 North Atlantic Treaty Organisation (NATO) Medal for Operation \"Eagle Assist\" Order",
      "citation": "SI/2003-153",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-25",
      "title": "Artists\u0027 Representatives (GST/HST) Regulations",
      "citation": "SOR/91-25",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-793",
      "title": "Arviat Airport Zoning Regulations",
      "citation": "SOR/90-793",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-341",
      "title": "Asbestos Mines and Mills Release Regulations",
      "citation": "SOR/90-341",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-260",
      "title": "Asbestos Products Regulations",
      "citation": "SOR/2007-260",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-137",
      "title": "Ash-Free Zone Regulations",
      "citation": "SOR/2004-137",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1305",
      "title": "Asian Development Bank Privileges and Immunities Order",
      "citation": "CRC, c 1305",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-101",
      "title": "Order Setting Aside or Referring Back to the CRTC Certain Decisions Respecting Various Undertakings",
      "citation": "SI/96-101",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-379",
      "title": "Assessable Activities, Exceptions and Executive Committee Projects Regulations",
      "citation": "SOR/2005-379",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-177",
      "title": "Assessment of Financial Institutions Regulations, 2001",
      "citation": "SOR/2001-177",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-317",
      "title": "Assessment of Pension Plans Regulations",
      "citation": "SOR/2011-317",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-293",
      "title": "Assessor\u0027s Rules of Procedure",
      "citation": "SOR/2003-293",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-450",
      "title": "Assets (Foreign Companies) Regulations",
      "citation": "SOR/2002-450",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-150",
      "title": "Assigning the Administration, Management and Control of Certain Lands from the Minister of Indian Affairs and Northern Development to the Minister of the Environment",
      "citation": "SI/86-150",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-20",
      "title": "Order Assigning the Honourable Gail Shea to Assist the member of the Queen\u0027s Privy Council for Canada",
      "citation": "SI/2013-20",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-21",
      "title": "Order Assigning the Honourable Steven Blaney to Assist the Minister of Foreign Affairs",
      "citation": "SI/2013-21",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-80-114",
      "title": "Assigning to the Minister of the Environment the Administration, Management and Control of Certain Public Lands",
      "citation": "SI/80-114",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-37",
      "title": "Assigning to the Minister of the Environment, the Administration, Management and Control of Certain Public Lands",
      "citation": "SI/85-37",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-138",
      "title": "Order Assigning to the Minister of the Environment the Administration, Management and Control of Certain Public Lands",
      "citation": "SI/88-138",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-151",
      "title": "Order Assigning to the Minister of the Environment the Administration, Management and Control of Certain Public Lands",
      "citation": "SI/89-151",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-107",
      "title": "Order Assigning to the Minister of the Environment, the Administration, Management and Control of Certain Public Lands",
      "citation": "SI/92-107",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-84-62",
      "title": "Assigning to the Minister of the Environment, the Administration, Management and Control of Certain Public Lands",
      "citation": "SI/84-62",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-84-77",
      "title": "Assigning to the Minister of the Environment, the Administration, Management and Control of Certain Public Lands",
      "citation": "SI/84-77",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-82-197",
      "title": "Assigning to the Minister of the Environment, the Administration, Management and Control of Certain Public Lands",
      "citation": "SI/82-197",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-142",
      "title": "Assigning to the Minister of the Environment the Administration, Management and Control of Certain Public Lands",
      "citation": "SI/85-142",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-80-129",
      "title": "Assigning to the Minister of the Environment the Administration, Management and Control of Certain Public Lands",
      "citation": "SI/80-129",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-81-90",
      "title": "Assigning to the Minister of the Environment the Administration, Management and Control of Certain Public Lands",
      "citation": "SI/81-90",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-82-30",
      "title": "Assigning to the Minister of the Environment the Administration, Management and Control of Certain Public Lands",
      "citation": "SI/82-30",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-56",
      "title": "Order Assigning to the Minister of the Environment the Administration of Certain Public Lands",
      "citation": "SI/95-56",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-194",
      "title": "Order Assigning to the Public Service Commission the Duty to Investigate Public Service Employee Complaints Respecting Personal Harassment",
      "citation": "SI/86-194",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-675",
      "title": "Assignment of Crown Debt Regulations",
      "citation": "CRC, c 675",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-19",
      "title": "Order Terminating the Assignment of the Honourable Bernard Valcourt",
      "citation": "SI/2013-19",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1578",
      "title": "Assistance Fund (W.V.A. and C.W.A.) Regulations",
      "citation": "CRC, c 1578",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-137",
      "title": "Assisted Human Reproduction (Section 8 Consent) Regulations",
      "citation": "SOR/2007-137",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-76-40",
      "title": "Astoria River Water Power Regulations",
      "citation": "SOR/76-40",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-524",
      "title": "Atlantic Enterprise Loan Insurance Regulations",
      "citation": "SOR/86-524",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-21",
      "title": "Atlantic Fishery Regulations, 1985",
      "citation": "SOR/86-21",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1004",
      "title": "Atlantic Pilotage Authority Non-compulsory Area Regulations",
      "citation": "SOR/86-1004",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1355",
      "title": "Atlantic Pilotage Authority Pension Regulations",
      "citation": "CRC, c 1355",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1264",
      "title": "Atlantic Pilotage Authority Regulations",
      "citation": "CRC, c 1264",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-586",
      "title": "Atlantic Pilotage Tariff Regulations, 1996",
      "citation": "SOR/95-586",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-710",
      "title": "Atlantic Pilotage Tariff Regulations — Newfoundland and Labrador Non-Compulsory Areas",
      "citation": "SOR/81-710",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-78-75",
      "title": "Atlantic Region Ferry Remission Order",
      "citation": "SI/78-75",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-39",
      "title": "Atomic Energy Control Board Cost Recovery Fees Remission Order",
      "citation": "SI/94-39",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-500",
      "title": "Atomic Energy of Canada Limited Authorization Order, 1988",
      "citation": "SOR/88-500",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-652",
      "title": "Atomic Energy of Canada Limited Authorization Order, 1988-2",
      "citation": "SOR/88-652",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1356",
      "title": "Atomic Energy of Canada Limited Pension Regulations",
      "citation": "CRC, c 1356",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-465",
      "title": "Australia and New Zealand Tariff Preference Maintenance Order",
      "citation": "SOR/95-465",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-35",
      "title": "Australia Tariff and New Zealand Tariff Rules of Origin Regulations",
      "citation": "SOR/98-35",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-31",
      "title": "Authority to Sell Veterinary Drugs Fees Regulations",
      "citation": "SOR/95-31",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-31",
      "title": "Authority to Sell Veterinary Drugs Fees Regulations",
      "citation": "SOR/95-31",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-573",
      "title": "Authorization of the Sale of the Shares of Nordair Limited Order",
      "citation": "SOR/84-573",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-207",
      "title": "Authorizations to Carry Restricted Firearms and Certain Handguns Regulations",
      "citation": "SOR/98-207",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-206",
      "title": "Authorizations to Transport Restricted Firearms and Prohibited Firearms Regulations",
      "citation": "SOR/98-206",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-575",
      "title": "Automatic Firearms Country Control List",
      "citation": "SOR/91-575",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-176",
      "title": "Automobile Operating Expense Benefit (GST/HST) Regulations",
      "citation": "SOR/99-176",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-610",
      "title": "Automotive Goods Rules of Origin Regulations",
      "citation": "SOR/93-610",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-966",
      "title": "Automotive Manufacturing Assistance Regulations",
      "citation": "CRC, c 966",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-36",
      "title": "Automotive Parts Tariff Removal Order, 1988",
      "citation": "SOR/89-36",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-132",
      "title": "Aveos Pension Plan Regulations",
      "citation": "SOR/2013-132",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-87",
      "title": "Aviation Occupational Health and Safety Regulations",
      "citation": "SOR/2011-87",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-182",
      "title": "Aviation Occupational Safety and Health Regulations",
      "citation": "SOR/87-182",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-67",
      "title": "Order Awarding of a Bar “ALLIED FORCE” to the General Campaign Star (Federal Republic of Yugoslavia, Albania, the Former Yugoslav Republic of Macedonia and Adriatic and Ionian Seas)",
      "citation": "SI/2004-67",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-68",
      "title": "Order Awarding of a Bar \"ALLIED FORCE\" to the General Service Medal (Aviano and Vicenza, Italy)",
      "citation": "SI/2004-68",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-66",
      "title": "Award Regulations",
      "citation": "SOR/96-66",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-293",
      "title": "Bagotville Airport Zoning Regulations",
      "citation": "SOR/93-293",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-75",
      "title": "Baie Comeau Airport Zoning Regulations",
      "citation": "CRC, c 75",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-896",
      "title": "Baie Verte Mines Inc. Regulations",
      "citation": "SOR/82-896",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-262",
      "title": "Baie Verte Mines Inc. Regulations, 1985",
      "citation": "SOR/86-262",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-180",
      "title": "Bait Services Fee Order",
      "citation": "SOR/96-180",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-794",
      "title": "Baker Lake Airport Zoning Regulations",
      "citation": "SOR/90-794",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-129",
      "title": "Ballast Water Control and Management Regulations",
      "citation": "SOR/2006-129",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-237",
      "title": "Ballast Water Control and Management Regulations",
      "citation": "SOR/2011-237",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-743",
      "title": "Ballet Shoes Remission Order",
      "citation": "CRC, c 743",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-28",
      "title": "Order Exempting Bands from the Operation of Section 32 of the Indian Act",
      "citation": "SOR/2010-28",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-199",
      "title": "Bank Holding Company Proposal Regulations",
      "citation": "SOR/2004-199",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-92",
      "title": "Banking Industry Commission-paid Salespeople Hours of Work Regulations",
      "citation": "SOR/2006-92",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-298",
      "title": "Bank of Canada Notes Regulations",
      "citation": "SOR/89-298",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-368",
      "title": "Bankruptcy and Insolvency General Rules",
      "citation": "CRC, c 368",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-133",
      "title": "Barley 1987 Period Stabilization Regulations",
      "citation": "SOR/89-133",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-542",
      "title": "Barley Stabilization Regulations, 1977-78",
      "citation": "SOR/79-542",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-69",
      "title": "Order Awarding of a Bar with the NATO Star Flanked with the Letters \"ISAF\" and \"FIAS\" to the General Campaign Star (Afghanistan)",
      "citation": "SI/2004-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-70",
      "title": "Order Awarding of a Bar with the NATO Star Flanked with the Letters \"ISAF\" and \"FIAS\" to the General Service Medal (Afghanistan)",
      "citation": "SI/2004-70",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-194",
      "title": "Base Metal Coins Regulations, 1996",
      "citation": "SOR/96-194",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-293",
      "title": "Basin Head Marine Protected Area Regulations",
      "citation": "SOR/2005-293",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-446",
      "title": "B.C. Place Replacement Turf Remission Order, 1995",
      "citation": "SOR/95-446",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-677",
      "title": "B.C. Tree Fruit Export Regulations",
      "citation": "SOR/81-677",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-204",
      "title": "Beef Calf Stabilization Regulations",
      "citation": "SOR/78-204",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-103",
      "title": "Beef Cattle Research, Market Development and Promotion Levies Order",
      "citation": "SOR/2005-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-158",
      "title": "Beef Cattle Research, Market Development and Promotion Levies Order",
      "citation": "SOR/2010-158",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-418",
      "title": "Beer Originating in the United States Remission Order, 1993",
      "citation": "SOR/93-418",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-493",
      "title": "Benzene in Gasoline Regulations",
      "citation": "SOR/97-493",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-217",
      "title": "Benzodiazepines and Other Targeted Substances Regulations",
      "citation": "SOR/2000-217",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-201",
      "title": "Billiard Cloth Remission Order, 1996",
      "citation": "SOR/96-201",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-4",
      "title": "Blainville Motor Vehicle Test Centre Benefit Eligibility Protection Regulations",
      "citation": "SOR/99-4",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-178",
      "title": "Blood Regulations",
      "citation": "SOR/2013-178",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-37",
      "title": "Order Approving Blood Sample Containers",
      "citation": "SOR/2005-37",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-332",
      "title": "Blouses and Shirts Remission Order",
      "citation": "SOR/88-332",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-89",
      "title": "Blouses, Shirts and Co-ordinates Remission Order, 1998",
      "citation": "SOR/98-89",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1405",
      "title": "Board of Steamship Inspection Scale of Fees",
      "citation": "CRC, c 1405",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-280",
      "title": "Boat and Fire Drill and Means of Exit Regulations",
      "citation": "SOR/2005-280",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1407",
      "title": "Boating Restriction Regulations",
      "citation": "CRC, c 1407",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-142",
      "title": "Regulations Designating a Body for the Purposes of Paragraph 91(2)(c) of the Immigration and Refugee Protection Act",
      "citation": "SOR/2011-142",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-124",
      "title": "Bolster and Sideframe Remission Order",
      "citation": "SOR/93-124",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-324",
      "title": "Book Importation Regulations",
      "citation": "SOR/99-324",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-676",
      "title": "Borrowing of Defence Materials and Equipment Order",
      "citation": "CRC, c 676",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-281",
      "title": "Borrowing (Property and Casualty Companies and Marine Companies) Regulations",
      "citation": "SOR/92-281",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-281",
      "title": "Borrowing (Property and Casualty Companies and Marine Companies) Regulations",
      "citation": "SOR/92-281",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-592",
      "title": "Proclamation Designating Botswana as a Designated State for Purposes of the Act",
      "citation": "SOR/87-592",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-215",
      "title": "Bottled Domestic Spirits Remission Order",
      "citation": "SI/85-215",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-45",
      "title": "Order Fixing the Boundaries of the Town of Banff in Banff National Park and Adding a Description of the Boundaries as Schedule IV to the Act",
      "citation": "SOR/90-45",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-167",
      "title": "Boundary Bay Airport Zoning Regulations",
      "citation": "SOR/81-167",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-124",
      "title": "Bowie Seamount Marine Protected Area Regulations",
      "citation": "SOR/2008-124",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-104",
      "title": "Notice of Branch Closure (Banks) Regulations",
      "citation": "SOR/2002-104",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-105",
      "title": "Notice of Branch Closure (Cooperative Credit Associations) Regulations",
      "citation": "SOR/2002-105",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-106",
      "title": "Notice of Branch Closure (Trust and Loan Companies) Regulations",
      "citation": "SOR/2002-106",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-509",
      "title": "Brandon Airport Zoning Regulations",
      "citation": "SOR/88-509",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-67",
      "title": "Proclamation Giving Notice that the Agreement on Social Security between Canada and the Federative Republic of Brazil Comes into Force on August 1, 2014",
      "citation": "SI/2014-67",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-566",
      "title": "Brewery Departmental Regulations",
      "citation": "CRC, c 566",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-565",
      "title": "Brewery Regulations",
      "citation": "CRC, c 565",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-51",
      "title": "Bridgestone Remission Order",
      "citation": "SI/89-51",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-424",
      "title": "British Columbia Chicken Order",
      "citation": "SOR/97-424",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-137",
      "title": "British Columbia Court of Appeal Criminal Appeal Rules, 1986",
      "citation": "SI/86-137",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-33",
      "title": "British Columbia Cranberry Order",
      "citation": "SOR/2011-33",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-140",
      "title": "British Columbia Egg Order",
      "citation": "CRC, c 140",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-1",
      "title": "British Columbia Forestry Revitalization Remission Order",
      "citation": "SI/2013-1",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-158",
      "title": "British Columbia Hog Order",
      "citation": "SOR/90-158",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-121",
      "title": "British Columbia HST Regulations",
      "citation": "SOR/2011-121",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-142",
      "title": "British Columbia Interior Vegetable Marketing Board (Interprovincial and Export) Regulations",
      "citation": "CRC, c 142",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-511",
      "title": "British Columbia Milk Order",
      "citation": "SOR/94-511",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-647",
      "title": "British Columbia Mushroom Marketing (Interprovincial and Export) Regulations",
      "citation": "SOR/80-647",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-444",
      "title": "British Columbia Mushroom Order",
      "citation": "SOR/78-444",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-1034",
      "title": "British Columbia Peach Stabilization Regulations, 1983",
      "citation": "SOR/85-1034",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-591",
      "title": "British Columbia Peach Stabilization Regulations, 1987",
      "citation": "SOR/88-591",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-592",
      "title": "British Columbia Pear Stabilization Regulations, 1987",
      "citation": "SOR/88-592",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-1033",
      "title": "British Columbia Pear Stabilization Regulations, 1983-84",
      "citation": "SOR/85-1033",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-1035",
      "title": "British Columbia Prune-plum Stabilization Regulations, 1983",
      "citation": "SOR/85-1035",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-593",
      "title": "British Columbia Prunes Stabilization Regulations, 1987",
      "citation": "SOR/88-593",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-569",
      "title": "British Columbia Rules of Practice Respecting Reduction in the Number of Years of Imprisonment Without Eligibility for Parole",
      "citation": "SOR/97-569",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-2",
      "title": "British Columbia Sex Offender Information Registration Regulations",
      "citation": "SOR/2005-2",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-137",
      "title": "British Columbia Sport Fishing Regulations, 1996",
      "citation": "SOR/96-137",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-818",
      "title": "British Columbia Tree Fruit Order",
      "citation": "SOR/79-818",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-147",
      "title": "British Columbia Tree Fruit Pooling Regulations",
      "citation": "CRC, c 147",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-150",
      "title": "British Columbia Turkey Marketing Board (Interprovincial and Export) Order",
      "citation": "CRC, c 150",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-149",
      "title": "British Columbia Turkey Marketing Levies Order",
      "citation": "CRC, c 149",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-148",
      "title": "British Columbia Turkey Order",
      "citation": "CRC, c 148",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-244",
      "title": "British Columbia Vegetable Marketing Levies Order",
      "citation": "SOR/2008-244",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-49",
      "title": "British Columbia Vegetable Order",
      "citation": "SOR/81-49",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-155",
      "title": "British Preferential Tariff Direct Shipment Without Transhipment Exemption Order—Hong Kong Special Administrative Region of the People\u0027s Republic of China",
      "citation": "SOR/97-155",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-555",
      "title": "Broadcasting Distribution Regulations",
      "citation": "SOR/97-555",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-430",
      "title": "Broadcasting Industry Commission Salesmen Hours of Work Regulations",
      "citation": "SOR/79-430",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-420",
      "title": "Broadcasting Information Regulations, 1993",
      "citation": "SOR/93-420",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-144",
      "title": "Broadcasting Licence Fee Regulations, 1997",
      "citation": "SOR/97-144",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-281",
      "title": "Broadcast Technical Data Services Fees Order",
      "citation": "SOR/82-281",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-222",
      "title": "Burlington Canal Regulations",
      "citation": "SOR/89-222",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-345",
      "title": "Burwash Airport Zoning Regulations",
      "citation": "SOR/95-345",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-347",
      "title": "2- Butoxyethanol Regulations",
      "citation": "SOR/2006-347",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-27",
      "title": "Cabbage Stabilization Regulations, 1982-83",
      "citation": "SOR/84-27",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-221",
      "title": "Calculation of Contribution Rates Regulations",
      "citation": "SOR/89-221",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-50",
      "title": "Calculation of Contribution Rates Regulations, 2007",
      "citation": "SOR/2008-50",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-593",
      "title": "Calculation of Default Contribution Rates Regulations",
      "citation": "SOR/98-593",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-631",
      "title": "Calculation of Interest Regulations",
      "citation": "SOR/87-631",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-77",
      "title": "Calgary International Airport Zoning Regulations",
      "citation": "CRC, c 77",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-184",
      "title": "Cambridge Bay Airport Zoning Regulations",
      "citation": "SOR/85-184",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-127",
      "title": "Campbell River Airport Zoning Regulations",
      "citation": "SOR/84-127",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-133",
      "title": "Camp Ipperwash Indian Settlement Remission Order, 2003",
      "citation": "SI/2003-133",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-22",
      "title": "Canada Assistance Plan (British Columbia) Remission Order",
      "citation": "SI/96-22",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-136",
      "title": "Order Transferring to the Canada Border Services Agency the Control and Supervision of Certain Portions in the Department of Citizenship and Immigration",
      "citation": "SI/2004-136",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-512",
      "title": "Canada Business Corporations Regulations, 2001",
      "citation": "SOR/2001-512",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-320",
      "title": "Canada-Chile Free Trade Agreement Fruit and Vegetable Aggregate Quantity Limit Order",
      "citation": "SOR/97-320",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-127",
      "title": "Canada Communication Group Divestiture Regulations",
      "citation": "SOR/97-127",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-256",
      "title": "Canada Cooperatives Regulations",
      "citation": "SOR/99-256",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-424",
      "title": "Canada Corporations Regulations",
      "citation": "CRC, c 424",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-8",
      "title": "Canada Customs and Revenue Agency Regulations",
      "citation": "SOR/2004-8",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-69",
      "title": "Canada Customs and Revenue Agency Regulations, No. 2",
      "citation": "SOR/2004-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-415",
      "title": "Canada Cycle and Motor Company Limited Enterprise Development Regulations",
      "citation": "SOR/78-415",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-236",
      "title": "Canada Deposit Insurance Corporation Application for Deposit Insurance By-law",
      "citation": "SOR/2006-236",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-515",
      "title": "Canada Deposit Insurance Corporation Application for Deposit Insurance By-law",
      "citation": "SOR/93-515",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-292",
      "title": "Canada Deposit Insurance Corporation Data and System Requirements By-law",
      "citation": "SOR/2010-292",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-542",
      "title": "Canada Deposit Insurance Corporation Deposit Insurance Information By-law",
      "citation": "SOR/96-542",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-516",
      "title": "Canada Deposit Insurance Corporation Deposit Insurance Policy By-law",
      "citation": "SOR/93-516",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-120",
      "title": "Canada Deposit Insurance Corporation Differential Premiums By-law",
      "citation": "SOR/99-120",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-279",
      "title": "Canada Deposit Insurance Corporation Joint and Trust Account Disclosure By-Law",
      "citation": "SOR/95-279",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-177",
      "title": "Canada Deposit Insurance Corporation Notice Regulations (Compensation in Respect of the Restructuring of Federal Member Institutions)",
      "citation": "SOR/2000-177",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-142",
      "title": "Canada Deposit Insurance Corporation Prescribed Practices Premium Surcharge By-law",
      "citation": "SOR/94-142",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-629",
      "title": "Determining that the Canada Development Corporation Need not Apply for a Certificate of Continuance Under the Act",
      "citation": "SOR/80-629",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-706",
      "title": "Canada Development Investment Corporation Registered Office Relocation Order",
      "citation": "SOR/90-706",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-186",
      "title": "Canada Disability Savings Regulations",
      "citation": "SOR/2008-186",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-151",
      "title": "Canada Education Savings Regulations",
      "citation": "SOR/2005-151",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-20",
      "title": "Regulation Adapting the Canada Elections Act for the Purposes of a Referendum",
      "citation": "SOR/2010-20",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-168",
      "title": "Order Transferring to the Canada Employment and Immigration Commission From the Department of External Affairs the Control and Supervision of the Consular, Immigration and Passport Affairs Branch and the Immigration Sections of the Programs Division",
      "citation": "SI/92-168",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-81-11",
      "title": "Canada Federal Court Reports Distribution Order",
      "citation": "SI/81-11",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-613",
      "title": "Canada Gazette (1978) Special Issue Regulations",
      "citation": "SOR/79-613",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-58",
      "title": "Canada Gazette Publication Order",
      "citation": "SI/2003-58",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-19",
      "title": "Canada Gazette Publication Order, 2014",
      "citation": "SI/2014-19",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-889",
      "title": "Canada Grain Regulations",
      "citation": "CRC, c 889",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-468",
      "title": "Canada Health and Social Transfer Regulations",
      "citation": "SOR/97-468",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-62",
      "title": "Canada Health Transfer and Canada Social Transfer Regulations",
      "citation": "SOR/2004-62",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-520",
      "title": "Canada Industrial Relations Board Regulations, 2012",
      "citation": "SOR/2001-520",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-520",
      "title": "Règlement de 2001 sur le Canada Industrial Relations Board Regulations, 2012",
      "citation": "SOR/2001-520",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-54",
      "title": "Canada Industrial Relations Regulations",
      "citation": "SOR/2002-54",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1013",
      "title": "Canada Industrial Relations Remuneration Regulations",
      "citation": "CRC, c 1013",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-96",
      "title": "Order Transferring to the Canada Information Office the Control and Supervision of the Portion of the Public Service Known as the Communications Coordination Services Branch of the Department of Public Works and Government Services",
      "citation": "SI/2001-96",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-986",
      "title": "Canada Labour Standards Regulations",
      "citation": "CRC, c 986",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-142",
      "title": "Canada Lands Surveyors Regulations",
      "citation": "SOR/99-142",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2015-5",
      "title": "Canada – Newfoundland and Labrador Offshore Area Diving Operations Safety Transitional Regulations",
      "citation": "SOR/2015-5",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2015-1",
      "title": "Canada – Newfoundland and Labrador Offshore Marine Installations and Structures Occupational Health and Safety Transitional Regulations",
      "citation": "SOR/2015-1",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2015-4",
      "title": "Canada – Newfoundland and Labrador Offshore Marine Installations and Structures Transitional Regulations",
      "citation": "SOR/2015-4",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-262",
      "title": "Canada-Newfoundland Oil and Gas Spills and Debris Liability Regulations",
      "citation": "SOR/88-262",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-223",
      "title": "Canada Not-for-profit Corporations Regulations",
      "citation": "SOR/2011-223",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2015-6",
      "title": "Canada – Nova Scotia Offshore Area Diving Operations Safety Transitional Regulations",
      "citation": "SOR/2015-6",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2015-2",
      "title": "Canada – Nova Scotia Offshore Marine Installations and Structures Occupational Health and Safety Transitional Regulations",
      "citation": "SOR/2015-2",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2015-3",
      "title": "Canada – Nova Scotia Offshore Marine Installations and Structures Transitional Regulations",
      "citation": "SOR/2015-3",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-123",
      "title": "Canada-Nova Scotia Oil and Gas Spills and Debris Liability Regulations",
      "citation": "SOR/95-123",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-304",
      "title": "Canada Occupational Health and Safety Regulations",
      "citation": "SOR/86-304",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-114",
      "title": "Canada Oil and Gas Certificate of Fitness Regulations",
      "citation": "SOR/96-114",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-600",
      "title": "Canada Oil and Gas Diving Regulations",
      "citation": "SOR/88-600",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1517",
      "title": "Canada Oil and Gas Drilling and Production Regulations",
      "citation": "CRC, c 1517",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-315",
      "title": "Canada Oil and Gas Drilling and Production Regulations",
      "citation": "SOR/2009-315",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-82",
      "title": "Canada Oil and Gas Drilling Regulations",
      "citation": "SOR/79-82",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-117",
      "title": "Canada Oil and Gas Geophysical Operations Regulations",
      "citation": "SOR/96-117",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-118",
      "title": "Canada Oil and Gas Installations Regulations",
      "citation": "SOR/96-118",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1518",
      "title": "Canada Oil and Gas Land Regulations",
      "citation": "CRC, c 1518",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-149",
      "title": "Canada Oil and Gas Operations Regulations",
      "citation": "SOR/83-149",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-791",
      "title": "Canada Oil and Gas Production and Conservation Regulations",
      "citation": "SOR/90-791",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-190",
      "title": "Canada Pension Plan Investment Board Regulations",
      "citation": "SOR/99-190",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-385",
      "title": "Canada Pension Plan Regulations",
      "citation": "CRC, c 385",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-386",
      "title": "Canada Pension Plan (Social Insurance Numbers) Regulations",
      "citation": "CRC, c 386",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-318",
      "title": "Canada Port Authority Environmental Assessment Regulations",
      "citation": "SOR/99-318",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-38",
      "title": "Canada Post Corporation Pension Plan Funding Regulations",
      "citation": "SOR/2014-38",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-375",
      "title": "Canada Post Corporation Withdrawal Regulations",
      "citation": "SOR/2000-375",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-377",
      "title": "Canada Post Systems Management Limited Authorization Order, 1990",
      "citation": "SOR/90-377",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-358",
      "title": "Canada Post Systems Management Limited Incorporation Authorization Order",
      "citation": "SOR/90-358",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-62",
      "title": "Canada Production Insurance Regulations",
      "citation": "SOR/2005-62",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-29",
      "title": "Order Transferring From Canada Revenue Agency to the Department of Transport the Control and Supervision of the Royal Canadian Mint and the Canada Post Corporation",
      "citation": "SI/2006-29",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-73",
      "title": "Canada Shipping Act Gold Franc Conversion Regulations",
      "citation": "SOR/78-73",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-527",
      "title": "Canada Small Business Financing (Establishment and Operation of Capital Leasing Pilot Project) Regulations",
      "citation": "SOR/2001-527",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-141",
      "title": "Canada Small Business Financing Regulations",
      "citation": "SOR/99-141",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-329",
      "title": "Canada Student Financial Assistance Regulations",
      "citation": "SOR/95-329",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-392",
      "title": "Canada Student Loans Regulations",
      "citation": "SOR/93-392",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-467",
      "title": "Canada Travelling Exhibitions Indemnification Regulations",
      "citation": "SOR/99-467",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-658",
      "title": "Canada Turkey Marketing Levies Order",
      "citation": "CRC, c 658",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-245",
      "title": "Canada Turkey Marketing Processors Levy Order",
      "citation": "SOR/98-245",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-142",
      "title": "Canada Turkey Marketing Producers Levy Order",
      "citation": "SOR/2002-142",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1240",
      "title": "Canada-United States Nuclear Liability Rules",
      "citation": "CRC, c 1240",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-122",
      "title": "Canadian Agricultural Loans Regulations",
      "citation": "SOR/99-122",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-122",
      "title": "Canadian Agricultural Loans Regulations",
      "citation": "SOR/99-122",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-51",
      "title": "Order Renaming the Canadian Armed Forces Naval Jack as the Canadian Naval Ensign and Approving its use by the Canadian Forces",
      "citation": "SI/2013-51",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-196",
      "title": "Canadian Arsenals Limited Pension Protection Regulations",
      "citation": "SOR/87-196",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-61-507",
      "title": "Canadian Arsenals Pension Regulations",
      "citation": "SOR/61-507",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-343",
      "title": "Canadian Artists And Producers Professional Relations Tribunal Procedural Regulations",
      "citation": "SOR/2003-343",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-433",
      "title": "Canadian Aviation Regulations",
      "citation": "SOR/96-433",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-111",
      "title": "Canadian Aviation Security Regulations",
      "citation": "SOR/2000-111",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-318",
      "title": "Canadian Aviation Security Regulations, 2012",
      "citation": "SOR/2011-318",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-48",
      "title": "Canadian Beef Cattle Research, Market Development and Promotion Agency Proclamation",
      "citation": "SOR/2002-48",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-394",
      "title": "Canadian Bill of Rights Examination Regulations",
      "citation": "CRC, c 394",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-516",
      "title": "Canadian Broiler Hatching Egg and Chick Licensing Regulations",
      "citation": "SOR/87-516",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-283",
      "title": "Canadian Broiler Hatching Egg And Chick Orderly Marketing Regulations",
      "citation": "SOR/2000-283",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-512",
      "title": "Canadian Broiler Hatching Egg (Interprovincial) Pricing Regulations",
      "citation": "SOR/89-512",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-92",
      "title": "Canadian Broiler Hatching Egg Marketing Levies Order",
      "citation": "SOR/2000-92",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-781",
      "title": "Canadian Charter of Rights and Freedoms Examination Regulations",
      "citation": "SOR/85-781",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-34",
      "title": "Canadian Chicken Anti-Dumping Regulations",
      "citation": "SOR/2002-34",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-22",
      "title": "Canadian Chicken Licensing Regulations",
      "citation": "SOR/2002-22",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-35",
      "title": "Canadian Chicken Marketing Levies Order",
      "citation": "SOR/2002-35",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-36",
      "title": "Canadian Chicken Marketing Quota Regulations",
      "citation": "SOR/2002-36",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-275",
      "title": "Canadian Computer Reservation Systems (CRS) Regulations",
      "citation": "SOR/95-275",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-38",
      "title": "Order Authorizing The Canadian Corps of Commissionaires to Wear the Medal of Commissionaire Long Service Medal",
      "citation": "SI/98-38",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-448",
      "title": "Canadian Cultural Property Export Control List",
      "citation": "CRC, c 448",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-654",
      "title": "Canadian Egg Anti-dumping Pricing Regulations",
      "citation": "CRC, c 654",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-242",
      "title": "Canadian Egg Licensing Regulations, 1987",
      "citation": "SOR/87-242",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-646",
      "title": "Canadian Egg Marketing Agency Proclamation",
      "citation": "CRC, c 646",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-8",
      "title": "Canadian Egg Marketing Agency Quota Regulations, 1986",
      "citation": "SOR/86-8",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-75",
      "title": "Canadian Egg Marketing Levies Order",
      "citation": "SOR/2003-75",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-280",
      "title": "Canadian Egg Marketing Levies Order",
      "citation": "SOR/95-280",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-657",
      "title": "Canadian Egg Pricing (Interprovincial and Export) Regulations",
      "citation": "CRC, c 657",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-79-73",
      "title": "Canadian Exploration Expense Remission Order",
      "citation": "SI/79-73",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-516",
      "title": "Canadian Film Certification Regulations",
      "citation": "CRC, c 516",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-550",
      "title": "Canadian Films and Video Tapes Certification Fees Order",
      "citation": "SOR/82-550",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-80",
      "title": "Order Transferring from the Canadian Firearms Centre to the Royal Canadian Mounted Police the Control and Supervision of the Canadian Firearms Centre",
      "citation": "SI/2006-80",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-420",
      "title": "Order Specifying That the Canadian Forces are a Portion of the Public Sector Employing One Hundred or More Employees for the Purpose of Subsection 4(1) of That Act",
      "citation": "SOR/2002-420",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-78",
      "title": "Canadian Forces Base Shearwater Airfield Zoning Regulations",
      "citation": "CRC, c 78",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-64-444",
      "title": "Canadian Forces Compulsory Early Superannuation Regulations",
      "citation": "SOR/64-444",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1045",
      "title": "Canadian Forces Dental Treatment by Civilians Regulations",
      "citation": "CRC, c 1045",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-421",
      "title": "Canadian Forces Employment Equity Regulations",
      "citation": "SOR/2002-421",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-294",
      "title": "Canadian Forces Grievance Board Rules of Procedure (Review of a Grievance by Way of a Hearing)",
      "citation": "SOR/2000-294",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-679",
      "title": "Canadian Forces Laundries Order",
      "citation": "CRC, c 679",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-617",
      "title": "Canadian Forces Materiel Remission Order",
      "citation": "SOR/94-617",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-50",
      "title": "Canadian Forces Members and Veterans Re-establishment and Compensation Regulations",
      "citation": "SOR/2006-50",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-395",
      "title": "Canadian Forces Special Election Regulations",
      "citation": "CRC, c 395",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-396",
      "title": "Canadian Forces Superannuation Regulations",
      "citation": "CRC, c 396",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-66-78",
      "title": "Canadian Forces Superannuation Special Election Regulations",
      "citation": "SOR/66-78",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-40",
      "title": "Canadian Hatching Egg Producers Proclamation",
      "citation": "SOR/87-40",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-144",
      "title": "Proclamation Amending the Canadian Hatching Egg Producers Proclamation",
      "citation": "SOR/2013-144",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-144",
      "title": "Canadian Hatching Egg Producers Proclamation",
      "citation": "SOR/2013-144",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-209",
      "title": "Canadian Hatching Egg Producers Quota Regulations",
      "citation": "SOR/87-209",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-168",
      "title": "Canadian Heraldic Authority Fee Order",
      "citation": "SOR/91-168",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-68",
      "title": "Canadian Human Rights Benefit Regulations",
      "citation": "SOR/80-68",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-223",
      "title": "By-law No. 3 of the Canadian Human Rights Commission",
      "citation": "SOR/78-223",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-850",
      "title": "Canadian Industrial Renewal Regulations",
      "citation": "SOR/81-850",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-602",
      "title": "Canadian International Trade Tribunal Procurement Inquiry Regulations",
      "citation": "SOR/93-602",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-35",
      "title": "Canadian International Trade Tribunal Regulations",
      "citation": "SOR/89-35",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-499",
      "title": "Canadian International Trade Tribunal Rules",
      "citation": "SOR/91-499",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-371",
      "title": "Canadian Judicial Council Inquiries and Investigations By-laws",
      "citation": "SOR/2002-371",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1026",
      "title": "Canadian Livestock Feed Board Headquarters Regulations",
      "citation": "CRC, c 1026",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-373",
      "title": "Canadian Manufactured Goods Exported Drawback Regulations",
      "citation": "SOR/78-373",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-160",
      "title": "Proclamation Declaring June 27 of each year as “Canadian Multiculturalism Day”",
      "citation": "SI/2002-160",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-375",
      "title": "Canadian National Railway Company Wholly-Owned Subsidiaries Incorporation Authorization Order, 1990",
      "citation": "SOR/90-375",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1376",
      "title": "Canadian National Railway Passenger Train Travel Rules and Regulations",
      "citation": "CRC, c 1376",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-212",
      "title": "Canadian Nuclear Safety Commission By-laws",
      "citation": "SOR/2000-212",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-212",
      "title": "Canadian Nuclear Safety Commission Cost Recovery Fees Regulations",
      "citation": "SOR/2003-212",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-282",
      "title": "Directive to the Canadian Nuclear Safety Commission Regarding the Health of Canadians",
      "citation": "SOR/2007-282",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-211",
      "title": "Canadian Nuclear Safety Commission Rules of Procedure",
      "citation": "SOR/2000-211",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-846",
      "title": "Canadian Ownership and Control Determination Forms Order, 1985",
      "citation": "SOR/85-846",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-431",
      "title": "Canadian Ownership and Control Determination Regulations, 1984",
      "citation": "SOR/84-431",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1377",
      "title": "Canadian Pacific Railway Traffic Rules and Regulations",
      "citation": "CRC, c 1377",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-81-86",
      "title": "Canadian Passport Order",
      "citation": "SI/81-86",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-94",
      "title": "Canadian Patrol Frigate Project Remission Order",
      "citation": "SI/85-94",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-174",
      "title": "Canadian Payments Association By-law No. 1 — General",
      "citation": "SOR/2003-174",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-346",
      "title": "Canadian Payments Association By-law No. 3 — Payment Items and Automated Clearing Settlement System",
      "citation": "SOR/2003-346",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-347",
      "title": "Canadian Payments Association By-law No. 6 — Compliance",
      "citation": "SOR/2003-347",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-175",
      "title": "Canadian Payments Association By-Law No. 2 — Finance",
      "citation": "SOR/2003-175",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-215",
      "title": "Canadian Payments Association Election of Directors Regulations",
      "citation": "SOR/2002-215",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-476",
      "title": "Canadian Payments Association Membership Requirements Regulations",
      "citation": "SOR/2001-476",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-410",
      "title": "Canadian Peacekeeping Service Medal Award Regulations",
      "citation": "SOR/99-410",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-183",
      "title": "Canadian Press Pension Plan Solvency Deficiency Funding Regulations",
      "citation": "SOR/2009-183",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-245",
      "title": "Canadian Press Pension Plan Solvency Deficiency Funding Regulations, 2010",
      "citation": "SOR/2010-245",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-277",
      "title": "Canadian Radio-television and Telecommunications Commission Rules of Practice and Procedure",
      "citation": "SOR/2010-277",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-430",
      "title": "Canadian Retailers Duty Remission Order, 1993",
      "citation": "SOR/93-430",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-81",
      "title": "Canadian Security Intelligence Service Act Deputy Heads of the Public Service of Canada Order",
      "citation": "SI/93-81",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-667",
      "title": "Canadian Telecommunications Common Carrier Ownership and Control Regulations",
      "citation": "SOR/94-667",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-15",
      "title": "Canadian Tourism Commission Divestiture Regulations",
      "citation": "SOR/2004-15",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-244",
      "title": "Canadian Transportation Agency Designated Provisions Regulations",
      "citation": "SOR/99-244",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-35",
      "title": "Canadian Transportation Agency General Rules",
      "citation": "SOR/2005-35",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-104",
      "title": "Canadian Transportation Agency Rules (Dispute Proceedings and Certain Rules Applicable to All Proceedings)",
      "citation": "SOR/2014-104",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-659",
      "title": "Canadian Turkey Anti-dumping Regulations",
      "citation": "CRC, c 659",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-660",
      "title": "Canadian Turkey Licensing Regulations",
      "citation": "CRC, c 660",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-647",
      "title": "Canadian Turkey Marketing Agency Proclamation",
      "citation": "CRC, c 647",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-231",
      "title": "Canadian Turkey Marketing Quota Regulations, 1990",
      "citation": "SOR/90-231",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-67-346",
      "title": "Canadian Vickers Dry Dock Regulations",
      "citation": "SOR/67-346",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-62",
      "title": "Canadian Volunteer Service Medal Order",
      "citation": "SI/94-62",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-602",
      "title": "Canadian Wheat Board Advisory Committee 1994 Election Regulations",
      "citation": "SOR/82-602",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-69",
      "title": "Canadian Wheat Board Contingency Fund Regulations",
      "citation": "SOR/2000-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-247",
      "title": "Canadian Wheat Board Direction Order",
      "citation": "SOR/2006-247",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-227",
      "title": "Canadian Wheat Board Direction Order",
      "citation": "SOR/2011-227",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-19",
      "title": "Canadian Wheat Board (Interim Operations) Regulations",
      "citation": "SOR/2013-19",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-397",
      "title": "Canadian Wheat Board Regulations",
      "citation": "CRC, c 397",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1564",
      "title": "Canal Regulations",
      "citation": "CRC, c 1564",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-18",
      "title": "Candles Regulations",
      "citation": "SOR/2011-18",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-134",
      "title": "Canola 1987 Period Stabilization Regulations",
      "citation": "SOR/89-134",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-242",
      "title": "Canola 1988 Period Stabilization Regulations",
      "citation": "SOR/91-242",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-243",
      "title": "Canola 1989 Period Stabilization Regulations",
      "citation": "SOR/91-243",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-379",
      "title": "Canola Stabilization Regulations, 1986",
      "citation": "SOR/88-379",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-576",
      "title": "Capacity Plates and Conformity Plates Charges Order",
      "citation": "SOR/98-576",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-97",
      "title": "Cape Breton Development Corporation Employees Pre-Retirement Remission Order",
      "citation": "SI/92-97",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-98",
      "title": "Cape Dorset Airport Zoning Regulations",
      "citation": "SOR/2012-98",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-831",
      "title": "Carbonated Beverage Glass Containers Regulations",
      "citation": "SOR/80-831",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-128",
      "title": "Cargo, Fumigation and Tackle Regulations",
      "citation": "SOR/2007-128",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1306",
      "title": "Caribbean Development Bank Privileges and Immunities Order",
      "citation": "CRC, c 1306",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-79",
      "title": "Carriage by Air Act Gold Franc Conversion Regulations",
      "citation": "SOR/83-79",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-379",
      "title": "Carriages and Strollers Regulations",
      "citation": "SOR/85-379",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-803",
      "title": "Carrier Exemption Regulations",
      "citation": "CRC, c 803",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-24",
      "title": "Carrot Stabilization Regulations, 1982-83",
      "citation": "SOR/84-24",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-79",
      "title": "Cartierville Airport Zoning Regulations",
      "citation": "CRC, c 79",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-180",
      "title": "CATSA Aerodrome Designation Regulations",
      "citation": "SOR/2002-180",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-66",
      "title": "CCFTA Remission Order",
      "citation": "SOR/2002-66",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-72",
      "title": "CCFTA Remission Order, 2003",
      "citation": "SOR/2003-72",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-323",
      "title": "CCFTA Rules of Origin for Casual Goods Regulations",
      "citation": "SOR/97-323",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-340",
      "title": "CCFTA Rules of Origin Regulations",
      "citation": "SOR/97-340",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-322",
      "title": "CCFTA Tariff Preference Regulations",
      "citation": "SOR/97-322",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-344",
      "title": "CCFTA Textile and Apparel Extension of Benefit Order",
      "citation": "SOR/97-344",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-132",
      "title": "CCOFTA Rules of Origin for Casual Goods Regulations",
      "citation": "SOR/2011-132",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-131",
      "title": "CCOFTA Rules of Origin Regulations",
      "citation": "SOR/2011-131",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-133",
      "title": "CCOFTA Tariff Preference Regulations",
      "citation": "SOR/2011-133",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-217",
      "title": "CCOFTA Verification of Origin Regulations",
      "citation": "SOR/2013-217",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-398",
      "title": "CCRFTA Non-entitlement to Preference Regulations",
      "citation": "SOR/2002-398",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-396",
      "title": "CCRFTA Rules of Origin for Casual Goods Regulations",
      "citation": "SOR/2002-396",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-395",
      "title": "CCRFTA Rules of Origin Regulations",
      "citation": "SOR/2002-395",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-400",
      "title": "CCRFTA Sugar Aggregate Quantity Limit Order",
      "citation": "SOR/2002-400",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-124",
      "title": "CCRFTA Sugar Aggregate Quantity Limit Remission Order",
      "citation": "SOR/2003-124",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-397",
      "title": "CCRFTA Tariff Preference Regulations",
      "citation": "SOR/2002-397",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-129",
      "title": "CCRFTA Verification of Origin Regulations",
      "citation": "SOR/2004-129",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-681",
      "title": "Cedar Point-Christian Island Ferry Fees Regulations",
      "citation": "CRC, c 681",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-199",
      "title": "CEFTA Rules of Origin for Casual Goods Regulations",
      "citation": "SOR/2009-199",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-198",
      "title": "CEFTA Rules of Origin Regulations",
      "citation": "SOR/2009-198",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-200",
      "title": "CEFTA Tariff Preference Regulations",
      "citation": "SOR/2009-200",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-214",
      "title": "CEFTA Verification of Origin of Exported Goods Regulations",
      "citation": "SOR/2013-214",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-215",
      "title": "CEFTA Verification of Origin of Imported Goods Regulations",
      "citation": "SOR/2013-215",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-547",
      "title": "Central Registry of Divorce Proceedings Fee Order",
      "citation": "SOR/86-547",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-600",
      "title": "Central Registry of Divorce Proceedings Regulations",
      "citation": "SOR/86-600",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-80-8",
      "title": "Certain Areas Covered With Water Proclaimed Public Harbours Effective January 1, 1980",
      "citation": "SI/80-8",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-1",
      "title": "Certain Canada Port Authorities Divestiture Regulations",
      "citation": "SOR/2000-1",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-260",
      "title": "Certain Chloroprene Sheets Remission Order",
      "citation": "SOR/2002-260",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-140",
      "title": "Certain Commercial Salmon Fishing Licences Fees Remission Order",
      "citation": "SI/99-140",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-76",
      "title": "Proclamation Designating Certain Countries as States Eligible to Participate in the Military Training Assistance Program",
      "citation": "SI/93-76",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-50",
      "title": "Order Directing that Certain Documents be Discontinued",
      "citation": "SI/2005-50",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-90",
      "title": "Order Directing that Certain Documents be Discontinued",
      "citation": "SI/2000-90",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-146",
      "title": "Order Directing that Certain Documents be Discontinued",
      "citation": "SI/2003-146",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-130",
      "title": "Order Directing that Certain Documents be Discontinued",
      "citation": "SI/99-130",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-34",
      "title": "Order Directing that Certain Documents be Discontinued",
      "citation": "SI/94-34",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-30",
      "title": "Order Directing that Certain Documents be Discontinued",
      "citation": "SI/93-30",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-261",
      "title": "Order Authorizing Certain Employees of the Government of Canada to Acquire Interests in Territorial Lands in the Northwest Territories (Order No. 1, 1998)",
      "citation": "SOR/98-261",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-77",
      "title": "Certain Fees Relating to Export Certificates Remission Order",
      "citation": "SI/2011-77",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2009-95",
      "title": "Certain Fees Relating to Export Certificates Remission Order",
      "citation": "SI/2009-95",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2008-23",
      "title": "Certain Foreign Marine Carriers Remission Order",
      "citation": "SI/2008-23",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-148",
      "title": "Certain Hidden Valley Golf Resort Association Members Remission Order",
      "citation": "SI/2004-148",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-35",
      "title": "Exemption Order for Certain Licences, Authorizations and Documents (Westslope Cutthroat Trout (Alberta Population))",
      "citation": "SOR/2013-35",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2008-99",
      "title": "Certain Marine Carriers Remission Order",
      "citation": "SI/2008-99",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-66",
      "title": "Certain Marine Carriers Remission Order, 2011",
      "citation": "SI/2011-66",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-65",
      "title": "Exclusion Approval Order for Certain Persons and Certain Positions (Air Traffic Control Group), 1991",
      "citation": "SI/91-65",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-226",
      "title": "Certain Refractory Products Remission Order",
      "citation": "SOR/2001-226",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-69",
      "title": "Certain Road Vehicle Supply Remission Order",
      "citation": "SI/2001-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-203",
      "title": "Certain Ships Remission Order, 2010",
      "citation": "SOR/2010-203",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-45",
      "title": "Certain Taxpayers Remission Order, 2000-3",
      "citation": "SI/2001-45",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2009-61",
      "title": "Certificates of Age and Origin for Distilled Spirits Produced or Packaged in Canada Order",
      "citation": "SI/2009-61",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-421",
      "title": "Certification of Countries Granting Equal Copyright Protection Notice",
      "citation": "CRC, c 421",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-332",
      "title": "Certification of Origin of Goods Exported to a Free Trade Partner Regulations",
      "citation": "SOR/97-332",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-517",
      "title": "Certified Seed Potatoes Regulations",
      "citation": "CRC, c 517",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-102",
      "title": "Champagne and Aishihik First Nations (GST) Remission Order",
      "citation": "SI/2000-102",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-337",
      "title": "Charges for Services Provided by the Office of the Superintendent of Financial Institutions Regulations 2002",
      "citation": "SOR/2002-337",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-536",
      "title": "Charitable Food Donations Anti-dumping and Countervailing Duty Remission Order",
      "citation": "SOR/98-536",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-8",
      "title": "Charitable Goods Remission Order",
      "citation": "SI/98-8",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-771",
      "title": "Charlo Airport Zoning Regulations",
      "citation": "SOR/78-771",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-649",
      "title": "Charlottetown Airport Zoning Regulations",
      "citation": "SOR/92-649",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-66",
      "title": "Order Designating Charlottetown, Prince Edward Island as Head Office of the Veterans Review and Appeal Board",
      "citation": "SI/96-66",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-480",
      "title": "Charter Flights Airport Facilities Reservation Regulations",
      "citation": "SOR/82-480",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-149",
      "title": "Charts and Nautical Publications Regulations, 1995",
      "citation": "SOR/95-149",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-173",
      "title": "Chatham Airport Zoning Regulations",
      "citation": "SOR/91-173",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-375",
      "title": "Cheese Stabilization 1980-81 Regulations",
      "citation": "SOR/81-375",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-82",
      "title": "Cheese Stabilization Regulations, 1979-80",
      "citation": "SOR/80-82",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-745",
      "title": "Cheese Stabilization Regulations, 1981-82",
      "citation": "SOR/81-745",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-995",
      "title": "Cheese Stabilization Regulations, 1982-83",
      "citation": "SOR/82-995",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-674",
      "title": "Cheese Stabilization Regulations, 1983-84",
      "citation": "SOR/83-674",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-383",
      "title": "Cheese Stabilization Regulations, 1984-85",
      "citation": "SOR/84-383",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-39",
      "title": "Cheque Holding Policy Disclosure (Banks) Regulations",
      "citation": "SOR/2002-39",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-240",
      "title": "Cheque Issue Regulations, 1997",
      "citation": "SOR/97-240",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-171",
      "title": "Chesterfield Inlet Airport Zoning Regulations",
      "citation": "SOR/91-171",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-509",
      "title": "Cheticamp/Grand Etang Fishermen\u0027s Co-operative Society Limited Regulations",
      "citation": "SOR/83-509",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-218",
      "title": "CHFTA Rules of Origin for Casual Goods Regulations",
      "citation": "SOR/2014-218",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-217",
      "title": "CHFTA Rules of Origin Regulations",
      "citation": "SOR/2014-217",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-220",
      "title": "CHFTA Sugar Aggregate Quantity Limit Order",
      "citation": "SOR/2014-220",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-219",
      "title": "CHFTA Tariff Preference Regulations",
      "citation": "SOR/2014-219",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-274",
      "title": "Chicken Farmers of Canada Delegation of Authority Order",
      "citation": "SOR/2003-274",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-158",
      "title": "Chicken Farmers of Canada Proclamation",
      "citation": "SOR/79-158",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-399",
      "title": "Children of Deceased Veterans Education Assistance Regulations",
      "citation": "CRC, c 399",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-132",
      "title": "Children\u0027s Jewellery Regulations",
      "citation": "SOR/2005-132",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-19",
      "title": "Children\u0027s Jewellery Regulations",
      "citation": "SOR/2011-19",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-15",
      "title": "Children\u0027s Sleepwear Regulations",
      "citation": "SOR/2011-15",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-12",
      "title": "Children\u0027s Special Allowance Regulations",
      "citation": "SOR/93-12",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-80",
      "title": "Chilliwack Airport Zoning Regulations",
      "citation": "CRC, c 80",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-156",
      "title": "China Direct Shipment Condition Exemption Order",
      "citation": "SOR/85-156",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-811",
      "title": "Chlor-Alkali Mercury Liquid Effluent Regulations",
      "citation": "CRC, c 811",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-130",
      "title": "Chlor-Alkali Mercury Release Regulations",
      "citation": "SOR/90-130",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1147",
      "title": "Chlorine Tank Car Unloading Facilities Regulations",
      "citation": "CRC, c 1147",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-152",
      "title": "Chlorobiphenyls Regulations",
      "citation": "SOR/91-152",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-162",
      "title": "Chromium Electroplating, Chromium Anodizing and Reverse Etching Regulations",
      "citation": "SOR/2009-162",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-926",
      "title": "Chrysler Canada Ltd. Regulations, 1982",
      "citation": "SOR/82-926",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-124",
      "title": "Churchill Airport Zoning Regulations",
      "citation": "SOR/88-124",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-65",
      "title": "CIFTA Remission Order",
      "citation": "SOR/2002-65",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-73",
      "title": "CIFTA Remission Order, 2003",
      "citation": "SOR/2003-73",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-65",
      "title": "CIFTA Rules of Origin for Casual Goods Regulations",
      "citation": "SOR/97-65",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-63",
      "title": "CIFTA Rules of Origin Regulations",
      "citation": "SOR/97-63",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-64",
      "title": "CIFTA Tariff Preference Regulations",
      "citation": "SOR/97-64",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-75",
      "title": "CIFTA Verification of Origin Regulations",
      "citation": "SOR/97-75",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-178",
      "title": "Cigarette Ignition Propensity Regulations",
      "citation": "SOR/2005-178",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-194",
      "title": "Cinematographic Works (Right to Remuneration) Regulations",
      "citation": "SOR/99-194",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-138",
      "title": "Regulations Prescribing Circumstances for Granting Waivers Pursuant to Section 147 of the Act",
      "citation": "SOR/2010-138",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-565",
      "title": "CITEL Meetings Privileges and Immunities Order",
      "citation": "SOR/94-565",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2007-37",
      "title": "Citizenship Fees Remission Order",
      "citation": "SI/2007-37",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-246",
      "title": "Citizenship Regulations",
      "citation": "SOR/93-246",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-246",
      "title": "Citizenship Regulations",
      "citation": "SOR/93-246",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-177",
      "title": "Designating the City of Ottawa as the Place for the Head Office of the Cree-Naskapi Commission",
      "citation": "SI/86-177",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-64",
      "title": "Designating the Civil Aviation Tribunal as a Department; the Minister of Transport as Appropriate Minister for Purposes of the Financial Administration Act and as Minister for Purposes of Section 31 of the Aeronautics Act",
      "citation": "SI/86-64",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-594",
      "title": "Civil Aviation Tribunal Rules",
      "citation": "SOR/86-594",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1046",
      "title": "Civilian Crews on Auxiliary Vessels Direction",
      "citation": "CRC, c 1046",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-682",
      "title": "Civilian Dental Treatment Regulations",
      "citation": "CRC, c 682",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-299",
      "title": "Civil Remedies (Banks and Bank Holding Companies) Regulations",
      "citation": "SOR/2006-299",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-300",
      "title": "Civil Remedies (Cooperative Credit Associations) Regulations",
      "citation": "SOR/2006-300",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-301",
      "title": "Civil Remedies (Insurance Companies and Insurance Holding Companies) Regulations",
      "citation": "SOR/2006-301",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-302",
      "title": "Civil Remedies (Trust and Loan Companies) Regulations",
      "citation": "SOR/2006-302",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-401",
      "title": "Civil Service Insurance Regulations",
      "citation": "CRC, c 401",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-180",
      "title": "CJFTA Rules of Origin for Casual Goods Regulations",
      "citation": "SOR/2012-180",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-179",
      "title": "CJFTA Rules of Origin Regulations",
      "citation": "SOR/2012-179",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-181",
      "title": "CJFTA Tariff Preference Regulations",
      "citation": "SOR/2012-181",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-283",
      "title": "CJFTA Verification of Origin Regulations",
      "citation": "SOR/2014-283",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-300",
      "title": "CKFTA Rules of Origin for Casual Goods Regulations",
      "citation": "SOR/2014-300",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-299",
      "title": "CKFTA Rules of Origin Regulations",
      "citation": "SOR/2014-299",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-301",
      "title": "CKFTA Tariff Preference Regulations",
      "citation": "SOR/2014-301",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-195",
      "title": "Regulations Clarifying the Application of Provisions of the Convention on the Privileges and Immunities of the United Nations",
      "citation": "SOR/2002-195",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-225",
      "title": "Classed Ships Inspection Regulations, 1988",
      "citation": "SOR/89-225",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-205",
      "title": "Class II Nuclear Facilities and Prescribed Equipment Regulations",
      "citation": "SOR/2000-205",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-204",
      "title": "Class I Nuclear Facilities Regulations",
      "citation": "SOR/2000-204",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-21",
      "title": "Closely Related Corporations (GST/HST) Regulations",
      "citation": "SOR/91-21",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-97",
      "title": "Clyde River Airport Zoning Regulations",
      "citation": "SOR/2012-97",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1244",
      "title": "C.N.R. Company Exemption Order",
      "citation": "CRC, c 1244",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-97",
      "title": "Coal Mining Occupational Health and Safety Regulations",
      "citation": "SOR/90-97",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-97",
      "title": "Coal Mining Occupational Health and Safety Regulations",
      "citation": "SOR/90-97",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-413",
      "title": "Coastal Fisheries Protection Regulations",
      "citation": "CRC, c 413",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-501",
      "title": "Coast Guard Radio Station Communications Charges Order, 1994",
      "citation": "SOR/94-501",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-88",
      "title": "Coffin or Casket Remission Order",
      "citation": "SI/83-88",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-21",
      "title": "Coin-Operated Devices Remission Order",
      "citation": "SI/99-21",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-166",
      "title": "Coin-operated Devices (Streamlined Accounting Users) Remission Order",
      "citation": "SI/2003-166",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-253",
      "title": "Cold Lake Airport Zoning Regulations",
      "citation": "SOR/91-253",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1416",
      "title": "Collision Regulations",
      "citation": "CRC, c 1416",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-708",
      "title": "Coloured Bean Stabilization Regulations, 1982-83",
      "citation": "SOR/84-708",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-81-124",
      "title": "Proclamation Giving Notice of Coming into Force Canada-Australia Convention Respecting Taxes on Income",
      "citation": "SI/81-124",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-81-27",
      "title": "Proclamation Giving Notice of Coming into Force Canada-Austria Convention Respecting Taxes on Income and Capital",
      "citation": "SI/81-27",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-70",
      "title": "Proclamation Giving Notice of Coming Into Force Canada-Bangladesh Convention Respecting Taxes on Income",
      "citation": "SI/85-70",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-81-48",
      "title": "Proclamation Giving Notice of Coming into Force Canada-Barbados Convention Respecting Taxes on Income and Capital",
      "citation": "SI/81-48",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-45",
      "title": "Proclamation Giving Notice of Coming into Force Canada-Brazil Convention Respecting Taxes on Income",
      "citation": "SI/86-45",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-183",
      "title": "Proclamation Giving Notice of Coming into Force Canada-Cameroon Convention Respecting Taxes on Income",
      "citation": "SI/88-183",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-35",
      "title": "Proclamation Giving Notice of Coming Into Force Canada-Republic of Zambia Convention Respecting Taxes on Income",
      "citation": "SI/91-35",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-26",
      "title": "Proclamation Giving Notice of the Coming into Force of Canada-United Kingdom Convention Respecting Judgments in Civil and Commercial Matters",
      "citation": "SI/87-26",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-87",
      "title": "Proclamation Giving Notice of the Coming into Force on December 31, 2004 of the Convention Between Canada and Romania for the Avoidance of Double Taxation",
      "citation": "SI/2005-87",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-86",
      "title": "Proclamation Giving Notice of the Coming into Force on October 6, 2004 of the Convention Between Canada and Belgium for the Avoidance of Double Taxation",
      "citation": "SI/2005-86",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-356",
      "title": "Commercial Loan (Cooperative Credit Associations) Regulations",
      "citation": "SOR/92-356",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-368",
      "title": "Commercial Loan (Insurance Companies, Societies, Insurance Holding Companies and Foreign Companies) Regulations",
      "citation": "SOR/2001-368",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-368",
      "title": "Commercial Loan (Insurance Companies, Societies, Insurance Holding Companies and Foreign Companies) Regulations",
      "citation": "SOR/2001-368",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-349",
      "title": "Commercial Loan (Trust and Loan Companies) Regulations",
      "citation": "SOR/92-349",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-751",
      "title": "Commercial Samples Remission Order",
      "citation": "CRC, c 751",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-313",
      "title": "Commercial Vehicle Drivers Hours of Service Regulations",
      "citation": "SOR/2005-313",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-716",
      "title": "Commercial Vehicle Drivers Hours of Service Regulations, 1994",
      "citation": "SOR/94-716",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-248",
      "title": "Commissioner\u0027s Standing Orders (Classification Redress Process for Members)",
      "citation": "SOR/2001-248",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-291",
      "title": "Commissioner\u0027s Standing Orders (Conduct)",
      "citation": "SOR/2014-291",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-362",
      "title": "Commissioner\u0027s Standing Orders (Disciplinary Action)",
      "citation": "SOR/88-362",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-141",
      "title": "Commissioner\u0027s Standing Orders (Dispute Resolution Process for Promotions and Job Requirements)",
      "citation": "SOR/2000-141",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-292",
      "title": "Commissioner\u0027s Standing Orders (Employment Requirements)",
      "citation": "SOR/2014-292",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-293",
      "title": "Commissioner\u0027s Standing Orders (General Administration)",
      "citation": "SOR/2014-293",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-181",
      "title": "Commissioner\u0027s Standing Orders (Grievances)",
      "citation": "SOR/2003-181",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-289",
      "title": "Commissioner\u0027s Standing Orders (Grievances and Appeals)",
      "citation": "SOR/2014-289",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-290",
      "title": "Commissioner\u0027s Standing Orders (Investigation and Resolution of Harassment Complaints)",
      "citation": "SOR/2014-290",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-367",
      "title": "Commissioner\u0027s Standing Orders (Practice and Procedure)",
      "citation": "SOR/88-367",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-522",
      "title": "Commissioner\u0027s Standing Orders (Public Complaints)",
      "citation": "SOR/88-522",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-366",
      "title": "Commissioner\u0027s Standing Orders (Qualifications)",
      "citation": "SOR/88-366",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-399",
      "title": "Commissioner\u0027s Standing Orders (Representation), 1997",
      "citation": "SOR/97-399",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-81",
      "title": "Commission for Environmental Cooperation (Director) Remission Order (Part IX of the Excise Tax Act)",
      "citation": "SI/99-81",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-80",
      "title": "Commission for Environmental Cooperation (Executive Director) Remission Order (Part IX of the Excise Tax Act)",
      "citation": "SI/99-80",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-450",
      "title": "Commission for Environmental Cooperation Privileges and Immunities in Canada Order",
      "citation": "SOR/97-450",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-79",
      "title": "Commission for Environmental Cooperation Remission Order (Part IX of the Excise Tax Act)",
      "citation": "SI/99-79",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-13",
      "title": "Order Designating the Commission of Inquiry into the Actions of Canadian Officials in Relation to Maher Arar as a Department and Designating the Prime Minister as Appropriate Minister with Respect to the Commission",
      "citation": "SI/2004-13",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-47",
      "title": "Order Designating the Commission of Inquiry into the Deployment of Canadian Forces to Somalia as a Department and the Prime Minister as Appropriate Minister",
      "citation": "SI/95-47",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-71",
      "title": "Order Designating the Commission of Inquiry into the Investigation of the Bombing of Air India Flight 182 as a Department and the Prime Minister as Appropriate Minister",
      "citation": "SI/2006-71",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-225",
      "title": "Order Designating the Commission of Inquiry on the Blood System in Canada as a Department and the Prime Minister as Appropriate Minister",
      "citation": "SI/93-225",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-60",
      "title": "Order Designating the Commission on the Future of Health Care in Canada as a Department and the Prime Minister as Appropriate Minister",
      "citation": "SI/2001-60",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-36",
      "title": "Commonwealth Caribbean Countries Tariff Rules of Origin Regulations",
      "citation": "SOR/98-36",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-575",
      "title": "Commonwealth of Learning, an Agency, Privileges and Immunities Order, 1988",
      "citation": "SOR/88-575",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1307",
      "title": "Commonwealth Secretariat Privileges and Immunities Order",
      "citation": "CRC, c 1307",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-14",
      "title": "Order Transferring from Communication Canada to the Privy Council Office the Control and Supervision of Certain Portions in Communication Canada",
      "citation": "SI/2004-14",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-16",
      "title": "Order Amalgamating and Combining Communication Canada with the Department of Public Works and Government Services",
      "citation": "SI/2004-16",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1341",
      "title": "Communications Security Establishment Appointments Regulations",
      "citation": "CRC, c 1341",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1342",
      "title": "Communications Security Establishment Exclusion of Positions and Employees Approval Order",
      "citation": "CRC, c 1342",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-255",
      "title": "Communications Security Establishment Regulations",
      "citation": "SOR/2011-255",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-96",
      "title": "Order Transferring to the Communications Security Establishment the Control and Supervision of Certain Portions of the Federal Public Administration in the Department of National Defence known as the Communications Security Establishment and the Communications Security Establishment Internal Services Unit",
      "citation": "SI/2011-96",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-803",
      "title": "Comox Airport Zoning Regulations",
      "citation": "SOR/80-803",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-219",
      "title": "Companies\u0027 Creditors Arrangement Regulations",
      "citation": "SOR/2009-219",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-150",
      "title": "Compensation for Certain Birds Destroyed in British Columbia (Avian Influenza) Regulations",
      "citation": "SOR/2004-150",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-233",
      "title": "Compensation for Destroyed Animals Regulations",
      "citation": "SOR/2000-233",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-53",
      "title": "Competency of Operators of Pleasure Craft Regulations",
      "citation": "SOR/99-53",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-290",
      "title": "Competition Tribunal Rules",
      "citation": "SOR/94-290",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-141",
      "title": "Competition Tribunal Rules",
      "citation": "SOR/2008-141",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-370",
      "title": "Complaint Information (Authorized Foreign Banks) Regulations",
      "citation": "SOR/2001-370",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-371",
      "title": "Complaint Information (Banks) Regulations",
      "citation": "SOR/2001-371",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-373",
      "title": "Complaint Information (Canadian Insurance Companies) Regulations",
      "citation": "SOR/2001-373",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-372",
      "title": "Complaint Information (Foreign Insurance Companies) Regulations",
      "citation": "SOR/2001-372",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-374",
      "title": "Complaint Information (Retail Associations) Regulations",
      "citation": "SOR/2001-374",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-375",
      "title": "Complaint Information (Trust and Loan Companies) Regulations",
      "citation": "SOR/2001-375",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-48",
      "title": "Complaints (Banks, Authorized Foreign Banks and External Complaints Bodies) Regulations",
      "citation": "SOR/2013-48",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-181",
      "title": "Proclamation Prescribing the Composition, Dimensions and Design of a Twenty-Five Cent Base Metal Coin",
      "citation": "SOR/92-181",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-638",
      "title": "Comprehensive Study List Regulations",
      "citation": "SOR/94-638",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-20",
      "title": "Computer Carrier Media Remission Order",
      "citation": "SI/85-20",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-752",
      "title": "Computer Generated Mailing List Remission Order",
      "citation": "CRC, c 752",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-17",
      "title": "Conditions for Exempted Persons Regulations",
      "citation": "SOR/2007-17",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-91",
      "title": "Regulations Establishing Conditions for Making Regulations Under Subsection 36(5.2) of the Fisheries Act",
      "citation": "SOR/2014-91",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-404",
      "title": "Conditions of Carriage Regulations",
      "citation": "SOR/2005-404",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-202",
      "title": "Conditions of Transferring Firearms and Other Weapons Regulations",
      "citation": "SOR/98-202",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-375",
      "title": "Confederation Bridge Area Provincial (P.E.I.) Laws Application Regulations",
      "citation": "SOR/97-375",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-309",
      "title": "Conference of American Armies Privileges and Immunities Order",
      "citation": "SOR/2003-309",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-102",
      "title": "Conference of Defence Ministers of the Americas (CDMA) Privileges and Immunities Order, 2008",
      "citation": "SOR/2008-102",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-118",
      "title": "Proclamation Declaring the Consolidated Agreements on Social Security Between Canada and the United Kingdom in Force December 1, 1995",
      "citation": "SI/95-118",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-967",
      "title": "Consolidated Computer Inc. Adjustment Assistance Regulations",
      "citation": "CRC, c 967",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-968",
      "title": "Consolidated Computer Inc. Enterprise Development Regulations",
      "citation": "CRC, c 968",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-693",
      "title": "Consolidated Computer Inc. Regulations",
      "citation": "SOR/81-693",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-538",
      "title": "Consolidated Regulations Delivery Regulations",
      "citation": "SOR/79-538",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-552",
      "title": "Construction Starts Fee Order",
      "citation": "SOR/81-552",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-30",
      "title": "Consular Fees (Specialized Services) Regulations",
      "citation": "SOR/2003-30",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-538",
      "title": "Consular Services Fees Regulations",
      "citation": "SOR/95-538",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-269",
      "title": "Consumer Chemicals and Containers Regulations, 2001",
      "citation": "SOR/2001-269",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-417",
      "title": "Consumer Packaging and Labelling Regulations",
      "citation": "CRC, c 417",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-273",
      "title": "Consumer Products Containing Lead (Contact with Mouth) Regulations",
      "citation": "SOR/2010-273",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-486",
      "title": "Contaminated Fuel Regulations",
      "citation": "SOR/91-486",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-204",
      "title": "Continuation of Amalgamated or Merged Corporations Regulations",
      "citation": "SOR/2003-204",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-156",
      "title": "Contractor and Homeowner Course Registration Fees Order",
      "citation": "SOR/84-156",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-313",
      "title": "Contraventions Regulations",
      "citation": "SOR/96-313",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-56",
      "title": "Order Transferring the Control and Supervision of Certain Portions of the Federal Public Administration in Passport Canada from the Department of Foreign Affairs and International Trade to the Department of Citizenship and Immigration and to the Department of Human Resources and Skills Development",
      "citation": "SI/2013-56",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-2",
      "title": "Controlled Access Zone Order (Halifax, Esquimalt and Nanoose Harbours)",
      "citation": "SI/2003-2",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-234",
      "title": "Controlled Drugs and Substances Act (Police Enforcement) Regulations",
      "citation": "SOR/97-234",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-32",
      "title": "Controlled Goods Regulations",
      "citation": "SOR/2001-32",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-66",
      "title": "Controlled Products Regulations",
      "citation": "SOR/88-66",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-56",
      "title": "Proclamation Giving Notice that the Protocol Amending the Convention Between the Government of Canada and the Government of the French Republic for the Avoidance of Double Taxation and the Prevention of Fiscal Evasion with respect to Taxes on Income and on Capital signed on May 2, 1975, as amended by the Protocol signed on January 16, 1987 and as further amended by the Protocol signed on November 30, 1995 Came into Force on December 27, 2013",
      "citation": "SI/2014-56",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2010-13",
      "title": "Proclamation Giving Notice That the Convention on Social Security Between the Government of Canada and the Government of the Kingdom of Morocco",
      "citation": "SI/2010-13",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-45",
      "title": "Convention Refugee Determination Division Rules",
      "citation": "SOR/93-45",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-65",
      "title": "Conversion From Analog to Digital Television Regulations",
      "citation": "SOR/2011-65",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-129",
      "title": "Converted Company Ownership Regulations",
      "citation": "SOR/99-129",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-418",
      "title": "Cooperatives Tariff of Fees",
      "citation": "CRC, c 418",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-31",
      "title": "Cooperative Vereniging International Post Corporation U.A. Acquisition of Shares Order",
      "citation": "SOR/89-31",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-181",
      "title": "Regulations Respecting the Coordination by Federal Authorities of Environmental Assessment Procedures and Requirements",
      "citation": "SOR/97-181",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-523",
      "title": "Coppermine Airport Zoning Regulations",
      "citation": "SOR/93-523",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-457",
      "title": "Copyright Regulations",
      "citation": "SOR/97-457",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-68",
      "title": "Coral Harbour Airport Zoning Regulations",
      "citation": "SOR/92-68",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-112",
      "title": "Corded Window Covering Products Regulations",
      "citation": "SOR/2009-112",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-131",
      "title": "Corn 1987 Period Stabilization Regulations",
      "citation": "SOR/89-131",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-476",
      "title": "Corn Stabilization Regulations, 1977-78",
      "citation": "SOR/79-476",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-57",
      "title": "Corporate Interrelationships (Banks and Bank Holding Companies) Regulations",
      "citation": "SOR/2008-57",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-58",
      "title": "Corporate Interrelationships (Cooperative Credit Associations) Regulations",
      "citation": "SOR/2008-58",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-59",
      "title": "Corporate Interrelationships (Insurance Companies and Insurance Holding Companies) Regulations",
      "citation": "SOR/2008-59",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-60",
      "title": "Corporate Interrelationships (Trust and Loan Companies) Regulations",
      "citation": "SOR/2008-60",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-151",
      "title": "Corporation de chauffage urbain de Montréal (CCUM) Incorporation and Sale Authorization Order",
      "citation": "SOR/90-151",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-125",
      "title": "Corporations and Labour Unions Returns Act Regulations",
      "citation": "SOR/84-125",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-13",
      "title": "Corporations Returns Regulations",
      "citation": "SOR/2014-13",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-114",
      "title": "By-law Repealing the Canada Deposit Insurance Corporation Standards of Sound Business and Financial Practices By-law",
      "citation": "SOR/2005-114",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-90",
      "title": "Corrected Certificates Remission Order",
      "citation": "SI/98-90",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-620",
      "title": "Corrections and Conditional Release Regulations",
      "citation": "SOR/92-620",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-869",
      "title": "Cosmetic Regulations",
      "citation": "CRC, c 869",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-262",
      "title": "Cost of Borrowing (Authorized Foreign Banks) Regulations",
      "citation": "SOR/2002-262",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-101",
      "title": "Cost of Borrowing (Banks) Regulations",
      "citation": "SOR/2001-101",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-102",
      "title": "Cost of Borrowing (Canadian Insurance Companies) Regulations",
      "citation": "SOR/2001-102",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-103",
      "title": "Cost of Borrowing (Foreign Insurance Companies) Regulations",
      "citation": "SOR/2001-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-263",
      "title": "Cost of Borrowing (Retail Associations) Regulations",
      "citation": "SOR/2002-263",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-104",
      "title": "Cost of Borrowing (Trust and Loan Companies) Regulations",
      "citation": "SOR/2001-104",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-146",
      "title": "Cost Recovery Regulations",
      "citation": "SOR/2012-146",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-61-472",
      "title": "C.O.T.C. Pension Regulations",
      "citation": "SOR/61-472",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1357",
      "title": "Counting of Service by Former Members of the Senate or House of Commons Regulations",
      "citation": "CRC, c 1357",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1394",
      "title": "Counting of Service by Former Members of the Senate or House of Commons Regulations, No. 2",
      "citation": "CRC, c 1394",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1596",
      "title": "Proclamation Designating Certain Countries as Designated States",
      "citation": "CRC, c 1596",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-84-167",
      "title": "Proclamation Designating Certain Countries as Designated States for Purposes of the Act",
      "citation": "SI/84-167",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-264",
      "title": "Proclamation Designating Certain Countries as Designated States for Purposes of the Act",
      "citation": "SOR/93-264",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-78",
      "title": "Proclamation Designating Certain Countries as Designated States for Purposes of the Act",
      "citation": "SOR/96-78",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-86",
      "title": "Proclamation Designating Certain Countries as Designated States for Purposes of the Act and Designating the Civilian Personnel of the Designated States",
      "citation": "SOR/99-86",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-23",
      "title": "Determination of Country of Origin for the Purposes of Marking Goods (NAFTA Countries) Regulations",
      "citation": "SOR/94-23",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-182",
      "title": "Courier Imports Remission Order",
      "citation": "SI/85-182",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-959",
      "title": "Court Martial Appeal Court Rules",
      "citation": "SOR/86-959",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-9",
      "title": "The Court of Appeal Criminal Appeal Rules (Saskatchewan)",
      "citation": "SI/2011-9",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-39",
      "title": "Court of Queen\u0027s Bench for Alberta Summary Conviction Appeal Rules",
      "citation": "SI/2012-39",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-81-97",
      "title": "The Court of Queen\u0027s Bench for Saskatchewan Summary Conviction Appeal Rules",
      "citation": "SI/81-97",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-20",
      "title": "Court of Queen\u0027s Bench for Saskatchewan Summary Conviction Appeal Rules",
      "citation": "SI/2011-20",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-78",
      "title": "Rule of the Court of Queen\u0027s Bench of New Brunswick Respecting Pre-Trial Conferences Under Subsection 553.1(2) of the Criminal Code of Canada",
      "citation": "SI/86-78",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-51",
      "title": "CPAFTA Rules of Origin for Casual Goods Regulations",
      "citation": "SOR/2013-51",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-50",
      "title": "CPAFTA Rules of Origin Regulations",
      "citation": "SOR/2013-50",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-52",
      "title": "CPAFTA Tariff Preference Regulations",
      "citation": "SOR/2013-52",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-215",
      "title": "CPFTA Rules of Origin for Casual Goods Regulations",
      "citation": "SOR/2009-215",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-214",
      "title": "CPFTA Rules of Origin Regulations",
      "citation": "SOR/2009-214",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-217",
      "title": "CPFTA Sugar Aggregate Quantity Limit Order",
      "citation": "SOR/2009-217",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-216",
      "title": "CPFTA Tariff Preference Regulations",
      "citation": "SOR/2009-216",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-216",
      "title": "CPFTA Verification of Origin Regulations",
      "citation": "SOR/2013-216",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-126",
      "title": "Cranbrook Airport Zoning Regulations",
      "citation": "SOR/88-126",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-257",
      "title": "Credit Business Practices (Banks, Authorized Foreign Banks, Trust and Loan Companies, Retail Associations, Canadian Insurance Companies and Foreign Insurance Companies) Regulations",
      "citation": "SOR/2009-257",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-282",
      "title": "Credit Enhancement Fund Use Regulations",
      "citation": "SOR/2010-282",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-57",
      "title": "Credit for Provincial Relief (HST) Regulations",
      "citation": "SOR/2011-57",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-11",
      "title": "Credit Information (Insurance Companies) Regulations",
      "citation": "SOR/97-11",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-44",
      "title": "Credit Note and Debit Note Information (GST/HST) Regulations",
      "citation": "SOR/91-44",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-315",
      "title": "Cree-Naskapi Band Elections Regulations",
      "citation": "SOR/88-315",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1070",
      "title": "Cree-Naskapi Land Registry Regulations",
      "citation": "SOR/86-1070",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-490",
      "title": "Cree-Naskapi Long-Term Borrowing Regulations, 1986",
      "citation": "SOR/86-490",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1418",
      "title": "Crew Accommodation Regulations",
      "citation": "CRC, c 1418",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-390",
      "title": "Crewing Regulations",
      "citation": "SOR/97-390",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-962",
      "title": "Cribs and Cradles Regulations",
      "citation": "SOR/86-962",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-261",
      "title": "Cribs, Cradles and Bassinets Regulations",
      "citation": "SOR/2010-261",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-169",
      "title": "Criminal Appeal Rules",
      "citation": "SI/93-169",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-78",
      "title": "Criminal Procedure Rules of the Supreme Court of the Northwest Territories",
      "citation": "SI/98-78",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-7",
      "title": "Criminal Proceedings Rules for the Superior Court of Justice (Ontario)",
      "citation": "SI/2012-7",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-303",
      "title": "Criminal Records Regulations",
      "citation": "SOR/2000-303",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-30",
      "title": "Criminal Rules of the Ontario Court of Justice",
      "citation": "SI/2012-30",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-140",
      "title": "Criminal Rules of the Supreme Court of British Columbia",
      "citation": "SI/97-140",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-68",
      "title": "Critical Habitats of the Northeast Pacific Northern and Southern Resident Populations of the Killer Whale (Orcinus Orca) Order",
      "citation": "SOR/2009-68",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-412",
      "title": "Cross-border Currency and Monetary Instruments Reporting Regulations",
      "citation": "SOR/2002-412",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-223",
      "title": "Crown Corporation Corporate Plan, Budget and Summaries Regulations",
      "citation": "SOR/95-223",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-226",
      "title": "Crown Corporation General Regulations, 1995",
      "citation": "SOR/95-226",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-1030",
      "title": "Crown Corporation Payments Regulations",
      "citation": "SOR/81-1030",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-104",
      "title": "Crown Corporations Involved in the Provision of Commercial Loans Environmental Assessment Regulations",
      "citation": "SOR/2006-104",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-604",
      "title": "Crown Liability and Proceedings (Provincial Court) Regulations",
      "citation": "SOR/91-604",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-113",
      "title": "Crown Share Adjustment Payments Regulations",
      "citation": "SOR/2012-113",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-11",
      "title": "Crown Waiver Order (Northwest Territories)",
      "citation": "SI/90-11",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-212",
      "title": "Order Referring Back to the CRTC a Decision Amending a Broadcasting Licence of Westcom Radio Group Limited",
      "citation": "SI/86-212",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-165",
      "title": "Order Referring Back to the CRTC a Decision Amending the Broadcasting Licence of Allarcom Limited",
      "citation": "SI/83-165",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-161",
      "title": "Order Referring Back to the CRTC a Decision Amending the Broadcasting Licence of Télécâble Vidéotron Ltée",
      "citation": "SI/83-161",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-1",
      "title": "Order Referring Back to the CRTC a Decision Issuing a Broadcasting Licence to Omineca Cablevision Limited",
      "citation": "SI/85-1",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-3",
      "title": "Order Referring back to the CRTC a Decision Respecting CHEZ-FM INC.",
      "citation": "SI/90-3",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-91",
      "title": "Order Referring back to the CRTC a Decision Respecting Nation\u0027s Capital Television Incorporated",
      "citation": "SI/87-91",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-98",
      "title": "Order Referring Back to the CRTC a Decision Respecting Telemedia Communications Inc. and Muskoka-Parry Sound Broadcasting Limited",
      "citation": "SI/87-98",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-257",
      "title": "Order Declining to Set Aside to the CRTC a Decision Respecting Telemedia Communications Inc. and Muskoka-Parry Sound Broadcasting Limited",
      "citation": "SI/87-257",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-34",
      "title": "Order Referring Back to the CRTC a Decision Respecting the Approval of an Application for a Licence by Ian Mackay",
      "citation": "SI/87-34",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-160",
      "title": "Order Referring Back to the CRTC Certain Decisions Amending Broadcasting Licences of Câblevision Nationale Ltée",
      "citation": "SI/83-160",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-54",
      "title": "Order Referring Back to the CRTC Decision CRTC 2001-757",
      "citation": "SI/2002-54",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-320",
      "title": "Directions to the CRTC (Direct-to-Home (DTH) Pay-Per-View Television Programming Undertakings) Order",
      "citation": "SOR/95-320",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-319",
      "title": "Directions to the CRTC (Direct-to-Home (DTH) Satellite Distribution Undertakings) Order",
      "citation": "SOR/95-319",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-192",
      "title": "Direction to the CRTC (Ineligibility of Non-Canadians)",
      "citation": "SOR/97-192",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-627",
      "title": "Direction to the CRTC (Ineligibility to Hold Broadcasting Licences)",
      "citation": "SOR/85-627",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-378",
      "title": "Direction to the CRTC (Reservation of Cable Channels)",
      "citation": "CRC, c 378",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-60",
      "title": "Direction to the CRTC (Reservation of Channels for the Distribution of CPAC)",
      "citation": "SOR/2005-60",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-284",
      "title": "Direction to the CRTC (Reservation of Frequencies for Toronto) Order",
      "citation": "SOR/98-284",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-375",
      "title": "CRTC Rules of Procedure",
      "citation": "CRC, c 375",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-555",
      "title": "CRTC Tariff Regulations",
      "citation": "SOR/79-555",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-554",
      "title": "CRTC Telecommunications Rules of Procedure",
      "citation": "SOR/79-554",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-190",
      "title": "Order Referring Back to the CRTC the Decisions Respecting Shaw Cablesystems Ltd. and Cogeco Radio-Télévision Inc.",
      "citation": "SI/93-190",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-503",
      "title": "Order Directing the CRTC to Issue a Notice to all Television Networks throughout Canada to Broadcast a Special Message from the Prime Minister",
      "citation": "SOR/95-503",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-78-57",
      "title": "Cruiser Remission Order",
      "citation": "SI/78-57",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-449",
      "title": "Cultural Property Export Regulations",
      "citation": "CRC, c 449",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-900",
      "title": "Currency Exchange for Customs Valuation Regulations",
      "citation": "SOR/85-900",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-754",
      "title": "Customs Accounting Document Error Remission Order",
      "citation": "CRC, c 754",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-196",
      "title": "Customs and Excise Human Rights Investigation Regulations",
      "citation": "SOR/83-196",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-46",
      "title": "Customs Bonded Warehouses Regulations",
      "citation": "SOR/96-46",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1067",
      "title": "Customs Brokers Licensing Regulations",
      "citation": "SOR/86-1067",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-127",
      "title": "Customs Controlled Areas Regulations",
      "citation": "SOR/2013-127",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-522",
      "title": "Customs Diplomatic Privileges Regulations",
      "citation": "CRC, c 522",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-337",
      "title": "Customs Drawback on Toy Skins Regulations",
      "citation": "SOR/80-337",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-487",
      "title": "Customs Drawback Shirting Fabrics Regulations",
      "citation": "CRC, c 487",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-301",
      "title": "Customs Duties Accelerated Reduction Order, No. 1",
      "citation": "SOR/90-301",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-379",
      "title": "Customs Duties Accelerated Reduction Order, No. 2",
      "citation": "SOR/90-379",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-399",
      "title": "Customs Duties Accelerated Reduction Order, No. 3",
      "citation": "SOR/90-399",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-406",
      "title": "Customs Duties Accelerated Reduction Order, No. 4",
      "citation": "SOR/90-406",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-548",
      "title": "Customs Duties Accelerated Reduction Order, No.5",
      "citation": "SOR/91-548",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-514",
      "title": "Customs Duties Accelerated Reduction Order, No. 6",
      "citation": "SOR/92-514",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-32",
      "title": "Regulations Respecting the Customs Duty Payable on Woollen Fabrics Originating in Commonwealth Countries",
      "citation": "SOR/98-32",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1065",
      "title": "Customs Sufferance Warehouses Regulations",
      "citation": "SOR/86-1065",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-79-72",
      "title": "Customs Tariff Remission Order, 1979",
      "citation": "SI/79-72",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-164",
      "title": "CyberFluor Inc. Shares Acquisition Order",
      "citation": "SOR/90-164",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-466",
      "title": "Dairy Products Marketing Regulations",
      "citation": "SOR/94-466",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-840",
      "title": "Dairy Products Regulations",
      "citation": "SOR/79-840",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-316",
      "title": "Dakota Tipi Band Council Elections Order",
      "citation": "SOR/2002-316",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-303",
      "title": "Dakota Tipi Band Council Method of Election Regulations",
      "citation": "SOR/2002-303",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-24",
      "title": "Dangerous Bulk Materials Regulations",
      "citation": "SOR/87-24",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-4",
      "title": "Dangerous Chemicals and Noxious Liquid Substances Regulations",
      "citation": "SOR/93-4",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-951",
      "title": "Dangerous Goods Shipping Regulations",
      "citation": "SOR/81-951",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-604",
      "title": "Dauphin Airport Zoning Regulations",
      "citation": "SOR/87-604",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-375",
      "title": "Dawson Creek Airport Zoning Regulations",
      "citation": "SOR/87-375",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-228",
      "title": "Order Declaring December 12, 1988 as the Day on which the First Supplement to the Revised Statutes of Canada, 1985 Comes into Force",
      "citation": "SI/88-228",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-231",
      "title": "Order Declaring November 1, 1989 as the Day on Which the Fourth Supplement to the Revised Statutes of Canada, 1985 Comes into Force",
      "citation": "SI/89-231",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-239",
      "title": "Order Declaring  December 12, 1988 as the Day on Which the Second Supplement to the Revised Statutes of Canada, 1985 Comes Into Force",
      "citation": "SI/88-239",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-123",
      "title": "Order Declaring May 1, 1989 as the Day on Which the Third Supplement to the Revised Statutes of Canada, 1985 Comes Into Force",
      "citation": "SI/89-123",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-244",
      "title": "Debt Reserve Fund Replenishment Regulations",
      "citation": "SOR/2006-244",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-602",
      "title": "Debt Write-off Regulations, 1994",
      "citation": "SOR/94-602",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1599",
      "title": "Deceased or Former Members Dependants Payment Order",
      "citation": "CRC, c 1599",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-380",
      "title": "Decision Body Time Periods and Consultation Regulations",
      "citation": "SOR/2005-380",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-496",
      "title": "Order Varying Decision No. 712-W-1993 of the National Transportation Agency",
      "citation": "SOR/94-496",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-110",
      "title": "Order Giving Notice of Decisions not to add Certain Species to the List of Endangered Species",
      "citation": "SI/2006-110",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-72",
      "title": "Order Giving Notice of Decisions not to add Certain Species to the List of Endangered Species",
      "citation": "SI/2005-72",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-2",
      "title": "Order Giving Notice of Decisions not to add Certain Species to the List of Endangered Species",
      "citation": "SI/2005-2",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2007-115",
      "title": "Order Giving Notice of Decisions not to add Certain Species to the List of Endangered Species",
      "citation": "SI/2007-115",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-61",
      "title": "Order Giving Notice of Decisions not to add Certain Species to the List of Endangered Species",
      "citation": "SI/2006-61",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-56",
      "title": "Order Giving Notice of Decisions not to add Certain Species to the List of Endangered Species",
      "citation": "SI/2011-56",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-56",
      "title": "Declarations Regulations (Chemical Weapons Convention)",
      "citation": "SOR/2010-56",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-55",
      "title": "Canadian Orders Decorations and Medals Directive, 1998",
      "citation": "SI/98-55",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-65",
      "title": "Deduction for Provincial Rebate (GST/HST) Regulations",
      "citation": "SOR/2001-65",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-299",
      "title": "Deep Panuke Offshore Production Platform Remission Order, 2007",
      "citation": "SOR/2007-299",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1135",
      "title": "Deer Lake Airport Zoning Regulations",
      "citation": "SOR/86-1135",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-686",
      "title": "Defence Clothing and Equipment Loan Order",
      "citation": "CRC, c 686",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-957",
      "title": "Defence Controlled Access Area Regulations",
      "citation": "SOR/86-957",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-687",
      "title": "Defence Floating Equipment Rental Order",
      "citation": "CRC, c 687",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-688",
      "title": "Defence Maps and Charts Transfer Order",
      "citation": "CRC, c 688",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-689",
      "title": "Defence Materiel and Equipment Loaned to Contractors Order",
      "citation": "CRC, c 689",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-691",
      "title": "Defence Materiel Loan or Transfer for Test and Evaluation Regulations",
      "citation": "CRC, c 691",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-690",
      "title": "Defence Materiel Loan Regulations",
      "citation": "CRC, c 690",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-755",
      "title": "Defence Production and Development Sharing Remission Order",
      "citation": "CRC, c 755",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1335",
      "title": "Defence Research Board Employees Regulations",
      "citation": "CRC, c 1335",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-554",
      "title": "Defence Services Pension Continuation Regulations",
      "citation": "CRC, c 554",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-55-416",
      "title": "Defence Services Pension Part V Regulations",
      "citation": "SOR/55-416",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-692",
      "title": "Defence Survival Training Indemnity Order",
      "citation": "CRC, c 692",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-567",
      "title": "Deficient Postage Regulations",
      "citation": "SOR/85-567",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-376",
      "title": "Definition of Promotion Regulations",
      "citation": "SOR/2005-376",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-257",
      "title": "Definition of Settler for the Purpose of Tariff Item No. 9807.00.00 Regulations",
      "citation": "SOR/2005-257",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-755",
      "title": "Definition of “Small Cable Transmission System” Regulations",
      "citation": "SOR/94-755",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-307",
      "title": "Definition of “Wireless Transmission System” Regulations",
      "citation": "SOR/98-307",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-766",
      "title": "Delegation of Powers (Customs) Order",
      "citation": "SOR/86-766",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-359",
      "title": "Delegation of Powers (Customs) Regulations",
      "citation": "SOR/86-359",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1588",
      "title": "Delegation of Powers (VLA) Regulations",
      "citation": "CRC, c 1588",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-756",
      "title": "Demonstration Aircraft Remission Order",
      "citation": "CRC, c 756",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-22",
      "title": "Denatured and Specially Denatured Alcohol Regulations",
      "citation": "SOR/2005-22",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-180",
      "title": "Denial of Licences for Family Orders and Agreements Enforcement Regulations",
      "citation": "SOR/97-180",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-333",
      "title": "Denim Apparel Fabrics Remission Order",
      "citation": "SOR/88-333",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-85",
      "title": "Order Transferring to the Department of Agriculture and Agri-Food the Control and Supervision of the Portion of the Public Service known as the Retail Food Program of the Department of Industry",
      "citation": "SI/95-85",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-140",
      "title": "Order Transferring to the Department of Agriculture the Control and Supervision of those Portions of the Public Service in the Department of Industry, Science and Technology Relating to Agri-Food Processing and Manufacturing and in the Department of Consumer and Corporate Affairs Relating to Agri-Food and Labelling",
      "citation": "SI/93-140",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-231",
      "title": "Order Transferring From the Department of Canadian Heritage to the Department of Human Resources Development the Control and Supervision of the Voluntary Sector Affairs Directorate",
      "citation": "SI/2003-231",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-112",
      "title": "Order Transferring From the Department of Canadian Heritage to the Parks Canada Agency the Control and Supervision of the Historic Places Policy Group and from the Minister of Canadian Heritage to the Minister of the Environment the Powers, Duties and Functions of the Department of Canadian Heritage Act Relating to Certain Programs",
      "citation": "SI/2004-112",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-135",
      "title": "Order Transferring to the Department of Citizenship and Immigration the Control and Supervision of Certain Portions within the Canada Border Services Agency and Transferring from the Deputy Minister and Minister of Public Safety and Emergency Preparedness to the Minister of Citizenship and Immigration Certain Powers, Duties and Functions",
      "citation": "SI/2004-135",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-202",
      "title": "Order Transferring to the Department of Communications the Control and Supervision of Certain Portions of the Public Service and to the Minister of Communications Certain Powers, Duties and Functions of the Secretary of State of Canada and of the Minister of Multiculturalism and Citizenship and Amalgamating and Combining the Departments of Communications and Multiculturalism and Citizenship Under the Minister of Communications",
      "citation": "SI/93-202",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-142",
      "title": "Order Transferring to the Department of Employment and Immigration the Control and Supervision of Certain Portions of the Public Service and Transferring to the Minister of Employment and Immigration the Powers, Duties and Functions of the Minister of Labour and Amalgamating and Combining the Department of Employment and Immigration and the Department of Labour Under the Minister of Employment and Immigration",
      "citation": "SI/93-142",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-35",
      "title": "Order Transferring from the Department of Employment and Immigration to the Department of National Revenue the Control and Supervision of those Portions of the Public Service in the Income Security Programs Branch relating to Child Tax Benefit Program and Special Allowances under the Children\u0027s Special Allowances Act",
      "citation": "SI/95-35",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-45",
      "title": "Order Transferring From the Department of Employment and Immigration to the Office of the Coordinator, Status of Women, the Control and Supervision of the Women\u0027s Program",
      "citation": "SI/95-45",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-311",
      "title": "Department of Employment and Social Development Regulations",
      "citation": "SOR/2005-311",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-311",
      "title": "Department of Employment and Social Development Regulations",
      "citation": "SOR/2005-311",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-203",
      "title": "Order Transferring to the Department of External Affairs from the Department of the Solicitor General, the Control and Supervision of the Passport Office",
      "citation": "SI/93-203",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-172b",
      "title": "Department of External Affairs Terms Under Three Months Regulations, 1993",
      "citation": "SOR/93-172b",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-146",
      "title": "Order Transferring to the Department of External Affairs the Control and Supervision of the Investment Development Division in Investment Canada",
      "citation": "SI/93-146",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-233",
      "title": "Order Transferring From the Department of Fisheries and Oceans to the Department of Transport of Certain Portions of the Directorate-General of Maritime Programs",
      "citation": "SI/2003-233",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-348a",
      "title": "Department of Forestry Terms Under Six Months Exclusion Approval Order, 1989",
      "citation": "SOR/89-348a",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-348b",
      "title": "Department of Forestry Terms Under Six Months Regulations, 1989",
      "citation": "SOR/89-348b",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-95b",
      "title": "Department of Forestry Terms Under Three Months Regulations, 1993",
      "citation": "SOR/93-95b",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-60",
      "title": "Order Transferring to the Department of Health and to the Public Health Agency of Canada the Control and Supervision of Certain Portions of the Federal Public Administration",
      "citation": "SI/2012-60",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-123",
      "title": "Order Transferring From the Department of Health to the Public Health Agency of Canada the Control and Supervision of the Population and Public Health Branch and Ordering the Minister of Health to Preside Over the Agency",
      "citation": "SI/2004-123",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-12",
      "title": "Order transferring to the Department of Indian Affairs and Northern Development the control and supervision of the portion of the federal public administration in the Department of Canadian Heritage known as the Urban Aboriginal Youth and Community Programs Unit",
      "citation": "SI/2012-12",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-141",
      "title": "Order Transferring to the Department of Industry, Science and Technology the Control and Supervision of Certain Portions of the Public Service and Transferring to the Minister of Industry, Science and Technology the Powers, Duties and Functions of the Minister of Consumer and Corporate Affairs and Amalgamating and Combining the Departments of Industry, Science and Technology and the Department of Consumer and Corporate Affairs under the Minister of Industry, Science and Technology",
      "citation": "SI/93-141",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-170",
      "title": "Order Transferring to the Department of Industry, Science and Technology the Control and Supervision of Certain Portions of the Public Service in the Department of Communications",
      "citation": "SI/93-170",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-135",
      "title": "Order Transferring From the Department of Industry to the Department of Indian Affairs and Northern Development the Control and Supervision of Aboriginal Business Canada and First Nations SchoolNet and Transferring from the Minister of Industry to the Minister of Indian Affairs and Northern Development the Powers, Duties and Functions Relating to the National Aboriginal Economic Development Board",
      "citation": "SI/2006-135",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-11",
      "title": "Order Amalgamating and Combining the Department of International Trade and the Department of Foreign Affairs and International Trade Under the Minister and the Deputy Minister of Foreign Affairs",
      "citation": "SI/2006-11",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-230",
      "title": "Order Transferring From the Department of Justice to the Department of the Solicitor General the Control and Supervision of the National Crime Prevention Centre",
      "citation": "SI/2003-230",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-201b",
      "title": "Department of National Defence Terms Under Three Months Regulations, 1992",
      "citation": "SOR/92-201b",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-229",
      "title": "Order Transferring From the Department of National Defence to the Department of the Solicitor General the Control and Supervision of the Office of Critical Infrastructure Protection and Emergency Preparedness",
      "citation": "SI/2003-229",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-167",
      "title": "Order Transferring to the Department of National Health and Welfare the Control and Supervision of Certain Portions of the Public Service in the Social Service Programs Branch of the Department of Employment and Immigration",
      "citation": "SI/93-167",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-145",
      "title": "Order Transferring to the Department of National Health and Welfare the Control and Supervision of the Product Safety Branch in the Department of Consumer and Corporate Affairs",
      "citation": "SI/93-145",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-31",
      "title": "Order Transferring to the Department of National Revenue the Control and Supervision of that Portion of the Public Service in the Office of the Superintendent of Financial Institutions known as the Pension Advice Section",
      "citation": "SI/96-31",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-48",
      "title": "Order Transferring From the Department of Natural Resources to the Department of the Environment the Control and Supervision of the Large Final Emitters Group",
      "citation": "SI/2005-48",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-83",
      "title": "Order Transferring From the Department of Public Works and Government Services to the Department of Human Resources and Skills Development the Control and Supervision of the Public Access Programs Sector",
      "citation": "SI/2005-83",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-139b",
      "title": "Department of Public Works Terms Under Three Months Regulations, 1993",
      "citation": "SOR/93-139b",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-309",
      "title": "Department of Social Development Regulations",
      "citation": "SOR/2005-309",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-138",
      "title": "Order Transferring to the Department of Supply and Services the Control and Supervision of the Government Telecommunications Agency and the Translation Bureau and Transferring to the Minister of Supply and Services the Powers, Duties and Functions of the Minister of Public Works and Amalgamating and Combining the Department of Supply and Services and the Department of Public Works Under the Minister of Supply and Services",
      "citation": "SI/93-138",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-36",
      "title": "Order Transferring from the Department of the Environment to the Office of Infrastructure of Canada the Responsibility of the Canada Mortgage and Housing Corporation, Canada Lands Company Limited and Queens Quay West Land Corporation",
      "citation": "SI/2004-36",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-35",
      "title": "Order Transferring to the Department of Transport Certain Portions of the Public Service in the Department of Fisheries and Oceans and of the Powers, Duties and Functions of the Minister of Fisheries and Oceans under certain paragraphs of the Canada Shipping Act",
      "citation": "SI/2004-35",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-7",
      "title": "Order transfering from the Department of Transport to Canada Customs and Revenue Agency the Control and Supervision of the Royal Canadian Mint and Canada Post Corporation",
      "citation": "SI/2004-7",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-46",
      "title": "Order Transferring from the Department of Transport to the Department of Fisheries and Oceans the Control and Supervision of the Canadian Coast Guard other Than the Harbours and Ports Directorate and the Regional Harbours and Ports Branches, the Marine Regulatory Directorate, and the Ship Inspection Directorate and the Regional Ship Inspection Branches",
      "citation": "SI/95-46",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-8",
      "title": "Order Transferring From the Department of Transport to the Department of the Environment the Control and Supervision of the Canada Mortgage and Housing Corporation, Canada Lands Company Limited and Queens Quay West Land Corporation",
      "citation": "SI/2004-8",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-62",
      "title": "Regulations Exempting Departments and Parent Crown Corporations from the Requirements of Subsections 65.1(1) and 131.1(1) of the Financial Administration Act",
      "citation": "SOR/2011-62",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-62",
      "title": "Regulations Exempting a Departments and Parent Crown Corporations from the Requirements of Subsections 65.1(1) and 131.1(1) of the Financial Administration Act",
      "citation": "SOR/2011-62",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-277",
      "title": "Deposit Insurance Application Fee By-law",
      "citation": "SOR/2000-277",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-91",
      "title": "Deposit Out of the Normal Course of Events Notification Regulations",
      "citation": "SOR/2011-91",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-279",
      "title": "Notices of Deposit Restrictions (Authorized Foreign Banks) Regulations",
      "citation": "SOR/99-279",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-98",
      "title": "Deposit Type Instruments Regulations",
      "citation": "SOR/2011-98",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-430",
      "title": "Designated Areas Firearms Order",
      "citation": "CRC, c 430",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-306",
      "title": "Designated Employees Pension Protection Regulations",
      "citation": "SOR/93-306",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-336",
      "title": "Designated Provisions (Customs) Regulations",
      "citation": "SOR/2002-336",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-112",
      "title": "Designated Provisions Regulations",
      "citation": "SOR/2000-112",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-278",
      "title": "Order Designating British Columbia for the Purposes of the Criminal Interest Rate Provisions of the Criminal Code",
      "citation": "SOR/2009-278",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-40",
      "title": "Order Designating Certain Countries and Military Service Agencies for the Purposes of Tariff Item No. 9810.00.00",
      "citation": "SOR/98-40",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-234",
      "title": "Regulations Designating Certain Countries for the Purposes of Tariff Item No. 9810.00.00, 1999-1",
      "citation": "SOR/99-234",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-212",
      "title": "Order Designating Manitoba for the Purposes of the Criminal Interest Rate Provisions of the Criminal Code",
      "citation": "SOR/2008-212",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-177",
      "title": "Order Designating Nova Scotia for the Purposes of the Criminal Interest Rate Provisions of the Criminal Code",
      "citation": "SOR/2009-177",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-277",
      "title": "Order Designating Designating Ontario for thePurposes of the Criminal Interest Rate Provisions of the Criminal Code",
      "citation": "SOR/2009-277",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-134",
      "title": "Regulations Designating Regulatory Provisions for Purposes of Enforcement (Canadian Environmental Protection Act, 1999)",
      "citation": "SOR/2012-134",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-49",
      "title": "Order Designating the Greater Toronto Airports Authority as a Designated Airport Authority and Designating the Date for the Transfer of the Toronto-Lester B. Pearson International Airport to the Greater Toronto Airport Authority",
      "citation": "SI/96-49",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-108",
      "title": "Order Designating the Minister for International Trade as Minister for Purposes of the Export Development Act and as Appropriate Minister with Respect to Export Development Canada for Purposes of the Financial Administration Act",
      "citation": "SI/93-108",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-128",
      "title": "Order Designating the Minister of Canadian Heritage as Minister for Purposes of the Act",
      "citation": "SI/99-128",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-227",
      "title": "Order Designating the Minister of Communications as Appropriate Minister with Respect to the National Battlefields Commission",
      "citation": "SI/93-227",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-173",
      "title": "Order Designating the Minister of Finance as Minister for Purposes of the Act",
      "citation": "SI/93-173",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-174",
      "title": "Order Designating the Minister of Finance as Minister for Purposes of the Act",
      "citation": "SI/93-174",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-58",
      "title": "Order Designating the Minister of Foreign Affairs, the Minister of National Defence and the Minister of Public Works and Government Services as the Ministers for Purposes of Certain Sections of the Act",
      "citation": "SI/99-58",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-128",
      "title": "Order Designating the Minister of Health as Appropriate Minister with Respect to the Assisted Human Reproduction Agency of Canada for Purposes of the Act",
      "citation": "SI/2006-128",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-229",
      "title": "Order Designating the Minister of Industry, Science and Technology as Minister for Purposes of the Act",
      "citation": "SI/93-229",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-33",
      "title": "Order Designating the Minister of Industry, Science and Technology as Minister for Purposes of the Act",
      "citation": "SI/90-33",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-6",
      "title": "Order Designating the Minister of Industry, Science and Technology as Minister for Purposes of the Canadian Space Agency Act, the Canadian Space Agency as a Department, and the President of the Agency as Deputy Head",
      "citation": "SI/91-6",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-235",
      "title": "Order Designating the Minister of Industry, Science and Technology as Minister for Purposes of the National Research Council Act and as Appropriate Minister with Respect to the National Research Council of Canada for Purposes of the Financial Administration Act",
      "citation": "SI/93-235",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-117",
      "title": "Order Designating the Minister of Industry, Science and Technology as Minister for Purposes of the Standards Council of Canada Act",
      "citation": "SI/93-117",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-39",
      "title": "Order Designating the Minister of Industry, Science and Technology as Minister for Purposes of the Statistics Act and for Purposes of the Financial Administration Act with Respect to Statistics Canada",
      "citation": "SI/90-39",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-109",
      "title": "Designating the Minister of Justice and the President of the Treasury Board as Ministers for Purposes of Certain Sections of the Act",
      "citation": "SI/83-109",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-113",
      "title": "Order Designating the Minister of Supply and Services as Minister for Purposes of the Act",
      "citation": "SI/93-113",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-23",
      "title": "Order Designating the Minister of Transport as Minister for Purposes of the Act",
      "citation": "SI/2006-23",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-28",
      "title": "Order Designating the Minister of Transport as Minister for Purposes of the Act",
      "citation": "SI/2006-28",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-85",
      "title": "Order Designating the Ottawa Macdonald-Cartier International Airport Authority as a Designated Airport Authority and Designating the Date for the Transfer of the Ottawa Macdonald-Cartier International Airport to the Ottawa Macdonald-Cartier International Airport Authority",
      "citation": "SI/96-85",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-84",
      "title": "Order Designating the Winnipeg Airports Authority Inc. as a Designated Airport Authority and Designating the Date for the Transfer of the Winnipeg International Airport to the Winnipeg Airports Authority Inc.",
      "citation": "SI/96-84",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-255",
      "title": "Designation of Countries (Standards Council of Canada) Order",
      "citation": "SI/93-255",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-82-221",
      "title": "Designation of Place of Inspection and Sale",
      "citation": "SI/82-221",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-4",
      "title": "Designer Remission Order, 2001",
      "citation": "SOR/2002-4",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-238",
      "title": "Destruction of Paid Instruments Regulations, 1996",
      "citation": "SOR/97-238",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-16",
      "title": "Determination of Country of Origin for the Purpose of Marking Goods (Non-NAFTA Countries) Regulations",
      "citation": "SOR/94-16",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-951",
      "title": "Determination of the Tariff Classification of Sugar, Molasses and Sugar Syrup Regulations",
      "citation": "SOR/86-951",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-44",
      "title": "Determination, Re-determination and Further Re-determination of Origin, Tariff Classification and Value for Duty Regulations",
      "citation": "SOR/98-44",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-957",
      "title": "Determining that Telesat Canada Need Not Apply for a Certificate of Continuance Under the Act",
      "citation": "SOR/80-957",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-322",
      "title": "Development Tax and Redevelopment Tax Grant Regulations",
      "citation": "CRC, c 322",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-61",
      "title": "DHC Shares Sale Regulations",
      "citation": "SOR/86-61",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-79-19",
      "title": "Diplomatic, Consular and International Organizations\u0027 Property Grants Order",
      "citation": "SI/79-19",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-757",
      "title": "Diplomatic (Excise Taxes) Remission Order",
      "citation": "CRC, c 757",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-78-162",
      "title": "Diplomatic Motor Vehicle Diversion Remission Order",
      "citation": "SI/78-162",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-485",
      "title": "Diplomatic Service (Special) Superannuation Regulations, 1988",
      "citation": "SOR/88-485",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-758",
      "title": "Dipropylacetic Acid (Also Known as 2-Propylpentanoic Acid) Remission Order",
      "citation": "CRC, c 758",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-109",
      "title": "Direction Applying the Auditor General Act Sustainable Development Strategy Requirements to Certain Departments",
      "citation": "SI/2004-109",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-165",
      "title": "Direction Applying the Auditor General Act Sustainable Development Strategy Requirements to Certain Departments",
      "citation": "SOR/2007-165",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-355",
      "title": "Order Issuing a Direction to the CRTC on Implementing the Canadian Telecommunications Policy Objectives",
      "citation": "SOR/2006-355",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-876",
      "title": "Direct Shipment of Goods Regulations",
      "citation": "SOR/86-876",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-278",
      "title": "Disclosure of Charges (Authorized Foreign Banks) Regulations",
      "citation": "SOR/99-278",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-324",
      "title": "Disclosure of Charges (Banks) Regulations",
      "citation": "SOR/92-324",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-297",
      "title": "Disclosure of Charges (Retail Associations) Regulations",
      "citation": "SOR/2003-297",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-328",
      "title": "Disclosure of Charges (Trust and Loan Companies) Regulations",
      "citation": "SOR/92-328",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-272",
      "title": "Disclosure of Interest (Authorized Foreign Banks) Regulations",
      "citation": "SOR/99-272",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-321",
      "title": "Disclosure of Interest (Banks) Regulations",
      "citation": "SOR/92-321",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-298",
      "title": "Disclosure of Interest (Retail Associations) Regulations",
      "citation": "SOR/2003-298",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-322",
      "title": "Disclosure of Interest (Trust and Loan Companies) Regulations",
      "citation": "SOR/92-322",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-38",
      "title": "Disclosure of Tax (GST/HST) Regulations",
      "citation": "SOR/91-38",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-471",
      "title": "Disclosure on Account Opening by Telephone Request (Authorized Foreign Banks) Regulations",
      "citation": "SOR/2001-471",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-472",
      "title": "Disclosure on Account Opening by Telephone Request (Banks) Regulations",
      "citation": "SOR/2001-472",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-299",
      "title": "Disclosure on Account Opening by Telephone Request (Retail Associations) Regulations",
      "citation": "SOR/2003-299",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-473",
      "title": "Disclosure on Account Opening by Telephone Request (Trust and Loan Companies) Regulations",
      "citation": "SOR/2001-473",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-267",
      "title": "Disclosure on Continuance Regulations (Federal Credit Unions)",
      "citation": "SOR/2012-267",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-383",
      "title": "Discontinuance and Continuance of Proceedings Order, 1996",
      "citation": "SOR/96-383",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-544",
      "title": "Discretionary Interest By-law",
      "citation": "SOR/96-544",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-177",
      "title": "Disposal at Sea Permit Application Regulations",
      "citation": "SOR/2014-177",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-275",
      "title": "Disposal at Sea Regulations",
      "citation": "SOR/2001-275",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-142",
      "title": "Order Respecting the Withdrawal From Disposal of Certain Lands in the Northwest Territories (Edéhzhíe (Horn Plateau), N.W.T.)",
      "citation": "SI/2002-142",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-59",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Edéhzhíe (Horn Plateau)) Order",
      "citation": "SI/2013-59",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-948",
      "title": "Disposal of Forfeited Goods and Chattels Regulations",
      "citation": "CRC, c 948",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-321",
      "title": "Distilled Spirits for Bottling in Bond Remission Order",
      "citation": "SOR/97-321",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-303",
      "title": "Distributing Bank and Distributing Bank Holding Company Regulations",
      "citation": "SOR/2006-303",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-305",
      "title": "Distributing Company and Distributing Insurance Holding Company Regulations",
      "citation": "SOR/2006-305",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-304",
      "title": "Distributing Cooperative Credit Association Regulations",
      "citation": "SOR/2006-304",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-306",
      "title": "Distributing Trust and Loan Company Regulations",
      "citation": "SOR/2006-306",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-50",
      "title": "Diversion of Imported Goods Exemption Regulations",
      "citation": "SOR/98-50",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-446",
      "title": "Divestiture of Service Transitional Coverage Regulations",
      "citation": "SOR/98-446",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-252",
      "title": "Division of Judges\u0027 Annuity Benefits Regulations",
      "citation": "SOR/2008-252",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-181",
      "title": "DNA Data Bank Advisory Committee Regulations",
      "citation": "SOR/2000-181",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-300",
      "title": "DNA Identification Regulations",
      "citation": "SOR/2000-300",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-836",
      "title": "Dogfish Exemption Notice",
      "citation": "CRC, c 836",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-514",
      "title": "DOMCO Remission Order No. 2",
      "citation": "SOR/88-514",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-698",
      "title": "Domestic Bonds of Canada Regulations",
      "citation": "CRC, c 698",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-321",
      "title": "Domestic Ferries Security Regulations",
      "citation": "SOR/2009-321",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-116",
      "title": "Domestic Spirits Destroyed Remission Order",
      "citation": "SI/87-116",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-759",
      "title": "Domestic Wine Spirits Remission Order",
      "citation": "CRC, c 759",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1378",
      "title": "Dominion Atlantic Railway Traffic Rules and Regulations",
      "citation": "CRC, c 1378",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-57-216",
      "title": "Dominion Succession Duty (1957) Regulations",
      "citation": "SOR/57-216",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-54-743",
      "title": "Dominion Succession Duty Regulations",
      "citation": "SOR/54-743",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1603",
      "title": "Dominion Water Power Regulations",
      "citation": "CRC, c 1603",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-424",
      "title": "Drug Evaluation Fees Regulations",
      "citation": "SOR/95-424",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-734",
      "title": "Dryden Airport Zoning Regulations",
      "citation": "SOR/84-734",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-430",
      "title": "Dundee Estates Limited Sale Authorization Order",
      "citation": "SOR/90-430",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-44",
      "title": "Duties Relief Regulations",
      "citation": "SOR/96-44",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1072",
      "title": "Duty Free Shop Regulations",
      "citation": "SOR/86-1072",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-972",
      "title": "Earlton Airport Zoning Regulations",
      "citation": "SOR/84-972",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-987",
      "title": "East Coast and Great Lakes Shipping Employees Hours of Work Regulations, 1985",
      "citation": "CRC, c 987",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-99",
      "title": "Eastern Canada Vessel Traffic Services Zone Regulations",
      "citation": "SOR/89-99",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-356",
      "title": "Eastern Townships Wood Producers\u0027 Levies (Interprovincial and Export Trade) Order",
      "citation": "SOR/94-356",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-356",
      "title": "Eastern Townships Wood Producers\u0027 Levies (Interprovincial and Export Trade) Order",
      "citation": "SOR/94-356",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-294",
      "title": "Eastport Marine Protected Areas Regulations",
      "citation": "SOR/2005-294",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-86",
      "title": "Edmonton Garrison Heliport Zoning Regulations",
      "citation": "SOR/2004-86",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-81",
      "title": "Edmonton International Airport Zoning Regulations",
      "citation": "CRC, c 81",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-296",
      "title": "Educational Program, Work and Other Subject-matter Record-keeping Regulations",
      "citation": "SOR/2001-296",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-84",
      "title": "EEC Aged Cheddar Cheese Export Regulations",
      "citation": "SOR/91-84",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-448",
      "title": "Eggplants and Tomatoes Production (Central Saanich) Restriction Regulations",
      "citation": "SOR/82-448",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-284",
      "title": "Egg Regulations",
      "citation": "CRC, c 284",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-67",
      "title": "Egmont Group of Financial Intelligence Units Privileges and Immunities Order",
      "citation": "SOR/2007-67",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-518",
      "title": "Eldorado Nuclear Limited Exemption Order",
      "citation": "SOR/88-518",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-414",
      "title": "Regulations Respecting the Election of Directors of the Canadian Wheat Board",
      "citation": "SOR/98-414",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-75",
      "title": "Proclamation Establishing Electoral Boundaries Commissions",
      "citation": "SI/2002-75",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-9",
      "title": "Proclamation Establishing Electoral Boundaries Commissions",
      "citation": "SI/2012-9",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-131",
      "title": "Electricity and Gas Inspection Regulations",
      "citation": "SOR/86-131",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-59",
      "title": "Electrolux Remission Order",
      "citation": "SI/2005-59",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-115",
      "title": "Electronic Alternatives Regulations for the Purposes of Subsection 254(1) of the Canada Labour Code",
      "citation": "SOR/2008-115",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-308",
      "title": "Electronic Alternatives Regulations for the Purposes of the Federal Real Property and Federal Immovables Act",
      "citation": "SOR/2004-308",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-221",
      "title": "Electronic Commerce Protection Regulations",
      "citation": "SOR/2013-221",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-36",
      "title": "Electronic Commerce Protection Regulations (CRTC)",
      "citation": "SOR/2012-36",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-117",
      "title": "Electronic Documents and Electronic Information Regulations",
      "citation": "SOR/2014-117",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-239",
      "title": "Electronic Documents (Banks and Bank Holding Companies) Regulations",
      "citation": "SOR/2010-239",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-242",
      "title": "Electronic Documents (Cooperative Credit Associations) Regulations",
      "citation": "SOR/2010-242",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-241",
      "title": "Electronic Documents (Insurance and Insurance Holding Companies) Regulations",
      "citation": "SOR/2010-241",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-240",
      "title": "Electronic Documents (Trust and Loan Companies) Regulations",
      "citation": "SOR/2010-240",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-150",
      "title": "Electronic Filing and Provision of Information (GST/HST) Regulations",
      "citation": "SOR/2010-150",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-129",
      "title": "Electronic Payments Regulations",
      "citation": "SOR/98-129",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-156",
      "title": "Electronic Registers and Ancillary Equipment Incorporated in Metering Assemblies Specifications",
      "citation": "SI/90-156",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-256",
      "title": "Eligible Financial Contract General Rules (Bankruptcy and Insolvency Act)",
      "citation": "SOR/2007-256",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-255",
      "title": "Eligible Financial Contract Regulations (Canada Deposit Insurance Corporation Act)",
      "citation": "SOR/2007-255",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-257",
      "title": "Eligible Financial Contract Regulations (Companies\u0027 Creditors Arrangement Act)",
      "citation": "SOR/2007-257",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-257",
      "title": "Eligible Financial Contract Regulations (Companies\u0027 Creditors Arrangement Act)",
      "citation": "SOR/2007-257",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-258",
      "title": "Eligible Financial Contract Regulations (Winding-up and Restructuring Act)",
      "citation": "SOR/2007-258",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-281",
      "title": "Eligible Mortgage Loan Regulations",
      "citation": "SOR/2012-281",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-266",
      "title": "Emergency Protection Orders Regulations",
      "citation": "SOR/2014-266",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-423",
      "title": "Regulations Adapting the Employment Equity Act in Respect of the Canadian Security Intelligence Service",
      "citation": "SOR/2002-423",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-30a",
      "title": "Employment Equity Programs Exclusion Approval Order",
      "citation": "SOR/89-30a",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-30b",
      "title": "Employment Equity Programs Regulations",
      "citation": "SOR/89-30b",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-470",
      "title": "Employment Equity Regulations",
      "citation": "SOR/96-470",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-445",
      "title": "Employment Insurance (Fishing) Regulations",
      "citation": "SOR/96-445",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-332",
      "title": "Employment Insurance Regulations",
      "citation": "SOR/96-332",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-208",
      "title": "Regulations Respecting Employment Related to Mines at the Donkin Coal Block",
      "citation": "SOR/2008-208",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-87",
      "title": "Endeavour Hydrothermal Vents Marine Protected Area Regulations",
      "citation": "SOR/2003-87",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1049",
      "title": "Energy Administration Act Sections 53 to 65 Non-application Order, 1986",
      "citation": "SOR/86-1049",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-590",
      "title": "Energy Conservation Equipment Exemption Regulations",
      "citation": "CRC, c 590",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-651",
      "title": "Energy Efficiency Regulations",
      "citation": "SOR/94-651",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-172",
      "title": "Energy Monitoring Regulations",
      "citation": "SOR/83-172",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-160",
      "title": "Energy Monitoring Regulations, 2006",
      "citation": "SOR/2007-160",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-282",
      "title": "Enhanced Survivor Annuity Regulations",
      "citation": "SOR/2001-282",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-969",
      "title": "Enterprise Development Regulations",
      "citation": "CRC, c 969",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-376",
      "title": "Entity Associated With a Foreign Bank Regulations",
      "citation": "SOR/2001-376",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2008-143",
      "title": "Order Recommending that Each Entity Listed as of July 23, 2008, in the Regulations Establishing a List of Entities Remain a Listed Entity",
      "citation": "SI/2008-143",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-132",
      "title": "Entity Member of Group Regulations",
      "citation": "SOR/2002-132",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-443",
      "title": "Environmental Assessment Review Panel Service Charges Order",
      "citation": "SOR/98-443",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-307",
      "title": "Environmental Emergency Regulations",
      "citation": "SOR/2003-307",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-275",
      "title": "Environmental Response Arrangements Regulations",
      "citation": "SOR/2008-275",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-641",
      "title": "Environmental Studies Research Fund Regions Regulations",
      "citation": "SOR/87-641",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1082",
      "title": "Equal Wages Guidelines, 1986",
      "citation": "SOR/86-1082",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-377",
      "title": "Equity of a Bank or a Bank Holding Company Regulations",
      "citation": "SOR/2001-377",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-378",
      "title": "Equity of a Cooperative Credit Association Regulations",
      "citation": "SOR/2001-378",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-380",
      "title": "Equity of an Insurance Company or Insurance Holding Company Regulations",
      "citation": "SOR/2001-380",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-379",
      "title": "Equity of a Trust and Loan Company Regulations",
      "citation": "SOR/2001-379",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-27",
      "title": "Equivalent Courses (GST/HST) Regulations",
      "citation": "SOR/91-27",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-79",
      "title": "E.S. Fox Limited Remission Order",
      "citation": "SOR/2008-79",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-105",
      "title": "Eskasoni Band Council Elections Order",
      "citation": "SOR/2004-105",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-187",
      "title": "Eskasoni Band Council Method of Election Regulations",
      "citation": "SOR/2004-187",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-332",
      "title": "Esquimalt Graving Dock Regulations",
      "citation": "SOR/89-332",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-332",
      "title": "Esquimalt Graving Dock Regulations",
      "citation": "SOR/89-332",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-4",
      "title": "Establishment Licensing Fees (Veterinary Drugs) Regulations",
      "citation": "SOR/98-4",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-4",
      "title": "Establishment Licensing Fees (Veterinary Drugs) Regulations",
      "citation": "SOR/98-4",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1048",
      "title": "Estates Regulations",
      "citation": "CRC, c 1048",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1308",
      "title": "European Communities Privileges and Immunities Order",
      "citation": "CRC, c 1308",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-53",
      "title": "European Community Monitor Mission Medal (Yugoslavia) Order",
      "citation": "SI/94-53",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-162",
      "title": "European Security and Defence Policy Service Medal Order",
      "citation": "SI/2004-162",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-464",
      "title": "European Space Agency Privileges Order",
      "citation": "SOR/82-464",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-317",
      "title": "European Union Surtax Order",
      "citation": "SOR/99-317",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-196",
      "title": "Evaluation of Impaired Operation (Drugs and Alcohol) Regulations",
      "citation": "SOR/2008-196",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-325",
      "title": "Exceptions for Educational Institutions, Libraries, Archives and Museums Regulations",
      "citation": "SOR/99-325",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-273",
      "title": "Exchange Rate (Authorized Foreign Banks) Regulations",
      "citation": "SOR/99-273",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-571",
      "title": "Excise Act Licence Fees Regulations",
      "citation": "CRC, c 571",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-680",
      "title": "Excise Duty Indexing Ratio Regulations",
      "citation": "SOR/83-680",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-115",
      "title": "Regulations Respecting Excise Licences and Registrations",
      "citation": "SOR/2003-115",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-681",
      "title": "Excise Tax Indexing Ratio Regulations",
      "citation": "SOR/83-681",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-63",
      "title": "Regulations Excluding Certain Indictable Offences From the Definition of \"Designated Offence\"",
      "citation": "SOR/2002-63",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-146",
      "title": "Regulations Excluding Certain Things From the Definition of \"Goods\" Under the Preclearance Act",
      "citation": "SOR/2002-146",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-125",
      "title": "Exclusion Approval Order for Certain Employees and Certain Positions (Air Traffic Control Group), 1989",
      "citation": "SI/89-125",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-639",
      "title": "Exclusion List Regulations",
      "citation": "SOR/94-639",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-108",
      "title": "Exclusion List Regulations, 2007",
      "citation": "SOR/2007-108",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-465",
      "title": "Regulations Prescribing Exclusions From Certain Definitions of the Criminal Code (International Sporting Competition Handguns)",
      "citation": "SOR/98-465",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1590",
      "title": "Execution of Purchase of Property Documents Regulations",
      "citation": "CRC, c 1590",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-63",
      "title": "Order Designating the Executive Director of the Public Appointments Commission Secretariat as Deputy Head",
      "citation": "SI/2006-63",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-381",
      "title": "Exempt Classes of Foreign Banks Regulations",
      "citation": "SOR/2001-381",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-382",
      "title": "Exempt Debt Obligation Transactions (Banks and Bank Holding Companies) Regulations",
      "citation": "SOR/2001-382",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-238",
      "title": "Exemption for Public Notices or Documents (Banks and Bank Holding Companies) Regulations",
      "citation": "SOR/2010-238",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-237",
      "title": "Exemption for Public Notices or Documents (Cooperative Credit Associations) Regulations",
      "citation": "SOR/2010-237",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-236",
      "title": "Exemption for Public Notices or Documents (Insurance Companies and Insurance Holding Companies) Regulations",
      "citation": "SOR/2010-236",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-235",
      "title": "Exemption for Public Notices or Documents (Trust and Loan Companies) Regulations",
      "citation": "SOR/2010-235",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-242",
      "title": "Exemption from Approval for Certain Investments in Intragroup Service Entities (Bank Act) Regulations",
      "citation": "SOR/2003-242",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-243",
      "title": "Exemption from Approval for Certain Investments in Intragroup Service Entities (Cooperative Credit Associations) Regulations",
      "citation": "SOR/2003-243",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-244",
      "title": "Exemption From Approval for Certain Investments in Intragroup Service Entities (Life Companies and Insurance Holding Companies) Regulations",
      "citation": "SOR/2003-244",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-245",
      "title": "Exemption From Approval for Certain Investments in Intragroup Service Entities (Trust and Loan Companies) Regulations",
      "citation": "SOR/2003-245",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-382",
      "title": "Exemption From Deposit Insurance By-law (Exemption Fee)",
      "citation": "SOR/99-382",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-384",
      "title": "Exemption From Deposit Insurance By-law (Foreign Currency Deposits)",
      "citation": "SOR/99-384",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-383",
      "title": "Exemption From Deposit Insurance By-law (Interest on Deposits)",
      "citation": "SOR/99-383",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-381",
      "title": "Exemption From Deposit Insurance By-law (Notice to Depositors)",
      "citation": "SOR/99-381",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-42",
      "title": "Exemption from Deposit Insurance By-Law (Prescribed Deposits)",
      "citation": "SOR/2000-42",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-383",
      "title": "Exemption from Restrictions on Investments (Banks, Bank Holding Companies and Foreign Banks) Regulations",
      "citation": "SOR/2001-383",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-384",
      "title": "Exemption From Restrictions on Investments (Cooperative Credit Associations) Regulations",
      "citation": "SOR/2001-384",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-385",
      "title": "Exemption From Restrictions on Investments (Insurance Companies, Insurance Holding Companies and Societies) Regulations",
      "citation": "SOR/2001-385",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-386",
      "title": "Exemption From Restrictions on Investments (Trust and Loan Companies) Regulations",
      "citation": "SOR/2001-386",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-13",
      "title": "Exemption List Regulations",
      "citation": "SOR/99-13",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-190",
      "title": "Exemption Order for Certain Licences, Authorizations and Documents (White Sturgeon)",
      "citation": "SOR/2006-190",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-429",
      "title": "Exemption Order No. 1, 1997 (Sending of Notices and Documents)",
      "citation": "SOR/97-429",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-116",
      "title": "Radiocommunication Act (Subsection 4(1) and Paragraph 9(1)(b)) Exemption Order (Security, Safety and International Relations), No. 2010–4",
      "citation": "SOR/2010-116",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-154",
      "title": "Exemption Regulations (Beef and Veal Imports)",
      "citation": "SOR/95-154",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-45",
      "title": "Exemption Regulations (Persons)",
      "citation": "SOR/2008-45",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-38",
      "title": "Exempt Personal Information Bank Order, No. 5 (ND)",
      "citation": "SOR/85-38",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-149",
      "title": "Exempt Personal Information Bank Order, No. 13 (RCMP)",
      "citation": "SOR/90-149",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-688",
      "title": "Exempt Personal Information Bank Order, No. 14 (CSIS)",
      "citation": "SOR/92-688",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-272",
      "title": "Exempt Personal Information Bank Order, No. 25 (RCMP)",
      "citation": "SOR/93-272",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2009-24",
      "title": "Order Respecting Ex-Gratia Payments to any Person or Entity That Should Receive Those Payments on Behalf of a Deceased Person who was a Chinese Head Tax Payer or who was in a Conjugal Relationship With a Head Tax Payer",
      "citation": "SI/2009-24",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-109",
      "title": "Order Respecting Ex-Gratia Payments to Chinese Head Tax Payers",
      "citation": "SI/2006-109",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-137",
      "title": "Order Respecting Ex-Gratia Payments to Persons who Were in Conjugal Relationships With now Deceased Chinese Head Tax Payers or to Designated Beneficiaries",
      "citation": "SI/2006-137",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2008-95",
      "title": "Order Respecting Ex Gratia Payments to Veterans and Science and Technology Workers Involved in Nuclear Weapons Testing or Nuclear Decontamination",
      "citation": "SI/2008-95",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-134",
      "title": "Order Respecting Ex Gratia Payments to Veterans Involved in Chemical Warfare Agent Testing",
      "citation": "SI/2006-134",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-95",
      "title": "Experimental Lakes Area Research Activities Regulations",
      "citation": "SOR/2014-95",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-449",
      "title": "Experts on Missions for the Organization for the Prohibition of Chemical Weapons Privileges and Immunities in Canada Order",
      "citation": "SOR/97-449",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-121",
      "title": "Expiry of Section 12.2 of the Softwood Lumber Products Export Charge Act, 2006 Regulations",
      "citation": "SOR/2014-121",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-130",
      "title": "Expiry of the Application of Section 12.1 of the Softwood Lumber Products Export Charge Act, 2006 Regulations",
      "citation": "SOR/2011-130",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-209",
      "title": "Exploitation of the Donkin Coal Block (Natural Resources) Regulations",
      "citation": "SOR/2008-209",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-599",
      "title": "Explosives Regulations",
      "citation": "CRC, c 599",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-211",
      "title": "Explosives Regulations, 2013",
      "citation": "SOR/2013-211",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-149",
      "title": "Export and Import of Hazardous Waste and Hazardous Recyclable Material Regulations",
      "citation": "SOR/2005-149",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-15",
      "title": "Export and Import of Rough Diamonds Regulations",
      "citation": "SOR/2003-15",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-245",
      "title": "Export and Import Permits and Certificates Fees Order",
      "citation": "SOR/95-245",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-202",
      "title": "Export Control List",
      "citation": "SOR/89-202",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-108",
      "title": "Export Control List Notification Regulations",
      "citation": "SOR/2000-108",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-410",
      "title": "Export Development Canada Exercise of Certain Powers Regulations",
      "citation": "SOR/94-410",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-80-11",
      "title": "Exported Dairy Products Assistance Payments Order",
      "citation": "SI/80-11",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-34",
      "title": "Exported Motor Vehicles Drawback Regulations",
      "citation": "SOR/96-34",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-762",
      "title": "Exported Vessels Remission Order",
      "citation": "CRC, c 762",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-71",
      "title": "Exporters\u0027 and Producers\u0027 Records Regulations",
      "citation": "SOR/97-71",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-3",
      "title": "Export Inspection and Certification Exemption Regulations",
      "citation": "SOR/91-3",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-609",
      "title": "Export of Consumable Stores Supplied to Vessels and Aircraft Permit",
      "citation": "CRC, c 609",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-611",
      "title": "Export of Goods for Special and Personal Use Permit",
      "citation": "CRC, c 611",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-612",
      "title": "Export of Logs Permit",
      "citation": "CRC, c 612",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-613",
      "title": "Export of One Cent Bronze Coins Permit",
      "citation": "CRC, c 613",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-616",
      "title": "Export of Specimens Permit",
      "citation": "CRC, c 616",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-88",
      "title": "Export of Substances on the Export Control List Regulations",
      "citation": "SOR/2013-88",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-317",
      "title": "Export of Substances Under the Rotterdam Convention Regulations",
      "citation": "SOR/2002-317",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-722",
      "title": "Export of Sugar Permit",
      "citation": "SOR/83-722",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-204",
      "title": "Export Permits Regulations",
      "citation": "SOR/97-204",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-319",
      "title": "Export Permits Regulations (Softwood Lumber Products)",
      "citation": "SOR/96-319",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-15",
      "title": "Export Permits Regulations (Softwood Lumber Products 2006)",
      "citation": "SOR/2007-15",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-321",
      "title": "Export Permit (Steel Monitoring) Regulations",
      "citation": "SOR/87-321",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-763",
      "title": "Exposed and Processed Film and Recorded Video Tape Remission Order",
      "citation": "CRC, c 763",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-300",
      "title": "Regulations Defining Certain Expressions for the Purpose of Section 21.1 of the Foreign Publishers Advertising Services Act",
      "citation": "SOR/99-300",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-62",
      "title": "Regulations Defining Certain Expressions for the Purposes of the Customs Tariff",
      "citation": "SOR/97-62",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-640",
      "title": "Expropriation Act Basic Rate Order",
      "citation": "CRC, c 640",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-142",
      "title": "Expropriation Fees Regulations",
      "citation": "SOR/2000-142",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-172a",
      "title": "Department of External Affairs Terms Under Three Months Exclusion Approval Order, 1993",
      "citation": "SOR/93-172a",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-259",
      "title": "Extra-billing and User Charges Information Regulations",
      "citation": "SOR/86-259",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-20",
      "title": "Face Protectors for Ice Hockey and Box Lacrosse Players Regulations",
      "citation": "SOR/2011-20",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-387",
      "title": "Factoring Entity Regulations",
      "citation": "SOR/2001-387",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1015",
      "title": "Fair Wages and Hours of Labour Regulations",
      "citation": "CRC, c 1015",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1621",
      "title": "Fair Wages Policy Order",
      "citation": "CRC, c 1621",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-181",
      "title": "Family Support Orders and Agreements Garnishment Regulations",
      "citation": "SOR/88-181",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-793",
      "title": "FAO Privileges and Immunities Order",
      "citation": "SOR/78-793",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-168",
      "title": "Farm Debt Mediation Regulations",
      "citation": "SOR/98-168",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-814",
      "title": "Farm Debt Secured Creditors Notice Regulations",
      "citation": "SOR/86-814",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-28",
      "title": "Farm Equipment Income Tax and Canada Pension Contributions Remission Order",
      "citation": "SI/92-28",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-645",
      "title": "Farm Improvement Loans Regulations",
      "citation": "CRC, c 645",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-508",
      "title": "Faro Airport Zoning Regulations",
      "citation": "SOR/95-508",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-280",
      "title": "Federal Authorities Regulations",
      "citation": "SOR/96-280",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-351",
      "title": "Federal Book Rebate (GST/HST) Regulations",
      "citation": "SOR/98-351",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-175",
      "title": "Federal Child Support Guidelines",
      "citation": "SOR/97-175",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-22",
      "title": "Federal Courts Immigration and Refugee Protection Rules",
      "citation": "SOR/93-22",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-106",
      "title": "Federal Courts Rules",
      "citation": "SOR/98-106",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-268",
      "title": "Federal Credit Union Conversion Regulations",
      "citation": "SOR/2012-268",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-64",
      "title": "Order Transferring to the Federal Economic Development Agency for Southern Ontario from the Department of Industry the Control and Supervision of that Portion of the Federal Public Administration in the Department of Industry Known as Canada Business Ontario",
      "citation": "SI/2013-64",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-53",
      "title": "Federal Elections Fees Tariff",
      "citation": "SOR/2004-53",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-298",
      "title": "Order Authorizing Federal Employees to Acquire Interests in Certain Lands in the Northwest Territories (Order No. 1, 1999)",
      "citation": "SOR/99-298",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-258",
      "title": "Order Authorizing Federal Employees to Acquire Interests in Certain Lands in the Northwest Territories (Order No. 1, 2001)",
      "citation": "SOR/2001-258",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-438",
      "title": "Order Authorizing Federal Employees to Acquire Interests in Certain Lands in the Northwest Territories (Order No. 2, 1998)",
      "citation": "SOR/98-438",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-422",
      "title": "Order Authorizing Federal Employees to Acquire Interests in Certain Lands in the Northwest Territories (Order No. 3, 1999)",
      "citation": "SOR/99-422",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-11",
      "title": "Order Authorizing Federal Employees to Acquire Interests in Certain Lands in the Northwest Territories (Order No. 05, 2001)",
      "citation": "SOR/2002-11",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-289",
      "title": "Federal Halocarbon Regulations, 2003",
      "citation": "SOR/2003-289",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-5",
      "title": "Federal Mobile PCB Treatment and Destruction Regulations",
      "citation": "SOR/90-5",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-587",
      "title": "Federal-Provincial Fiscal Arrangements and Established Programs Financing Regulations, 1977",
      "citation": "SOR/78-587",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-784",
      "title": "Federal-Provincial Fiscal Arrangements and Established Programs Financing Regulations, 1982",
      "citation": "SOR/83-784",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-194",
      "title": "Federal Provincial Fiscal Arrangements Regulations",
      "citation": "SOR/93-194",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-240",
      "title": "Federal-Provincial Fiscal Arrangements Regulations, 1987",
      "citation": "SOR/87-240",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-100",
      "title": "Federal-Provincial Fiscal Arrangements Regulations, 1999",
      "citation": "SOR/2000-100",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-303",
      "title": "Federal-Provincial Fiscal Arrangements Regulations, 2007",
      "citation": "SOR/2007-303",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-664",
      "title": "Federal-Provincial Reciprocal Taxation Agreements Payment Regulations",
      "citation": "CRC, c 664",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-502",
      "title": "Federal Real Property Regulations",
      "citation": "SOR/92-502",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-433",
      "title": "Federal Referendum Fees Tariff",
      "citation": "SOR/92-433",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-10",
      "title": "Federal Registration of Storage Tank Systems for Petroleum Products and Allied Petroleum Products on Federal Lands or Aboriginal Lands Regulations",
      "citation": "SOR/97-10",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-52",
      "title": "Federal Sales Tax Inventory Rebate Regulations",
      "citation": "SOR/91-52",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-53",
      "title": "Federal Sales Tax New Housing Rebate Regulations",
      "citation": "SOR/91-53",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1027",
      "title": "Feed Grain Transportation and Storage Assistance Regulations",
      "citation": "CRC, c 1027",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-593",
      "title": "Feeds Regulations, 1983",
      "citation": "SOR/83-593",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1028",
      "title": "Fees for Documents Regulations",
      "citation": "SOR/86-1028",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-207",
      "title": "Regulations Respecting the Fees for the Examination of Instruments and the Provision of Tables",
      "citation": "SOR/2003-207",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-79",
      "title": "Fees in Respect of Drugs and Medical Devices Regulations",
      "citation": "SOR/2011-79",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-414",
      "title": "Fees in Respect of Mail Regulations",
      "citation": "SOR/92-414",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-432",
      "title": "Fees in Respect of Medical Devices Regulations",
      "citation": "SOR/98-432",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-173",
      "title": "Regulations Prescribing the Fees to be Paid for a Pest Control Product Application Examination Service Provided by or on Behalf of her Majesty in Right of Canada, for a Right or Privilege to Manufacture or Sell a Pest Control Product in Canada and for Establishing a Maximum Residue Limit in relation to a Pest Control Product",
      "citation": "SOR/97-173",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-418",
      "title": "Order Prescribing the Fee to be Paid by Foreign Nationals to Participate in an International Youth Exchange Program in Canada",
      "citation": "SOR/2000-418",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-338",
      "title": "Fenner Dunlop (Bracebridge) Inc. Remission Order",
      "citation": "SOR/2005-338",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-202",
      "title": "Ferry-Boats, Tankers and Cargo Vessels Remission Order, 2010",
      "citation": "SOR/2010-202",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1026",
      "title": "Ferry Cable Regulations",
      "citation": "SOR/86-1026",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-666",
      "title": "Fertilizers Regulations",
      "citation": "CRC, c 666",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-49",
      "title": "FIFA Men\u0027s U-20 World Cup Canada 2007 Remission Order",
      "citation": "SOR/2007-49",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-388",
      "title": "Finance Entity Regulations",
      "citation": "SOR/2001-388",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-474",
      "title": "Financial Consumer Agency of Canada Assessment of Financial Institutions Regulations",
      "citation": "SOR/2001-474",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-101",
      "title": "Financial Consumer Agency of Canada Designated Violations Regulations",
      "citation": "SOR/2002-101",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-389",
      "title": "Financial Leasing Entity Regulations",
      "citation": "SOR/2001-389",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-26",
      "title": "Financial Services and Financial Institutions (GST/HST) Regulations",
      "citation": "SOR/91-26",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-26",
      "title": "Financial Services and Financial Institutions (GST/HST) Regulations",
      "citation": "SOR/91-26",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-201",
      "title": "Financing Secured by Other Revenues Regulations",
      "citation": "SOR/2011-201",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-131",
      "title": "Fingerprinting, Palmprinting and Photography Order",
      "citation": "SI/92-131",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-83",
      "title": "Fire and Boat Drills Regulations",
      "citation": "SOR/2010-83",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-462",
      "title": "Regulations Prescribing Certain Firearms and Other Weapons, Components and Parts of Weapons, Accessories, Cartridge Magazines, Ammunition and Projectiles as Prohibited or Restricted",
      "citation": "SOR/98-462",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-204",
      "title": "Firearms Fees Regulations",
      "citation": "SOR/98-204",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-54",
      "title": "Firearms Fees Remission Order",
      "citation": "SI/2000-54",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-79",
      "title": "Firearms Fees Remission Order (Licenses)",
      "citation": "SI/2006-79",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-99",
      "title": "Firearms Fees Remission Order (Registration Certificate)",
      "citation": "SI/2001-99",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-138",
      "title": "Firearms Information Regulations (Non-restricted Firearms)",
      "citation": "SOR/2012-138",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-199",
      "title": "Firearms Licences Regulations",
      "citation": "SOR/98-199",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-275",
      "title": "Firearms Marking Regulations",
      "citation": "SOR/2004-275",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-213",
      "title": "Firearms Records Regulations",
      "citation": "SOR/98-213",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-198",
      "title": "Firearms Records Regulations (Classification)",
      "citation": "SOR/2014-198",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-201",
      "title": "Firearms Registration Certificates Regulations",
      "citation": "SOR/98-201",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1422",
      "title": "Fire Detection and Extinguishing Equipment Regulations",
      "citation": "CRC, c 1422",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-54",
      "title": "Proclamation Designating the “Fire Prevention Week”",
      "citation": "SI/2004-54",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-102",
      "title": "Proclamation Designating the “Fire Prevention Week”",
      "citation": "SI/2006-102",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-67",
      "title": "Proclamation Designating the “Fire Prevention Week”",
      "citation": "SI/2005-67",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2007-76",
      "title": "Proclamation Designating the “Fire Prevention Week”",
      "citation": "SI/2007-76",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-30",
      "title": "Proclamation Designating the “Fire Prevention Week”",
      "citation": "SI/2000-30",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-70",
      "title": "Proclamation Designating the “Fire Prevention Week”",
      "citation": "SI/2001-70",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-127",
      "title": "Proclamation Designating the “Fire Prevention Week”",
      "citation": "SI/2002-127",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-110",
      "title": "Proclamation Designating the “Fire Prevention Week”",
      "citation": "SI/2003-110",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-380",
      "title": "Order Respecting the First Legislative Assembly of Nunavut",
      "citation": "SOR/98-380",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-47",
      "title": "First Nation of Nacho Nyak Dun (GST) Remission Order",
      "citation": "SI/2001-47",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-241",
      "title": "First Nations Assessment Appeal Regulations",
      "citation": "SOR/2007-241",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-242",
      "title": "First Nations Assessment Inspection Regulations",
      "citation": "SOR/2007-242",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-231",
      "title": "First Nations Land Registry Regulations",
      "citation": "SOR/2007-231",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-240",
      "title": "First Nations Local Revenue Law Review Regulations",
      "citation": "SOR/2007-240",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-254",
      "title": "First Nations Oil and Gas and Moneys Management Voting Regulations",
      "citation": "SOR/2006-254",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-272",
      "title": "First Nations Oil and Gas Environmental Assessment Regulations",
      "citation": "SOR/2007-272",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-277",
      "title": "First Nations Property Assessment and Taxation (Railway Rights-of-Way) Regulations",
      "citation": "SOR/2007-277",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-244",
      "title": "First Nations Rates and Expenditure Laws Timing Regulations",
      "citation": "SOR/2007-244",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-243",
      "title": "First Nations Taxation Enforcement Regulations",
      "citation": "SOR/2007-243",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-243",
      "title": "First Nations Tax Commissioner Appointment Regulations",
      "citation": "SOR/2006-243",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-239",
      "title": "First Nations Tax Commission Review Procedures Regulations",
      "citation": "SOR/2007-239",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-210",
      "title": "Fiscal Equalization Payments Regulations, 1992",
      "citation": "SOR/92-210",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-299",
      "title": "Fiscal Equalization Payments Regulations, 1994",
      "citation": "SOR/94-299",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-177",
      "title": "Fiscal Equalization Payments Regulations, 1999",
      "citation": "SOR/99-177",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-864",
      "title": "Fisheries Improvement Loans Regulations",
      "citation": "CRC, c 864",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-53",
      "title": "Fishery (General) Regulations",
      "citation": "SOR/93-53",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-812",
      "title": "Fish Health Protection Regulations",
      "citation": "CRC, c 812",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-767",
      "title": "Fishing and Recreational Harbours Regulations",
      "citation": "SOR/78-767",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1549",
      "title": "Fishing Zones of Canada (Zone 6) Order",
      "citation": "CRC, c 1549",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1548",
      "title": "Fishing Zones of Canada (Zones 4 and 5) Order",
      "citation": "CRC, c 1548",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1547",
      "title": "Fishing Zones of Canada (Zones 1, 2 and 3) Order",
      "citation": "CRC, c 1547",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-802",
      "title": "Fish Inspection Regulations",
      "citation": "CRC, c 802",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-258",
      "title": "Fish Toxicant Regulations",
      "citation": "SOR/88-258",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-868",
      "title": "Fitness and Amateur Sport Regulations",
      "citation": "CRC, c 868",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-227",
      "title": "Order Fixing December 12, 1988 as the Day on Which the Revised Statutes of Canada, 1985 Come Into Force",
      "citation": "SI/88-227",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-946",
      "title": "Proclamation Fixing Valuation Days",
      "citation": "CRC, c 946",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1148",
      "title": "Flammable Liquids Bulk Storage Regulations",
      "citation": "CRC, c 1148",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-69",
      "title": "Flin Flon Airport Zoning Regulations",
      "citation": "SOR/92-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-10",
      "title": "Flying Accidents Compensation Regulations",
      "citation": "CRC, c 10",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-457",
      "title": "Fondation québécoise du cancer Inc. Remission Order",
      "citation": "SOR/91-457",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1309",
      "title": "Food and Agricultural Organization Remission Order",
      "citation": "CRC, c 1309",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-870",
      "title": "Food and Drug Regulations",
      "citation": "CRC, c 870",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-52",
      "title": "Order Transferring Food Inspection Activities From the Minister of National Health and Welfare to the Minister of Agriculture and Agri-Food and Transferring Food Safety Policy Activities From the Department of Agriculture and Agri-Food to the Department of National Health and Welfare",
      "citation": "SI/96-52",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-478",
      "title": "Food Research and Development Centre Fees Order",
      "citation": "SOR/92-478",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-970",
      "title": "Footwear and Tanning Industries Assistance Regulations",
      "citation": "CRC, c 970",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-559",
      "title": "Ford Liard Airport Zoning Regulations",
      "citation": "SOR/95-559",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-220",
      "title": "Ford New Holland, Inc. Loan Regulations",
      "citation": "SOR/87-220",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-209",
      "title": "Foreign Aircraft Servicing Equipment Remission Order, 1992",
      "citation": "SI/92-209",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-299",
      "title": "Foreign Bank Representative Offices Regulations",
      "citation": "SOR/92-299",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-66-506",
      "title": "Foreign Claims (Bulgaria) Settlement Regulations",
      "citation": "SOR/66-506",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-81",
      "title": "Foreign Claims (Cuba) Settlement Regulations",
      "citation": "SOR/81-81",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-73-681",
      "title": "Foreign Claims (Czechoslovakia) Settlement Regulations",
      "citation": "SOR/73-681",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-70-527",
      "title": "Foreign Claims (Hungary) Settlement Regulations",
      "citation": "SOR/70-527",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-381",
      "title": "Foreign Claims (People\u0027s Republic of China) Settlement Regulations",
      "citation": "SOR/82-381",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-72-395",
      "title": "Foreign Claims (Poland) Settlement Regulations",
      "citation": "SOR/72-395",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-72-90",
      "title": "Foreign Claims (Romania) Settlement Regulations",
      "citation": "SOR/72-90",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-570",
      "title": "Foreign Companies Prescribed Transactions Regulations",
      "citation": "SOR/98-570",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-584",
      "title": "Foreign Extraterritorial Measures (United States) Order, 1992",
      "citation": "SOR/92-584",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-185",
      "title": "Foreign Institutions Subject to the Canadian Residency Requirements Regulations (Insurance Companies)",
      "citation": "SOR/2003-185",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-186",
      "title": "Foreign Institutions Subject to the Canadian Residency Requirements Regulations (Trust and Loan Companies)",
      "citation": "SOR/2003-186",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-84-50",
      "title": "Foreign Organizations Remission Order, 1983",
      "citation": "SI/84-50",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-416",
      "title": "Foreign Ownership of Land Regulations",
      "citation": "SOR/79-416",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-767",
      "title": "Foreign Ports Transhipped Goods Remission Order",
      "citation": "CRC, c 767",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-815",
      "title": "Foreign Vessel Fishing Regulations",
      "citation": "CRC, c 815",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-95a",
      "title": "Department of Forestry Terms Under Three Months Exclusion Approval Order, 1993",
      "citation": "SOR/93-95a",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-76",
      "title": "Forfeited Property Sharing Regulations",
      "citation": "SOR/95-76",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1331",
      "title": "Formal Documents Regulations",
      "citation": "CRC, c 1331",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-320",
      "title": "Former Members of Parliament Elections for Joint and Survivor Benefits Regulations",
      "citation": "SOR/96-320",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-547",
      "title": "Form of Deeds Relating to Certain Successions of Cree and Naskapi Beneficiaries Regulations",
      "citation": "SOR/89-547",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-1084",
      "title": "Form of Instrument of Cession Regulations",
      "citation": "SOR/85-1084",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-390",
      "title": "Form of Proxy (Banks and Bank Holding Companies) Regulations",
      "citation": "SOR/2001-390",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-591",
      "title": "Formula Refunds Regulations",
      "citation": "CRC, c 591",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-133",
      "title": "Fort Frances Airport Zoning Regulations",
      "citation": "SOR/96-133",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-79",
      "title": "Fort McKay First Nation Oil Sands Regulations",
      "citation": "SOR/2007-79",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-82",
      "title": "Fort McMurray Airport Zoning Regulations",
      "citation": "CRC, c 82",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-119",
      "title": "Fort McPherson Airport Zoning Regulations",
      "citation": "SOR/94-119",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-712",
      "title": "Fort Nelson Airport Zoning Regulations",
      "citation": "SOR/82-712",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-388",
      "title": "Fort Norman Airport Zoning Regulations",
      "citation": "SOR/93-388",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-545",
      "title": "Fort Resolution Airport Zoning Regulations",
      "citation": "SOR/95-545",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-706",
      "title": "Fort Simpson Airport Zoning Regulations",
      "citation": "SOR/81-706",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-567",
      "title": "Fort Smith Airport Zoning Regulations",
      "citation": "SOR/81-567",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-468",
      "title": "Fort St. John Airport Zoning Regulations",
      "citation": "SOR/82-468",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-86",
      "title": "Fort William First Nation Sawmill Regulations",
      "citation": "SOR/2011-86",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-452",
      "title": "Fredericton Airport Zoning Regulations",
      "citation": "SOR/81-452",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-72",
      "title": "Free Trade Agreement Advance Rulings Regulations",
      "citation": "SOR/97-72",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-78",
      "title": "Freezing Assets of Corrupt Foreign Officials (Tunisia and Egypt) Regulations",
      "citation": "SOR/2011-78",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-44",
      "title": "Freezing Assets of Corrupt Foreign Officials (Ukraine) Regulations",
      "citation": "SOR/2014-44",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-285",
      "title": "Fresh Fruit and Vegetable Regulations",
      "citation": "CRC, c 285",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-58",
      "title": "Fresh Fruits and Vegetables (United States Tariff) Import Regulations",
      "citation": "SOR/89-58",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-80-176",
      "title": "Front End Wheel Loader Remission Order",
      "citation": "SI/80-176",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-26",
      "title": "Frontier Lands Petroleum Royalty Regulations",
      "citation": "SOR/92-26",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-230",
      "title": "Frontier Lands Registration Regulations",
      "citation": "SOR/88-230",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-357",
      "title": "Fruit and Vegetable Remission Order, 2003",
      "citation": "SOR/2003-357",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-137",
      "title": "Fruit Remission Order, 2006",
      "citation": "SOR/2006-137",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-236",
      "title": "Fruit Remission Order, 2008",
      "citation": "SOR/2008-236",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-215",
      "title": "Fruit Remission Order, 2011",
      "citation": "SOR/2011-215",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-407",
      "title": "Fuels Information Regulations, No. 1",
      "citation": "CRC, c 407",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-73",
      "title": "Order Transferring Functions From the Department of Human Resources Development to the Canada Customs and Revenue Agency of the National Collection Services and of the Collections Litigation and Advisory Services",
      "citation": "SI/2005-73",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-196",
      "title": "G8 Summit Privileges and Immunities Order, 2002",
      "citation": "SOR/2002-196",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-336",
      "title": "G8 Summit Privileges and Immunities Order, 2010",
      "citation": "SOR/2009-336",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-13",
      "title": "G8 Summit Privileges and Immunities Order, No. 2010-2",
      "citation": "SOR/2010-13",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-62",
      "title": "G20 Summit Privileges and Immunities Order, 2010",
      "citation": "SOR/2010-62",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-95",
      "title": "Gallantry Awards Order",
      "citation": "SI/90-95",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1610",
      "title": "Game Declared in Danger of Becoming Extinct",
      "citation": "CRC, c 1610",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1236",
      "title": "Game Declared in Danger of Becoming Extinct",
      "citation": "CRC, c 1236",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-28",
      "title": "Games of Chance (GST/HST) Regulations",
      "citation": "SOR/91-28",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-83",
      "title": "Gander Airport Zoning Regulations",
      "citation": "CRC, c 83",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1424",
      "title": "Garbage Pollution Prevention Regulations",
      "citation": "CRC, c 1424",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-212",
      "title": "Garnishment and Attachment Regulations",
      "citation": "SOR/83-212",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-43",
      "title": "Gasoline and Gasoline Blend Dispensing Flow Rate Regulations",
      "citation": "SOR/2000-43",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-253",
      "title": "Gasoline Excise Tax Regulations",
      "citation": "SOR/2000-253",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-247",
      "title": "Gasoline Regulations",
      "citation": "SOR/90-247",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-190",
      "title": "Gas Pipeline Uniform Accounting Regulations",
      "citation": "SOR/83-190",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-407",
      "title": "Gatineau Region Wood Producers\u0027 Levy (Interprovincial and Export Trade) Order",
      "citation": "SOR/88-407",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-971",
      "title": "General Adjustment Assistance Regulations",
      "citation": "CRC, c 971",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-594",
      "title": "General Excise and Sales Tax Regulations",
      "citation": "CRC, c 594",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-89",
      "title": "General Export Permit no. 43 — Nuclear Goods and Technology to Certain Destinations",
      "citation": "SOR/2012-89",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-90",
      "title": "General Export Permit no. 44 — Nuclear-Related Dual-use Goods and Technology to Certain Destinations",
      "citation": "SOR/2012-90",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-160",
      "title": "General Export Permit no. 45 — Cryptography for the Development or Production of a Product",
      "citation": "SOR/2012-160",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-107",
      "title": "General Export Permit No. 12 — United States Origin Goods",
      "citation": "SOR/97-107",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-264",
      "title": "General Export Permit No. 37 — Toxic Chemicals and Precursors to the United States",
      "citation": "SOR/98-264",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-265",
      "title": "General Export Permit No. 38 — CWC Toxic Chemical and Precursor Mixtures",
      "citation": "SOR/98-265",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-238",
      "title": "General Export Permit No. 39 — Mass Market Cryptographic Software",
      "citation": "SOR/99-238",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-180",
      "title": "Order Issuing the General Export Permit No. 40 — Certain Industrial Chemicals",
      "citation": "SOR/2003-180",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-1",
      "title": "General Export Permit No. 46 — Cryptography for Use by Certain Consignees",
      "citation": "SOR/2013-1",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-121",
      "title": "General Export Permit No. Ex. 18—Portable Personal Computers and Associated Software",
      "citation": "SI/89-121",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-735",
      "title": "General Export Permit No. Ex. 29 — Eligible Industial Goods",
      "citation": "SOR/94-735",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-734",
      "title": "General Export Permit No. Ex. 30—Certain Industrial Goods to Eligible Countries and Territories",
      "citation": "SOR/94-734",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-41",
      "title": "General Export Permit No. Ex. 31—Peanut Butter",
      "citation": "SOR/95-41",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-322",
      "title": "General Export Permit No. Ex. 82 — Carbon Steel Products",
      "citation": "SOR/87-322",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-158",
      "title": "General Export Permit No. EX. 15 — Eggs",
      "citation": "SOR/84-158",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-159",
      "title": "General Export Permit No. EX. 16 — Egg Products",
      "citation": "SOR/84-159",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-580",
      "title": "General Export Permit No. EX 27 — Nuclear-Related Dual-Use Goods",
      "citation": "SOR/93-580",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-17",
      "title": "General Import Permit no. 80 — Carbon Steel",
      "citation": "SOR/2012-17",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-18",
      "title": "General Import Permit no. 81 — Specialty Steel Products",
      "citation": "SOR/2012-18",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-40",
      "title": "General Import Permit No. 1 — Dairy Products for Personal Use",
      "citation": "SOR/95-40",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-39",
      "title": "General Import Permit No. 2 — Chickens and Chicken Products for Personal Use",
      "citation": "SOR/95-39",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-396",
      "title": "General Import Permit No. 3—Wheat and Wheat Products and Barley and Barley Products for Personal Use",
      "citation": "SOR/95-396",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-80",
      "title": "General Import Permit No. 6—Roses for Personal Use",
      "citation": "SOR/97-80",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-38",
      "title": "General Import Permit No. 7 — Turkeys and Turkey Products for Personal Use",
      "citation": "SOR/95-38",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-42",
      "title": "General Import Permit No. 8—Eggs for Personal Use",
      "citation": "SOR/95-42",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-43",
      "title": "General Import Permit No. 13 — Beef and Veal for Personal Use",
      "citation": "SOR/95-43",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-44",
      "title": "General Import Permit No. 14—Margarine for Personal Use",
      "citation": "SOR/95-44",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-384",
      "title": "General Import Permit No. 19",
      "citation": "SOR/78-384",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-400",
      "title": "General Import Permit No. 20—Wheat and Wheat Products and Barley and Barley Products",
      "citation": "SOR/95-400",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-57",
      "title": "General Import Permit No. 80—Carbon Steel",
      "citation": "SOR/97-57",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-58",
      "title": "General Import Permit No. 81—Specialty Steel Products",
      "citation": "SOR/97-58",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-37",
      "title": "General Import Permit No. 100 — Eligible Agriculture Goods",
      "citation": "SOR/95-37",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-287",
      "title": "General Import Permit No. 102—Yarn or Fabric",
      "citation": "SOR/96-287",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-170",
      "title": "General Import Permit No. 106—Apparel Goods or Other Textile Articles",
      "citation": "SOR/97-170",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-266",
      "title": "General Import Permit No. 108 — CWC Toxic Chemicals and Precursors",
      "citation": "SOR/98-266",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-77",
      "title": "General Import Permit No. 193—Roses",
      "citation": "SOR/97-77",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1425",
      "title": "General Load Line Rules",
      "citation": "CRC, c 1425",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-202",
      "title": "General Nuclear Safety and Control Regulations",
      "citation": "SOR/2000-202",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-132",
      "title": "General Pilotage Regulations",
      "citation": "SOR/2000-132",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-34",
      "title": "General Preferential Tariff and Least Developed Country Tariff Rules of Origin Regulations",
      "citation": "SOR/98-34",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-165",
      "title": "General Preferential Tariff and Least Developed Country Tariff Rules of Origin Regulations",
      "citation": "SOR/2013-165",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-6",
      "title": "General Preferential Tariff Reduction Order, 1996",
      "citation": "SOR/96-6",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-696",
      "title": "General Preferential Tariff Reduction Order, No. 1",
      "citation": "SOR/91-696",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-5",
      "title": "General Preferential Tariff Withdrawal (Certain Automotive Goods) Order",
      "citation": "SOR/96-5",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-161",
      "title": "General Preferential Tariff Withdrawal Order (2013 GPT Review)",
      "citation": "SOR/2013-161",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-173",
      "title": "General Preferential Tariff Withdrawal Order (Bulgaria and Romania)",
      "citation": "SOR/2007-173",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-93",
      "title": "General Preferential Tariff Withdrawal Order (Certain Countries Acceding to the European Union)",
      "citation": "SOR/2004-93",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-174",
      "title": "General Preferential Tariff Withdrawal Order (Republic of Belarus)",
      "citation": "SOR/2007-174",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-295",
      "title": "Gilbert Bay Marine Protected Regulations",
      "citation": "SOR/2005-295",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-409",
      "title": "Gillam Airport Zoning Regulations",
      "citation": "SOR/93-409",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-467",
      "title": "Ginn and Company (Canada) and GLC Publishers Limited Acquisition of Shares Order",
      "citation": "SOR/88-467",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-90",
      "title": "Proclamation Giving Notice That the Agreement on Social Security Between Canada and Romania and the Administrative Agreement Between the Government of Canada and the Government of Romania for the Application of the Agreement on Social Security Between Canada and Romania Comes Into Force on November 1, 2011",
      "citation": "SI/2011-90",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-89",
      "title": "Proclamation Giving Notice That the Agreement on Social Security Between Canada and the Republic of Macedonia Comes Into Force on November 1, 2011",
      "citation": "SI/2011-89",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-50",
      "title": "Proclamation Giving Notice That the Memorandum of Understanding Between the Government of Canada and the Government of the United Kingdom of Great Britain and Northern Ireland Concerning Co-operation and Mutual Assistance in the Administration of Social Security Programmes Comes Into Force April 1, 1998",
      "citation": "SI/98-50",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-325",
      "title": "Gjoa Haven Airport Zoning Regulations",
      "citation": "SOR/93-325",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-110",
      "title": "Glass Doors and Enclosures Regulations",
      "citation": "SOR/2009-110",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-176",
      "title": "Glazed Ceramics and Glassware Regulations",
      "citation": "SOR/98-176",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-210",
      "title": "GM Loan Regulations",
      "citation": "SOR/87-210",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-307",
      "title": "Going-Private Transaction (Banks and Bank Holding Companies) Regulations",
      "citation": "SOR/2006-307",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-308",
      "title": "Going-Private Transaction (Insurance Companies and Insurance Holding Companies) Regulations",
      "citation": "SOR/2006-308",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-309",
      "title": "Going-Private Transaction (Trust and Loan Companies) Regulations",
      "citation": "SOR/2006-309",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-330",
      "title": "Golden Nematode Compensation Regulations",
      "citation": "SOR/2006-330",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-260",
      "title": "Golden Nematode Order",
      "citation": "SOR/80-260",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-33",
      "title": "Goods and Services Tax Builders Remission Order",
      "citation": "SI/95-33",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-768",
      "title": "Goods for Emergency Use Remission Order",
      "citation": "CRC, c 768",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-82-246",
      "title": "Goods for Production of Aircraft Components Remission Order No. 2",
      "citation": "SI/82-246",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-948",
      "title": "Goods for Ships and Aircraft (Excise) Drawback Regulations",
      "citation": "SOR/86-948",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-38",
      "title": "Regulations Exempting Certain Goods From the Application of Subsection 118(1) of the Customs Tariff",
      "citation": "SOR/98-38",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-489",
      "title": "Goods Imported and Exported Drawback Regulations",
      "citation": "CRC, c 489",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-42",
      "title": "Goods Imported and Exported Refund and Drawback Regulations",
      "citation": "SOR/96-42",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-57",
      "title": "Goods Imported by Designated Foreign Countries, Military Service Agencies and Institutions (Tariff Item No. 9810.00.00) Regulations",
      "citation": "SOR/98-57",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-102",
      "title": "Goods Imported for Certification Remission Order",
      "citation": "SI/87-102",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-326",
      "title": "Regulations Exempting Goods of Chile From the Application of Anti-dumping Measures",
      "citation": "SOR/97-326",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-6",
      "title": "Goodyear Remission Order",
      "citation": "SI/88-6",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-518",
      "title": "Goose Bay Airport Zoning Regulations",
      "citation": "SOR/2001-518",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-74",
      "title": "Gore Bay Airport Zoning Regulations",
      "citation": "SOR/93-74",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-373",
      "title": "Government Airport Concession Operations Regulations",
      "citation": "SOR/79-373",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-214",
      "title": "Government and Long-Term Corporate Debt Obligations Remission Order",
      "citation": "SI/85-214",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-879",
      "title": "Government Annuities Regulations",
      "citation": "CRC, c 879",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-402",
      "title": "Government Contracts Regulations",
      "citation": "SOR/87-402",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-43",
      "title": "Proclamation Declaring That the Government Corporations Operation Act is Applicable to BDC Capital Inc.",
      "citation": "SOR/2006-43",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-312",
      "title": "Proclamation Declaring That the Government Corporations Operation Act is Applicable to Canada Lands Company Limited Beginning on September 16, 2003",
      "citation": "SOR/2003-312",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-313",
      "title": "Proclamation Declaring That the Government Corporations Operation Act is Applicable to Parc Downsview Park Inc. Beginning on September 16, 2003",
      "citation": "SOR/2003-313",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-578",
      "title": "Proclamation Declaring That the Government Corporations Operation Act is Applicable to the Federal Bridge Corporation Limited Beginning on December 1, 1998",
      "citation": "SOR/98-578",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-791",
      "title": "Government Employees Compensation Place of Employment Regulations",
      "citation": "SOR/86-791",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-880",
      "title": "Government Employees Compensation Regulations",
      "citation": "CRC, c 880",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-300",
      "title": "Government Employees Land Acquisition Order, 1994, No. 1",
      "citation": "SOR/94-300",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-478",
      "title": "Government Employees Land Acquisition Order, 1994, No. 3",
      "citation": "SOR/94-478",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-782",
      "title": "Government Employees Land Acquisition Order, 1994, No. 5",
      "citation": "SOR/94-782",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-92",
      "title": "Government Employees Land Acquisition Order, 1995, No. 1",
      "citation": "SOR/95-92",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-302",
      "title": "Government Employees Land Acquisition Order, 1995, No. 2",
      "citation": "SOR/95-302",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-552",
      "title": "Government Employees Land Acquisition Order, 1995, No. 3",
      "citation": "SOR/95-552",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-134",
      "title": "Government Employees Land Acquisition Order, 1996, No. 1",
      "citation": "SOR/96-134",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-202",
      "title": "Government Employees Land Acquisition Order, 1996, No. 2",
      "citation": "SOR/96-202",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-93",
      "title": "Government of Jamaica Remission Order",
      "citation": "SI/2001-93",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-887",
      "title": "Government Property Traffic Regulations",
      "citation": "CRC, c 887",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-71",
      "title": "Regulations Excluding Certain Government Ships From the Application of the Canada Shipping Act",
      "citation": "SOR/2000-71",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-355",
      "title": "Governor in Council Authority Delegation Order",
      "citation": "CRC, c 355",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-275",
      "title": "Grade Crossings Regulations",
      "citation": "SOR/2014-275",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1427",
      "title": "Grain Cargo Regulations",
      "citation": "CRC, c 1427",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-84",
      "title": "Grande Prairie Airport Zoning Regulations",
      "citation": "CRC, c 84",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-387",
      "title": "Grand Manan Airport Zoning Regulations",
      "citation": "SOR/93-387",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1379",
      "title": "Grand River Railway Traffic Rules and Regulations",
      "citation": "CRC, c 1379",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2008-148",
      "title": "Order Governing the Grant of the Memorial Cross",
      "citation": "SI/2008-148",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-141",
      "title": "Order Governing the Grant of the Memorial Cross (Canadian Forces)",
      "citation": "SI/2006-141",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-690",
      "title": "Great Lakes Fishery Commission Privileges and Immunities Order",
      "citation": "SOR/81-690",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1266",
      "title": "Great Lakes Pilotage Regulations",
      "citation": "CRC, c 1266",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-253",
      "title": "Great Lakes Pilotage Tariff Regulations",
      "citation": "SOR/84-253",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1429",
      "title": "Great Lakes Sewage Pollution Prevention Regulations",
      "citation": "CRC, c 1429",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-587",
      "title": "Greenhouse Cucumber Stabilization Regulations, 1979",
      "citation": "SOR/80-587",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-16",
      "title": "Greenhouse Cucumber Stabilization Regulations, 1980",
      "citation": "SOR/82-16",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-201",
      "title": "Greenhouse Cucumber Stabilization Regulations, 1981",
      "citation": "SOR/83-201",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-22",
      "title": "Greenhouse Cucumber Stabilization Regulations, 1982",
      "citation": "SOR/84-22",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-188",
      "title": "Greenhouse Tomato Stabilization Regulations, 1981",
      "citation": "SOR/83-188",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-558",
      "title": "Greenwood Airport Zoning Regulations",
      "citation": "SOR/95-558",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-205",
      "title": "Gros Morne National Park of Canada Snowshoe Hare Regulations",
      "citation": "SOR/2005-205",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-204",
      "title": "Gros Morne National Park of Canada Timber Harvesting Regulations",
      "citation": "SOR/2005-204",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1245",
      "title": "Grosse Isle, P.Q., Prohibited Place Order",
      "citation": "CRC, c 1245",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-62",
      "title": "Gross Revenue Insurance Plan Account for the Province of Manitoba Remission Order",
      "citation": "SI/2002-62",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-61",
      "title": "Gross Revenue Insurance Program Account for the Province of Alberta Remission Order",
      "citation": "SI/2002-61",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-399",
      "title": "Groupe CSL Inc. Acquisition Exemption Order",
      "citation": "SOR/88-399",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-13",
      "title": "GST Federal Government Departments Remission Order",
      "citation": "SI/91-13",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1579",
      "title": "Guardianship of Veterans\u0027 Property Regulations",
      "citation": "CRC, c 1579",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-278",
      "title": "Guidelines Respecting Control in Fact for the Purpose of Section 377.2 of the Bank Act",
      "citation": "SOR/2012-278",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-163",
      "title": "Guidelines Respecting Control in Fact for the Purpose of Subsection 377(1) of the Bank Act",
      "citation": "SOR/2002-163",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-162",
      "title": "Guidelines Respecting Control in Fact for the Purpose of Subsection 407.2(1) of the Insurance Companies Act",
      "citation": "SOR/2002-162",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-680",
      "title": "Guitar Parts Remission Order, 1994",
      "citation": "SOR/94-680",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-112",
      "title": "Gully Marine Protected Area Regulations",
      "citation": "SOR/2004-112",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-211",
      "title": "Gun Shows Regulations",
      "citation": "SOR/98-211",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-93",
      "title": "Gwaii Haanas National Park Reserve Order",
      "citation": "SOR/96-93",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-293",
      "title": "Haisla Nation Liquefied Natural Gas Facility Regulations",
      "citation": "SOR/2012-293",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-58",
      "title": "Haiti Deemed Direct Shipment (General Preferential Tariff and Least Developed Country Tariff) Regulations",
      "citation": "SOR/2010-58",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-241",
      "title": "Halifax International Airport Zoning Regulations",
      "citation": "SOR/94-241",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-142",
      "title": "Hall Beach Airport Zoning Regulations",
      "citation": "SOR/92-142",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-5",
      "title": "Hamilton Airport Zoning Regulations",
      "citation": "SOR/84-5",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-68",
      "title": "Hampton Place and Taylor Way Remission Order",
      "citation": "SI/2001-68",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-15",
      "title": "Handling of Carloads of Explosives on Railway Trackage Regulations",
      "citation": "SOR/79-15",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-62",
      "title": "Harbourfront Remission Order",
      "citation": "SI/95-62",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-4",
      "title": "Hatchery Exclusion Regulations",
      "citation": "SOR/91-4",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1023",
      "title": "Hatchery Regulations",
      "citation": "CRC, c 1023",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-87",
      "title": "Hay River Airport Zoning Regulations",
      "citation": "CRC, c 87",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-86",
      "title": "Hazardous Materials Information Review Act Appeal Board Procedures Regulations",
      "citation": "SOR/91-86",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-456",
      "title": "Hazardous Materials Information Review Regulations",
      "citation": "SOR/88-456",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-446",
      "title": "Hazardous Products (Booster Cushions) Regulations",
      "citation": "SOR/89-446",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-923",
      "title": "Hazardous Products (Carpet) Regulations",
      "citation": "CRC, c 923",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-732",
      "title": "Hazardous Products (Cellulose Insulation) Regulations",
      "citation": "SOR/79-732",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-924",
      "title": "Hazardous Products (Charcoal) Regulations",
      "citation": "CRC, c 924",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-443",
      "title": "Hazardous Products (Children\u0027s Sleepwear) Regulations",
      "citation": "SOR/87-443",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-151",
      "title": "Hazardous Products (Child Restraint Systems) Regulations",
      "citation": "SOR/88-151",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-440",
      "title": "Hazardous Products (Crocidolite Asbestos) Regulations",
      "citation": "SOR/89-440",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-39",
      "title": "Hazardous Products (Expansion Gates and Expandable Enclosures) Regulations",
      "citation": "SOR/90-39",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-271",
      "title": "Hazardous Products (Infant Feeding Bottle Nipples) Regulations",
      "citation": "SOR/84-271",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-927",
      "title": "Hazardous Products (Kettles) Regulations",
      "citation": "CRC, c 927",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-514",
      "title": "Hazardous Products (Lighters) Regulations",
      "citation": "SOR/89-514",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-929",
      "title": "Hazardous Products (Matches) Regulations",
      "citation": "CRC, c 929",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-810",
      "title": "Hazardous Products (Mattresses) Regulations",
      "citation": "SOR/80-810",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-930",
      "title": "Hazardous Products (Pacifiers) Regulations",
      "citation": "CRC, c 930",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-245",
      "title": "Hazardous Products (Tents) Regulations",
      "citation": "SOR/90-245",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-931",
      "title": "Hazardous Products (Toys) Regulations",
      "citation": "CRC, c 931",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-23",
      "title": "Health Care Services (GST/HST) Regulations",
      "citation": "SOR/91-23",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-399",
      "title": "Health Information Custodians in the Province of Ontario Exemption Order",
      "citation": "SOR/2005-399",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-296",
      "title": "Health of Animals Regulations",
      "citation": "CRC, c 296",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-267",
      "title": "Health Reform Transfer Regulations",
      "citation": "SOR/2003-267",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1151",
      "title": "Heating and Power Boilers Regulations",
      "citation": "CRC, c 1151",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-24",
      "title": "Heavy-duty Vehicle and Engine Greenhouse Gas Emission Regulations",
      "citation": "SOR/2013-24",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1182",
      "title": "Height of Wires of Telegraph and Telephone Lines Regulations",
      "citation": "CRC, c 1182",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-122",
      "title": "Heritage Railway Stations Regulations",
      "citation": "SOR/91-122",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-774",
      "title": "Hibernia Development Project Offshore Application Regulations",
      "citation": "SOR/90-774",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-202",
      "title": "Proclamations Certifying who are the High Contracting Parties to the Warsaw Convention",
      "citation": "SI/85-202",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-25",
      "title": "Proclamations Certifying who are the High Contracting Parties to the Warsaw Convention",
      "citation": "SI/85-25",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-398",
      "title": "Proclamation Certifying who are the High Contracting Parties to the Warsaw Convention",
      "citation": "CRC, c 398",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1183",
      "title": "Highway Crossings Protective Devices Regulations",
      "citation": "CRC, c 1183",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-220",
      "title": "Historic Canals Regulations",
      "citation": "SOR/93-220",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-515",
      "title": "Hog Stabilization 1979-80 Regulations",
      "citation": "SOR/80-515",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-609",
      "title": "Hog Stabilization 1980-81 Regulations",
      "citation": "SOR/81-609",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-707",
      "title": "Hog Stabilization Regulations, 1983-84",
      "citation": "SOR/84-707",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-505",
      "title": "Holman Airport Zoning Regulations",
      "citation": "SOR/95-505",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-75-24",
      "title": "Home Buyer Grant Regulations",
      "citation": "SOR/75-24",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1430",
      "title": "Home-Trade, Inland and Minor Waters Voyages Regulations",
      "citation": "CRC, c 1430",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-136",
      "title": "Honeybee Importation Prohibition Regulations, 2004",
      "citation": "SOR/2004-136",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-287",
      "title": "Honey Regulations",
      "citation": "CRC, c 287",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-59",
      "title": "Honeywell Remission Order",
      "citation": "SI/98-59",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-207",
      "title": "Hong Kong Economic and Trade Office Privileges and Immunities Order",
      "citation": "SOR/96-207",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-70",
      "title": "Order Assigning the Honourable Leona Aglukkaq to Assist the Minister of Foreign Affairs",
      "citation": "SI/2012-70",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-77-664",
      "title": "Horseshoe Falls Water Power Regulations",
      "citation": "SOR/77-664",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-232",
      "title": "Housing Loan (Insurance, Guarantee and Protection) Regulations",
      "citation": "SOR/2012-232",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1431",
      "title": "Hull Construction Regulations",
      "citation": "CRC, c 1431",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1432",
      "title": "Hull Inspection Regulations",
      "citation": "CRC, c 1432",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-558",
      "title": "Human Pathogens Importation Regulations",
      "citation": "SOR/94-558",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-394",
      "title": "Human Rights Tribunal Appeal Regulations",
      "citation": "SOR/80-394",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-353",
      "title": "Hunter Douglas Canada Remission Order",
      "citation": "SOR/2005-353",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-798",
      "title": "IAC Limited Shareholders Order",
      "citation": "SOR/81-798",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-563",
      "title": "ICAO Privileges and Immunities Order",
      "citation": "SOR/94-563",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-21",
      "title": "Ice Hockey Helmet Regulations",
      "citation": "SOR/2011-21",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-10",
      "title": "Icewine Regulations",
      "citation": "SOR/2014-10",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-748",
      "title": "I.C.S.P.R.C.P. Status Order",
      "citation": "SOR/78-748",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-82",
      "title": "Identity Screening Regulations",
      "citation": "SOR/2007-82",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-100",
      "title": "Igloolik Airport Zoning Regulations",
      "citation": "SOR/92-100",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1311",
      "title": "I.I.C.A. Privileges and Immunities Order",
      "citation": "CRC, c 1311",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-227",
      "title": "Immigration and Refugee Protection Regulations",
      "citation": "SOR/2002-227",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-230",
      "title": "Immigration Appeal Division Rules",
      "citation": "SOR/2002-230",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-229",
      "title": "Immigration Division Rules",
      "citation": "SOR/2002-229",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-80-125",
      "title": "Immigration Guidelines",
      "citation": "SI/80-125",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-686",
      "title": "Immigration Investigation Regulations",
      "citation": "SOR/80-686",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-121",
      "title": "Order Restricting Certain Immunity In Relation to the United States",
      "citation": "SOR/97-121",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-204",
      "title": "Regulations Implementing the United Nations Resolution on Lebanon",
      "citation": "SOR/2007-204",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-287",
      "title": "Regulations Implementing the United Nations Resolution on the Democratic People\u0027s Republic of Korea",
      "citation": "SOR/2006-287",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-44",
      "title": "Regulations Implementing the United Nations Resolutions on Iran",
      "citation": "SOR/2007-44",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-360",
      "title": "Regulations Implementing the United Nations Resolutions on the Suppression of Terrorism",
      "citation": "SOR/2001-360",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-36",
      "title": "Import Allocation Regulations",
      "citation": "SOR/95-36",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1504",
      "title": "Import and Export Returns Order",
      "citation": "CRC, c 1504",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-214",
      "title": "Importation and Exportation of Firearms Regulations (Businesses)",
      "citation": "SOR/98-214",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-215",
      "title": "Importation and Exportation of Firearms Regulations (Individuals)",
      "citation": "SOR/98-215",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-533",
      "title": "Importation of Periodicals Regulations",
      "citation": "CRC, c 533",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-202",
      "title": "Importations by Certain Processing Service Companies (GST) Remission Order",
      "citation": "SI/92-202",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-603",
      "title": "Import Certificate Regulations",
      "citation": "CRC, c 603",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-604",
      "title": "Import Control List",
      "citation": "CRC, c 604",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-771",
      "title": "Imported Demonstration Aircraft Remission Order",
      "citation": "CRC, c 771",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1011",
      "title": "Imported Goods Records Regulations",
      "citation": "SOR/86-1011",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-151",
      "title": "Imported Spirits for Blending Remission Order",
      "citation": "SI/83-151",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1078",
      "title": "Import of Arms Permit",
      "citation": "SOR/86-1078",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-72",
      "title": "Import of Chickens Permit",
      "citation": "SOR/79-72",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-623",
      "title": "Import of Dairy Products for Personal Use Permit",
      "citation": "CRC, c 623",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-625",
      "title": "Import of Eggs Permit",
      "citation": "CRC, c 625",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-632",
      "title": "Import of Specimens (Personal or Household) Permit",
      "citation": "CRC, c 632",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-122",
      "title": "Import of Turkeys and Turkey Products Permit",
      "citation": "SOR/81-122",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-5",
      "title": "Import Permits Regulations",
      "citation": "SOR/79-5",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-419",
      "title": "Imports of Certain Textile and Apparel Goods From Chile Customs Duty Remission Order",
      "citation": "SOR/98-419",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-399",
      "title": "Imports of Certain Textile and Apparel Goods From Costa Rica Customs Duty Remission Order",
      "citation": "SOR/2002-399",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-221",
      "title": "Imports of Certain Textile and Apparel Goods from Honduras Customs Duty Remission Order",
      "citation": "SOR/2014-221",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-420",
      "title": "Imports of Certain Textile and Apparel Goods FRrom Mexico or the United States Customs Duty Remission Order",
      "citation": "SOR/98-420",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-637",
      "title": "Inclusion List Regulations",
      "citation": "SOR/94-637",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-96",
      "title": "Income Earned in Quebec Income Tax Remission Order, 1982",
      "citation": "SOR/83-96",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-157",
      "title": "Income Earned in Quebec Income Tax Remission Order, 1988",
      "citation": "SI/89-157",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-82",
      "title": "Income Tax Paid by Investors, Other Than Promoters Remission Order",
      "citation": "SI/96-82",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-80",
      "title": "Income Tax Paid by Investors, Other Than Promoters Remission Order",
      "citation": "SI/96-80",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-945",
      "title": "Income Tax Regulations",
      "citation": "CRC, c 945",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-21",
      "title": "Income Tax Remission Order (Canada Pension Plan)",
      "citation": "SI/95-21",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-108",
      "title": "Indemnification and Advances Regulations for Directors and Officers of Crown Corporations",
      "citation": "SOR/2011-108",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-102",
      "title": "Index-linked Deposits Interest Disclosure Regulations",
      "citation": "SOR/2002-102",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-353",
      "title": "Indian Affairs and Northern Development Aboriginal Peoples Employment Equity Program Appointments Regulations",
      "citation": "SOR/97-353",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-949",
      "title": "Indian Band Council Borrowing Regulations",
      "citation": "CRC, c 949",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-950",
      "title": "Indian Band Council Procedure Regulations",
      "citation": "CRC, c 950",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-952",
      "title": "Indian Band Election Regulations",
      "citation": "CRC, c 952",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-297",
      "title": "Indian Band Revenue Moneys Order",
      "citation": "SOR/90-297",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-138",
      "title": "Indian Bands Council Elections Order",
      "citation": "SOR/97-138",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-46",
      "title": "Indian Bands Council Method of Election Regulations",
      "citation": "SOR/90-46",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-953",
      "title": "Indian Bands Revenue Moneys Regulations",
      "citation": "CRC, c 953",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-89",
      "title": "Indian Bridge Workers Remission Order",
      "citation": "SI/91-89",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-22",
      "title": "Indian Economic Development Direct Loan Order",
      "citation": "SOR/78-22",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-21",
      "title": "Indian Economic Development Guarantee Order",
      "citation": "SOR/78-21",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-954",
      "title": "Indian Estates Regulations",
      "citation": "CRC, c 954",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-44",
      "title": "Indian Income Tax Remission Order",
      "citation": "SI/93-44",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-166",
      "title": "Indian Income Tax Remission Order",
      "citation": "SI/93-166",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-956",
      "title": "Indian Mining Regulations",
      "citation": "CRC, c 956",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-753",
      "title": "Indian Oil and Gas Regulations, 1995",
      "citation": "SOR/94-753",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-957",
      "title": "Indian Referendum Regulations",
      "citation": "CRC, c 957",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-144",
      "title": "Indian Remission Order",
      "citation": "SI/85-144",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-959",
      "title": "Indian Reserve Traffic Regulations",
      "citation": "CRC, c 959",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-960",
      "title": "Indian Reserve Waste Disposal Regulations",
      "citation": "CRC, c 960",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2009-45",
      "title": "Order Designating the Chairperson of the Indian Residential Schools Truth and Reconciliation Commission as Deputy Head",
      "citation": "SI/2009-45",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-102",
      "title": "Indians and Bands on Certain Indian Settlements Remission Order",
      "citation": "SI/92-102",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-127",
      "title": "Indians and Bands on Certain Indian Settlements Remission Order (1997)",
      "citation": "SI/97-127",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-71",
      "title": "Indians and the War Lake First Nation Band on the Ilford Indian Settlement Remission Order",
      "citation": "SI/94-71",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-70",
      "title": "Indians and Webequie Band on the Webequie Indian Settlement Remission Order",
      "citation": "SI/94-70",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-69",
      "title": "Indian Settlements Remission Order (2000)",
      "citation": "SI/2000-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-128",
      "title": "Indian Statute-Barred Income Tax Assessment Remission Order",
      "citation": "SI/89-128",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-109",
      "title": "Indian Timber Harvesting Regulations",
      "citation": "SOR/2002-109",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-961",
      "title": "Indian Timber Regulations",
      "citation": "CRC, c 961",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-599",
      "title": "Industrial and Regional Development Regulations",
      "citation": "SOR/83-599",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-460",
      "title": "Industrial Design Regulations",
      "citation": "SOR/99-460",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-156",
      "title": "Industrial Hemp Regulations",
      "citation": "SOR/98-156",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-265",
      "title": "Industrial Milk and Cream 1989 Period Stabilization Regulations",
      "citation": "SOR/89-265",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-509",
      "title": "Industrial Milk and Cream Stabilization 1978-79 Regulations",
      "citation": "SOR/78-509",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-477",
      "title": "Industrial Milk and Cream Stabilization 1979-80 Regulations",
      "citation": "SOR/79-477",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-336",
      "title": "Industrial Milk and Cream Stabilization 1980-81 Regulations",
      "citation": "SOR/80-336",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-456",
      "title": "Industrial Milk and Cream Stabilization 1981-82 Regulations",
      "citation": "SOR/81-456",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-996",
      "title": "Industrial Milk and Cream Stabilization 1982-83 Regulations",
      "citation": "SOR/82-996",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-581",
      "title": "Industrial Milk and Cream Stabilization 1983-84 Regulations",
      "citation": "SOR/83-581",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-394",
      "title": "Industrial Milk and Cream Stabilization 1984-85 Regulations",
      "citation": "SOR/84-394",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-546",
      "title": "Industrial Milk and Cream Stabilization 1985-86 Regulations",
      "citation": "SOR/85-546",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-687",
      "title": "Industrial Milk and Cream Stabilization 1986-87 Regulations",
      "citation": "SOR/86-687",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-318",
      "title": "Industrial Milk and Cream Stabilization 1988-89 Regulations",
      "citation": "SOR/88-318",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-672",
      "title": "Industrial Product Price Indices Adjustment Regulations",
      "citation": "SOR/90-672",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1600",
      "title": "Infant or Person of Unsound Mind Payment Order",
      "citation": "CRC, c 1600",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-391",
      "title": "Information Processing Activities (Banks and Authorized Foreign Banks) Regulations",
      "citation": "SOR/2001-391",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-47",
      "title": "Regulations Concerning Information Required by Foreign States",
      "citation": "SOR/2002-47",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-60",
      "title": "Information Technology Activities (Authorized Foreign Banks) Regulations",
      "citation": "SOR/2003-60",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-62",
      "title": "Information Technology Activities (Bank Holding Companies) Regulations",
      "citation": "SOR/2003-62",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-61",
      "title": "Information Technology Activities (Banks) Regulations",
      "citation": "SOR/2003-61",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-63",
      "title": "Information Technology Activities (Canadian Societies) Regulations",
      "citation": "SOR/2003-63",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-64",
      "title": "Information Technology Activities (Cooperative Credit Associations) Regulations",
      "citation": "SOR/2003-64",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-65",
      "title": "Information Technology Activities (Foreign Banks) Regulations",
      "citation": "SOR/2003-65",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-66",
      "title": "Information Technology Activities (Insurance Holding Companies) Regulations",
      "citation": "SOR/2003-66",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-67",
      "title": "Information Technology Activities (Life Companies) Regulations",
      "citation": "SOR/2003-67",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-68",
      "title": "Information Technology Activities (Property and Casualty Companies and Marine Companies) Regulations",
      "citation": "SOR/2003-68",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-68",
      "title": "Information Technology Activities (Property and Casualty Companies and Marine Companies) Regulations",
      "citation": "SOR/2003-68",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-69",
      "title": "Information Technology Activities (Trust and Loan Companies) Regulations",
      "citation": "SOR/2003-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-201",
      "title": "Regulations Respecting the Information to be Displayed on Alcohol Containers and Their Packaging",
      "citation": "SOR/2003-201",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-89",
      "title": "Infrastructure Projects Environmental Assessment Adaptation Regulations",
      "citation": "SOR/2009-89",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-64",
      "title": "Ingredient Disclosure List",
      "citation": "SOR/88-64",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-45",
      "title": "Input Tax Credit Information (GST/HST) Regulations",
      "citation": "SOR/91-45",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-310",
      "title": "Insider Reports (Banks and Bank Holding Companies) Regulations",
      "citation": "SOR/2006-310",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-311",
      "title": "Insider Reports (Cooperative Credit Associations) Regulations",
      "citation": "SOR/2006-311",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-113",
      "title": "Insider Reports Exemptions (Banks) Regulations",
      "citation": "SOR/2000-113",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-312",
      "title": "Insider Reports (Insurance Companies and Insurance Holding Companies) Regulations",
      "citation": "SOR/2006-312",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-207",
      "title": "Insider Reports Regulations",
      "citation": "SOR/82-207",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-313",
      "title": "Insider Reports (Trust and Loan Companies) Regulations",
      "citation": "SOR/2006-313",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-958",
      "title": "Inspection and Search Defence Regulations",
      "citation": "SOR/86-958",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-707",
      "title": "Inspection of Defence Materiel Order",
      "citation": "CRC, c 707",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-33",
      "title": "Insurable Earnings and Collection of Premiums Regulations",
      "citation": "SOR/97-33",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-282",
      "title": "Insurable Housing Loan Regulations",
      "citation": "SOR/2012-282",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-270",
      "title": "Insurance Business (Authorized Foreign Banks) Regulations",
      "citation": "SOR/99-270",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-330",
      "title": "Insurance Business (Banks and Bank Holding Companies) Regulations",
      "citation": "SOR/92-330",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-300",
      "title": "Insurance Business (Cooperative Credit Associations) Regulations",
      "citation": "SOR/2003-300",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-331",
      "title": "Insurance Business (Trust and Loan Companies) Regulations",
      "citation": "SOR/92-331",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-182",
      "title": "Insurance Companies Assessed Expenses Recovery Regulations",
      "citation": "SOR/99-182",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-212",
      "title": "Integrated Circuit Topography Regulations",
      "citation": "SOR/93-212",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1312",
      "title": "Inter-American Development Bank Privileges and Immunities Order",
      "citation": "CRC, c 1312",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-188",
      "title": "Interest and Administrative Charges Regulations",
      "citation": "SOR/96-188",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-224",
      "title": "Interest Payable on Certain Deposits By-law",
      "citation": "SOR/99-224",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-19",
      "title": "Interest Rate (Excise Tax Act) Regulations",
      "citation": "SOR/91-19",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1121",
      "title": "Interest Rate for Customs Purposes Regulations",
      "citation": "SOR/86-1121",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-267",
      "title": "Interest Rates (Air Travellers Security Charge Act) Regulations",
      "citation": "SOR/2007-267",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-229",
      "title": "Interest Rates (Excise Act, 2001) Regulations",
      "citation": "SOR/2006-229",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-230",
      "title": "Interest Rates (Excise Tax Act) Regulations",
      "citation": "SOR/2006-230",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-26",
      "title": "Order Respecting the Interim Federal Health Program, 2012",
      "citation": "SI/2012-26",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-226",
      "title": "Interim Payments and Recovery of Overpayments Regulations",
      "citation": "SOR/81-226",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-59",
      "title": "Order Transferring the Internal E-staffing Unit From Human Resources and Skills Development to the Public Service Commission",
      "citation": "SI/2011-59",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-144",
      "title": "Order Designating the Internal Inquiry Into the Actions of Canadian Officials in Relation to Abdullah Almalki, Ahmad Abou-Emaati and Muayyed Nureddin as a Department for Purposes of the Act and the Prime Minister as Appropriate Minister",
      "citation": "SI/2006-144",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-445",
      "title": "International Boundary Waters Regulations",
      "citation": "SOR/2002-445",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-17",
      "title": "International Bridges and Tunnels Regulations",
      "citation": "SOR/2009-17",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1314",
      "title": "International Commission for the Northwest Atlantic Fisheries Privileges and Immunities Order",
      "citation": "CRC, c 1314",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-112",
      "title": "International Cospas-Sarsat Programme Privileges and Immunities Order",
      "citation": "SOR/2005-112",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-156",
      "title": "International Criminal Court Privileges and Immunities Order",
      "citation": "SOR/2004-156",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-137",
      "title": "International Development Association, International Finance Corporation and Multilateral Investment Guarantee Agency Privileges and Immunities Order",
      "citation": "SOR/2014-137",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-39",
      "title": "International Force East Timor (INTERFET) Medal Order",
      "citation": "SI/2001-39",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1315",
      "title": "International Joint Commission Immunity Order",
      "citation": "CRC, c 1315",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-807",
      "title": "International Letter-post Items Regulations",
      "citation": "SOR/83-807",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-581",
      "title": "International Maritime Satellite Organization (INMARSAT) Privileges and Immunities Order",
      "citation": "SOR/94-581",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1316",
      "title": "International North Pacific Fisheries Commission Privileges and Immunities Order",
      "citation": "CRC, c 1316",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-87",
      "title": "International Organization for Migration Privileges and Immunities Order",
      "citation": "SOR/2012-87",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-982",
      "title": "International River Improvements Regulations",
      "citation": "CRC, c 982",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-488",
      "title": "International Submarine Cable Licences Regulations",
      "citation": "SOR/98-488",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-907",
      "title": "International Telecommunications Satellite Organization (INTELSAT) Privileges and Immunities Order",
      "citation": "SOR/81-907",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-230",
      "title": "International Telecommunication Union Privileges and Immunities Order",
      "citation": "SOR/2000-230",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-292",
      "title": "Internet Child Pornography Reporting Regulations",
      "citation": "SOR/2011-292",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-301",
      "title": "Interprovincial Movement of Hazardous Waste Regulations",
      "citation": "SOR/2002-301",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-113",
      "title": "Introduced Forest Pest Compensation Regulations",
      "citation": "SOR/2004-113",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-868",
      "title": "Inuk of Fort George Observer Regulations",
      "citation": "SOR/86-868",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-707",
      "title": "Inuvik Airport Zoning Regulations",
      "citation": "SOR/81-707",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-133",
      "title": "Investigation Drugs, Placebos and Emergency Drugs Remission Order",
      "citation": "SI/85-133",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-611",
      "title": "Investment Canada Regulations",
      "citation": "SOR/85-611",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-392",
      "title": "Investment Limits (Bank Holding Companies) Regulations",
      "citation": "SOR/2001-392",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-393",
      "title": "Investment Limits (Banks) Regulations",
      "citation": "SOR/2001-393",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-394",
      "title": "Investment Limits (Canadian Societies) Regulations",
      "citation": "SOR/2001-394",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-395",
      "title": "Investment Limits (Cooperative Credit Associations) Regulations",
      "citation": "SOR/2001-395",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-274",
      "title": "Investment Limits (Foreign Companies) Regulations",
      "citation": "SOR/92-274",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-274",
      "title": "Investment Limits (Foreign Companies) Regulations",
      "citation": "SOR/92-274",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-396",
      "title": "Investment Limits (Insurance Companies) Regulations",
      "citation": "SOR/2001-396",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-397",
      "title": "Investment Limits (Insurance Holding Companies) Regulations",
      "citation": "SOR/2001-397",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-398",
      "title": "Investment Limits (Trust and Loan Companies) Regulations",
      "citation": "SOR/2001-398",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-301",
      "title": "Investments in Associations and Cooperatively-owned Entities Regulations",
      "citation": "SOR/2003-301",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-708",
      "title": "Investors\u0027 Indemnity Regulations",
      "citation": "CRC, c 708",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-43",
      "title": "Investors in the Norbourg and Evolution Funds Remission Order",
      "citation": "SI/2012-43",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-772",
      "title": "Involuntary Retirements Remission Order",
      "citation": "CRC, c 772",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-161",
      "title": "Iodinated Contrast Media Anti-Dumping Duty Remission Order",
      "citation": "SOR/2001-161",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-770",
      "title": "I.O.S. Income Tax Remission Order",
      "citation": "CRC, c 770",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-136",
      "title": "IRSTD System Remission Order",
      "citation": "SI/85-136",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-587",
      "title": "Issuance of Certificates Regulations",
      "citation": "SOR/93-587",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-188",
      "title": "Order Prohibiting the Issuance of Interests at Lapierre House Historic Site in the Yukon Territory",
      "citation": "SOR/98-188",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-540",
      "title": "Order Prohibiting the Issuance of Interests at Rampart House in the Yukon Territory",
      "citation": "SOR/97-540",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-183",
      "title": "Order Authorizing the Issuance of Non-circulation Coins of the Denomination of Eight Dollars",
      "citation": "SI/2003-183",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-181",
      "title": "Issuance of Writs of Referendum Authorization Order",
      "citation": "SI/92-181",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-138",
      "title": "Order Authorizing the Issue and Determining the Composition, Dimensions and Designs of a Circulation Coin",
      "citation": "SOR/2001-138",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-61",
      "title": "Order Authorizing the Issue and Determining the Composition, Dimensions and Designs of a Fifty Cent Circulation Coin",
      "citation": "SOR/2002-61",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-104",
      "title": "Order Autorizing the Issue and Determining the Composition, Dimensions and Designs of a Five Cent Circulation Coin",
      "citation": "SOR/2005-104",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-133",
      "title": "Order Authorizing the Issue and Determining the Composition, Dimensions and Designs of a One Dollar Circulation Coin",
      "citation": "SOR/2004-133",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-141",
      "title": "Order Authorizing the Issue and Determining the Composition, Dimensions and Designs of a One Dollar Circulation Coin Featuring the Canadian Olympic Committee Symbol",
      "citation": "SOR/2004-141",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-234",
      "title": "Order Authorizing the Issue and Determining the Composition, Dimensions and Designs of a Twenty-five Cent Circulation Coin",
      "citation": "SOR/2006-234",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-278",
      "title": "Order Authorizing the Issue and Determining the Composition, Dimensions and Designs of a Twenty-five Cent Circulation Coin",
      "citation": "SOR/2005-278",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-18",
      "title": "Order Authorizing the Issue and Determining the Composition, Dimensions and Designs of a Twenty-five Cent Circulation Coin",
      "citation": "SOR/2006-18",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-81",
      "title": "Order Authorizing the Issue and Determining the Composition, Dimensions and Designs of a Twenty-five Cent Circulation Coin",
      "citation": "SOR/2004-81",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-153",
      "title": "Order Authorizing the Issue and Determining the Composition, Dimensions and Designs of a Twenty-five Cent Circulation Coin",
      "citation": "SOR/2005-153",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-154",
      "title": "Order Authorizing the Issue and Determining the Composition, Dimensions and Designs of a Twenty-five Cent Circulation Coin",
      "citation": "SOR/2005-154",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-19",
      "title": "Order Authorizing the Issue and Determining the Composition, Dimensions and Designs of a Two Dollar Ciculation Coin",
      "citation": "SOR/2006-19",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-192",
      "title": "Order Authorizing the Issue and Determining the Composition, Dimensions and Designs of Circulation Coins",
      "citation": "SOR/2002-192",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-14",
      "title": "Order Authorizing the Issue and Determining the Composition, Dimensions and Designs of Circulation Coins",
      "citation": "SOR/2002-14",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-20",
      "title": "Order Authorizing the Issue and Determining the Composition, Dimensions and Designs of Various Ciculation Coins",
      "citation": "SOR/2006-20",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-297",
      "title": "Order Authorizing the Issue and Determining the Composition, Dimensions and Designs of Various Circulation Coins",
      "citation": "SOR/2006-297",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-316",
      "title": "Order Authorizing the Issue and Determining the Composition, Dimensions and Designs of Various Circulation Coins",
      "citation": "SOR/2000-316",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-356",
      "title": "Order Authorizing the Issue and Determining the Design of a One Dollar Circulation Coin",
      "citation": "SOR/2005-356",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-191",
      "title": "Order Authorizing the Issue and Determining the Designs of Circulation Coins",
      "citation": "SOR/2003-191",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-541",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Commemorative One Dollar Precious Metal Coin (300th Anniversary of Henry Kelsey\u0027s Voyage of Exploration of the Canadian Prairies)",
      "citation": "SOR/89-541",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-594",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Commemorative One Dollar Precious Metal Coin (Sir Alexander Mackenzie\u0027s Discovery of the Mackenzie River)",
      "citation": "SOR/88-594",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-10",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Commemorative One Hundred Dollar Precious Metal Coin (Sainte-Marie, Ontario—350th Anniversary—1639-1989)",
      "citation": "SOR/89-10",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-76",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Fifty Cent, a Twenty-five Cent, a Ten Cent and a Five Cent Precious Metal Coins",
      "citation": "SOR/96-76",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-142",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Fifty Dollar Precious Metal Coin",
      "citation": "SOR/98-142",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-570",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Five Dollar Precious Metal Coin",
      "citation": "SOR/97-570",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-513",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Five Dollar Precious Metal Coin",
      "citation": "SOR/98-513",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-391",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Cent Base Metal Coin",
      "citation": "SOR/96-391",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-51",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Cent Base Metal Coin",
      "citation": "SOR/97-51",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-512",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Cent Precious Metal Coin",
      "citation": "SOR/98-512",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-99",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Dollar Base Metal Coin",
      "citation": "SOR/97-99",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-144",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Dollar Precious Metal Coin",
      "citation": "SOR/94-144",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-751",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Dollar Precious Metal Coin",
      "citation": "SOR/92-751",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-552",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Dollar Precious Metal Coin",
      "citation": "SOR/96-552",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-69",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Dollar Precious Metal Coin",
      "citation": "SOR/96-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-241",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Dollar Precious Metal Coin",
      "citation": "SOR/97-241",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-485",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Dollar Precious Metal Coin",
      "citation": "SOR/97-485",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-47",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Dollar Precious Metal Coin",
      "citation": "SOR/95-47",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-195",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Hundred Dollar Precious Metal Coin",
      "citation": "SOR/94-195",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-749",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Hundred Dollar Precious Metal Coin",
      "citation": "SOR/92-749",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-455",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Hundred Dollar Precious Metal Coin",
      "citation": "SOR/96-455",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-71",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Hundred Dollar Precious Metal Coin",
      "citation": "SOR/96-71",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-46",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Hundred Dollar Precious Metal Coin",
      "citation": "SOR/95-46",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-44",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Hundred Dollar Precious Metal Coin",
      "citation": "SOR/99-44",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-111",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Hundred Dollar Precious Metal Coin",
      "citation": "SOR/98-111",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-90",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a One Hundred Dollar Precious Metal Coin (International Year of Literacy)",
      "citation": "SOR/90-90",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-360",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Ten Cent Precious Metal Coin",
      "citation": "SOR/97-360",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-93",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Three Hundred and Fifty Dollar Precious Metal Coin",
      "citation": "SOR/98-93",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-87",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Three Hundred and Fifty Dollar Precious Metal Coin",
      "citation": "SOR/99-87",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-489",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Two Dollar Base Metal Coin",
      "citation": "SOR/95-489",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-489",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Two Dollar Precious Metal Coin",
      "citation": "SOR/96-489",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-105",
      "title": "Issue and Prescribing the Composition, Dimensions and Design of a Two Dollar Precious Metal Coin",
      "citation": "SOR/96-105",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-196",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Two Hundred Dollar Precious Metal Coin",
      "citation": "SOR/94-196",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-265",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Two Hundred Dollar Precious Metal Coin",
      "citation": "SOR/93-265",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-456",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Two Hundred Dollar Precious Metal Coin",
      "citation": "SOR/96-456",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-70",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Two Hundred Dollar Precious Metal Coin",
      "citation": "SOR/96-70",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-94",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Two Hundred Dollar Precious Metal Coin",
      "citation": "SOR/95-94",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-42",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Two Hundred Dollar Precious Metal Coin",
      "citation": "SOR/99-42",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-112",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Design of a Two Hundred Dollar Precious Metal Coin",
      "citation": "SOR/98-112",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-93",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of a Three Hundred, a One Hundred and Fifty, a Seventy-five and a Thirty Dollar Precious Metal Coins",
      "citation": "SOR/95-93",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-184",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of a Three Hundred, a One Hundred and Fifty, a Seventy-five and a Thirty Dollar Precious Metal Coins",
      "citation": "SOR/99-184",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-143",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of a Three Hundred Dollar, a One Hundred and Fifty Dollar, a Seventy-five Dollar and a Thirty Dollar Precious Metal Coin",
      "citation": "SOR/94-143",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-500",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of a Three Hundred Dollar, a One Hundred and Fifty Dollar, a Seventy-five Dollar and a Thirty Dollar Precious Metal Coin",
      "citation": "SOR/97-500",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-73",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of a Three Hundred Dollar, a One Hundred and Fifty Dollar, a Seventy-five Dollar and a Thirty Dollar Precious Metal Coins",
      "citation": "SOR/96-73",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-98",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of a Three Hundred Dollar, a One Hundred and Fifty Dollar, a Seventy-five Dollar and a Thirty Dollar Precious Metal Coins",
      "citation": "SOR/97-98",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-750",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of a Twenty Dollar Gold and Silver Precious Metal Coins",
      "citation": "SOR/92-750",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-482",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of a Twenty-five Cent Precious Metal Coin and a Twenty-five Cent Base Metal Coin",
      "citation": "SOR/92-482",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-483",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of a Twenty-five Cent Precious Metal Coin and a Twenty-five Cent Base Metal Coin",
      "citation": "SOR/92-483",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-576",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of a Twenty-five Cent Precious Metal Coin and a Twenty-five Cent Base Metal Coin",
      "citation": "SOR/92-576",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-641",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of a Twenty-five Cent Precious Metal Coin and a Twenty-five Cent Base Metal Coin",
      "citation": "SOR/92-641",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-643",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of a Twenty-five Cent Precious Metal Coin and a Twenty-five Cent Base Metal Coin",
      "citation": "SOR/92-643",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-540",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of a Two Hundred Dollar Precious Metal Coin",
      "citation": "SOR/92-540",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-411",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Certain Coins",
      "citation": "SOR/88-411",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-481",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Certain Precious Metal Coins",
      "citation": "SOR/92-481",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-106",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Certain Precious Metal Coins",
      "citation": "SOR/93-106",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-230",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Certain Precious Metal Coins",
      "citation": "SOR/93-230",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-476",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Certain Precious Metal Coins",
      "citation": "SOR/90-476",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-118",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Commemorative One Hundred and One Dollar Precious Metal Coins (100th and 175th Anniversaries of the Launch of the First C.P. Steamship and the Launch of the First Steamboat to operate on Lake Ontario)",
      "citation": "SOR/91-118",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-95",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Fifteen Dollar Precious Metal Coins",
      "citation": "SOR/98-95",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-49",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Four Fifty-cent Precious Metal Coins",
      "citation": "SOR/95-49",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-74",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Four Fifty Cent Precious Metal Coins",
      "citation": "SOR/96-74",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-477",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Four Fifty Cent Precious Metal Coins",
      "citation": "SOR/97-477",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-183",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Four Fifty Cent Precious Metal Coins",
      "citation": "SOR/99-183",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-88",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Four Fifty Cent Precious Metal Coins",
      "citation": "SOR/99-88",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-97",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Four Fifty Cent Precious Metal Coins",
      "citation": "SOR/97-97",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-243",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Four Precious Metal Fifty Cent Coins",
      "citation": "SOR/98-243",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-43",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Two One Dollar Precious Metal Coins",
      "citation": "SOR/99-43",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-145",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Two Twenty Dollar Precious Metal Coins",
      "citation": "SOR/94-145",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-490",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Two Twenty Dollar Precious Metal Coins",
      "citation": "SOR/96-490",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-72",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Two Twenty Dollar Precious Metal Coins",
      "citation": "SOR/96-72",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-450",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Two Twenty Dollar Precious Metal Coins",
      "citation": "SOR/95-450",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-193",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Two Twenty Dollar Precious Metal Coins",
      "citation": "SOR/98-193",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-89",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Two Twenty Dollars Precious Metal Coins",
      "citation": "SOR/99-89",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-401",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Various Fifty Dollar Precious Metal Coins",
      "citation": "SOR/97-401",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-413",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Various Precious Metal Coins",
      "citation": "SOR/98-413",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-41",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Various Twenty-five Cent Metal Base Coins",
      "citation": "SOR/99-41",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-40",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, Dimensions and Designs of Various Twenty-five Cent Precious Metal Coins",
      "citation": "SOR/99-40",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-416",
      "title": "Proclamation Authorizing the Issue and Prescribing the Composition, the Dimensions and Designs of Two Twenty Dollar Precious Metal Coins",
      "citation": "SOR/92-416",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-200",
      "title": "Order Authorizing the Issue and Prescribing the Design and Dimensions of a Two Dollar Circulation Coin",
      "citation": "SOR/99-200",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-245",
      "title": "Order Authorizing the Issue of a Two Dollar Circulation Coin Commemorating the Millenium and Specifying its Characteristics",
      "citation": "SOR/2000-245",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-162",
      "title": "Order Authorizing the Issue of Circulation Coins of Fifty Cents, Twenty-five Cents, Ten Cents and Five Cents and Specifying Their Characteristics",
      "citation": "SOR/2000-162",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-178",
      "title": "Order Authorizing the Issue of Circulation Coins of One Dollar and Specifying Their Characteristics",
      "citation": "SOR/2007-178",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-218",
      "title": "Order Authorizing the Issue of Nineteen Non- circulation Coins",
      "citation": "SOR/99-218",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-119",
      "title": "Order Authorizing the Issue of Non-circulation Coins of the Denomination of Five Hundred Dollars",
      "citation": "SI/2006-119",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-108",
      "title": "Order Authorizing the Issue of Non-circulation Coins of the Denomination of One Hundred and Twenty-five Dollars",
      "citation": "SI/2005-108",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2007-16",
      "title": "Order Authorizing the Issue of Non-circulation Coins of the Denomination of One Million Dollars",
      "citation": "SI/2007-16",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-3",
      "title": "Order Authorizing the Issue of Non-circulation Coins of the Denomination of Three Hundred and Fifty Dollars",
      "citation": "SI/2006-3",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-120",
      "title": "Order Authorizing the Issue of Non-circulation Coins of the Denomination of Three Hundred and Fifty Dollars",
      "citation": "SI/2006-120",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-107",
      "title": "Order Authorizing the Issue of Non-circulation Coins of the Denomination of Twenty-five Dollars",
      "citation": "SI/2005-107",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-109",
      "title": "Order Authorizing the Issue of Non-circulation Coins of the Denomination of Two Hundred and Fifty Dollars",
      "citation": "SI/2005-109",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-110",
      "title": "Order Authorizing the Issue of Non-circulation Coins of the Denomination of Two Thousand Five Hundred Dollars",
      "citation": "SI/2005-110",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-251",
      "title": "Order Authorizing the Issue of Non-circulation Coins of the Denominations of Three Dollars and Four Dollars",
      "citation": "SOR/2003-251",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-21",
      "title": "Order Authorizing the Issue of Twenty-five Cent Circulation Coins and Specifying Their Characteristics and Determining Their Designs",
      "citation": "SOR/2000-21",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-568",
      "title": "The Jacques-Cartier and Champlain Bridges Inc. Regulations",
      "citation": "SOR/98-568",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-9",
      "title": "Joint Canada-United States Government Projects Remission Order",
      "citation": "SI/91-9",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1185",
      "title": "Joint Use of Poles Regulations",
      "citation": "CRC, c 1185",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-36",
      "title": "Joint Venture (GST/HST) Regulations",
      "citation": "SOR/91-36",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-57-176",
      "title": "Joyceville Institution Proclaimed a Penitentiary for Ontario",
      "citation": "SOR/57-176",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-984",
      "title": "Judges Act (Removal Allowance) Order",
      "citation": "CRC, c 984",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-84-134",
      "title": "Designating the Judges of her Majesty\u0027s Court of Queen\u0027s Bench for Manitoba to Exercise all the Functions and Powers of a Judge Under the Act",
      "citation": "SI/84-134",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-88",
      "title": "Kamloops Airport Zoning Regulations",
      "citation": "CRC, c 88",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-473",
      "title": "Kananaskis Falls and Horseshoe Falls Water Power Regulations",
      "citation": "SOR/97-473",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-77-665",
      "title": "Kananaskis Falls Water Power Regulations",
      "citation": "SOR/77-665",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-100",
      "title": "Kapuskasing Airport Zoning Regulations",
      "citation": "SOR/87-100",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-89",
      "title": "Kelowna Airport Zoning Regulations",
      "citation": "CRC, c 89",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-729",
      "title": "Kemano Completion Project Guidelines Order",
      "citation": "SOR/90-729",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-524",
      "title": "Kenora Airport Zoning Regulations",
      "citation": "SOR/93-524",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-96",
      "title": "Kimmirut Airport Zoning Regulations",
      "citation": "SOR/2012-96",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-22",
      "title": "Kindersley Airport Zoning Regulations",
      "citation": "SOR/92-22",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-32",
      "title": "Kinepouch and Kinestik Explosives Order, 1981",
      "citation": "SI/83-32",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-447",
      "title": "Kingston Airport Zoning Regulations",
      "citation": "SOR/88-447",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-95",
      "title": "Kugaaruk Airport Zoning Regulations",
      "citation": "SOR/2012-95",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-773",
      "title": "L-5 Hydroxytryptophan Remission Order",
      "citation": "CRC, c 773",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-663",
      "title": "Labelle Area Wood Producers\u0027 Levy (Interprovincial and Export Trade) Order",
      "citation": "SOR/87-663",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-134",
      "title": "Labrador Innu Settlements Remission Order, 2003",
      "citation": "SI/2003-134",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1380",
      "title": "Lake Erie and Northern Railway Traffic Rules and Regulations",
      "citation": "CRC, c 1380",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-90",
      "title": "Lakehead Airport Zoning Regulations",
      "citation": "CRC, c 90",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-84-51",
      "title": "Laminated Glass Dinnerware Remission Order",
      "citation": "SI/84-51",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-349",
      "title": "Lancaster Sound Designated Area Regulations",
      "citation": "SOR/98-349",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1021",
      "title": "Lands Surveys Tariff",
      "citation": "CRC, c 1021",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-91",
      "title": "Langley Airport Zoning Regulations",
      "citation": "CRC, c 91",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1435",
      "title": "Large Fishing Vessel Inspection Regulations",
      "citation": "CRC, c 1435",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-281",
      "title": "By-law No. 7 Respecting the Large Value Transfer System",
      "citation": "SOR/2001-281",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-603",
      "title": "La Ronge Airport Zoning Regulations",
      "citation": "SOR/87-603",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-239",
      "title": "LASMO Drilling Rig Remission Order",
      "citation": "SOR/92-239",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-58",
      "title": "Laurentian Pilotage Authority District No. 3 Regulations",
      "citation": "SOR/87-58",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-234",
      "title": "Order Approving the Proposed Regulations Repealing the Laurentian Pilotage Authority District No. 3 Regulations",
      "citation": "SOR/2014-234",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1268",
      "title": "Laurentian Pilotage Authority Regulations",
      "citation": "CRC, c 1268",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-84",
      "title": "Laurentian Pilotage Tariff Regulations",
      "citation": "SOR/2001-84",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-636",
      "title": "Law List Regulations",
      "citation": "SOR/94-636",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-408",
      "title": "Lawrence Bay Airways Ltd. Exemption Order, 1989",
      "citation": "SOR/89-408",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-38",
      "title": "Order Designating the Leader of the Government in the House of Commons as Minister for Purposes of the Act",
      "citation": "SI/2006-38",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-39",
      "title": "Order Designating the Leader of the Government in the House of Commons as Responsible Minister for Purposes of the Electoral Boundaries Readjustment Suspension Act, 1994",
      "citation": "SI/2006-39",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-30",
      "title": "Order Designating the Leader of the Government in the House of Commons as the appropriate Minister with respect to the Canadian Transportation Accident Investigation and Safety Board for the purposes of the Act",
      "citation": "SI/2013-30",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-29",
      "title": "Order Transferring to the Leader of the Government in the House of Commons the powers, duties and functions of the President of the Queen\u0027s Privy Council for Canada under the Canadian Transportation Accident Investigation and Safety Board Act",
      "citation": "SI/2013-29",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-547",
      "title": "Leamy Lake Navigation Channel Regulations",
      "citation": "SOR/98-547",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-162",
      "title": "Least Developed Country Tariff Withdrawal Order (2013 GPT Review)",
      "citation": "SOR/2013-162",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-337",
      "title": "Legal Deposit of Publications Regulations",
      "citation": "SOR/2006-337",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-503",
      "title": "Les Collections Shan Remission Order, 1997",
      "citation": "SOR/97-503",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-92",
      "title": "Lethbridge Airport Zoning Regulations",
      "citation": "CRC, c 92",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-481",
      "title": "Letter Definition Regulations",
      "citation": "SOR/83-481",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-430",
      "title": "Letter Mail Regulations",
      "citation": "SOR/88-430",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-5",
      "title": "Licensed Dealers for Controlled Drugs and Narcotics (Veterinary Use) Fees Regulations",
      "citation": "SOR/98-5",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-5",
      "title": "Licensed Dealers for Controlled Drugs and Narcotics (Veterinary Use) Fees Regulations",
      "citation": "SOR/98-5",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-432",
      "title": "Licensing and Arbitration Regulations",
      "citation": "SOR/84-432",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1022",
      "title": "Lieutenant Governors Superannuation Regulations",
      "citation": "CRC, c 1022",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-277",
      "title": "Life Companies Borrowing Regulations",
      "citation": "SOR/92-277",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1436",
      "title": "Life Saving Equipment Regulations",
      "citation": "CRC, c 1436",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-231",
      "title": "Lighters Regulations",
      "citation": "SOR/2008-231",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-143",
      "title": "Limitation of the Right to Equitable Remuneration of Certain Rome Convention Countries Statement",
      "citation": "SOR/99-143",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1152",
      "title": "Liquefied Petroleum Gases Bulk Storage Regulations",
      "citation": "CRC, c 1152",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-293",
      "title": "Listed Airport Regulations",
      "citation": "SOR/2004-293",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-282",
      "title": "List of Countries to Which Canada Accords Reciprocal Protection Under the Act",
      "citation": "SOR/93-282",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-284",
      "title": "Regulations Establishing a List of Entities",
      "citation": "SOR/2002-284",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-170",
      "title": "Order Establishing a List of Foreign State Supporters of Terrorism",
      "citation": "SOR/2012-170",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-636",
      "title": "List of Hazardous Waste Authorities",
      "citation": "SOR/92-636",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-114",
      "title": "List of Pest Control Product Formulants and Contaminants of Health or Environmental Concern",
      "citation": "SI/2005-114",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-45",
      "title": "List of Wildlife Species at Risk (Decisions not to add Certain Species) Order",
      "citation": "SI/2012-45",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2010-14",
      "title": "List of Wildlife Species at Risk (Decisions Not to Add Certain Species) Order",
      "citation": "SI/2010-14",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-27",
      "title": "List of Wildlife Species at Risk (Decisions Not to Add Certain Species) Order",
      "citation": "SI/2013-27",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-28",
      "title": "List of Wildlife Species at Risk (referral back to COSEWIC) Order",
      "citation": "SI/2013-28",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-103",
      "title": "Little Salmon/Carmacks First Nation (Gst) Remission Order",
      "citation": "SI/2000-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-541",
      "title": "Livestock and Poultry Carcass Grading Regulations",
      "citation": "SOR/92-541",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-182",
      "title": "Livestock Pedigree Associations Winding-up Regulations",
      "citation": "SOR/80-182",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-80-108",
      "title": "Load Line Assignment Authorization Order",
      "citation": "SI/80-108",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-81-105",
      "title": "Load Line Exemption Order",
      "citation": "SI/81-105",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-99",
      "title": "Load Line Regulations",
      "citation": "SOR/2007-99",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1440",
      "title": "Load Line Regulations (Inland)",
      "citation": "CRC, c 1440",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1441",
      "title": "Load Line Regulations (Sea)",
      "citation": "CRC, c 1441",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1442",
      "title": "Load Line Rules for Lakes and Rivers",
      "citation": "CRC, c 1442",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-710",
      "title": "Loaning of Purebred Sires Regulations",
      "citation": "CRC, c 710",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-709",
      "title": "Loan of Defence Materiel to Canadian Contractors Order",
      "citation": "CRC, c 709",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-579",
      "title": "Lobbyists Registration Regulations",
      "citation": "SOR/95-579",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-116",
      "title": "Lobbyists Registration Regulations",
      "citation": "SOR/2008-116",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-152",
      "title": "Locally-Engaged Staff Employment Regulations",
      "citation": "SOR/95-152",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-245",
      "title": "Local Revenue Management Implementation Regulations",
      "citation": "SOR/2007-245",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-254",
      "title": "Local Signal and Distant Signal Regulations",
      "citation": "SOR/89-254",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-93",
      "title": "London Airport Zoning Regulations",
      "citation": "CRC, c 93",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-227",
      "title": "Long-Range Identification and Tracking of Vessels Regulations",
      "citation": "SOR/2010-227",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-206",
      "title": "Losses of Bulk Spirits and Packaged Alcohol Regulations",
      "citation": "SOR/2003-206",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1111",
      "title": "Lower St. Lawrence Region Wood Producers\u0027 Levy (Interprovincial and Export Trade) Order",
      "citation": "SOR/86-1111",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-152",
      "title": "Low Level Air Defence System Remission Order",
      "citation": "SI/87-152",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-797",
      "title": "Machinery Sales Tax Remission Order",
      "citation": "CRC, c 797",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-429",
      "title": "Mackenzie Valley Land Use Regulations",
      "citation": "SOR/98-429",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-32",
      "title": "Mail and Courier Imports (GST/HST) Regulations",
      "citation": "SOR/91-32",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-743",
      "title": "Mail Receptacles Regulations",
      "citation": "SOR/83-743",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-376",
      "title": "Regulations Maintaining Certain Reciprocal Transfer Agreements",
      "citation": "SOR/2000-376",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-791",
      "title": "Maislin Industries Ltd. Regulations",
      "citation": "SOR/82-791",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-351",
      "title": "Management of Contaminated Fisheries Regulations",
      "citation": "SOR/90-351",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-465",
      "title": "Manitoba Chicken Marketing (Interprovincial and Export) Order",
      "citation": "SOR/88-465",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-533",
      "title": "Manitoba Chicken Marketing Levies Order",
      "citation": "SOR/85-533",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-152",
      "title": "Manitoba Chicken Order",
      "citation": "CRC, c 152",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-35",
      "title": "Manitoba Court of Queen\u0027s Bench Rules (Criminal)",
      "citation": "SI/92-35",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-106",
      "title": "Manitoba Criminal Appeal Rules",
      "citation": "SI/92-106",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-515",
      "title": "Manitoba Egg Marketing Administration Levies Order",
      "citation": "SOR/81-515",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-549",
      "title": "Manitoba Egg Order",
      "citation": "SOR/81-549",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-509",
      "title": "Manitoba Fishery Regulations, 1987",
      "citation": "SOR/87-509",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-461",
      "title": "Manitoba Hog Marketing Administration Levies (Interprovincial and Export Trade) Order, 1998",
      "citation": "SOR/98-461",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-154",
      "title": "Manitoba Hog Order",
      "citation": "CRC, c 154",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-530",
      "title": "Manitoba Milk Marketing (Interprovincial and Export) Order",
      "citation": "SOR/90-530",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-315",
      "title": "Manitoba Milk Marketing Levies Order, 1991",
      "citation": "SOR/91-315",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-155",
      "title": "Manitoba Milk Order",
      "citation": "CRC, c 155",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-531",
      "title": "Manitoba Pullet Marketing Administration Levies Order",
      "citation": "SOR/80-531",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-445",
      "title": "Manitoba Pullet Marketing (Interprovincial and Export) Order",
      "citation": "SOR/89-445",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-352",
      "title": "Manitoba Pullet Order",
      "citation": "SOR/80-352",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-595",
      "title": "Manitoba Rules of Practice Respecting Reduction in the Number of Years of Imprisonment Without Eligibility for Parole",
      "citation": "SOR/88-595",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-310",
      "title": "Manitoba Sex Offender Information Registration Regulations",
      "citation": "SOR/2004-310",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-527",
      "title": "Manitoba Turkey Marketing (Interprovincial and Export Trade) Order",
      "citation": "SOR/88-527",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-835",
      "title": "Manitoba Turkey Marketing Levies Order",
      "citation": "SOR/84-835",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-160",
      "title": "Manitoba Turkey Order",
      "citation": "CRC, c 160",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-19",
      "title": "Manitoba Vegetable Marketing (Interprovincial and Export) Order",
      "citation": "SOR/89-19",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-544",
      "title": "Manitoba Vegetable Order",
      "citation": "SOR/85-544",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-399",
      "title": "Manner of Calculation (Foreign Banks) Regulations",
      "citation": "SOR/2001-399",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-145",
      "title": "Manner of Disposal of Detained, Seized or Forfeited Goods Regulations (Preclearance Act)",
      "citation": "SOR/2002-145",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-98",
      "title": "Regulations Respecting the Manner of Selection and Term of Office of the Members of the Coal Mining Safety Commission",
      "citation": "SOR/90-98",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-289",
      "title": "Maple Products Regulations",
      "citation": "CRC, c 289",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-36",
      "title": "Proclamation Designating the Maple Tree as National Arboreal Emblem of Canada",
      "citation": "SI/96-36",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-261",
      "title": "Marihuana Exemption (Food and Drugs Act) Regulations",
      "citation": "SOR/2003-261",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-120",
      "title": "Marihuana Exemption (Food and Drugs Act) Regulations",
      "citation": "SOR/2013-120",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-119",
      "title": "Marihuana for Medical Purposes Regulations",
      "citation": "SOR/2013-119",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-227",
      "title": "Marihuana Medical Access Regulations",
      "citation": "SOR/2001-227",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-76",
      "title": "Marine Activities in the Saguenay-St. Lawrence Marine Park Regulations",
      "citation": "SOR/2002-76",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-391",
      "title": "Marine Certification Regulations",
      "citation": "SOR/97-391",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1127",
      "title": "Marine Industries Limited Regulations",
      "citation": "SOR/86-1127",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-307",
      "title": "Marine Liability Regulations",
      "citation": "SOR/2002-307",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-264",
      "title": "Marine Machinery Regulations",
      "citation": "SOR/90-264",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-56",
      "title": "Marine Mammal Regulations",
      "citation": "SOR/93-56",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-183",
      "title": "Marine Occupational Safety and Health Regulations",
      "citation": "SOR/87-183",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-115",
      "title": "Marine Personnel Regulations",
      "citation": "SOR/2007-115",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-10",
      "title": "Marine Spark-Ignition Engine, Vessel and Off-Road Recreational Vehicle Emission Regulations",
      "citation": "SOR/2011-10",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-144",
      "title": "Marine Transportation Security Regulations",
      "citation": "SOR/2004-144",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-120",
      "title": "Maritime Occupational Health and Safety Regulations",
      "citation": "SOR/2010-120",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-55",
      "title": "Maritime Provinces Fishery Regulations",
      "citation": "SOR/93-55",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-202",
      "title": "Marketing Authorization for Food Additives that may be Used as Anticaking Agents",
      "citation": "SOR/2012-202",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-203",
      "title": "Marketing Authorization for Food Additives that may be Used as Bleaching, Maturing or Dough Conditioning Agents",
      "citation": "SOR/2012-203",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-216",
      "title": "Marketing Authorization for Food Additives that may be Used as Carrier or Extraction Solvents",
      "citation": "SOR/2012-216",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-204",
      "title": "Marketing Authorization for Food Additives that may be Used as Colouring Agents",
      "citation": "SOR/2012-204",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-205",
      "title": "Marketing Authorization for Food Additives that may be Used as Emulsifying, Gelling, Stabilizing or Thickening Agents",
      "citation": "SOR/2012-205",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-207",
      "title": "Marketing Authorization for Food Additives That may be Used as Firming Agents",
      "citation": "SOR/2012-207",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-206",
      "title": "Marketing Authorization for Food Additives that may be Used as Food Enzymes",
      "citation": "SOR/2012-206",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-208",
      "title": "Marketing Authorization for Food Additives that may be Used as Glazing or Polishing Agents",
      "citation": "SOR/2012-208",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-211",
      "title": "Marketing Authorization for Food Additives that may be Used as pH Adjusting Agents, Acid-reacting Materials or Water Correcting Agents",
      "citation": "SOR/2012-211",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-212",
      "title": "Marketing Authorization for Food Additives that may be Used as Preservatives",
      "citation": "SOR/2012-212",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-213",
      "title": "Marketing Authorization for Food Additives that may be Used as Sequestering Agents",
      "citation": "SOR/2012-213",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-214",
      "title": "Marketing Authorization for Food Additives that may be Used as Starch-modifying Agents",
      "citation": "SOR/2012-214",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-210",
      "title": "Marketing Authorization for Food Additives that may be Used as Sweeteners",
      "citation": "SOR/2012-210",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-215",
      "title": "Marketing Authorization for Food Additives that may be Used as Yeast Foods",
      "citation": "SOR/2012-215",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-209",
      "title": "Marketing Authorization for Food Additives with Other Generally Accepted Uses",
      "citation": "SOR/2012-209",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-87",
      "title": "Marketing Authorization for Maximum Residue Limits for Veterinary Drugs in Foods",
      "citation": "SOR/2013-87",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-535",
      "title": "Marking of Imported Goods Order",
      "citation": "CRC, c 535",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-10",
      "title": "Marking of Imported Goods Regulations",
      "citation": "SOR/94-10",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-261",
      "title": "Masked Name Regulations",
      "citation": "SOR/94-261",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-163",
      "title": "Material Banking Group Percentage Regulations",
      "citation": "SOR/2008-163",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-400",
      "title": "Material Percentage Regulations",
      "citation": "SOR/2001-400",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1283",
      "title": "Materials for the Use of the Blind Regulations",
      "citation": "CRC, c 1283",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-128",
      "title": "2014 Summit on Maternal, Newborn and Child Health — Privileges and Immunities Order",
      "citation": "SOR/2014-128",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-15",
      "title": "Maximum Amount for Agreements with Governments Order",
      "citation": "SOR/99-15",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-295",
      "title": "Maximum Amounts for Informal Procedure Regulations",
      "citation": "SOR/93-295",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-474",
      "title": "Order Prescribing the Maximum Total Amount for Agreements Involving Leases of Residential Property",
      "citation": "SOR/98-474",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-346",
      "title": "Mayo Airport Zoning Regulations",
      "citation": "SOR/95-346",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-128",
      "title": "McIntyre Lands Income Tax Remission Order",
      "citation": "SI/2005-128",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-55",
      "title": "Order Imposing Measures to Address the Extraordinary Disruption to the National Transportation System in Relation to Grain Movement",
      "citation": "SOR/2014-55",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-818",
      "title": "Meat and Poultry Products Plant Liquid Effluent Regulations",
      "citation": "CRC, c 818",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-96",
      "title": "Meat Inspection Overtime Inspection Fees Remission Order",
      "citation": "SI/83-96",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-288",
      "title": "Meat Inspection Regulations, 1990",
      "citation": "SOR/90-288",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-24",
      "title": "Medical Devices (GST) Regulations",
      "citation": "SOR/91-24",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-282",
      "title": "Medical Devices Regulations",
      "citation": "SOR/98-282",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-538",
      "title": "Medical High Technology International Inc. Shares Acquisition Order",
      "citation": "SOR/89-538",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-316",
      "title": "Meeting and Proposals (Insurance Companies and Insurance Holding Companies) Regulations",
      "citation": "SOR/2006-316",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-123",
      "title": "19th Meeting of the Parties to the Montreal Protocol Privileges and Immunities Order",
      "citation": "SOR/2007-123",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-314",
      "title": "Meetings and Proposals (Banks and Bank Holding Companies) Regulations",
      "citation": "SOR/2006-314",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-315",
      "title": "Meetings and Proposals (Cooperative Credit Associations) Regulations",
      "citation": "SOR/2006-315",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-317",
      "title": "Meetings and Proposals (Trust and Loan Companies) Regulations",
      "citation": "SOR/2006-317",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-225",
      "title": "Members of Committees and Special Committees (NAFTA) Regulations",
      "citation": "SOR/94-225",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-71",
      "title": "Members of Committees Regulations",
      "citation": "SOR/89-71",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-117",
      "title": "Members of Panels (NAFTA) Regulations",
      "citation": "SOR/94-117",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-70",
      "title": "Members of Panels Regulations",
      "citation": "SOR/89-70",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-373",
      "title": "Members of Parliament Disability Allowance Regulations",
      "citation": "SOR/2002-373",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1033",
      "title": "Members of Parliament Retiring Allowances Regulations",
      "citation": "CRC, c 1033",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-33",
      "title": "Order Authorizing Members of the Canadian Armed Forces to Accept and Wear the United Nations Iraq/Kuwait Observer Mission (UNIKOM) Medal",
      "citation": "SI/92-33",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-103",
      "title": "Order Authorizing Members of the Canadian Armed Forces to Accept and Wear the United Nations Observer Group in Central America (ONUCA) Medal",
      "citation": "SI/90-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-145",
      "title": "Order Authorizing Members of the Canadian Armed Forces to Accept and Wear the United Nations Protection Force in Yugoslavia (UNPROFOR) Medal",
      "citation": "SI/92-145",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-146",
      "title": "Order Authorizing Members of the Canadian Armed Forces to Accept and Wear the United Nations Transitional Authority in Cambodia (UNTAC) Medal",
      "citation": "SI/92-146",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-146",
      "title": "Order Placing Members of the Canadian Force on Active Service (Western Sahara)",
      "citation": "SI/91-146",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-95",
      "title": "Order Authorizing Members of the Canadian Forces and Canadian Civil Police to Accept and Wear the United Nations Advance Mission in Cambodia (UNAMIC) Medal",
      "citation": "SI/93-95",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-94",
      "title": "Order Authorizing Members of the Canadian Forces and Canadian Civil Police to Accept and Wear the United Nations Protection Force (UNPROFOR) Medal (Yugoslavia)",
      "citation": "SI/93-94",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-96",
      "title": "Order Authorizing Members of the Canadian Forces and Canadian Civil Police to Accept and Wear the United Nations Transitional Authority in Cambodia (UNTAC) Medal",
      "citation": "SI/93-96",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-92",
      "title": "Order Authorizing Members of the Canadian Forces and the Royal Canadian Mounted Police to Accept and Wear the United Nations Transition Assistance Group (UNTAG) Medal",
      "citation": "SI/90-92",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-111",
      "title": "Order Placing Members of the Canadian Forces on Active Service (Arabian Peninsula)",
      "citation": "SI/90-111",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-43",
      "title": "Order Placing Members of the Canadian Forces on Active Service (Cambodia)",
      "citation": "SI/92-43",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-103",
      "title": "Order Placing Members of the Canadian Forces on Active Service for the Purpose of Fulfilling Canada\u0027s Obligations Under the North Atlantic Treaty",
      "citation": "SI/89-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-129",
      "title": "Order Placing Members of the Canadian Forces on Active Service (Iran-Iraq)",
      "citation": "SI/88-129",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-15",
      "title": "Order Placing Members of the Canadian Forces on Active Service (ONUCA—Central America)",
      "citation": "SI/90-15",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-222",
      "title": "Order No. 2 Placing Members of the Canadian Forces on Active Service (Somalia)",
      "citation": "SI/92-222",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-169",
      "title": "Order Placing Members of the Canadian Forces on Active Service (Somalia)",
      "citation": "SI/92-169",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-104",
      "title": "Order Placing Members of the Canadian Forces on Active Service (UNTAG-Namibia)",
      "citation": "SI/89-104",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-42",
      "title": "Order Placing Members of the Canadian Forces on Active Service (Yugoslavia)",
      "citation": "SI/92-42",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-152",
      "title": "Order Authorizing Members of the Canadian Forces to Accept and Wear the United Nations Iran-Iraq Military Observer Group (UNIIMOG) Medal",
      "citation": "SI/89-152",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1622",
      "title": "Memorial Cross Order (World War I)",
      "citation": "CRC, c 1622",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1623",
      "title": "Memorial Cross Order (World War II)",
      "citation": "CRC, c 1623",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-219",
      "title": "Merchandise for Photographic Layouts Remission Order",
      "citation": "SI/85-219",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-520",
      "title": "Merchant Seamen Compensation Order, 1992",
      "citation": "SOR/92-520",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-222",
      "title": "Metal Mining Effluent Regulations",
      "citation": "SOR/2002-222",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-155",
      "title": "Metering Assemblies Incorporating Electronic ATCs Specifications",
      "citation": "SI/90-155",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-972",
      "title": "Metric Commission Order",
      "citation": "CRC, c 972",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-37",
      "title": "Mexico Deemed Direct Shipment (General Preferential Tariff) Regulations",
      "citation": "SOR/98-37",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-592",
      "title": "Mexico Fruit and Vegetable Aggregate Quantity Limit Order",
      "citation": "SOR/93-592",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-533",
      "title": "Miawpukek Band Order",
      "citation": "SOR/89-533",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-204",
      "title": "The Micronutrient Initiative Divestiture Regulations",
      "citation": "SOR/2009-204",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-226",
      "title": "MicroSD Cards Exclusion Regulations (Copyright Act)",
      "citation": "SOR/2012-226",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1036",
      "title": "Migratory Bird Sanctuary Regulations",
      "citation": "CRC, c 1036",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1035",
      "title": "Migratory Birds Regulations",
      "citation": "CRC, c 1035",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-14",
      "title": "Military Police Professional Code of Conduct",
      "citation": "SOR/2000-14",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1049",
      "title": "Military Rules of Evidence",
      "citation": "CRC, c 1049",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-301",
      "title": "Mingan Archipelago National Park Reserve of Canada Snowshoe Hare Regulations",
      "citation": "SOR/2004-301",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-189",
      "title": "Order Specifying the Minimum Amount of Grain to be Moved",
      "citation": "SOR/2014-189",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-276",
      "title": "Minimum Amount of Grain to Be Moved, No. 2, Order Specifying the",
      "citation": "SOR/2014-276",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-104",
      "title": "Mining Near Lines of Railways Regulations",
      "citation": "SOR/91-104",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-427",
      "title": "Minister Designation Order (Canada Business Corporations Act)",
      "citation": "CRC, c 427",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1241",
      "title": "Minister Designation Order (Nuclear Liability Act)",
      "citation": "CRC, c 1241",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1302",
      "title": "Minister Designation Order (Precious Metals Marking Act)",
      "citation": "CRC, c 1302",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-8",
      "title": "Order Designating the Minister for International Trade as Minister for Purposes of Sections 1 to 9 and Parts I and III of the Act",
      "citation": "SI/94-8",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-228",
      "title": "Order Designating the Minister for Purposes of the Act",
      "citation": "SI/93-228",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-228",
      "title": "Order Minister for Purposes of the Act",
      "citation": "SI/93-228",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-631",
      "title": "Ministerial Regulations Authorization Order",
      "citation": "SOR/86-631",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-44",
      "title": "Order Transferring From the Minister of Agriculture and Agri-Food to the Minister of National Health and Welfare, the Powers, Duties and Functions Under Certain Acts",
      "citation": "SI/95-44",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-138",
      "title": "Minister of Agriculture Authority to Prescribe Fees Order",
      "citation": "SI/86-138",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-140",
      "title": "Minister of Agriculture Authority to Prescribe Fees Order",
      "citation": "SI/86-140",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-139",
      "title": "Minister of Agriculture Authority to Prescribe Fees Order",
      "citation": "SI/86-139",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-712",
      "title": "Minister of Agriculture Authority to Prescribe Fees Order",
      "citation": "CRC, c 712",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-147",
      "title": "Order Authorizing the Minister of Agriculture to Prescribe Fees",
      "citation": "SI/94-147",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-147",
      "title": "Order Authorizing the Minister of Agriculture to Prescribe Fees",
      "citation": "SI/92-147",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-124",
      "title": "Order Designating the Minister of Canadian Heritage for Purposes of Section 23 of the Act",
      "citation": "SI/2005-124",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-125",
      "title": "Order Designating the Minister of Canadian Heritage for Purposes of Section 110 of the Act",
      "citation": "SI/2005-125",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-73",
      "title": "Order Designating the Minister of Canadian Heritage to be the Minister referred to in the Act",
      "citation": "SI/2014-73",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-85",
      "title": "Order Transferring From the Minister of Canadian Heritage to the Minister of Indian Affairs and Northern Development the Control and Supervision of the Office of Indian Residential Schools Resolution of Canada and Ordering the Minister of Indian Affairs and Northern Development to Preside Over the Office",
      "citation": "SI/2006-85",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-226",
      "title": "Order Transferring From the Minister of Canadian Heritage to the Minister of the Environment the Control and Supervision of the Parks Canada Agency",
      "citation": "SI/2003-226",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-120",
      "title": "Order Designating the Minister of Citizenship and Immigration as the Minister Responsible for the Administration of That Act",
      "citation": "SI/2001-120",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-214",
      "title": "Order Transferring From the Minister of Citizenship and Immigration to the Deputy Prime Minister and Minister of Public Safety and Emergency Preparedness the Control and Supervision of the Canada Border Services Agency",
      "citation": "SI/2003-214",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-200",
      "title": "Order Transferring to the Minister of Communications and to the Minister of Multiculturalism and Citizenship Certain Powers, Duties and Functions of the Secretary of State of Canada and De-amalgamating the Departments of the Secretary of State of Canada, Multiculturalism and Citizenship and Communications",
      "citation": "SI/93-200",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-536",
      "title": "Order Designating the Minister of Communications as Appropriate Minister With Respect to the National Archives, the National Film Board, the National Library, the National Gallery, the Canadian Museum of Nature, the Canadian Museum of Civilization and the National Museum of Science and Technology",
      "citation": "SOR/93-536",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-231",
      "title": "Order Designating the Minister of Communications as Appropriate Minister With Respect to the National Capital Commission",
      "citation": "SI/93-231",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-232",
      "title": "Order Designating the Minister of Communications as Minister for Purposes of Section 47 of the Act",
      "citation": "SI/93-232",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-230",
      "title": "Order Designating the Minister of Communications as Minister for Purposes of the act",
      "citation": "SI/93-230",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-79-10",
      "title": "Minister of Communications Authority to Prescribe Fees Order",
      "citation": "SI/79-10",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-82-104",
      "title": "Minister of Communications Authority to Prescribe Fees Order",
      "citation": "SI/82-104",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-81-19",
      "title": "Minister of Communications Authority to Prescribe Fees Order",
      "citation": "SI/81-19",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-32",
      "title": "Minister of Communications Authority to Prescribe Fees Order",
      "citation": "SI/93-32",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-106",
      "title": "Minister of Communications (National Archives of Canada) Authority to Prescribe Fees Order",
      "citation": "SI/94-106",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-87",
      "title": "Minister of Communications (National Library) Authority to Prescribe Fees Order",
      "citation": "SI/94-87",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-93",
      "title": "Minister of Consumer and Corporate Affairs Authority to Prescribe Fees or Charges Order",
      "citation": "SI/87-93",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-84-35",
      "title": "Minister of Consumer and Corporate Affairs Authority to Prescribe Fees Order",
      "citation": "SI/84-35",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-84-38",
      "title": "Minister of Consumer and Corporate Affairs Authority to Prescribe Fees Order",
      "citation": "SI/84-38",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-46",
      "title": "Minister of Consumer and Corporate Affairs Authority to Prescribe Fees Order",
      "citation": "SI/88-46",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-47",
      "title": "Minister of Consumer and Corporate Affairs Authority to Prescribe Fees Order",
      "citation": "SI/88-47",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-130",
      "title": "Minister of Employment and Immigration Authority to Prescribe Fees or Charges Order",
      "citation": "SI/88-130",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-36",
      "title": "Order Transferring From the Minister of Employment and Immigration to the Minister of Labour, all the Powers, Duties and Functions of the Minister of Employment and Immigration Under Certain Acts",
      "citation": "SI/95-36",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-144",
      "title": "Order Transferring to the Minister of Energy, Mines and Resources the Powers, Duties and Functions of the Minister of Forestry and Amalgamating and Combining the Department of Energy, Mines and Resources and the Department of Forestry Under the Minister of Energy, Mines and Resources",
      "citation": "SI/93-144",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-2",
      "title": "Order Authorizing the Minister of Energy, Mines and Resources to Prescribe Fees",
      "citation": "SI/95-2",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-98",
      "title": "Order Designating the Minister of Finance as Minister for Purposes of the Act",
      "citation": "SI/89-98",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-69",
      "title": "Order Transferring to the Minister of Finance the Powers, Duties and Functions of the Minister of Indian Affairs and Northern Development in Respect of the Administration of the Transfer Payments to the Territorial Governments Program",
      "citation": "SI/95-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-122",
      "title": "Order Transferring to the Minister of Finance the Powers, Duties and Functions of the Minister of Supply and Services Under Part I of the Currency Act",
      "citation": "SI/88-122",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-187",
      "title": "Order Authorizing the Minister of Finance to Exercise the Powers of the Minister of Fisheries and Oceans Referred to in the Act in Respect of National Sea Products Limited",
      "citation": "SOR/89-187",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-16",
      "title": "Order Transferring From the Minister of Finance to the Minister of Industry, the Powers, Duties and Functions in Relation to Regional Economic Development in Quebec and the Control and Supervision of the Economic Development Agency of Canada for the Regions of Quebec",
      "citation": "SI/96-16",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-52",
      "title": "Order Transferring From the Minister of Finance, to the Minister of National Health and Welfare and to the Minister of Employment and Immigration Certain Powers, Duties and Functions With Respect to Certain Programs",
      "citation": "SI/95-52",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-41",
      "title": "Minister of Fisheries and Oceans Authority to Prescribe Fees or Charges Orders",
      "citation": "SI/88-41",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-48",
      "title": "Order Designating the Minister of Foreign Affairs for the Purposes of the National Capital Act",
      "citation": "SI/2011-48",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-208",
      "title": "Order Transferring From the Minister of Foreign Affairs to the Minister of International Trade the Control and Supervision of the Department of International Trade",
      "citation": "SI/2003-208",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-110",
      "title": "Order Repealing Order in Council P.C. 2008-12 and Designating the Minister of Health as the appropriate Minister with Respect to the Canadian Food Inspection Agency for the Purposes of the Act",
      "citation": "SI/2013-110",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-41",
      "title": "Order Transferring to the Minister of Health the powers, duties and functions of the Minister of Agriculture and Agri-Food and from the Canadian Food Inspection Agency to the Public Health Agency of Canada the control and supervision of that portion of the Federal Public Administration known as the Domestic Terrestrial Animal Pathogen Unit",
      "citation": "SI/2013-41",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-90",
      "title": "Order Transferring from the Minister of Health to the Minister of the Canadian Northern Economic Development Agency the Control and Supervision of the portion of the federal public administration known as the Canadian Northern Economic Development Agency and Ordering the Minister of the Canadian Northern Economic Development Agency to Preside over the Agency",
      "citation": "SI/2013-90",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-78-65",
      "title": "Minister of Indian Affairs and Northern Development Authority to Prescribe Fees Order",
      "citation": "SI/78-65",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-713",
      "title": "Minister of Indian Affairs and Northern Development Authority to Prescribe Fees Order",
      "citation": "CRC, c 713",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-87",
      "title": "Order Authorizing the Minister of Indian Affairs and Northern Development to Act as Federal Interlocutor for Métis and Non-status Indians",
      "citation": "SI/2004-87",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-94",
      "title": "Order Transferring From the Minister of Indian Affairs and Northern Development to the Minister of Health the Control and Supervision of the Canadian Northern Economic Development Agency and Ordering the Minister of Health to Preside over the Agency",
      "citation": "SI/2011-94",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-32",
      "title": "Order Designating the Minister of Industry for the Purposes of the Act, as the Minister in Relation to Small Businesses in Quebec, the Minister for the Purposes of the Atlantic Canada Opportunities Agency Act in Relation to Small Businesses in the Four Maritime Provinces and the Minister of Western Diversification in Relation to Small Businesses in the Four Western Provinces",
      "citation": "SI/99-32",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-60",
      "title": "Order Designating the Minister of Industry for the Purposes of the Canada Not-for-profit Corporations Act",
      "citation": "SI/2011-60",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-234",
      "title": "Order Designating the Minister of Industry, Science and Technology as Minister for Purposes of the Natural Sciences and Engineering Research Council Act and as Appropriate Minister With Respect to the Natural Sciences and Engineering Research Council for Purposes of the Financial Administration Act",
      "citation": "SI/93-234",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-18",
      "title": "Order Designating the Minister of Industry, Science and Technology as Minister for Purposes of the Social Sciences and Humanities Research Council Act and as Appropriate Minister With Respect to the Social Sciences and Humanities Research Council for Purposes of the Financial Administration Act",
      "citation": "SI/94-18",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-238",
      "title": "Order Transferring From the Minister of Industry, Science and Technology to the Minister of National Health and Welfare the Powers, Duties and Functions of the Minister of Industry, Science and Technology under the Hazardous Materials Information Review Act",
      "citation": "SI/93-238",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-227",
      "title": "Order Transferring From the Minister of Industry to the Minister of Environment the Control and Supervision of the Office of Infrastructure of Canada",
      "citation": "SI/2003-227",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-49",
      "title": "Order Transferring From the Minister of Industry to the Minister of Finance, the Powers, Duties and Functions under Certain Sections of the Act",
      "citation": "SI/95-49",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-101",
      "title": "Order Transferring From the Minister of Industry to the Minister of the Economic Development Agency of Canada for the Regions of Quebec and Minister Responsible for the Francophonie the Control and Supervision of the Agency and Ordering the Minister to Preside Over the Agency",
      "citation": "SI/2004-101",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-16",
      "title": "Order Transferring From the Minister of Industry to the Registrar of Lobbyists the Control and Supervision of the Office of the Registrar of Lobbyists",
      "citation": "SI/2006-16",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-108",
      "title": "Designating the Minister of Justice and the President of the Treasury Board as Ministers for Purposes of Certain Sections of the Act",
      "citation": "SI/83-108",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-98",
      "title": "Order Transferring From the Minister of National Health and Welfare to the Minister of Labour, the Powers, Duties and Functions in Relation to Regional Economic Development in Quebec and the Control and Supervision of the Federal Office of Regional Development—Quebec",
      "citation": "SI/93-98",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-8",
      "title": "Minister of National Revenue Authority to Prescribe Fees Order",
      "citation": "SI/93-8",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-188",
      "title": "Minister of National Revenue Authority to Prescribe Fees Order",
      "citation": "SI/88-188",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-189",
      "title": "Minister of National Revenue Authority to Prescribe Fees Order",
      "citation": "SI/88-189",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-102",
      "title": "Order Transferring to the Minister of National Revenue, the Powers, Duties and Functions of the Minister of Industry under the Tax Rebate Discounting Act",
      "citation": "SI/95-102",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-27",
      "title": "Order Transferring From the Minister of National Revenue to the Minister of Transport the Control and Supervision of the Royal Canadian Mint",
      "citation": "SI/2006-27",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-47",
      "title": "Order Tranferring From the Minister of Public Safety and Emergency Preparedness to the Minister of Canadian Heritage the Control and Supervision of the Office of Indian Residential Schools Resolution of Canada",
      "citation": "SI/2006-47",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-58",
      "title": "Order Designating the Minister of Public Works and Government Services for the purposes of the Act",
      "citation": "SI/2012-58",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-228",
      "title": "Order Transferring From the Minister of Public Works and Government Services to the President of the Queen\u0027s Privy Council for Canada the Control and Supervision of the Office of Indian Residential Schools Resolution",
      "citation": "SI/2003-228",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-764",
      "title": "Order Authorizing the Minister of Regional Industrial Expansion to Exercise the Powers of the Minister of Fisheries and Oceans Under Subsection 6(1) of the Act With Respect to Certain Companies",
      "citation": "SOR/84-764",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-202",
      "title": "Order Transferring From the Minister of Social Development to the Minister of Human Resources and Skills Development the Control and Supervision of the Department of Human Resources and Skills Development",
      "citation": "SI/2003-202",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-204",
      "title": "Order Transferring From the Minister of Social Development to the Minister of Human Resources and Skills Development the Powers, Duties and Functions Relating to the Canada Millennium Scholarships Foundation",
      "citation": "SI/2003-204",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-22",
      "title": "Order Transferring From the Minister of State (Infrastructure and Communities) to the Minister of Transport the Control and Supervision of the Office of Infrastructure of Canada",
      "citation": "SI/2006-22",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-91",
      "title": "Order Authorizing the Minister of State (Privatization and Regulatory Affairs) to Exercise the Powers Given to the Minister of Fisheries and Oceans Under the Act in Respect of Certain Companies",
      "citation": "SOR/87-91",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-111",
      "title": "Order Designating the Minister of Supply and Services as Appropriate Minister With Respect to the Defence Construction (1951) Ltd.",
      "citation": "SI/93-111",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-112",
      "title": "Order Designating the Minister of Supply and Services as Appropriate Minister with Respect to the Queens Quay West Land Corporation",
      "citation": "SI/93-112",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-21",
      "title": "Order Designating the Minister of the Environment as the Minister Responsible for the Administration and Enforcement of Subsections 36(3) to (6) of the Fisheries Act",
      "citation": "SI/2014-21",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-103",
      "title": "Order Assigning to the Minister of the Environment the Administration of Certain Public Lands",
      "citation": "SI/98-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-42",
      "title": "Order Assigning to the Minister of the Environment, the Administration of Certain Public Lands",
      "citation": "SI/95-42",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-93",
      "title": "Order Authorizing the Minister of the Environment to Exercise, With the Concurrence of the Minister of National Defence, the Administration of Certain Public Lands",
      "citation": "SI/2003-93",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-89",
      "title": "Order Authorizing the Minister of the Environment to Prescribe Charges",
      "citation": "SI/98-89",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-102",
      "title": "Order Transferring From the Minister of the Environment to the Minister of State (Infrastructure and Communities) the Control and Supervision of the Office of Infrastructure of Canada and Ordering the Minister to Preside Over the Office",
      "citation": "SI/2004-102",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-25",
      "title": "Order Designating the Minister of Transport as Appropriate Minister for the Canada Lands Company Limited for Purposes of the Act",
      "citation": "SI/2006-25",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-26",
      "title": "Order Designating the Minister of Transport as Appropriate Minister for the Queens Quay West Land Corporation for Purposes of the Act",
      "citation": "SI/2006-26",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-32",
      "title": "Order Designating the Minister of Transport as Appropriate Minister With Respect to the National Capital Commission for Purposes of the Act",
      "citation": "SI/2006-32",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-31",
      "title": "Order Designating the Minister of Transport as Minister for Purposes of the Act",
      "citation": "SI/2006-31",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-78",
      "title": "Order Designating the Minister of Transport as the appropriate Minister for the Windsor-Detroit Bridge Authority",
      "citation": "SI/2012-78",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-81-16",
      "title": "Minister of Transport Authority to Prescribe Charges Order",
      "citation": "SI/81-16",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-81-61",
      "title": "Minister of Transport Authority to Prescribe Fees or Charges Order",
      "citation": "SI/81-61",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-217",
      "title": "Minister of Transport Authority to Prescribe Fees Order",
      "citation": "SI/83-217",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-97",
      "title": "Order Authorizing the Minister of Transport, Infrastructure and Communities to Enter Into a Contribution Agreement With Harbourfront Corporation (1990)",
      "citation": "SI/2006-97",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-240",
      "title": "Order Transferring From the Minister of Transport to the Minister of National Revenue the Control and Supervision of the Royal Canadian Mint and the Powers, Duties and Functions Under the Royal Canadian Mint Act",
      "citation": "SI/2003-240",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-87",
      "title": "Order Transferring from the Minister of Transport to the President of the Queen\u0027s Privy Council for Canada the Control and Supervision of the portion of the federal public administration known as the Office of Infrastructure of Canada and Ordering the President of the Queen\u0027s Privy Council for Canada to Preside over the Office",
      "citation": "SI/2013-87",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-116",
      "title": "Order Designating the Minister of Western Economic Diversification as Minister for Purposes of the Act, Except Section 3 in Relation to the Western Provinces",
      "citation": "SI/88-116",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-112",
      "title": "Order Transferring to the Minister of Western Economic Diversification the Powers, Duties and Functions of the Minister of Indian Affairs and Northern Development Under Certain Acts in Relation to the Western Provinces",
      "citation": "SI/88-112",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-81-81",
      "title": "Minister Responsible for the Canada Mortgage and Housing Corporation Authority to Prescribe Fees Order",
      "citation": "SI/81-81",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-425",
      "title": "Ministers Designation Order (Canada Corporations Act)",
      "citation": "CRC, c 425",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-141",
      "title": "Proclamation Exempting the Mink Arm portion of South McMahon Lake from the Operation of Section 22 of the Act",
      "citation": "SOR/2013-141",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-401",
      "title": "Minority Investment (Bank Holding Companies) Regulations",
      "citation": "SOR/2001-401",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-402",
      "title": "Minority Investment (Banks) Regulations",
      "citation": "SOR/2001-402",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-403",
      "title": "Minority Investment (Cooperative Credit Associations) Regulations",
      "citation": "SOR/2001-403",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-404",
      "title": "Minority Investment (Insurance Companies) Regulations",
      "citation": "SOR/2001-404",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-405",
      "title": "Minority Investment (Insurance Holding Companies) Regulations",
      "citation": "SOR/2001-405",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-406",
      "title": "Minority Investment (Trust and Loan Companies) Regulations",
      "citation": "SOR/2001-406",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1448",
      "title": "Minor Waters Order",
      "citation": "CRC, c 1448",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-107",
      "title": "Mobile Offshore Drilling Units Remissions Order, 2004",
      "citation": "SOR/2004-107",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-68",
      "title": "Mohawk Council of Akwesasne Sewage Treatment Plant Remission Order",
      "citation": "SI/95-68",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-818",
      "title": "Moncton Airport Zoning Regulations",
      "citation": "SOR/90-818",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-475",
      "title": "Money Market Mutual Fund Conditions Regulations",
      "citation": "SOR/2001-475",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-95",
      "title": "Mont-Joli Airport Zoning Regulations",
      "citation": "CRC, c 95",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-96",
      "title": "Montreal International Airport Zoning Regulations",
      "citation": "CRC, c 96",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-484",
      "title": "Montreal Port Warden\u0027s Tariff of Fees Order (1988)",
      "citation": "SOR/88-484",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-138",
      "title": "Moose Jaw Airport Zoning Regulations",
      "citation": "SOR/2002-138",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-236",
      "title": "Moosonee Airport Zoning Regulations",
      "citation": "SOR/89-236",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-68",
      "title": "Mortgage Insurance Business (Banks, Authorized Foreign Banks, Trust and Loan Companies, Retail Associations, Canadian Insurance Companies and Canadian Societies) Regulations",
      "citation": "SOR/2010-68",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-69",
      "title": "Mortgage Insurance Disclosure (Banks, Authorized Foreign Banks, Trust and Loan Companies, Retail Associations, Canadian Insurance Companies and Canadian Societies) Regulations",
      "citation": "SOR/2010-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-303",
      "title": "Most-Favoured-Nation Tariff Extension of Benefit Order—Ferro-Chromium",
      "citation": "SOR/90-303",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-33",
      "title": "Most-Favoured-Nation Tariff Rules of Origin Regulations",
      "citation": "SOR/98-33",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-180",
      "title": "Motor Carrier Safety Fitness Certificate Regulations",
      "citation": "SOR/2005-180",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-990",
      "title": "Motor Vehicle Operators Hours of Work Regulations",
      "citation": "CRC, c 990",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-159",
      "title": "Motor Vehicle Restraint Systems and Booster Cushions Safety Regulations",
      "citation": "SOR/98-159",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-90",
      "title": "Motor Vehicle Restraint Systems and Booster Seats Safety Regulations",
      "citation": "SOR/2010-90",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1038",
      "title": "Motor Vehicle Safety Regulations",
      "citation": "CRC, c 1038",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-71",
      "title": "Motor Vehicles Tariff Order, 1988",
      "citation": "SOR/88-71",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-198",
      "title": "Motor Vehicle Tire Safety Regulations",
      "citation": "SOR/2013-198",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-148",
      "title": "Motor Vehicle Tire Safety Regulations, 1995",
      "citation": "SOR/95-148",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-167",
      "title": "Multinational Force and Observers Medal Order",
      "citation": "SI/86-167",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-127",
      "title": "Rules of the Municipal Courts",
      "citation": "SI/2005-127",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-415",
      "title": "Mushuau Innu First Nation Band Order",
      "citation": "SOR/2002-415",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-567",
      "title": "Muskoka Airport Zoning Regulations",
      "citation": "SOR/84-567",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-354",
      "title": "Musquash Estuary Marine Protected Area Regulations",
      "citation": "SOR/2006-354",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-128",
      "title": "Mutual Company (Life Insurance) Conversion Regulations",
      "citation": "SOR/99-128",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-151",
      "title": "“MV Sonia” Remission Order, 2007",
      "citation": "SOR/2007-151",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-333",
      "title": "NAFTA and CCFTA Verification of Origin Regulations",
      "citation": "SOR/97-333",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-49",
      "title": "NAFTA Marking Determination, Re-determination and Further Re-determination Regulations",
      "citation": "SOR/98-49",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-608",
      "title": "NAFTA Prescribed Class of Goods Regulations",
      "citation": "SOR/93-608",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-593",
      "title": "NAFTA Rules of Origin for Casual Goods Regulations",
      "citation": "SOR/93-593",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-14",
      "title": "NAFTA Rules of Origin Regulations",
      "citation": "SOR/94-14",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-17",
      "title": "NAFTA Tariff Preference Regulations",
      "citation": "SOR/94-17",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-158",
      "title": "Name Use (Affiliates of Banks or Bank Holding Companies) Regulations",
      "citation": "SOR/2008-158",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-103",
      "title": "Name Use (Affiliates of Banks, or Bank Holding Companies, that Are Not Widely Held) Regulations",
      "citation": "SOR/2002-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-408",
      "title": "Name Use by Non-financial Businesses (Excluded Entities) Regulations",
      "citation": "SOR/2001-408",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-260",
      "title": "Name Use (Cooperative Credit Associations) Regulations",
      "citation": "SOR/92-260",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-407",
      "title": "Name Use (Foreign Banks) Regulations",
      "citation": "SOR/2001-407",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-156",
      "title": "Name Use (Foreign Banks) Regulations",
      "citation": "SOR/2008-156",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-409",
      "title": "Name Use in Securities-Related Transactions (Banks and Bank Holding Companies) Regulations",
      "citation": "SOR/2001-409",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-410",
      "title": "Name Use in Securities-Related Transactions (Insurance Companies and Insurance Holding Companies) Regulations",
      "citation": "SOR/2001-410",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-257",
      "title": "Name Use (Trust and Loan Companies) Regulations",
      "citation": "SOR/92-257",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1041",
      "title": "Narcotic Control Regulations",
      "citation": "CRC, c 1041",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-272",
      "title": "Nardeux Canada Ltée Regulations",
      "citation": "SOR/87-272",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-55",
      "title": "Proclamation Declaring June 21 of Each Year as National Aboriginal Day",
      "citation": "SI/96-55",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-519",
      "title": "National Battlefields Park By-Law",
      "citation": "SOR/91-519",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-164",
      "title": "National Capital Commission Animal Regulations",
      "citation": "SOR/2002-164",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1044",
      "title": "National Capital Commission Traffic and Property Regulations",
      "citation": "CRC, c 1044",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-61",
      "title": "Order Designating the National Capital Region as the Place in Canada at Which the Principal Office of the National Round Table on the Environment and the Economy Shall be Located",
      "citation": "SI/94-61",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-89",
      "title": "Order Specifying the National Capital Region as the Place in Which Shall be Located the Head Office of Certain Museums",
      "citation": "SI/90-89",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-137",
      "title": "Order Designating the National Capital Region as the Place of the Head Office of the Hazardous Materials Information Review Commission",
      "citation": "SI/88-137",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-27",
      "title": "Proclamation Declaring May 9, 2014 as a “National Day of Honour” (Afghanistan Mission)",
      "citation": "SI/2014-27",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-716",
      "title": "National Defence Official Mementos Regulations",
      "citation": "CRC, c 716",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-201a",
      "title": "Department of National Defence Terms Under Three Months Exclusion Approval Order, 1992",
      "citation": "SOR/92-201a",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-244",
      "title": "National Energy Board Act Part VI (Oil and Gas) Regulations",
      "citation": "SOR/96-244",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-7",
      "title": "National Energy Board Cost Recovery Regulations",
      "citation": "SOR/91-7",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-130",
      "title": "National Energy Board Electricity Regulations",
      "citation": "SOR/97-130",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-563",
      "title": "National Energy Board Export and Import Reporting Regulations",
      "citation": "SOR/95-563",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-294",
      "title": "National Energy Board Onshore Pipeline Regulations",
      "citation": "SOR/99-294",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-294",
      "title": "National Energy Board Onshore Pipeline Regulations",
      "citation": "SOR/99-294",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1055",
      "title": "National Energy Board Order No. MO-62-69",
      "citation": "CRC, c 1055",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-528",
      "title": "National Energy Board Pipeline Crossing Regulations, Part I",
      "citation": "SOR/88-528",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-529",
      "title": "National Energy Board Pipeline Crossing Regulations, Part II",
      "citation": "SOR/88-529",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-39",
      "title": "National Energy Board Processing Plant Regulations",
      "citation": "SOR/2003-39",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-208",
      "title": "National Energy Board Rules of Practice and Procedure, 1995",
      "citation": "SOR/95-208",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-191",
      "title": "National Energy Board Substituted Service Regulations",
      "citation": "SOR/83-191",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-263",
      "title": "National Historic Parks General Regulations",
      "citation": "SOR/82-263",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-613",
      "title": "National Historic Parks Wildlife and Domestic Animals Regulations",
      "citation": "SOR/81-613",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1112",
      "title": "National Historic Sites of Canada Order",
      "citation": "CRC, c 1112",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-199",
      "title": "National Library Book Deposit Regulations, 1995",
      "citation": "SOR/95-199",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-64",
      "title": "National Mining Week Proclamation",
      "citation": "SI/95-64",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1114",
      "title": "National Parks Building Regulations",
      "citation": "CRC, c 1114",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-127",
      "title": "National Parks Camping Regulations",
      "citation": "SOR/80-127",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-677",
      "title": "National Parks Cemetery Regulations",
      "citation": "SOR/83-677",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-217",
      "title": "National Parks Garbage Regulations",
      "citation": "SOR/80-217",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-213",
      "title": "National Parks General Regulations",
      "citation": "SOR/78-213",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1126",
      "title": "National Parks Highway Traffic Regulations",
      "citation": "CRC, c 1126",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-23",
      "title": "National Parks Land Rents Remission Order",
      "citation": "SI/2000-23",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-50",
      "title": "National Parks Land Rents Remission Order, No. 2",
      "citation": "SI/2001-50",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-150",
      "title": "National Parks of Canada Aircraft Access Regulations",
      "citation": "SOR/97-150",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-455",
      "title": "National Parks of Canada Businesses Regulations",
      "citation": "SOR/98-455",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-398",
      "title": "National Parks of Canada Cottages Regulations",
      "citation": "SOR/79-398",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-177",
      "title": "National Parks of Canada Domestic Animals Regulations",
      "citation": "SOR/98-177",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-946",
      "title": "National Parks of Canada Fire Protection Regulations",
      "citation": "SOR/80-946",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1120",
      "title": "National Parks of Canada Fishing Regulations",
      "citation": "CRC, c 1120",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-65",
      "title": "National Parks of Canada Land Rents Remission Order, No. 3",
      "citation": "SI/2002-65",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-25",
      "title": "National Parks of Canada Lease and Licence of Occupation Regulations",
      "citation": "SOR/92-25",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1134",
      "title": "National Parks of Canada Water and Sewer Regulations",
      "citation": "CRC, c 1134",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1130",
      "title": "National Parks Signs Regulations",
      "citation": "CRC, c 1130",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-8",
      "title": "National Parks Town, Visitor Centre and Resort Subdivision Designation Regulations",
      "citation": "SOR/91-8",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-387",
      "title": "National Parks Wilderness Area Declaration Regulations",
      "citation": "SOR/2000-387",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-401",
      "title": "National Parks Wildlife Regulations",
      "citation": "SOR/81-401",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-547",
      "title": "National Revenue Transfer of Property Regulations",
      "citation": "SOR/81-547",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-271",
      "title": "National Security Review of Investments Regulations",
      "citation": "SOR/2009-271",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-39",
      "title": "Order Designating the National Transportation Act Review Commission as a Department and the Prime Minister as Appropriate Minister",
      "citation": "SI/92-39",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-84",
      "title": "Proclamation Declaring October 31st of Each Year to be “National UNICEF Day”",
      "citation": "SI/2000-84",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1321",
      "title": "NATO Common Infrastructure Project Remission Order",
      "citation": "CRC, c 1321",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-73",
      "title": "Natural and Man-made Harbour Navigation and Use Regulations",
      "citation": "SOR/2005-73",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-359",
      "title": "Natural Gas and Gas Liquids Tax Regulations",
      "citation": "SOR/82-359",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-196",
      "title": "Natural Health Products Regulations",
      "citation": "SOR/2003-196",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-171",
      "title": "Natural Health Products (Unprocessed Product Licence Applications) Regulations",
      "citation": "SOR/2010-171",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-281",
      "title": "Nautical Charts and Related Publications Fees Order",
      "citation": "SOR/94-281",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-479",
      "title": "NAV CANADA Divestiture Regulations",
      "citation": "SOR/96-479",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1231",
      "title": "Navigable Waters Bridges Regulations",
      "citation": "CRC, c 1231",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1232",
      "title": "Navigable Waters Works Regulations",
      "citation": "CRC, c 1232",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-134",
      "title": "Navigation Safety Regulations",
      "citation": "SOR/2005-134",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-23",
      "title": "Negative Option Billing Regulations",
      "citation": "SOR/2012-23",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-232",
      "title": "Order Authorizing Negotiations for the Settlement of the Dispute Causing the Extraordinary Disruption of the National Transportation System in Relation to Container Movements Into and out of Certain Ports in British Columbia",
      "citation": "SOR/2005-232",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-348",
      "title": "Regulations Prescribing Networks (Copyright Act)",
      "citation": "SOR/99-348",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-238",
      "title": "New Brunswick Blueberry Order",
      "citation": "SOR/2013-238",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-80-117",
      "title": "New Brunswick Court of Queen\u0027s Bench Summary Conviction Appeal Rules",
      "citation": "SI/80-117",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-82-13",
      "title": "New Brunswick Criminal Appeal Rule 63 With Respect to Criminal Appeals to the Court of Appeal",
      "citation": "SI/82-13",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-162",
      "title": "New Brunswick Egg Order",
      "citation": "CRC, c 162",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-1",
      "title": "New Brunswick Hog Marketing Levies (Interprovincial and Export Trade) Order",
      "citation": "SOR/84-1",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-42",
      "title": "New Brunswick Hog Order",
      "citation": "SOR/82-42",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-627",
      "title": "New Brunswick Milk Order",
      "citation": "SOR/94-627",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-401",
      "title": "New Brunswick Potato Marketing Levies (Interprovincial and Export) Order",
      "citation": "SOR/85-401",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-307",
      "title": "New Brunswick Potato Marketing Levies (Interprovincial and Export) Order — No. 2",
      "citation": "SOR/88-307",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-726",
      "title": "New Brunswick Potato Order",
      "citation": "SOR/80-726",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-228",
      "title": "New Brunswick Primary Forest Products Order",
      "citation": "SOR/2000-228",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-264",
      "title": "New Brunswick Rules of Practice Respecting Reduction in the Number of Years of Imprisonment Without Eligibility for Parole",
      "citation": "SOR/2004-264",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-16",
      "title": "New Brunswick Sex Offender Information Registration Regulations",
      "citation": "SOR/2005-16",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-2",
      "title": "New Brunswick Summary Conviction Appeal Rule 64 With Respect to Summary Conviction Appeals to the Court of Queen\u0027s Bench",
      "citation": "SI/92-2",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-9",
      "title": "New Brunswick Translated Documents Regulations",
      "citation": "SOR/93-9",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-164",
      "title": "New Brunswick Turkey Order",
      "citation": "CRC, c 164",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-230",
      "title": "New Classes of Practitioners Regulations",
      "citation": "SOR/2012-230",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-443",
      "title": "Newfoundland and Labrador Fishery Regulations",
      "citation": "SOR/78-443",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-192",
      "title": "Newfoundland and Labrador Offshore Area Line Regulations",
      "citation": "SOR/2003-192",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-302",
      "title": "Newfoundland and Labrador Offshore Revenue Fiscal Equalization Offset Payments Regulations",
      "citation": "SOR/2002-302",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-322",
      "title": "Newfoundland and Labrador Sex Offender Information Registration Regulations",
      "citation": "SOR/2004-322",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-347",
      "title": "Newfoundland Offshore Area Oil and Gas Operations Regulations",
      "citation": "SOR/88-347",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-601",
      "title": "Newfoundland Offshore Area Petroleum Diving Regulations",
      "citation": "SOR/88-601",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-334",
      "title": "Newfoundland Offshore Area Petroleum Geophysical Operations Regulations",
      "citation": "SOR/95-334",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-103",
      "title": "Newfoundland Offshore Area Petroleum Production and Conservation Regulations",
      "citation": "SOR/95-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-263",
      "title": "Newfoundland Offshore Area Registration Regulations",
      "citation": "SOR/88-263",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-100",
      "title": "Newfoundland Offshore Certificate of Fitness Regulations",
      "citation": "SOR/95-100",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-316",
      "title": "Newfoundland Offshore Petroleum Drilling and Production Regulations",
      "citation": "SOR/2009-316",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-23",
      "title": "Newfoundland Offshore Petroleum Drilling Regulations",
      "citation": "SOR/93-23",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-104",
      "title": "Newfoundland Offshore Petroleum Installations Regulations",
      "citation": "SOR/95-104",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-257",
      "title": "Newfoundland Offshore Petroleum Resource Revenue Fund Regulations",
      "citation": "SOR/95-257",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-569",
      "title": "Newfoundland Railway Reimbursement Regulations",
      "citation": "SOR/81-569",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-297",
      "title": "Newfoundland Rules of Practice Respecting Reduction in the Number of Years of Imprisonment Without Eligibility for Parole",
      "citation": "SOR/89-297",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-117",
      "title": "New Harmonized Value-added Tax System Regulations",
      "citation": "SOR/2010-117",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-151",
      "title": "New Harmonized Value-added Tax System Regulations, No. 2",
      "citation": "SOR/2010-151",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-98",
      "title": "New Montreal International Airport (Mirabel) Zoning Regulations",
      "citation": "CRC, c 98",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-710",
      "title": "New Motor Vehicles Exported Drawback Regulations",
      "citation": "SOR/82-710",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-374",
      "title": "New Substances Fees Regulations",
      "citation": "SOR/2002-374",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-247",
      "title": "New Substances Notification Regulations (Chemicals and Polymers)",
      "citation": "SOR/2005-247",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-248",
      "title": "New Substances Notification Regulations (Organisms)",
      "citation": "SOR/2005-248",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-14",
      "title": "New Zealand Service Aircraft Over-flight Regulations",
      "citation": "CRC, c 14",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-641",
      "title": "NFPMC General Rules of Procedure",
      "citation": "SOR/82-641",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-686",
      "title": "N.H.A. Designated Areas Order",
      "citation": "SOR/84-686",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1107",
      "title": "N.H.A. Maximum Interest Rates Regulations",
      "citation": "CRC, c 1107",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-39",
      "title": "Nisga\u0027a Final Agreement Indian Remission Order",
      "citation": "SI/2000-39",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-610",
      "title": "Non-alcoholic Wine Remission Order",
      "citation": "SOR/90-610",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-151",
      "title": "Non-article 5 North Atlantic Treaty Organisation (NATO) Medal for Operations in the Balkans Order",
      "citation": "SI/2003-151",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-87",
      "title": "Non-article 5 North Atlantic Treaty Organisation (NATO) Medal for the NATO Training Mission in Iraq Order",
      "citation": "SI/2006-87",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-96",
      "title": "Non-canadian Ships Safety Order",
      "citation": "SI/97-96",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-79-39",
      "title": "Non-commercial Importations Remission Order",
      "citation": "SI/79-39",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-10",
      "title": "Non-mailable Matter Regulations",
      "citation": "SOR/90-10",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-659",
      "title": "Non-Pleasure Craft Sewage Pollution Prevention Regulations",
      "citation": "SOR/91-659",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-10",
      "title": "Non-profit Grazing Associations (GST) Remission Order",
      "citation": "SI/93-10",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-361a",
      "title": "Non-public Funds Staff Exclusion Approval Order",
      "citation": "SOR/82-361a",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-361b",
      "title": "Non-public Funds Staff Regulations",
      "citation": "SOR/82-361b",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-42",
      "title": "Non-resident Rebate (GST) Regulations",
      "citation": "SOR/91-42",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-720",
      "title": "Non-residents\u0027 Temporary Importation of Baggage and Conveyances Regulations",
      "citation": "SOR/87-720",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-112",
      "title": "Non-signatory Yukon First Nations Remission Order",
      "citation": "SI/2003-112",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-21",
      "title": "Non-smokers\u0027 Health Regulations",
      "citation": "SOR/90-21",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-31",
      "title": "Non-taxable Imported Goods (GST/HST) Regulations",
      "citation": "SOR/91-31",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-421",
      "title": "Nordion and Theratronics Employees Pension Protection Regulations",
      "citation": "SOR/95-421",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-162",
      "title": "Nordion Europe S.A. Incorporation Authorization Order",
      "citation": "SOR/90-162",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-163",
      "title": "Nordion International Incorporation Authorization Order (Incorporation of Corporations)",
      "citation": "SOR/90-163",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-502",
      "title": "Nordion International Inc./Theratronics International Limited Shares Order, 1988",
      "citation": "SOR/88-502",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-296",
      "title": "Norman Wells Airport Zoning Regulations",
      "citation": "SOR/82-296",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-194",
      "title": "North American Leaders\u0027 Summit 2007 Privileges and Immunities Order",
      "citation": "SOR/2007-194",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-103",
      "title": "North Atlantic Treaty Organization (NATO) Medal for the Former Yugoslav Republic of Macedonia (FYROM) Order",
      "citation": "SI/2002-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-105",
      "title": "North Atlantic Treaty Organization (NATO) Medal Order",
      "citation": "SI/95-105",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-14",
      "title": "North Atlantic Treaty Organization (NATO) Medal with \"Kosovo\" (KFOR) Bar Order",
      "citation": "SI/2000-14",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-648",
      "title": "North Battleford Airport Zoning Regulations",
      "citation": "SOR/92-648",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-99",
      "title": "North Bay Airport Zoning Regulations",
      "citation": "CRC, c 99",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-76",
      "title": "Northern Canada Power Commission Yukon Assets Disposal Authorization Act",
      "citation": "SI/2014-76",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-127",
      "title": "Northern Canada Vessel Traffic Services Zone Regulations",
      "citation": "SOR/2010-127",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-332",
      "title": "Northern Mineral Exploration Assistance Regulations",
      "citation": "CRC, c 332",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-503",
      "title": "Northern Ontario Loan Insurance Regulations",
      "citation": "SOR/88-503",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-133",
      "title": "Northern Pipeline Agency Cost Recovery Charge Remission Order, 2004",
      "citation": "SI/2004-133",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-235",
      "title": "Northern Pipeline Notice of Objection Regulations",
      "citation": "SOR/81-235",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-81-22",
      "title": "Northern Pipeline Socio-Economic and Environmental Terms and Conditions for Northern British Columbia",
      "citation": "SI/81-22",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-80-128",
      "title": "Northern Pipeline Socio-Economic and Environmental Terms and Conditions for Southern British Columbia",
      "citation": "SI/80-128",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-80-127",
      "title": "Northern Pipeline Socio-Economic and Environmental Terms and Conditions for the Province of Alberta",
      "citation": "SI/80-127",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-81-23",
      "title": "Northern Pipeline Socio-Economic and Environmental Terms and Conditions for the Province of Saskatchewan",
      "citation": "SI/81-23",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-81-21",
      "title": "Northern Pipeline Socio-Economic and Environmental Terms and Conditions for the Swift River Portion of the Pipeline in the Province of British Columbia",
      "citation": "SI/81-21",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-658",
      "title": "Northern Transportation Company Limited Exemption and Transfer Order",
      "citation": "SOR/85-658",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-562",
      "title": "North Pacific Anadromous Fish Commission Privileges and Immunities Order",
      "citation": "SOR/94-562",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-611",
      "title": "North Pacific Marine Science Organization Privileges and Immunities Order",
      "citation": "SOR/93-611",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-64",
      "title": "Northwest Atlantic Fisheries Organization Privileges and Immunities Order",
      "citation": "SOR/80-64",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-576",
      "title": "Northwestel Inc. Shares Sale Order",
      "citation": "SOR/88-576",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1516",
      "title": "Northwest Territories and Nunavut Mining Regulations",
      "citation": "CRC, c 1516",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-219",
      "title": "Northwest Territories Archaeological Sites Regulations",
      "citation": "SOR/2001-219",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-39",
      "title": "Northwest Territories Borrowing Limits Regulations",
      "citation": "SOR/2013-39",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-847",
      "title": "Northwest Territories Fishery Regulations",
      "citation": "CRC, c 847",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-53",
      "title": "Northwest Territories Mining District and Nunavut Mining District Order",
      "citation": "SI/2000-53",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-68",
      "title": "Northwest Territories Mining Regulations",
      "citation": "SOR/2014-68",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1238",
      "title": "Northwest Territories Reindeer Regulations",
      "citation": "CRC, c 1238",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-392",
      "title": "Northwest Territories Rules of Practice Respecting Applications and Hearings concerning a Reduction in the Number of Years of Imprisonment Without Eligibility for Parole",
      "citation": "SOR/98-392",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-1",
      "title": "Northwest Territories Sex Offender Information Registration Regulations",
      "citation": "SOR/2005-1",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-303",
      "title": "Northwest Territories Waters Regulations",
      "citation": "SOR/93-303",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-36",
      "title": "Order Giving Notice That a Tax Agreement Between Canada and the People\u0027s Republic of China Came Into Force on December 29, 1986",
      "citation": "SI/87-36",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-51",
      "title": "Proclamation Giving Notice That the Administrative Understanding on Mutual Assistance Shall be Effective Upon the Entering Into Force of the Second Supplementary Agreement Between Canada and the United States of America",
      "citation": "SI/98-51",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-120",
      "title": "Proclamation Giving Notice That the Agreement between Canada and Saint Vincent and the Grenadines is in Force as of November 1, 1998",
      "citation": "SI/98-120",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-4",
      "title": "Proclamation Giving Notice That the Agreement on Social Security Between Canada and Australia Comes into Force on January 1, 2003",
      "citation": "SI/2003-4",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-11",
      "title": "Proclamation Giving Notice That the Agreement on Social Security Between Canada and Grenada is in Force as of February 1, 1999",
      "citation": "SI/99-11",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2008-25",
      "title": "Proclamation Giving Notice That the Agreement on Social Security Between Canada and Japan Comes into Force on March 1, 2008",
      "citation": "SI/2008-25",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-42",
      "title": "Proclamation Giving Notice That the Agreement on Social Security Between Canada and Sweden Comes into Force on April 1, 2003",
      "citation": "SI/2003-42",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-5",
      "title": "Proclamation Giving Notice That the Agreement on Social Security Between Canada and the Czech Republic Comes into Force on January 1, 2003",
      "citation": "SI/2003-5",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-1",
      "title": "Proclamation Giving Notice That the Agreement on Social Security Between Canada and the Eastern Republic of Uruguay is in Force as of January 1, 2002",
      "citation": "SI/2002-1",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-34",
      "title": "Proclamation Giving Notice That the Agreement on Social Security Between Canada and the Kingdom of the Netherlands will be in Force on April 1, 2004",
      "citation": "SI/2004-34",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-121",
      "title": "Proclamation Giving Notice that the Agreement on Social Security Between Canada and the Republic of Estonia Comes Into Force on November 1, 2006",
      "citation": "SI/2006-121",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-159",
      "title": "Proclamation Giving Notice That the Agreement on Social Security Between Canada and the Republic of Hungary Comes Into Force on October 1, 2003",
      "citation": "SI/2003-159",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-122",
      "title": "Proclamation Giving Notice That the Agreement on Social Security Between Canada and the Republic of Latvia Comes Into Force on November 1, 2006",
      "citation": "SI/2006-122",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-123",
      "title": "Proclamation Giving Notice That the Agreement on Social Security Between Canada and the Republic of Lithuania Comes Into Force on November 1, 2006",
      "citation": "SI/2006-123",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-72",
      "title": "Proclamation Giving Notice That  the Agreement on Social Security Between Canada and the Republic of Slovenia Came Into Force on January 1, 2001",
      "citation": "SI/2001-72",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-76",
      "title": "Proclamation Giving Notice That the Agreement on Social Security Between Canada and the Republic of Trinidad and Tobago is in Force as of July 1, 1999",
      "citation": "SI/99-76",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-6",
      "title": "Proclamation Giving Notice That the Agreement on Social Security between Canada and the Slovak Republic Comes into Force on January 1, 2003",
      "citation": "SI/2003-6",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-62",
      "title": "Proclamation Giving Notice That the Agreement on Social Security Between the Government of Canada and the Government of the Republic of Chile Comes into Force on June 1, 1998",
      "citation": "SI/98-62",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-49",
      "title": "Proclamation Giving Notice That the Agreement on Social Security between the Government of Canada and the Government of the Republic of Croatia is in Force as of May 1, 1999",
      "citation": "SI/99-49",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-50",
      "title": "Proclamation Giving Notice That the Agreement on Social Security Between the Government of Canada and the Government of the Republic of Korea is in Force as of May 1, 1999",
      "citation": "SI/99-50",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-164",
      "title": "Proclamation Giving Notice that the Agreement on Social Security Between the Government of Canada and the Government of the Republic of Turkey Comes Into Force on January 1, 2005",
      "citation": "SI/2004-164",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-19",
      "title": "Proclamation Giving Notice that the Annexed November 30, 1995 Supplementary Agreement, Entitled Protocol to the Tax Convention Between the Government of Canada and the Government of the French Republic signed on May 2, 1975 and Amended by the Protocol of January 16, 1987, Came into force on September 1, 1998",
      "citation": "SI/99-19",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-48",
      "title": "Proclamation giving notice that the attached supplementary convention, which alters and adds to the Convention set out in Schedule II to the Act, came into force on January 29, 2001",
      "citation": "SI/2001-48",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-16",
      "title": "Proclamation Giving Notice that the Convention between the Government of Canada and the Government of the Italian Republic for the Avoidance of Double Taxation with Respect to Taxes on Income and the Prevention of Fiscal Evasion Came into Force on November 25, 2011",
      "citation": "SI/2012-16",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-94",
      "title": "Proclamation Giving Notice that the Convention Between the Government of Canada and the Swiss Federal Council for the Avoidance of Double Taxation with Respect to Taxes on Income and on Capital Came into Force on April 21, 1998",
      "citation": "SI/98-94",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-49",
      "title": "Proclamation giving notice that the Convention on Social Security between the Government of Canada and the Government of the United Kingdom of Great Britain and Northern Ireland comes into force on April 1, 1998",
      "citation": "SI/98-49",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-155",
      "title": "Proclamation Giving Notice that the Interim Agreement on Social Security between Canada and Israel Comes Into Force on September 1, 2003",
      "citation": "SI/2003-155",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-119",
      "title": "Proclamation Giving Notice that the Protocol Amending the Convention between Canada and Australia Came into Force on December 18, 2002",
      "citation": "SI/2003-119",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-77",
      "title": "Proclamation giving notice that the Protocol Amending the Convention between Canada and the Republic of Indonesia for the Avoidance of Double Taxation and the Prevention of Fiscal Evasion with respect to Taxes on Income and on Capital came into force on December 31, 1998",
      "citation": "SI/99-77",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-173",
      "title": "Proclamation Giving Notice that the Supplementary Agreement to the Agreement on Social Security between Canada and the Federal Republic of Germany Comes Into Force on December 1, 2003",
      "citation": "SI/2003-173",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-81",
      "title": "Proclamation Giving Notice that the Supplementary Agreement to the Agreement on Social Security Between Canada and the Republic of the Philippines is in Force as of July 1, 2001",
      "citation": "SI/2001-81",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-348",
      "title": "Notifiable Transactions Regulations",
      "citation": "SOR/87-348",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-446",
      "title": "Nova BUS Corporation Exemption Order",
      "citation": "SOR/96-446",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-319",
      "title": "Nova Scotia and Newfoundland and Labrador Additional Fiscal Equalization Offset Payments Regulations",
      "citation": "SOR/2008-319",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-165",
      "title": "Nova Scotia Chicken Order",
      "citation": "CRC, c 165",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-166",
      "title": "Nova Scotia Egg Order",
      "citation": "CRC, c 166",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-167",
      "title": "Nova Scotia Hog Order",
      "citation": "CRC, c 167",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-99",
      "title": "Nova Scotia HST Regulations, 2010",
      "citation": "SOR/2010-99",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-626",
      "title": "Nova Scotia Milk Order",
      "citation": "SOR/94-626",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-189",
      "title": "Nova Scotia Offshore Area Petroleum Diving Regulations",
      "citation": "SOR/95-189",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-144",
      "title": "Nova Scotia Offshore Area Petroleum Geophysical Operations Regulations",
      "citation": "SOR/95-144",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-190",
      "title": "Nova Scotia Offshore Area Petroleum Production and Conservation Regulations",
      "citation": "SOR/95-190",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-187",
      "title": "Nova Scotia Offshore Certificate of Fitness Regulations",
      "citation": "SOR/95-187",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-317",
      "title": "Nova Scotia Offshore Petroleum Drilling and Production Regulations",
      "citation": "SOR/2009-317",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-676",
      "title": "Nova Scotia Offshore Petroleum Drilling Regulations",
      "citation": "SOR/92-676",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-191",
      "title": "Nova Scotia Offshore Petroleum Installations Regulations",
      "citation": "SOR/95-191",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-441",
      "title": "Nova Scotia Offshore Revenue Account Regulations",
      "citation": "SOR/93-441",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-249",
      "title": "Nova Scotia Offshore Revenue Fiscal Equalization Offset Payments Regulations",
      "citation": "SOR/96-249",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-16",
      "title": "Nova Scotia Public Service Long Term Disability Plan Trust Fund Remission Order",
      "citation": "SI/2000-16",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-168",
      "title": "Nova Scotia Resources (Ventures) Limited Drilling Assistance Regulations",
      "citation": "SOR/94-168",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-10",
      "title": "Nova Scotia Rules of Practice Respecting Applications and Hearings Concerning a Reduction in the Number of Years of Imprisonment without Eligibility for Parole",
      "citation": "SOR/2005-10",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-305",
      "title": "Nova Scotia Sex Offender Information Registration Regulations",
      "citation": "SOR/2004-305",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-848",
      "title": "Nova Scotia Share of Offshore Revenue Interim Period Payment Regulations",
      "citation": "SOR/84-848",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-912",
      "title": "Nova Scotia Share of Offshore Sales Tax Payment Regulations",
      "citation": "SOR/85-912",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-169",
      "title": "Nova Scotia Turkey Marketing Levies Order",
      "citation": "CRC, c 169",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-168",
      "title": "Nova Scotia Turkey Order",
      "citation": "CRC, c 168",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-719",
      "title": "Nova Scotia Wheat Order",
      "citation": "SOR/82-719",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-403",
      "title": "N.R.C. Sales and Loans Regulations",
      "citation": "SOR/81-403",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-210",
      "title": "Nuclear Non-proliferation Import and Export Control Regulations",
      "citation": "SOR/2000-210",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-209",
      "title": "Nuclear Security Regulations",
      "citation": "SOR/2000-209",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-207",
      "title": "Nuclear Substances and Radiation Devices Regulations",
      "citation": "SOR/2000-207",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-220",
      "title": "Nunavut Archaeological and Palaeontological Sites Regulations",
      "citation": "SOR/2001-220",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-40",
      "title": "Nunavut Borrowing Limits Regulations",
      "citation": "SOR/2013-40",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-69",
      "title": "Nunavut Mining Regulations",
      "citation": "SOR/2014-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-321",
      "title": "Nunavut Sex Offender Information Registration Regulations",
      "citation": "SOR/2004-321",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-69",
      "title": "Nunavut Waters Regulations",
      "citation": "SOR/2013-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-334",
      "title": "Nursing Home Care Benefits Regulations",
      "citation": "CRC, c 334",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-231",
      "title": "Oath or Solemn Affirmation of Office Rules (Immigration and Refugee Board)",
      "citation": "SOR/2002-231",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-255",
      "title": "Oath or Solemn Affirmation of Office Rules (Immigration and Refugee Board)",
      "citation": "SOR/2012-255",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1239",
      "title": "Oaths of Allegiance and Office and Seat of Government Order (N.W.T.)",
      "citation": "CRC, c 1239",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1611",
      "title": "Oaths of Allegiance and Office Order (Yukon)",
      "citation": "CRC, c 1611",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1242",
      "title": "Oaths of Office Regulations",
      "citation": "CRC, c 1242",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-132",
      "title": "Oats 1987 Period Stabilization Regulations",
      "citation": "SOR/89-132",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-282",
      "title": "Ocean Data and Services Fees Order, 1993",
      "citation": "SOR/94-282",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-114",
      "title": "Ocean Dumping Permit Fee Regulations (Site Monitoring)",
      "citation": "SOR/99-114",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-161",
      "title": "Regulations Prescribing Certain Offences to be Serious Offences",
      "citation": "SOR/2010-161",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-890",
      "title": "Off Grades of Grain and Grades of Screenings Order",
      "citation": "CRC, c 890",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-74",
      "title": "Order Designating the Office of Indian Residential Schools Resolution of Canada as a Department and the Executive Director and Deputy Head as the Deputy Head for Purposes of that Act",
      "citation": "SI/2001-74",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-100",
      "title": "Order Transferring from the Office of Infrastructure of Canada to the Department of Human Resources and Skills Development the Control and Supervision of the Crown Corporations Secretariat Relating to the Canada Mortgage and Housing Corporation",
      "citation": "SI/2004-100",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-30",
      "title": "Order Transferring from the Office of Infrastructure of Canada to the Department of Transport the Control and Supervision of the Canada Lands Company Limited and the Queens Quay West Land Corporation",
      "citation": "SI/2006-30",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-76-44",
      "title": "Designating the Office of the Co-ordinator, Status of Women as a Department, the Minister of National Health and Welfare as Appropriate Minister and the Co-ordinator as Deputy Head",
      "citation": "SI/76-44",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-114",
      "title": "Office of the Governor General\u0027s Secretary Employment Regulations",
      "citation": "SOR/2006-114",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-89",
      "title": "Office of the Governor-General\u0027s Secretary Exclusion Approval Order",
      "citation": "SI/2006-89",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-344",
      "title": "Office of the Registrar General Fees for Services Order",
      "citation": "SOR/88-344",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-18",
      "title": "Order Designating the Office of the Registrar of Lobbyists as a Department and the Registrar as Deputy Head",
      "citation": "SI/2006-18",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1066",
      "title": "Officers Authorized to Exercise Powers or Perform Duties of the Minister of National Revenue Regulations",
      "citation": "SOR/86-1066",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-48",
      "title": "Official Languages (Communications with and Services to the Public) Regulations",
      "citation": "SOR/92-48",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-32",
      "title": "Off-Road Compression-Ignition Engine Emission Regulations",
      "citation": "SOR/2005-32",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-355",
      "title": "Off-Road Small Spark-Ignition Engine Emission Regulations",
      "citation": "SOR/2003-355",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-49",
      "title": "Offset of Taxes (GST/HST) Regulations",
      "citation": "SOR/91-49",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-592",
      "title": "Offshore Area Exclusion Order",
      "citation": "SOR/84-592",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-7",
      "title": "Offshore Minerals Revenue Election Regulations",
      "citation": "SOR/2000-7",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1521",
      "title": "Oil and Gas Land Order No. 1-1962",
      "citation": "CRC, c 1521",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1520",
      "title": "Oil and Gas Land Order No. 2-1961",
      "citation": "CRC, c 1520",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-612",
      "title": "Oil and Gas Occupational Safety and Health Regulations",
      "citation": "SOR/87-612",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-331",
      "title": "Oil and Gas Spills and Debris Liability Regulations",
      "citation": "SOR/87-331",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-335",
      "title": "Oil Import Compensation Regulations No. 1, 1975",
      "citation": "CRC, c 335",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1058",
      "title": "Oil Pipeline Uniform Accounting Regulations",
      "citation": "CRC, c 1058",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-3",
      "title": "Oil Pollution Prevention Regulations",
      "citation": "SOR/93-3",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-216",
      "title": "Oil Product Designation Regulations",
      "citation": "SOR/88-216",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1246",
      "title": "Old Age Security Regulations",
      "citation": "CRC, c 1246",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-120",
      "title": "Old Crow Airport Zoning Regulations",
      "citation": "SOR/94-120",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-294",
      "title": "Olympic and Paralympic Marks Regulations",
      "citation": "SOR/2007-294",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-306",
      "title": "2010 Olympic and Paralympic Winter Games Remission Order",
      "citation": "SOR/2008-306",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-184",
      "title": "On Board Trains Occupational Safety and Health Regulations",
      "citation": "SOR/87-184",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1007",
      "title": "One-day Local Employment Seminars Fee or Charge Order",
      "citation": "SOR/86-1007",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-2",
      "title": "On-Road Vehicle and Engine Emission Regulations",
      "citation": "SOR/2003-2",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-170",
      "title": "Ontario Apple Order",
      "citation": "CRC, c 170",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-859",
      "title": "Ontario Asparagus Growers\u0027 Marketing-for-Processing Order",
      "citation": "SOR/78-859",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-176",
      "title": "Ontario Asparagus Marketing (Interprovincial and Export) Regulations",
      "citation": "SOR/80-176",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-175",
      "title": "Ontario Asparagus Pricing (Interprovincial and Export) Regulations",
      "citation": "SOR/80-175",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-177",
      "title": "Ontario Bean Marketing (Interprovincial and Export) Regulations",
      "citation": "CRC, c 177",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-176",
      "title": "Ontario Bean Order",
      "citation": "CRC, c 176",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-178",
      "title": "Ontario Berry-for-Processing Order",
      "citation": "CRC, c 178",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-179",
      "title": "Ontario Cheese Order",
      "citation": "CRC, c 179",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-180",
      "title": "Ontario Chicken Order",
      "citation": "CRC, c 180",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-99",
      "title": "Ontario Court of Justice Criminal Proceedings Rules",
      "citation": "SI/92-99",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-216",
      "title": "Ontario Cream Producers\u0027 Marketing Levies Order",
      "citation": "CRC, c 216",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-182",
      "title": "Ontario Egg Marketing Levies Order",
      "citation": "CRC, c 182",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-181",
      "title": "Ontario Egg Order",
      "citation": "CRC, c 181",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-93",
      "title": "Ontario Fishery Regulations, 1989",
      "citation": "SOR/89-93",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-237",
      "title": "Ontario Fishery Regulations, 2007",
      "citation": "SOR/2007-237",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-239",
      "title": "Ontario Flue-Cured Tobacco Licence Charges (Interprovincial and Export) Order",
      "citation": "SOR/86-239",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-185",
      "title": "Ontario Flue-Cured Tobacco Licensing (Interprovincial and Export) Regulations",
      "citation": "CRC, c 185",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-186",
      "title": "Ontario Flue-Cured Tobacco Marketing (Interprovincial and Export) Regulations",
      "citation": "CRC, c 186",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-183",
      "title": "Ontario Flue-Cured Tobacco Order",
      "citation": "CRC, c 183",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-307",
      "title": "Ontario Fresh Grape Information (Interprovincial and Export) Regulations",
      "citation": "SOR/78-307",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-194",
      "title": "Ontario Fresh Grape Order",
      "citation": "CRC, c 194",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-198",
      "title": "Ontario Fresh Grape Service Charge (Interprovincial and Export) Regulations",
      "citation": "CRC, c 198",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-553",
      "title": "Ontario Fresh Potato Order",
      "citation": "SOR/87-553",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-243",
      "title": "Ontario Grain Order",
      "citation": "SOR/2013-243",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-259",
      "title": "Ontario Grapes-for-Processing Marketing Levies Order",
      "citation": "SOR/2010-259",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-201",
      "title": "Ontario Grapes-for-Processing Order",
      "citation": "CRC, c 201",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-210",
      "title": "Ontario Greenhouse Vegetable Appointed Shippers\u0027 Procedures (Interprovincial and Export) Regulations",
      "citation": "CRC, c 210",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-212",
      "title": "Ontario Greenhouse Vegetable Handling (Interprovincial and Export) Regulations",
      "citation": "CRC, c 212",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-213",
      "title": "Ontario Greenhouse Vegetable Information (Interprovincial and Export) Regulations",
      "citation": "CRC, c 213",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-208",
      "title": "Ontario Greenhouse Vegetable Marketing (Interprovincial and Export) Regulations",
      "citation": "CRC, c 208",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-207",
      "title": "Ontario Greenhouse Vegetable Order",
      "citation": "CRC, c 207",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-209",
      "title": "Ontario Greenhouse Vegetable Pricing (Interprovincial and Export) Regulations",
      "citation": "CRC, c 209",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-211",
      "title": "Ontario Greenhouse Vegetable Service Charge (Interprovincial and Export) Regulations",
      "citation": "CRC, c 211",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-440",
      "title": "Ontario Hog Charges (Interprovincial and Export) Order",
      "citation": "SOR/96-440",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-181",
      "title": "Ontario Hydro Nuclear Facilities Exclusion from Part III of the Canada Labour Code Regulations (Labour Standards)",
      "citation": "SOR/98-181",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-180",
      "title": "Ontario Hydro Nuclear Facilities Exclusion from Part II of the Canada Labour Code Regulations (Occupational Health and Safety)",
      "citation": "SOR/98-180",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-179",
      "title": "Ontario Hydro Nuclear Facilities Exclusion from Part I of the Canada Labour Code Regulations (Industrial Relations)",
      "citation": "SOR/98-179",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-182",
      "title": "Ontario Hydro Nuclear Facilities Exclusion Regulations (Use of Tobacco)",
      "citation": "SOR/98-182",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-217",
      "title": "Ontario Milk Marketing Levies Order",
      "citation": "CRC, c 217",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-215",
      "title": "Ontario Milk Order",
      "citation": "CRC, c 215",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-62-1",
      "title": "Ontario Onion Growers\u0027 Marketing Order",
      "citation": "SOR/62-1",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-218",
      "title": "Ontario Onion Order",
      "citation": "CRC, c 218",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-418",
      "title": "Ontario Pork Producers\u0027 Marketing Order",
      "citation": "SOR/79-418",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-249",
      "title": "Ontario Review of Parole Ineligibility Rules (Rule 50)",
      "citation": "SOR/2013-249",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-270",
      "title": "Ontario Rules of Practice Respecting Reduction in the Number of Years of Imprisonment Without Eligibility for Parole",
      "citation": "SOR/92-270",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-306",
      "title": "Ontario Sex Offender Information Registration Regulations",
      "citation": "SOR/2004-306",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-328",
      "title": "Ontario Soya-Bean Marketing Levies Order",
      "citation": "SOR/80-328",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-183",
      "title": "Ontario Soya-Bean Order",
      "citation": "SOR/80-183",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-678",
      "title": "Ontario Tender Fruit Order",
      "citation": "SOR/79-678",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-669",
      "title": "Ontario Tender Fruit Service Charge (Interprovincial and Export) Regulations, 1986",
      "citation": "SOR/86-669",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-221",
      "title": "Ontario Turkey Marketing Levies Order",
      "citation": "CRC, c 221",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-220",
      "title": "Ontario Turkey Order",
      "citation": "CRC, c 220",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-860",
      "title": "Ontario Vegetable Growers\u0027 Marketing-for-Processing Order",
      "citation": "SOR/78-860",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-224",
      "title": "Ontario Wheat Marketing (Interprovincial and Export) Regulations",
      "citation": "CRC, c 224",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-223",
      "title": "Ontario Wheat Order",
      "citation": "CRC, c 223",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-192",
      "title": "Regulations on Operational Terms for Rail Level of Services Arbitration",
      "citation": "SOR/2014-192",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-283",
      "title": "Optional Survivor Annuity Regulations",
      "citation": "SOR/2001-283",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2009-46",
      "title": "Order Designating the Indian Residential Schools Truth and Reconciliation Commission as a Department and the Chairperson as the Deputy Head for Purposes of the Act",
      "citation": "SI/2009-46",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-369",
      "title": "Orderly Payment of Debts Regulations",
      "citation": "CRC, c 369",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-338",
      "title": "Organic Products Regulations",
      "citation": "SOR/2006-338",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-176",
      "title": "Organic Products Regulations, 2009",
      "citation": "SOR/2009-176",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-574",
      "title": "Organisation internationale de la Francophonie and the Institut de l\u0027énergie et de l\u0027environnement de la Francophonie Privileges and Immunities Order",
      "citation": "SOR/88-574",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-107",
      "title": "Organization for Economic Co-operation and Development Privileges and Immunities Order",
      "citation": "SOR/91-107",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-350",
      "title": "Organization of American States Privileges and Immunities in Canada Order",
      "citation": "SOR/99-350",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-219",
      "title": "Organizations in the Province of Alberta Exemption Order",
      "citation": "SOR/2004-219",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-220",
      "title": "Organizations in the Province of British Columbia Exemption Order",
      "citation": "SOR/2004-220",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-374",
      "title": "Organizations in the Province of Quebec Exemption Order",
      "citation": "SOR/2003-374",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-348",
      "title": "Origin Regulations (Safeguard Measures in Respect of the People\u0027s Republic of China)",
      "citation": "SOR/2002-348",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-205",
      "title": "Orion Bus Industries Ltd. Remission Order",
      "citation": "SOR/2001-205",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-358",
      "title": "Oshawa Airport Zoning Regulations",
      "citation": "SOR/85-358",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-910",
      "title": "Proclamation Establishing the Oshawa Harbour Commission",
      "citation": "CRC, c 910",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-911",
      "title": "Oshawa Harbour Commission By-laws",
      "citation": "CRC, c 911",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-100",
      "title": "Ottawa International Airport Zoning Regulations",
      "citation": "CRC, c 100",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-231",
      "title": "Ottawa Macdonald-Cartier International Airport Zoning Regulations",
      "citation": "SOR/2009-231",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-88",
      "title": "Outerwear Apparel Remission Order, 1998",
      "citation": "SOR/98-88",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-90",
      "title": "Outerwear Fabrics Remission Order, 1998",
      "citation": "SOR/98-90",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-86",
      "title": "Outerwear Greige Fabrics Remission Order, 1998",
      "citation": "SOR/98-86",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-138",
      "title": "Outward Processing Remission Order (Textiles and Apparel)",
      "citation": "SOR/2008-138",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2008-69",
      "title": "Overpayments of Canada Education Savings Grants Remission Order",
      "citation": "SI/2008-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-99",
      "title": "Overpayments to Certain Members of the Canadian Forces Remission Order",
      "citation": "SI/99-99",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-7",
      "title": "Ozone-Depleting Substances Regulations, 1998",
      "citation": "SOR/99-7",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-270",
      "title": "Pacific Aquaculture Regulations",
      "citation": "SOR/2010-270",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-215",
      "title": "Pacific Fishery Management Area Regulations",
      "citation": "SOR/82-215",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-77",
      "title": "Pacific Fishery Management Area Regulations, 2007",
      "citation": "SOR/2007-77",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-54",
      "title": "Pacific Fishery Regulations, 1993",
      "citation": "SOR/93-54",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-750",
      "title": "Pacific Hake Exemption Notice",
      "citation": "SOR/86-750",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1270",
      "title": "Pacific Pilotage Regulations",
      "citation": "CRC, c 1270",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-583",
      "title": "Pacific Pilotage Tariff Regulations",
      "citation": "SOR/85-583",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-501",
      "title": "Pacific Salmon Commission Privileges and Immunities Order",
      "citation": "SOR/86-501",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-208",
      "title": "Packaging and Transport of Nuclear Substances Regulations",
      "citation": "SOR/2000-208",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-94",
      "title": "Pangnirtung Airport Zoning Regulations",
      "citation": "SOR/2012-94",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-210",
      "title": "Pardon Services Fees Order",
      "citation": "SOR/95-210",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-365",
      "title": "Pari-Mutuel Betting Supervision Regulations",
      "citation": "SOR/91-365",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-72",
      "title": "Pari-Mutuel Payments Order",
      "citation": "SI/83-72",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1247",
      "title": "Parliamentary Secretaries Expenses Regulations",
      "citation": "CRC, c 1247",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-33a",
      "title": "Part-time Work Exclusion Approval Order",
      "citation": "SOR/81-33a",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-201",
      "title": "Passenger Automobile and Light Truck Greenhouse Gas Emission Regulations",
      "citation": "SOR/2010-201",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-219",
      "title": "Passenger Information (Customs) Regulations",
      "citation": "SOR/2003-219",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-147",
      "title": "Passenger Information Regulations (Preclearance Act)",
      "citation": "SOR/2002-147",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-10",
      "title": "Passover Products Remission Order",
      "citation": "SI/91-10",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-253",
      "title": "Passport and Other Travel Document Services Fees Regulations",
      "citation": "SOR/2012-253",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-719",
      "title": "Passport Services Fees Regulations",
      "citation": "CRC, c 719",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-106",
      "title": "Pasta Remission Order 1988",
      "citation": "SI/88-106",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-247",
      "title": "Patented Medicine Prices Review Board Rules of Practice and Procedure",
      "citation": "SOR/2012-247",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-133",
      "title": "Patented Medicines (Notice of Compliance) Regulations",
      "citation": "SOR/93-133",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-688",
      "title": "Patented Medicines Regulations",
      "citation": "SOR/94-688",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-423",
      "title": "Patent Rules",
      "citation": "SOR/96-423",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-67-619",
      "title": "Pay and Allowances Regulations (DSPCA)",
      "citation": "SOR/67-619",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-130",
      "title": "Payments and Settlements Requisitioning Regulations, 1997",
      "citation": "SOR/98-130",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-29",
      "title": "Payments in Lieu of Taxes Regulations",
      "citation": "SOR/81-29",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-239",
      "title": "Payments to Estates Regulations, 1996",
      "citation": "SOR/97-239",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-252",
      "title": "Payments to the Provinces Regulations",
      "citation": "SOR/2007-252",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-105",
      "title": "Pay Television Regulations, 1990",
      "citation": "SOR/90-105",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-273",
      "title": "PCB Regulations",
      "citation": "SOR/2008-273",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-109",
      "title": "PCB Waste Export Regulations, 1996",
      "citation": "SOR/97-109",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-188",
      "title": "Peace River Airport Zoning Regulations",
      "citation": "SOR/85-188",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-606",
      "title": "Pear Stabilization Regulations, 1982",
      "citation": "SOR/83-606",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-549",
      "title": "Pêcheurs Unis du Québec Regulations",
      "citation": "SOR/83-549",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-237",
      "title": "P.E.I. Vegetable Directed Sales (Interprovincial and Export) Regulations",
      "citation": "CRC, c 237",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-240",
      "title": "P.E.I. Vegetable Information (Interprovincial and Export) Regulations",
      "citation": "CRC, c 240",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-238",
      "title": "P.E.I. Vegetable Licensing (Interprovincial and Export) Regulations",
      "citation": "CRC, c 238",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-239",
      "title": "P.E.I. Vegetable Seizure (Interprovincial and Export) Regulations",
      "citation": "CRC, c 239",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-125",
      "title": "Pembroke Airport Zoning Regulations",
      "citation": "SOR/88-125",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-65-262",
      "title": "Pensionable Service Order",
      "citation": "SOR/65-262",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-620",
      "title": "Pension and Allowance Adjustment Regulations",
      "citation": "SOR/91-620",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-390",
      "title": "Pension Appeals Board Rules of Procedure (Benefits)",
      "citation": "CRC, c 390",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-612",
      "title": "Pension Benefits Division Regulations",
      "citation": "SOR/94-612",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-19",
      "title": "Pension Benefits Standards Regulations, 1985",
      "citation": "SOR/87-19",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-48",
      "title": "Pension Diversion Regulations",
      "citation": "SOR/84-48",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1581",
      "title": "Pensioners Training Regulations",
      "citation": "CRC, c 1581",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-58-417",
      "title": "Pension Increase Regulations, 1958",
      "citation": "SOR/58-417",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-445",
      "title": "Pension Plan Transfer Agreements Regulations",
      "citation": "SOR/98-445",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-101",
      "title": "Penticton Airport Zoning Regulations",
      "citation": "CRC, c 101",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-65",
      "title": "Proclamation Requesting the People of Canada to Observe June 23 of Every Year as a National Day of Remembrance for Victims of Terrorism",
      "citation": "SI/2005-65",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-778",
      "title": "Perfluorinated Ion-Exchange Membranes Remission Order",
      "citation": "CRC, c 778",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-178",
      "title": "Perfluorooctane Sulfonate and its Salts and Certain Other Compounds Regulations",
      "citation": "SOR/2008-178",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-15",
      "title": "Regulations Adding Perfluorooctane Sulfonate and Its Salts to the Virtual Elimination List",
      "citation": "SOR/2009-15",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-252",
      "title": "Period for Entering into an Agreement for the Purpose of Jointly Establishing a Review Panel Regulations",
      "citation": "SOR/2006-252",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-375",
      "title": "Regulations Establishing Periods of Probation and Periods of Notice of Termination of Employment During Probation",
      "citation": "SOR/2005-375",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-143",
      "title": "Regulations Establishing the Periods Within Which Eligible Authors, Eligible Performers and Eligible Makers not Represented by Collective Societies Can Claim Private Copying Remuneration",
      "citation": "SOR/2013-143",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-164",
      "title": "Regulations Establishing the Period Within Which Owners of Copyright not Represented by Collective Societies Can Claim Retransmission Royalties",
      "citation": "SOR/97-164",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-140",
      "title": "Permits Authorizing an Activity Affecting Listed Wildlife Species Regulations",
      "citation": "SOR/2013-140",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-107",
      "title": "Persistence and Bioaccumulation Regulations",
      "citation": "SOR/2000-107",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-265",
      "title": "Personal Health Information Custodians in New Brunswick Exemption Order",
      "citation": "SOR/2011-265",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-72",
      "title": "Personal Health Information Custodians in Newfoundland and Labrador Exemption Order",
      "citation": "SI/2012-72",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-42",
      "title": "Personnel Training for the Assistance of Persons with Disabilities Regulations",
      "citation": "SOR/94-42",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-148",
      "title": "Regulations Designating Persons and Categories of Persons — Other Than Travellers Destined for the United States — Who May Enter a Preclearance Area",
      "citation": "SOR/2002-148",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-418",
      "title": "Persons Authorized to Account for Casual Goods Regulations",
      "citation": "SOR/95-418",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-15",
      "title": "Order Authorizing Persons Specified Therein to Be Parties to Certain Commercial Arrangements and Providing Specific Directives to the Vancouver Port Authority and the Fraser River Port Authority",
      "citation": "SOR/2006-15",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-328",
      "title": "Order Authorizing Certain Persons to Be a Party to Certain Commercial Arrangements and Providing Specific Directives to the Vancouver Port Authority and the Fraser River Port Authority",
      "citation": "SOR/2005-328",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1140",
      "title": "P.E.S.R.A. Regulations and Rules of Procedure",
      "citation": "SOR/86-1140",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-260",
      "title": "Pest Control Products Incident Reporting Regulations",
      "citation": "SOR/2006-260",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1253",
      "title": "Pest Control Products Regulations",
      "citation": "CRC, c 1253",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-124",
      "title": "Pest Control Products Regulations",
      "citation": "SOR/2006-124",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-261",
      "title": "Pest Control Products Sales and Information Reporting Regulations",
      "citation": "SOR/2006-261",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1254",
      "title": "Pesticide Residue Compensation Regulations",
      "citation": "CRC, c 1254",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-123",
      "title": "Peterborough Airport Zoning Regulations",
      "citation": "SOR/94-123",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-61",
      "title": "Pet Food for Testing Remission Order",
      "citation": "SI/86-61",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-469",
      "title": "Petro-Canada—Inter-City Gas Corporation Transactions Authorization Order",
      "citation": "SOR/89-469",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-714",
      "title": "Petro-Canada Limited Pension Protection Regulations",
      "citation": "SOR/91-714",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-75",
      "title": "Petro-Canada Limited Transactions Authorization Order, 1991",
      "citation": "SOR/91-75",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-58",
      "title": "Petro-Canada Transactions Authorization Order, 1990",
      "citation": "SOR/90-58",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-503",
      "title": "Petroleum and Gas Revenue Tax Regulations",
      "citation": "SOR/82-503",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-461",
      "title": "Petroleum Incentives Program Advance Ruling Fees Order",
      "citation": "SOR/83-461",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-828",
      "title": "Petroleum Refinery Liquid Effluent Regulations",
      "citation": "CRC, c 828",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-337",
      "title": "Pharmaceutical Industry Development Assistance Regulations",
      "citation": "CRC, c 337",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-501",
      "title": "Phosphorus in Certain Cleaning Products Regulations",
      "citation": "SOR/89-501",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-501",
      "title": "Phosphorus in Certain Cleaning Products Regulations",
      "citation": "SOR/89-501",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-298",
      "title": "Phthalates Regulations",
      "citation": "SOR/2010-298",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-147",
      "title": "Regulations Designating Physical Activities",
      "citation": "SOR/2012-147",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-135",
      "title": "Phytophthora Ramorum Compensation Regulations",
      "citation": "SOR/2007-135",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-212",
      "title": "Pickering Airport Site Zoning Regulations",
      "citation": "SOR/2004-212",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-297",
      "title": "Order Declaring the Pickering Lands as an Airport Site",
      "citation": "SOR/2001-297",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-787",
      "title": "Pipeline Arbitration Committee Procedure Rules, 1986",
      "citation": "SOR/86-787",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-102",
      "title": "Pitt Meadows Airport Zoning Regulations",
      "citation": "CRC, c 102",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-170",
      "title": "Place of Supply (GST/HST) Regulations",
      "citation": "SOR/2001-170",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-594",
      "title": "Plant Breeders\u0027 Rights Regulations",
      "citation": "SOR/91-594",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-212",
      "title": "Plant Protection Regulations",
      "citation": "SOR/95-212",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-932",
      "title": "Playpens Regulations",
      "citation": "CRC, c 932",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-661",
      "title": "Pleasure Craft Sewage Pollution Prevention Regulations",
      "citation": "SOR/91-661",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-211",
      "title": "Plum Pox Virus Compensation Regulations",
      "citation": "SOR/2001-211",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-131",
      "title": "Plum Pox Virus Compensation Regulations, 2004",
      "citation": "SOR/2005-131",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-76",
      "title": "Point Lepreau, New Brunswick Nuclear Facility Exclusion Regulations (Parts I, II and III of the Canada Labour Code and the Non-Smokers\u0027 Health Act)",
      "citation": "SOR/2008-76",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-409",
      "title": "Polar Bear Pass Withdrawal Order",
      "citation": "SOR/84-409",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-97",
      "title": "Proclamation Declaring the last Sunday in September of each year to be \"Police and Peace Officers\u0027 National Memorial Day\"",
      "citation": "SI/98-97",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-234",
      "title": "Policyholders Disclosure Regulations",
      "citation": "SOR/2010-234",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-373",
      "title": "Political Activities Regulations",
      "citation": "SOR/2005-373",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-351",
      "title": "Pollutant Discharge Reporting Regulations, 1995",
      "citation": "SOR/95-351",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1458",
      "title": "Pollutant Substances Pollution Prevention Regulations",
      "citation": "CRC, c 1458",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-218",
      "title": "Polybrominated Diphenyl Ethers Regulations",
      "citation": "SOR/2008-218",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-93",
      "title": "Pond Inlet Airport Zoning Regulations",
      "citation": "SOR/2012-93",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-222",
      "title": "Pooled Registered Pension Plans Regulations",
      "citation": "SOR/2012-222",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-294",
      "title": "Pooled Registered Pension Plans Regulations",
      "citation": "SOR/2012-294",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-676",
      "title": "Population Recovery Adjustment Payments Regulations",
      "citation": "SOR/82-676",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-915",
      "title": "Port Alberni Harbour Small Vessel Facilities Tariff By-law",
      "citation": "CRC, c 915",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-101",
      "title": "Port Authorities Management Regulations",
      "citation": "SOR/99-101",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-55",
      "title": "Port Authorities Operations Regulations",
      "citation": "SOR/2000-55",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-484",
      "title": "Port Hardy Airport Zoning Regulations",
      "citation": "SOR/91-484",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-215",
      "title": "Order Transferring Certain Portions from the Department of Citizenship and Immigration to the Canada Border Services Agency",
      "citation": "SI/2003-215",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-216",
      "title": "Order Transferring Certain Portions of the Canada Customs and Revenue Agency to the Canada Border Services Agency",
      "citation": "SI/2003-216",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-60",
      "title": "Portions of the Canada Ports Corporation Divestiture Regulations",
      "citation": "SOR/2000-60",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-291",
      "title": "Portions of the Department of Agriculture and Agri-Food Divestiture Regulations",
      "citation": "SOR/2010-291",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-3",
      "title": "Portions of the Department of Citizenship and Immigration Divestiture Regulations",
      "citation": "SOR/99-3",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-209",
      "title": "Order Transferring Certain Portions of the Department of Foreign Affairs to the Department of International Trade",
      "citation": "SI/2003-209",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-286",
      "title": "Portions of the Department of Health Divestiture Regulations",
      "citation": "SOR/2003-286",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-203",
      "title": "Order Transferring Certain Portions of the Department of Human Resources Development to the Department of Human Resources and Skills Development",
      "citation": "SI/2003-203",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-210",
      "title": "Order Transferring Certain Portions of the Department of Industry to the Department of International Trade",
      "citation": "SI/2003-210",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-230",
      "title": "Portions of the Department of National Defence Divestiture Regulations",
      "citation": "SOR/98-230",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-231",
      "title": "Portions of the Department of Public Works and Government Services Divestiture Regulations",
      "citation": "SOR/98-231",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-217",
      "title": "Order Transferring Certain Portions of the Operations Branch of the Canadian Food Inspection Agency to the Canada Border Services Agency",
      "citation": "SI/2003-217",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-43",
      "title": "Order Transferring Certain Portions of the Public Service Commission to the Canada School of Public Service",
      "citation": "SI/2004-43",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-44",
      "title": "Order Transferring Certain Portions of the Public Service Commission to the Public Service Human Resources Management Agency of Canada",
      "citation": "SI/2004-44",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1336",
      "title": "Designation of Certain Portions of the Public Service Order",
      "citation": "CRC, c 1336",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-247",
      "title": "Portions of the Royal Canadian Mounted Police Divestiture Regulations",
      "citation": "SOR/99-247",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-222",
      "title": "Order Transferring Portions of the Treasury Board Secretariat to the Public Service Human Resources Management Agency of Canada",
      "citation": "SI/2003-222",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-154",
      "title": "Port Wardens Tariff",
      "citation": "SOR/79-154",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-4",
      "title": "Order Designating the Position of Deputy Minister of National Defence to the Position of Deputy Head in Respect of the Communications Security Establishment",
      "citation": "SI/2006-4",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-205",
      "title": "Regulations Respecting the Possession of Non-duty-paid Packaged Alcohol",
      "citation": "SOR/2003-205",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-203",
      "title": "Regulations Respecting the Possession of Tobacco Products That Are Not Stamped",
      "citation": "SOR/2003-203",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-748",
      "title": "Postage Meter Regulations",
      "citation": "SOR/83-748",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-220",
      "title": "Postage Meter Regulations, 2010",
      "citation": "SOR/2010-220",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-181",
      "title": "Postal Imports Remission Order",
      "citation": "SI/85-181",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-259",
      "title": "Postal Services Interruption Regulations",
      "citation": "SOR/87-259",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1288",
      "title": "Posting Abroad of Letter-Post Items Regulations",
      "citation": "CRC, c 1288",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-1955-p-2561",
      "title": "Post Office Savings Bank Regulations",
      "citation": "CRC 1955, p 2561",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1105",
      "title": "Potable Water Regulations for Common Carriers",
      "citation": "CRC, c 1105",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-731",
      "title": "1985 Potatoes Stabilization Regulations",
      "citation": "SOR/87-731",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-829",
      "title": "Potato Processing Plant Liquid Effluent Regulations",
      "citation": "CRC, c 829",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-186",
      "title": "Potato Production and Sale (Central Saanich) Restriction Regulations",
      "citation": "SOR/82-186",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-445",
      "title": "Potato Stabilization Regulations, 1977",
      "citation": "SOR/78-445",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-55",
      "title": "Potato Stabilization Regulations, 1978",
      "citation": "SOR/80-55",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-295",
      "title": "Potato Stabilization Regulations, 1979",
      "citation": "SOR/81-295",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-451",
      "title": "Potato Wart Compensation Regulations",
      "citation": "SOR/2001-451",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-211",
      "title": "Potato Wart Compensation Regulations, 2003",
      "citation": "SOR/2004-211",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-500",
      "title": "Power Line Crossing Regulations",
      "citation": "SOR/95-500",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-228",
      "title": "Order Transfering the Powers, Duties and Functions of the International Expositions Division to the Minister of Communication",
      "citation": "SI/92-228",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-310",
      "title": "Prairie Dog and Certain Other Rodents Importation Prohibition Regulations",
      "citation": "SOR/2003-310",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-442",
      "title": "Prairie Farm Rehabilitation Administration Employees Regulations",
      "citation": "SOR/78-442",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1301",
      "title": "Prairie Grain Advance Payments Regulations",
      "citation": "CRC, c 1301",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1303",
      "title": "Precious Metals Marking Regulations",
      "citation": "CRC, c 1303",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-359",
      "title": "Precursor Control Regulations",
      "citation": "SOR/2002-359",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-229",
      "title": "Regulations Exempting Certain Precursors and Controlled Substances from the Application of the Controlled Drugs and Substances Act",
      "citation": "SOR/97-229",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-12",
      "title": "Preliminary Screening Requirement Regulations",
      "citation": "SOR/99-12",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-209",
      "title": "Prepaid Payment Products Regulations",
      "citation": "SOR/2013-209",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-7",
      "title": "Regulations Respecting Prescribed Brands of Manufactured Tobacco and Prescribed Cigarettes",
      "citation": "SOR/2011-7",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-46",
      "title": "Prescribed Classes of Persons in Respect of Diversion of Imported Goods Regulations",
      "citation": "SOR/98-46",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-53",
      "title": "Prescribed Deposits (Authorized Foreign Banks) Regulations",
      "citation": "SOR/2000-53",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-54",
      "title": "Prescribed Deposits (Banks without Deposit Insurance) Regulations",
      "citation": "SOR/2000-54",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-67",
      "title": "Prescribed Deposits (Retail Associations Without Deposit Insurance) Regulations",
      "citation": "SOR/2008-67",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-66",
      "title": "Prescribed Deposits (Trust and Loan Companies Without Deposit Insurance) Regulations",
      "citation": "SOR/2008-66",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-230",
      "title": "Prescribed Entities and Classes of Mortgages and Hypothecs Regulations",
      "citation": "SOR/2011-230",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-167",
      "title": "Prescribed Group of Consumers Regulations",
      "citation": "SOR/2008-167",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-148",
      "title": "Prescribed Information for the Description of a Designated Project Regulations",
      "citation": "SOR/2012-148",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-68",
      "title": "Prescribed Persons and Organizations Regulations",
      "citation": "SOR/96-68",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-100",
      "title": "Prescribed Products Regulations",
      "citation": "SOR/2011-100",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-391",
      "title": "Prescribed Province Pension Regulations",
      "citation": "CRC, c 391",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-280",
      "title": "Regulations Prescribing an Oath of Secrecy",
      "citation": "SOR/2014-280",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-280",
      "title": "Proclamation Prescribing the Composition, Dimensions and Design of a One Dollar Base Metal Coin",
      "citation": "SOR/94-280",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-417",
      "title": "Proclamation Prescribing the Composition, Dimensions and Design of a One Dollar Base Metal Coin",
      "citation": "SOR/92-417",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-48",
      "title": "Proclamation Prescribing the Composition, Dimensions and Design of a One Dollar Base Metal Coin",
      "citation": "SOR/95-48",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-337",
      "title": "Proclamation Prescribing the Designs of Gold, Platinum and Silver Coins",
      "citation": "SOR/89-337",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-323",
      "title": "Presentation of Persons (2003) Regulations",
      "citation": "SOR/2003-323",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-68",
      "title": "Order Designating the President of Shared Services Canada as Deputy Head in Respect of that Entity",
      "citation": "SI/2011-68",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-13",
      "title": "Order Designating the President of the Queen\u0027s Privy Council for Canada as the appropriate Minister for The Jacques-Cartier and Champlain Bridges Inc",
      "citation": "SI/2014-13",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-63",
      "title": "Order Designating the President of the Queen\u0027s Privy Council for Canada to be the Minister for the purposes of the Act",
      "citation": "SI/2014-63",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-88",
      "title": "Order Transferring from the President of the Queen\u0027s Privy Council for Canada to the Deputy Prime Minister and Minister of Public Safety and Emergency Preparedness the Control and Supervision of the Office of Indian Residential Schools Resolution of Canada and Ordering the Deputy Prime Minister to Preside Over the Office",
      "citation": "SI/2004-88",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-90",
      "title": "Order Transferring from the President of the Queen\u0027s Privy Council for Canada to the President of the Treasury Board the Control and Supervision of the Public Service Human Resources Management Agency of Canada and Ordering the President of the Treasury Board to Preside Over the Agency",
      "citation": "SI/2004-90",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-221",
      "title": "Order Transferring from the President of the Treasury Board to the President of the Queen\u0027s Privy Council for Canada the Control and Supervision of the Public Service Human Resources Management Agency of Canada",
      "citation": "SI/2003-221",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-86",
      "title": "Regulations for the Prevention of Pollution from Ships and for Dangerous Chemicals",
      "citation": "SOR/2007-86",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-62",
      "title": "Order Designating the Prime Minister as Presiding Minister for the Public Appointments Commission Secretariat",
      "citation": "SI/2006-62",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-153",
      "title": "Prime Minister Authority to Prescribe Fees Order (Canadian Heraldic Authority)",
      "citation": "SI/90-153",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-46",
      "title": "Order Transferring from the Prime Minister to the President of the Queen\u0027s Privy Council for Canada the Control and Supervision of the Canadian Intergovernmental Conference Secretariat",
      "citation": "SI/2006-46",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-51",
      "title": "Order transferring from the Prime Minister to the President of the Treasury Board the control and supervision of The Leadership Network",
      "citation": "SI/2001-51",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-705",
      "title": "Prince Albert Airport Zoning Regulations",
      "citation": "SOR/87-705",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2007-61",
      "title": "Prince Albert National Park of Canada Land Rents Remission Order",
      "citation": "SI/2007-61",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-382",
      "title": "Prince Edward Island Cattle Marketing Levies Order",
      "citation": "SOR/93-382",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-292",
      "title": "Prince Edward Island Cattle Order",
      "citation": "SOR/88-292",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-225",
      "title": "Prince Edward Island Chicken Order",
      "citation": "CRC, c 225",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-109",
      "title": "Prince Edward Island – Criminal Appeal Rules of Court",
      "citation": "SI/2011-109",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-383",
      "title": "Prince Edward Island Criminal Rule of Practice Respecting Reduction in the Number of Years of Imprisonment Without Eligibility for Parole",
      "citation": "SOR/92-383",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-226",
      "title": "Prince Edward Island Egg Order",
      "citation": "CRC, c 226",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-277",
      "title": "Order Designating Prince Edward Island for the Purposes of the Criminal Interest Rate Provisions of the Criminal Code",
      "citation": "SOR/2014-277",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-229",
      "title": "Prince Edward Island Hog Marketing Levies Order",
      "citation": "CRC, c 229",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-228",
      "title": "Prince Edward Island Hog Order",
      "citation": "CRC, c 228",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-628",
      "title": "Prince Edward Island Milk Order",
      "citation": "SOR/94-628",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-232",
      "title": "Prince Edward Island Pedigreed Seed Order",
      "citation": "CRC, c 232",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-234",
      "title": "Prince Edward Island Potato Marketing Levies Order",
      "citation": "CRC, c 234",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-233",
      "title": "Prince Edward Island Potato Order",
      "citation": "CRC, c 233",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-307",
      "title": "Prince Edward Island Sex Offender Information Registration Regulations",
      "citation": "SOR/2004-307",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-235",
      "title": "Prince Edward Island Tobacco Order",
      "citation": "CRC, c 235",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-236",
      "title": "Prince Edward Island Vegetable Order",
      "citation": "CRC, c 236",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-103",
      "title": "Prince George Airport Zoning Regulations",
      "citation": "CRC, c 103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-180",
      "title": "Principal Protected Notes Regulations",
      "citation": "SOR/2008-180",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-360",
      "title": "Printed Material for Foreign Carriers Remission Order, 1995",
      "citation": "SOR/95-360",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-134",
      "title": "Prisoner-of-War Status Determination Regulations",
      "citation": "SOR/91-134",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-41",
      "title": "Prison Manufactured or Produced Goods Regulations, 1998",
      "citation": "SOR/98-41",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-553",
      "title": "Privacy Act Extension Order No. 1",
      "citation": "SOR/83-553",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-206",
      "title": "Privacy Act Extension Order, No. 2",
      "citation": "SOR/89-206",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-114",
      "title": "Privacy Act Heads of Government Institutions Designation Order",
      "citation": "SI/83-114",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-508",
      "title": "Privacy Regulations",
      "citation": "SOR/83-508",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-335",
      "title": "Private Buoy Regulations",
      "citation": "SOR/99-335",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1317",
      "title": "Privileges and Immunities Accession Order (United Nations)",
      "citation": "CRC, c 1317",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-409",
      "title": "Order respecting Privileges and Immunities in relation to the Preparatory Commission for the Comprehensive Nuclear Test-Ban Treaty Organization and its Provisional Technical Secretariat",
      "citation": "SOR/99-409",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1318",
      "title": "Privileges and Immunities (International Labour Organization) Order",
      "citation": "CRC, c 1318",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-612",
      "title": "Privileges and Immunities of the European Bank for Reconstruction and Development Order",
      "citation": "SOR/93-612",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-501",
      "title": "Privileges and Immunities of the Secretariat of the Convention on Biological Diversity Order",
      "citation": "SOR/97-501",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-401",
      "title": "Order Respecting the Privileges and Immunities of the United Nations Educational, Scientific and Cultural Organization and Its Institute for Statistics",
      "citation": "SOR/2002-401",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-45",
      "title": "Order Transferring from Privy Council Office to the Department of Canadian Heritage the Control and Supervision of the Official Languages Secretariat",
      "citation": "SI/2006-45",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-74",
      "title": "Order Transferring from the Privy Council Office to the Department of Human Resources and Skills Development the Control and Supervision of the Policy Research Initiative",
      "citation": "SI/2006-74",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-86",
      "title": "Order Transferring from the Privy Council Office to the Department of Indian Affairs and Northern Development the Control and Supervision of the Aboriginal Affairs Secretariat",
      "citation": "SI/2004-86",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-51",
      "title": "Order Transferring from the Privy Council Office to the Department of Indian Affairs and Northern Development the Control and Supervision of the Aboriginal Affairs Secretariat",
      "citation": "SI/2006-51",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-72",
      "title": "Order Transferring from the Privy Council Office to the Department of Industry the Control and Supervision of the National Science Advisor Secretariat",
      "citation": "SI/2006-72",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-73",
      "title": "Order Transferring from the Privy Council Office to the Department of Public Safety and Emergency Preparedness the Control and Supervision of the Borders Task Force",
      "citation": "SI/2006-73",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-103",
      "title": "Order Transferring from the Privy Council Office to the Office of Infrastructure of Canada the Control and Supervision of the Cities Secretariat",
      "citation": "SI/2004-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-101",
      "title": "Order Transferring from the Privy Council Office to the Public Service Human Resources Management Agency of Canada the Control and Supervision of the Office of the Senior Advisor Responsible for Diversity and Special Projects",
      "citation": "SI/2006-101",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-76",
      "title": "Order Transferring from the Privy Council Office to the Treasury Board the Control and Supervision of the Regional Offices in the Communications and Consultations Secretariat",
      "citation": "SI/2006-76",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-75",
      "title": "Order Transferring from the Privy Council Office to the Treasury Board the Control and Supervision of the Regulatory Affairs and Orders in Council Secretariat Except Orders in Council Division",
      "citation": "SI/2006-75",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-126",
      "title": "Order Amalgamating and Combining the Privy Council Office with the Federal-Provincial Relations Office under the Prime Minister",
      "citation": "SI/93-126",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-292",
      "title": "Proceeds of Crime (Money Laundering) and Terrorist Financing Administrative Monetary Penalties Regulations",
      "citation": "SOR/2007-292",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-121",
      "title": "Proceeds of Crime (Money Laundering) and Terrorist Financing Registration Regulations",
      "citation": "SOR/2007-121",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-184",
      "title": "Proceeds of Crime (Money Laundering) and Terrorist Financing Regulations",
      "citation": "SOR/2002-184",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-317",
      "title": "Proceeds of Crime (Money Laundering) and Terrorist Financing Suspicious Transaction Reporting Regulations",
      "citation": "SOR/2001-317",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-290",
      "title": "Processed Egg Regulations",
      "citation": "CRC, c 290",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-291",
      "title": "Processed Products Regulations",
      "citation": "CRC, c 291",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-254",
      "title": "Processing and Distribution of Semen for Assisted Conception Regulations",
      "citation": "SOR/96-254",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-143",
      "title": "Procurement Ombudsman Regulations",
      "citation": "SOR/2008-143",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-41",
      "title": "Procurement Review Board Regulations",
      "citation": "SOR/89-41",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-254",
      "title": "Products Containing Mercury Regulations",
      "citation": "SOR/2014-254",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-436",
      "title": "Programming Undertaking Regulations",
      "citation": "SOR/93-436",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-155",
      "title": "Program Registration Fees Order",
      "citation": "SOR/84-155",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-411",
      "title": "Prohibited Activities Respecting Real Property (Foreign Banks) Regulations",
      "citation": "SOR/2001-411",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-377",
      "title": "Regulations Prohibiting Deployments into the Executive Group",
      "citation": "SOR/2005-377",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-99",
      "title": "Prohibition of Certain Toxic Substances Regulations, 2003",
      "citation": "SOR/2003-99",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-41",
      "title": "Prohibition of Certain Toxic Substances Regulations, 2005",
      "citation": "SOR/2005-41",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-285",
      "title": "Prohibition of Certain Toxic Substances Regulations, 2012",
      "citation": "SOR/2012-285",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-491",
      "title": "Projects Outside Canada Environmental Assessment Regulations",
      "citation": "SOR/96-491",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-178",
      "title": "Promotion of Tobacco Products and Accessories Regulations (Prohibited Terms)",
      "citation": "SOR/2011-178",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-52",
      "title": "Proof of Origin of Imported Goods Regulations",
      "citation": "SOR/98-52",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-110",
      "title": "Propane Dispenser Specifications",
      "citation": "SI/91-110",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-493",
      "title": "Property Assessment and Taxation (Railway Right-of-Way) Regulations",
      "citation": "SOR/2001-493",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-66",
      "title": "Property Supplied by Auction (GST/HST) Regulations",
      "citation": "SOR/2001-66",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-65",
      "title": "Proposed Acquisition of Increased Interest in Greyhound Lines of Canada Ltd. Exemption Order",
      "citation": "SOR/93-65",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-210",
      "title": "Prosecution of Provincial Offences Regulations (Donkin Coal Block)",
      "citation": "SOR/2008-210",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-145",
      "title": "Prospector\u0027s Assistance Terms and Conditions Order",
      "citation": "SOR/81-145",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-318",
      "title": "Prospectus (Banks and Bank Holding Companies) Regulations",
      "citation": "SOR/2006-318",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-412",
      "title": "Prospectus (Banks and Bank Holding Companies) Regulations",
      "citation": "SOR/2001-412",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-319",
      "title": "Prospectus (Cooperative Credit Associations) Regulations",
      "citation": "SOR/2006-319",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-340",
      "title": "Prospectus (Cooperative Credit Associations) Regulations",
      "citation": "SOR/95-340",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-414",
      "title": "Prospectus Exemptions (Bank Holding Companies) Regulations",
      "citation": "SOR/2001-414",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-73",
      "title": "Prospectus Exemptions (Banks) Regulations",
      "citation": "SOR/94-73",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-74",
      "title": "Prospectus Exemptions (Cooperative Credit Associations) Regulations",
      "citation": "SOR/94-74",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-75",
      "title": "Prospectus Exemptions (Insurance Companies) Regulations",
      "citation": "SOR/94-75",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-418",
      "title": "Prospectus Exemptions (Insurance Holding Companies) Regulations",
      "citation": "SOR/2001-418",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-76",
      "title": "Prospectus Exemptions (Trust and Loan Companies) Regulations",
      "citation": "SOR/94-76",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-265",
      "title": "Prospectus (Federal Credit Unions) Regulations",
      "citation": "SOR/2012-265",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-320",
      "title": "Prospectus (Insurance Companies and Insurance Holding Companies) Regulations",
      "citation": "SOR/2006-320",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-413",
      "title": "Prospectus (Insurance Companies and Insurance Holding Companies) Regulations",
      "citation": "SOR/2001-413",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-321",
      "title": "Prospectus (Trust and Loan Companies) Regulations",
      "citation": "SOR/2006-321",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-342",
      "title": "Prospectus (Trust and Loan Companies) Regulations",
      "citation": "SOR/95-342",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-463",
      "title": "Protection for the Income of Milk Producers Regulations (1994)",
      "citation": "SOR/94-463",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-295",
      "title": "Protection for the Income of Milk Producers Regulations (1996)",
      "citation": "SOR/96-295",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-312",
      "title": "Protection for the Income of Milk Producers Regulations (1993-1)",
      "citation": "SOR/93-312",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-313",
      "title": "Protection for the Income of Milk Producers Regulations (1993-2)",
      "citation": "SOR/93-313",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-294",
      "title": "Protection for the Income of Milk Producers Regulations (1995-1)",
      "citation": "SOR/95-294",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-295",
      "title": "Protection for the Income of Milk Producers Regulations (1995-2)",
      "citation": "SOR/95-295",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-352",
      "title": "Protection of Assets (Banks) Regulations",
      "citation": "SOR/92-352",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-354",
      "title": "Protection of Assets (Cooperative Credit Associations) Regulations",
      "citation": "SOR/92-354",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-81",
      "title": "Protection of Assets (Fraternal Benefit Societies) Regulations",
      "citation": "SOR/94-81",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-353",
      "title": "Protection of Assets (Insurance Companies) Regulations",
      "citation": "SOR/92-353",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-350",
      "title": "Protection of Assets (Trust and Loan Companies) Regulations",
      "citation": "SOR/92-350",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-346",
      "title": "Protection of Passenger Information Regulations",
      "citation": "SOR/2005-346",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-145",
      "title": "Protection of Personal Information Regulations",
      "citation": "SOR/78-145",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-440",
      "title": "Protection of Privacy Regulations",
      "citation": "CRC, c 440",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-231",
      "title": "Protection of Residential Mortgage or Hypothecary Insurance Regulations",
      "citation": "SOR/2012-231",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-202",
      "title": "Emergency Order for the Protection of the Greater Sage-Grouse",
      "citation": "SOR/2013-202",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-3",
      "title": "Proclamation Giving Notice that the Protocol Amending the Convention between the Government of Canada and the Government of the Republic of Singapore for the Avoidance of Double Taxation and the Prevention of Fiscal Evasion with respect to Taxes on Income Came into Force on August 31, 2012",
      "citation": "SI/2013-3",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-17",
      "title": "Proclamation Giving Notice that the Protocol Amending the Convention between the Government of Canada and the Swiss Federal Council for the Avoidance of Double Taxation with respect to Taxes on Income and on Capital, Came into Force on December 16, 2011",
      "citation": "SI/2012-17",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-288",
      "title": "Order Designating the Province of Manitoba for the Purposes of the Definition “applicable guidelines” in Subsection 2(1) of the Divorce Act",
      "citation": "SOR/98-288",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-256",
      "title": "Order Designating the Province of New Brunswick for the Purposes of the Definition “applicable guidelines” in Subsection 2(1) of the Divorce Act",
      "citation": "SOR/98-256",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-9",
      "title": "Order Designating the Province of Prince Edward Island for the Purposes of the Definition “applicable guidelines” in Subsection 2(1) of the Divorce Act",
      "citation": "SOR/98-9",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-237",
      "title": "Order Designating the Province of Quebec for the Purposes of the Definition “applicable guidelines” in Subsection 2(1) of the Divorce Act",
      "citation": "SOR/97-237",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-104",
      "title": "Provincial Court of British Columbia Criminal Caseflow Management Rules",
      "citation": "SI/99-104",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-288",
      "title": "Regulations Prescribing a Provincial Tax",
      "citation": "SOR/88-288",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-721",
      "title": "Provision of Materiel and Services to Foreign Ships Order",
      "citation": "CRC, c 721",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-156",
      "title": "Order Declaring that all Provisions of Part X of the Financial Administration Act Apply to Parc Downsview Park Inc.",
      "citation": "SI/2003-156",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-608",
      "title": "Prune Stabilization Regulations, 1982",
      "citation": "SOR/83-608",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-348",
      "title": "P.S.S.R.B. Regulations and Rules of Procedure, 1993",
      "citation": "SOR/93-348",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-133",
      "title": "Public Accountability Statements (Banks, Insurance Companies, Trust and Loan Companies) Regulations",
      "citation": "SOR/2002-133",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-203",
      "title": "Public Agents Firearms Regulations",
      "citation": "SOR/98-203",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-64",
      "title": "Order Designating the Public Appointments Commission Secretariat as a Department and the Executive Director as Deputy Head",
      "citation": "SI/2006-64",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-406",
      "title": "Publication of Standards Regulations",
      "citation": "SOR/95-406",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1367",
      "title": "Publication of Statutes Regulations",
      "citation": "CRC, c 1367",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-43",
      "title": "Publications Supplied by a Registrant (GST/HST) Regulations",
      "citation": "SOR/91-43",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-50",
      "title": "Order Transferring to the Public Health Agency of Canada the Control and Supervision of Certain Portions of the Federal Public Administration in the Department of Health known as the International Affairs Directorate and the Emergency Preparedness and Response Unit",
      "citation": "SI/2012-50",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-40",
      "title": "Order Transferring to the Public Health Agency of Canada the Control and Supervision of the Portion of the Federal Public Administration in the Department of Health known as the Travelling Public Program Unit",
      "citation": "SI/2013-40",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-276",
      "title": "Public Inquiry (Authorized Foreign Banks) Rules",
      "citation": "SOR/99-276",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-722",
      "title": "Public Lands Licensing Order",
      "citation": "CRC, c 722",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-13",
      "title": "Public Lands Mineral Regulations",
      "citation": "SOR/96-13",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1326",
      "title": "Public Lands Oil and Gas Regulations",
      "citation": "CRC, c 1326",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-117",
      "title": "Designated Public Office Holder Regulations",
      "citation": "SOR/2008-117",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-200",
      "title": "Public Office Holders and Reporting Public Office Holders under Section 62.2 of the Conflict of Interest Act",
      "citation": "SOR/2014-200",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-466",
      "title": "Regulations Prescribing Public Officers",
      "citation": "SOR/98-466",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-723",
      "title": "Public Officers Guarantee Regulations",
      "citation": "CRC, c 723",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-134",
      "title": "Public Opinion Research Contract Regulations",
      "citation": "SOR/2007-134",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-154",
      "title": "Public Ports and Public Port Facilities Regulations",
      "citation": "SOR/2001-154",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-745",
      "title": "Public Property Loan Regulations",
      "citation": "SOR/92-745",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-77",
      "title": "Public Sector Pension Investment Board Regulations",
      "citation": "SOR/2000-77",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-170",
      "title": "Public Servants Disclosure Protection Tribunal Rules of Procedure",
      "citation": "SOR/2011-170",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1332",
      "title": "Public Servants Inventions Regulations",
      "citation": "CRC, c 1332",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-37",
      "title": "Public Service Body Rebate (GST/HST) Regulations",
      "citation": "SOR/91-37",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-334",
      "title": "Public Service Employment Regulations",
      "citation": "SOR/2005-334",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-59",
      "title": "Public Service Labour Relations Act Separate Agency Designation Order",
      "citation": "SOR/2005-59",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-79",
      "title": "Public Service Labour Relations Regulations",
      "citation": "SOR/2005-79",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-79",
      "title": "Public Service Labour Relations Regulations",
      "citation": "SOR/2005-79",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-347",
      "title": "Public Service Official Languages Appointment Regulations",
      "citation": "SOR/2005-347",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-118",
      "title": "Public Service Official Languages Exclusion Approval Order",
      "citation": "SI/2005-118",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-33b",
      "title": "Public Service Part-time Regulations",
      "citation": "SOR/81-33b",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1352",
      "title": "Public Service Pension Adjustment Regulations",
      "citation": "CRC, c 1352",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-6",
      "title": "Public Service Staffing Complaints Regulations",
      "citation": "SOR/2006-6",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-6",
      "title": "Public Service Staffing Complaints Regulations",
      "citation": "SOR/2006-6",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1358",
      "title": "Public Service Superannuation Regulations",
      "citation": "CRC, c 1358",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1359",
      "title": "Public Service Superannuation Special Election Regulations",
      "citation": "CRC, c 1359",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1364",
      "title": "Public Works Leasing Regulations",
      "citation": "CRC, c 1364",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1365",
      "title": "Public Works Nuisances Regulations",
      "citation": "CRC, c 1365",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-139a",
      "title": "Department of Public Works Terms Under Three Months Exclusion Approval Order, 1993",
      "citation": "SOR/93-139a",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-269",
      "title": "Pulp and Paper Effluent Regulations",
      "citation": "SOR/92-269",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-268",
      "title": "Pulp and Paper Mill Defoamer and Wood Chip Regulation",
      "citation": "SOR/92-268",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-267",
      "title": "Pulp and Paper Mill Effluent Chlorinated Dioxins and Furans Regulations",
      "citation": "SOR/92-267",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-51",
      "title": "Regulations for the Purposes of the Canada Strategic Infrastructure Fund Act",
      "citation": "SOR/2003-51",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-180",
      "title": "Qalipu Mi\u0027kmaq First Nation Band Order",
      "citation": "SOR/2011-180",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-91",
      "title": "Qikiqtarjuaq Airport Zoning Regulations",
      "citation": "SOR/2012-91",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-594",
      "title": "Qualifications for Designations as Analysts Regulations",
      "citation": "SOR/98-594",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1368",
      "title": "Quarantine Regulations",
      "citation": "CRC, c 1368",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-104",
      "title": "Quebec Airport Zoning Regulations",
      "citation": "CRC, c 104",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-60-481",
      "title": "Quebec Apple Growers\u0027 Marketing Order",
      "citation": "SOR/60-481",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-293",
      "title": "Quebec Beef Cattle Order",
      "citation": "SOR/92-293",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-108",
      "title": "Quebec Beef Cattle Producers\u0027 Levies or Charges (Interprovincial and Export Trade) Order",
      "citation": "SOR/93-108",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1381",
      "title": "Quebec Central Railway Traffic Rules and Regulations",
      "citation": "CRC, c 1381",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-100",
      "title": "Quebec Domestic Help Charities Remission Order",
      "citation": "SI/2011-100",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-246",
      "title": "Quebec Family Allowances (1988) Remission Order",
      "citation": "SI/88-246",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-116",
      "title": "Quebec Family Allowances (1989) Remission Order",
      "citation": "SI/90-116",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-153",
      "title": "Quebec Family Allowances Income Tax Remission Order, 1991",
      "citation": "SI/91-153",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-42",
      "title": "Quebec Family Allowances Income Tax Remission Order, 1992",
      "citation": "SI/93-42",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-214",
      "title": "Quebec Fishery Regulations, 1990",
      "citation": "SOR/90-214",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-243",
      "title": "Quebec Flue-Cured Tobacco Order",
      "citation": "CRC, c 243",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-55",
      "title": "Quebec Gross Revenue Insurance Program Conditional Remission Order",
      "citation": "SI/2004-55",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-549",
      "title": "Quebec Hog Marketing Levies (Interprovincial and Export Trade) Order",
      "citation": "SOR/94-549",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-727",
      "title": "Quebec Hog Marketing Order",
      "citation": "SOR/79-727",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-154",
      "title": "Quebec Maple Sap and Maple Syrup Order",
      "citation": "SOR/93-154",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-195",
      "title": "Quebec Maple Syrup Producers\u0027 Levy (Interprovincial and Export Trade) Order",
      "citation": "SOR/93-195",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-720",
      "title": "Quebec Milk Order",
      "citation": "SOR/94-720",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-1134",
      "title": "Quebec Port Warden\u0027s Tariff of Fees Order",
      "citation": "SOR/85-1134",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-75",
      "title": "Quebec Rules of Practice Respecting Reduction in the Number of Years of Imprisonment without Eligibility for Parole",
      "citation": "SOR/2002-75",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-6",
      "title": "Quebec Sex Offender Information Registration Regulations",
      "citation": "SOR/2005-6",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-91",
      "title": "Quebec Sheep and Wool Order",
      "citation": "SOR/92-91",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-258",
      "title": "Quebec-South Maple Products Order",
      "citation": "CRC, c 258",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-260",
      "title": "Quebec Turkey Marketing Levies Order",
      "citation": "CRC, c 260",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-259",
      "title": "Quebec Turkey Order",
      "citation": "CRC, c 259",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-118",
      "title": "Quebec Vegetables for Processing Order",
      "citation": "SOR/92-118",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-713",
      "title": "Quebec Wood Order, 1983",
      "citation": "SOR/83-713",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-277",
      "title": "Quebec Wood Producers\u0027 Levies (Interprovincial and Export Trade) Order",
      "citation": "SOR/98-277",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-453",
      "title": "Quesnel Airport Zoning Regulations",
      "citation": "SOR/87-453",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-340",
      "title": "Rabies Indemnification Regulations",
      "citation": "CRC, c 340",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-730",
      "title": "Racing Products Transfer and Loan Regulations",
      "citation": "SOR/85-730",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1370",
      "title": "Radiation Emitting Devices Regulations",
      "citation": "CRC, c 1370",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-203",
      "title": "Radiation Protection Regulations",
      "citation": "SOR/2000-203",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-206",
      "title": "Radiocommunication Act (Paragraph 9(1)(c)) Exemption Order",
      "citation": "SOR/2001-206",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-16",
      "title": "Radiocommunication Act (Paragraph 9(1)(c)) Exemption Order (National Defence and Security)",
      "citation": "SOR/2002-16",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-67",
      "title": "Radiocommunication Act (subsection 4(1) and paragraph 9(1)(b)) Exemption Order (Security, Safety and International Relations), No. 2009–1",
      "citation": "SOR/2009-67",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-223",
      "title": "Radiocommunication Act (Subsection 4(1) and Paragraph 9(1)(b)) Exemption Order (Security and Safety, International Relations and National Defence), No. 2002-1",
      "citation": "SOR/2002-223",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-224",
      "title": "Radiocommunication Act (Subsection 4(1) and Paragraph 9(1)(b)) Exemption Order (Security and Safety, International Relations and National Defence), No. 2002-2",
      "citation": "SOR/2002-224",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-34",
      "title": "Radiocommunication Act (Subsection 4(1) and Paragraph 9(1)(b)) Exemption Order (Security, Safety and International Relations), no. 2012‑1",
      "citation": "SOR/2012-34",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-284",
      "title": "Radiocommunication Act (Subsection 4(1) and Paragraph 9(1)(b)) Exemption Order (Security, Safety and International Relations), No. 2004-1",
      "citation": "SOR/2004-284",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-193",
      "title": "Radiocommunication Act (Subsection 4(1) and Paragraph 9(1)(b)) Exemption Order (Security, Safety and International Relations), No. 2007–1",
      "citation": "SOR/2007-193",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-292",
      "title": "Radiocommunication Act (Subsection 4(1) and Paragraph 9(1)(b)) Exemption Order (Security, Safety and International Relations), No. 2008-1",
      "citation": "SOR/2008-292",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-29",
      "title": "Radiocommunication Act (Subsection 4(1) and Paragraph 9(1)(b)) Exemption Order (Security, Safety and International Relations), No. 2010-1",
      "citation": "SOR/2010-29",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-52",
      "title": "Radiocommunication Act (Subsection 4(1) and Paragraph 9(1)(b)) Exemption Order (Security, Safety and International Relations), No. 2010–2",
      "citation": "SOR/2010-52",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-156",
      "title": "Radiocommunication Act (Subsection 4(1) and Paragraph 9(1)(b)) Exemption Order (Security, Safety and International Relations), No. 2010–3",
      "citation": "SOR/2010-156",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-135",
      "title": "Radiocommunication Act (Subsection 4(1) and Paragraph 9(1)(b)) Exemption Order (Security, Safety and International Relations), No. 2010–5",
      "citation": "SOR/2010-135",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-484",
      "title": "Radiocommunication Regulations",
      "citation": "SOR/96-484",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-982",
      "title": "Radio Regulations, 1986",
      "citation": "SOR/86-982",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-76-31",
      "title": "Rail Rate Freeze Compensatory Payments Order",
      "citation": "SI/76-31",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1382",
      "title": "Railway Abandonment Regulations",
      "citation": "CRC, c 1382",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-207",
      "title": "Railway Company Pay Out of Excess Revenue for the Movement of Grain Regulations",
      "citation": "SOR/2001-207",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-310",
      "title": "Railway Costing Regulations",
      "citation": "SOR/80-310",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-150",
      "title": "Railway Employee Qualification Standards Regulations",
      "citation": "SOR/87-150",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-748",
      "title": "Railway-Highway Crossing at Grade Regulations",
      "citation": "SOR/80-748",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-192",
      "title": "Railway Hygiene Regulations",
      "citation": "SOR/85-192",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-41",
      "title": "Railway Interswitching Regulations",
      "citation": "SOR/88-41",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-258",
      "title": "Railway Operating Certificate Regulations",
      "citation": "SOR/2014-258",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-342",
      "title": "Railway Passenger Services Adjustment Assistance Regulations",
      "citation": "CRC, c 342",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-286",
      "title": "Railway Passenger Services Contract Regulations",
      "citation": "SOR/78-286",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-1015",
      "title": "Railway Prevention of Electric Sparks Regulations",
      "citation": "SOR/82-1015",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-780",
      "title": "Railway Rolling Stock (Canadian Manufactured) Remission Order, No. 1",
      "citation": "CRC, c 780",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-161",
      "title": "Railway Rolling Stock (Temporary Domestic Service U.S. Built Multi-Level) Remission Order",
      "citation": "SI/86-161",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-991",
      "title": "Railway Running-Trades Employees Hours of Work Regulations",
      "citation": "CRC, c 991",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-233",
      "title": "Railway Safety Administrative Monetary Penalties Regulations",
      "citation": "SOR/2014-233",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1171",
      "title": "Railway Safety Appliance Standards Regulations",
      "citation": "CRC, c 1171",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-37",
      "title": "Railway Safety Management System Regulations",
      "citation": "SOR/2001-37",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-337",
      "title": "Railway Third Party Liability Insurance Coverage Regulations",
      "citation": "SOR/96-337",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-21",
      "title": "Railway Track Scales for In-Motion Weighing Specifications",
      "citation": "SI/91-21",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-338",
      "title": "Railway Traffic and Passenger Tariffs Regulations",
      "citation": "SOR/96-338",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-488",
      "title": "Railway Traffic Liability Regulations",
      "citation": "SOR/91-488",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-103",
      "title": "Notice of Railway Works Regulations",
      "citation": "SOR/91-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2007-31",
      "title": "Rainy River First Nations Settlement Agreement Remission Order",
      "citation": "SI/2007-31",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-351",
      "title": "Order Raising the Limit on the Amount that the Canadian Broadcasting Corporation may Expend Under a Lease or Other Rental Agreement",
      "citation": "SOR/2002-351",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-253",
      "title": "Ranch-raised Fur Pelts Designation Regulations",
      "citation": "SOR/99-253",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-326",
      "title": "Rankin Inlet Airport Zoning Regulations",
      "citation": "SOR/93-326",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-313",
      "title": "RCMP External Review Committee Rules of Practice and Procedure",
      "citation": "SOR/88-313",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-886",
      "title": "R.C.M.P. Stoppage of Pay and Allowances Regulations",
      "citation": "SOR/84-886",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-343",
      "title": "Real Property Grants Regulations",
      "citation": "CRC, c 343",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-128",
      "title": "Receipt and Deposit of Public Money Regulations, 1997",
      "citation": "SOR/98-128",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-91",
      "title": "Order Acknowledging Receipt of the Assessments Done Pursuant to Subsection 23(1) of the Act",
      "citation": "SI/2011-91",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-46",
      "title": "Order Acknowledging Receipt of the Assessments Done Pursuant to Subsection 23(1) of the Act",
      "citation": "SI/2012-46",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-8",
      "title": "Proclamation Declaring the Reciprocal Agreement on Social Security Between Canada and Australia in Force September 1, 1989",
      "citation": "SI/90-8",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-52",
      "title": "Proclamation Declaring the Reciprocal Agreement on Social Security Between Canada and Iceland in Force October 1, 1989",
      "citation": "SI/90-52",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-27",
      "title": "Order According Reciprocal Protection to Switzerland Under the Act",
      "citation": "SOR/94-27",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-57",
      "title": "Proclamation recognizing the outstanding service to Canadians by employees in the public service of Canada in times of natural disaster",
      "citation": "SI/98-57",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-103",
      "title": "Order Accepting the Recommendation of the Minister of Public Safety and Emergency Preparedness Concerning the Two-year Review of the List of Entities",
      "citation": "SI/2012-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2010-90",
      "title": "Order Accepting the Recommendation of the Minister of Public Safety and Emergency Preparedness That Each Entity Listed as of July 23, 2010, in the Regulations Establishing a List of Entities Remain a Listed Entity",
      "citation": "SI/2010-90",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-155",
      "title": "Order Recommending that Each Entity Listed as of July 23, 2004, in the Regulations Establishing a List of Entities Remain a Listed Entity",
      "citation": "SI/2004-155",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-133",
      "title": "Order Recommending that Each Entity Listed as of July 23, 2006, in the Regulations Establishing a List of entities Remain a Listed Entity",
      "citation": "SI/2006-133",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-63",
      "title": "Reconsideration Request Regulations.",
      "citation": "SOR/2013-63",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-568",
      "title": "Recovery of Overpayments Made to Former Members of Parliament Regulations",
      "citation": "SOR/97-568",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-264",
      "title": "Regulations for the Redemption of One Cent Coins",
      "citation": "SOR/2012-264",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-450",
      "title": "Redemption of Subsidiary Coin Regulations",
      "citation": "CRC, c 450",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-556",
      "title": "Red Lake Airport Zoning Regulations",
      "citation": "SOR/94-556",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-4",
      "title": "Red River Basin (Manitoba) Flood Relief Payments Remission Order",
      "citation": "SI/98-4",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-167",
      "title": "Reduction of Carbon Dioxide Emissions from Coal-fired Generation of Electricity Regulations",
      "citation": "SOR/2012-167",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-265",
      "title": "Order Declaring that the Reduction of Carbon Dioxide Emissions from Coal-fired Generation of Electricity Regulations do not apply in Nova Scotia",
      "citation": "SOR/2014-265",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-36",
      "title": "Order Declining to Refer Back to the CRTC Decision CRTC 2005-15",
      "citation": "SI/2005-36",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-82",
      "title": "Order Declining to Refer Back to the CRTC Decision CRTC 2005-248",
      "citation": "SI/2005-82",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-108",
      "title": "Order Declining to Refer Back to the CRTC Decision CRTC 2006-193",
      "citation": "SI/2006-108",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-69",
      "title": "Order Declining to Refer Back to the CRTC Decision CRTC 2012-308",
      "citation": "SI/2012-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-129",
      "title": "Order Declining to Refer Back to the CRTC Decisions CRTC 2006-323 and 2006-324",
      "citation": "SI/2006-129",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-180",
      "title": "Proclamation Directing a Referendum Relating to the Constitution of Canada",
      "citation": "SI/92-180",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-11",
      "title": "Order Referring Back to the CRTC Various Decisions",
      "citation": "SI/97-11",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-257",
      "title": "Refugee Appeal Division Rules",
      "citation": "SOR/2012-257",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-228",
      "title": "Refugee Protection Division Rules",
      "citation": "SOR/2002-228",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-256",
      "title": "Refugee Protection Division Rules",
      "citation": "SOR/2012-256",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-56",
      "title": "Refund of Duties on Obsolete or Surplus Goods Regulations",
      "citation": "SOR/98-56",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-48",
      "title": "Refund of Duties Regulations",
      "citation": "SOR/98-48",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-105",
      "title": "Regina Airport Zoning Regulations",
      "citation": "CRC, c 105",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1593",
      "title": "Regional Advisory Committee Regulations",
      "citation": "CRC, c 1593",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1387",
      "title": "Regional Development Incentives Designated Regions Order, 1974",
      "citation": "CRC, c 1387",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1386",
      "title": "Regional Development Incentives Regulations",
      "citation": "CRC, c 1386",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1388",
      "title": "Regional Development Incentives Regulations, 1974",
      "citation": "CRC, c 1388",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-78",
      "title": "Region of Waterloo International Airport Zoning Regulations",
      "citation": "SOR/2006-78",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-763",
      "title": "Registered Charities Information Return Fee Order",
      "citation": "SOR/90-763",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-99",
      "title": "Registered Products Regulations",
      "citation": "SOR/2011-99",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-216",
      "title": "Order Designating the Registrar of the Competition Tribunal as Deputy Head with Respect to the Registry of the Competition Tribunal",
      "citation": "SI/86-216",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-17",
      "title": "Order Designating the Registrar of the Office of the Registrar of Lobbyists as Deputy Head",
      "citation": "SI/2006-17",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-301",
      "title": "Registration of Bank Special Security Regulations",
      "citation": "SOR/92-301",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-84",
      "title": "Regulations Implementing the United Nations Resolution on Eritrea",
      "citation": "SOR/2010-84",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-213",
      "title": "Regulations Implementing the United Nations Resolution on Yemen",
      "citation": "SOR/2014-213",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-294",
      "title": "Regulations Amending Certain Regulations Made Under the Broadcasting Act",
      "citation": "SOR/2009-294",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-420",
      "title": "Regulatory Capital (Bank Holding Companies) Regulations",
      "citation": "SOR/2001-420",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-531",
      "title": "Regulatory Capital (Banks) Regulations",
      "citation": "SOR/92-531",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-528",
      "title": "Regulatory Capital (Cooperative Credit Associations) Regulations",
      "citation": "SOR/92-528",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-529",
      "title": "Regulatory Capital (Insurance Companies) Regulations",
      "citation": "SOR/92-529",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-424",
      "title": "Regulatory Capital (Insurance Holding Companies) Regulations",
      "citation": "SOR/2001-424",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-530",
      "title": "Regulatory Capital (Trust and Loan Companies) Regulations",
      "citation": "SOR/92-530",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-298",
      "title": "Reinsurance (Canadian Companies) Regulations",
      "citation": "SOR/92-298",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-302",
      "title": "Reinsurance (Foreign Companies) Regulations",
      "citation": "SOR/92-302",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-138",
      "title": "Related Party of a Retail Association Regulations",
      "citation": "SOR/2005-138",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-309",
      "title": "Related Party Transactions (Banks) Regulations",
      "citation": "SOR/92-309",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-275",
      "title": "Related Party Transactions (Cooperative Credit Associations) Regulations",
      "citation": "SOR/96-275",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-276",
      "title": "Related Party Transactions (Insurance Companies) Regulations",
      "citation": "SOR/96-276",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-277",
      "title": "Related Party Transactions (Trust and Loan Companies) Regulations",
      "citation": "SOR/96-277",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-90",
      "title": "Release and Environmental Emergency Notification Regulations",
      "citation": "SOR/2011-90",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-315",
      "title": "Release of Information for Family Orders and Agreements Enforcement Regulations",
      "citation": "SOR/87-315",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-202",
      "title": "Regulations Relieving Special Duty on Certain Tobacco Products",
      "citation": "SOR/2003-202",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-479",
      "title": "Order Respecting the Remission of Anti-Dumping and Countervailing Duties on Certain Specialty Sugar Products",
      "citation": "SOR/97-479",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-148",
      "title": "Order Respecting the Remission of Anti-dumping Duties on Automotive Galvannealed Coil Steel",
      "citation": "SOR/99-148",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-480",
      "title": "Order Respecting the Remission of Anti-Dumping Duties on Certain Hot-Dipped Galvannealed Steel Sheet for use in the Manufacture of Non-Exposed Motor Vehicle Parts",
      "citation": "SOR/97-480",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-135",
      "title": "Order Respecting the Remission of Anti-Dumping Duties on Vitreous Type I Cold-Rolled Steel",
      "citation": "SOR/98-135",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-108",
      "title": "Remission of Application Fee for Certificate of Citizenship Order",
      "citation": "SI/85-108",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-41",
      "title": "Order Respecting the Remission of Equalization Overpayments Made to Certain Provinces under the Federal-Provincial Fiscal Arrangements Act",
      "citation": "SI/2004-41",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-23",
      "title": "Order Respecting the Remission of Interest Payments Owing to the Government of Canada on Loans Made to the Republic of Cameroon, the Republic of Madagascar and the Republic of Zambia",
      "citation": "SI/2001-23",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-100",
      "title": "Remission Order Concerning Certain Expenses Incurred as a Result of the Tsunami Disaster of December 26, 2004 in South and Southeast Asia",
      "citation": "SI/2006-100",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-103",
      "title": "Remission Order Concerning Certain Passport Services Fees and Consular Services Fees",
      "citation": "SI/2003-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-60",
      "title": "Remission Order Concerning Interest Accruing on Certain Unemployment Insurance Overpayments",
      "citation": "SI/2002-60",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-23",
      "title": "Remission Order in Respect of Goods Under the Customs Duties Reduction or Removal Order, 1988",
      "citation": "SOR/98-23",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-66",
      "title": "Remote Sensing Space Systems Regulations",
      "citation": "SOR/2007-66",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-189",
      "title": "Renewable Fuels Regulations",
      "citation": "SOR/2010-189",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-88",
      "title": "Renewable Fuel Used as Ships\u0027 Stores Remission Order",
      "citation": "SOR/2010-88",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-82-131",
      "title": "Repair Abroad of Canadian Civil Aircraft, Canadian Aircraft Engines and Flight Simulators Remission Order",
      "citation": "SI/82-131",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-127",
      "title": "Repayment of Receipts Regulations, 1997",
      "citation": "SOR/98-127",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-41",
      "title": "Order Repealing Order in Council P.C. 2006-62",
      "citation": "SI/2014-41",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-344",
      "title": "Regulations Repealing the Hazardous Products (Ice Hockey Helmets) Regulations (Miscellaneous Program)",
      "citation": "SOR/2005-344",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-2",
      "title": "Reportable Diseases Regulations",
      "citation": "SOR/91-2",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-23",
      "title": "Reporting of Exported Goods Regulations",
      "citation": "SOR/2005-23",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-873",
      "title": "Reporting of Imported Goods Regulations",
      "citation": "SOR/86-873",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-479",
      "title": "Report on the State of Canada\u0027s Forests Regulations",
      "citation": "SOR/95-479",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-785",
      "title": "Representational Gifts Remission Order",
      "citation": "CRC, c 785",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-154",
      "title": "Proclamation Declaring the Representation Order to be in Force Effective on the First Dissolution of Parliament That Occurs After August 25, 2004",
      "citation": "SI/2003-154",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-102",
      "title": "Proclamation Declaring the Representation Order to be in Force Effective on the First Dissolution of Parliament That Occurs after May 1, 2014",
      "citation": "SI/2013-102",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-5",
      "title": "Reproduction of Federal Law Order",
      "citation": "SI/97-5",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1292",
      "title": "Reproduction of Postage Stamps Regulations",
      "citation": "CRC, c 1292",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-106",
      "title": "Republic of Cameroon Remission Order",
      "citation": "SI/2005-106",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-14",
      "title": "Proclamation Designating the Republic of Croatia as a Designated State for Purposes of the Visiting Forces Act",
      "citation": "SOR/2009-14",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-343",
      "title": "Repulse Bay Airport Zoning Regulations",
      "citation": "SOR/92-343",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-139",
      "title": "Order Rescinding Decision No. 618-W-2005 of the Canadian Transportation Agency",
      "citation": "SOR/2006-139",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-152",
      "title": "Regulations Respecting Research, Market Development and Technical Assistance (Wheat and Barley)",
      "citation": "SOR/2012-152",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-856",
      "title": "Reservation to the Crown Waiver Order, 1983, No. 1",
      "citation": "SOR/83-856",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-962",
      "title": "Reservation to the Crown Waiver Order, 1984, No. 1",
      "citation": "SOR/84-962",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-1003",
      "title": "Reservation to the Crown Waiver Order, 1985, No. 4",
      "citation": "SOR/85-1003",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-461",
      "title": "Reservation to the Crown Waiver Order, 1986, No. 2",
      "citation": "SOR/86-461",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-9",
      "title": "Reservation to the Crown Waiver Order (Bed of Great Slave Lake)",
      "citation": "SI/2005-9",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-9",
      "title": "Reservation to the Crown Waiver Order (Contwoyto Lake, N.W.T.)",
      "citation": "SI/99-9",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-80",
      "title": "Reservation to the Crown Waiver Order (Dolomite Lake, N.W.T.)",
      "citation": "SI/2012-80",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-7",
      "title": "Reservation to the Crown Waiver Order (Grace Lake, N.W.T.)",
      "citation": "SI/2013-7",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-8",
      "title": "Reservation to the Crown Waiver Order (Great Slave Lake)",
      "citation": "SI/2005-8",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-28",
      "title": "Reservation to the Crown Waiver Order (Great Slave Lake, N.T.)",
      "citation": "SI/94-28",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-9",
      "title": "Reservation to the Crown Waiver Order (Great Slave Lake, N.W.T.)",
      "citation": "SI/2002-9",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-205",
      "title": "Reservation to the Crown Waiver Order (Great Slave Lake, N.W.T.)",
      "citation": "SI/92-205",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-82",
      "title": "Reservation to the Crown Waiver Order (Great Slave Lake, N.W.T.)",
      "citation": "SI/2012-82",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2010-54",
      "title": "Reservation to the Crown Waiver Order (Great Slave Lake, N.W.T.)",
      "citation": "SI/2010-54",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-69",
      "title": "Reservation to the Crown Waiver Order (Hay River, N.W.T.)",
      "citation": "SI/2006-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-139",
      "title": "Reservation to the Crown Waiver Order (Hay River, N.W.T.)",
      "citation": "SI/99-139",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-119",
      "title": "Reservation to the Crown Waiver Order (Hay River, N.W.T.)",
      "citation": "SI/92-119",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-76",
      "title": "Reservation to the Crown Waiver Order (Hecla and Griper Bay, N.W.T.)",
      "citation": "SI/2012-76",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-27",
      "title": "Reservation to the Crown Waiver Order (Kam Lake, N.W.T.)",
      "citation": "SI/93-27",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-10",
      "title": "Reservation to the Crown Waiver Order (La Martre Lake, N.W.T.)",
      "citation": "SI/2002-10",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-130",
      "title": "Reservation to the Crown Waiver Order (Mackenzie River, N.W.T.)",
      "citation": "SI/94-130",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-114",
      "title": "Reservation to the Crown Waiver Order (Mackenzie River, N.W.T.)",
      "citation": "SI/90-114",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-81",
      "title": "Reservation to the Crown Waiver Order (Mackenzie River, N.W.T.)",
      "citation": "SI/2012-81",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-96",
      "title": "Reservation to the Crown Waiver Order (Niven Lake, N.W.T.)",
      "citation": "SI/96-96",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-60",
      "title": "Reservation to the Crown Waiver Order (Port Radium, N.W.T.)",
      "citation": "SI/99-60",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-240",
      "title": "Reservation to the Crown Waiver Order (Prosperous Lake, N.T.)",
      "citation": "SI/93-240",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-92",
      "title": "Reservation to the Crown Waiver Order (Seepage Lake, N.W.T.)",
      "citation": "SI/98-92",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-171",
      "title": "Reservation to the Crown Waiver Order (Shell Lake, N.W.T.)",
      "citation": "SI/2003-171",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-115",
      "title": "Reservation to the Crown Waiver Order (Tuktoyaktuk Harbour, N.W.T.)",
      "citation": "SI/90-115",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-82",
      "title": "Reservation to the Crown Waiver Order (Tuktoyaktuk, N.W.T.)",
      "citation": "SI/2000-82",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-91",
      "title": "Reservation to the Crown Waiver Order (Yellowknife Bay, N.W.T.)",
      "citation": "SI/98-91",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-83",
      "title": "Reservation to the Crown Waiver Order (Yellowknife Bay, N.W.T.)",
      "citation": "SI/2012-83",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-113",
      "title": "Reservation to the Crown Waiver Order (Yellowknife, N.W.T.)",
      "citation": "SI/94-113",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-32",
      "title": "Reserve Force Pension Plan Regulations",
      "citation": "SOR/2007-32",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1050",
      "title": "Reserve Forces Training Leave Regulations",
      "citation": "CRC, c 1050",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-282",
      "title": "Resident Canadian (Banks) Regulations",
      "citation": "SOR/92-282",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-285",
      "title": "Resident Canadian (Cooperative Credit Associations) Regulations",
      "citation": "SOR/92-285",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-284",
      "title": "Resident Canadian (Insurance Companies) Regulations",
      "citation": "SOR/92-284",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-283",
      "title": "Resident Canadian (Trust and Loan Companies) Regulations",
      "citation": "SOR/92-283",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-193",
      "title": "Residential Detectors Regulations",
      "citation": "SOR/2009-193",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-137",
      "title": "Residents of India Remission Order",
      "citation": "SI/91-137",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-106",
      "title": "Resolute Bay Airport Zoning Regulations",
      "citation": "CRC, c 106",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-30",
      "title": "Remission Order Respecting Imports of Goods Originating in Commonwealth Developing Countries",
      "citation": "SOR/98-30",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-405",
      "title": "Response Organizations and Oil Handling Facilities Regulations",
      "citation": "SOR/95-405",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-120",
      "title": "Order Setting Out the Respective Responsibilities of the Minister of Citizenship and Immigration and the Minister of Public Safety and Emergency Preparedness Under the Act",
      "citation": "SI/2005-120",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-16",
      "title": "Restraint Systems and Booster Seats for Motor Vehicles Regulations",
      "citation": "SOR/2011-16",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-47",
      "title": "Restricted Components Regulations",
      "citation": "SOR/2008-47",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-416",
      "title": "Restrictive Trade Practices Commission Rules",
      "citation": "CRC, c 416",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-216",
      "title": "Retail Association Regulations",
      "citation": "SOR/2002-216",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-785",
      "title": "Retirement Compensation Arrangements Regulations, No. 1",
      "citation": "SOR/94-785",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-169",
      "title": "Retirement Compensation Arrangements Regulations, No. 2",
      "citation": "SOR/95-169",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "dors-91-690",
      "title": "Retransmission Royalties Criteria Regulations",
      "citation": "DORS/91-690",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-48",
      "title": "Returnable Beverage Container (GST/HST) Regulations",
      "citation": "SOR/2008-48",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-345",
      "title": "Returned Contributions Interest Regulations",
      "citation": "CRC, c 345",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1390",
      "title": "Returned Soldiers\u0027 Insurance Regulations",
      "citation": "CRC, c 1390",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-61",
      "title": "Returning Persons Exemption Regulations",
      "citation": "SOR/98-61",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-200",
      "title": "Return of Packaged Alcohol to an Excise Warehouse Regulations",
      "citation": "SOR/2003-200",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-48",
      "title": "Revenue Canada Technical Publication Subscription Service Fees Order",
      "citation": "SOR/93-48",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-730",
      "title": "Revenue Trust Account Regulations",
      "citation": "CRC, c 730",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-22",
      "title": "Review Panel Regulations",
      "citation": "SOR/2008-22",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-19",
      "title": "Review Tribunal Rules of Procedure",
      "citation": "SOR/92-19",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-50",
      "title": "Revised Statutes of Canada, 1985 Distribution Direction",
      "citation": "SI/88-50",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-122",
      "title": "Revised Statutes of Canada, 1985 Distribution Direction « No. 2 »",
      "citation": "SI/89-122",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1328",
      "title": "Rideau Canal Lease Regulations",
      "citation": "CRC, c 1328",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2009-110",
      "title": "Right of Permanent Residence Fees Remission Order",
      "citation": "SI/2009-110",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-20",
      "title": "Right of Permanent Residence Fees Remission Order (2012)",
      "citation": "SI/2012-20",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-101",
      "title": "Rimouski Airport Zoning Regulations",
      "citation": "SOR/91-101",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-837",
      "title": "Roe Herring Exemption Notice",
      "citation": "CRC, c 837",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-72-91",
      "title": "Romania Claims Fund Regulations",
      "citation": "SOR/72-91",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-73-558",
      "title": "Rosabel Jeanette Pryce Housing Regulations",
      "citation": "SOR/73-558",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-386",
      "title": "Ross River Airport Zoning Regulations",
      "citation": "SOR/94-386",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-537",
      "title": "Royal Canadian Mounted Police Automated Fingerprint Identification System Fees Regulations",
      "citation": "SOR/94-537",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-421",
      "title": "Royal Canadian Mounted Police Case Examination Fees (Bermuda) Regulations",
      "citation": "SOR/89-421",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-253",
      "title": "Royal Canadian Mounted Police Casual Employment Regulations",
      "citation": "SOR/2014-253",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-482",
      "title": "Royal Canadian Mounted Police, Criminal Record Verification for Certain Enterprises Fees Regulations",
      "citation": "SOR/93-482",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-485",
      "title": "Royal Canadian Mounted Police, Criminal Record Verification for Civil Purposes Fee Regulations",
      "citation": "SOR/93-485",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-122",
      "title": "Royal Canadian Mounted Police (Dependants) Pension Fund Increase in Benefits Order",
      "citation": "SOR/2011-122",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-70",
      "title": "Royal Canadian Mounted Police (Dependants) Pension Fund Increase in Benefits Order",
      "citation": "SOR/2014-70",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-397",
      "title": "Royal Canadian Mounted Police External Review Committee Security and Confidentiality Regulations",
      "citation": "SOR/88-397",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-481",
      "title": "Royal Canadian Mounted Police Fingerprinting for Visa, Licensing or Security Clearance Purposes Fee Regulations",
      "citation": "SOR/93-481",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-480",
      "title": "Royal Canadian Mounted Police, Forensic Laboratory Services Fees Regulations",
      "citation": "SOR/93-480",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-422",
      "title": "Order Specifying the Royal Canadian Mounted Police for the Purposes of the Employment Equity Act",
      "citation": "SOR/2002-422",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1392",
      "title": "Royal Canadian Mounted Police Pension Continuation Regulations",
      "citation": "CRC, c 1392",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-116",
      "title": "Royal Canadian Mounted Police, Police Information Retrieval System Fees Order",
      "citation": "SOR/90-116",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-484",
      "title": "Royal Canadian Mounted Police, Private Communications Interception Fee Regulations",
      "citation": "SOR/93-484",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-17",
      "title": "Royal Canadian Mounted Police Public Complaints Commission Rules of Practice",
      "citation": "SOR/93-17",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-361",
      "title": "Royal Canadian Mounted Police Regulations, 1988",
      "citation": "SOR/88-361",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-281",
      "title": "Royal Canadian Mounted Police Regulations, 2014",
      "citation": "SOR/2014-281",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1393",
      "title": "Royal Canadian Mounted Police Superannuation Regulations",
      "citation": "CRC, c 1393",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-115",
      "title": "Royal Canadian Mounted Police Training Academy Fees Order",
      "citation": "SOR/90-115",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2009-23",
      "title": "Rule 63 — Summary Conviction Appeal",
      "citation": "SI/2009-23",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2009-22",
      "title": "Rule 64 — Prerogative Writ",
      "citation": "SI/2009-22",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2009-3",
      "title": "Rule 91 — Criminal Appeal",
      "citation": "SI/2009-3",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-308",
      "title": "Rules Governing Proceedings at Public Inquiries into Objections (Banks)",
      "citation": "SOR/92-308",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-307",
      "title": "Rules Governing Proceedings at Public Inquiries into Objections (Cooperative Credit Associations)",
      "citation": "SOR/92-307",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-306",
      "title": "Rules Governing Proceedings at Public Inquiries into Objections (Insurance Companies)",
      "citation": "SOR/92-306",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-296",
      "title": "Rules Governing Proceedings at Public Inquiries into Objections (Trust and Loan Companies)",
      "citation": "SOR/92-296",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-42",
      "title": "Rules of Practice and Procedure of the Tax Court of Canada in Respect of Appeals Under Part IX of the Excise Tax Act (Informal Procedure)",
      "citation": "SOR/92-42",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-10",
      "title": "Rules of Practice in Criminal Matters in the Court of Appeal of Quebec",
      "citation": "SI/99-10",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-46",
      "title": "Rules of Practice of the Superior Court of the Province of Quebec, Criminal Division, 2002",
      "citation": "SI/2002-46",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-28",
      "title": "Rules of Procedure for Boards of Review",
      "citation": "SOR/2003-28",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-241",
      "title": "Rules of Procedure for Hearings Before the Military Police Complaints Commission",
      "citation": "SOR/2002-241",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-94",
      "title": "Rules of Procedure for Rail Level of Service Arbitration",
      "citation": "SOR/2014-94",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-306",
      "title": "Rules of the Board of Arbitration (Agriculture and Agri-Food)",
      "citation": "SOR/2000-306",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-142",
      "title": "Rules of the Court of Appeal of Quebec in Criminal Matters",
      "citation": "SI/2006-142",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-68",
      "title": "Rules of the Court of Appeals for the Northwest Territories as to A. Criminal Appeals B. Bail on Appeals",
      "citation": "SOR/78-68",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-133",
      "title": "Rules of the Ontario Court of Justice in Criminal Proceedings",
      "citation": "SI/97-133",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-134",
      "title": "Rules of the Provincial Court of Newfoundland and Labrador in Criminal Proceedings",
      "citation": "SI/2004-134",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-451",
      "title": "Rules of the Review Tribunal (Agriculture and Agri-Food)",
      "citation": "SOR/99-451",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-156",
      "title": "Rules of the Supreme Court of Canada",
      "citation": "SOR/2002-156",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-168",
      "title": "Certain Ruminants and Their Products Importation Prohibition Regulations, No. 2",
      "citation": "SOR/2006-168",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-25",
      "title": "Rutabaga Stabilization Regulations, 1982-83",
      "citation": "SOR/84-25",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1465",
      "title": "Sable Island Regulations",
      "citation": "CRC, c 1465",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-506",
      "title": "Sachs Harbour Airport Zoning Regulations",
      "citation": "SOR/95-506",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-1038",
      "title": "Safe Containers Convention Regulations",
      "citation": "SOR/82-1038",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-22",
      "title": "Safeguard Surtax Regulations, 1995-1",
      "citation": "SOR/95-22",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-380",
      "title": "Safeguard Surtax Regulations, 1995-2",
      "citation": "SOR/95-380",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-381",
      "title": "Safeguard Surtax Regulations, 1995-3",
      "citation": "SOR/95-381",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-305",
      "title": "Safety and Health Committees and Representatives Regulations",
      "citation": "SOR/86-305",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-933",
      "title": "Safety Glass Regulations",
      "citation": "CRC, c 933",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-348",
      "title": "Safety Management Regulations",
      "citation": "SOR/98-348",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-118",
      "title": "Safety of Human Cells, Tissues and Organs for Transplantation Regulations",
      "citation": "SOR/2007-118",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1354",
      "title": "Safety or Security Positions Bargaining Direction",
      "citation": "CRC, c 1354",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1467",
      "title": "Safe Working Practices Regulations",
      "citation": "CRC, c 1467",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-114",
      "title": "Saint John Airport Zoning Regulations",
      "citation": "CRC, c 114",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-731",
      "title": "Sale of Defence Materiel and Services to United Nations Order",
      "citation": "CRC, c 731",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1010",
      "title": "Sale of Goods Regulations",
      "citation": "SOR/86-1010",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2008-28",
      "title": "Sale of the Digby Fisherman\u0027s Wharf Remission Order",
      "citation": "SI/2008-28",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-52",
      "title": "Sales or Trades (Authorized Foreign Banks) Regulations",
      "citation": "SOR/2000-52",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-304",
      "title": "Samples of Bodily Substances Regulations",
      "citation": "SOR/2014-304",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-786",
      "title": "Samples of Negligible Value Remission Order",
      "citation": "CRC, c 786",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-78",
      "title": "Sandy Bay Band Council Elections Order",
      "citation": "SOR/2003-78",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-125",
      "title": "Sandy Bay Band Council Method of Election Regulations",
      "citation": "SOR/2003-125",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-92",
      "title": "Sanikiluaq Airport Zoning Regulations",
      "citation": "SOR/2012-92",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-602",
      "title": "Sarnia Airport Zoning Regulations",
      "citation": "SOR/87-602",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-345",
      "title": "Saskatchewan Alfalfa Seed Order",
      "citation": "SOR/2001-345",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-268",
      "title": "Saskatchewan Broiler Chicken Order",
      "citation": "CRC, c 268",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-241",
      "title": "Saskatchewan Canaryseed Order",
      "citation": "SOR/2013-241",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-608",
      "title": "Saskatchewan Canola Development Commission Levies Order",
      "citation": "SOR/92-608",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-225",
      "title": "Saskatchewan Canola Order",
      "citation": "SOR/2003-225",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-158",
      "title": "Saskatchewan Court of Queen\u0027s Bench Rules Respecting Pre-Trial Conferences",
      "citation": "SI/86-158",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-270",
      "title": "Saskatchewan Egg Marketing Levies Order",
      "citation": "CRC, c 270",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-269",
      "title": "Saskatchewan Egg Order",
      "citation": "CRC, c 269",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-233",
      "title": "Saskatchewan Fishery Regulations, 1995",
      "citation": "SOR/95-233",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-344",
      "title": "Saskatchewan Flax Order",
      "citation": "SOR/2001-344",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-240",
      "title": "Saskatchewan Forage Seed Order",
      "citation": "SOR/2013-240",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-204",
      "title": "Order Designating Saskatchewan for the Purposes of the Criminal Interest Rate Provisions of the Criminal Code",
      "citation": "SOR/2011-204",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-275",
      "title": "Saskatchewan Hog Farm Registration (Interprovincial and Export) Regulations",
      "citation": "CRC, c 275",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-276",
      "title": "Saskatchewan Hog Information (Interprovincial and Export) Regulations",
      "citation": "CRC, c 276",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-273",
      "title": "Saskatchewan Hog Licensing (Interprovincial and Export) Regulations",
      "citation": "CRC, c 273",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-272",
      "title": "Saskatchewan Hog Marketing (Interprovincial and Export) Regulations",
      "citation": "CRC, c 272",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-277",
      "title": "Saskatchewan Hog Marketing Levies Order",
      "citation": "CRC, c 277",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-271",
      "title": "Saskatchewan Hog Order",
      "citation": "CRC, c 271",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-274",
      "title": "Saskatchewan Hog Service Charge (Interprovincial and Export) Regulations",
      "citation": "CRC, c 274",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-122",
      "title": "Saskatchewan Indian Federated College Remission Order, 2003",
      "citation": "SI/2003-122",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-96",
      "title": "Saskatchewan Liquor Board Employees\u0027 Remission Order",
      "citation": "SI/92-96",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-721",
      "title": "Saskatchewan Milk Order",
      "citation": "SOR/94-721",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-239",
      "title": "Saskatchewan Mustard Order",
      "citation": "SOR/2013-239",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-56-290",
      "title": "Saskatchewan Penitentiary proclaimed a penitentiary for Alberta and Saskatchewan",
      "citation": "SOR/56-290",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-701",
      "title": "Saskatchewan Pulse Crop Marketing Levies (Interprovincial and Export) Order",
      "citation": "SOR/87-701",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-403",
      "title": "Saskatchewan Pulse Crop Order",
      "citation": "SOR/87-403",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-99",
      "title": "Saskatchewan Review of Parole Ineligibility Rules",
      "citation": "SOR/2005-99",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-304",
      "title": "Saskatchewan Sex Offender Information Registration Regulations",
      "citation": "SOR/2004-304",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-281",
      "title": "Saskatchewan Turkey Marketing Levies Order",
      "citation": "CRC, c 281",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-280",
      "title": "Saskatchewan Turkey Order",
      "citation": "CRC, c 280",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-115",
      "title": "Saskatchewan Uranium Mines and Mills Exclusion Regulations",
      "citation": "SOR/2001-115",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-242",
      "title": "Saskatchewan Winter Cereals Order",
      "citation": "SOR/2013-242",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-706",
      "title": "Saskatoon Airport Zoning Regulations",
      "citation": "SOR/87-706",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-96",
      "title": "Satellite Remote Sensing Services Fees Order, 1987",
      "citation": "SOR/87-96",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-82-229",
      "title": "Satellites and Satellite Subsystems Remission Order, 1982",
      "citation": "SI/82-229",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-532",
      "title": "Satellites and Satellite Subsystems Remission Order, 1988",
      "citation": "SOR/88-532",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-109",
      "title": "Sault Ste. Marie Airport Zoning Regulations",
      "citation": "CRC, c 109",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-155",
      "title": "Schedule 1 Chemicals Regulations (Chemical Weapons Convention)",
      "citation": "SOR/2004-155",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-145",
      "title": "Order Amending the Schedule to the Customs Tariff (Conditions for Special Provisions for the Purposes of the United States Tariff (UST))",
      "citation": "SOR/2009-145",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-29",
      "title": "School Cafeteria Food and Beverages (GST/HST) Regulations",
      "citation": "SOR/91-29",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-49-42",
      "title": "Regulations re school fees and transportation costs for children of certain Government employees",
      "citation": "SOR/49-42",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-934",
      "title": "Science Education Sets Regulations",
      "citation": "CRC, c 934",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-82",
      "title": "Scientific or Exploratory Expeditions Remission Order",
      "citation": "SOR/95-82",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-789",
      "title": "Scientific or Technical Assistance Fees Order (Research Branch—Agriculture)",
      "citation": "SOR/94-789",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-732",
      "title": "Search, Reproduction and Certification of Documents Regulations",
      "citation": "CRC, c 732",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-96",
      "title": "“SeaRose FPSO” Remission Order, 2004",
      "citation": "SOR/2004-96",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-123",
      "title": "“SeaRose FPSO” Repair or Alteration Remission Order, 2012",
      "citation": "SOR/2012-123",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-84",
      "title": "Seasonal Residents\u0027 Remission Order, 1991",
      "citation": "SI/91-84",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-569",
      "title": "The Seaway International Bridge Corporation, Ltd. Regulation",
      "citation": "SOR/98-569",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-105",
      "title": "Seaway Property Regulations",
      "citation": "SOR/2003-105",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-155",
      "title": "Secondary Lead Smelter Release Regulations",
      "citation": "SOR/91-155",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-54",
      "title": "Proclamation Giving Notice that the Second Protocol Amending the Convention Between Canada and the Republic of Austria for the Avoidance of Double Taxation and the Prevention of Fiscal Evasion with respect to Taxes on Income and on Capital, done at Vienna on December 9, 1976, as Amended by the Protocol done at Vienna on June 15, 1999 Came into Force on October 1, 2013",
      "citation": "SI/2014-54",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-111",
      "title": "Proclamation Declaring the Second Supplementary Agreement Amending the Agreement on Social Security Between Canada and the United States of America in Force October 1, 1997",
      "citation": "SI/97-111",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-28",
      "title": "Secretariat of the Convention on Biological Diversity Remission Order (Part IX of the Excise Tax Act)",
      "citation": "SI/2001-28",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-172",
      "title": "Secretary of State Authority to Prescribe Fees and Charges Order (Canadian Centre for Management Development)",
      "citation": "SI/89-172",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-201",
      "title": "Order Transferring to the Secretary of State of Canada Certain Powers, Duties and Functions of the Ministers of Employment and Immigration, the Solicitor General of Canada and Multiculturalism and Citizenship Relating to Immigration and Citizenship and to the Department of the Secretary of State of Canada the Control and Supervision of Certain Portions of the Public Service Relating Thereto",
      "citation": "SI/93-201",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-30",
      "title": "Secure Electronic Signature Regulations",
      "citation": "SOR/2005-30",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-275",
      "title": "Securities Dealing Restrictions (Authorized Foreign Banks) Regulations",
      "citation": "SOR/99-275",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-279",
      "title": "Securities Dealing Restrictions (Banks) Regulations",
      "citation": "SOR/92-279",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-278",
      "title": "Securities Dealing Restrictions (Cooperative Credit Associations) Regulations",
      "citation": "SOR/92-278",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-280",
      "title": "Securities Dealing Restrictions (Insurance Companies) Regulations",
      "citation": "SOR/92-280",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-272",
      "title": "Securities Dealing Restrictions (Trust and Loan Companies) Regulations",
      "citation": "SOR/92-272",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-426",
      "title": "Security Certificate Transfer Fee (Banks, Bank Holding Companies, Insurance Companies and Insurance Holding Companies) Regulations",
      "citation": "SOR/2001-426",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-263",
      "title": "Security Certificate Transfer Fee (Cooperative Credit Associations) Regulations",
      "citation": "SOR/92-263",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-256",
      "title": "Security Certificate Transfer Fee (Trust and Loan Companies) Regulations",
      "citation": "SOR/92-256",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-505",
      "title": "Security for Debts Due to Her Majesty Regulations",
      "citation": "SOR/87-505",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-55",
      "title": "Security Interest (GST/HST) Regulations",
      "citation": "SOR/2011-55",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1400",
      "title": "Seeds Regulations",
      "citation": "CRC, c 1400",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-303",
      "title": "Seized Property Disposition Regulations",
      "citation": "SOR/94-303",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-171",
      "title": "Selected Listed Financial Institutions Attribution Method (GST/HST) Regulations",
      "citation": "SOR/2001-171",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-101",
      "title": "Selkirk First Nation (GST) Remission Order",
      "citation": "SI/2000-101",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-331",
      "title": "Selkirk Marine Railway Dry Dock Regulations, 1989",
      "citation": "SOR/89-331",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-337",
      "title": "Senate Sessional Allowance (Deductions for Non-attendance) Regulations",
      "citation": "SOR/98-337",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-126",
      "title": "Senate Sessional Allowance (Suspension) Regulations",
      "citation": "SOR/98-126",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-657",
      "title": "Sept-Iles Airport Zoning Regulations",
      "citation": "SOR/78-657",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-922",
      "title": "Service Equipment Cars Regulations",
      "citation": "SOR/86-922",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-63",
      "title": "Service of Documents Required or Authorized to be Served Under Sections 53 to 57 of the Conflict of Interest Act Regulations",
      "citation": "SOR/2007-63",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-69",
      "title": "Order Designating Shared Services Canada as a Department and the President as Deputy Head for Purposes of the Act",
      "citation": "SI/2011-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-67",
      "title": "Order Transfering to Shared Services Canada the Control and Supervision of Certain Portions of the Federal Public Administration in the Department of Public Works and Government Services",
      "citation": "SI/2011-67",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-150",
      "title": "Order Declining to Set Aside a CRTC Decision Respecting Westcom Radio Group Limited",
      "citation": "SI/87-150",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-92",
      "title": "Order Declining to Set Aside Decision CRTC 96-479",
      "citation": "SI/96-92",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-69",
      "title": "Order Declining to Set Aside Decision CRTC 97-137",
      "citation": "SI/97-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-84-69",
      "title": "Order Declining to Set Aside or Refer Back to the CRTC a Decision Issuing a Broadcasting Licence to Saskatoon Telecable Ltd.",
      "citation": "SI/84-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-139",
      "title": "Order Declining to Set Aside or Refer Back to the CRTC Decision CRTC 2003-115",
      "citation": "SI/2003-139",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-99",
      "title": "Order Declining to Set Aside or to Refer Back a Decision to the CRTC for Reconsideration and Hearing",
      "citation": "SI/87-99",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-41",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Amending the Broadcasting Licence of Moffat Communications Limited",
      "citation": "SI/85-41",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-84-231",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Amending the Broadcasting Licence of Premier Cablesystems Limited",
      "citation": "SI/84-231",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-84-145",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Issuing a Broadcasting Licence to Chinavision Canada Corporation",
      "citation": "SI/84-145",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-95",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting a Broadcasting Licence",
      "citation": "SI/96-95",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-25",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting a Broadcasting Licence",
      "citation": "SI/97-25",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-170",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting Al Cablesystems Inc.",
      "citation": "SI/90-170",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-56",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting Armadale Communications Limited",
      "citation": "SI/87-56",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-213",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting Campbell River T.V. Association",
      "citation": "SI/85-213",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-103",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting Canadian Learning Television",
      "citation": "SI/96-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-51",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting CFJO-FM",
      "citation": "SI/97-51",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-50",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting CFMB Limited",
      "citation": "SI/97-50",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-120",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting CHAM Hamilton",
      "citation": "SI/95-120",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-110",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting CHUM Limited",
      "citation": "SI/94-110",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-69",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting CHUM Limited",
      "citation": "SI/93-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-199",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting CHUM Limited",
      "citation": "SI/93-199",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-1",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting CION FM Inc. and CIBM-FM Mont-Bleu Ltée",
      "citation": "SI/89-1",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-118",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting CJFP-FM Rivière-du-Loup",
      "citation": "SI/96-118",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-117",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting CTEQ Télévision Inc.",
      "citation": "SI/95-117",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-139",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting KAYU-TV (Fox)",
      "citation": "SI/94-139",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-123",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting New Brunswick Broadcasting Company Ltd.",
      "citation": "SI/88-123",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-111",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting Premier Choix: TVEC Inc.",
      "citation": "SI/94-111",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-242",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting Racetrack Management (1983)",
      "citation": "SI/87-242",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-37",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting Radio Beauce Inc.",
      "citation": "SI/87-37",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-38",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting Radio Station CIWW Ottawa",
      "citation": "SI/95-38",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-154",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting Rawlco Communications Ltd.",
      "citation": "SI/90-154",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-39",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting Rogers Cable TV Limited, Shaw Cablesystems (B.C.) Ltd. and Shaw Cablesystems (Ontario) Ltd.",
      "citation": "SI/95-39",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-171",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting Shaw Cablesystems (B.C.) Ltd.",
      "citation": "SI/92-171",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-27",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting SkyCable Inc.",
      "citation": "SI/96-27",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-256",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting Sun Country Cablevision Ltd.",
      "citation": "SI/87-256",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-88",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting TELUS Cable Holdings Inc.",
      "citation": "SI/97-88",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-112",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting the Canadian Broadcasting Corporation",
      "citation": "SI/94-112",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-104",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting TreeHouse TV",
      "citation": "SI/96-104",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-119",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting Various Cable Distribution Undertakings",
      "citation": "SI/95-119",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-33",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting Victoria Communications Limited",
      "citation": "SI/88-33",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-143",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC a Decision Respecting Viking Cable T.V. Limited",
      "citation": "SI/87-143",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-76",
      "title": "Order declining to set aside or to refer back to the CRTC Certain Decisions",
      "citation": "SI/98-76",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-211",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Certain Decisions",
      "citation": "SI/86-211",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-2",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Certain Decisions",
      "citation": "SI/85-2",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-212",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Certain Decisions",
      "citation": "SI/85-212",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-215",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Certain Decisions",
      "citation": "SI/88-215",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-88-29",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Certain Decisions",
      "citation": "SI/88-29",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-119",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Certain Decisions Respecting Various Undertakings",
      "citation": "SI/96-119",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-25",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Certain Decisions Respecting Various Undertakings",
      "citation": "SI/96-25",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-26",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Certain Decisions Respecting Various Undertakings",
      "citation": "SI/96-26",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-28",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Certain Decisions Respecting Various Undertakings",
      "citation": "SI/96-28",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-87",
      "title": "Order declining to set aside or to refer back to the CRTC Decision CRTC 2000-205",
      "citation": "SI/2000-87",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-91",
      "title": "Order Declining to set Aside or to Refer Back to the CRTC Decision CRTC 2000-219",
      "citation": "SI/2000-91",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-43",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Decision CRTC 99-19",
      "citation": "SI/99-43",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-119",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Decision CRTC 97-293",
      "citation": "SI/97-119",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-120",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Decision CRTC 97-294",
      "citation": "SI/97-120",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-83",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Decision CRTC 2002-39",
      "citation": "SI/2002-83",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-94",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Decision CRTC 2002-81",
      "citation": "SI/2002-94",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-104",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Decision CRTC 2002-82",
      "citation": "SI/2002-104",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-86",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Decision CRTC 2000-203",
      "citation": "SI/2000-86",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-14",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Decision CRTC 2001-628",
      "citation": "SI/2002-14",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-41",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Decision CRTC 2001-678",
      "citation": "SI/2002-41",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-32",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Decision CRTC 2002-377",
      "citation": "SI/2003-32",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-15",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Decision CRTC 2004-538",
      "citation": "SI/2005-15",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-95",
      "title": "Order declining to Set Aside or to Refer Back to the CRTC Decisions CRTC 99-109 to CRTC 99-112",
      "citation": "SI/99-95",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-110",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Decisions CRTC 2001-457 and CRTC 2001-458",
      "citation": "SI/2001-110",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-81",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Decisions CRTC 2005-246 and CRTC 2005-247",
      "citation": "SI/2005-81",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-103",
      "title": "Order Declining to Set Aside or to Refer Back to the CRTC Decisions CRTC 2005-338 and CRTC 2005-339",
      "citation": "SI/2005-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-126",
      "title": "Order Declining to Set Aside or to Refer Decision CRTC 97-362",
      "citation": "SI/97-126",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-130",
      "title": "Order Declining to Set Aside or to Refer Decision CRTC 97-370",
      "citation": "SI/97-130",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-895",
      "title": "Order Setting Aside Amendments to Certain Broadcasting Licences Issued by the Canadian Radio-television and Telecommunications Commission",
      "citation": "SOR/82-895",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-194",
      "title": "Order Setting Aside Amendments to the Broadcasting Licence of M.S.A. Cablevision Ltd. Issued by the CRTC",
      "citation": "SI/83-194",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-195",
      "title": "Order Setting Aside Amendments to the Broadcasting Licence of Western Cablevision Limited Issued by the CRTC",
      "citation": "SI/83-195",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-76-153",
      "title": "Order Setting Aside Certain Broadcasting Licences Issued by the Canadian Radio-Television and Telecommunications Commission",
      "citation": "SI/76-153",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-790",
      "title": "Settlers\u0027 Effects Acquired with Blocked Currencies Remission Order",
      "citation": "CRC, c 790",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-195",
      "title": "Set Top Boxes Remission Order",
      "citation": "SOR/2006-195",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-247",
      "title": "Sex Offender Information Registration Regulations (Canadian Forces)",
      "citation": "SOR/2008-247",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-339",
      "title": "Shared Premises Regulations (Banks)",
      "citation": "SOR/2002-339",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-69",
      "title": "Shared Premises Regulations (Retail Associations)",
      "citation": "SOR/2008-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-68",
      "title": "Shared Premises Regulations (Trust and Loan Companies)",
      "citation": "SOR/2008-68",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-95",
      "title": "Order Transferring to Shared Services Canada the Control and Supervision of Certain Portions of the Federal Public Administration in each Department and Portion of the Federal Public Administration known as the Email, Data Centre and Network Services Unit and the Email, Data Centre and Network Services Support Unit",
      "citation": "SI/2011-95",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-43",
      "title": "Order Transferring to Shared Services Canada the control and supervision of the portion of the Federal Public Administration in the Acquisitions Branch of the Department of Public Works and Government Services known as the End User Devices Procurement Unit",
      "citation": "SI/2013-43",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-59",
      "title": "Order Transferring to Shared Services Canada the Control and Supervision of the Portion of the Federal Public Administration in the Acquisitions Branch of the Department of Public Works and Government Services known as the Information Technology Shared Services Procurement Directorate",
      "citation": "SI/2012-59",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-494",
      "title": "Sheep Stabilization 1981 Regulations",
      "citation": "SOR/82-494",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-414",
      "title": "Sheshatshiu Innu First Nation Band Order",
      "citation": "SOR/2002-414",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-348",
      "title": "Shipbuilding Industry Assistance Regulations",
      "citation": "CRC, c 348",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-349",
      "title": "Shipbuilding Temporary Assistance Program Regulations",
      "citation": "CRC, c 349",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-347",
      "title": "Ship Construction Subsidy Regulations",
      "citation": "CRC, c 347",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-106",
      "title": "Ship Fumigation Regulations",
      "citation": "SOR/89-106",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-514",
      "title": "Shipping Casualties Reporting Regulations",
      "citation": "SOR/85-514",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1479",
      "title": "Shipping Inquiries and Investigations Rules",
      "citation": "CRC, c 1479",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-356",
      "title": "Shipping Safety Control Zones Order",
      "citation": "CRC, c 356",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1472",
      "title": "Ship Radio Inspection Fees Regulations",
      "citation": "CRC, c 1472",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-70",
      "title": "Ship Registration and Tonnage Regulations",
      "citation": "SOR/2000-70",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1480",
      "title": "Ships\u0027 Crews Food and Catering Regulations",
      "citation": "CRC, c 1480",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1482",
      "title": "Ships\u0027 Elevator Regulations",
      "citation": "CRC, c 1482",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-40",
      "title": "Ships\u0027 Stores Regulations",
      "citation": "SOR/96-40",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-260",
      "title": "Ship Station (Radio) Regulations, 1999",
      "citation": "SOR/2000-260",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-265",
      "title": "Ship Station (Radio) Technical Regulations, 1999",
      "citation": "SOR/2000-265",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-87",
      "title": "Shirting Fabrics Remission Order, 1998",
      "citation": "SOR/98-87",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-212",
      "title": "Shooting Clubs and Shooting Ranges Regulations",
      "citation": "SOR/98-212",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-245",
      "title": "Short-term Pooled Investment Fund Regulations",
      "citation": "SOR/2006-245",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-791",
      "title": "Side Shows and Concessions Remission Order",
      "citation": "CRC, c 791",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-700",
      "title": "Sinai Multinational Force and Observers Privileges and Immunities Order, 1990",
      "citation": "SOR/90-700",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-27",
      "title": "Single-Entry Visitor Visa Fee (Sporting and Cultural Events) Remission Order",
      "citation": "SI/2001-27",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1501",
      "title": "Small Businesses Loans Regulations",
      "citation": "CRC, c 1501",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-169",
      "title": "Small Business Loans Regulations, 1993",
      "citation": "SOR/93-169",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1330",
      "title": "Small Craft Harbours Leasing Regulations",
      "citation": "CRC, c 1330",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1486",
      "title": "Small Fishing Vessel Inspection Regulations",
      "citation": "CRC, c 1486",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-498",
      "title": "Small Manufacturers or Producers Exemption Regulations",
      "citation": "SOR/82-498",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-572",
      "title": "Small Manufacturers Production Equipment Exemption Regulations",
      "citation": "SOR/82-572",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-255",
      "title": "Definition of Small Retransmission Systems Regulations",
      "citation": "SOR/89-255",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1487",
      "title": "Small Vessel Regulations",
      "citation": "CRC, c 1487",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-91",
      "title": "Small Vessel Regulations",
      "citation": "SOR/2010-91",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-916",
      "title": "Smithers Airport Zoning Regulations",
      "citation": "SOR/81-916",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-10",
      "title": "Order Amalgamating the Department of Social Development and the Department of Human Resources and Skills Development under the Authority of the Minister and the Deputy Minister",
      "citation": "SI/2006-10",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-41",
      "title": "Social Insurance Number Disclosure Regulations",
      "citation": "SOR/91-41",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-82",
      "title": "Social Insurance Number Regulations",
      "citation": "SOR/2013-82",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-464",
      "title": "Social Insurance Number Replacement Card Fees Order, 1988",
      "citation": "SOR/88-464",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-60",
      "title": "Social Security Tribunal Regulations",
      "citation": "SOR/2013-60",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2007-13",
      "title": "Softwood Lumber Products Charge on Duty Deposit Refunds Remission Order, No. 1",
      "citation": "SI/2007-13",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2007-14",
      "title": "Softwood Lumber Products Charge on Duty Deposit Refunds Remission Order, No. 2",
      "citation": "SI/2007-14",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-16",
      "title": "Softwood Lumber Products Export Allocations Regulations",
      "citation": "SOR/2007-16",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-317",
      "title": "Softwood Lumber Products Export Permit Fees Regulations",
      "citation": "SOR/96-317",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1295",
      "title": "Solicitations by Mail Regulations",
      "citation": "CRC, c 1295",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-215",
      "title": "Solicitor General of Canada Authority to Prescribe Fees or Charges Order (Police Information Retrieval System)",
      "citation": "SI/89-215",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-214",
      "title": "Solicitor General of Canada Authority to Prescribe Fees or Charges Order (RCMP Training Academy-Regina)",
      "citation": "SI/89-214",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-59",
      "title": "Solicitor General of Canada Authority to Prescribe Fees Order",
      "citation": "SI/95-59",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-275",
      "title": "Solvency Funding Relief Regulations",
      "citation": "SOR/2006-275",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-182",
      "title": "Solvency Funding Relief Regulations, 2009",
      "citation": "SOR/2009-182",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-283",
      "title": "Solvent Degreasing Regulations",
      "citation": "SOR/2003-283",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-163",
      "title": "Soquip Alberta Inc. (Petroleum and Gas Revenue Tax Act) Remission Order",
      "citation": "SI/90-163",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-17",
      "title": "Sour Cherry Stabilization Regulations, 1980",
      "citation": "SOR/82-17",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-214",
      "title": "Sour Cherry Stabilization Regulations, 1982",
      "citation": "SOR/83-214",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-381",
      "title": "Sour Cherry Stabilization Regulations, 1987",
      "citation": "SOR/88-381",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-405",
      "title": "Special Areas Industrial Renewal Order",
      "citation": "SOR/82-405",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-208",
      "title": "Special Authority to Possess Regulations (Firearms Act)",
      "citation": "SOR/98-208",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-631",
      "title": "Special Counting of Prior Service Regulations (Health Care Services for the Benefit of the Canadian Forces)",
      "citation": "SOR/90-631",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-496",
      "title": "Special Duty Area Pension Order",
      "citation": "SOR/2001-496",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-286",
      "title": "Special Economic Measures (Burma) Permit Authorization Order",
      "citation": "SOR/2007-286",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-285",
      "title": "Special Economic Measures (Burma) Regulations",
      "citation": "SOR/2007-285",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-168",
      "title": "Special Economic Measures (Democratic People\u0027s Republic of Korea) Permit Authorization Order",
      "citation": "SOR/2011-168",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-167",
      "title": "Special Economic Measures (Democratic People\u0027s Republic of Korea) Regulations",
      "citation": "SOR/2011-167",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-166",
      "title": "Special Economic Measures (Iran) Permit Authorization Order",
      "citation": "SOR/2010-166",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-165",
      "title": "Special Economic Measures (Iran) Regulations",
      "citation": "SOR/2010-165",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-52",
      "title": "Special Economic Measures (Libya) Permit Authorization Order",
      "citation": "SOR/2011-52",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-59",
      "title": "Special Economic Measures (Russia) Permit Authorization Order",
      "citation": "SOR/2014-59",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-58",
      "title": "Special Economic Measures (Russia) Regulations",
      "citation": "SOR/2014-58",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-236",
      "title": "Special Economic Measures (South Sudan) Permit Authorization Order",
      "citation": "SOR/2014-236",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-235",
      "title": "Special Economic Measures (South Sudan) Regulations",
      "citation": "SOR/2014-235",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-115",
      "title": "Special Economic Measures (Syria) Permit Authorization Order",
      "citation": "SOR/2011-115",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-18",
      "title": "Special Economic Measures (Syria) Permit Authorization Order",
      "citation": "SOR/2014-18",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-18",
      "title": "Order Repealing the Special Economic Measures (Syria) Permit Authorization Order",
      "citation": "SOR/2014-18",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-114",
      "title": "Special Economic Measures (Syria) Regulations",
      "citation": "SOR/2011-114",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-61",
      "title": "Special Economic Measures (Ukraine) Permit Authorization Order",
      "citation": "SOR/2014-61",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-60",
      "title": "Special Economic Measures (Ukraine) Regulations",
      "citation": "SOR/2014-60",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-249",
      "title": "Special Economic Measures (Zimbabwe) Permit Authorization Order",
      "citation": "SOR/2008-249",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-248",
      "title": "Special Economic Measures (Zimbabwe) Regulations",
      "citation": "SOR/2008-248",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1586",
      "title": "Special Force Superannuation Regulations",
      "citation": "CRC, c 1586",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-927",
      "title": "Special Import Measures Regulations",
      "citation": "SOR/84-927",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-477",
      "title": "Specialized Financing (Bank Holding Companies) Regulations",
      "citation": "SOR/2001-477",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-428",
      "title": "Specialized Financing (Banks) Regulations",
      "citation": "SOR/2001-428",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-427",
      "title": "Specialized Financing (Cooperative Credit Associations) Regulations",
      "citation": "SOR/2001-427",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-432",
      "title": "Specialized Financing (Foreign Banks) Regulations",
      "citation": "SOR/2001-432",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-478",
      "title": "Specialized Financing (Insurance Holding Companies) Regulations",
      "citation": "SOR/2001-478",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-429",
      "title": "Specialized Financing (Life Companies) Regulations",
      "citation": "SOR/2001-429",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-430",
      "title": "Specialized Financing (Retail Associations) Regulations",
      "citation": "SOR/2001-430",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-431",
      "title": "Specialized Financing (Trust and Loan Companies) Regulations",
      "citation": "SOR/2001-431",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-121",
      "title": "Special-purpose Vessels Regulations",
      "citation": "SOR/2008-121",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-124",
      "title": "Special Service Medal Bar Order \"ALERT\"",
      "citation": "SI/95-124",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-50",
      "title": "Special Service Medal Bar Order “EXPEDITION”",
      "citation": "SI/2014-50",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-24",
      "title": "Special Service Medal Bar Order \"Humanitas\"",
      "citation": "SI/97-24",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-126",
      "title": "Special Service Medal Bar Order \"NATO—OTAN\"",
      "citation": "SI/95-126",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-125",
      "title": "Special Service Medal Bar Order \"PEACE—PAIX\"",
      "citation": "SI/95-125",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-116",
      "title": "Special Service Medal Bar Order \"Ranger\"",
      "citation": "SI/99-116",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-218",
      "title": "Special Service Medal Clasp Order \"Pakistan 1989-90\"",
      "citation": "SI/92-218",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1296",
      "title": "Special Services and Fees Regulations",
      "citation": "CRC, c 1296",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1012",
      "title": "Special Services (Customs) Regulations",
      "citation": "SOR/86-1012",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-689",
      "title": "Special Services (Excise) Regulations",
      "citation": "SOR/87-689",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-478",
      "title": "Special Services Regulations",
      "citation": "CRC, c 478",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-956",
      "title": "Special Status Order, 1982",
      "citation": "SOR/82-956",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-148",
      "title": "Special Status Order, 1983",
      "citation": "SOR/83-148",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-106",
      "title": "Specialty Services Regulations, 1990",
      "citation": "SOR/90-106",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-148",
      "title": "Special Voting Rules",
      "citation": "SOR/78-148",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-133",
      "title": "Special Voting Rules Referendum Fees Tariff",
      "citation": "SI/92-133",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-677",
      "title": "Specific Agreement Confirmation Regulations",
      "citation": "SOR/92-677",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-101",
      "title": "Specification 112 and 114 Tank Cars Regulations",
      "citation": "SOR/79-101",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-734",
      "title": "Specification of Statutory Provisions (Treasury Board) Order",
      "citation": "CRC, c 734",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-482",
      "title": "Specifications and Books of Reference",
      "citation": "SOR/80-482",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-81",
      "title": "Specifications Relating to Non-automatic Weighing Devices (1998)",
      "citation": "SI/98-81",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-119",
      "title": "Specific Claims Tribunal Rules of Practice and Procedure",
      "citation": "SOR/2011-119",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-175",
      "title": "Specified Crown Agents (GST/HST) Regulations",
      "citation": "SOR/99-175",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-20",
      "title": "Specified Tangible Personal Property (GST/HST) Regulations",
      "citation": "SOR/91-20",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-6",
      "title": "Regulations Specifying Investigative Bodies",
      "citation": "SOR/2001-6",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-7",
      "title": "Regulations Specifying Publicly Available Information",
      "citation": "SOR/2001-7",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-67",
      "title": "Spence Bay Airport Zoning Regulations",
      "citation": "SOR/92-67",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-793",
      "title": "Spirit Destruction Remission Order",
      "citation": "CRC, c 793",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-115",
      "title": "Springbank Airport Zoning Regulations",
      "citation": "CRC, c 115",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-118",
      "title": "Spruce Falls Remission Order",
      "citation": "SI/90-118",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-131",
      "title": "Order Designating the Staff of the Non-Public Funds, Canadian Forces, a separate employer, for the purposes of paragraph 62(1)(a) of the Act",
      "citation": "SOR/2000-131",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-392",
      "title": "Stainless Steel Round Wire Products Anti-dumping Duty Remission Order",
      "citation": "SOR/2005-392",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-288",
      "title": "Stamping and Marking of Tobacco Products Regulations",
      "citation": "SOR/2003-288",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-66",
      "title": "Standard Trust Depositors Remission Order",
      "citation": "SI/91-66",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-948",
      "title": "St. Andrews Airport Zoning Regulations",
      "citation": "SOR/81-948",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-444",
      "title": "St. Andrew\u0027s Lock Regulations",
      "citation": "SOR/91-444",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-121",
      "title": "St. Anthony Airport Zoning Regulations",
      "citation": "SOR/94-121",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-611",
      "title": "St. Anthony Fisheries Limited Regulations",
      "citation": "SOR/82-611",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-181",
      "title": "Statement Limiting the Right to Equitable Remuneration of Certain Rome Convention or WPPT Countries",
      "citation": "SOR/2014-181",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-246",
      "title": "Statistical Data Disclosure Regulations",
      "citation": "SOR/2006-246",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-153a",
      "title": "Statistics Canada 1996 Census of Population Terms Exclusion Approval Order",
      "citation": "SOR/95-153a",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-153b",
      "title": "Statistics Canada 1996 Census of Population Terms Regulations",
      "citation": "SOR/95-153b",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-257",
      "title": "Statistics Canada 2006 Census Term Employees Exclusion Approval Order / Regulations on the employment with Statistics Canada for the purpose of the 2006 Census",
      "citation": "SOR/2004-257",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2010-81",
      "title": "Statistics Canada Census and Survey Related Term Employment Exclusion Approval Order",
      "citation": "SI/2010-81",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-147",
      "title": "Statistics Canada Census and Survey Related Term Employment Regulations",
      "citation": "SOR/2010-147",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-147",
      "title": "Statistics Canada Census and Survey Related Term Employment Regulations",
      "citation": "SOR/2010-147",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-147",
      "title": "Statistics Canada Census and Survey Related Term Employment Regulations",
      "citation": "SOR/2010-147",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2010-45",
      "title": "Statistics Canada Census-Related Term Employment Exclusion Approval Order",
      "citation": "SI/2010-45",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-176",
      "title": "Status of the Artist Act Procedural Regulations",
      "citation": "SOR/2014-176",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-191",
      "title": "Status of the Artist Act Professional Category Regulations",
      "citation": "SOR/99-191",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-83-207",
      "title": "Statutes of Canada Distribution Direction",
      "citation": "SI/83-207",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-89-140",
      "title": "Statutes of Canada Distribution Direction, No. 2",
      "citation": "SI/89-140",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-158",
      "title": "Statutes of Canada Loose-leaf Consolidation Updates Distribution Order",
      "citation": "SI/92-158",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1509",
      "title": "Statutory Instruments Regulations",
      "citation": "CRC, c 1509",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-901",
      "title": "St. Catharines Airport Zoning Regulations",
      "citation": "SOR/84-901",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-335",
      "title": "St. Clair and Detroit River Navigation Safety Regulations",
      "citation": "SOR/84-335",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-810",
      "title": "Steering Appliances and Equipment Regulations",
      "citation": "SOR/83-810",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-71",
      "title": "Stelco Steel Remission Order",
      "citation": "SI/91-71",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-449",
      "title": "Stephenville Airport Zoning Regulations",
      "citation": "SOR/87-449",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-111",
      "title": "St. Hubert Airport Zoning Regulations",
      "citation": "CRC, c 111",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-112",
      "title": "St. Johns (Quebec) Airport Zoning Regulations",
      "citation": "CRC, c 112",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-113",
      "title": "St. John\u0027s (Torbay) Airport Zoning Regulations",
      "citation": "CRC, c 113",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-144",
      "title": "St. Lawrence Seaway Authority Divestiture Regulations",
      "citation": "SOR/99-144",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-522",
      "title": "St. Leonard Airport Zoning Regulations",
      "citation": "SOR/93-522",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-239",
      "title": "St. Marys Paper Inc. Regulations",
      "citation": "SOR/87-239",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-80",
      "title": "Storage Charges Remission Order, 1993",
      "citation": "SI/93-80",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-210",
      "title": "Storage, Display and Transportation of Firearms and Other Weapons by Businesses Regulations",
      "citation": "SOR/98-210",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-209",
      "title": "Storage, Display, Transportation and Handling of Firearms by Individuals Regulations",
      "citation": "SOR/98-209",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-991",
      "title": "Storage of Goods Regulations",
      "citation": "SOR/86-991",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-507",
      "title": "Storage of PCB Material Regulations",
      "citation": "SOR/92-507",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-197",
      "title": "Storage Tank Systems for Petroleum Products and Allied Petroleum Products Regulations",
      "citation": "SOR/2008-197",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-214",
      "title": "Stranded Steel Wire Remission Order",
      "citation": "SOR/2000-214",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-495",
      "title": "Stratford Municipal Airport Zoning Regulations",
      "citation": "SOR/95-495",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-51",
      "title": "Streamlined Accounting (GST/HST) Regulations",
      "citation": "SOR/91-51",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2010-46",
      "title": "Student Employment Programs Participants Exclusion Approval Order",
      "citation": "SI/2010-46",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-148",
      "title": "Student Employment Programs Participants Regulations",
      "citation": "SOR/2010-148",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-194",
      "title": "Student Employment Programs Regulations",
      "citation": "SOR/97-194",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-872",
      "title": "Submission of Samples of Fabrics for Importation Regulations",
      "citation": "SOR/86-872",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-314",
      "title": "Subsidiaries Holding Association Shares (Cooperative Credit Associations) Regulations",
      "citation": "SOR/92-314",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-313",
      "title": "Subsidiaries Holding Bank Shares (Banks) Regulations",
      "citation": "SOR/92-313",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-315",
      "title": "Subsidiaries Holding Company Shares (Insurance Companies) Regulations",
      "citation": "SOR/92-315",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-297",
      "title": "Subsidiaries Holding Company Shares (Trust and Loan Companies) Regulations",
      "citation": "SOR/92-297",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-433",
      "title": "Subsidiaries that Hold Bank Holding Company Shares Regulations",
      "citation": "SOR/2001-433",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-434",
      "title": "Subsidiaries that Hold Insurance Holding Company Shares Regulations",
      "citation": "SOR/2001-434",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-116",
      "title": "Sudbury Airport Zoning Regulations",
      "citation": "CRC, c 116",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-379",
      "title": "Sugar Beet Stabilization Regulations, 1977",
      "citation": "SOR/79-379",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-736",
      "title": "Sugar Beet Stabilization Regulations, 1981",
      "citation": "SOR/83-736",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-665",
      "title": "Sugar Beet Stabilization Regulations, 1982",
      "citation": "SOR/84-665",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-28",
      "title": "Sugar Beet Stabilization Regulations, 1985",
      "citation": "SOR/86-28",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-278",
      "title": "Sugar Beet Stabilization Regulations, 1986",
      "citation": "SOR/87-278",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-254",
      "title": "Sulphur in Diesel Fuel Regulations",
      "citation": "SOR/2002-254",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-236",
      "title": "Sulphur in Gasoline Regulations",
      "citation": "SOR/99-236",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-110",
      "title": "Notice Respecting the 2010 G8 and G20 Summits Railway Transportation Security Measures",
      "citation": "SOR/2010-110",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-58",
      "title": "Supervisory Information (Authorized Foreign Banks) Regulations",
      "citation": "SOR/2001-58",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-480",
      "title": "Supervisory Information (Bank Holding Companies) Regulations",
      "citation": "SOR/2001-480",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-59",
      "title": "Supervisory Information (Banks) Regulations",
      "citation": "SOR/2001-59",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-57",
      "title": "Supervisory Information (Cooperative Credit Associations) Regulations",
      "citation": "SOR/2001-57",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-56",
      "title": "Supervisory Information (Insurance Companies) Regulations",
      "citation": "SOR/2001-56",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-484",
      "title": "Supervisory Information (Insurance Holding Companies) Regulations",
      "citation": "SOR/2001-484",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-55",
      "title": "Supervisory Information (Trust and Loan Companies) Regulations",
      "citation": "SOR/2001-55",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-100",
      "title": "Proclamation Declaring the Supplementary Agreement to the Agreement on Social Security between Canada and the Republic of Austria in Force on December 1, 1996",
      "citation": "SI/96-100",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1360",
      "title": "Supplementary Death Benefit Regulations",
      "citation": "CRC, c 1360",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1511",
      "title": "Supplementary Retirement Benefits Regulations",
      "citation": "CRC, c 1511",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-264",
      "title": "Support Orders and Support Provisions (Banks and Authorized Foreign Banks) Regulations",
      "citation": "SOR/2002-264",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-265",
      "title": "Support Orders and Support Provisions (Retail Associations) Regulations",
      "citation": "SOR/2002-265",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-266",
      "title": "Support Orders and Support Provisions (Trust and Loan Companies) Regulations",
      "citation": "SOR/2002-266",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-62-91",
      "title": "Support Organizations Superannuation Regulations",
      "citation": "SOR/62-91",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-96",
      "title": "Supreme Court of Newfoundland and Labrador — Court of Appeal Criminal Appeal Rules (2002)",
      "citation": "SI/2002-96",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-33",
      "title": "Supreme Court of Newfoundland, Trial Division Rules for Orders in the Nature of Certiorari, Habeas Corpus, Mandamus and Prohibition",
      "citation": "SI/2000-33",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-64",
      "title": "Supreme Court of Yukon Summary Conviction Appeal Rules, 2009",
      "citation": "SI/2012-64",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-109",
      "title": "Surface Coating Materials Regulations",
      "citation": "SOR/2005-109",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-12",
      "title": "Surtax on Boneless Beef Order No. 2, 1994",
      "citation": "SOR/94-12",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-739",
      "title": "Swan River Airport Zoning Regulations",
      "citation": "SOR/90-739",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-151",
      "title": "Sweet Cherry Stabilization Regulations",
      "citation": "SOR/78-151",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-408",
      "title": "Swift Current Airport Zoning Regulations",
      "citation": "SOR/93-408",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-117",
      "title": "Sydney Airport Zoning Regulations",
      "citation": "CRC, c 117",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-67-56",
      "title": "Sydney Pilots\u0027 Pension Regulations",
      "citation": "SOR/67-56",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-794",
      "title": "Syncrude Remission Order",
      "citation": "CRC, c 794",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-31",
      "title": "Systems Software Development Contract Remission Order",
      "citation": "SI/86-31",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1494",
      "title": "Tackle Regulations",
      "citation": "CRC, c 1494",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-291",
      "title": "Tailored Collar Shirts Remission Order, 1997",
      "citation": "SOR/97-291",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-256",
      "title": "Tariff Classification Advance Rulings Regulations",
      "citation": "SOR/2005-256",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-701",
      "title": "Tariff Item No. 9805.00.00 Exemption Order",
      "citation": "SOR/81-701",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-225",
      "title": "Tariff Item No. 9807.00.00 Exemption Order",
      "citation": "SOR/90-225",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-47",
      "title": "Tariff Item Nos. 9971.00.00 and 9992.00.00 Accounting Regulations",
      "citation": "SOR/98-47",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-308",
      "title": "Tariff of Costs",
      "citation": "SOR/99-308",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1495",
      "title": "Tariff of Fees of Shipping Masters",
      "citation": "CRC, c 1495",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-190",
      "title": "Tarium Niryutait Marine Protected Areas Regulations",
      "citation": "SOR/2010-190",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-53",
      "title": "Order Giving Notice that a Tax Agreement Between Canada and Papua New Guinea Came into Force on December 21, 1989",
      "citation": "SI/91-53",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-52",
      "title": "Order Giving Notice that a Tax Agreement Between Canada and the Polish People\u0027s Republic Came into Force on November 30, 1989",
      "citation": "SI/91-52",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-156",
      "title": "Taxation Statistical Analyses and Data Processing Services Fees Order",
      "citation": "SOR/92-156",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-94",
      "title": "Taxation Statistics Diskette Fee Order",
      "citation": "SOR/91-94",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-218",
      "title": "Tax Collection Agreements and Federal Post-Secondary Education and Health Contributions Regulations, 1987",
      "citation": "SOR/87-218",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-108",
      "title": "Order Giving Notice that a Tax Convention Between Canada and the Grand Duchy of Luxembourg Came into Force on July 8, 1991",
      "citation": "SI/91-108",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-110",
      "title": "Order Giving Notice that a Tax Convention Between Canada and the United Mexican States Came into Force May 11, 1992",
      "citation": "SI/92-110",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-688a",
      "title": "Tax Court of Canada Rules (General Procedure)",
      "citation": "SOR/90-688a",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-688b",
      "title": "Tax Court of Canada Rules (Informal Procedure)",
      "citation": "SOR/90-688b",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-689",
      "title": "Tax Court of Canada Rules of Procedure respecting the Canada Pension Plan",
      "citation": "SOR/90-689",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-99",
      "title": "Tax Court of Canada Rules of Procedure Respecting the Customs Act (Informal Procedure)",
      "citation": "SOR/2004-99",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-690",
      "title": "Tax Court of Canada Rules of Procedure respecting the Employment Insurance Act",
      "citation": "SOR/90-690",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-102",
      "title": "Tax Court of Canada Rules of Procedure respecting the Excise Act, 2001 (Informal Procedure)",
      "citation": "SOR/2004-102",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-34",
      "title": "Taxes, Duties and Fees (GST/HST) Regulations",
      "citation": "SOR/91-34",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-109",
      "title": "Order Giving Notice that a Tax Information Exchange Convention Between Canada and the United Mexican States Came into Force April 27, 1992",
      "citation": "SI/92-109",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-108",
      "title": "Tax Rebate Discounting Regulations",
      "citation": "SOR/86-108",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-475",
      "title": "Technical Assistance Regulations",
      "citation": "SOR/86-475",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-345",
      "title": "Technical Data Control Regulations",
      "citation": "SOR/86-345",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-143",
      "title": "Order Varying Telecom Decision CRTC 94-19",
      "citation": "SI/94-143",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-322",
      "title": "Order Varying Telecom Decision CRTC 95-14 and Requiring the CRTC to Report on the Matter of Directory Subscriber Listings",
      "citation": "SOR/96-322",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-8",
      "title": "Order Varying Telecom Decision CRTC 95-21",
      "citation": "SOR/96-8",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-288",
      "title": "Order Varying Telecom Decision CRTC 2005-28",
      "citation": "SOR/2006-288",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-71",
      "title": "Order Varying Telecom Decision CRTC 2006-15",
      "citation": "SOR/2007-71",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-895",
      "title": "Telecommunication Programming Services Tax Regulations",
      "citation": "SOR/86-895",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-319",
      "title": "Telecommunications and Electronic Maintenance Services Fees Order",
      "citation": "SOR/81-319",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-365",
      "title": "Telecommunications Apparatus Assessment and Testing Fees Order",
      "citation": "SOR/93-365",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-532",
      "title": "Telecommunications Apparatus Regulations",
      "citation": "SOR/2001-532",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-157",
      "title": "Telecommunications Fees Regulations, 1995",
      "citation": "SOR/95-157",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-65",
      "title": "Telecommunications Fees Regulations, 2010",
      "citation": "SOR/2010-65",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-379",
      "title": "Teleglobe Canada Pension Protection Regulations",
      "citation": "SOR/87-379",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1225",
      "title": "Telegraph and Cable Messages Terms and Conditions Order",
      "citation": "CRC, c 1225",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-90-117",
      "title": "Telesat Canada Financing Remission Order",
      "citation": "SI/90-117",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-82",
      "title": "Telesat Canada Remission Order",
      "citation": "SI/99-82",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-49",
      "title": "Television Broadcasting Regulations, 1987",
      "citation": "SOR/87-49",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-799",
      "title": "Temporary Export of Aircraft Remission Order",
      "citation": "CRC, c 799",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-427",
      "title": "Temporary Importation (Excise Levies and Additional Duties) Regulations",
      "citation": "SOR/89-427",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-840",
      "title": "Temporary Importation of Conveyances by Residents of Canada Regulations",
      "citation": "SOR/82-840",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-122",
      "title": "Temporary Importation of Vessels Remission Order, No. 9",
      "citation": "SI/95-122",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-352",
      "title": "Temporary Importation Remission Order, No. 1 (Customs Tariff)",
      "citation": "SOR/2005-352",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-58",
      "title": "Temporary Importation (Tariff Item No. 9993.00.00) Regulations",
      "citation": "SOR/98-58",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-79",
      "title": "Temporary Storage Period Regulations",
      "citation": "SOR/88-79",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-54",
      "title": "Terms and Conditions of Employment of the Federal Ombudsman for Victims of Crime",
      "citation": "SOR/2007-54",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-577",
      "title": "Terra Nova Telecommunications Inc. Shares Sale Order",
      "citation": "SOR/88-577",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1522",
      "title": "Territorial Coal Regulations",
      "citation": "CRC, c 1522",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1523",
      "title": "Territorial Dredging Regulations",
      "citation": "CRC, c 1523",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1526",
      "title": "Territorial Lands Act Exclusion Order",
      "citation": "CRC, c 1526",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1525",
      "title": "Territorial Lands Regulations",
      "citation": "CRC, c 1525",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1524",
      "title": "Territorial Land Use Regulations",
      "citation": "CRC, c 1524",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1527",
      "title": "Territorial Quarrying Regulations",
      "citation": "CRC, c 1527",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-872",
      "title": "Territorial Sea Geographical Coordinates (Area 7) Order",
      "citation": "SOR/85-872",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1550",
      "title": "Territorial Sea Geographical Coordinates Order",
      "citation": "CRC, c 1550",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-61",
      "title": "Regulations prescribing a territory for the purposes of the definition “country” in the Customs Tariff",
      "citation": "SOR/97-61",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-591",
      "title": "Teslin Airport Zoning Regulations",
      "citation": "SOR/94-591",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-74-431",
      "title": "Teslin Exploration Limited Order",
      "citation": "SOR/74-431",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-104",
      "title": "Teslin Tlingit Council (GST) Remission Order",
      "citation": "SI/2000-104",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2007-87",
      "title": "Testing of Unregistered US Military Herbicides, including Agent Orange, at CFB Gagetown Ex Gratia Payments Order",
      "citation": "SI/2007-87",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-79",
      "title": "Tetrachloroethylene (Use in Dry Cleaning and Reporting Requirements) Regulations",
      "citation": "SOR/2003-79",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-581",
      "title": "Textile and Apparel Extension of Benefit Order",
      "citation": "SOR/93-581",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-22",
      "title": "Textile Flammability Regulations",
      "citation": "SOR/2011-22",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1551",
      "title": "Textile Labelling and Advertising Regulations",
      "citation": "CRC, c 1551",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-278",
      "title": "Textiles and Apparel Remission Order, 2014",
      "citation": "SOR/2014-278",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-25",
      "title": "Order Establishing the Text of a Resolution Providing for the Extension of the Application of Sections 83.28, 83.29 and 83.3 of the Criminal Code",
      "citation": "SOR/2007-25",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-87-28",
      "title": "The Criminal Appeal Rules of the Supreme Court of Newfoundland, Trial Division",
      "citation": "SI/87-28",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-60",
      "title": "The de Havilland Shares Sale Order",
      "citation": "SOR/86-60",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-82-627",
      "title": "The Lake Group Ltd. Regulations",
      "citation": "SOR/82-627",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-410",
      "title": "The Pas Airport Zoning Regulations",
      "citation": "SOR/93-410",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-501",
      "title": "Theratronics International Limited Authorization Order, 1988",
      "citation": "SOR/88-501",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-308",
      "title": "Theratronics International Limited Divestiture Regulations",
      "citation": "SOR/98-308",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-172",
      "title": "Thompson Airport Zoning Regulations",
      "citation": "SOR/91-172",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-182",
      "title": "Proclamation establishing three different time zones in Nunavut, for the purposes of the definition of \"standard time\" in the Interpretation Act",
      "citation": "SOR/2001-182",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-972",
      "title": "Thunder Bay Harbour Commission Administrative By-Law",
      "citation": "SOR/86-972",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-974",
      "title": "Thunder Bay Harbour Tariff By-Law",
      "citation": "SOR/86-974",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1496",
      "title": "Timber Cargo Regulations",
      "citation": "CRC, c 1496",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1552",
      "title": "Timber Marking Rules",
      "citation": "CRC, c 1552",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-118",
      "title": "Timber Regulations, 1993",
      "citation": "SOR/94-118",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-115",
      "title": "Order Extending the Time for the Assessment of the Status of Wildlife Species",
      "citation": "SOR/2006-115",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-215",
      "title": "Order Extending the Time for the Assessment of the Status of Wildlife Species",
      "citation": "SOR/2003-215",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-59",
      "title": "Time Limit for the Application of Subsection 118(1) of the Customs Tariff Regulations",
      "citation": "SOR/98-59",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-139",
      "title": "Establishing Timelines for Comprehensive Studies Regulations",
      "citation": "SOR/2011-139",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-118",
      "title": "Timmins Airport Zoning Regulations",
      "citation": "CRC, c 118",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-93",
      "title": "Tobacco (Access) Regulations",
      "citation": "SOR/99-93",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-272",
      "title": "Tobacco Products Information Regulations",
      "citation": "SOR/2000-272",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-177",
      "title": "Tobacco Products Labelling Regulations (Cigarettes and Little Cigars)",
      "citation": "SOR/2011-177",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-273",
      "title": "Tobacco Reporting Regulations",
      "citation": "SOR/2000-273",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-94",
      "title": "Tobacco (Seizure and Restoration) Regulations",
      "citation": "SOR/99-94",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-319",
      "title": "Toll Information Regulations",
      "citation": "SOR/79-319",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-226",
      "title": "Proclamation Exempting Tom MacKay Lake from the Operation of Section 22 of the Act",
      "citation": "SOR/2005-226",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-332",
      "title": "Toronto Area Rail Transportation of Dangerous Goods Advisory Council Order",
      "citation": "SOR/86-332",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-148",
      "title": "Toronto/Buttonville Airport Zoning Regulations",
      "citation": "SOR/88-148",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-515",
      "title": "Toronto Island Airport Zoning Regulations",
      "citation": "SOR/85-515",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-123",
      "title": "Toronto/Lester B. Pearson International Airport Zoning Regulations",
      "citation": "SOR/99-123",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-120",
      "title": "Toronto Port Authority Regulations",
      "citation": "SOR/2005-120",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-435",
      "title": "Total Assets for Public Holding Requirements (Trust and Loan Companies) Regulations",
      "citation": "SOR/2001-435",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-436",
      "title": "Total Assets for Supervisability and Public Holding Requirements (Banks and Bank Holding Companies) Regulations",
      "citation": "SOR/2001-436",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-437",
      "title": "Total Assets for Supervisability and Public Holding Requirements (Insurance Companies and Insurance Holding Companies) Regulations",
      "citation": "SOR/2001-437",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1498",
      "title": "Towboat Crew Accommodation Regulations",
      "citation": "CRC, c 1498",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-174",
      "title": "Town of Jasper Streetworks Taxes Regulations",
      "citation": "SOR/88-174",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1111",
      "title": "Town of Jasper Zoning Regulations",
      "citation": "CRC, c 1111",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-17",
      "title": "Toys Regulations",
      "citation": "SOR/2011-17",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-195",
      "title": "Trade-marks Regulations",
      "citation": "SOR/96-195",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1560",
      "title": "Trade Unions Regulations",
      "citation": "CRC, c 1560",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-102",
      "title": "Traffic on the Land Side of Airports Regulations",
      "citation": "SOR/2006-102",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-58",
      "title": "Transfer of a Portion of the Canadian Food Inspection Agency Regulations",
      "citation": "SOR/2013-58",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-109",
      "title": "Transfer of Duties Order",
      "citation": "SI/2013-109",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-96",
      "title": "Transfer of Duties Order (International Experience Canada Program Unit)",
      "citation": "SI/2013-96",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-11",
      "title": "Transfer of Duties Order (Unit Responsible for Federal Bridges in the Region of Montreal)",
      "citation": "SI/2014-11",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-735",
      "title": "Transfer of Military Films to NATO Order",
      "citation": "CRC, c 735",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-246",
      "title": "Transfer of Portions of the Canada Revenue Agency Regulations",
      "citation": "SOR/2011-246",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-430",
      "title": "Transfer of Portions of the Canadian Food Inspection Agency Regulations",
      "citation": "SOR/2003-430",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-247",
      "title": "Transfer of Portions of the Canadian Food Inspection Agency Regulations",
      "citation": "SOR/2011-247",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-248",
      "title": "Transfer of Portions of the Canadian Nuclear Safety Commission Regulations",
      "citation": "SOR/2011-248",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-249",
      "title": "Transfer of Portions of the Financial Transactions and Reports Analysis Centre of Canada Regulations",
      "citation": "SOR/2011-249",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-251",
      "title": "Transfer of Portions of the National Research Council of Canada Regulations",
      "citation": "SOR/2011-251",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-250",
      "title": "Transfer of Portions of the Parks Canada Agency Regulations",
      "citation": "SOR/2011-250",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-736",
      "title": "Transfer of Royal Military College Academic Hoods Order",
      "citation": "CRC, c 736",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-254",
      "title": "Transfer of the Communications Security Establishment in the Department of National Defence Regulations",
      "citation": "SOR/2011-254",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-36",
      "title": "Transfer of the Crown Corporation Secretariat Regulations (Royal Canadian Mint and Canada Post Corporation)",
      "citation": "SOR/2006-36",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2003-429",
      "title": "Transfer of the Customs Services Regulations",
      "citation": "SOR/2003-429",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-73",
      "title": "Order Transferring from the Minister of Indian Affairs and Northern Development to the Deputy Prime Minister and Minister of State the Control and Supervision of the Office of Indian Residential Schools Resolution of Canada",
      "citation": "SI/2001-73",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-52",
      "title": "Order Transferring to the Minister of Energy, Mines and Resources, the Powers, Duties and Functions of the Minister of Industry, Science and Technology Under Sections 9 and 15 of the Department of Industry, Science and Technology Act Respecting the Hibernia Development Project",
      "citation": "SI/93-52",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-606",
      "title": "Transhipment Regulations",
      "citation": "CRC, c 606",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-65-410",
      "title": "Transitional Assistance Benefit Regulations",
      "citation": "SOR/65-410",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1561",
      "title": "Translation Bureau Regulations",
      "citation": "CRC, c 1561",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-130",
      "title": "Transportation Appeal Tribunal of Canada Certificate Regulations",
      "citation": "SOR/2004-130",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-334",
      "title": "Transportation Information Regulations",
      "citation": "SOR/96-334",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-334",
      "title": "Transportation Information Regulations",
      "citation": "SOR/96-334",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-153",
      "title": "Transportation of Dangerous Goods General Policy Advisory Council Order",
      "citation": "SOR/90-153",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-286",
      "title": "Transportation of Dangerous Goods Regulations",
      "citation": "SOR/2001-286",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-34",
      "title": "Transportation of Dangerous Goods Regulations",
      "citation": "SOR/2008-34",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1064",
      "title": "Transportation of Goods Regulations",
      "citation": "SOR/86-1064",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-446",
      "title": "Transportation Safety Board Regulations",
      "citation": "SOR/92-446",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-37",
      "title": "Transportation Safety Board Regulations",
      "citation": "SOR/2014-37",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1566",
      "title": "Transport Control Regulations",
      "citation": "CRC, c 1566",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-1123",
      "title": "Treasury Board Delegation of Powers Order",
      "citation": "SOR/86-1123",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-232",
      "title": "Order Transferring from the Treasury Board Secretariat to the Department of Public Works and Government Services Certain Portions of the Chief Information Officer Branch of the Treasury Board Secretariat",
      "citation": "SI/2003-232",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-6",
      "title": "Order Transferring from the Treasury Board to the Department of Public Works and Government Services the Control and Supervision of the Government Travel Modernization Office",
      "citation": "SI/2004-6",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-1",
      "title": "Treaty Land Entitlement (Manitoba) Remission Order",
      "citation": "SI/2001-1",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-47",
      "title": "Treaty Land Entitlement (Saskatchewan) Remission Order",
      "citation": "SI/94-47",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-401",
      "title": "Trenton Airport Zoning Regulations",
      "citation": "SOR/96-401",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-66",
      "title": "Tributyltetradecylphosphonium Chloride Regulations",
      "citation": "SOR/2000-66",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-325",
      "title": "Trident Aircraft Ltd. Regulations",
      "citation": "SOR/80-325",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-99",
      "title": "Tr\u0027ondëk Hwëch\u0027in (GST) Remission Order",
      "citation": "SI/2000-99",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-347",
      "title": "Tuktoyaktuk Airport Zoning Regulations",
      "citation": "SOR/95-347",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1298",
      "title": "Undeliverable and Redirected Mail Regulations",
      "citation": "CRC, c 1298",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-83-813",
      "title": "Unemployment Insurance Account Advance Regulations",
      "citation": "SOR/83-813",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-54",
      "title": "Unemployment Insurance Benefits Order, 1990",
      "citation": "SOR/90-54",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-74",
      "title": "Unemployment Insurance Benefits Order, 1991",
      "citation": "SOR/91-74",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1575",
      "title": "Unemployment Insurance (Collection of Premiums) Regulations",
      "citation": "CRC, c 1575",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-98",
      "title": "Unemployment Insurance Premiums (Spousal Employment) Remission Order",
      "citation": "SI/92-98",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-588",
      "title": "Unemployment Insurance Rate of Premium, 1996, Order",
      "citation": "SOR/95-588",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-59",
      "title": "Unemployment Insurance Rates of Premium, 1990, Order",
      "citation": "SOR/90-59",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1319",
      "title": "UNESCO Remission Order",
      "citation": "CRC, c 1319",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-82-42",
      "title": "Unfinished Leather Remission Order",
      "citation": "SI/82-42",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-388",
      "title": "Notices of Uninsured Deposits Regulations (Banks)",
      "citation": "SOR/99-388",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-65",
      "title": "Notices of Uninsured Deposits Regulations (Retail Associations)",
      "citation": "SOR/2008-65",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-64",
      "title": "Notices of Uninsured Deposits Regulations (Trust and Loan Companies)",
      "citation": "SOR/2008-64",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-15",
      "title": "United Kingdom Service Aircraft Over-Flight Regulations",
      "citation": "CRC, c 15",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-444",
      "title": "United Nations Al-Qaida and Taliban Regulations",
      "citation": "SOR/99-444",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-99-444",
      "title": "United Nations Al-Qaida and Taliban Regulations",
      "citation": "SOR/99-444",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-91",
      "title": "United Nations Assistance Mission in Rwanda (UNAMIR) Medal Order",
      "citation": "SI/94-91",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-127",
      "title": "United Nations Côte d\u0027Ivoire Regulations",
      "citation": "SOR/2005-127",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-222",
      "title": "United Nations Democratic Republic of the Congo Regulations",
      "citation": "SOR/2004-222",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-291",
      "title": "United Nations Framework Convention on Climate Change Privileges and Immunities Order",
      "citation": "SOR/2005-291",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-136",
      "title": "United Nations Headquarters Medal for Special Service Order (New York City)",
      "citation": "SI/99-136",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-221",
      "title": "United Nations Iraq Regulations",
      "citation": "SOR/2004-221",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2001-261",
      "title": "United Nations Liberia Regulations",
      "citation": "SOR/2001-261",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-30",
      "title": "United Nations Medal for Special Service Order",
      "citation": "SI/98-30",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-106",
      "title": "United Nations Medal for Special Service Order",
      "citation": "SI/97-106",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-127",
      "title": "United Nations Medal for Special Services Order (MINURCA)",
      "citation": "SI/99-127",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-11",
      "title": "United Nations Mission in Bosnia-Herzegovina Medal Order",
      "citation": "SI/2000-11",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-37",
      "title": "United Nations Mission in East Timor and the United Nations Transitional Administration in East Timor (UNAMET/UNTAET) Medal Order",
      "citation": "SI/2001-37",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-104",
      "title": "United Nations Mission in Haiti (UNMIH) Medal Order",
      "citation": "SI/95-104",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-13",
      "title": "United Nations Mission in Kosovo Medal Order",
      "citation": "SI/2000-13",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-36",
      "title": "United Nations Mission in Sierra Leone (UNAMSIL) Medal Order",
      "citation": "SI/2001-36",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-12",
      "title": "United Nations Mission of Observers in Prevlaka (Croatia) Medal Order",
      "citation": "SI/2000-12",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-103",
      "title": "United Nations Observer Mission in Uganda/Rwanda (UNOMUR) Medal Order",
      "citation": "SI/95-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-86",
      "title": "United Nations Operation in Côte d\u0027Ivoire (UNOCI) Medal Order",
      "citation": "SI/2006-86",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-90",
      "title": "United Nations Operation in Mozambique (UNOMOZ) Medal Order",
      "citation": "SI/94-90",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-94-55",
      "title": "United Nations Operation in Somalia (UNOSOM) Medal Order",
      "citation": "SI/94-55",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-87",
      "title": "United Nations Organization Mission in Ethiopia and Eritrea (UNMEE) Medal Order",
      "citation": "SI/2001-87",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-38",
      "title": "United Nations Organization Mission in the Democratic Republic of the Congo (MONUC) Medal Order",
      "citation": "SI/2001-38",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-10",
      "title": "United Nations Preventive Deployment Force (Macedonia) Medal Order",
      "citation": "SI/2000-10",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1320",
      "title": "United Nations Remission Order",
      "citation": "CRC, c 1320",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-51",
      "title": "Regulations Implementing the United Nations Resolution on Libya and Taking Special Economic Measures",
      "citation": "SOR/2011-51",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-92",
      "title": "Regulations Implementing the United Nations Resolutions on Somalia",
      "citation": "SOR/2009-92",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-163",
      "title": "Regulations Implementing the United Nations Resolutions on the Central African Republic",
      "citation": "SOR/2014-163",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-94-582",
      "title": "United Nations Rwanda Regulations",
      "citation": "SOR/94-582",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-400",
      "title": "United Nations Sierra Leone Regulations",
      "citation": "SOR/98-400",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-90",
      "title": "United Nations Stabilization Mission in Haiti (MINUSTAH) Medal Order",
      "citation": "SI/2005-90",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2004-197",
      "title": "United Nations Sudan Regulations",
      "citation": "SOR/2004-197",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-97-506",
      "title": "United States Barley and Barley Products Remission Order",
      "citation": "SOR/97-506",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-86-41",
      "title": "United States Diplomatic and Consular Staff Remission Order",
      "citation": "SI/86-41",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-16",
      "title": "United States Service Aircraft Over-Flight Regulations",
      "citation": "CRC, c 16",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-78-179",
      "title": "Declaring the United States to be a Reciprocating Country for Purposes of the Act",
      "citation": "SI/78-179",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-7",
      "title": "Unsolicited Telecommunications Fees Regulations",
      "citation": "SOR/2013-7",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-206",
      "title": "Uranium Mines and Mills Regulations",
      "citation": "SOR/2000-206",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-181",
      "title": "Uranium Mines (Ontario) Employment Exclusion Order",
      "citation": "SOR/87-181",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-435",
      "title": "Uranium Mines (Ontario) Occupational Health and Safety Regulations",
      "citation": "SOR/84-435",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1385",
      "title": "Urban Development and Transportation Plans Regulations",
      "citation": "CRC, c 1385",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-549",
      "title": "Used Mattress Materials Regulations",
      "citation": "CRC, c 549",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-42",
      "title": "Used or Second-hand Motor Vehicles Regulations",
      "citation": "SOR/98-42",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-143",
      "title": "Use of Patented Products for International Humanitarian Purposes Regulations",
      "citation": "SOR/2005-143",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-124",
      "title": "Valley Gospel Mission Remission Order",
      "citation": "SI/2006-124",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-792",
      "title": "Valuation for Duty Regulations",
      "citation": "SOR/86-792",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-91-30",
      "title": "Value of Imported Goods (GST/HST) Regulations",
      "citation": "SOR/91-30",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-298",
      "title": "Vancouver 2010 Aviation Security Regulations",
      "citation": "SOR/2009-298",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-902",
      "title": "Vancouver International Airport Zoning Regulations",
      "citation": "SOR/80-902",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2010-10",
      "title": "Vancouver Organizing Committee for the 2010 Olympic and Paralympic Winter Games (GST/HST) Remission Order",
      "citation": "SI/2010-10",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-488",
      "title": "Order Varying Certain National Transportation Agency Orders Respecting Railway Companies",
      "citation": "SOR/89-488",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-590",
      "title": "Order Varying CTC Abandonment Orders Respecting the Avonlea Subdivision Between Parry and Avonlea",
      "citation": "SOR/88-590",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-98-45",
      "title": "Verification of Origin (Non-Free Trade Partners), Tariff Classification and Value for Duty of Imported Goods Regulations",
      "citation": "SOR/98-45",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-95",
      "title": "Versatile Pacific Shipyards Inc. Regulations",
      "citation": "SOR/87-95",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-31",
      "title": "Vessel Certificates Regulations",
      "citation": "SOR/2007-31",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-125",
      "title": "Vessel Clearance Regulations",
      "citation": "SOR/2007-125",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-127",
      "title": "Vessel Detention Orders Review Regulations",
      "citation": "SOR/2007-127",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-304",
      "title": "Vessel Duties Reduction or Removal Regulations",
      "citation": "SOR/90-304",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-120",
      "title": "Vessel Operation Restriction Regulations",
      "citation": "SOR/2008-120",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-69",
      "title": "Vessel Pollution and Dangerous Chemicals Regulations",
      "citation": "SOR/2012-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-126",
      "title": "Vessel Registration and Tonnage Regulations",
      "citation": "SOR/2007-126",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-172",
      "title": "Vessels Registry Fees Tariff",
      "citation": "SOR/2002-172",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-98",
      "title": "Vessel Traffic Services Zones Regulations",
      "citation": "SOR/89-98",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-276",
      "title": "Vested Assets (Foreign Companies) Regulations",
      "citation": "SOR/92-276",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1582",
      "title": "Vetcraft Shops Regulations",
      "citation": "CRC, c 1582",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1602",
      "title": "Veterans Allowance Regulations",
      "citation": "CRC, c 1602",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-200",
      "title": "Veterans Burial Regulations, 2005",
      "citation": "SOR/2005-200",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1584",
      "title": "Veterans Estates Regulations",
      "citation": "CRC, c 1584",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-90-594",
      "title": "Veterans Health Care Regulations",
      "citation": "SOR/90-594",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1587",
      "title": "Veterans Insurance Regulations",
      "citation": "CRC, c 1587",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1594",
      "title": "Veterans\u0027 Land Regulations",
      "citation": "CRC, c 1594",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-67",
      "title": "Veterans Review and Appeal Board Regulations",
      "citation": "SOR/96-67",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1585",
      "title": "Veterans Treatment Regulations",
      "citation": "CRC, c 1585",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-143",
      "title": "Veterinary Drug Evaluation Fees Regulations",
      "citation": "SOR/96-143",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-364",
      "title": "VHF Radiotelephone Practices and Procedures Regulations",
      "citation": "SOR/81-364",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-817",
      "title": "By-laws Nos. 6 and 8 of VIA Rail Canada Inc.",
      "citation": "SOR/79-817",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-89-366",
      "title": "Victim Fine Surcharge Regulations",
      "citation": "SOR/89-366",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-121",
      "title": "Victoria International Airport Zoning Regulations",
      "citation": "CRC, c 121",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1384",
      "title": "Victoria Jubilee Bridge Traffic By-law",
      "citation": "CRC, c 1384",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-568",
      "title": "Villeneuve Airport Zoning Regulations",
      "citation": "SOR/81-568",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-631",
      "title": "Vinyl Chloride Release Regulations, 1992",
      "citation": "SOR/92-631",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2006-298",
      "title": "Virtual Elimination List",
      "citation": "SOR/2006-298",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-122",
      "title": "Visiting Forces and Visiting Forces Personnel Alcoholic Beverages Remission Order",
      "citation": "SI/85-122",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-85-122",
      "title": "Visiting Forces and Visiting Forces Personnel Alcoholic Beverages Remission Order",
      "citation": "SI/85-122",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1597",
      "title": "Visiting Forces Attachment and Serving Together Regulations",
      "citation": "CRC, c 1597",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-210",
      "title": "Visiting Forces (Part IX of the Excise Tax Act) Remission Order",
      "citation": "SI/92-210",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1598",
      "title": "Visiting Forces Regulations",
      "citation": "CRC, c 1598",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-109",
      "title": "Visitor Visa Fee (World Youth Day 2002) Remission Order",
      "citation": "SI/2001-109",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-264",
      "title": "Volatile Organic Compound (VOC) Concentration Limits for Architectural Coatings Regulations",
      "citation": "SOR/2009-264",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-197",
      "title": "Volatile Organic Compound (VOC) Concentration Limits for Automotive Refinishing Products Regulations",
      "citation": "SOR/2009-197",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2011-203",
      "title": "Voyage Data Recorder Regulations",
      "citation": "SOR/2011-203",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-400",
      "title": "Voyageur Colonial Limited Acquisition Exemption Order",
      "citation": "SOR/88-400",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-100",
      "title": "Vuntut Gwitchin First Nation (GST) Remission Order",
      "citation": "SI/2000-100",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-7",
      "title": "Wabush Airport Zoning Regulations",
      "citation": "SOR/87-7",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2008-222",
      "title": "Wage Earner Protection Program Regulations",
      "citation": "SOR/2008-222",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-67",
      "title": "Wapusk National Park of Canada Park Use Regulations",
      "citation": "SOR/2010-67",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-54-578",
      "title": "War Claims Regulations",
      "citation": "SOR/54-578",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1601",
      "title": "War Service Grants Regulations",
      "citation": "CRC, c 1601",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2012-139",
      "title": "Wastewater Systems Effluent Regulations",
      "citation": "SOR/2012-139",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2014-279",
      "title": "Order Declaring that the Wastewater Systems Effluent Regulations Do Not Apply in Yukon",
      "citation": "SOR/2014-279",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-76",
      "title": "Waterloo-Guelph Airport Zoning Regulations",
      "citation": "SOR/93-76",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-142",
      "title": "Proclamation Exempting the Waters of Lake Pignac and Lake B from the Operation of Section 22 of the Act",
      "citation": "SOR/2013-142",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2009-202",
      "title": "Proclamation Exempting the Waters of Sandy Pond from the Operation of Section 22 of the Navigable Waters Protection Act",
      "citation": "SOR/2009-202",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-349",
      "title": "Water Supply and Sewage Disposal Services Charges Order (Peace River-Liard Regional District)",
      "citation": "SOR/84-349",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-970",
      "title": "Watson Lake Airport Zoning Regulations",
      "citation": "SOR/84-970",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-290",
      "title": "Watson Lake Hospital Divestiture Regulations",
      "citation": "SOR/2010-290",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1604",
      "title": "Weather Modification Information Regulations",
      "citation": "CRC, c 1604",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-42",
      "title": "Proclamation declaring the Wednesday of Canadian Environment Week in June of each Year to be Clean Air Day Canada",
      "citation": "SI/99-42",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-86-836",
      "title": "Weed Seeds Order, 1986",
      "citation": "SOR/86-836",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-220",
      "title": "Weed Seeds Order, 2005",
      "citation": "SOR/2005-220",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1605",
      "title": "Weights and Measures Regulations",
      "citation": "CRC, c 1605",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2007-232",
      "title": "Westbank First Nation Land Registry Regulations",
      "citation": "SOR/2007-232",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-992",
      "title": "West Coast Shipping Employees Hours of Work Regulations",
      "citation": "CRC, c 992",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1607",
      "title": "Western Grain Stabilization Regulations",
      "citation": "CRC, c 1607",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-314",
      "title": "Western Grain Transition Payments Regulations",
      "citation": "SOR/95-314",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-92-21",
      "title": "Weyburn Airport Zoning Regulations",
      "citation": "SOR/92-21",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-93-327",
      "title": "Whale Cove Airport Zoning Regulations",
      "citation": "SOR/93-327",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-108",
      "title": "Wheat Stabilization 1977 Regulations",
      "citation": "SOR/79-108",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-122",
      "title": "Whitehorse Airport Zoning Regulations",
      "citation": "CRC, c 122",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-76",
      "title": "White Pea Bean Stabilization 1978 Regulations",
      "citation": "SOR/80-76",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-609",
      "title": "White Pea Bean Stabilization Regulations",
      "citation": "SOR/78-609",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-26",
      "title": "White Pea Bean Stabilization Regulations, 1982-83",
      "citation": "SOR/84-26",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-286",
      "title": "White Pea Bean Stabilization Regulations, 1985-86",
      "citation": "SOR/87-286",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-450",
      "title": "Wiarton Airport Zoning Regulations",
      "citation": "SOR/88-450",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-96-263",
      "title": "Wild Animal and Plant Trade Regulations",
      "citation": "SOR/96-263",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1609",
      "title": "Wildlife Area Regulations",
      "citation": "CRC, c 1609",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-123",
      "title": "Windsor Airport Zoning Regulations",
      "citation": "CRC, c 123",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-708",
      "title": "Winnipeg International Airport Zoning Regulations",
      "citation": "SOR/81-708",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-136",
      "title": "Winter Pears Stabilization Regulations",
      "citation": "SOR/78-136",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-20",
      "title": "Winter Wheat Stabilization Regulations, 1982-83",
      "citation": "SOR/84-20",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1195",
      "title": "Wire Crossings and Proximities Regulations",
      "citation": "CRC, c 1195",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-132",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands in Nunavut (Northern Bathurst Island National Park, Nunavut)",
      "citation": "SI/2004-132",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-153",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands in Nunavut (the Eeyou Marine Region, Nunavut)",
      "citation": "SI/2004-153",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-60",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands in Nunavut (The Nunavik Marine Region, Nunavut)",
      "citation": "SI/2006-60",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-149",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands in Nunavut (The Nunavik Marine Region, Nunavut)",
      "citation": "SI/2003-149",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2007-103",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands in the Northwest Territories (Akaitcho Dene First Nations, N.W.T.)",
      "citation": "SI/2007-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2007-68",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands in the Northwest Territories (Deh Cho First Nations, N.W.T.)",
      "citation": "SI/2007-68",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-148",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands in the Northwest Territories (Deh Cho First Nations, N.W.T.)",
      "citation": "SI/2003-148",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2000-77",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands in the Northwest Territories (Dogrib Settlement Agreement, North Slave Region, N.W.T.)",
      "citation": "SI/2000-77",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2007-101",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands in the Northwest Territories (East Arm of Great Slave Lake, N.W.T.)",
      "citation": "SI/2007-101",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-154",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands in the Northwest Territories (Ezodziti, N.W.T.)",
      "citation": "SI/2002-154",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-55",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands in the Northwest Territories (Giant Mine)",
      "citation": "SI/2005-55",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2005-113",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands in the Northwest Territories (Sahoyúé - ?ehdacho (Grizzly Bear Mountain and Scented Grass Hills), National Historic Site, N.W.T.)",
      "citation": "SI/2005-113",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2001-26",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands in the Northwest Territories (Sahyoue/Edacho (Grizzly Bear Mountain and Scented Grass Hills), N.W.T.)",
      "citation": "SI/2001-26",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2006-98",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands in the Northwest Territories (Salt River First Nation, N.W.T.)",
      "citation": "SI/2006-98",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2002-123",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands in the Northwest Territories (Salt River First Nation, N.W.T.)",
      "citation": "SI/2002-123",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2007-102",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands in the Northwest Territories (Ts\u0027ude niline Tu\u0027eyeta (Ramparts River and Wetlands))",
      "citation": "SI/2007-102",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-97-93",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands in the Yukon Territory (Dawson First Nation, Y.T.)",
      "citation": "SI/97-93",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-38",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands (Tuktut Nogait National Park, Northwest Territories and Nunavut)",
      "citation": "SI/2003-38",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2004-122",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Lands (Ukkusiksalik National Park, Nunavut)",
      "citation": "SI/2004-122",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-36",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Subsurface Lands in the Northwest Territories",
      "citation": "SI/2003-36",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-78",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in Nunavut (Eeyou Marine Region) Order",
      "citation": "SI/2011-78",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2008-104",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands In Nunavut (Eeyou Marine Region) Order",
      "citation": "SI/2008-104",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-62",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in Nunavut (Kivalliq area) Order",
      "citation": "SI/2013-62",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2009-104",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in Nunavut (Northern Bathurst Island National Park) Order",
      "citation": "SI/2009-104",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-109",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in Nunavut (Qausuittuq National Park) Order",
      "citation": "SI/2014-109",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2009-103",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in Nunavut (Ukkusiksalik National Park) Order",
      "citation": "SI/2009-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-39",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Central and Eastern Portions of the South Slave Region) Order",
      "citation": "SI/2014-39",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2008-103",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Dehcho First Nations) Order",
      "citation": "SI/2008-103",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2010-83",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Dehcho First Nations) Order",
      "citation": "SI/2010-83",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-92",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Dehcho Region) Order",
      "citation": "SI/2011-92",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-126",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Dehcho Region) Order",
      "citation": "SI/2013-126",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-63",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Eastern Portion of the South Slave Region) Order",
      "citation": "SI/2013-63",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2008-102",
      "title": "Withdrawal fromDisposal of Certain Tracts of Territorial Lands in the Northwest Territories (Edéhzhíe (Horn Plateau)) Order",
      "citation": "SI/2008-102",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-111",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Edéhzhíe (Horn Plateau)) Order",
      "citation": "SI/2011-111",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2010-84",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Edéhzhíe (Horn Plateau)) Order",
      "citation": "SI/2010-84",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-106",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Kwets\u0027ootl\u0027àà (North Arm of Great Slave Lake)) Order",
      "citation": "SI/2013-106",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2008-27",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Nááts\u0027ihch\u0027oh National Park Reserve)",
      "citation": "SI/2008-27",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-25",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Nááts\u0027ihch\u0027oh National Park Reserve) Order",
      "citation": "SI/2012-25",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-37",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Nááts\u0027ihch\u0027oh National Park Reserve) Order",
      "citation": "SI/2014-37",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2008-101",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Nahanni National Park Reserve of Canada) Order",
      "citation": "SI/2008-101",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2008-20",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Nành\u0027 Geenjit Gwitr\u0027it Tigwaa\u0027in/Working for the Land: Gwich\u0027in Land use Plan), N.W.T.",
      "citation": "SI/2008-20",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2003-107",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Nành\u0027Geenjit Gwitr\u0027it Tigwaa\u0027in/Working for the Land: Gwich\u0027in Land use Plan), N.W.T.",
      "citation": "SI/2003-107",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2009-73",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Salt River First Nation) Order",
      "citation": "SI/2009-73",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2009-94",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Saoyú–Æehdacho (Grizzly Bear Mountain and Scented Grass Hills) National Historic Site) Order",
      "citation": "SI/2009-94",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-35",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (South Slave and North Slave Regions) Order",
      "citation": "SI/2014-35",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-24",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (South Slave / North Slave Regions) Order",
      "citation": "SI/2012-24",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-6",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (South Slave Region) Order",
      "citation": "SI/2013-6",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2012-23",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Thaidene Nene (East Arm of Great Slave Lake) National Park Reserve) Order",
      "citation": "SI/2012-23",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-36",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Thaidene Nene (East Arm of Great Slave Lake) National Park Reserve) Order",
      "citation": "SI/2014-36",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-93",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Ts\u0027ude niline Tu\u0027eyeta (Ramparts River and Wetlands)) Order",
      "citation": "SI/2011-93",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-125",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Ts\u0027ude niline Tu\u0027eyeta (Ramparts River and Wetlands)) Order",
      "citation": "SI/2013-125",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2013-53",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Tuktut Nogait National Park of Canada) Order",
      "citation": "SI/2013-53",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-58",
      "title": "Withdrawal from Disposal of Certain Tracts of Territorial Lands in the Northwest Territories (Tuktut Nogait National Park) Order",
      "citation": "SI/2011-58",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2008-26",
      "title": "Order Respecting the Withdrawal from Disposal of Certain Tracts of Territorial Lands (Tuktut Nogait National Park, in the Northwest Territories and Nunavut)",
      "citation": "SI/2008-26",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2011-14",
      "title": "Withdrawal from Disposal of the Subsurface Rights in Certain Tracts of Territorial Lands in the Northwest Territories (Nành\u0027 Geenjit Gwitr\u0027it Tigwaa\u0027in/Working for the Land: Gwich\u0027in Land use Plan) Order",
      "citation": "SI/2011-14",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-99-53",
      "title": "Withdrawal from Disposal Order Aulavik (Banks Island) National Park, Northwest Territories",
      "citation": "SI/99-53",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-150",
      "title": "Withdrawal from Disposal Order (Contwoyto Lake, N.W.T.)",
      "citation": "SI/93-150",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-98-3",
      "title": "Withdrawal from Disposal Order (East Arm of Great Slave Lake National Park — Great Slave Lake, N.W.T.)",
      "citation": "SI/98-3",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-95-23",
      "title": "Withdrawal from Disposal Order (Kelly Lake Area, N.W.T.)",
      "citation": "SI/95-23",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-96-87",
      "title": "Withdrawal from Disposal Order (Wager Bay National Park, N.W.T.)",
      "citation": "SI/96-87",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2010-18",
      "title": "Withdrawal from Disposal, Setting Apart and Appropriation of Certain Tracts of Territorial Lands in the Northwest Territories (Reindeer Grazing Reserve) Order",
      "citation": "SI/2010-18",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2014-38",
      "title": "Withdrawal from Disposal, Setting Apart and Appropriation of Certain Tracts of Territorial Lands in the Northwest Territories (Reindeer Grazing Reserve) Order",
      "citation": "SI/2014-38",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1530",
      "title": "Withdrawal of Certain Lands (Baffin Island N.W.T.) from Disposal Order",
      "citation": "CRC, c 1530",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1533",
      "title": "Withdrawal of Certain Lands (Dubawnt Lake N.W.T.) from Disposal Order",
      "citation": "CRC, c 1533",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-111",
      "title": "Withdrawal of Certain Lands (Fry Inlet on Contwoyto Lake, N.W.T.) from Disposal Order",
      "citation": "SI/92-111",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1535",
      "title": "Withdrawal of Certain Lands (Great Slave Lake N.W.T.) from Disposal Order",
      "citation": "CRC, c 1535",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2007-69",
      "title": "Withdrawal of Certain Lands in the Northwest Territories (Edéhzhie (Horn Plateau), N.W.T.) from Disposal Order",
      "citation": "SI/2007-69",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-2007-80",
      "title": "Withdrawal of Certain Lands in the Northwest Territories (Nahanni National Park Reserve of Canada) from Disposal Order",
      "citation": "SI/2007-80",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1537",
      "title": "Withdrawal of Certain Lands (Lockhart River N.W.T.) from Disposal Order",
      "citation": "CRC, c 1537",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-74",
      "title": "Withdrawal of Certain Lands (Mackenzie Delta Region, N.W.T.) from Disposal Order",
      "citation": "SI/92-74",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-91-155",
      "title": "Withdrawal of Certain Lands (Mayo, Y.T.) from Disposal Order",
      "citation": "SI/91-155",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-139",
      "title": "Withdrawal of Certain Lands (North and South Baffin, Kitikmeot East and West and Keewatin, N.W.T.) from Disposal Order",
      "citation": "SI/92-139",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-41",
      "title": "Withdrawal of Certain Lands (North Baffin Island, N.W.T.) from Disposal Order",
      "citation": "SI/92-41",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-92-73",
      "title": "Withdrawal of Certain Lands (Peel River Basin, Y.T.) from Disposal Order",
      "citation": "SI/92-73",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1539",
      "title": "Withdrawal of Certain Lands (Sawmill Bay N.W.T.) from Disposal Order",
      "citation": "CRC, c 1539",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1540",
      "title": "Withdrawal of Certain Lands (South Nahanni River N.W.T.) from Disposal Order, 1971",
      "citation": "CRC, c 1540",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-1541",
      "title": "Withdrawal of Certain Lands (South Nahanni River N.W.T.) from Disposal Order, 1972",
      "citation": "CRC, c 1541",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2000-143",
      "title": "Withdrawal of Entities Regulations",
      "citation": "SOR/2000-143",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-78-830",
      "title": "Wood Buffalo National Park Game Regulations",
      "citation": "SOR/78-830",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-263",
      "title": "14th IAAF World Half Marathon Championships Remission Order",
      "citation": "SOR/2005-263",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2010-130",
      "title": "13th IAAF World Junior Championships in Athletics Remission Order, 2010",
      "citation": "SOR/2010-130",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-85-154",
      "title": "World Meteorological Organization Privileges and Immunities Order",
      "citation": "SOR/85-154",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2002-259",
      "title": "Woven Fabrics and Shells of Woven Fabrics Remission Order",
      "citation": "SOR/2002-259",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-95-507",
      "title": "Wrigley Airport Zoning Regulations",
      "citation": "SOR/95-507",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-125",
      "title": "Yarmouth Airport Zoning Regulations",
      "citation": "CRC, c 125",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-81-472",
      "title": "Yellowknife Airport Zoning Regulations",
      "citation": "SOR/81-472",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-84-23",
      "title": "Yellow Seed Onions Stabilization Regulations, 1982-83",
      "citation": "SOR/84-23",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-79-276",
      "title": "Yellow Seed Onion Stabilization 1977 Regulations",
      "citation": "SOR/79-276",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-80-167",
      "title": "Yellow Seed Onion Stabilization 1978 Regulations",
      "citation": "SOR/80-167",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-87-605",
      "title": "Yorkton Airport Zoning Regulations",
      "citation": "SOR/87-605",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2013-38",
      "title": "Yukon Borrowing Limits Regulations",
      "citation": "SOR/2013-38",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-2005-43",
      "title": "Yukon Sex Offender Information Registration Regulations",
      "citation": "SOR/2005-43",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "si-93-53",
      "title": "Yukon Territory Court of Appeal Criminal Appeal Rules, 1993",
      "citation": "SI/93-53",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "crc-c-854",
      "title": "Yukon Territory Fishery Regulations",
      "citation": "CRC, c 854",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-73-71",
      "title": "Yukon Territory Pension Regulations",
      "citation": "SOR/73-71",
      "type": "REGULATION"
    },
    {
      "databaseId": "car",
      "legislationId": "sor-88-427",
      "title": "Yukon Territory Supreme Court Rules for Pre-hearing Conferences in Criminal Matters",
      "citation": "SOR/88-427",
      "type": "REGULATION"
    }
  ]
}`
	var v struct {
		Legis []canlii.Legislation `json:"legislations"`
	}
	err := json.NewDecoder(strings.NewReader(resp)).Decode(&v)
	if err != nil {
		t.Fatal(err)
	}
	wantLegis := v.Legis
	api := "derp"
	path := "/v1/legislationBrowse/en/car"

	client, close := testClient(t, path, resp, api)
	defer close()

	gotLegis, _, err := client.LegislationBrowse.ListLegislations("car")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantLegis, gotLegis) {
		t.Logf("want=%+v", wantLegis)
		t.Logf(" got=%+v", gotLegis)
		t.Error("mismatch yo")
	}
}

func TestLegislationBrowseCaseMetadata(t *testing.T) {
	resp := `{
  "legislationId": "si-2005-87",
  "url": "http://canlii.ca/t/7wwq",
  "title": "Proclamation Giving Notice of the Coming into Force on December 31, 2004 of the Convention Between Canada and Romania for the Avoidance of Double Taxation",
  "citation": "SI/2005-87",
  "type": "REGULATION",
  "language": "en",
  "dateScheme": "ENTRY_INTO_FORCE",
  "startDate": "2006-03-22",
  "endDate": "",
  "repealed": "NO",
  "content": [
    {
      "partId": "1",
      "partName": "Main"
    }
  ]
}`
	var wantLegis canlii.Legislation
  err := json.NewDecoder(strings.NewReader(resp)).Decode(&wantLegis)
	if err != nil {
		t.Fatal(err)
	}
	api := "derp"
	path := "/v1/legislationBrowse/en/car/si-2005-87"

	client, close := testClient(t, path, resp, api)
	defer close()

	gotLegis, _, err := client.LegislationBrowse.LegislationMetadata("car", "si-2005-87")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(wantLegis, gotLegis) {
		t.Logf("want=%+v", wantLegis)
		t.Logf(" got=%+v", gotLegis)
		t.Error("mismatch yo")
	}
}
