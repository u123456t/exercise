package chef

import (
	"fmt"
	"time"

	"github.com/u123456t/exercise/order"
)

// Chef : the struct of the Chef
type Chef struct {
	chefID       string
	responseChan chan string
}

// GetResponseChan : return the response chan of the chef
func (c *Chef) GetResponseChan() chan string {
	return c.responseChan
}

// CreateChef : create a chef
func CreateChef(chefid string) Chef {
	chef := Chef{chefid, make(chan string, 1)}
	return chef
}

// CookFood : simulate the cook process of a chef
func (c *Chef) CookFood(food order.Food) {
	fmt.Printf("Chef start to cook food %s at time %s\n", food.GetName(), time.Now().String())
	time.Sleep(time.Second * time.Duration(food.GetCostTime()))
	response := fmt.Sprintf("%s cook finished, chief is available\n", food.GetName())
	c.responseChan <- response
}
