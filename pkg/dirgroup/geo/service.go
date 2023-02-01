package geo

import (
	"fmt"
	"net/http"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
)

const (
	serviceName = "DirGroupGeoIPGroup"
	// https://api.ultradns.com/accounts/{accountName}/dirgroups/geo/
)

type Service struct {
	c *client.Client
}

func New(cfn client.Config) (*Service, error) {
	c, err := client.NewClient(cfn)
	if err != nil {
		return nil, errors.ServiceConfigError(serviceName, err)
	}

	return &Service{c}, nil
}

func Get(c *client.Client) (*Service, error) {
	if c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	return &Service{c}, nil
}

func (s *Service) dirGroupURI(groupName string) string {
	return fmt.Sprintf("accounts/%s/dirgroups/geo/%s", "testing", groupName)
}

func (s *Service) Read(groupName string) (*http.Response, *DirGroupGeoIP, error) {
	target := client.Target(&DirGroupGeoIP{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, s.dirGroupURI(groupName), nil, target)
	if err != nil {
		return nil, nil, errors.ReadError(serviceName, groupName, err)
	}

	groupRes := target.Data.(*DirGroupGeoIP)

	return res, groupRes, nil
}
