package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"reviser_script/internal/models"
)

func MakeGraphqlRequest(body map[string]any, user *models.User) *http.Response {
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "https://leetcode.com/graphql", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "LEETCODE_SESSION="+user.GetSession())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Failed to connect to LeetCode:", err)
		os.Exit(1)
	}
	return resp
}
func MakeRESTRequest(payload map[string]any, endpoint string, method string, headers map[string]string) error {

	client := &http.Client{}
	var req *http.Request

	payloadBytes, _ := json.Marshal(payload)
	req, _ = http.NewRequest(method, endpoint, bytes.NewBuffer(payloadBytes))

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	res, err := client.Do(req)
	if err != nil || res.StatusCode >= 300 {
		if res.StatusCode == 409 {
			return fmt.Errorf("conflict Error: It seems like the record already exists in the database")
		}
		return fmt.Errorf("failed Request: Status Code %d", res.StatusCode)
	}
	defer res.Body.Close()
	return nil
}
