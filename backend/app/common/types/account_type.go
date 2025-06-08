package commonType

// Bank, Credit, Wallet, Investment
type AccountType int

const (
	AccountTypeBank AccountType = iota + 1
	AccountTypeCredit
	AccountTypeWallet
	AccountTypeInvestment
)

func (t AccountType) String() string {
	return [...]string{"Bank", "Credit", "Wallet", "Investment"}[t]
}

func (t AccountType) IsValid() bool {
	return t >= AccountTypeBank && t <= AccountTypeInvestment
}
