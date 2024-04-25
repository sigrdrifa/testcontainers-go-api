package main

import (
	"fmt"
	"os"

	http_client "github.com/sigrdrifa/go-api-testcontainers/internal/http-client"
)

func main() {
	client, err := http_client.NewClient("https://api.agify.io")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := client.GetAge("Sig")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(res)

}
