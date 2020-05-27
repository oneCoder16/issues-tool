package command

import (
	"fmt"
	"github.com/oneCoder16/issues-tool/pkg/api_github"
	"github.com/oneCoder16/issues-tool/pkg/cli_gui"
	"github.com/spf13/viper"
)

type CommentsCommand struct {
	api       *api_github.Api
	user      string
	repo      string
	issues_id int
}

func NewCommentsCommand() Command {
	return &CommentsCommand{
		api:       api_github.New(viper.GetString("token")),
		user:      "",
		repo:      "",
		issues_id: 0,
	}
}

func (command *CommentsCommand) SetP(params map[string]interface{}) Command {
	if _, ok := params["user"]; ok {
		command.user = params["user"].(string)
	}

	if _, ok := params["repo"]; ok {
		command.repo = params["repo"].(string)
	}

	if _, ok := params["issues_id"]; ok {
		command.issues_id = params["issues_id"].(int)
	}

	return command
}

func (command *CommentsCommand) Execute() error {
	comments, err := command.api.GetComments(command.user, command.repo, command.issues_id)
	if err != nil {
		return err
	}

	fmt.Println(cli_gui.Table(comments))
	return nil
}
