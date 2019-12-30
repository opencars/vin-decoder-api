package govin

type Country struct {
	From string
	To   string
	Name string
}

var countries = []Country{
	{
		From: "AA",
		To:   "AH",
		Name: "South Africa",
	},
	{
		From: "AJ",
		To:   "AN",
		Name: "Ivory Coast",
	},
	{
		From: "BA",
		To:   "BE",
		Name: "Angola",
	},
	{
		From: "BF",
		To:   "BK",
		Name: "Kenya",
	},
	{
		From: "BL",
		To:   "BR",
		Name: "Tanzania",
	},
	{
		From: "CA",
		To:   "CE",
		Name: "Benin",
	},
	{
		From: "CF",
		To:   "CK",
		Name: "Madagascar",
	},
	{
		From: "CL",
		To:   "CR",
		Name: "Tunisia",
	},
	{
		From: "DA",
		To:   "DE",
		Name: "Egypt",
	},
	{
		From: "DF",
		To:   "DK",
		Name: "Morocco",
	},
	{
		From: "DL",
		To:   "DR",
		Name: "Zambia",
	},
	{
		From: "EA",
		To:   "EE",
		Name: "Ethiopia",
	},
	{
		From: "EF",
		To:   "EK",
		Name: "Mozambique",
	},
	{
		From: "FA",
		To:   "FE",
		Name: "Ghana",
	},
	{
		From: "FF",
		To:   "FK",
		Name: "Nigeria",
	},
	{
		From: "JA",
		To:   "J0",
		Name: "Japan",
	},
	{
		From: "KA",
		To:   "KE",
		Name: "Sri Lanka",
	},
	{
		From: "KF",
		To:   "KK",
		Name: "Israel",
	},
	{
		From: "KL",
		To:   "KR",
		Name: "Korea (South)",
	},
	{
		From: "KS",
		To:   "K0",
		Name: "Kazakhstan",
	},
	{
		From: "LA",
		To:   "L0",
		Name: "China",
	},
	{
		From: "MA",
		To:   "ME",
		Name: "India",
	},
	{
		From: "MF",
		To:   "MK",
		Name: "Indonesia",
	},
	{
		From: "ML",
		To:   "MR",
		Name: "Thailand",
	},
	{
		From: "NA",
		To:   "NE",
		Name: "Iran",
	},
	{
		From: "NF",
		To:   "NK",
		Name: "Pakistan",
	},
	{
		From: "NL",
		To:   "NR",
		Name: "Turkey",
	},
	{
		From: "PA",
		To:   "PE",
		Name: "Philippines",
	},
	{
		From: "PF",
		To:   "PK",
		Name: "Singapore",
	},
	{
		From: "PL",
		To:   "PR",
		Name: "Malaysia",
	},
	{
		From: "RA",
		To:   "RE",
		Name: "United Arab Emirates",
	},
	{
		From: "RF",
		To:   "RK",
		Name: "Taiwan",
	},
	{
		From: "RL",
		To:   "RR",
		Name: "Vietnam",
	},
	{
		From: "RS",
		To:   "R0",
		Name: "Saudi Arabia",
	},
	{
		From: "SA",
		To:   "SM",
		Name: "United Kingdom",
	},
	{
		From: "SN",
		To:   "ST",
		Name: "Germany",
	},
	{
		From: "SU",
		To:   "SZ",
		Name: "Poland",
	},
	{
		From: "S1",
		To:   "S4",
		Name: "Latvia",
	},
	{
		From: "TA",
		To:   "TH",
		Name: "Switzerland",
	},
	{
		From: "TJ",
		To:   "TP",
		Name: "Czech Republic",
	},
	{
		From: "TR",
		To:   "TV",
		Name: "Hungary",
	},
	{
		From: "TW",
		To:   "T1",
		Name: "Portugal",
	},
	{
		From: "UH",
		To:   "UM",
		Name: "Denmark",
	},
	{
		From: "UN",
		To:   "UT",
		Name: "Ireland",
	},
	{
		From: "UU",
		To:   "UZ",
		Name: "Romania",
	},
	{
		From: "U5",
		To:   "U7",
		Name: "Slovakia",
	},
	{
		From: "VA",
		To:   "VE",
		Name: "Austria",
	},
	{
		From: "VF",
		To:   "VR",
		Name: "France",
	},
	{
		From: "VS",
		To:   "VW",
		Name: "Spain",
	},
	{
		From: "VX",
		To:   "V2",
		Name: "Serbia",
	},
	{
		From: "V3",
		To:   "V5",
		Name: "Croatia",
	},
	{
		From: "V6",
		To:   "V0",
		Name: "Estonia",
	},
	{
		From: "WA",
		To:   "W0",
		Name: "Germany",
	},
	{
		From: "XA",
		To:   "XE",
		Name: "Bulgaria",
	},
	{
		From: "XF",
		To:   "XK",
		Name: "Greece",
	},
	{
		From: "XL",
		To:   "XR",
		Name: "Netherlands",
	},
	{
		From: "XS",
		To:   "XW",
		Name: "Russia",
	},
	{
		From: "XX",
		To:   "X2",
		Name: "Luxembourg",
	},
	{
		From: "X3",
		To:   "X0",
		Name: "Russia",
	},
	{
		From: "YA",
		To:   "YE",
		Name: "Belgium",
	},
	{
		From: "YF",
		To:   "YK",
		Name: "Finland",
	},
	{
		From: "YL",
		To:   "YR",
		Name: "Malta",
	},
	{
		From: "YS",
		To:   "YW",
		Name: "Sweden",
	},
	{
		From: "YX",
		To:   "Y2",
		Name: "Norway",
	},
	{
		From: "Y3",
		To:   "Y5",
		Name: "Belarus",
	},
	{
		From: "Y6",
		To:   "Y0",
		Name: "Ukraine",
	},
	{
		From: "ZA",
		To:   "ZR",
		Name: "Italy",
	},
	{
		From: "ZX",
		To:   "Z2",
		Name: "Slovenia",
	},
	{
		From: "Z3",
		To:   "Z5",
		Name: "Lithuania",
	},
	{
		From: "1A",
		To:   "10",
		Name: "United States",
	},
	{
		From: "2A",
		To:   "20",
		Name: "Canada",
	},
	{
		From: "3A",
		To:   "37",
		Name: "Mexico",
	},
	{
		From: "38",
		To:   "30",
		Name: "Cayman Islands",
	},
	{
		From: "4A",
		To:   "40",
		Name: "United States",
	},
	{
		From: "5A",
		To:   "50",
		Name: "United States",
	},
	{
		From: "6A",
		To:   "6W",
		Name: "Australia",
	},
	{
		From: "7A",
		To:   "7E",
		Name: "New Zealand",
	},
	{
		From: "8A",
		To:   "8E",
		Name: "Argentina",
	},
	{
		From: "8F",
		To:   "8K",
		Name: "Chile",
	},
	{
		From: "8L",
		To:   "8R",
		Name: "Ecuador",
	},
	{
		From: "8S",
		To:   "8W",
		Name: "Peru",
	},
	{
		From: "8X",
		To:   "82",
		Name: "Venezuela",
	},
	{
		From: "9A",
		To:   "9E",
		Name: "Brazil",
	},
	{
		From: "9F",
		To:   "9K",
		Name: "Colombia",
	},
	{
		From: "9L",
		To:   "9R",
		Name: "Paraguay",
	},
	{
		From: "9S",
		To:   "9W",
		Name: "Uruguay",
	},
	{
		From: "9X",
		To:   "92",
		Name: "Trinidad & Tobago",
	},
	{
		From: "93",
		To:   "99",
		Name: "Brazil",
	},
}
