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

// コンポーネントの最大CPU数
var ApplicationMaxCPUs = []string{
	(string)(v1.PostApplicationBodyComponentsItemMaxCPU05),
	(string)(v1.PostApplicationBodyComponentsItemMaxCPU1),
	(string)(v1.PostApplicationBodyComponentsItemMaxCPU2),
}

// コンポーネントの最大メモリ
var ApplicationMaxMemories = []string{
	(string)(v1.PostApplicationBodyComponentsItemMaxMemory1Gi),
	(string)(v1.PostApplicationBodyComponentsItemMaxMemory2Gi),
	(string)(v1.PostApplicationBodyComponentsItemMaxMemory4Gi),
}

// ソート順
var ApplicationSortOrders = []string{
	(string)(v1.ListApplicationsSortOrderAsc),
	(string)(v1.ListApplicationsSortOrderDesc),
}

// アプリケーションステータス
var ApplicationStatuses = []string{
	(string)(v1.HandlerGetApplicationOnlyStatusStatusHealthy),
	(string)(v1.HandlerGetApplicationOnlyStatusStatusDeploying),
	(string)(v1.HandlerGetApplicationOnlyStatusStatusUnHealthy),
}

type ApplicationAPI interface {
	// List アプリケーション一覧を取得
	List(ctx context.Context, params *v1.ListApplicationsParams) (*v1.HandlerListApplications, error)
	// Create アプリケーションを作成
	Create(ctx context.Context, params *v1.PostApplicationBody) (*v1.HandlerPostApplication, error)
	// Read アプリケーション詳細を取得
	Read(ctx context.Context, id string) (*v1.HandlerGetApplication, error)
	// Update アプリケーションを部分的に変更
	Update(ctx context.Context, id string, params *v1.PatchApplicationBody) (*v1.HandlerPatchApplication, error)
	// Delete アプリケーションを削除
	Delete(ctx context.Context, id string) error
	// ReadStatus アプリケーションステータスを取得
	ReadStatus(ctx context.Context, id string) (*v1.HandlerGetApplicationOnlyStatus, error)
}

var _ ApplicationAPI = (*applicationOp)(nil)

type applicationOp struct {
	client *Client
}

// NewApplicationOp アプリケーション操作関連API
func NewApplicationOp(client *Client) ApplicationAPI {
	return &applicationOp{client: client}
}

func (op *applicationOp) List(ctx context.Context, params *v1.ListApplicationsParams) (*v1.HandlerListApplications, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	reqParams := v1.ListApplicationsParams{}
	if params != nil {
		reqParams = *params
	}
	resp, err := apiClient.ListApplications(ctx, reqParams)
	if err != nil {
		return nil, err
	}
	switch result := resp.(type) {
	case *v1.HandlerListApplications:
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected list applications response: %T", resp)
	}
}

func (op *applicationOp) Create(ctx context.Context, params *v1.PostApplicationBody) (*v1.HandlerPostApplication, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	if params == nil {
		return nil, fmt.Errorf("params is nil")
	}
	resp, err := apiClient.PostApplication(ctx, params)
	if err != nil {
		return nil, err
	}
	switch result := resp.(type) {
	case *v1.HandlerPostApplication:
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected post application response: %T", resp)
	}
}

func (op *applicationOp) Update(ctx context.Context, id string, params *v1.PatchApplicationBody) (*v1.HandlerPatchApplication, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	if params == nil {
		return nil, fmt.Errorf("params is nil")
	}
	resp, err := apiClient.PatchApplication(ctx, params, v1.PatchApplicationParams{ID: id})
	if err != nil {
		return nil, err
	}
	switch result := resp.(type) {
	case *v1.HandlerPatchApplication:
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected patch application response: %T", resp)
	}
}

func (op *applicationOp) Read(ctx context.Context, id string) (*v1.HandlerGetApplication, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.GetApplication(ctx, v1.GetApplicationParams{ID: id})
	if err != nil {
		return nil, err
	}
	switch result := resp.(type) {
	case *v1.HandlerGetApplication:
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected get application response: %T", resp)
	}
}

func (op *applicationOp) Delete(ctx context.Context, id string) error {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return err
	}
	resp, err := apiClient.DeleteApplication(ctx, v1.DeleteApplicationParams{ID: id})
	if err != nil {
		return err
	}
	switch resp.(type) {
	case *v1.DeleteApplicationNoContent:
		return nil
	default:
		return fmt.Errorf("unexpected delete application response: %T", resp)
	}
}

func (op *applicationOp) ReadStatus(ctx context.Context, id string) (*v1.HandlerGetApplicationOnlyStatus, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.GetApplicationStatus(ctx, v1.GetApplicationStatusParams{ID: id})
	if err != nil {
		return nil, err
	}
	switch result := resp.(type) {
	case *v1.HandlerGetApplicationOnlyStatus:
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected get application status response: %T", resp)
	}
}
