package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := genai.NewClient(
		ctx,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Models.GenerateContent(
		ctx,
		"gemini-3.5-flash",
		genai.Text("What is the skies color only respond with one word"),
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(resp.Text())
}
