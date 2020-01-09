package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/skratchdot/open-golang/open"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func main() {
	os.Exit(run())
}

func run() int {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	r, err := git.PlainOpenWithOptions(wd, &git.PlainOpenOptions{
		DetectDotGit: true,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	h, err := r.Reference(plumbing.HEAD, false)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	currentBranch := strings.TrimPrefix(h.Target().String(), "refs/heads/")

	list, err := r.Remotes()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	var gitURL string
	for _, r := range list {
		if r.Config().Name != "origin" {
			continue
		}
		gitURL = r.Config().URLs[0]
	}
	if gitURL == "" {
		fmt.Fprintln(os.Stderr, "not set url")
	}

	// openUrl is format of "xxx://xxx/xxx"
	e := getOpenURLElements(gitURL, currentBranch)
	domain := resolveDomain(e.domain)
	openURL := fmt.Sprintf("%s://%s/%s", e.protocol, domain, e.urlPath)
	err = open.Run(openURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}
