package sub

import "fmt"

type Ledger struct {
}

func NewLedger() *Ledger {
	return &Ledger{}
}

func (s *Ledger) MakeEntry(accountID, txnType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %d\n", accountID, txnType, amount)
	return
}
