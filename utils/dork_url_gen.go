package utils

import (
	"fmt"
	"net/url"
	"strings"

	"emailscraper/constants"
)

func GenerateDorkUrl(jobTitle []string, domains *[]string) []string {
	var d []string
	if domains != nil {
		d = *domains
	} else {
		d = constants.Domains
	}

	var urls = make([]string, len(d))
	jb := strings.Join(jobTitle, " ")

	for i, domain := range d {
		urls[i] = fmt.Sprintf(
			"https://www.google.com/search?q=%s", url.QueryEscape(fmt.Sprintf("%s AND intitle:\"@%s\"", jb, domain)),
		)
	}

	return urls
}
