package utils

import (
	"fmt"
	"net/url"

	"emailscraper/constants"
)

func GenerateDorkUrl(jobTitle string, keyword *string, domains *[]string) []string {
	var d []string
	if domains != nil {
		d = *domains
	} else {
		d = constants.Domains
	}

	var urls = make([]string, len(d))

	var kw string

	if keyword != nil {
		kw = fmt.Sprintf(" intitle:%s ", *keyword)
	} else {
		kw = " "
	}

	for i, domain := range d {
		urls[i] = fmt.Sprintf(
			"https://www.google.com/search?q=%s",
			url.QueryEscape(fmt.Sprintf("%s%sAND intitle:\"@%s\"", jobTitle, kw, domain)),
		)
	}

	return urls
}
