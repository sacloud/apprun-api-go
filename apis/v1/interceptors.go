package v1

import (
	"context"
	"net/http"
)

// AppRunAuthInterceptor AppRun APIリクエストに認証情報の注入を行う
func AppRunAuthInterceptor(token, secret string) func(context.Context, *http.Request) error {
	return func(ctx context.Context, req *http.Request) error {
		req.SetBasicAuth(token, secret)
		return nil
	}
}
