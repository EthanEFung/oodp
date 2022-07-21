package main

import "fmt"

type paymentValidationHandler struct {
	next handler
}

func (h *paymentValidationHandler) setNext(next handler) {
	h.next = next
}
func (h *paymentValidationHandler) execute(order *order) {
	payment := order.payment
	if payment.network == "Master Card" && payment.cardNumber[0] == '5' {
		fmt.Println("payment type -> Master Card")
		h.next.execute(order)
		return
	}
	fmt.Println("payment not accepted")
}
