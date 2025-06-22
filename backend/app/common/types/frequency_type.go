package commonType

// Daily, Weekly, onthly, Yearly
type FrequencyType int8

const (
	FrequencyTypeDaily FrequencyType = iota + 1
	FrequencyTypeWeekly
	FrequencyTypeMonthly
	FrequencyTypeYearly
)

func (t FrequencyType) String() string {
	return [...]string{"Daily", "Weekly", "Monthly", "Yearly"}[t]
}

func (t FrequencyType) IsValid() bool {
	return t >= FrequencyTypeDaily && t <= FrequencyTypeYearly
}
