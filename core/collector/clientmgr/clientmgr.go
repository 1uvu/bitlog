package clientmgr

import (
	"context"

	"github.com/1uvu/bitlog/pkg/config"
)

type CollectorClientMgr struct {
	// ctx  context.Context // TODO

	rpc   *RPClient
	log   *LogClient
	shell *ShellClient
}

func NewCollectorClientMgr(_ctx context.Context, _conf *config.CollectorConfig) (*CollectorClientMgr, error) {
	rc, err := NewRPClient(_conf.RPC)
	if err != nil {
		return nil, err
	}
	lc, err := NewLogClient()
	if err != nil {
		return nil, err
	}
	sc, err := NewShellClient(_conf.Shell)
	if err != nil {
		return nil, err
	}
	mgr := new(CollectorClientMgr)
	mgr.rpc = rc
	mgr.log = lc
	mgr.shell = sc
	return mgr, nil
}

func (mgr *CollectorClientMgr) ClientRPC() *RPClient {
	return mgr.rpc
}

func (mgr *CollectorClientMgr) ClientLog() *LogClient {
	return mgr.log
}

func (mgr *CollectorClientMgr) ClientShell() *ShellClient {
	return mgr.shell
}
