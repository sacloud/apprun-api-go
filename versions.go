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

// ソート順
var VersionSortOrders = []string{
	(string)(v1.ListApplicationVersionsSortOrderAsc),
	(string)(v1.ListApplicationVersionsSortOrderDesc),
}

// バージョンステータス
var VersionStatuses = []string{
	(string)(v1.HandlerGetApplicationVersionOnlyStatusStatusHealthy),
	(string)(v1.HandlerGetApplicationVersionOnlyStatusStatusDeploying),
	(string)(v1.HandlerGetApplicationVersionOnlyStatusStatusUnHealthy),
}

type VersionAPI interface {
	// List アプリケーションバージョン一覧を取得
	List(ctx context.Context, appId string, params *v1.ListApplicationVersionsParams) (*v1.HandlerListVersions, error)
	// Read アプリケーションバージョン詳細を取得
	Read(ctx context.Context, appId, versionId string) (*v1.HandlerGetVersion, error)
	// Delete アプリケーションバージョンを削除
	Delete(ctx context.Context, appId, versionId string) error

	ReadStatus(ctx context.Context, appId, versionId string) (*v1.HandlerGetApplicationVersionOnlyStatus, error)
}

var _ VersionAPI = (*versionOp)(nil)

type versionOp struct {
	client *Client
}

// NewVersionOp アプリケーションバージョン操作関連API
func NewVersionOp(client *Client) VersionAPI {
	return &versionOp{client: client}
}

func (op *versionOp) List(ctx context.Context, appId string, params *v1.ListApplicationVersionsParams) (*v1.HandlerListVersions, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	reqParams := v1.ListApplicationVersionsParams{ID: appId}
	if params != nil {
		reqParams = *params
		reqParams.ID = appId
	}
	resp, err := apiClient.ListApplicationVersions(ctx, reqParams)
	if err != nil {
		return nil, err
	}
	switch result := resp.(type) {
	case *v1.HandlerListVersions:
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected list versions response: %T", resp)
	}
}

func (op *versionOp) Read(ctx context.Context, appId, versionId string) (*v1.HandlerGetVersion, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.GetApplicationVersion(ctx, v1.GetApplicationVersionParams{ID: appId, VersionID: versionId})
	if err != nil {
		return nil, err
	}
	switch result := resp.(type) {
	case *v1.HandlerGetVersion:
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected get version response: %T", resp)
	}
}

func (op *versionOp) Delete(ctx context.Context, appId, versionId string) error {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return err
	}
	resp, err := apiClient.DeleteApplicationVersion(ctx, v1.DeleteApplicationVersionParams{ID: appId, VersionID: versionId})
	if err != nil {
		return err
	}
	switch resp.(type) {
	case *v1.DeleteApplicationVersionNoContent:
		return nil
	default:
		return fmt.Errorf("unexpected delete version response: %T", resp)
	}
}

func (op *versionOp) ReadStatus(ctx context.Context, appId, versionId string) (*v1.HandlerGetApplicationVersionOnlyStatus, error) {
	apiClient, err := op.client.apiClient()
	if err != nil {
		return nil, err
	}
	resp, err := apiClient.GetApplicationVersionStatus(ctx, v1.GetApplicationVersionStatusParams{ID: appId, VersionID: versionId})
	if err != nil {
		return nil, err
	}
	switch result := resp.(type) {
	case *v1.HandlerGetApplicationVersionOnlyStatus:
		return result, nil
	default:
		return nil, fmt.Errorf("unexpected get version status response: %T", resp)
	}
}
