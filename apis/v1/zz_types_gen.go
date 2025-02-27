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

// Package v1 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package v1

import (
	"encoding/json"
	"time"

	"github.com/oapi-codegen/runtime"
)

// Defines values for ApplicationStatus.
const (
	ApplicationStatusDeploying ApplicationStatus = "Deploying"
	ApplicationStatusHealthy   ApplicationStatus = "Healthy"
	ApplicationStatusUnHealthy ApplicationStatus = "UnHealthy"
)

// Defines values for VersionStatus.
const (
	VersionStatusDeploying VersionStatus = "Deploying"
	VersionStatusHealthy   VersionStatus = "Healthy"
	VersionStatusUnHealthy VersionStatus = "UnHealthy"
)

// Defines values for HandlerGetApplicationStatusStatus.
const (
	HandlerGetApplicationStatusStatusDeploying HandlerGetApplicationStatusStatus = "Deploying"
	HandlerGetApplicationStatusStatusHealthy   HandlerGetApplicationStatusStatus = "Healthy"
	HandlerGetApplicationStatusStatusUnHealthy HandlerGetApplicationStatusStatus = "UnHealthy"
)

// Defines values for HandlerGetVersionStatus.
const (
	HandlerGetVersionStatusDeploying HandlerGetVersionStatus = "Deploying"
	HandlerGetVersionStatusHealthy   HandlerGetVersionStatus = "Healthy"
	HandlerGetVersionStatusUnHealthy HandlerGetVersionStatus = "UnHealthy"
)

// Defines values for HandlerListApplicationsDataStatus.
const (
	HandlerListApplicationsDataStatusDeploying HandlerListApplicationsDataStatus = "Deploying"
	HandlerListApplicationsDataStatusHealthy   HandlerListApplicationsDataStatus = "Healthy"
	HandlerListApplicationsDataStatusUnHealthy HandlerListApplicationsDataStatus = "UnHealthy"
)

// Defines values for HandlerListApplicationsMetaSortOrder.
const (
	HandlerListApplicationsMetaSortOrderAsc  HandlerListApplicationsMetaSortOrder = "asc"
	HandlerListApplicationsMetaSortOrderDesc HandlerListApplicationsMetaSortOrder = "desc"
)

// Defines values for HandlerListVersionsMetaSortOrder.
const (
	HandlerListVersionsMetaSortOrderAsc  HandlerListVersionsMetaSortOrder = "asc"
	HandlerListVersionsMetaSortOrderDesc HandlerListVersionsMetaSortOrder = "desc"
)

// Defines values for HandlerPatchApplicationStatus.
const (
	HandlerPatchApplicationStatusDeploying HandlerPatchApplicationStatus = "Deploying"
	HandlerPatchApplicationStatusHealthy   HandlerPatchApplicationStatus = "Healthy"
	HandlerPatchApplicationStatusUnHealthy HandlerPatchApplicationStatus = "UnHealthy"
)

// Defines values for ModelErrorLocationType.
const (
	ModelErrorLocationTypeBody      ModelErrorLocationType = "body"
	ModelErrorLocationTypeHeader    ModelErrorLocationType = "header"
	ModelErrorLocationTypeParameter ModelErrorLocationType = "parameter"
	ModelErrorLocationTypeQuery     ModelErrorLocationType = "query"
)

