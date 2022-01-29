package main

import (
	"context"
	"encoding/csv"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/refs/roi/pkg/human"

	"github.com/google/go-github/v42/github"
)

type Repository struct {
	Author      string // convenient way to access the parsed url repository author.
	Name        string // convenient way to access the parsed url repository name.
	URL         string // text encoded url. i.e: https://github.com/authed/spicedb.
	Description string // description of the repository.
	Stargazers  int    // number of starts a repo has. Vanity metric essentially, but good measure of popularity.

	LatestPushToDefault string // last time a push event happened in the default branch.
	LastPushed          string // last time it was a push event to the repository.
	LastUpdated         string // last time the repository was updated.
}

func main() {
	// use a default HTTP client, don't really mind about timeouts.
	client := github.NewClient(nil)
	repositories := make(map[string][]Repository, 0)
	bundles := bundleByCategories()

	for category, repos := range bundles {
		for _, repoURL := range repos {
			u, err := url.Parse(repoURL)
			if err != nil {
				panic(err)
			}

			var authorN, repoN string
			parts := strings.Split(u.Path, "/")
			if len(parts) == 3 {
				authorN, repoN = parts[1], parts[2]
			} else {
				// TODO find a way to deal with the errors. Panic is not an option. Accumulate and log them at the end. Fail silently.
				panic("malformed url")
			}

			// list a public repository.
			repo, _, err := client.Repositories.Get(context.Background(), authorN, repoN)
			if err != nil {
				panic(err)
			}

			// in order to get the latest commit, we need to identify the default tree, get the tree sha and request the commits.
			defaultTree, _, err := client.Git.GetTree(context.Background(), authorN, repoN, *repo.DefaultBranch, false)
			if err != nil {
				panic(err)
			}

			latestCommit, _, err := client.Git.GetCommit(context.Background(), authorN, repoN, defaultTree.GetSHA())
			if err != nil {
				panic(err)
			}

			r := Repository{
				Author:              authorN,
				Name:                repoN,
				URL:                 repoURL,
				Description:         repo.GetDescription(),
				Stargazers:          repo.GetStargazersCount(),
				LatestPushToDefault: human.Duration(time.Since(*latestCommit.Author.Date)),
				LastPushed:          human.Duration(time.Since(repo.GetPushedAt().Time)),
				LastUpdated:         human.Duration(time.Since(repo.GetUpdatedAt().Time)),
			}

			repositories[category] = append(repositories[category], r)
		}

	}

	t, err := template.New("README.gotmpl").ParseFiles("../configs/README.gotmpl")
	if err != nil {
		panic(err)
	}

	if err := t.Execute(os.Stdout, repositories); err != nil {
		panic(err)
	}
}

type bundle map[string][]string

// bundleByCategories will parse the csv sources files and bundle repositories by its categories on a go struct.
// It is convenient and should help aid understanding and separation of concerns.
func bundleByCategories() bundle {
	b := make(bundle)

	// TODO do not hardcode file location. Read source from ENV variables / cli flags.
	f, err := os.Open("../configs/examples.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		b[line[1]] = append(b[line[1]], line[0])
	}

	return b
}
