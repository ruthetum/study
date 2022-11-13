package sub

import "fmt"

type Notification struct {
}

func NewNotification() *Notification {
	return &Notification{}
}

func (n *Notification) SendWalletCreditNotification() {
	fmt.Println("Sending wallet credit notification")
}

func (n *Notification) SendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}
