package handler

import (
	"encoding/json"
	"gohexa/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	custService service.CustomerService
}

func NewCustomerHandler(custService service.CustomerService) customerHandler {
	return customerHandler{custService: custService}
}

func (h customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.custService.GetCustomers()
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customers)

}

func (h customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	customer, err := h.custService.GetCustomer(id)
	if err != nil {

		handleError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")

	json.NewEncoder(w).Encode(customer)
}
