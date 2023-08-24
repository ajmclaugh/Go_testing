package catapi

import(
	"testing"
)

func TestGetCatAPIRequest(t *testing.T) {
	cat_api_result := GetCatAPIRequest("https://catfact.ninja/fact")
	
	if cat_api_result == "" {
		t.Fatalf("no value returned")

	}
}