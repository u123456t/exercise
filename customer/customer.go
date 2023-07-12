package customer

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/u123456t/exercise/order"
)

// Customer : struct for customer
type Customer struct {
	CustomerID string
	Food       order.Food
	WatingTime int
}

func (c *Customer) setFood(f order.Food) {
	c.Food = f
}

func (c *Customer) setWaitingTime(t int) {
	c.WatingTime = t
}

// InitCustomer : init the customers
func InitCustomer() []Customer {
	foodMap := order.InitMenu()
	numberOfFood := len(foodMap)
	customerQueue := []Customer{}
	for i := 0; i < 10; i++ {
		waitingTime := rand.Intn(10)
		customerQueue = append(customerQueue, Customer{
			CustomerID: strconv.Itoa(i + 1),
			Food:       foodMap[rand.Intn(numberOfFood)],
			WatingTime: waitingTime,
		})
	}
	return customerQueue
}

// OrderFood : simulate the process for the customers order food
func OrderFood(customerQueue []Customer) {
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for _, customer := range customerQueue {
		foodname := customer.Food.GetName()
		waitingTime := customer.WatingTime
		customerID := customer.CustomerID
		fmt.Printf("Cusotomer %v waiting %v seconds and will order %s.\n", customer.CustomerID, customer.WatingTime, foodname)
		go func() {
			time.Sleep(time.Duration(waitingTime) * time.Second)
			requestURL := "http://localhost:8090/serve?food=" + foodname + "&customer=" + customerID
			response, err := http.Get(requestURL)
			if err != nil {
				log.Fatal(err)
			}
			defer response.Body.Close()
			_, err = io.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
