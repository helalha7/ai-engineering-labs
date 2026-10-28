package main

import (
	"ai-tests/ai/classification"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal(err)
	}

	client, err := genai.NewClient(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	classifier := classification.Classifier{
		Client: client,
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		message = strings.TrimSpace(message)

		if message == "end" {
			break
		}

		c, err := classifier.Classify(message)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("%+v\n", c)
	}

}
