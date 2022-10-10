package tests

import (
	"github.com/1uvu/bitlog/collector/logclient"
	"github.com/1uvu/bitlog/pkg/config"
	"testing"
	"time"
)

func TestLogClient(t *testing.T) {
	var (
		conf = &config.CollectorConfig{
			Base:      &config.BaseConfig{BasePath: ""},
			Node:      &config.NodeConfig{LoggerName: "test_logclient"},
			LogClient: nil,
			LogServer: &config.LogServerConfig{Address: "http://localhost:8080"},
		}
		stopCh = make(chan struct{})
	)
	c, err := logclient.NewLogClient(conf, stopCh)
	if err != nil {
		t.Error(err)
	}
	go func() {
		time.Sleep(60 * time.Second)
		stopCh <- struct{}{}
	}()
	c.Run()
}
