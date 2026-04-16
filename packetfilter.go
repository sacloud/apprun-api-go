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

	"github.com/google/uuid"
	v1 "github.com/sacloud/apprun-api-go/apis/v1"
)

type PacketFilterAPI interface {
	// Read パケットフィルタ詳細を取得
	Read(ctx context.Context, appId string) (*v1.HandlerGetPacketFilter, error)
	// Update パケットフィルタを部分的に変更
	Update(ctx context.Context, appId string, params *v1.PatchPacketFilter) (*v1.HandlerPatchPacketFilter, error)
}

var _ PacketFilterAPI = (*packetFilterOp)(nil)

type packetFilterOp struct {
	client *Client
}

// NewPacketFilterOp アプリケーショントラフィック分散関連API
func NewPacketFilterOp(client *Client) PacketFilterAPI {
	return &packetFilterOp{client: client}
}

func (op *packetFilterOp) Read(ctx context.Context, appId string) (*v1.HandlerGetPacketFilter, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(appId)
	if err != nil {
		return nil, err
	}

	resp, err := apiClient.GetPacketFilter(ctx, v1.GetPacketFilterParams{ID: id})
	if err != nil {
		return nil, err
	}
	switch result := resp.(type) {
	case *v1.HandlerGetPacketFilter:
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected get packet filter response: %T", resp)
	}
}

func (op *packetFilterOp) Update(ctx context.Context, appId string, params *v1.PatchPacketFilter) (*v1.HandlerPatchPacketFilter, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(appId)
	if err != nil {
		return nil, err
	}

	if params == nil {
		return nil, fmt.Errorf("params is nil")
	}
	resp, err := apiClient.PatchPacketFilter(ctx, params, v1.PatchPacketFilterParams{ID: id})
	if err != nil {
		return nil, err
	}
	switch result := resp.(type) {
	case *v1.HandlerPatchPacketFilter:
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected patch packet filter response: %T", resp)
	}
}
