package utils

import (
	"strings"

	"emailscraper/constants"
)

// GetCountryTld get country tld
func GetCountryTld(countryName string) map[string]string {
	tld := constants.CountryTldList[strings.ToLower(countryName)]

	if tld != nil {
		return tld
	}

	// Default to random smartproxy gate
	return constants.CountryTldList[""]
}