// Defines values for PatchApplicationBodyComponentMaxCpu.
const (
	PatchApplicationBodyComponentMaxCpuN01 PatchApplicationBodyComponentMaxCpu = "0.1"
	PatchApplicationBodyComponentMaxCpuN02 PatchApplicationBodyComponentMaxCpu = "0.2"
	PatchApplicationBodyComponentMaxCpuN03 PatchApplicationBodyComponentMaxCpu = "0.3"
	PatchApplicationBodyComponentMaxCpuN04 PatchApplicationBodyComponentMaxCpu = "0.4"
	PatchApplicationBodyComponentMaxCpuN05 PatchApplicationBodyComponentMaxCpu = "0.5"
	PatchApplicationBodyComponentMaxCpuN06 PatchApplicationBodyComponentMaxCpu = "0.6"
	PatchApplicationBodyComponentMaxCpuN07 PatchApplicationBodyComponentMaxCpu = "0.7"
	PatchApplicationBodyComponentMaxCpuN08 PatchApplicationBodyComponentMaxCpu = "0.8"
	PatchApplicationBodyComponentMaxCpuN09 PatchApplicationBodyComponentMaxCpu = "0.9"
	PatchApplicationBodyComponentMaxCpuN1  PatchApplicationBodyComponentMaxCpu = "1"
)

// Defines values for PatchApplicationBodyComponentMaxMemory.
const (
	PatchApplicationBodyComponentMaxMemoryN1Gi   PatchApplicationBodyComponentMaxMemory = "1Gi"
	PatchApplicationBodyComponentMaxMemoryN256Mi PatchApplicationBodyComponentMaxMemory = "256Mi"
	PatchApplicationBodyComponentMaxMemoryN2Gi   PatchApplicationBodyComponentMaxMemory = "2Gi"
	PatchApplicationBodyComponentMaxMemoryN512Mi PatchApplicationBodyComponentMaxMemory = "512Mi"
)

// Defines values for PostApplicationBodyComponentMaxCpu.
const (
	PostApplicationBodyComponentMaxCpuN01 PostApplicationBodyComponentMaxCpu = "0.1"
	PostApplicationBodyComponentMaxCpuN02 PostApplicationBodyComponentMaxCpu = "0.2"
	PostApplicationBodyComponentMaxCpuN03 PostApplicationBodyComponentMaxCpu = "0.3"
	PostApplicationBodyComponentMaxCpuN04 PostApplicationBodyComponentMaxCpu = "0.4"
	PostApplicationBodyComponentMaxCpuN05 PostApplicationBodyComponentMaxCpu = "0.5"
	PostApplicationBodyComponentMaxCpuN06 PostApplicationBodyComponentMaxCpu = "0.6"
	PostApplicationBodyComponentMaxCpuN07 PostApplicationBodyComponentMaxCpu = "0.7"
	PostApplicationBodyComponentMaxCpuN08 PostApplicationBodyComponentMaxCpu = "0.8"
	PostApplicationBodyComponentMaxCpuN09 PostApplicationBodyComponentMaxCpu = "0.9"
	PostApplicationBodyComponentMaxCpuN1  PostApplicationBodyComponentMaxCpu = "1"
)

// Defines values for PostApplicationBodyComponentMaxMemory.
const (
	PostApplicationBodyComponentMaxMemoryN1Gi   PostApplicationBodyComponentMaxMemory = "1Gi"
	PostApplicationBodyComponentMaxMemoryN256Mi PostApplicationBodyComponentMaxMemory = "256Mi"
	PostApplicationBodyComponentMaxMemoryN2Gi   PostApplicationBodyComponentMaxMemory = "2Gi"
	PostApplicationBodyComponentMaxMemoryN512Mi PostApplicationBodyComponentMaxMemory = "512Mi"
)

// Defines values for ListApplicationsParamsSortOrder.
const (
	ListApplicationsParamsSortOrderAsc  ListApplicationsParamsSortOrder = "asc"
	ListApplicationsParamsSortOrderDesc ListApplicationsParamsSortOrder = "desc"
)

// Defines values for ListApplicationVersionsParamsSortOrder.
const (
	ListApplicationVersionsParamsSortOrderAsc  ListApplicationVersionsParamsSortOrder = "asc"
	ListApplicationVersionsParamsSortOrderDesc ListApplicationVersionsParamsSortOrder = "desc"
)

