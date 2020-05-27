package command

type Command interface {
	SetP(parms map[string]interface{}) Command
	Execute() error
}
