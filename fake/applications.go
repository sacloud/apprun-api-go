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
	"time"

	"github.com/google/uuid"
	v1 "github.com/sacloud/apprun-api-go/apis/v1"
)

func (engine *Engine) CreateApplication(reqBody *v1.PostApplicationBody) (*v1.HandlerPostApplication, error) {
	defer engine.lock()()
	id, err := newId()
	if err != nil {
		return nil, newError(
			ErrorTypeUnknown, "application", nil,
			"IDの生成に失敗しました。")
	}

	var components []v1.HandlerPostApplicationComponent
	for _, reqComponent := range reqBody.Components {
		var cr v1.HandlerPostApplicationComponentDataSourceContainerRegistry
		if reqComponent.Datasource.ContainerRegistry != nil {
			cr.Image = reqComponent.Datasource.ContainerRegistry.Image
			cr.Server = reqComponent.Datasource.ContainerRegistry.Server
			cr.Username = reqComponent.Datasource.ContainerRegistry.Username
		}

		var env []v1.HandlerPostApplicationComponentEnv
		if reqComponent.Env != nil {
			for _, e := range *reqComponent.Env {
				env = append(env, (v1.HandlerPostApplicationComponentEnv)(e))
			}
		}

		var component v1.HandlerPostApplicationComponent
		component.Name = reqComponent.Name
		component.MaxCpu = string(reqComponent.MaxCpu)
		component.MaxMemory = string(reqComponent.MaxMemory)
		component.Datasource.ContainerRegistry = &cr
		component.Env = &env
		component.Probe = (*v1.HandlerPostApplicationComponentProbe)(reqComponent.Probe)
		components = append(components, component)
	}

	status := v1.HandlerPostApplicationStatusSuccess
	url := fmt.Sprintf("https://example.com/apprun/dummy/%s", id)
	createdAt := time.Now().UTC().Truncate(time.Second)
	engine.Application = &v1.HandlerPostApplication{
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
	return engine.Application, nil
}

func newId() (string, error) {
	id, err := uuid.NewRandom()
	return id.String(), err
}