// Application defines model for Application.
type Application struct {
	// Components アプリケーションのコンポーネント情報
	Components []HandlerApplicationComponent `json:"components"`

	// CreatedAt 作成日時
	CreatedAt time.Time `json:"created_at"`

	// Id アプリケーションID
	Id string `json:"id"`

	// MaxScale アプリケーション全体の最大スケール数
	MaxScale int `json:"max_scale"`

	// MinScale アプリケーション全体の最小スケール数
	MinScale int `json:"min_scale"`

	// Name アプリケーション名
	Name string `json:"name"`

	// Port アプリケーションがリクエストを待ち受けるポート番号
	Port int `json:"port"`

	// PublicUrl 公開URL
	PublicUrl string `json:"public_url"`

	// Status アプリケーションステータス
	Status ApplicationStatus `json:"status"`

	// TimeoutSeconds アプリケーションの公開URLにアクセスして、インスタンスが起動してからレスポンスが返るまでの時間制限
	TimeoutSeconds int `json:"timeout_seconds"`
}

// ApplicationStatus アプリケーションステータス
type ApplicationStatus string

// Traffic defines model for Traffic.
type Traffic struct {
	// IsLatestVersion 最新バージョンかどうか
	IsLatestVersion bool `json:"is_latest_version"`

	// Percent トラフィック分散の割合
	Percent int `json:"percent"`

	// VersionName バージョン名
	VersionName string `json:"version_name"`
}

// Version defines model for Version.
type Version struct {
	// CreatedAt 作成日時
	CreatedAt time.Time `json:"created_at"`

	// Id バージョンID
	Id string `json:"id"`

	// Name バージョン名
	Name string `json:"name"`

	// Status ステータス
	Status VersionStatus `json:"status"`
}

// VersionStatus ステータス
type VersionStatus string

// HandlerApplicationComponent defines model for handler.ApplicationComponent.
type HandlerApplicationComponent struct {
	// DeploySource コンポーネントを構成するソース
	DeploySource HandlerApplicationComponentDeploySource `json:"deploy_source"`

	// Env コンポーネントに渡す環境変数
	Env *[]HandlerApplicationComponentEnv `json:"env,omitempty"`

	// MaxCpu コンポーネントの最大CPU数
	MaxCpu string `json:"max_cpu"`

	// MaxMemory コンポーネントの最大メモリ
	MaxMemory string `json:"max_memory"`

	// Name コンポーネント名
	Name string `json:"name"`

	// Probe コンポーネントのプローブ設定
	Probe *HandlerApplicationComponentProbe `json:"probe"`
}

// HandlerApplicationComponentDeploySource コンポーネントを構成するソース
type HandlerApplicationComponentDeploySource struct {
	// ContainerRegistry コンテナレジストリ
	ContainerRegistry *HandlerApplicationComponentDeploySourceContainerRegistry `json:"container_registry,omitempty"`
}

// HandlerApplicationComponentDeploySourceContainerRegistry コンテナレジストリ
type HandlerApplicationComponentDeploySourceContainerRegistry struct {
	// Image コンテナイメージ名
	Image string `json:"image"`

	// Server コンテナレジストリのサーバー名
	Server *string `json:"server,omitempty"`

	// Username コンテナレジストリの認証情報
	Username *string `json:"username,omitempty"`
}

// HandlerApplicationComponentEnv defines model for handler.ApplicationComponentEnv.
type HandlerApplicationComponentEnv struct {
	// Key 環境変数名
	Key *string `json:"key,omitempty"`

	// Value 環境変数の値
	Value *string `json:"value,omitempty"`
}

// HandlerApplicationComponentProbe コンポーネントのプローブ設定
type HandlerApplicationComponentProbe struct {
	// HttpGet HTTP Getプローブタイプ
	HttpGet *HandlerApplicationComponentProbeHttpGet `json:"http_get"`
}

