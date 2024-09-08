package main

import (
	"net/url"
	"strings"
)

func normalizeURL(urlStr string) (string, error) {
	url, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	host := url.Hostname()
	path := strings.TrimSuffix(url.Path, "/")
	return host + path, nil
}
