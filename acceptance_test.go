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

//go:build acctest
// +build acctest

package apprun_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/sacloud/apprun-api-go"
	v1 "github.com/sacloud/apprun-api-go/apis/v1"
	"github.com/stretchr/testify/require"
)

const appName = "app-for-acceptance"

// TestUserAPI ユーザー関連APIの操作テスト
// ユーザーは削除もできないため、2回目以降は既にユーザーが存在する状態でのテストとなってしまうことに注意する。
func TestUserAPI(t *testing.T) {
	skipIfNoAPIKey(t)

	ctx := context.Background()
	userOp := apprun.NewUserOp(client)

	// Create
	_, err := userOp.Create(ctx)
	require.NoError(t, err)

	// Read
	res, err := userOp.Read(ctx)
	require.NoError(t, err)
	require.Equal(t, res.StatusCode, 200)
}

// TestApplicationAPI アプリケーションの一連の操作テスト
// 以下のシナリオでテストを行う
//   - アプリケーションを作成
//   - アプリケーションの一覧を取得
//   - アプリケーションを更新
//   - アプリケーションが更新できたかどうかを確認
//   - アプリケーションのステータスを取得
//   - アプリケーションを削除
func TestApplicationAPI(t *testing.T) {
	skipIfNoAPIKey(t)

	if err := cleanupTestApplication(); err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	appOp := apprun.NewApplicationOp(client)

	// Create
	application, err := appOp.Create(ctx, &v1.PostApplicationBody{
		Name:           appName,
		TimeoutSeconds: 100,
		Port:           80,
		MinScale:       0,
		MaxScale:       1,
		Components: []v1.PostApplicationBodyComponent{
			{
				Name:      "component1",
				MaxCpu:    "0.1",
				MaxMemory: "256Mi",
				DeploySource: v1.PostApplicationBodyComponentDeploySource{
					ContainerRegistry: &v1.PostApplicationBodyComponentDeploySourceContainerRegistry{
						Image: "apprun-test.sakuracr.jp/apprun/test1:latest",
					},
				},
				Probe: &v1.PostApplicationBodyComponentProbe{
					HttpGet: &v1.PostApplicationBodyComponentProbeHttpGet{
						Path: "/",
						Port: 80,
					},
				},
			},
		},
	})
	require.NoError(t, err)

	// Read
	application, err = appOp.Read(ctx, application.Id)
	require.NoError(t, err)
	require.Equal(t, application.Name, appName)

	// Update
	timeoutUpdated := 20
	appOp.Update(ctx, application.Id, &v1.PatchApplicationBody{
		TimeoutSeconds: &timeoutUpdated,
	})

	// Read
	application, err = appOp.Read(ctx, application.Id)
	require.NoError(t, err)
	require.Equal(t, application.TimeoutSeconds, timeoutUpdated)

	// Read Status
	// ヘルスチェックが完了するまでタイムラグがあるため暫く待つ
	time.Sleep(30 * time.Second)

	res, err := appOp.ReadStatus(ctx, application.Id)
	require.NoError(t, err)
	require.Equal(t, res.Status, v1.HandlerGetApplicationStatusStatusHealthy)

	// Delete
	err = appOp.Delete(ctx, application.Id)
	require.NoError(t, err)
}

// TestVersionAPI アプリケーションバージョンの一連の操作テスト
// 以下のシナリオでテストを行う
//   - アプリケーションを作成
//   - アプリケーションを更新
//   - アプリケーションバージョンの一覧を取得
//   - アプリケーションバージョンを削除
//   - アプリケーションバージョンを確認し、削除できていることを確認
//   - アプリケーションを削除
func TestVersionAPI(t *testing.T) {
	skipIfNoAPIKey(t)

	if err := cleanupTestApplication(); err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	appOp := apprun.NewApplicationOp(client)
	versionOp := apprun.NewVersionOp(client)

	// Application Create
	application, _ := appOp.Create(ctx, &v1.PostApplicationBody{
		Name:           appName,
		TimeoutSeconds: 100,
		Port:           80,
		MinScale:       0,
		MaxScale:       1,
		Components: []v1.PostApplicationBodyComponent{
			{
				Name:      "component1",
				MaxCpu:    "0.1",
				MaxMemory: "256Mi",

				DeploySource: v1.PostApplicationBodyComponentDeploySource{
					ContainerRegistry: &v1.PostApplicationBodyComponentDeploySourceContainerRegistry{
						Image: "apprun-test.sakuracr.jp/apprun/test1:latest",
					},
				},
				Probe: &v1.PostApplicationBodyComponentProbe{
					HttpGet: &v1.PostApplicationBodyComponentProbeHttpGet{
						Path: "/",
						Port: 80,
					},
				},
			},
		},
	})

	// Update Application
	timeoutSeconds := 10
	appOp.Update(ctx, application.Id, &v1.PatchApplicationBody{
		TimeoutSeconds: &timeoutSeconds,
	})

	// List Version
	versions, err := versionOp.List(ctx, application.Id, &v1.ListApplicationVersionsParams{})
	require.NoError(t, err)
	require.Equal(t, len(versions.Data), 2)

	// Delete Version
	err = versionOp.Delete(ctx, application.Id, versions.Data[1].Id)
	require.NoError(t, err)

	// List Version
	versions, err = versionOp.List(ctx, application.Id, &v1.ListApplicationVersionsParams{})
	require.NoError(t, err)
	require.Equal(t, len(versions.Data), 1)

	// Delete Application
	appOp.Delete(ctx, application.Id)
}

