package ip

import (
	"net/http"

	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/errors"
)

const (
	serviceName = "AccGroupGeoGroup"
	basePath    = "dirgroups/" // will be: account/{account}/dirgroup/{group}
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

func (s *Service) Read(groupName string) (*http.Response, *Response, error) {
	target := client.Target(&Response{})

	if s.c == nil {
		return nil, nil, errors.ServiceError(serviceName)
	}

	res, err := s.c.Do(http.MethodGet, basePath+groupName, nil, target)
	if err != nil {
		return nil, nil, errors.ReadError(serviceName, groupName, err)
	}

	groupResponse := target.Data.(*Response)
	return res, groupResponse, nil
}
