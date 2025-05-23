package validate

import (
	"fmt"
	"log"
	"os"

	"reviser_script/internal/models"
	"reviser_script/internal/request"
)

func CheckDBCredentialsAreLoaded(user models.User) {
	if user.GetDbKey() == "" || user.GetDbURL() == "" {
		log.Println("DB credentials are not loaded!!!")
		os.Exit(1)
	}
	err := checkSupabaseConnection(user)
	if err != nil {
		log.Println("Error pinging database!!", err)
		os.Exit(1)
	}
}

func checkSupabaseConnection(user models.User) error {

	headers := map[string]string{
		"apikey":        user.GetDbKey(),
		"Authorization": "Bearer " + user.GetDbKey(),
	}

	err := request.MakeRESTRequest(nil,
		user.GetDbURL()+"/rest/v1/leetcode_questions?limit=1",
		"GET",
		headers)

	if err != nil {
		return fmt.Errorf("invalid response from Supabase: %s", err)
	}

	return nil
}
