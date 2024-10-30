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

	v1 "github.com/sacloud/apprun-api-go/apis/v1"
)

func (engine *Engine) createVersion(app *v1.Application) error {
	versionId, err := newId()
	if err != nil {
		return err
	}
	name := fmt.Sprintf("version-%03d", engine.nextVersionId())
	createdAt := time.Now().UTC().Truncate(time.Second)

	v := v1.Version{
		Id:        &versionId,
		Name:      &name,
		Status:    (*v1.VersionStatus)(app.Status),
		CreatedAt: &createdAt,
	}
	engine.Versions = append(engine.Versions, &v)

	// 内部的にVersionとApplicationのリレーションを保持する
	engine.appVersionRelations[*app.Id] = append(engine.appVersionRelations[*app.Id], &appVersionRelation{
		application: app,
		version:     &v,
	})

	return nil
}
