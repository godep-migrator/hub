package main

type AuthUserCommand struct{}

var AuthUser AuthUserCommand

func (a *AuthUserCommand) Execute(args []string) error {
	return nil
}

func (a *AuthUserCommand) Usage() string {
	return ""
}