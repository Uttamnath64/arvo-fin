package commonType

// Bank, Credit, Wallet, Investment
type AccountType int8

const (
	AccountTypeBank AccountType = iota + 1
	AccountTypeCash
	AccountTypeWallet
	AccountTypeCredit
	AccountTypeLoan
	AccountTypeInvestment
)

func (t AccountType) String() string {
	return [...]string{"Bank", "Cash", "Wallet", "Credit", "Loan", "Investment"}[t]
}

func (t AccountType) IsValid() bool {
	return t >= AccountTypeBank && t <= AccountTypeInvestment
}
