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
	"time"
)

// Defines values for ApplicationStatus.
const (
	ApplicationStatusFail    ApplicationStatus = "Fail"
	ApplicationStatusSuccess ApplicationStatus = "Success"
	ApplicationStatusUnknown ApplicationStatus = "Unknown"
)

// Defines values for HandlerGetApplicationStatusStatus.
const (
	HandlerGetApplicationStatusStatusFail    HandlerGetApplicationStatusStatus = "Fail"
	HandlerGetApplicationStatusStatusSuccess HandlerGetApplicationStatusStatus = "Success"
	HandlerGetApplicationStatusStatusUnknown HandlerGetApplicationStatusStatus = "Unknown"
)

// Defines values for HandlerGetVersionStatus.
const (
	HandlerGetVersionStatusFail    HandlerGetVersionStatus = "Fail"
	HandlerGetVersionStatusSuccess HandlerGetVersionStatus = "Success"
	HandlerGetVersionStatusUnknown HandlerGetVersionStatus = "Unknown"
)

// Defines values for HandlerListApplicationsDataStatus.
const (
	HandlerListApplicationsDataStatusFail    HandlerListApplicationsDataStatus = "Fail"
	HandlerListApplicationsDataStatusSuccess HandlerListApplicationsDataStatus = "Success"
	HandlerListApplicationsDataStatusUnknown HandlerListApplicationsDataStatus = "Unknown"
)

// Defines values for HandlerListApplicationsMetaSortOrder.
const (
	HandlerListApplicationsMetaSortOrderAsc  HandlerListApplicationsMetaSortOrder = "asc"
	HandlerListApplicationsMetaSortOrderDesc HandlerListApplicationsMetaSortOrder = "desc"
)

// Defines values for HandlerListVersionsDataStatus.
const (
	HandlerListVersionsDataStatusFail    HandlerListVersionsDataStatus = "Fail"
	HandlerListVersionsDataStatusSuccess HandlerListVersionsDataStatus = "Success"
	HandlerListVersionsDataStatusUnknown HandlerListVersionsDataStatus = "Unknown"
)

// Defines values for HandlerListVersionsMetaSortOrder.
const (
	HandlerListVersionsMetaSortOrderAsc  HandlerListVersionsMetaSortOrder = "asc"
	HandlerListVersionsMetaSortOrderDesc HandlerListVersionsMetaSortOrder = "desc"
)

// Defines values for HandlerPatchApplicationStatus.
const (
	HandlerPatchApplicationStatusFail    HandlerPatchApplicationStatus = "Fail"
	HandlerPatchApplicationStatusSuccess HandlerPatchApplicationStatus = "Success"
	HandlerPatchApplicationStatusUnknown HandlerPatchApplicationStatus = "Unknown"
)

// Defines values for ModelErrorsLocationType.
const (
	ModelErrorsLocationTypeBody      ModelErrorsLocationType = "body"
	ModelErrorsLocationTypeHeader    ModelErrorsLocationType = "header"
	ModelErrorsLocationTypeParameter ModelErrorsLocationType = "parameter"
	ModelErrorsLocationTypeQuery     ModelErrorsLocationType = "query"
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
	Components *[]HandlerApplicationComponent `json:"components,omitempty"`

	// CreatedAt 作成日時
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// Id アプリケーションID
	Id *string `json:"id,omitempty"`

	// MaxScale アプリケーション全体の最大スケール数
	MaxScale *int `json:"max_scale,omitempty"`

	// MinScale アプリケーション全体の最小スケール数
	MinScale *int `json:"min_scale,omitempty"`

	// Name アプリケーション名
	Name *string `json:"name,omitempty"`

	// Port アプリケーションがリクエストを待ち受けるポート番号
	Port *int `json:"port,omitempty"`

	// PublicUrl 公開URL
	PublicUrl *string `json:"public_url,omitempty"`

	// Status アプリケーションステータス
	Status *ApplicationStatus `json:"status,omitempty"`

	// TimeoutSeconds アプリケーションの公開URLにアクセスして、インスタンスが起動してからレスポンスが返るまでの時間制限
	TimeoutSeconds *int `json:"timeout_seconds,omitempty"`
}

