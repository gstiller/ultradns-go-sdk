package geo

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

const (
	serviceName = "DirGroupGeoIPGroup"
	basePath    = "accounts/"
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

func (s *Service) CreateDirGroupGeoIP(dirGroupGeoIP *DirGroupGeoIP, cfn *client.Config) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	createPath := basePath + cfn.Username + "/dirgroups/geo/" + dirGroupGeoIP.Name
	res, err := s.c.Do(http.MethodPost, createPath, dirGroupGeoIP, target)

	if err != nil {
		geoIPGroupName := ""
		geoIPGroupName = dirGroupGeoIP.Name

		return nil, errors.CreateError(serviceName, geoIPGroupName, err)
	}

	return res, nil
}

func (s *Service) dirGroupURI(groupName string) string {
	return fmt.Sprintf("accounts/%s/dirgroups/geo/%s", "testing", groupName)
}

func (s *Service) ReadDirGroupGeoIP(dirGroupGeoIPName string, cfn *client.Config) (*http.Response, *Response, error) {
	target := client.Target(&Response{})
	dirGroupGeoIPName = url.PathEscape(dirGroupGeoIPName)

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	readPath := basePath + cfn.Username + "/dirgroups/geo/" + dirGroupGeoIPName
	res, err := s.c.Do(http.MethodGet, readPath, nil, target)
	if err != nil {
		return nil, nil, errors.ReadError(serviceName, dirGroupGeoIPName, err)
	}

	dirGroupGeoIPResponse := target.Data.(*Response)

	return res, dirGroupGeoIPResponse, nil
}

func (s *Service) UpdateDirGroupGeoIP(dirGroupGeoIPName string, dirGroupGeoIP *DirGroupGeoIP, cfn *client.Config) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	dirGroupGeoIPName = url.PathEscape(dirGroupGeoIPName)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	updatePath := basePath + cfn.Username + "/dirgroups/geo/" + dirGroupGeoIPName
	res, err := s.c.Do(http.MethodPut, updatePath, dirGroupGeoIP, target)

	if err != nil {
		return nil, errors.UpdateError(serviceName, dirGroupGeoIPName, err)
	}

	return res, nil
}

func (s *Service) PartialUpdateDirGroupGeoIP(dirGroupGeoIPName string, dirGroupGeoIP *DirGroupGeoIP, cfn *client.Config) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	dirGroupGeoIPName = url.PathEscape(dirGroupGeoIPName)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	partiallyUpdatePath := basePath + cfn.Username + "/dirgroups/geo/" + dirGroupGeoIPName
	res, err := s.c.Do(http.MethodPatch, partiallyUpdatePath, dirGroupGeoIP, target)

	if err != nil {
		return nil, errors.PartialUpdateError(serviceName, dirGroupGeoIPName, err)
	}

	return res, nil
}

func (s *Service) DeleteDirGroupGeoIP(dirGroupGeoIPName string, cfn *client.Config) (*http.Response, error) {
	target := client.Target(&client.SuccessResponse{})
	dirGroupGeoIPName = url.PathEscape(dirGroupGeoIPName)

	if s.c == nil {
		return nil, errors.ServiceError(serviceName)
	}

	deletePath := basePath + cfn.Username + "/dirgroups/geo/" + dirGroupGeoIPName
	res, err := s.c.Do(http.MethodDelete, deletePath, nil, target)

	if err != nil {
		return nil, errors.DeleteError(serviceName, dirGroupGeoIPName, err)
	}

	return res, nil
}

func (s *Service) ListDirGroupGeoIP(queryInfo *helper.QueryInfo, cfn client.Config) (*http.Response, *ResponseList, error) {
	target := client.Target(&ResponseList{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	listPath := basePath + cfn.Username + "/dirgroups/geo"
	res, err := s.c.Do(http.MethodGet, listPath+queryInfo.URI(), nil, target)

	if err != nil {
		return nil, nil, errors.ListError(serviceName, listPath+queryInfo.URI(), err)
	}

	dirGroupGeoIPListResponse := target.Data.(*ResponseList)

	return res, dirGroupGeoIPListResponse, nil
}
