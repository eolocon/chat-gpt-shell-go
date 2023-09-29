package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Print(banner)

	authToken := os.Getenv("OPENAI_AUTH_TOKEN")
	if authToken == "" {
		fmt.Println("Openai auth token not found.\nPlease set OPENAI_AUTH_TOKEN env variable.")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	client := GetClient(
		&ChatGptClientConfiguration{
			authToken:   authToken,
			model:       model,
			n:           numberOfChoices,
			temperature: temperature,
		},
	)

	fmt.Print(mainShellWelcomeMessage)
	// input loop
	for {
		fmt.Print(mainShellHeader)
		scanner.Scan()
		fmt.Println()

		// application ends if error occurs during input
		if err := scanner.Err(); err != nil {
			fmt.Printf(genericErrorMessage, err)
			os.Exit(1)
		}

		switch scanner.Text() {
		case "quit":
			fallthrough
		case "exit":
			fmt.Print(mainShellExitMessage)
			os.Exit(0)
		case "help":
			fmt.Print(mainShellHelpMessage)
		case "hello-world":
			helloWorldCommandHandler(client)
		case "echo":
			echoCommandHandler(client)
		case "text-adventure":
			textAdventureCommandHandler(client)
		default:
			fmt.Printf(mainShellDefaultCaseMessage, scanner.Text())
		}

	}
}

func helloWorldCommandHandler(client *ChatGptClient) {

	message := client.convert(helloWorldPrompt)
	response, err := client.send(message)

	if err != nil {
		fmt.Printf(chatGptClientErrorMessage, err)
	} else {
		fmt.Printf(chatGptHeader, response.Choices[0].Message.Content)
	}
}

func echoCommandHandler(client *ChatGptClient) {
	fmt.Print(echoShellWelcomeMessage)

	scanner := bufio.NewScanner(os.Stdin)

	for endLoop := false; !endLoop; {
		fmt.Print(echoShellHeader)
		scanner.Scan()
		fmt.Println()

		// application ends if error occurs during input
		if err := scanner.Err(); err != nil {
			fmt.Printf(genericErrorMessage, err)
			os.Exit(1)
		}

		switch strings.Trim(scanner.Text(), " \t\n") {
		case "quit":
			fallthrough
		case "exit":
			fmt.Print(echoShellExitMessage)
			endLoop = true
		case "":
			fmt.Print(emptyStringMessage)
		default:
			message := client.convert(scanner.Text())
			response, err := client.send(message)

			if err != nil {
				fmt.Printf(chatGptClientErrorMessage, err)
			} else {
				fmt.Printf(chatGptHeader, response.Choices[0].Message.Content)
			}
		}

	}
}

func textAdventureCommandHandler(client *ChatGptClient) {
	fmt.Print(textAdventureShellWelcomeMessage)

	instructions := []string{taskPrompt, settingPrompt, allowedActionPrompt, obstaclesPrompt, potentiallyHarmfulActionsPrompt, goalPrompt, finalPrompt}

	messages := client.convert(instructions...)
	_, err := client.send(messages)

	// application ends if error occurs during input
	if err != nil {
		fmt.Printf(chatGptClientErrorMessage, err)
		os.Exit(1)
	}

	fmt.Print(textAdventureInstructionsMessage)

	scanner := bufio.NewScanner(os.Stdin)

	for endLoop := false; !endLoop; {
		fmt.Print(textAdventureShellHeader)
		scanner.Scan()
		fmt.Println()

		// application ends if error occurs during input
		if err := scanner.Err(); err != nil {
			fmt.Printf(genericErrorMessage, err)
			os.Exit(1)
		}

		switch strings.Trim(scanner.Text(), " \t\n") {
		case "quit":
			fallthrough
		case "exit":
			fmt.Print(textAdventureExitMessage)
			endLoop = true
		case "":
			fmt.Print(emptyStringMessage)
		default:
			messages = append(messages, client.convert(scanner.Text())...)
			response, err := client.send(messages)

			if err != nil {
				fmt.Printf(chatGptClientErrorMessage, err)
			} else {
				fmt.Printf(chatGptHeader, response.Choices[0].Message.Content)
				messages = append(messages, response.Choices[0].Message)
			}
		}
	}

}