// HandlerApplicationComponentProbeHttpGet HTTP Getプローブタイプ
type HandlerApplicationComponentProbeHttpGet struct {
	Headers *[]HandlerApplicationComponentProbeHttpGetHeader `json:"headers,omitempty"`

	// Path HTTPサーバーへアクセスしプローブをチェックする際のパス
	Path string `json:"path"`

	// Port HTTPサーバーへアクセスしプローブをチェックする際のポート番号
	Port int `json:"port"`
}

// HandlerApplicationComponentProbeHttpGetHeader defines model for handler.ApplicationComponentProbeHttpGetHeader.
type HandlerApplicationComponentProbeHttpGetHeader struct {
	// Name ヘッダーフィールド名
	Name *string `json:"name,omitempty"`

	// Value ヘッダーフィールド値
	Value *string `json:"value,omitempty"`
}

// HandlerGetApplication defines model for handler.getApplication.
type HandlerGetApplication = Application

// HandlerGetApplicationStatusResponse defines model for handler.getApplicationStatus.
type HandlerGetApplicationStatusResponse struct {
	// Message ステータス失敗時のメッセージ
	Message string `json:"message"`

	// Status アプリケーションステータス
	Status HandlerGetApplicationStatusStatus `json:"status"`
}

// HandlerGetApplicationStatusStatus アプリケーションステータス
type HandlerGetApplicationStatusStatus string

// HandlerGetVersion defines model for handler.getVersion.
type HandlerGetVersion struct {
	// Components バージョンのコンポーネント情報
	Components []HandlerApplicationComponent `json:"components"`

	// CreatedAt 作成日時
	CreatedAt time.Time `json:"created_at"`

	// Id バージョンID
	Id string `json:"id"`

	// MaxScale バージョンの最大スケール数
	MaxScale int `json:"max_scale"`

	// MinScale バージョンの最小スケール数
	MinScale int `json:"min_scale"`

	// Name バージョン名
	Name string `json:"name"`

	// Port アプリケーションがリクエストを待ち受けるポート番号
	Port int `json:"port"`

	// Status バージョンステータス
	Status HandlerGetVersionStatus `json:"status"`

	// TimeoutSeconds アプリケーションの公開URLにアクセスして、インスタンスが起動してからレスポンスが返るまでの時間制限
	TimeoutSeconds int `json:"timeout_seconds"`
}

// HandlerGetVersionStatus バージョンステータス
type HandlerGetVersionStatus string

// HandlerListApplications defines model for handler.listApplications.
type HandlerListApplications struct {
	Data []HandlerListApplicationsData `json:"data"`
	Meta HandlerListApplicationsMeta   `json:"meta"`
}

// HandlerListApplicationsData defines model for handler.listApplicationsData.
type HandlerListApplicationsData struct {
	// CreatedAt 作成日時
	CreatedAt time.Time `json:"created_at"`

	// Id アプリケーションID
	Id string `json:"id"`

	// Name アプリケーション名
	Name string `json:"name"`

	// PublicUrl 公開URL
	PublicUrl string `json:"public_url"`

	// Status アプリケーションステータス
	Status HandlerListApplicationsDataStatus `json:"status"`
}

// HandlerListApplicationsDataStatus アプリケーションステータス
type HandlerListApplicationsDataStatus string

// HandlerListApplicationsMeta defines model for handler.listApplicationsMeta.
type HandlerListApplicationsMeta struct {
	ObjectTotal int                                  `json:"object_total"`
	PageNum     int                                  `json:"page_num"`
	PageSize    int                                  `json:"page_size"`
	SortField   string                               `json:"sort_field"`
	SortOrder   HandlerListApplicationsMetaSortOrder `json:"sort_order"`
}

// HandlerListApplicationsMetaSortOrder defines model for HandlerListApplicationsMeta.SortOrder.
type HandlerListApplicationsMetaSortOrder string

// HandlerListTraffics defines model for handler.listTraffics.
type HandlerListTraffics struct {
	Data []Traffic               `json:"data"`
	Meta *map[string]interface{} `json:"meta"`
}

