package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/skratchdot/open-golang/open"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func main() {
	flag.Parse()
	r, b := flag.Arg(0), flag.Arg(1)
	os.Exit(run(r, b))
}

func run(targetRemote, targetBranch string) int {
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

	list, err := r.Remotes()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	var gitURL string
	// default remote: origin
	if targetRemote == "" {
		targetRemote = "origin"
	}
	for _, r := range list {
		if r.Config().Name != targetRemote {
			continue
		}
		gitURL = r.Config().URLs[0]
	}
	if gitURL == "" {
		fmt.Fprintf(os.Stderr, "not set %s url\n", targetRemote)
		return 1
	}

	// default: current branch
	if targetBranch == "" {
		targetBranch = strings.TrimPrefix(h.Target().String(), "refs/heads/")
	}

	// openUrl is format of "xxx://xxx/xxx"
	e := getOpenURLElements(gitURL, targetBranch)
	domain := resolveDomain(e.domain)
	openURL := fmt.Sprintf("%s://%s/%s", e.protocol, domain, e.urlPath)
	err = open.Run(openURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	return 0
}
