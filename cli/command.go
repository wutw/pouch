package main

import (
	"github.com/spf13/cobra"
)

// Command define some interfaces that the command must implement them.
type Command interface {
	Init(*Cli)
	Cmd() *cobra.Command
}

type baseCommand struct {
	cmd *cobra.Command
	cli *Cli
}

func (b *baseCommand) Init(cli *Cli) {}

func (b *baseCommand) Cmd() *cobra.Command {
	return b.cmd
}
