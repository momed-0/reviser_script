package models

type Submission struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	TitleSlug string `json:"titleSlug"`
	Timestamp string `json:"timestamp"`
}
