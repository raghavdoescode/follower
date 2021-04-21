package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/nerdthatnoonelikes/follower/src/structs"
)

var token string = "token here"

func main() {
	client := &http.Client{}
	user := os.Args[1]
	// followUser(user, client)
	followers := GetFollowers(user, client)

	for i := 0; i < len(followers); i++ {
		followUser(followers[i].Login, client)
		fmt.Printf(color.GreenString("Followed %s\n"), followers[i].Login)
	}
}

func followUser(username string, client *http.Client) {
	// Make a PUT request
	req, _ := http.NewRequest("PUT", "https://api.github.com/user/following/"+username, nil)
	req.Header.Set("Authorization", "token "+token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	_, err := client.Do(req)

	// Check if an error occured
	if err != nil {
		fmt.Println("An error occured")
	}
}

func GetFollowers(username string, client *http.Client) structs.UsersFollowers {
	req, _ := http.NewRequest("GET", "https://api.github.com/users/"+username+"/followers", nil)
	req.Header.Set("Authorization", "token "+token)

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("An error occured")
	}

	var response structs.UsersFollowers
	bodyBytes, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(bodyBytes, &response)

	return response
}
