// https://wyliemaster.github.io/gddocs/#

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const(
	URL = "http://www.boomlings.com/database/getGJUserInfo20.php"
	GAMEVERSION = 22
	BINARYVERSION = 36
	SECRET = "Wmfd2893gb7"
)

func getUserInfo(client *http.Client, targetAccountID string) string {
	data := url.Values{}
	data.Set("secret", SECRET)
	data.Set("targetAccountID", targetAccountID)

	req, err := http.NewRequest("POST", URL, strings.NewReader(data.Encode()))
	if err != nil { 
		log.Fatalf("Failed to prepare POST request: %s", err)
	}

	req.Header.Add("User-Agent", "")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send POST request: %s", err)
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read response body: %s", err)
		}
		return string(bodyBytes)
	} else {
		log.Fatalf("Status Code: %s", err)
		return "No Data :("
	}
}

// TODO  make sure this works
func getUsers(client *http.Client, username string) string {
	data := url.Values{}
	data.Set("secret", SECRET)
	data.Set("str", username)

	req, err := http.NewRequest("POST", "http://boomlings.com/database/getGJUsers20.php", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalf("Failed to prepare POST request: %s", err)
	}

	req.Header.Add("User-Agent", "")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send POST request: %s", err)
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read response body: %s", err)
		}
		return string(bodyBytes)
	} else {
		log.Fatalf("Status Code: %s", err)
		return "No Data :("
	}
}

// https://wyliemaster.github.io/gddocs/#/resources/server/user
func parseData(data string) []string {
	return strings.Split(data, ":")
}

func main() {
	client := &http.Client{}
	info := getUserInfo(client, "8498828")
	data := parseData(info)
	fmt.Println(data)

	fmt.Printf("Username: %s\n", data[1])
	fmt.Printf("Stars: %s\n", data[15])
}
