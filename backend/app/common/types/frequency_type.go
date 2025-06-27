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
	names := [...]string{"", "Daily", "Weekly", "Monthly", "Yearly"}
	if !t.IsValid() {
		return "Unknown"
	}
	return names[t]
}

func (t FrequencyType) IsValid() bool {
	return t >= FrequencyTypeDaily && t <= FrequencyTypeYearly
}
