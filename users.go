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

type UserAPI interface {
	// Read ログイン中のユーザー情報を取得
	Read(ctx context.Context) (*v1.HandlerGetUser, error)
	// Create さくらのAppRunにサインアップ
	Create(ctx context.Context) (*v1.HandlerPostUser, error)
}

var _ UserAPI = (*userOp)(nil)

type userOp struct {
	client *Client
}

// NewUserOp ユーザー操作関連API
func NewUserOp(client *Client) UserAPI {
	return &userOp{client: client}
}

func (op *userOp) Read(ctx context.Context) (*v1.HandlerGetUser, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.GetUser(ctx)
	if err != nil {
		return nil, err
	}
	switch result := resp.(type) {
	case *v1.HandlerGetUser:
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected get user response: %T", resp)
	}
}

func (op *userOp) Create(ctx context.Context) (*v1.HandlerPostUser, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.PostUser(ctx)
	if err != nil {
		return nil, err
	}
	switch result := resp.(type) {
	case *v1.HandlerPostUser:
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected post user response: %T", resp)
	}
}
