package group

import (
	"context"
	"os"
	"syscall"
	"testing"
	"time"
)

func teardown() {
	initHook()
}

func taskFunc(ctx context.Context) {
	select {
	case <-ctx.Done():
	case <-time.After(time.Second):
	}
}

func TestGraceExit(t *testing.T) {
	defer teardown()
	graceExitCh := make(chan struct{})
	signalWatchedCh := make(chan struct{})
	graceExitHook = func() {
		close(graceExitCh)
	}
	signalWatchedHook = func() {
		close(signalWatchedCh)
	}
	Go(taskFunc)
	<-signalWatchedCh
	pid := os.Getpid()
	syscall.Kill(pid, os.Interrupt.(syscall.Signal))
	Wait()
	select {
	case <-graceExitCh:
	default:
		t.Fatal("grace exit failed")
	}

}
