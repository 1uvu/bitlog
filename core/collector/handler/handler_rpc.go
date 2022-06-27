package handler

import (
	"fmt"
	"log"
	"math/rand"
	"sync"

	"github.com/1uvu/bitlog/core/collector/conn"
)

// Put on hold

type RPCHandler struct {
	mux sync.Mutex

	rookie, live, idle map[int]*conn.RPConn
	connIdx            int // a increasement id
	connNum, poolSize  int // conn num and pool size, conn num may > pool size

	option *conn.RPCOption
}

func NewRPCHandler(poolSize int, option *conn.RPCOption) (*RPCHandler, error) {
	rookie := make(map[int]*conn.RPConn)
	for i := 0; i < poolSize; i++ {
		c, err := conn.NewRPConn(i, option)
		if err != nil {
			return nil, err
		}
		rookie[i] = c
	}
	h := &RPCHandler{
		rookie:   rookie,
		live:     make(map[int]*conn.RPConn),
		idle:     make(map[int]*conn.RPConn),
		connIdx:  poolSize,
		connNum:  poolSize,
		poolSize: poolSize,
		option:   option,
	}
	return h, nil
}

func (h *RPCHandler) String() string {
	return fmt.Sprintf("%d, %d, %d, %d, %d, %d\n", len(h.rookie), len(h.live), len(h.idle), h.connIdx, h.connNum, h.poolSize)
}

// select a conn to process the rpc call, and if no rookie conn will create a new automatically, will update the conn types
// after call success will backtrack the conn types
func (h *RPCHandler) Call(fcn string, arg *conn.ConnCallArg, reply *conn.ConnCallReply) {
	h.mux.Lock()
	defer h.mux.Unlock()
	c, err := h.selectConn()
	if err != nil {
		reply.Err = err
		return
	}
	// c -> live
	h.switchConnStatus(c, conn.Live)
	// call
	c.Call(fcn, arg, reply)
	// c -> idle
	h.switchConnStatus(c, conn.Idle)
	// try release idle conn
	h.releaseConn()
}

func (h *RPCHandler) selectConn() (*conn.RPConn, error) {
	m, n := len(h.idle), len(h.rookie)
	if m == 0 && n == 0 {
		// no idle and rookie conn
		c, err := h.createConn()
		if err != nil {
			return nil, err
		}
		return c, nil
	}
	var key = 0
	// if m > 0, there has probability of 80% to select conn from idle
	p := rand.Intn(100)
	randKey := func(m map[int]*conn.RPConn, l int) int {
		var rk int
		r := rand.Intn(l)
		i := 0
		for k := range m {
			i++
			rk = k
			if i > r {
				break
			}
		}
		return rk
	}
	if m > 0 && p < 80 {
		key = randKey(h.idle, m)
	} else {
		key = randKey(h.rookie, n)
	}
	c, ok := h.rookie[key]
	// TODO
	if !ok {
		log.Println(key)
		panic("c is nil")
	}
	return c, nil
}

func (h *RPCHandler) createConn() (*conn.RPConn, error) {
	c, err := conn.NewRPConn(h.connIdx+1, h.option)
	if err != nil {
		return nil, err
	}
	h.connIdx++
	h.connNum++
	h.rookie[c.ID] = c
	return c, nil
}

func (h *RPCHandler) switchConnStatus(c *conn.RPConn, newStatus conn.RPConnStatus) {
	// 1 update types
	oldStatus := c.SwitchStatus(newStatus)
	// 2 update old set
	switch oldStatus {
	case conn.Rookie:
		delete(h.rookie, c.ID)
	case conn.Live:
		delete(h.live, c.ID)
	case conn.Idle:
		delete(h.idle, c.ID)
	}
	// 3 update new set
	switch newStatus {
	case conn.Rookie:
		h.rookie[c.ID] = c
	case conn.Live:
		h.live[c.ID] = c
	case conn.Idle:
		h.idle[c.ID] = c
	}
}

func (h *RPCHandler) releaseConn() {
	m, n := len(h.idle), len(h.rookie)
	needRelease := m != 0 || n != 0                     // has idle or rookie conn
	needRelease = needRelease && h.connNum > h.poolSize // curr conn num > pool size
	needRelease = needRelease && h.poolSize/m < 5       // idle conn > pool size/5
	if !needRelease {
		return
	}
	releaseCount := m / 2 // c from idle
	for i := 0; i < releaseCount; i++ {
		// 1 select
		c, err := h.selectConn()
		if err != nil {
			// dont need to handle error, just return
			return
		}
		// 3 release from map
		delete(h.idle, c.ID)
		h.connNum--
		// 4 close conn
		c.Shutdown()
	}
}
