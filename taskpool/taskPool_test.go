package taskpool

import (
	"github.com/u123456t/exercise/order"
	"testing"
	"time"
)

func TestGetRunningChefs(t *testing.T) {
	p := NewPool(3)
	got := p.GetRunningChefs()
	want := 0
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}

func TestAddTask(t *testing.T) {
	p := NewPool(3)
	recieveTime := time.Now()
	foodMap := order.InitMenu()
	task := &Task{
		RecieveTime: recieveTime,
		Order:       foodMap[0],
	}
	p.AddTask(task)
	p.AddTask(task)
	p.AddTask(task)
	p.AddTask(task)
}

func TestGetCapacity(t *testing.T) {
	p := NewPool(3)
	got := p.GetCapacity()
	want := 3
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}
