// Copyright 2021-2024 The sacloud/apprun-api-go authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package apprun

import (
	"context"
	"fmt"

	v1 "github.com/sacloud/apprun-api-go/apis/v1"
)

type TrafficAPI interface {
	// List アプリケーショントラフィック分散を取得
	List(ctx context.Context, appId string) (*v1.HandlerListTraffics, error)
	// Update アプリケーショントラフィック分散を変更
	Update(ctx context.Context, appId string, params *v1.PutTrafficsBody) (*v1.HandlerPutTraffics, error)
}

var _ TrafficAPI = (*trafficOp)(nil)

type trafficOp struct {
	client *Client
}

// NewTrafficOp アプリケーショントラフィック分散関連API
func NewTrafficOp(client *Client) TrafficAPI {
	return &trafficOp{client: client}
}

func (op *trafficOp) List(ctx context.Context, appId string) (*v1.HandlerListTraffics, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.ListApplicationTraffics(ctx, v1.ListApplicationTrafficsParams{ID: appId})
	if err != nil {
		return nil, err
	}
	switch result := resp.(type) {
	case *v1.HandlerListTraffics:
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected list traffics response: %T", resp)
	}
}

func (op *trafficOp) Update(ctx context.Context, appId string, params *v1.PutTrafficsBody) (*v1.HandlerPutTraffics, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	if params == nil {
		return nil, fmt.Errorf("params is nil")
	}
	resp, err := apiClient.PutApplicationTraffic(ctx, *params, v1.PutApplicationTrafficParams{ID: appId})
	if err != nil {
		return nil, err
	}
	switch result := resp.(type) {
	case *v1.HandlerPutTraffics:
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected put traffics response: %T", resp)
	}
}
