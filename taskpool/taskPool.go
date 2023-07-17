package taskpool

import (
	"fmt"
	"github.com/u123456t/exercise/chef"
	"github.com/u123456t/exercise/order"
	"strconv"
	"sync"
	"time"
)

// Task : struct for Task
type Task struct {
	RecieveTime time.Time
	Order       order.Food
}

// Pool : struct for task pool
type Pool struct {
	capacity     int
	runningChefs int
	chOrder      chan *Task
	sync.Mutex
}

// singleton for task pool
var instance *Pool

// GetRunningChefs : return the number of the running chefs in the task pool
func (p *Pool) GetRunningChefs() int {
	return p.runningChefs
}

// GetCapacity : return the number of the total capacity of the task pool
func (p *Pool) GetCapacity() int {
	return p.capacity
}

// NewPool : create the task pool
func NewPool(capacity int) *Pool {
	if instance == nil {
		instance = &Pool{
			capacity: capacity,
			chOrder:  make(chan *Task, capacity),
		}
	}
	return instance
}

func (p *Pool) run(f order.Food) {
	p.runningChefs++
	go func() {
		defer func() {
			p.runningChefs--
		}()
		for task := range p.chOrder {
			//start to handle task
			handleTask(task, p)
		}
	}()
}

// AddTask : add the task to the task pool
func (p *Pool) AddTask(t *Task) {
	p.Lock()
	defer p.Unlock()
	p.chOrder <- t

	if p.GetRunningChefs() < p.GetCapacity() {
		p.run(t.Order)
	}
}

func handleTask(t *Task, p *Pool) {
	ch := chef.CreateChef("chef" + strconv.Itoa(p.GetRunningChefs()))
	ch.CookFood(t.Order)
	costTime := time.Since(t.RecieveTime).Seconds()
	fmt.Printf("Food %s cooke finished, the price is %d, and it cost %v seconds\n", t.Order.GetName(), t.Order.GetPrice(), int(costTime))
	message := <-ch.GetResponseChan()
	fmt.Print(message)
}