// TestTrafficAPI アプリケーショントラフィックの一連の操作テスト
// 以下のシナリオでテストを行う
//   - アプリケーションを作成
//   - アプリケーションを更新
//   - アプリケーショントラフィックを変更
//   - アプリケーショントラフィックを確認
//   - アプリケーションを削除
func TestTrafficAPI(t *testing.T) {
	skipIfNoAPIKey(t)

	if err := cleanupTestApplication(); err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	appOp := apprun.NewApplicationOp(client)
	versionOp := apprun.NewVersionOp(client)
	trafficOp := apprun.NewTrafficOp(client)

	// Application Create
	application, _ := appOp.Create(ctx, &v1.PostApplicationBody{
		Name:           appName,
		TimeoutSeconds: 100,
		Port:           80,
		MinScale:       0,
		MaxScale:       1,
		Components: []v1.PostApplicationBodyComponent{
			{
				Name:      "component1",
				MaxCpu:    "0.1",
				MaxMemory: "256Mi",
				DeploySource: v1.PostApplicationBodyComponentDeploySource{
					ContainerRegistry: &v1.PostApplicationBodyComponentDeploySourceContainerRegistry{
						Image: "apprun-test.sakuracr.jp/apprun/test1:latest",
					},
				},
				Probe: &v1.PostApplicationBodyComponentProbe{
					HttpGet: &v1.PostApplicationBodyComponentProbeHttpGet{
						Path: "/",
						Port: 80,
					},
				},
			},
		},
	})

	// Update Application
	timeoutSeconds := 10
	appOp.Update(ctx, application.Id, &v1.PatchApplicationBody{
		TimeoutSeconds: &timeoutSeconds,
	})

	// Update Application Traffic
	versions, _ := versionOp.List(ctx, application.Id, &v1.ListApplicationVersionsParams{})

	v0IsLatestVersion := true
	v0Percent := 90

	v1Name := versions.Data[1].Name
	v1IsLatestVersion := false
	v1Percent := 10

	_, err := trafficOp.Update(ctx, application.Id, &[]v1.Traffic{
		{
			IsLatestVersion: v0IsLatestVersion,
			Percent:         v0Percent,
		},
		{
			VersionName:     v1Name,
			IsLatestVersion: v1IsLatestVersion,
			Percent:         v1Percent,
		},
	})
	require.NoError(t, err)

	// List Application Traffic
	blankName := ""
	traffics, err := trafficOp.List(ctx, application.Id)
	require.NoError(t, err)
	require.Equal(t, traffics.Data, []v1.Traffic{
		{
			VersionName:     blankName,
			IsLatestVersion: v0IsLatestVersion,
			Percent:         v0Percent,
		},
		{
			VersionName:     v1Name,
			IsLatestVersion: v1IsLatestVersion,
			Percent:         v1Percent,
		},
	})

	// Delete Application
	appOp.Delete(ctx, application.Id)
}

var client = &apprun.Client{}

// skipIfNoEnv 指定の環境変数のいずれかが空の場合はt.SkipNow()する
func skipIfNoEnv(t *testing.T, envs ...string) {
	var emptyEnvs []string
	for _, env := range envs {
		if os.Getenv(env) == "" {
			emptyEnvs = append(emptyEnvs, env)
		}
	}
	if len(emptyEnvs) > 0 {
		for _, env := range emptyEnvs {
			t.Logf("environment variable %q is not set", env)
		}
		t.SkipNow()
	}
}

func skipIfNoAPIKey(t *testing.T) {
	skipIfNoEnv(t, "SAKURACLOUD_ACCESS_TOKEN", "SAKURACLOUD_ACCESS_TOKEN_SECRET")
}

func cleanupTestApplication() error {
	ctx := context.Background()
	appOp := apprun.NewApplicationOp(client)

	apps, err := appOp.List(ctx, &v1.ListApplicationsParams{})
	if err != nil {
		return err
	}
	if apps.Data == nil {
		return nil
	}

	for _, app := range apps.Data {
		if app.Name == appName {
			if err := appOp.Delete(ctx, app.Id); err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}
