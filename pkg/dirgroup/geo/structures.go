package geo

import "github.com/ultradns/ultradns-go-sdk/pkg/helper"

// Account Level Group is defined as:
// {
//	 "name": "accountGeoGroup",
//	 "description": "A sample group",
//	 "codes": ["Z6", "RU"]
// }

type DirGroupGeoIP struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Codes       []string `json:"codes"`
}

type Response struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Codes       []string `json:"codes,omitempty"`
}

type ResponseList struct {
	QueryInfo      *helper.QueryInfo  `json:"queryInfo,omitempty"`
	CursorInfo     *helper.CursorInfo `json:"cursorInfo,omitempty"`
	DirGroupGeoIPs []*Response        `json:"zones,omitempty"`
}