// HandlerListVersions defines model for handler.listVersions.
type HandlerListVersions struct {
	Data []Version               `json:"data"`
	Meta HandlerListVersionsMeta `json:"meta"`
}

// HandlerListVersionsMeta defines model for handler.listVersionsMeta.
type HandlerListVersionsMeta struct {
	ObjectTotal int                              `json:"object_total"`
	PageNum     int                              `json:"page_num"`
	PageSize    int                              `json:"page_size"`
	SortField   string                           `json:"sort_field"`
	SortOrder   HandlerListVersionsMetaSortOrder `json:"sort_order"`
}

// HandlerListVersionsMetaSortOrder defines model for HandlerListVersionsMeta.SortOrder.
type HandlerListVersionsMetaSortOrder string

// HandlerPatchApplication defines model for handler.patchApplication.
type HandlerPatchApplication struct {
	// Components アプリケーションのコンポーネント情報
	Components []HandlerApplicationComponent `json:"components"`

	// Id アプリケーションID
	Id string `json:"id"`

	// MaxScale アプリケーション全体の最大スケール数
	MaxScale int `json:"max_scale"`

	// MinScale アプリケーション全体の最小スケール数
	MinScale int `json:"min_scale"`

	// Name アプリケーション名
	Name string `json:"name"`

	// Port アプリケーションがリクエストを待ち受けるポート番号
	Port int `json:"port"`

	// PublicUrl 公開URL
	PublicUrl string `json:"public_url"`

	// Status アプリケーションステータス
	Status HandlerPatchApplicationStatus `json:"status"`

	// TimeoutSeconds アプリケーションの公開URLにアクセスして、インスタンスが起動してからレスポンスが返るまでの時間制限
	TimeoutSeconds int `json:"timeout_seconds"`

	// UpdatedAt 更新日時
	UpdatedAt time.Time `json:"updated_at"`
}

// HandlerPatchApplicationStatus アプリケーションステータス
type HandlerPatchApplicationStatus string

// HandlerPostApplication defines model for handler.postApplication.
type HandlerPostApplication = Application

// HandlerPutTraffics defines model for handler.putTraffics.
type HandlerPutTraffics struct {
	Data []Traffic               `json:"data"`
	Meta *map[string]interface{} `json:"meta"`
}

// ModelAppRunError defines model for model.appRunError.
type ModelAppRunError struct {
	union json.RawMessage
}

// ModelCloudctrlError defines model for model.cloudctrlError.
type ModelCloudctrlError struct {
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	IsFatal   bool   `json:"is_fatal"`
	Serial    string `json:"serial"`
	Status    string `json:"status"`
}

// ModelDefaultError defines model for model.defaultError.
type ModelDefaultError struct {
	Detail struct {
		Code    int         `json:"code"`
		Errors  ModelErrors `json:"errors"`
		Message string      `json:"message"`
	} `json:"error"`
}

// ModelError defines model for model.error.
type ModelError struct {
	Domain       *string                 `json:"domain"`
	Location     *string                 `json:"location"`
	LocationType *ModelErrorLocationType `json:"location_type"`
	Message      *string                 `json:"message"`
	Reason       *string                 `json:"reason"`
}

// ModelErrorLocationType defines model for ModelError.LocationType.
type ModelErrorLocationType string

// ModelErrors defines model for model.errors.
type ModelErrors = []ModelError

