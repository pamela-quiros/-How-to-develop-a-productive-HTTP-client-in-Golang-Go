package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pamela-quiros/g-http-client/gohttp"
)

var (
	githubHttpClient = getGithubClient() //singleton
)

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()

	commonHeaders := make(http.Header)
	commonHeaders.Set("Authorization", "Bearer ABC-123")
	client.SetHeaders(commonHeaders)

	return client
}

func main() {
	getUrls()
}

type User struct {
	FirstName string `jsnon:"first_name "`
	LasttName string `jsnon:"last_name "`
}

func getUrls() {

	response, err := githubHttpClient.Get("https://api.github.com", nil)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}

func createUser(user User) {
	response, err := githubHttpClient.Post("https://api.github.com", nil, user)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
