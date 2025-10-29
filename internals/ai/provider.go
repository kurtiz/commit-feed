package ai

import "github.com/kurtiz/commit-feed/internals/git"

type GeneratedPosts struct {
	LinkedIn string
	Twitter  string
}

type Provider interface {
	GeneratePosts(commits []git.Commit, platforms []string, projectContext string) (*GeneratedPosts, error)
}
