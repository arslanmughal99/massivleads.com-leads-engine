package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"emailscraper/dtos"
	"emailscraper/exceptions"
	"emailscraper/utils"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var (
	_          = godotenv.Load()
	webhookURL = os.Getenv("WEBHOOK_URL")
)

// ScrapeEmailLeads scrape email leads
func ScrapeEmailLeads(dto dtos.EmailScraper) error {
	var emails []string
	mtx := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, url := range utils.GenerateDorkUrl(dto.JobTitle, dto.Keyword, dto.Domains) {
		wg.Add(1)
		go func(_url string) {
			_emails := scrapeUrlPages(_url)
			mtx.Lock()
			emails = append(emails, _emails...)
			mtx.Unlock()
			wg.Done()
		}(url)
	}

	wg.Wait()

	if len(emails) < 5 {
		exp := exceptions.NewBaseException(http.StatusInternalServerError, dto.Id, "Not enough emails found.")
		resp, _ := json.Marshal(exp)
		postWebhookResult(resp)
		return nil
	}
	response := new(dtos.EmailScraperWebhookResp)
	response.Id = dto.Id
	response.Emails = emails
	resp, _ := json.Marshal(response)
	postWebhookResult(resp)
	return nil
}

// scrapeUrlPages scrape first 10 pages for given dork url
func scrapeUrlPages(url string) []string {
	var emails []string
	mtx := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(page int) {
			body, err := getBody(url, page)

			if err != nil {
				wg.Done()
				return
			}

			_emails := utils.ParseEmails(*body)

			mtx.Lock()
			emails = append(emails, _emails...)
			mtx.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
	return emails
}

// getBody Get body of http response
func getBody(dorkUrl string, page int) (*string, error) {
	resp, err := utils.HttpClient.Get(fmt.Sprintf("%s&start=%d", dorkUrl, page*10))
	if err != nil {
		log.Error().Err(err).Int("Page", page).Str("Url", dorkUrl).Msg("Http request error")
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		//log.Error().Int("Page", page).Str("Url", dorkUrl).Str("Status", resp.Status).Msg("Bad response")
		return nil, errors.New("response not ok")
	}

	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Int("Page", page).Str("Url", dorkUrl).Msg("IO read body failed")
		return nil, err
	}

	body := string(rawBody)

	return &body, nil
}

// postWebhookResult post result on webhook with retry
func postWebhookResult(result []byte) {
	for attempt := 0; attempt <= 9; attempt++ {
		res, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(result))
		if err != nil {
			log.Error().Err(err).Msg("Http error when calling webhook")
			time.Sleep(time.Second * 5)
			continue
		}

		if res.StatusCode != http.StatusOK {
			log.Error().Err(err).Msg("Http error when calling webhook")
			time.Sleep(time.Second * 5)
			continue
		}
		log.Debug().Msg("Posted data on webhook")
		break
	}
}
