package db

import (
	"log"
	"reviser_script/internal/models"
)

func InsertSubmissionToDB(user *models.User, sub models.Submission, code string, description string) error {
	if sub.TitleSlug == "" || code == "" || description == "" {
		log.Println("Skipping DB insert because one of the fields is empty")
		log.Printf("Slug: %s; code: %s; description: %s", sub.TitleSlug, code, description)
		return nil
	}

	err := UpsertQuestion(sub, description, user)

	if err != nil {
		return err
	}

	err = InsertSubmission(sub, code, user)

	if err != nil {
		return err
	}

	return nil
}
