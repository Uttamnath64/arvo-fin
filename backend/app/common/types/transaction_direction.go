package commonType

// "Borrowed", "Repaid", "Given", "Received", "Saved"
type TransactionDirection int8

const (
	TransactionDirectionBorrowed TransactionDirection = iota + 1 // You borrowed from someone
	TransactionDirectionRepaid                                   // You repaid the borrowed amount
	TransactionDirectionGiven                                    // You gave someone money
	TransactionDirectionReceived                                 // You received back what you gave
	TransactionDirectionSaved                                    // You saved money (e.g., into a savings account or envelope)

)

func (d TransactionDirection) String() string {
	names := [...]string{
		"", "Borrowed", "Repaid", "Given", "Received", "Saved",
	}
	if !d.IsValid() {
		return "Unknown"
	}
	return names[d]
}

func (d TransactionDirection) IsValid() bool {
	return d >= TransactionDirectionBorrowed && d <= TransactionDirectionSaved
}
