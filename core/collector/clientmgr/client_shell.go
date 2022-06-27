package clientmgr

import (
	"github.com/1uvu/bitlog/core/pkg/config"
)

type ShellClient struct {
}

func NewShellClient(conf *config.ShellConfig) (*ShellClient, error) {
	return nil, nil
}

func (c *ShellClient) Shutdown() {
}

func (c *ShellClient) WaitForShutdown() {
}
