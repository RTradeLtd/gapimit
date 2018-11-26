package krab_test

import (
	"testing"

	"github.com/RTradeLtd/config"
	"github.com/RTradeLtd/grpc/backends/krab"
)

func TestServer(t *testing.T) {
	cfg, err := config.LoadConfig(testCfgPath)
	if err != nil {
		t.Fatal(err)
	}
	if err = krab.NewServer(cfg.Endpoints.Krab.URL, "tcp", cfg); err != nil {
		t.Fatal(err)
	}
}