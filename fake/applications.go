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

package fake

import (
	"fmt"
	"sort"
	"time"

	"github.com/google/uuid"
	v1 "github.com/sacloud/apprun-api-go/apis/v1"
)

func (engine *Engine) ReadApplication(id string) (*v1.Application, error) {
	defer engine.rLock()()

	if len(engine.Applications) == 0 {
		return nil, newError(
			ErrorTypeNotFound, "application", nil,
			"アプリケーションが見つかりませんでした。")
	}

	for _, app := range engine.Applications {
		if app != nil && app.Id != nil && *app.Id == id {
			return app, nil
		}
	}

	return nil, newError(
		ErrorTypeNotFound, "application", nil,
		"アプリケーションが見つかりませんでした。")
}

// sort_field, sort_orderは無視する
func (engine *Engine) ListApplications(param v1.ListApplicationsParams) (*v1.HandlerListApplications, error) {
	defer engine.rLock()()

	apps := engine.Applications
	appsLen := len(apps)
	if appsLen == 0 {
		return nil, newError(
			ErrorTypeNotFound, "application", nil,
			"アプリケーションが見つかりませんでした。")
	}

	sort.Slice(apps, func(i int, j int) bool {
		switch *param.SortField {
		case "created_at":
			if *param.SortOrder == "desc" {
				return apps[i].CreatedAt.After(*apps[j].CreatedAt)
			}
			return apps[i].CreatedAt.Before(*apps[j].CreatedAt)
		// sort_fieldのデフォルト値であるcreated_at以外は未サポート
		default:
			return false
		}
	})

	start := (*param.PageNum - 1) * *param.PageSize
	// 範囲外の場合nilを返す
	if start > len(apps) {
		return nil, nil
	}

	end := start + *param.PageSize
	if end > len(apps) {
		end = len(apps)
	}

	var data []v1.HandlerListApplicationsData
	for _, app := range apps[start:end] {
		if app != nil {
			d := v1.HandlerListApplicationsData{
				Id:        app.Id,
				Name:      app.Name,
				Status:    (*v1.HandlerListApplicationsDataStatus)(app.Status),
				PublicUrl: app.PublicUrl,
				CreatedAt: app.CreatedAt,
			}

			data = append(data, d)
		}
	}

	meta := v1.HandlerListApplicationsMeta{
		PageNum:     param.PageNum,
		PageSize:    param.PageSize,
		ObjectTotal: &appsLen,
		SortField:   param.SortField,
		SortOrder:   (*v1.HandlerListApplicationsMetaSortOrder)(param.SortOrder),
	}

	return &v1.HandlerListApplications{
		Data: &data,
		Meta: &meta,
	}, nil
}

func (engine *Engine) DeleteApplication(id string) error {
	defer engine.lock()()

	var idx int
	for i, app := range engine.Applications {
		if *app.Id == id {
			idx = i
			break
		}
	}

	engine.Applications[idx] = engine.Applications[len(engine.Applications)-1]
	engine.Applications = engine.Applications[:len(engine.Applications)-1]

	return nil
}

