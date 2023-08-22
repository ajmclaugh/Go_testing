package main


import (
	"net/http"
	"log"
	"io"
	"fmt"
)

func main () {

	endpoint := "https://catfact.ninja/fact"

	req, err := http.Get(endpoint)

	if err != nil {
		log.Fatalf("An Error Occurred %v", err)
	}

	defer req.Body.Close()

	response, err := io.ReadAll(req.Body)

	fmt.Print(string(response))

}