// ApplicationStatus アプリケーションステータス
type ApplicationStatus string

// HandlerApplicationComponent defines model for handler.ApplicationComponent.
type HandlerApplicationComponent struct {
	Datasource HandlerApplicationComponentDataSource `json:"datasource"`

	// Env コンポーネントに渡す環境変数
	Env *[]HandlerApplicationComponentEnv `json:"env,omitempty"`

	// MaxCpu コンポーネントの最大CPU数
	MaxCpu string `json:"max_cpu"`

	// MaxMemory コンポーネントの最大メモリ
	MaxMemory string `json:"max_memory"`

	// Name コンポーネント名
	Name  string                            `json:"name"`
	Probe *HandlerApplicationComponentProbe `json:"probe"`
}

// HandlerApplicationComponentDataSource defines model for handler.ApplicationComponentDataSource.
type HandlerApplicationComponentDataSource struct {
	ContainerRegistry *HandlerApplicationComponentDataSourceContainerRegistry `json:"container_registry,omitempty"`
}

// HandlerApplicationComponentDataSourceContainerRegistry defines model for handler.ApplicationComponentDataSourceContainerRegistry.
type HandlerApplicationComponentDataSourceContainerRegistry struct {
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

// HandlerApplicationComponentProbe defines model for handler.ApplicationComponentProbe.
type HandlerApplicationComponentProbe struct {
	// HttpGet HTTP Getプローブタイプ
	HttpGet *struct {
		Headers *[]struct {
			// Name ヘッダーフィールド名
			Name *string `json:"name,omitempty"`

			// Value ヘッダーフィールド値
			Value *string `json:"value,omitempty"`
		} `json:"headers,omitempty"`

		// Path HTTPサーバーへアクセスしプローブをチェックする際のパス
		Path string `json:"path"`

		// Port HTTPサーバーへアクセスしプローブをチェックする際のポート番号
		Port int `json:"port"`
	} `json:"http_get"`
}

// HandlerGetApplication defines model for handler.getApplication.
type HandlerGetApplication = Application

// HandlerGetApplicationStatusResponse defines model for handler.getApplicationStatus.
type HandlerGetApplicationStatusResponse struct {
	// Message ステータス失敗時のメッセージ
	Message *string `json:"message,omitempty"`

	// Status アプリケーションステータス
	Status *HandlerGetApplicationStatusStatus `json:"status,omitempty"`
}

// HandlerGetApplicationStatusStatus アプリケーションステータス
type HandlerGetApplicationStatusStatus string

// HandlerGetVersion defines model for handler.getVersion.
type HandlerGetVersion struct {
	// Components バージョンのコンポーネント情報
	Components *[]struct {
		// Datasource コンポーネントを構成するソース
		Datasource struct {
			// ContainerRegistry コンテナレジストリ
			ContainerRegistry *struct {
				// Image コンテナイメージ名
				Image string `json:"image"`

				// Server コンテナレジストリのサーバー名
				Server *string `json:"server,omitempty"`

				// Username コンテナレジストリの認証情報
				Username *string `json:"username,omitempty"`
			} `json:"container_registry,omitempty"`
		} `json:"datasource"`

		// Env コンポーネントに渡す環境変数
		Env *[]struct {
			// Key 環境変数名
			Key *string `json:"key,omitempty"`

			// Value 環境変数の値
			Value *string `json:"value,omitempty"`
		} `json:"env,omitempty"`

		// MaxCpu コンポーネントの最大CPU数
		MaxCpu string `json:"max_cpu"`

		// MaxMemory コンポーネントの最大メモリ
		MaxMemory string `json:"max_memory"`

		// Name コンポーネント名
		Name string `json:"name"`

		// Probe コンポーネントのプローブ設定
		Probe *struct {
			// HttpGet HTTP Getプローブタイプ
			HttpGet *struct {
				Headers *[]struct {
					// Name ヘッダーフィールド名
					Name *string `json:"name,omitempty"`

					// Value ヘッダーフィールド値
					Value *string `json:"value,omitempty"`
				} `json:"headers,omitempty"`

				// Path HTTPサーバーへアクセスしプローブをチェックする際のパス
				Path string `json:"path"`

				// Port HTTPサーバーへアクセスしプローブをチェックする際のポート番号
				Port int `json:"port"`
			} `json:"http_get"`
		} `json:"probe"`
	} `json:"components,omitempty"`

	// CreatedAt 作成日時
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// Id バージョンID
	Id *string `json:"id,omitempty"`

	// MaxScale バージョンの最大スケール数
	MaxScale *int `json:"max_scale,omitempty"`

	// MinScale バージョンの最小スケール数
	MinScale *int `json:"min_scale,omitempty"`

	// Name バージョン名
	Name *string `json:"name,omitempty"`

	// Port アプリケーションがリクエストを待ち受けるポート番号
	Port *int `json:"port,omitempty"`

	// Status バージョンステータス
	Status *HandlerGetVersionStatus `json:"status,omitempty"`

	// TimeoutSeconds アプリケーションの公開URLにアクセスして、インスタンスが起動してからレスポンスが返るまでの時間制限
	TimeoutSeconds *int `json:"timeout_seconds,omitempty"`
}

// HandlerGetVersionStatus バージョンステータス
type HandlerGetVersionStatus string

// HandlerListApplications defines model for handler.listApplications.
type HandlerListApplications struct {
	Data *[]HandlerListApplicationsData `json:"data,omitempty"`
	Meta *HandlerListApplicationsMeta   `json:"meta,omitempty"`
}

// HandlerListApplicationsData defines model for handler.listApplicationsData.
type HandlerListApplicationsData struct {
	// CreatedAt 作成日時
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// Id アプリケーションID
	Id *string `json:"id,omitempty"`

	// Name アプリケーション名
	Name *string `json:"name,omitempty"`

	// PublicUrl 公開URL
	PublicUrl *string `json:"public_url,omitempty"`

	// Status アプリケーションステータス
	Status *HandlerListApplicationsDataStatus `json:"status,omitempty"`
}

// HandlerListApplicationsDataStatus アプリケーションステータス
type HandlerListApplicationsDataStatus string

// HandlerListApplicationsMeta defines model for handler.listApplicationsMeta.
type HandlerListApplicationsMeta struct {
	ObjectTotal *int                                  `json:"object_total,omitempty"`
	PageNum     *int                                  `json:"page_num,omitempty"`
	PageSize    *int                                  `json:"page_size,omitempty"`
	SortField   *string                               `json:"sort_field,omitempty"`
	SortOrder   *HandlerListApplicationsMetaSortOrder `json:"sort_order,omitempty"`
}

// HandlerListApplicationsMetaSortOrder defines model for HandlerListApplicationsMeta.SortOrder.
type HandlerListApplicationsMetaSortOrder string

// HandlerListTraffics defines model for handler.listTraffics.
type HandlerListTraffics struct {
	Data *[]struct {
		// IsLatestVersion 最新バージョンかどうか
		IsLatestVersion *bool `json:"is_latest_version,omitempty"`

		// Percent トラフィック分散の割合
		Percent *int `json:"percent,omitempty"`

		// VersionName バージョン名
		VersionName *string `json:"version_name,omitempty"`
	} `json:"data,omitempty"`
	Meta *map[string]interface{} `json:"meta"`
}

// HandlerListVersions defines model for handler.listVersions.
type HandlerListVersions struct {
	Data *[]struct {
		// CreatedAt 作成日時
		CreatedAt *time.Time `json:"created_at,omitempty"`

		// Id バージョンID
		Id *string `json:"id,omitempty"`

		// Name バージョン名
		Name *string `json:"name,omitempty"`

		// Status ステータス
		Status *HandlerListVersionsDataStatus `json:"status,omitempty"`
	} `json:"data,omitempty"`
	Meta *struct {
		ObjectTotal *int                              `json:"object_total,omitempty"`
		PageNum     *int                              `json:"page_num,omitempty"`
		PageSize    *int                              `json:"page_size,omitempty"`
		SortField   *string                           `json:"sort_field,omitempty"`
		SortOrder   *HandlerListVersionsMetaSortOrder `json:"sort_order,omitempty"`
	} `json:"meta,omitempty"`
}

// HandlerListVersionsDataStatus ステータス
type HandlerListVersionsDataStatus string

// HandlerListVersionsMetaSortOrder defines model for HandlerListVersions.Meta.SortOrder.
type HandlerListVersionsMetaSortOrder string

// HandlerPatchApplication defines model for handler.patchApplication.
type HandlerPatchApplication struct {
	// Components アプリケーションのコンポーネント情報
	Components *[]HandlerApplicationComponent `json:"components,omitempty"`

	// Id アプリケーションID
	Id *string `json:"id,omitempty"`

	// MaxScale アプリケーション全体の最大スケール数
	MaxScale *int `json:"max_scale,omitempty"`

	// MinScale アプリケーション全体の最小スケール数
	MinScale *int `json:"min_scale,omitempty"`

	// Name アプリケーション名
	Name *string `json:"name,omitempty"`

	// Port アプリケーションがリクエストを待ち受けるポート番号
	Port *int `json:"port,omitempty"`

	// PublicUrl 公開URL
	PublicUrl *string `json:"public_url,omitempty"`

	// Status アプリケーションステータス
	Status *HandlerPatchApplicationStatus `json:"status,omitempty"`

	// TimeoutSeconds アプリケーションの公開URLにアクセスして、インスタンスが起動してからレスポンスが返るまでの時間制限
	TimeoutSeconds *int `json:"timeout_seconds,omitempty"`

	// UpdatedAt 更新日時
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// HandlerPatchApplicationStatus アプリケーションステータス
type HandlerPatchApplicationStatus string

// HandlerPostApplication defines model for handler.postApplication.
type HandlerPostApplication = Application

// HandlerPutTraffics defines model for handler.putTraffics.
type HandlerPutTraffics struct {
	Data *[]struct {
		// IsLatestVersion 最新バージョンかどうか
		IsLatestVersion *bool `json:"is_latest_version,omitempty"`

		// Percent トラフィック分散の割合
		Percent *int `json:"percent,omitempty"`

		// VersionName バージョン名
		VersionName *string `json:"version_name,omitempty"`
	} `json:"data,omitempty"`
	Meta *map[string]interface{} `json:"meta"`
}

// ModelDefaultError defines model for model.defaultError.
type ModelDefaultError struct {
	Error *struct {
		Code    *float32     `json:"code,omitempty"`
		Errors  *ModelErrors `json:"errors,omitempty"`
		Message *string      `json:"message,omitempty"`
	} `json:"error,omitempty"`
}

// ModelErrors defines model for model.errors.
type ModelErrors = []struct {
	Domain       *string                  `json:"domain"`
	Location     *string                  `json:"location"`
	LocationType *ModelErrorsLocationType `json:"location_type"`
	Message      *string                  `json:"message"`
	Reason       *string                  `json:"reason"`
}

// ModelErrorsLocationType defines model for ModelErrors.LocationType.
type ModelErrorsLocationType string

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

	// Name アプリケーション名
	Name *string `json:"name,omitempty"`

	// Port アプリケーションがリクエストを待ち受けるポート番号
	Port *int `json:"port,omitempty"`

	// TimeoutSeconds アプリケーションの公開URLにアクセスして、インスタンスが起動してからレスポンスが返るまでの時間制限
	TimeoutSeconds *int `json:"timeout_seconds,omitempty"`
}

// PatchApplicationBodyComponent defines model for patchApplicationBodyComponent.
type PatchApplicationBodyComponent struct {
	Datasource PatchApplicationBodyComponentDataSource `json:"datasource"`
	Env        *PatchApplicationBodyComponentEnv       `json:"env"`

	// MaxCpu コンポーネントの最大CPU数
	MaxCpu PatchApplicationBodyComponentMaxCpu `json:"max_cpu"`

	// MaxMemory コンポーネントの最大メモリ
	MaxMemory PatchApplicationBodyComponentMaxMemory `json:"max_memory"`

	// Name コンポーネント名
	Name  string                              `json:"name"`
	Probe *PatchApplicationBodyComponentProbe `json:"probe"`
}

// PatchApplicationBodyComponentMaxCpu コンポーネントの最大CPU数
type PatchApplicationBodyComponentMaxCpu string

// PatchApplicationBodyComponentMaxMemory コンポーネントの最大メモリ
type PatchApplicationBodyComponentMaxMemory string

// PatchApplicationBodyComponentDataSource defines model for patchApplicationBodyComponentDataSource.
type PatchApplicationBodyComponentDataSource struct {
	ContainerRegistry *PatchApplicationBodyComponentDataSourceContainerRegistry `json:"container_registry,omitempty"`
}

// PatchApplicationBodyComponentDataSourceContainerRegistry defines model for patchApplicationBodyComponentDataSourceContainerRegistry.
type PatchApplicationBodyComponentDataSourceContainerRegistry struct {
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
type PatchApplicationBodyComponentEnv = []struct {
	// Key 環境変数名
	Key *string `json:"key,omitempty"`

	// Value 環境変数の値
	Value *string `json:"value,omitempty"`
}

// PatchApplicationBodyComponentProbe defines model for patchApplicationBodyComponentProbe.
type PatchApplicationBodyComponentProbe struct {
	// HttpGet HTTP Getプローブタイプ
	HttpGet *struct {
		Headers *[]struct {
			// Name ヘッダーフィールド名
			Name *string `json:"name,omitempty"`

			// Value ヘッダーフィールド値
			Value *string `json:"value,omitempty"`
		} `json:"headers,omitempty"`

		// Path HTTPサーバーへアクセスしプローブをチェックする際のパス
		Path string `json:"path"`

		// Port HTTPサーバーへアクセスしプローブをチェックする際のポート番号
		Port int `json:"port"`
	} `json:"http_get"`
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
	Datasource PostApplicationBodyComponentDataSource `json:"datasource"`

	// Env コンポーネントに渡す環境変数
	Env *[]PostApplicationBodyComponentEnv `json:"env"`

	// MaxCpu コンポーネントの最大CPU数
	MaxCpu PostApplicationBodyComponentMaxCpu `json:"max_cpu"`

	// MaxMemory コンポーネントの最大メモリ
	MaxMemory PostApplicationBodyComponentMaxMemory `json:"max_memory"`

	// Name コンポーネント名
	Name  string                             `json:"name"`
	Probe *PostApplicationBodyComponentProbe `json:"probe"`
}

// PostApplicationBodyComponentMaxCpu コンポーネントの最大CPU数
type PostApplicationBodyComponentMaxCpu string

// PostApplicationBodyComponentMaxMemory コンポーネントの最大メモリ
type PostApplicationBodyComponentMaxMemory string

// PostApplicationBodyComponentDataSource defines model for postApplicationBodyComponentDataSource.
type PostApplicationBodyComponentDataSource struct {
	ContainerRegistry *PostApplicationBodyComponentDataSourceContainerRegistry `json:"container_registry,omitempty"`
}

// PostApplicationBodyComponentDataSourceContainerRegistry defines model for postApplicationBodyComponentDataSourceContainerRegistry.
type PostApplicationBodyComponentDataSourceContainerRegistry struct {
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

// PostApplicationBodyComponentProbe defines model for postApplicationBodyComponentProbe.
type PostApplicationBodyComponentProbe struct {
	// HttpGet HTTP Getプローブタイプ
	HttpGet *struct {
		Headers *[]struct {
			// Name ヘッダーフィールド名
			Name *string `json:"name,omitempty"`

			// Value ヘッダーフィールド値
			Value *string `json:"value,omitempty"`
		} `json:"headers,omitempty"`

		// Path HTTPサーバーへアクセスしプローブをチェックする際のパス
		Path string `json:"path"`

		// Port HTTPサーバーへアクセスしプローブをチェックする際のポート番号
		Port int `json:"port"`
	} `json:"http_get"`
}

// PutTrafficsBody defines model for putTrafficsBody.
type PutTrafficsBody = []struct {
	// IsLatestVersion 最新バージョンかどうか
	IsLatestVersion *bool `json:"is_latest_version,omitempty"`

	// Percent トラフィック分散の割合
	Percent *int `json:"percent,omitempty"`

	// VersionName バージョン名
	VersionName *string `json:"version_name,omitempty"`
}

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
