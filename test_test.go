package bitclient

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestName(t *testing.T) {
	const (
		proj     = "JOH"
		repoSlug = "dgraph"
	)

	client := NewBitClient("http://localhost:7990", os.Getenv("USER"), os.Getenv("PASS"))
	prResp, err := client.CreatePullRequest(proj, repoSlug, CreatePullRequestParams{
		Title:       "oneone",
		Description: "desc",
		FromRef: BranchRef{
			Id: "harshil-goel/searchable-search",
			Repository: Repository{
				Slug: "dgraph",
				Project: Project{
					Key: "JOH",
				},
			},
		},
		ToRef: BranchRef{
			Id: "master",
			Repository: Repository{
				Slug: "dgraph",
				Project: Project{
					Key: "JOH",
				},
			},
		},
		Reviewers: []Participant{
			{User: User{Name: "johnny"}},
		},
		CloseSourceBranch: false,
	})
	require.NoError(t, err)

	c1, err := client.CreatePullRequestComment(proj, repoSlug, strconv.Itoa(prResp.Id), CreatePullRequestCommentParams{
		Text: "salam popolam!",
	})
	require.NoError(t, err)

	_, err = client.CreatePullRequestComment(proj, repoSlug, strconv.Itoa(prResp.Id), CreatePullRequestCommentParams{
		Text: "salam popolam!",
		Parent: &CreatePullRequestCommentParentParams{
			Id: c1.Id,
		},
	})
	require.NoError(t, err)
}
