package utils

import (
	"fmt"
	"net/url"
)

func GenerateDorkUrl(jobTitle string, countryCode string, keyword *string, domains *[]string) []string {
	var d []string
	if domains != nil {
		d = *domains
	} else {
		d = GetEmailDomains(countryCode)
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
			"https://www.google.%s/search?q=%s", countryCode,
			url.QueryEscape(fmt.Sprintf("%s%sAND (intitle:\"@%s\" OR intext:\"@%s\")", jobTitle, kw, domain, domain)),
		)
	}

	return urls
}
