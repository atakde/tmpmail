package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		printCommands()
		// exit
		return
	}

	switch os.Args[1] {
	case "generate-random":
		fmt.Println(generateRandomEmail(10))
		break
	case "messages":
		if len(os.Args) < 3 {
			fmt.Println("go run main.go messages <email>")
			return
		}

		login, domain := parseEmail(os.Args[2])
		messages := getMessages(login, domain)
		fmt.Println(messages)
		break
	case "message":
		if len(os.Args) < 4 {
			fmt.Println("go run main.go message <email> <id>")
			return
		}

		login, domain := parseEmail(os.Args[2])
		intId, _ := strconv.Atoi(os.Args[3])
		messageBody := fetchMessage(login, domain, intId)
		fmt.Println(messageBody)
		break
	}
}

func printCommands() {
	fmt.Println("commands:")
	fmt.Println("generate-random")
	fmt.Println("messages <email>")
	fmt.Println("message <email> <id>")
}
