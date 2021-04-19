package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hikarivina/go-http.git/gohttp"
)

var (
	githubHttpClient = getGithubClient()
)

func getGithubClient() gohttp.HttpClient {
	client := gohttp.New()

	commonHeader := make(http.Header)
	commonHeader.Set("Authorization", "Bearer ABC-123")
	client.SetHeaders(commonHeader)

	return client
}

func main() {
	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
}
