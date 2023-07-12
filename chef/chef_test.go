package chef

import (
	"testing"

	"github.com/u123456t/exercise/order"
)

func TestGetResponseChan(t *testing.T) {
	ch := CreateChef("test")
	response := ch.GetResponseChan()

	if response == nil {
		t.Failed()
	}
}

func TestCookfood(t *testing.T) {
	ch := CreateChef("test")
	foodmap := order.InitMenu()
	food := foodmap[1]
	ch.CookFood(food)
}
