package main

import (
	"log"
	"os"
	"time"

	"reviser_script/internal/db"
	"reviser_script/internal/leetcode"
	"reviser_script/internal/models"
	"reviser_script/internal/validate"

	"github.com/joho/godotenv"
)

func init() {
	if os.Getenv("ENV") != "PROD" {
		log.Println("Running the script in Local Development Enviroment")
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env found! Exiting....")
			os.Exit(1)
		}
	} else {
		log.Println("Running the script in Production")
	}
}

func main() {
	user := models.CreateUser()

	validate.CheckDBCredentialsAreLoaded(user)

	submissions := leetcode.GetRecentAcceptedSubmissions(user)

	log.Printf("No of Submissions Fetched: %d\n", len(submissions))

	for _, sub := range submissions {
		log.Printf("Trying to fetch details for %s\n", sub.TitleSlug)

		description := leetcode.GetProblemDescription(sub.TitleSlug, user)
		code := leetcode.GetSubmissionCodeByID(sub.ID, user)

		err := db.InsertSubmissionToDB(user, sub, code, description)
		if err != nil {
			log.Println("Error inserting:", err)
		} else {
			log.Printf("Inserted %s to DB\n", sub.TitleSlug)
		}
		time.Sleep(1 * time.Second) // be kind to LeetCode
	}
}
