package testcheck

import (
	"fmt"
	"testing"
	"time"
)

func TestA(t *testing.T) {
	time.Sleep(time.Second * 1)
	t.Cleanup(func() {
		fmt.Println("a-1")
	})
	t.Cleanup(func() {
		fmt.Println("a-2")
	})
	t.Cleanup(func() {
		fmt.Println("a-3")
	})
	t.Run("B", TestB)
}

func TestB(t *testing.T) {
	t.Log("this B")
	time.Sleep(time.Second * 1)
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}
