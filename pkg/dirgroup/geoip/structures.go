package geoip

import "github.com/ultradns/ultradns-go-sdk/pkg/helper"

// Account Level Group is defined as:
// {
//	 "name": "accountGeoGroup",
//	 "description": "A sample group",
//	 "codes": ["Z6", "RU"]
// }

type DirGroupGeoIP struct {
	Account     string   `json:"account_name,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Codes       []string `json:"codes,omitempty"`
}

type Response struct {
	Account     string   `json:"account_name,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Codes       []string `json:"codes,omitempty"`
}

type ResponseList struct {
	QueryInfo      *helper.QueryInfo  `json:"queryInfo,omitempty"`
	CursorInfo     *helper.CursorInfo `json:"cursorInfo,omitempty"`
	DirGroupGeoIPs []*Response        `json:"dirgroupgeoips,omitempty"`
}
