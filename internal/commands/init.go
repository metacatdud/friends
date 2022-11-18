package commands

import (
	"friends/internal/usecase"
	"friends/pkg/command"
	"github.com/spf13/cobra"
)

type cmdInit struct {
	client *cobra.Command

	Store string
}

func NewCmdInit() command.Command {
	c := &cmdInit{}

	c.client = buildCmdInit(c)

	return c
}

func (c *cmdInit) Cmd() *cobra.Command {
	return c.client
}

func (c *cmdInit) Execute(cmd *cobra.Command, args []string) error {

	req := usecase.InitReq{
		Storage: c.Store,
	}

	if err := req.Validate(); err != nil {
		return err
	}

	if _, err := usecase.Init(req); err != nil {
		return err
	}

	return nil
}

func buildCmdInit(c *cmdInit) *cobra.Command {
	cc := &cobra.Command{
		Use:   "init",
		Short: "init project",
		RunE:  c.Execute,
	}

	cc.Flags().StringVarP(&c.Store, "storage", "s", "", "Set storage folder")

	return cc
}
