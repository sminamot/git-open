package main

import (
	"fmt"
	"strings"

	"github.com/kevinburke/ssh_config"
)

const defaultProtocol = "https"

type openURLElements struct {
	protocol string
	domain   string
	urlPath  string
}

func getOpenURLElements(gitURL, currentBranch string) *openURLElements {
	var domain, urlPath string
	protocol := defaultProtocol

	s := strings.Split(gitURL, `://`)
	if len(s) > 1 {
		gitProtocol := s[0]

		if gitProtocol == "http" {
			protocol = "http"
		}

		uri := strings.Join(s[1:], "")
		s = strings.Split(uri, "@")
		if len(s) > 1 {
			uri = strings.Join(s[1:], "")
		}

		s = strings.SplitN(uri, "/", 2)
		domain = s[0]

		if len(s) > 1 {
			urlPath = s[1]
		}
	} else {
		s = strings.Split(gitURL, "@")
		uri := strings.Join(s[len(s)-1:], "")
		s = strings.Split(uri, ":")
		domain = s[0]
		urlPath = strings.Join(s[1:], "")
	}

	urlPath = strings.Trim(urlPath, "/")
	urlPath = strings.TrimSuffix(urlPath, ".git")
	if currentBranch != "master" {
		// escape "%" and "#"
		currentBranch = strings.ReplaceAll(currentBranch, "%", "%25")
		currentBranch = strings.ReplaceAll(currentBranch, "#", "%23")
		branchRef := fmt.Sprintf("/tree/%s", currentBranch)
		urlPath += branchRef
	}

	return &openURLElements{
		protocol: protocol,
		domain:   domain,
		urlPath:  urlPath,
	}
}

func resolveDomain(domain string) string {
	hostname := ssh_config.Get(domain, "Hostname")
	if hostname == "" {
		return domain
	}
	return hostname
}
