package examples

import (
	"net/http"
	"time"

	"github.com/pamela-quiros/go-http-client/gohttp"
	"github.com/pamela-quiros/go-http-client/gomime"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)

	client := gohttp.NewBuilder().
		SetHeaders(headers).
		//SetHttpClient(nil).
		SetConnectionTimeout(2 * time.Second).
		SetResponsetTimeout(3 * time.Second).
		SetUserAgent("Pam's").
		Build()
	return client
}
