package db

import (
	"fmt"
	"reviser_script/internal/models"
	"reviser_script/internal/request"
	"strconv"
	"time"
)

func InsertSubmission(sub models.Submission, code string, user *models.User) error {

	timestampInt, _ := strconv.ParseInt(sub.Timestamp, 10, 64)
	timestamp := time.Unix(timestampInt, 0).Format(time.RFC3339)

	// Insert leetcode_submissions
	subPayload := map[string]interface{}{
		"submission_id": sub.ID,
		"question_slug": sub.TitleSlug,
		"submitted_at":  timestamp,
		"code":          code,
	}

	headers := map[string]string{
		"apikey":        user.GetDbKey(),
		"Authorization": "Bearer " + user.GetDbKey(),
		"Content-Type":  "application/json",
	}

	err := request.MakeRESTRequest(subPayload,
		user.GetDbURL()+"/rest/v1/leetcode_submissions",
		"POST",
		headers)

	if err != nil {
		return fmt.Errorf("submission insert failed: %v", err)
	}

	return nil
}

func UpsertQuestion(sub models.Submission, description string, user *models.User) error {
	// UPSERT leetcode_questions
	questionPayload := map[string]any{
		"slug":        sub.TitleSlug,
		"title":       sub.Title,
		"description": description,
	}

	headers := map[string]string{
		"apikey":        user.GetDbKey(),
		"Authorization": "Bearer " + user.GetDbKey(),
		"Content-Type":  "application/json",
		"Prefer":        "resolution=merge-duplicates",
	}

	err := request.MakeRESTRequest(questionPayload,
		user.GetDbURL()+"/rest/v1/leetcode_questions",
		"POST",
		headers)

	if err != nil {
		return fmt.Errorf("question upsert failed: %v", err)
	}

	return nil
}
