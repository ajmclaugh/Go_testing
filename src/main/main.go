package main


import (
	"catapi"
	"fmt"
	"os"
)

// API to test ->"https://catfact.ninja/fact"


func main () {

	endpoint := os.Args[1] 

	response := catapi.GetCatAPIRequest(endpoint)
	
	fmt.Print(string(response))

}