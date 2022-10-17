package examples

import (
	"time"

	"github.com/pamela-quiros/go-http-client/gohttp"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponsetTimeout(3 * time.Second).
		Build()

	return client
}
