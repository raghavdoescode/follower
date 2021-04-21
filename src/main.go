package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/nerdthatnoonelikes/follower/src/structs"
)

func main() {
	client := &http.Client{}
	user := os.Args[1]
	// followUser(user, client)
	GetFollowers(user, client)
}

func followUser(username string, client *http.Client) {
	// Make a PUT request
	req, _ := http.NewRequest("PUT", "https://api.github.com/user/following/"+username, nil)
	req.Header.Set("Authorization", "token "+"tokenhere")
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	_, err := client.Do(req)

	// Check if an error occured
	if err != nil {
		fmt.Println("An error occured")
	}
}

func GetFollowers(username string, client *http.Client) structs.UsersFollowers {
	req, _ := http.NewRequest("GET", "https://api.github.com/users/"+username+"/followers", nil)
	req.Header.Set("Authorization", "token "+"tokenhere")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("An error occured")
	}

	var response structs.UsersFollowers
	bodyBytes, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(bodyBytes, &response)

	return response
}
