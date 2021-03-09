/*
@Time : 2021/3/9
@Author : jzd
@Project: go-learning
*/
package go_kong

import (
	"context"
	"github.com/hbagdi/go-kong/kong"
	"net/http"
)

func NewServiceClient(url string) *ServiceClient {
	c := &http.Client{}
	kongClient, _ := kong.NewClient(kong.String(url), c)
	return &ServiceClient{c: kongClient}
}

type ServiceClient struct {
	c *kong.Client
}

func (s *ServiceClient) RemoveAll() error {
	services, err := s.c.Services.ListAll(context.Background())
	if err != nil {
		return err
	}
	for _, service := range services {
		if err := s.c.Services.Delete(context.Background(), service.ID); err != nil {
			return err
		}
	}
	return nil
}