// PatchApplicationBody defines model for patchApplicationBody.
type PatchApplicationBody struct {
	// AllTrafficAvailable アプリケーションを最新のバージョンにすべてのトラフィックを割り当てるかどうか
	AllTrafficAvailable *bool `json:"all_traffic_available,omitempty"`

	// Components アプリケーションのコンポーネント情報
	Components *[]PatchApplicationBodyComponent `json:"components,omitempty"`

	// MaxScale アプリケーション全体の最大スケール数
	MaxScale *int `json:"max_scale,omitempty"`

	// MinScale アプリケーション全体の最小スケール数
	MinScale *int `json:"min_scale,omitempty"`

	// Port アプリケーションがリクエストを待ち受けるポート番号
	Port *int `json:"port,omitempty"`

	// TimeoutSeconds アプリケーションの公開URLにアクセスして、インスタンスが起動してからレスポンスが返るまでの時間制限
	TimeoutSeconds *int `json:"timeout_seconds,omitempty"`
}

// PatchApplicationBodyComponent defines model for patchApplicationBodyComponent.
type PatchApplicationBodyComponent struct {
	// DeploySource コンポーネントを構成するソース
	DeploySource PatchApplicationBodyComponentDeploySource `json:"deploy_source"`

	// Env コンポーネントに渡す環境変数
	Env *[]PatchApplicationBodyComponentEnv `json:"env"`

	// MaxCpu コンポーネントの最大CPU数
	MaxCpu PatchApplicationBodyComponentMaxCpu `json:"max_cpu"`

	// MaxMemory コンポーネントの最大メモリ
	MaxMemory PatchApplicationBodyComponentMaxMemory `json:"max_memory"`

	// Name コンポーネント名
	Name string `json:"name"`

	// Probe コンポーネントのプローブ設定
	Probe *PatchApplicationBodyComponentProbe `json:"probe"`
}

// PatchApplicationBodyComponentMaxCpu コンポーネントの最大CPU数
type PatchApplicationBodyComponentMaxCpu string

// PatchApplicationBodyComponentMaxMemory コンポーネントの最大メモリ
type PatchApplicationBodyComponentMaxMemory string

// PatchApplicationBodyComponentDeploySource コンポーネントを構成するソース
type PatchApplicationBodyComponentDeploySource struct {
	// ContainerRegistry コンテナレジストリ
	ContainerRegistry *PatchApplicationBodyComponentDeploySourceContainerRegistry `json:"container_registry,omitempty"`
}

// PatchApplicationBodyComponentDeploySourceContainerRegistry コンテナレジストリ
type PatchApplicationBodyComponentDeploySourceContainerRegistry struct {
	// Image コンテナイメージ名
	Image string `json:"image"`

	// Password コンテナレジストリの認証情報
	Password *string `json:"password"`

	// Server コンテナレジストリのサーバー名
	Server *string `json:"server"`

	// Username コンテナレジストリの認証情報
	Username *string `json:"username"`
}

// PatchApplicationBodyComponentEnv defines model for patchApplicationBodyComponentEnv.
type PatchApplicationBodyComponentEnv struct {
	// Key 環境変数名
	Key *string `json:"key,omitempty"`

	// Value 環境変数の値
	Value *string `json:"value,omitempty"`
}

// PatchApplicationBodyComponentProbe コンポーネントのプローブ設定
type PatchApplicationBodyComponentProbe struct {
	// HttpGet HTTP Getプローブタイプ
	HttpGet *PatchApplicationBodyComponentProbeHttpGet `json:"http_get"`
}

// PatchApplicationBodyComponentProbeHttpGet HTTP Getプローブタイプ
type PatchApplicationBodyComponentProbeHttpGet struct {
	Headers *[]PatchApplicationBodyComponentProbeHttpGetHeader `json:"headers,omitempty"`

	// Path HTTPサーバーへアクセスしプローブをチェックする際のパス
	Path string `json:"path"`

	// Port HTTPサーバーへアクセスしプローブをチェックする際のポート番号
	Port int `json:"port"`
}

// PatchApplicationBodyComponentProbeHttpGetHeader defines model for patchApplicationBodyComponentProbeHttpGetHeader.
type PatchApplicationBodyComponentProbeHttpGetHeader struct {
	// Name ヘッダーフィールド名
	Name *string `json:"name,omitempty"`

	// Value ヘッダーフィールド値
	Value *string `json:"value,omitempty"`
}

