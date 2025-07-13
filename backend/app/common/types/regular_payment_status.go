package commonType

// "active", "closed", "cancelled", "overdue", "partial"
type RegularPaymentStatus int8

const (
	RegularStatusActive RegularPaymentStatus = iota + 1
	RegularStatusClosed
	RegularStatusCancelled
	RegularStatusOverdue
	RegularStatusPartial
)

func (s RegularPaymentStatus) String() string {
	names := [...]string{"", "active", "closed", "cancelled", "overdue", "partial"}
	if !s.IsValid() {
		return "unknown"
	}
	return names[s]
}

func (s RegularPaymentStatus) IsValid() bool {
	return s >= RegularStatusActive && s <= RegularStatusPartial
}
