package command

import (
	"fmt"
)

type Broker struct {
	commands map[string]Command
}

func NewBroker() *Broker {
	broker := &Broker{
		commands: make(map[string]Command),
	}
	broker.RegisterC()
	return broker
}

func (this *Broker) RegisterC() {
	this.addC("get_issues_list", NewIssuesListCommand())
	this.addC("get_comment_list", NewCommentsCommand())
	this.addC("add_comment", NewAddCommentCommand())
	this.addC("edit_comment", NewEditCommentCommand())
}

func (this *Broker) addC(name string, command Command) {
	this.commands[name] = command
}

func (this *Broker) Execute(name string, params map[string]interface{}) error {
	command, ok := this.commands[name]
	if !ok {
		return fmt.Errorf("command %s not exists", name)
	}

	return command.SetP(params).Execute()
}
