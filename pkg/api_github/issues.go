package api_github

import (
	"encoding/json"
	"fmt"
	"github.com/idoubi/goz"
)

var (
	reposIssuesUrl string = "https://api.github.com/repos/%s/%s/issues"
	commentsUrl    string = "https://api.github.com/repos/%s/%s/issues/%d/comments"
	addCommentUrl  string = "https://api.github.com/repos/%s/%s/issues/%d/comments"
	editCommentUrl string = "https://api.github.com/repos/%s/%s/issues/%d/comments/%d"
)

type Issues struct {
	HtmlUrl string `json:"html_url"`
	Title   string `json:"title"`
	Number  int    `json:"number"`
}

type Comment struct {
	Id       int    `json:"id" table:"id"`
	Body     string `json:"body" table:"内容"`
	CreateAt string `json:"created_at" table:"发布时间"`
}

func (api *Api) GetIssuess(user, repo string) (issues []Issues, err error) {
	cli := goz.NewClient()
	resp, err := cli.Get(fmt.Sprintf(reposIssuesUrl, user, repo), goz.Options{
		Headers: map[string]interface{}{
			"Authorization": "token " + api.token,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("GetIssuess api error: %w", err)
	}

	body, err := resp.GetBody()
	if err != nil {
		return nil, fmt.Errorf("GetIssuess GetBody error: %w", err)
	}

	if err := json.Unmarshal([]byte(body.GetContents()), &issues); err != nil {
		return nil, fmt.Errorf("GetIssuess json.Unmarshal content: %s error: %w", body.GetContents(), err)
	}

	return issues, nil
}

func (api *Api) GetComments(user, repo string, issuesId int) (comments []Comment, err error) {
	cli := goz.NewClient()
	resp, err := cli.Get(fmt.Sprintf(commentsUrl, user, repo, issuesId), goz.Options{
		Headers: map[string]interface{}{
			"Authorization": "token " + api.token,
		},
	})

	if err != nil {
		return nil, fmt.Errorf("GetComments api error: %w", err)
	}

	body, err := resp.GetBody()
	if err != nil {
		return nil, fmt.Errorf("GetComments GetBody error: %w", err)

	}

	if err := json.Unmarshal([]byte(body.GetContents()), &comments); err != nil {
		return nil, fmt.Errorf("GetComments json.Unmarshal content:%s error: %w", body.GetContents(), err)
	}

	return comments, nil
}

func (api *Api) AddComment(user, repo string, issuesId int, comment Comment) error {
	cli := goz.NewClient()

	resp, err := cli.Post(fmt.Sprintf(addCommentUrl, user, repo, issuesId), goz.Options{
		Headers: map[string]interface{}{
			"Authorization": "token " + api.token,
		},
		JSON: map[string]interface{}{
			"body": comment.Body,
		},
	})

	if err != nil {
		return fmt.Errorf("AddComment api error: %w", err)
	}

	_, err = resp.GetBody()
	if err != nil {
		return fmt.Errorf("AddComment GetBody error: %w", err)
	}

	return nil
}

func (api *Api) EditComment(user, repo string, issuesId int, comment Comment) error {
	cli := goz.NewClient()

	resp, err := cli.Post(fmt.Sprintf(editCommentUrl, user, repo, issuesId, comment.Id), goz.Options{
		Headers: map[string]interface{}{
			"Authorization": "token " + api.token,
		},
		JSON: map[string]interface{}{
			"body": comment.Body,
		},
	})

	if err != nil {
		return fmt.Errorf("EditComment api error: %w", err)
	}

	_, err = resp.GetBody()
	if err != nil {
		return fmt.Errorf("EditComment GetBody error: %w", err)
	}

	return nil
}
