package entity

type Fact string

type Languages struct {
	Lang []Lang `json:"lang"`
}

type Lang struct {
	LocaleCode  string `json:"locale_code"`
	IsoCode     string `json:"iso_code"`
	FullCode    string `json:"full_code"`
	LocalName   string `json:"local_name"`
	EnglishName string `json:"english_name"`
	FullName    string `json:"full_name"`
	FactCount   int    `json:"fact_count"`
}
