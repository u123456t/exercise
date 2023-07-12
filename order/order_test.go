package order

import (
	"testing"
)

func TestGetPrice(t *testing.T) {
	menu := InitMenu()
	for _, i := range menu {
		i.GetPrice()
	}
}

func TestGetCostTime(t *testing.T) {
	menu := InitMenu()
	for _, i := range menu {
		i.GetCostTime()
	}
}

func TestGetName(t *testing.T) {
	menu := InitMenu()
	for _, i := range menu {
		i.GetName()
	}
}

func TestCheckFood(t *testing.T) {
	menu := InitMenu()
	CheckFood("pizza", menu)

	CheckFood("bread", menu)

	CheckFood("x", menu)
}
