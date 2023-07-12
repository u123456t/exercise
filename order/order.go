package order

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

// FoodInterface : interface for food
type FoodInterface interface {
	GetName() string
	GetPrice() int16
	GetCostTime() int16
}

// Food food struct
type Food struct {
	Name  string
	Price int16
	Time  int16
}

// GetPrice get the price
func (f Food) GetPrice() int16 {
	return f.Price
}

// GetCostTime get the cost time
func (f Food) GetCostTime() int16 {
	return f.Time
}

// GetName get the name of the food
func (f Food) GetName() string {
	return f.Name
}

// Initmenu : init the menu
func InitMenu() []Food {

	bytes, err := os.ReadFile("/home/ccloud/src/training/dave/go_exercise_new/src/order/menu.json")
	if err != nil {
		log.Fatal(err)
	}
	var foodlist []Food
	json.Unmarshal(bytes, &foodlist)
	return foodlist
}

// CheckFood : check whether the ordered food is on the menu
func CheckFood(foodName string, menu []Food) (*Food, error) {

	for _, f := range menu {
		if f.GetName() == foodName {
			return &f, nil
		}
	}
	return nil, errors.New("Food is not on the menu")
}
