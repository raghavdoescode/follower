package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	user := os.Args[1]
	followUser(user, client)
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
