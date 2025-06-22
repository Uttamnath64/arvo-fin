package commonType

// Bank, Cash, Wallet, Credit, Loan, Investment
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
	names := [...]string{"", "Bank", "Cash", "Wallet", "Credit", "Loan", "Investment"}
	if !t.IsValid() {
		return "Unknown"
	}
	return names[t]
}

func (t AccountType) IsValid() bool {
	return t >= AccountTypeBank && t <= AccountTypeInvestment
}
