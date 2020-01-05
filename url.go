package main

import (
	"fmt"
	"strings"
)

const defaultProtocol = "https"

func getOpenURL(gitURL, currentBranch string) string {
	s := strings.Split(gitURL, `://`)
	if len(s) < 2 {
		return ""
	}

	gitProtocol := s[0]

	uri := strings.Join(s[1:], "")
	s = strings.Split(uri, "@")
	if len(s) > 1 {
		uri = strings.Join(s[1:], "")
	}

	s = strings.SplitN(uri, "/", 2)
	domain := s[0]

	var urlPath string
	if len(s) > 1 {
		urlPath = s[1]
	}
	urlPath = strings.Trim(urlPath, "/")
	urlPath = strings.TrimSuffix(urlPath, ".git")

	protocol := defaultProtocol
	if gitProtocol == "http" {
		protocol = "http"
	}

	openURL := fmt.Sprintf("%s://%s/%s", protocol, domain, urlPath)
	if currentBranch != "master" {
		// escape "%" and "#"
		currentBranch = strings.ReplaceAll(currentBranch, "%", "%25")
		currentBranch = strings.ReplaceAll(currentBranch, "#", "%23")
		branchRef := fmt.Sprintf("/tree/%s", currentBranch)
		openURL += branchRef
	}

	return openURL
}
