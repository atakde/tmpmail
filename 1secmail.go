package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Message struct {
	Id      int    `json:"id"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	Date    string `json:"date"`
}

type Email struct {
	Id       int    `json:"id"`
	From     string `json:"from"`
	Subject  string `json:"subject"`
	Date     string `json:"date"`
	Body     string `json:"body"`
	TextBody string `json:"textBody"`
	HTMLBody string `json:"htmlBody"`
}

func getRandomEmailAddresses(count int) (emails []string) {
	if count == 0 {
		count = 10
	}

	apiUrl := "https://www.1secmail.com/api/v1/?action=genRandomMailbox&count=" + strconv.Itoa(count)
	request, error := http.NewRequest("GET", apiUrl, nil)
	if error != nil {
		fmt.Println(error)
	}

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		fmt.Println(error)
	}

	responseBody, error := io.ReadAll(response.Body)
	if error != nil {
		fmt.Println(error)
	}

	err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&emails)
	if err != nil {
		fmt.Println(err)
		return
	}

	if response.StatusCode != 200 {
		fmt.Println("Status Code: ", response.StatusCode)
		fmt.Println("Response Body: ", emails)
		return
	}

	defer response.Body.Close()

	return emails
}

func getDomains() (domains []string) {
	apiUrl := "https://www.1secmail.com/api/v1/?action=getDomainList"
	request, error := http.NewRequest("GET", apiUrl, nil)
	if error != nil {
		fmt.Println(error)
	}

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		fmt.Println(error)
	}

	responseBody, error := io.ReadAll(response.Body)
	if error != nil {
		fmt.Println(error)
	}

	err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&domains)
	if err != nil {
		fmt.Println(err)
		return
	}

	if response.StatusCode != 200 {
		fmt.Println("Status Code: ", response.StatusCode)
		fmt.Println("Response Body: ", domains)
		return
	}

	defer response.Body.Close()

	return domains
}

func getMessages(login string, domain string) (messages []Message) {
	apiUrl := "https://www.1secmail.com/api/v1/?action=getMessages&login=" + login + "&domain=" + domain
	request, error := http.NewRequest("GET", apiUrl, nil)
	if error != nil {
		fmt.Println(error)
	}

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		fmt.Println(error)
	}

	responseBody, error := io.ReadAll(response.Body)
	if error != nil {
		fmt.Println(error)
	}

	err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&messages)
	if err != nil {
		fmt.Println(err)
		return
	}

	if response.StatusCode != 200 {
		fmt.Println("Status Code: ", response.StatusCode)
		fmt.Println("Response Body: ", messages)
		return
	}

	defer response.Body.Close()

	return messages
}

func fetchMessage(login string, domain string, id int) (message Email) {
	apiUrl := "https://www.1secmail.com/api/v1/?action=readMessage&login=" + login + "&domain=" + domain + "&id=" + strconv.Itoa(id)
	request, error := http.NewRequest("GET", apiUrl, nil)
	if error != nil {
		fmt.Println(error)
	}

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		fmt.Println(error)
	}

	responseBody, error := io.ReadAll(response.Body)
	if error != nil {
		fmt.Println(error)
	}

	err := json.NewDecoder(bytes.NewReader(responseBody)).Decode(&message)
	if err != nil {
		fmt.Println(err)
		return
	}

	if response.StatusCode != 200 {
		fmt.Println("Status Code: ", response.StatusCode)
		fmt.Println("Response Body: ", message)
		return
	}

	defer response.Body.Close()

	return message
}

func generateRandomEmail(length int) string {
	rand.Seed(time.Now().UnixNano())

	domains := getDomains()
	randomDomain := domains[rand.Intn(len(domains))]

	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	randomString := make([]byte, length)
	for i := range randomString {
		randomString[i] = charset[rand.Intn(len(charset))]
	}

	return fmt.Sprintf("%s@%s", string(randomString), randomDomain)
}

func parseEmail(email string) (login string, domain string) {
	splitEmail := bytes.Split([]byte(email), []byte("@"))
	login = string(splitEmail[0])
	domain = string(splitEmail[1])

	return login, domain
}
