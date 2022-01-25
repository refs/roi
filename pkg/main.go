package main

import (
	"context"
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/google/go-github/v42/github"
)

func main() {
	// use a default HTTP client, don't really mind about timeouts.
	client := github.NewClient(nil)

	// list a public repository.
	repo, _, err := client.Repositories.Get(context.Background(), "authzed", "spicedb")
	if err != nil {
		panic(err)
	}

	// in order to get the latest commit, we need to identify the default tree, get the tree sha and request the commits.
	defaultTree, _, err := client.Git.GetTree(context.Background(), "authzed", "spicedb", *repo.DefaultBranch, false)
	if err != nil {
		panic(err)
	}

	latestCommit, _, err := client.Git.GetCommit(context.Background(), "authzed", "spicedb", defaultTree.GetSHA())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Description: %s\nStars: %v ⭐️\nlast pushed %v ago\nlast updated %s ago\nlatest push to main branch %v ago",
		*repo.Description,
		*repo.StargazersCount,
		humanize.Time(repo.PushedAt.Time),
		humanize.Time(repo.UpdatedAt.Time),
		humanize.Time(*latestCommit.Author.Date),
	)
}
