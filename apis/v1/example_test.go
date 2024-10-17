package v1_test

import (
	"context"
	"fmt"
	"io"
	"os"

	v1 "github.com/sacloud/apprun-api-go/apis/v1"
)

var serverURL = "https://secure.sakura.ad.jp/cloud/api/apprun/1.0/apprun/api"

func Example_getUser() {
	token := os.Getenv("SAKURACLOUD_ACCESS_TOKEN")
	secret := os.Getenv("SAKURACLOUD_ACCESS_TOKEN_SECRET")

	client, err := v1.NewClientWithResponses(serverURL, func(c *v1.Client) error {
		c.RequestEditors = []v1.RequestEditorFn{
			v1.AppRunAuthInterceptor(token, secret),
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	resp, err := client.GetUser(context.Background())
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	// Output:
}
