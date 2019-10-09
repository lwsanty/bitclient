package bitclient

import (
	"fmt"
)

func (bc *BitClient) GetPullRequestSettings(projectKey string, repositorySlug string) (PullRequestSettings, error) {
	response := PullRequestSettings{}

	_, err := bc.DoGet(
		fmt.Sprintf("/projects/%s/repos/%s/settings/pull-requests", projectKey, repositorySlug),
		nil,
		&response,
	)

	return response, err
}

func (bc *BitClient) SetPullRequestSettings(projectKey string, repositorySlug string, params PullRequestSettings) error {

	// Remove disable mergeConfig strategies or they'll get enabled :/
	var filteredStrategies []PullRequestStrategy
	for _, strategy := range params.MergeConfig.Strategies {
		if strategy.Enabled == true {
			filteredStrategies = append(filteredStrategies, strategy)
		}
	}
	params.MergeConfig.Strategies = filteredStrategies

	_, err := bc.DoPost(
		fmt.Sprintf("/projects/%s/repos/%s/settings/pull-requests", projectKey, repositorySlug),
		params,
		nil,
	)

	return err
}

// GetPullRequestsRequest defines the available parameters when requesting the list of pull requests
// from a repository.
type GetPullRequestsRequest struct {
	PagedRequest
	Direction      string `url:"direction,omitempty"`
	At             string `url:"at,omitempty"`
	State          string `url:"state,omitempty"`
	Order          string `url:"order,omitempty"`
	WithAttributes bool   `url:"withAttributes,omitempty"`
	WithProperties bool   `url:"withProperties,omitempty"`
}

// GetPullRequestsResponse holds the API response data
type GetPullRequestsResponse struct {
	PagedResponse
	Values []PullRequest
}

// GetPullRequests lists the pull requests from given repository
func (bc *BitClient) GetPullRequests(projectKey string, repositorySlug string, params GetPullRequestsRequest) (GetPullRequestsResponse, error) {
	response := new(GetPullRequestsResponse)

	_, err := bc.DoGet(
		fmt.Sprintf("/projects/%s/repos/%s/pull-requests", projectKey, repositorySlug),
		nil,
		response,
	)

	return *response, err
}

type CreatePullRequestParams struct {
	Title             string        `json:"title"`
	Description       string        `json:"description"`
	FromRef           BranchRef     `json:"fromRef"`
	ToRef             BranchRef     `json:"toRef"`
	Reviewers         []Participant `json:"reviewers"`
	CloseSourceBranch bool          `json:"close_source_branch"`
}

type BranchRef struct {
	Id         string     `json:"id"`
	Repository Repository `json:"repository"`
}

func (bc *BitClient) CreatePullRequest(projectKey string, repositorySlug string, params CreatePullRequestParams) error {
	_, err := bc.DoPost(
		fmt.Sprintf("/projects/%s/repos/%s/pull-requests", projectKey, repositorySlug),
		params,
		nil,
	)
	return err
}

// pull-request data example
/*
{
  "title": "put some title1",
  "description": "put some desc",
  "fromRef": {
    "id": "harshil-goel/searchable-search",
    "repository": {
      "slug": "dgraph",
      "name": null,
      "project": {
        "key": "JOH"
      }
    }
  },
  "toRef": {
    "id": "master",
    "repository": {
      "slug": "dgraph",
      "name": null,
      "project": {
        "key": "JOH"
      }
    }
  },
  "reviewers": [
    {
      "user": {
        "name": "johnny"
      }
    }
  ],
  "close_source_branch": false
}
*/
