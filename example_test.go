package group_test

import (
	"context"
	"time"

	"github.com/wwq1988/group"
)

func taskFunc(ctx context.Context) {
	select {
	case <-ctx.Done():
	case <-time.After(time.Second):
	}
}
func ExampleApp() {
	group.Go(taskFunc)
	group.Wait()
}
