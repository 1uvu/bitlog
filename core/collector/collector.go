package collector

import (
	"context"

	"github.com/1uvu/bitlog/core/collector/clientmgr"
	"github.com/1uvu/bitlog/core/collector/conn"
	"github.com/1uvu/bitlog/core/collector/handler"
	"github.com/1uvu/bitlog/core/pkg/config"
	"github.com/1uvu/bitlog/core/pkg/pool"
)

// todo mod, move collector schedule into app backend

type Collector struct {
	// ctx       context.Context TODO
	clientmgr *clientmgr.CollectorClientMgr
}

func NewCollector(_ctx context.Context, _conf *config.CollectorConfig) (*Collector, error) {
	ctx := context.Background() // TODO
	// 1 base the conf, get the clientmgr and add observer for log client, each client has its controller
	clientmgr, err := clientmgr.NewCollectorClientMgr(ctx, _conf)
	if err != nil {
		return nil, err
	}
	collector := new(Collector)
	// collector.ctx = _ctx
	collector.clientmgr = clientmgr
	return collector, nil
}

func (c *Collector) ClientMgr() *clientmgr.CollectorClientMgr {
	return c.clientmgr
}

func (c *Collector) HandlerLog(_bufferSize int32, _pool *pool.Pool, _filterStr string) (*handler.LogHandler, error) {
	return handler.NewLogHandler(c.ClientMgr().ClientLog(), _bufferSize, _pool, _filterStr)
}

func (c *Collector) DefaultHandlerLog() (*handler.LogHandler, error) {
	return c.HandlerLog(0, nil, "")
}

func (c *Collector) HandlerRPC(_poolSize int, _option *conn.RPCOption) (*handler.RPCHandler, error) {
	return handler.NewRPCHandler(_poolSize, _option)
}

func (c *Collector) DefaultHandlerRPC() (*handler.RPCHandler, error) {
	return c.HandlerRPC(2, &conn.RPCOption{Conf: c.ClientMgr().ClientRPC().Conf})
}
