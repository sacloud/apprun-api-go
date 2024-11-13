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
	"fmt"
	"runtime"
	"strings"
	"sync"

	client "github.com/sacloud/api-client-go"
	v1 "github.com/sacloud/apprun-api-go/apis/v1"
)

// DefaultAPIRootURL デフォルトのAPIルートURL
const DefaultAPIRootURL = "https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api"

// UserAgent APIリクエスト時のユーザーエージェント
var UserAgent = fmt.Sprintf(
	"apprun-api-go/%s (%s/%s; +https://github.com/sacloud/apprun-api-go) %s",
	Version,
	runtime.GOOS,
	runtime.GOARCH,
	client.DefaultUserAgent,
)

// Client APIクライアント
type Client struct {
	// Profile usacloud互換プロファイル名
	Profile string

	// Token APIキー: トークン
	Token string
	// Token APIキー: シークレット
	Secret string

	// APIRootURL APIのリクエスト先URLプレフィックス、省略可能
	APIRootURL string

	// Options HTTPクライアント関連オプション
	Options *client.Options

	// DisableProfile usacloud互換プロファイルからの設定読み取りを無効化
	DisableProfile bool
	// DisableEnv 環境変数からの設定読み取りを無効化
	DisableEnv bool

	initOnce sync.Once
	factory  *client.Factory
}

func (c *Client) serverURL() string {
	v := DefaultAPIRootURL
	if c.APIRootURL != "" {
		v = c.APIRootURL
	}
	if !strings.HasSuffix(v, "/") {
		v += "/"
	}
	return v
}

func (c *Client) init() error {
	var initError error
	c.initOnce.Do(func() {
		var opts []*client.Options
		// 1: Profile
		if !c.DisableProfile {
			o, err := client.OptionsFromProfile(c.Profile)
			if err != nil {
				initError = err
				return
			}
			opts = append(opts, o)
		}

		// 2: Env
		if !c.DisableEnv {
			opts = append(opts, client.OptionsFromEnv())
		}

		// 3: UserAgent
		opts = append(opts, &client.Options{
			UserAgent: UserAgent,
		})

		// 4: Options
		if c.Options != nil {
			opts = append(opts, c.Options)
		}

		// 5: フィールドのAPIキー
		opts = append(opts, &client.Options{
			AccessToken:       c.Token,
			AccessTokenSecret: c.Secret,
		})

		c.factory = client.NewFactory(opts...)
	})
	return initError
}

func (c *Client) apiClient() (*v1.ClientWithResponses, error) {
	if err := c.init(); err != nil {
		return nil, err
	}

	return &v1.ClientWithResponses{
		ClientInterface: &v1.Client{
			Server: c.serverURL(),
			Client: c.factory.NewHttpRequestDoer(),
		},
	}, nil
}
