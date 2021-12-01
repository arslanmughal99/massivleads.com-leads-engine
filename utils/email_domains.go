package utils

import (
	"strings"

	"emailscraper/constants"
)

func GetEmailDomains(countryCode string) []string {
	var selDomains []string

	for _, domain := range constants.Domains {
		if strings.Contains(domain, countryCode) {
			selDomains = append(selDomains, domain)
		}
	}

	if len(selDomains) < 10 {
		selDomains = constants.Domains
	}

	return selDomains
}
