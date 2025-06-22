package commonType

// Income, Expense
type TransactionType int8

const (
	TransactionTypeIncome TransactionType = iota + 1
	TransactionTypeExpense
)

func (t TransactionType) String() string {
	names := [...]string{"Income", "Expense"}
	if !t.IsValid() {
		return "Unknown"
	}
	return names[t]
}

func (t TransactionType) IsValid() bool {
	return t >= TransactionTypeIncome && t <= TransactionTypeExpense
}
