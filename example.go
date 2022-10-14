package main

import (
	"fmt"

	"github.com/pamela-quiros/g-http-client/gohttp"
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
	FirstName string `jsnon:"first_name "`
	LasttName string `jsnon:"last_name "`
}

func getUrls() {
	response, err := githubHttpClient.Get("https://api.github.com", nil)

	if err != nil {
		panic(err)
	}

	//Using our custome response:
	var user User
	if err := response.UnmarshalJson(&user); err != nil {
		panic(err)
	}
	fmt.Println(response.Status())
	fmt.Println(response.StatusCode())
	fmt.Println(response.String())

	/*Using default http.Response:
	response.Body.Close()

	bytes2, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var user2 User
	if err := json.Unmarshal(bytes, &user2); err != nil {
		panic(err)
	}
	fmt.Println(user2.FirstName)*/

}

/* func createUser(user User) {
	response, err := githubHttpClient.Post("https://api.github.com", nil, user)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
*/
