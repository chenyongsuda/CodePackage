package timecache

import (
	"testing"
	"time"
)

func TestDemo(t *testing.T) {
	group := Cache("test")

	group.Add("key-1", "test-1", 4*time.Second)
	//group.Add("key-2", "test-2", 150*time.Millisecond)
	//group.Add("key-3", "test-3", 500*time.Millisecond)
	//group.Add("key-4", "test-3", 115*time.Millisecond)
	time.Sleep(5 * time.Second)
}
