package command

import (
	"fmt"
	"github.com/oneCoder16/issues-tool/pkg/api_github"
	"github.com/oneCoder16/issues-tool/pkg/cli_gui"
	"github.com/spf13/viper"
)

type IssuesListCommand struct {
	api  *api_github.Api
	user string
	repo string
}

func NewIssuesListCommand() Command {
	return &IssuesListCommand{
		api:  api_github.New(viper.GetString("token")),
		user: "",
		repo: "",
	}
}

func (command *IssuesListCommand) SetP(params map[string]interface{}) Command {
	if _, ok := params["user"]; ok {
		command.user = params["user"].(string)
	}

	if _, ok := params["repo"]; ok {
		command.repo = params["repo"].(string)
	}

	return command
}

func (command *IssuesListCommand) Execute() error {
	issuesList, err := command.api.GetIssuess(command.user, command.repo)
	if err != nil {
		return fmt.Errorf("error:%w", err)
	}

	fmt.Println(cli_gui.Table(issuesList))
	return nil
}
