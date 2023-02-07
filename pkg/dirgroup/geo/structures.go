package geo

import (
	"fmt"
	"net/url"

	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

type DirGroupGeo struct {
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

func (d *DirGroupGeo) DirGroupGeoURI() string {
	d.Account = url.PathEscape(d.Account)
	d.Name = url.PathEscape(d.Name)

	return fmt.Sprintf("accounts/%s/dirgroups/geo/%s", d.Account, d.Name)
}

func (d *DirGroupGeo) DirGroupGeoListURI() string {
	d.Account = url.PathEscape(d.Account)

	return fmt.Sprintf("accounts/%s/dirgroups/geo", d.Account)
}
