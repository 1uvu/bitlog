package handler

import (
	"errors"
	"sync"

	"github.com/1uvu/bitlog/collector/clientmgr"
	"github.com/1uvu/bitlog/collector/event"
	"github.com/1uvu/bitlog/pkg/pool"
)

type LogHandler struct {
	mux sync.Mutex
	wg  sync.WaitGroup

	c        *clientmgr.LogClient
	shutdown chan bool

	listenerSet sync.Map // map[*CollectorEventListener]struct{}
	pool        *pool.Pool

	eventQueue []event.CollectorEvent // need lock
	queueSize  int32
	head, tail int32

	eventFilter *event.CollectorEventFilter
}

func NewLogHandler(_c *clientmgr.LogClient, _bufferSize int32, _pool *pool.Pool, _filterStr string) (*LogHandler, error) {
	if _c == nil {
		return nil, errors.New("error from _c == nil")
	}
	h := &LogHandler{
		c:           _c,
		shutdown:    make(chan bool),
		listenerSet: sync.Map{},
		pool:        _pool,
		eventQueue:  make([]event.CollectorEvent, _bufferSize+1),
		queueSize:   _bufferSize,
		head:        0,
		tail:        1,
		eventFilter: event.FromStrFilter(_filterStr),
	}
	return h, nil
}

// register a listener and disable GetNextEvent
func (h *LogHandler) Loop(l event.CollectorEventListener) {
	h.listenerSet.Store(&l, struct{}{})
}

func (h *LogHandler) GetNextEvent() (event.CollectorEvent, error) {
	return h.dequeueEvent()
}

func (h *LogHandler) Shutdown() {
	h.mux.Lock()
	defer h.mux.Unlock()
	h.doShutdown()
}

func (h *LogHandler) doShutdown() {
	h.shutdown <- true
	close(h.shutdown)
	_ = h.c.Watcher().Close()
}

func (h *LogHandler) Run() {
	h.wg.Add(1)
	go h.run()
}

func (h *LogHandler) run() {
	for {
		select {
		case e, ok := <-h.c.Watcher().Events:
			if ok {
				ce := event.FromFsEvent(e)
				if !h.eventFilter.FilterOut(ce) {
					h.enqueueEvent(ce)
					h.gotEvent(ce)
				}
			}
		case x := <-h.shutdown:
			if x {
				h.wg.Done()
				return
			}
		case <-h.c.Watcher().Errors:
		}
	}
}

func (h *LogHandler) Wait() {
	h.wg.Wait()
}

func (h *LogHandler) RunAndWait() {
	h.Run()
	h.wg.Wait()
}

func (h *LogHandler) gotEvent(ce event.CollectorEvent) {
	// loop if exist listener
	h.listenerSet.Range(func(key, value interface{}) bool {
		l := key.(*event.CollectorEventListener)
		h.tryRunWithPool(ce, *l)
		return true
	})
}

func (h *LogHandler) enqueueEvent(ce event.CollectorEvent) error {
	h.mux.Lock()
	defer h.mux.Unlock()
	// push TODO

	return nil
}

func (h *LogHandler) dequeueEvent() (event.CollectorEvent, error) {
	h.mux.Lock()
	defer h.mux.Unlock()
	// pop TODO

	return event.CollectorEvent{}, nil
}

func (h *LogHandler) tryRunWithPool(ce event.CollectorEvent, l event.CollectorEventListener) {
	if h.pool != nil {
		h.pool.RunFunc(func() {
			l(ce)
		})
	} else {
		l(ce)
	}
}
