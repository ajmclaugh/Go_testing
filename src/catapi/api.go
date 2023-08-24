package catapi

import(
	"log"
	"io"
	"net/http"
)



func GetCatAPIRequest (endpoint string) string {

	req, err := http.Get(endpoint)

	if err != nil {
		log.Fatalf("An error occurred during the GET request %v", err)
	}

	defer req.Body.Close()

	response, err := io.ReadAll(req.Body)

	
	if err != nil {
		log.Fatalf("An Error Occurred %v", err)
	}

	return string(response)
	
}