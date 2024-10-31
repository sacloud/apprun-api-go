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

import v1 "github.com/sacloud/apprun-api-go/apis/v1"

func (engine *Engine) ListTraffics(appId string) (*v1.HandlerListTraffics, error) {
	var ts []v1.Traffic
	if rs, ok := engine.appTrafficRelations[appId]; ok {
		for _, r := range rs {
			ts = append(ts, *r.traffic)
		}
	} else {
		return nil, newError(
			ErrorTypeNotFound, "traffic", nil,
			"アプリケーションが見つかりませんでした。")
	}

	return &v1.HandlerListTraffics{
		Meta: nil,
		Data: &ts,
	}, nil
}

func (engine *Engine) initTraffic(app *v1.Application) {
	isLatestVersion := true
	percent := 100
	versionName := ""
	t := v1.Traffic{
		IsLatestVersion: &isLatestVersion,
		Percent:         &percent,
		VersionName:     &versionName,
	}
	engine.Traffics = append(engine.Traffics, &t)

	// 内部的にTrafficとApplicationのリレーションを保持する
	engine.appTrafficRelations[*app.Id] = append(engine.appTrafficRelations[*app.Id],
		&appTrafficRelation{
			application: app,
			traffic:     &t,
		},
	)
}
