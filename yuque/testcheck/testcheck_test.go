package testcheck

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	res := Add(3, 5)
	time.Sleep(time.Second * 5)
	if res != 8 {
		t.Fatal()
	}
}
