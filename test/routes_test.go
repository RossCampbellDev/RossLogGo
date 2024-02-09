package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/rcampbell-sec/RossLogGo/internal/types"
)

var (
	url  string
	port string
)

func init() {
	url = "http://localhost"
	port = "8080"
}

func TestGetJSONRequest(t *testing.T) { // []types.Entry
	queryString := "logs"
	url := fmt.Sprintf("%s:%s/%s", url, port, queryString)

	res, err := http.Get(url)
	if err != nil {
		t.Errorf("request failed")
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("could not read request body")
	}

	var entries []types.Entry
	if err := json.Unmarshal(resBody, &entries); err != nil {
		t.Errorf("could not unmarshal request body into slice of entries")
	}

	if len(entries) == 0 || entries == nil {
		t.Errorf("No entries in resulting slice")
	}
}
