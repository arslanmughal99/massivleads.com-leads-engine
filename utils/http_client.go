package utils

import (
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

var (
	_ = godotenv.Load()

	proxyURL = &url.URL{
		Scheme: "http",
		Host:   os.Getenv("PROXY_HOST"),
		User:   url.UserPassword(os.Getenv("PROXY_USERNAME"), os.Getenv("PROXY_PASSWORD")),
	}

	HttpClient = http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}
)
