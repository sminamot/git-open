package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetOpenURLElements(t *testing.T) {
	tests := []struct {
		gitURL        string
		currentBranch string
		want          *openURLElements
	}{
		{
			gitURL:        "https://github.com/sminamot/git-open.git",
			currentBranch: "master",
			want: &openURLElements{
				protocol: "https",
				domain:   "github.com",
				urlPath:  "sminamot/git-open",
			},
		},
		{
			gitURL:        "https://github.com/sminamot/git-open.git",
			currentBranch: "develop",
			want: &openURLElements{
				protocol: "https",
				domain:   "github.com",
				urlPath:  "sminamot/git-open/tree/develop",
			},
		},
		{
			gitURL:        "http://github.com/sminamot/git-open.git",
			currentBranch: "master",
			want: &openURLElements{
				protocol: "http",
				domain:   "github.com",
				urlPath:  "sminamot/git-open",
			},
		},
		{
			gitURL:        "http://github.com/sminamot/git-open.git",
			currentBranch: "develop",
			want: &openURLElements{
				protocol: "http",
				domain:   "github.com",
				urlPath:  "sminamot/git-open/tree/develop",
			},
		},
		{
			gitURL:        "ssh://github.com/sminamot/git-open.git",
			currentBranch: "master",
			want: &openURLElements{
				protocol: "https",
				domain:   "github.com",
				urlPath:  "sminamot/git-open",
			},
		},
		{
			gitURL:        "ssh://github.com/sminamot/git-open.git",
			currentBranch: "develop",
			want: &openURLElements{
				protocol: "https",
				domain:   "github.com",
				urlPath:  "sminamot/git-open/tree/develop",
			},
		},
		{
			gitURL:        "ssh://git@github.com/sminamot/git-open.git",
			currentBranch: "master",
			want: &openURLElements{
				protocol: "https",
				domain:   "github.com",
				urlPath:  "sminamot/git-open",
			},
		},
		{
			gitURL:        "ssh://git@github.com/sminamot/git-open.git",
			currentBranch: "develop",
			want: &openURLElements{
				protocol: "https",
				domain:   "github.com",
				urlPath:  "sminamot/git-open/tree/develop",
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("%s-%s", tt.currentBranch, tt.gitURL), func(t *testing.T) {
			t.Parallel()
			if got := getOpenURLElements(tt.gitURL, tt.currentBranch); cmp.Diff(got, tt.want, cmp.AllowUnexported(openURLElements{})) != "" {
				t.Fatalf("want = %v, got = %v", tt.want, got)
			}
		})
	}
}
