package main

import (
	"fmt"

	"github.com/pamela-quiros/go-http-client/gohttp"
)

var (
	githubHttpClient = getGithubClient() //singleton
)

func getGithubClient() gohttp.Client {

	client := gohttp.NewBuilder().
		DisableTimeouts(true).
		SetMaxIdleConnections(5).
		Build()
	return client
}

func main() {
	getUrls()
}

type User struct {
	FirstName string `jsnon:"first_name"`
	LasttName string `jsnon:"last_name"`
}

func getUrls() {
	response, err := githubHttpClient.Get("https://api.github.com", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
	fmt.Println(response.String())
}
