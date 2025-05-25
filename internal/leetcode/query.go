package leetcode

import (
	"encoding/json"
	"log"
	"reviser_script/internal/models"
	"reviser_script/internal/request"
)

func GetRecentAcceptedSubmissions(user *models.User) []models.Submission {
	query := `
	query recentAcSubmissions($username: String!, $limit: Int!) {
		recentAcSubmissionList(username: $username, limit: $limit) {
			id
			title
			titleSlug
			timestamp
		}
	}`
	// fetch past 8 submissions
	variables := map[string]interface{}{
		"username": user.GetUser(),
		"limit":    8,
	}

	body := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	resp := request.MakeGraphqlRequest(body, user)
	defer resp.Body.Close()

	type RespData struct {
		Data struct {
			RecentSubmissionList []models.Submission `json:"recentAcSubmissionList"`
		} `json:"data"`
	}

	var data RespData
	json.NewDecoder(resp.Body).Decode(&data)

	if len(data.Data.RecentSubmissionList) == 0 {
		log.Println("No recent submission list found!! Check the query")
	}
	return data.Data.RecentSubmissionList
}

func GetProblemDescription(slug string, user *models.User) string {
	query := `
	query questionContent($titleSlug: String!) {
		question(titleSlug: $titleSlug) {
			content
		}
	}`

	variables := map[string]interface{}{
		"titleSlug": slug,
	}

	body := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	resp := request.MakeGraphqlRequest(body, user)
	defer resp.Body.Close()

	type RespData struct {
		Data struct {
			Question struct {
				Content string `json:"content"`
			} `json:"question"`
		} `json:"data"`
	}

	var data RespData
	json.NewDecoder(resp.Body).Decode(&data)

	if data.Data.Question.Content == "" {
		log.Printf("Description for %s returned empty response!!\n", slug)
	}
	return data.Data.Question.Content
}

func GetSubmissionCodeByID(id string, user *models.User) string {
	query := `
	query submissionDetails($submissionId: Int!) {
		submissionDetails(submissionId: $submissionId) {
			code
		}
	}`

	variables := map[string]interface{}{
		"submissionId": id,
	}

	body := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	resp := request.MakeGraphqlRequest(body, user)
	defer resp.Body.Close()

	type RespData struct {
		Data struct {
			SubmissionDetails struct {
				Code string `json:"code"`
			} `json:"submissionDetails"`
		} `json:"data"`
	}

	var data RespData
	json.NewDecoder(resp.Body).Decode(&data)

	if data.Data.SubmissionDetails.Code == "" {
		log.Println("Empty Code fetched! Try updating the cookie!")
	}
	return data.Data.SubmissionDetails.Code
}
