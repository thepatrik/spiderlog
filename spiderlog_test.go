package spiderlog_test

import (
	"testing"

	"github.com/spider-pigs/spiderlog"
)

func TestLogger(t *testing.T) {
	logger := spiderlog.New(spiderlog.StdoutEnabled(true))
	logger.Info("you should not see this (with go test -v)")
}
