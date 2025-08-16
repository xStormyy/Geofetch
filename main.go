package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const(
	GAMEVERSION = 22
	BINARYVERSION = 36
	SECRET = "Wmfd2893gb7"
)

func getUserInfo(client *http.Client, targetAccountID string) string {
	data := url.Values{}
	data.Set("secret", SECRET)
	data.Set("targetAccountID", targetAccountID)

	req, err := http.NewRequest("POST", "http://www.boomlings.com/database/getGJUserInfo20.php", strings.NewReader(data.Encode()))
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

func getUsers(client *http.Client, username string) string {
	data := url.Values{}
	data.Set("secret", SECRET)
	data.Set("str", username)

	req, err := http.NewRequest("POST", "http://www.boomlings.com/database/getGJUsers20.php", strings.NewReader(data.Encode()))
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

func parseData(data string) []string {
	// TODO try to remove the numbers indicating some sort of index and sort them by https://wyliemaster.github.io/gddocs/#/resources/server/user
	return strings.Split(data, ":")
}

func main() {
	client := &http.Client{}

	if (len(os.Args) > 1) {
		id := getUsers(client, os.Args[1])
		info := getUserInfo(client, parseData(id)[21])
		data := parseData(info)
		fmt.Println(data)

		fmt.Printf("Username: %s\n", data[1])
		fmt.Printf("Stars: %s\n", data[15])
		fmt.Printf("Moons: %s\n", data[17])
		fmt.Printf("Secret Coins: %s\n", data[6])
		fmt.Printf("User Coins: %s\n", data[8])
		fmt.Printf("Demons: %s\n", data[22])
	} else {
		fmt.Println("ERROR: Please provide a username!\n\nExample:\n$ geofetch RobTop")
	}
}
