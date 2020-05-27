package command

import (
	"github.com/oneCoder16/issues-tool/pkg/api_github"
	"github.com/spf13/viper"
)

type AddCommentCommand struct {
	api       *api_github.Api
	user      string
	repo      string
	issues_id int
	content   api_github.Comment
}

func NewAddCommentCommand() Command {
	return &AddCommentCommand{
		api:  api_github.New(viper.GetString("token")),
		user: "",
		repo: "",
		content: api_github.Comment{
			Id:   0,
			Body: "",
		},
	}
}

func (command *AddCommentCommand) SetP(params map[string]interface{}) Command {
	if _, ok := params["user"]; ok {
		command.user = params["user"].(string)
	}

	if _, ok := params["repo"]; ok {
		command.repo = params["repo"].(string)
	}

	if _, ok := params["issues_id"]; ok {
		command.issues_id = params["issues_id"].(int)
	}

	if _, ok := params["content"]; ok {
		command.content.Body = params["content"].(string)
	}

	return command
}

func (command *AddCommentCommand) Execute() error {
	err := command.api.AddComment(command.user, command.repo, command.issues_id, command.content)
	if err != nil {
		return err
	}

	return nil
}
