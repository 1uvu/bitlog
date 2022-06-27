package conn

import (
	"sync"

	"github.com/1uvu/bitlog/core/collector/clientmgr"
	"github.com/1uvu/bitlog/core/pkg/config"
)

type RPConn struct {
	mux    sync.Mutex
	client *clientmgr.RPClient

	ID     int
	Status RPConnStatus
	// IdleTime time.Duration // conn in idle for a IdleTime
}

type RPConnStatus int8

const (
	Rookie RPConnStatus = iota // conn in pool
	Live                       // conn is live
	Idle                       // conn is live but idle (not shutdown)
)

type RPCOption struct {
	Conf *config.RPConfig
	// IdleTime time.Duration
}

func NewRPConn(id int, option *RPCOption) (*RPConn, error) {
	c, err := clientmgr.NewRPClient(option.Conf)
	if err != nil {
		return nil, err
	}
	return &RPConn{
		client: c,
		ID:     id,
		Status: Rookie,
		// IdleTime: option.IdleTime,
	}, nil
}

func (c *RPConn) Lock() {
	c.mux.Lock()
}

func (c *RPConn) Unlock() {
	c.mux.Unlock()
}

func (c *RPConn) SwitchStatus(newStatus RPConnStatus) RPConnStatus {
	c.Lock()
	defer c.Unlock()
	oldStatus := c.Status
	if oldStatus != newStatus {
		c.Status = newStatus
	}
	return oldStatus
}

func (c *RPConn) Call(fcn string, arg *ConnCallArg, reply *ConnCallReply) {
	c.Lock()
	defer c.Unlock()
	if f, ok := callHandleFuncs[fcn]; ok {
		f(c, arg, reply)
	}
}

func (c *RPConn) Shutdown() {
	c.Lock()
	defer c.Unlock()
	c.client.Shutdown()
}

func (c *RPConn) WaitForShutdown() {
	c.Lock()
	defer c.Unlock()
	c.client.WaitForShutdown()
}
