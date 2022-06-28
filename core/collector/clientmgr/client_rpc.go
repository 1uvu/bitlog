package clientmgr

import (
	"strconv"
	"strings"

	"github.com/1uvu/bitlog/pkg/config"

	"github.com/btcsuite/btcd/rpcclient"
)

type RPClient struct {
	*rpcclient.Client
	Conf *config.RPConfig
}

func NewRPClient(conf *config.RPConfig) (*RPClient, error) {
	connCfg := &rpcclient.ConnConfig{
		Host:         strings.Join([]string{conf.Address, strconv.Itoa(conf.Port)}, ":"),
		User:         conf.Username,
		Pass:         conf.Password,
		HTTPPostMode: true,
		DisableTLS:   true,
	}
	rc, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return nil, err
	}
	rpcClient := new(RPClient)
	rpcClient.Client = rc
	rpcClient.Conf = conf
	return rpcClient, nil
}
