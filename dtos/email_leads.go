package dtos

type EmailScraper struct {
	Id       string    `json:"id" form:"id" valid:"required~Unique id for task is required."`
	Domains  *[]string `json:"domains" form:"domains"`
	Keyword  *string   `json:"keyword" form:"keyword"`
	JobTitle string    `json:"jobTitle" form:"jobTitle" valid:"required~Job title required."`
}

type EmailScraperResp struct {
	Id string `json:"id"`
}

type EmailScraperWebhookResp struct {
	Id     string   `json:"id"`
	Emails []string `json:"emails"`
}
