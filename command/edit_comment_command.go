package command

import (
	"github.com/oneCoder16/issues-tool/pkg/api_github"
	"github.com/spf13/viper"
)

type EditCommentCommand struct {
	api       *api_github.Api
	user      string
	repo      string
	issues_id int
	content   api_github.Comment
}

func NewEditCommentCommand() Command {
	return &EditCommentCommand{
		api:  api_github.New(viper.GetString("token")),
		user: "",
		repo: "",
		content: api_github.Comment{
			Id:   0,
			Body: "",
		},
	}
}

func (command *EditCommentCommand) SetP(params map[string]interface{}) Command {
	if _, ok := params["user"]; ok {
		command.user = params["user"].(string)
	}

	if _, ok := params["repo"]; ok {
		command.repo = params["repo"].(string)
	}

	if _, ok := params["issues_id"]; ok {
		command.issues_id = params["issues_id"].(int)
	}

	if _, ok := params["comment_id"]; ok {
		command.content.Id = params["comment_id"].(int)
	}

	if _, ok := params["content"]; ok {
		command.content.Body = params["content"].(string)
	}

	return command
}

func (command *EditCommentCommand) Execute() error {
	err := command.api.EditComment(command.user, command.repo, command.issues_id, command.content)
	if err != nil {
		return err
	}

	return nil
}
