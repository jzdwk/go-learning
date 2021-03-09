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

func NewRouteClient(url string) *RouteClient {
	c := &http.Client{}
	kongClient, _ := kong.NewClient(kong.String(url), c)
	return &RouteClient{c: kongClient}
}

type RouteClient struct {
	c *kong.Client
}

func (r *RouteClient) RemoveAll() error {
	routes, err := r.c.Routes.ListAll(context.Background())
	if err != nil {
		return err
	}
	for _, route := range routes {
		if err := r.c.Routes.Delete(context.Background(), route.ID); err != nil {
			return err
		}
	}
	return nil
}
