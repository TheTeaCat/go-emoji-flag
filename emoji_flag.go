package emojiflag

import (
	"strings"
)

var countryMap = map[string]string{
	"AFG": "AF",
	"ALA": "AX",
	"ALB": "AL",
	"DZA": "DZ",
	"ASM": "AS",
	"AND": "AD",
	"AGO": "AO",
	"AIA": "AI",
	"ATA": "AQ",
	"ATG": "AG",
	"ARG": "AR",
	"ARM": "AM",
	"ABW": "AW",
	"AUS": "AU",
	"AUT": "AT",
	"AZE": "AZ",
	"BHS": "BS",
	"BHR": "BH",
	"BGD": "BD",
	"BRB": "BB",
	"BLR": "BY",
	"BEL": "BE",
	"BLZ": "BZ",
	"BEN": "BJ",
	"BMU": "BM",
	"BTN": "BT",
	"BOL": "BO",
	"BIH": "BA",
	"BWA": "BW",
	"BVT": "BV",
	"BRA": "BR",
	"VGB": "VG",
	"IOT": "IO",
	"BRN": "BN",
	"BGR": "BG",
	"BFA": "BF",
	"BDI": "BI",
	"KHM": "KH",
	"CMR": "CM",
	"CAN": "CA",
	"CPV": "CV",
	"CYM": "KY",
	"CAF": "CF",
	"TCD": "TD",
	"CHL": "CL",
	"CHN": "CN",
	"HKG": "HK",
	"MAC": "MO",
	"CXR": "CX",
	"CCK": "CC",
	"COL": "CO",
	"COM": "KM",
	"COG": "CG",
	"COD": "CD",
	"COK": "CK",
	"CRI": "CR",
	"CIV": "CI",
	"HRV": "HR",
	"CUB": "CU",
	"CYP": "CY",
	"CZE": "CZ",
	"DNK": "DK",
	"DJI": "DJ",
	"DMA": "DM",
	"DOM": "DO",
	"ECU": "EC",
	"EGY": "EG",
	"SLV": "SV",
	"GNQ": "GQ",
	"ERI": "ER",
	"EST": "EE",
	"ETH": "ET",
	"FLK": "FK",
	"FRO": "FO",
	"FJI": "FJ",
	"FIN": "FI",
	"FRA": "FR",
	"GUF": "GF",
	"PYF": "PF",
	"ATF": "TF",
	"GAB": "GA",
	"GMB": "GM",
	"GEO": "GE",
	"DEU": "DE",
	"GHA": "GH",
	"GIB": "GI",
	"GRC": "GR",
	"GRL": "GL",
	"GRD": "GD",
	"GLP": "GP",
	"GUM": "GU",
	"GTM": "GT",
	"GGY": "GG",
	"GIN": "GN",
	"GNB": "GW",
	"GUY": "GY",
	"HTI": "HT",
	"HMD": "HM",
	"VAT": "VA",
	"HND": "HN",
	"HUN": "HU",
	"ISL": "IS",
	"IND": "IN",
	"IDN": "ID",
	"IRN": "IR",
	"IRQ": "IQ",
	"IRL": "IE",
	"IMN": "IM",
	"ISR": "IL",
	"ITA": "IT",
	"JAM": "JM",
	"JPN": "JP",
	"JEY": "JE",
	"JOR": "JO",
	"KAZ": "KZ",
	"KEN": "KE",
	"KIR": "KI",
	"PRK": "KP",
	"KOR": "KR",
	"KWT": "KW",
	"KGZ": "KG",
	"LAO": "LA",
	"LVA": "LV",
	"LBN": "LB",
	"LSO": "LS",
	"LBR": "LR",
	"LBY": "LY",
	"LIE": "LI",
	"LTU": "LT",
	"LUX": "LU",
	"MKD": "MK",
	"MDG": "MG",
	"MWI": "MW",
	"MYS": "MY",
	"MDV": "MV",
	"MLI": "ML",
	"MLT": "MT",
	"MHL": "MH",
	"MTQ": "MQ",
	"MRT": "MR",
	"MUS": "MU",
	"MYT": "YT",
	"MEX": "MX",
	"FSM": "FM",
	"MDA": "MD",
	"MCO": "MC",
	"MNG": "MN",
	"MNE": "ME",
	"MSR": "MS",
	"MAR": "MA",
	"MOZ": "MZ",
	"MMR": "MM",
	"NAM": "NA",
	"NRU": "NR",
	"NPL": "NP",
	"NLD": "NL",
	"ANT": "AN",
	"NCL": "NC",
	"NZL": "NZ",
	"NIC": "NI",
	"NER": "NE",
	"NGA": "NG",
	"NIU": "NU",
	"NFK": "NF",
	"MNP": "MP",
	"NOR": "NO",
	"OMN": "OM",
	"PAK": "PK",
	"PLW": "PW",
	"PSE": "PS",
	"PAN": "PA",
	"PNG": "PG",
	"PRY": "PY",
	"PER": "PE",
	"PHL": "PH",
	"PCN": "PN",
	"POL": "PL",
	"PRT": "PT",
	"PRI": "PR",
	"QAT": "QA",
	"REU": "RE",
	"ROU": "RO",
	"RUS": "RU",
	"RWA": "RW",
	"BLM": "BL",
	"SHN": "SH",
	"KNA": "KN",
	"LCA": "LC",
	"MAF": "MF",
	"SPM": "PM",
	"VCT": "VC",
	"WSM": "WS",
	"SMR": "SM",
	"STP": "ST",
	"SAU": "SA",
	"SEN": "SN",
	"SRB": "RS",
	"SYC": "SC",
	"SLE": "SL",
	"SGP": "SG",
	"SVK": "SK",
	"SVN": "SI",
	"SLB": "SB",
	"SOM": "SO",
	"ZAF": "ZA",
	"SGS": "GS",
	"SSD": "SS",
	"ESP": "ES",
	"LKA": "LK",
	"SDN": "SD",
	"SUR": "SR",
	"SJM": "SJ",
	"SWZ": "SZ",
	"SWE": "SE",
	"CHE": "CH",
	"SYR": "SY",
	"TWN": "TW",
	"TJK": "TJ",
	"TZA": "TZ",
	"THA": "TH",
	"TLS": "TL",
	"TGO": "TG",
	"TKL": "TK",
	"TON": "TO",
	"TTO": "TT",
	"TUN": "TN",
	"TUR": "TR",
	"TKM": "TM",
	"TCA": "TC",
	"TUV": "TV",
	"UGA": "UG",
	"UKR": "UA",
	"ARE": "AE",
	"GBR": "GB",
	"USA": "US",
	"UMI": "UM",
	"URY": "UY",
	"UZB": "UZ",
	"VUT": "VU",
	"VEN": "VE",
	"VNM": "VN",
	"VIR": "VI",
	"WLF": "WF",
	"ESH": "EH",
	"YEM": "YE",
	"ZMB": "ZM",
	"ZWE": "ZW",
	"XKX": "XK",
}

// GetFlag returns a best attempt at matching a country code flag
func GetFlag(countryCode string) string {
	code := strings.ToUpper(countryCode)
	if len(code) == 3 {
		if cc, ok := countryMap[code]; ok {
			return string(0x1F1E6+rune(cc[0])-'A') + string(0x1F1E6+rune(cc[1])-'A')
		}
	}

	if len(code) == 2 {
		for _, v := range countryMap {
			if code == v {
				return string(0x1F1E6+rune(code[0])-'A') + string(0x1F1E6+rune(code[1])-'A')
			}
		}
	}

	return ""
}