// PostApplicationBody defines model for postApplicationBody.
type PostApplicationBody struct {
	// Components アプリケーションのコンポーネント情報
	Components []PostApplicationBodyComponent `json:"components"`

	// MaxScale アプリケーション全体の最大スケール数
	MaxScale int `json:"max_scale"`

	// MinScale アプリケーション全体の最小スケール数
	MinScale int `json:"min_scale"`

	// Name アプリケーション名
	Name string `json:"name"`

	// Port アプリケーションがリクエストを待ち受けるポート番号
	Port int `json:"port"`

	// TimeoutSeconds アプリケーションの公開URLにアクセスして、インスタンスが起動してからレスポンスが返るまでの時間制限
	TimeoutSeconds int `json:"timeout_seconds"`
}

// PostApplicationBodyComponent defines model for postApplicationBodyComponent.
type PostApplicationBodyComponent struct {
	// DeploySource コンポーネントを構成するソース
	DeploySource PostApplicationBodyComponentDeploySource `json:"deploy_source"`

	// Env コンポーネントに渡す環境変数
	Env *[]PostApplicationBodyComponentEnv `json:"env"`

	// MaxCpu コンポーネントの最大CPU数
	MaxCpu PostApplicationBodyComponentMaxCpu `json:"max_cpu"`

	// MaxMemory コンポーネントの最大メモリ
	MaxMemory PostApplicationBodyComponentMaxMemory `json:"max_memory"`

	// Name コンポーネント名
	Name string `json:"name"`

	// Probe コンポーネントのプローブ設定
	Probe *PostApplicationBodyComponentProbe `json:"probe"`
}

// PostApplicationBodyComponentMaxCpu コンポーネントの最大CPU数
type PostApplicationBodyComponentMaxCpu string

// PostApplicationBodyComponentMaxMemory コンポーネントの最大メモリ
type PostApplicationBodyComponentMaxMemory string

// PostApplicationBodyComponentDeploySource コンポーネントを構成するソース
type PostApplicationBodyComponentDeploySource struct {
	// ContainerRegistry コンテナレジストリ
	ContainerRegistry *PostApplicationBodyComponentDeploySourceContainerRegistry `json:"container_registry,omitempty"`
}

// PostApplicationBodyComponentDeploySourceContainerRegistry コンテナレジストリ
type PostApplicationBodyComponentDeploySourceContainerRegistry struct {
	// Image コンテナイメージ名
	Image string `json:"image"`

	// Password コンテナレジストリの認証情報
	Password *string `json:"password"`

	// Server コンテナレジストリのサーバー名
	Server *string `json:"server"`

	// Username コンテナレジストリの認証情報
	Username *string `json:"username"`
}

// PostApplicationBodyComponentEnv defines model for postApplicationBodyComponentEnv.
type PostApplicationBodyComponentEnv struct {
	// Key 環境変数名
	Key *string `json:"key,omitempty"`

	// Value 環境変数の値
	Value *string `json:"value,omitempty"`
}

// PostApplicationBodyComponentProbe コンポーネントのプローブ設定
type PostApplicationBodyComponentProbe struct {
	// HttpGet HTTP Getプローブタイプ
	HttpGet *PostApplicationBodyComponentProbeHttpGet `json:"http_get"`
}

// PostApplicationBodyComponentProbeHttpGet HTTP Getプローブタイプ
type PostApplicationBodyComponentProbeHttpGet struct {
	Headers *[]PostApplicationBodyComponentProbeHttpGetHeader `json:"headers,omitempty"`

	// Path HTTPサーバーへアクセスしプローブをチェックする際のパス
	Path string `json:"path"`

	// Port HTTPサーバーへアクセスしプローブをチェックする際のポート番号
	Port int `json:"port"`
}

