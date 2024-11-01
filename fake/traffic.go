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
	if _, ok := engine.Traffics[appId]; !ok {
		return nil, newError(
			ErrorTypeNotFound, "traffic", nil,
			"アプリケーションが見つかりませんでした。")
	}

	var ts []v1.Traffic
	for _, t := range engine.Traffics[appId] {
		ts = append(ts, *t)
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

	// 内部的にTrafficとApplicationのリレーションを保持する
	engine.Traffics[*app.Id] = append(engine.Traffics[*app.Id], &t)
}
