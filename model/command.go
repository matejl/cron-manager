package model

type command struct {
	ShellCmd string `json:"shellCmd"`
}

func NewShellCommand(shellCmd string) command {
	return command{
		ShellCmd: shellCmd,
	}
}