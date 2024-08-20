package cato_go_sdk

import (
	"context"
	"net/http"

	"github.com/Yamashou/gqlgenc/clientv2"
)

func New(url string, token string, httpClient http.Client) (*Client, error) {
	catoClient := &Client{
		Client: clientv2.NewClient(&httpClient, url, nil,
			func(ctx context.Context, req *http.Request, gqlInfo *clientv2.GQLRequestInfo, res interface{}, next clientv2.RequestInterceptorFunc) error {
				req.Header.Set("x-api-key", token)

				return next(ctx, req, gqlInfo, res)
			}),
	}

	return catoClient, nil
}
