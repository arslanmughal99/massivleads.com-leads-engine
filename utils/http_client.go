package utils

import (
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var (
	_          = godotenv.Load()
	timeout, _ = strconv.ParseInt(os.Getenv("HTTP_TIMEOUT"), 10, 64)
	userPass   = url.UserPassword(os.Getenv("PROXY_USERNAME"), os.Getenv("PROXY_PASSWORD"))
)

// NewProxyClient Create a new http client with proxy
func NewProxyClient(proxy string) *http.Client {
	proxyURL := &url.URL{
		Host:   proxy,
		Scheme: "http",
		User:   userPass,
	}

	c := http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	return &c
}
