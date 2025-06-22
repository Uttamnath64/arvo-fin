package commonType

// Income, Expense
type TransactionType int8

const (
	TransactionTypeIncome TransactionType = iota + 1
	TransactionTypeExpense
)

func (t TransactionType) String() string {
	return [...]string{"Income", "Expense"}[t]
}

func (t TransactionType) IsValid() bool {
	return t >= TransactionTypeIncome && t <= TransactionTypeExpense
}
