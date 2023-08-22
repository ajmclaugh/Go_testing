package catapi

import(
	"log"
	"io"
	"net/http"
)



func GetCatAPIRequest (endpoint string) []byte {

	req, err := http.Get(endpoint)

	if err != nil {
		log.Fatalf("An Error Occurred %v", err)
	}

	defer req.Body.Close()

	response, err := io.ReadAll(req.Body)

	
	if err != nil {
		log.Fatalf("An Error Occurred %v", err)
	}


	return response
	
}