package commonType

type PaymentFrequency int8

const (
	PaymentFrequencyDaily PaymentFrequency = iota + 1
	PaymentFrequencyWeekly
	PaymentFrequencyBiWeekly // every 2 weeks
	PaymentFrequencyMonthly
	PaymentFrequencyBiMonthly   // every 2 months
	PaymentFrequencyQuarterly   // every 3 months
	PaymentFrequencyFourMonthly // every 4 months
	PaymentFrequencyHalfYearly  // every 6 months
	PaymentFrequencyYearly
)

func (f PaymentFrequency) String() string {
	names := [...]string{
		"",
		"daily",
		"weekly",
		"bi-weekly",
		"monthly",
		"bi-monthly",
		"quarterly",
		"four-monthly",
		"half-yearly",
		"yearly",
	}
	if !f.IsValid() {
		return "unknown"
	}
	return names[f]
}

func (f PaymentFrequency) IsValid() bool {
	return f >= PaymentFrequencyDaily && f <= PaymentFrequencyYearly
}
