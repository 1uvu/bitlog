package tests

import (
	"log"
	"testing"

	"github.com/1uvu/bitlog/core/pkg/config"
)

func TestNewConfig(t *testing.T) {
	confPath := "../config/_example/collector_config.yaml"

	conf, err := config.NewCollectorConfig(confPath, "yaml")
	if err != nil {
		t.Error(err.Error())
		return
	}

	log.Println(conf)
}
