package dtos

type EmailScraper struct {
	Id string `json:"id" form:"id" valid:"required~Unique id for task is required."`
	//Limit     *uint     `json:"limit" form:"limit"`
	Domains   *[]string `json:"domains" form:"domains"`
	JobTitles []string  `json:"jobTitles" form:"jobTitles" valid:"required~Job title(s) required."`
}

type EmailScraperResp struct {
	Id string `json:"id"`
}

type EmailScraperWebhookResp struct {
	Id     string   `json:"id"`
	Emails []string `json:"emails"`
}
