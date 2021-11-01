package models

import "emailscraper/exceptions"

type Result struct {
	Result    interface{}
	Exception *exceptions.BaseException
}
