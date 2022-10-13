package gohttpclient

import "github.com/pamela-quiros/g-http-client/gohttp"

func basicExample() {
	client := gohttp.New()

	client.Get()
}
