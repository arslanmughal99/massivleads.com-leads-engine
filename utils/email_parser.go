package utils

import (
	"regexp"
)

var (
	emailRegex     = regexp.MustCompile(`^[a-zA-Z0-9._{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)
	grabEmailRegex = regexp.MustCompile("[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?")
)

func ParseEmails(body string) []string {
	emails := grabEmailRegex.FindAllString(body, -1)
	emails = RemoveDuplicate(emails)

	var checkedEmails []string

	for _, email := range emails {
		if ok := emailRegex.MatchString(email); ok {
			checkedEmails = append(checkedEmails, email)
		}
	}

	return checkedEmails
}