// PostApplicationBodyComponentProbeHttpGetHeader defines model for postApplicationBodyComponentProbeHttpGetHeader.
type PostApplicationBodyComponentProbeHttpGetHeader struct {
	// Name ヘッダーフィールド名
	Name *string `json:"name,omitempty"`

	// Value ヘッダーフィールド値
	Value *string `json:"value,omitempty"`
}

// PutTrafficsBody defines model for putTrafficsBody.
type PutTrafficsBody = []Traffic

// ListApplicationsParams defines parameters for ListApplications.
type ListApplicationsParams struct {
	// PageNum 表示したいページ番号
	PageNum *int `form:"page_num,omitempty" json:"page_num,omitempty"`

	// PageSize 表示したい1ページあたりのサイズ
	PageSize *int `form:"page_size,omitempty" json:"page_size,omitempty"`

	// SortField ソートしたいフィールド名
	SortField *string `form:"sort_field,omitempty" json:"sort_field,omitempty"`

	// SortOrder ソート順（昇順、降順）
	SortOrder *ListApplicationsParamsSortOrder `form:"sort_order,omitempty" json:"sort_order,omitempty"`
}

// ListApplicationsParamsSortOrder defines parameters for ListApplications.
type ListApplicationsParamsSortOrder string

// ListApplicationVersionsParams defines parameters for ListApplicationVersions.
type ListApplicationVersionsParams struct {
	// PageNum 表示したいページ番号
	PageNum *int `form:"page_num,omitempty" json:"page_num,omitempty"`

	// PageSize 表示したい1ページあたりのサイズ
	PageSize *int `form:"page_size,omitempty" json:"page_size,omitempty"`

	// SortField ソートしたいフィールド名
	SortField *string `form:"sort_field,omitempty" json:"sort_field,omitempty"`

	// SortOrder ソート順（昇順、降順）
	SortOrder *ListApplicationVersionsParamsSortOrder `form:"sort_order,omitempty" json:"sort_order,omitempty"`
}

// ListApplicationVersionsParamsSortOrder defines parameters for ListApplicationVersions.
type ListApplicationVersionsParamsSortOrder string

// PostApplicationJSONRequestBody defines body for PostApplication for application/json ContentType.
type PostApplicationJSONRequestBody = PostApplicationBody

// PatchApplicationJSONRequestBody defines body for PatchApplication for application/json ContentType.
type PatchApplicationJSONRequestBody = PatchApplicationBody

// PutApplicationTrafficJSONRequestBody defines body for PutApplicationTraffic for application/json ContentType.
type PutApplicationTrafficJSONRequestBody = PutTrafficsBody

// AsModelDefaultError returns the union data inside the ModelAppRunError as a ModelDefaultError
func (t ModelAppRunError) AsModelDefaultError() (ModelDefaultError, error) {
	var body ModelDefaultError
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromModelDefaultError overwrites any union data inside the ModelAppRunError as the provided ModelDefaultError
func (t *ModelAppRunError) FromModelDefaultError(v ModelDefaultError) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeModelDefaultError performs a merge with any union data inside the ModelAppRunError, using the provided ModelDefaultError
func (t *ModelAppRunError) MergeModelDefaultError(v ModelDefaultError) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsModelCloudctrlError returns the union data inside the ModelAppRunError as a ModelCloudctrlError
func (t ModelAppRunError) AsModelCloudctrlError() (ModelCloudctrlError, error) {
	var body ModelCloudctrlError
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromModelCloudctrlError overwrites any union data inside the ModelAppRunError as the provided ModelCloudctrlError
func (t *ModelAppRunError) FromModelCloudctrlError(v ModelCloudctrlError) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeModelCloudctrlError performs a merge with any union data inside the ModelAppRunError, using the provided ModelCloudctrlError
func (t *ModelAppRunError) MergeModelCloudctrlError(v ModelCloudctrlError) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t ModelAppRunError) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ModelAppRunError) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}
