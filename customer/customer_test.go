package customer

import (
	"net/http"
	"os"
	"testing"
)

func TestInitCustomer(t *testing.T) {
	InitCustomer()
}

func TestOrderFood(t *testing.T) {
	customers := InitCustomer()
	go func() {
		http.HandleFunc("/serve", mockhandler)
		http.ListenAndServe(":8090", nil)
	}()
	go func() {
		OrderFood(customers)
	}()

}

func mockhandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	os.Exit(0)
}
