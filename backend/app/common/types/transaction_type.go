package commonType

// Income, Expense, Transfer
type TransactionType int8

const (
	TransactionTypeIncome TransactionType = iota + 1
	TransactionTypeExpense
	TransactionTypeTransfer
)

func (t TransactionType) String() string {
	names := [...]string{"", "Income", "Expense", "Transfer"}
	if !t.IsValid() {
		return "Unknown"
	}
	return names[t]
}

func (t TransactionType) IsValid() bool {
	return t >= TransactionTypeIncome && t <= TransactionTypeTransfer
}