func (engine *Engine) CreateApplication(reqBody *v1.PostApplicationBody) (*v1.Application, error) {
	defer engine.lock()()
	id, err := newId()
	if err != nil {
		return nil, newError(
			ErrorTypeUnknown, "application", nil,
			"IDの生成に失敗しました。")
	}

	var components []v1.HandlerApplicationComponent
	for _, reqComponent := range reqBody.Components {
		var cr v1.HandlerApplicationComponentDataSourceContainerRegistry
		if reqComponent.Datasource.ContainerRegistry != nil {
			cr.Image = reqComponent.Datasource.ContainerRegistry.Image
			cr.Server = reqComponent.Datasource.ContainerRegistry.Server
			cr.Username = reqComponent.Datasource.ContainerRegistry.Username
		}

		var env []v1.HandlerApplicationComponentEnv
		if reqComponent.Env != nil {
			for _, e := range *reqComponent.Env {
				env = append(env, (v1.HandlerApplicationComponentEnv)(e))
			}
		}

		var component v1.HandlerApplicationComponent
		component.Name = reqComponent.Name
		component.MaxCpu = string(reqComponent.MaxCpu)
		component.MaxMemory = string(reqComponent.MaxMemory)
		component.Datasource.ContainerRegistry = &cr
		component.Env = &env
		component.Probe = (*v1.HandlerApplicationComponentProbe)(reqComponent.Probe)
		components = append(components, component)
	}

	status := v1.ApplicationStatusSuccess
	url := fmt.Sprintf("https://example.com/apprun/dummy/%s", id)
	createdAt := time.Now().UTC().Truncate(time.Second)
	app := &v1.Application{
		Id:             &id,
		Name:           &reqBody.Name,
		TimeoutSeconds: &reqBody.TimeoutSeconds,
		Port:           &reqBody.Port,
		MinScale:       &reqBody.MinScale,
		MaxScale:       &reqBody.MaxScale,
		Components:     &components,
		Status:         &status,
		PublicUrl:      &url,
		CreatedAt:      &createdAt,
	}
	engine.Applications = append(engine.Applications, app)

	return app, nil
}

func (engine *Engine) PatchApplication(id string, reqBody *v1.PatchApplicationBody) (*v1.HandlerPatchApplication, error) {
	defer engine.lock()()
	var patchedApp v1.HandlerPatchApplication
	for _, app := range engine.Applications {
		if *app.Id == id {
			if reqBody.Name != nil {
				app.Name = reqBody.Name
			}
			if reqBody.TimeoutSeconds != nil {
				app.TimeoutSeconds = reqBody.TimeoutSeconds
			}
			if reqBody.Port != nil {
				app.Port = reqBody.Port
			}
			if reqBody.MinScale != nil {
				app.MinScale = reqBody.MinScale
			}
			if reqBody.MaxScale != nil {
				app.MaxScale = reqBody.MaxScale
			}
			if reqBody.Components != nil {
				var components []v1.HandlerApplicationComponent
				for _, reqComponent := range *reqBody.Components {
					var cr v1.HandlerApplicationComponentDataSourceContainerRegistry
					if reqComponent.Datasource.ContainerRegistry != nil {
						cr.Image = reqComponent.Datasource.ContainerRegistry.Image
						cr.Server = reqComponent.Datasource.ContainerRegistry.Server
						cr.Username = reqComponent.Datasource.ContainerRegistry.Username
					}

					var env []v1.HandlerApplicationComponentEnv
					if reqComponent.Env != nil {
						for _, e := range *reqComponent.Env {
							env = append(env, (v1.HandlerApplicationComponentEnv)(e))
						}
					}

					var component v1.HandlerApplicationComponent
					component.Name = reqComponent.Name
					component.MaxCpu = string(reqComponent.MaxCpu)
					component.MaxMemory = string(reqComponent.MaxMemory)
					component.Datasource.ContainerRegistry = &cr
					component.Env = &env
					component.Probe = (*v1.HandlerApplicationComponentProbe)(reqComponent.Probe)
					components = append(components, component)
				}

				if len(components) > 0 {
					app.Components = &components
				}
			}
		}

		updatedAt := time.Now().UTC().Truncate(time.Second)
		patchedApp = v1.HandlerPatchApplication{
			Name:           app.Name,
			TimeoutSeconds: app.TimeoutSeconds,
			Port:           app.Port,
			MinScale:       app.MinScale,
			MaxScale:       app.MaxScale,
			Components:     app.Components,
			Status:         (*v1.HandlerPatchApplicationStatus)(app.Status),
			PublicUrl:      app.PublicUrl,
			UpdatedAt:      &updatedAt,
		}
	}

	return &patchedApp, nil
}

func newId() (string, error) {
	id, err := uuid.NewRandom()
	return id.String(), err
}
