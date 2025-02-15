package common

// OTP type enum
type OtpType int

const (
	Register OtpType = iota + 1
	ResetPassword
)

func (o OtpType) String() string {
	return [...]string{"Register", "ResetPassword"}[o]
}

func (o OtpType) IsValid() bool {
	return o >= Register && o <= ResetPassword
}

const (

	// account types
	ACCOUNT_TYPE_BANK       = "bank"
	ACCOUNT_TYPE_CREDIT     = "credit"
	ACCOUNT_TYPE_WALLET     = "wallet"
	ACCOUNT_TYPE_INVESTMENT = "investment"

	// category types
	CATEGORY_TYPE_INCOME  = "income"
	CATEGORY_TYPE_EXPENSE = "expense"

	// transaction types
	TRANSACTION_TYPE_INCOME   = "income"
	TRANSACTION_TYPE_EXPENSE  = "expense"
	TRANSACTION_TYPE_TRANSFER = "transfer"

	// recurring transaction types
	RECURRING_TRANSACTION_TYPE_INCOME  = "income"
	RECURRING_TRANSACTION_TYPE_EXPENSE = "expense"

	// recurring transaction frequencys
	RECURRING_TRANSACTION_FREQUENCY_DAILY   = "daily"
	RECURRING_TRANSACTION_FREQUENCY_WEEKLY  = "weekly"
	RECURRING_TRANSACTION_FREQUENCY_MONTHLY = "monthly"
	RECURRING_TRANSACTION_FREQUENCY_YEARLY  = "yearly"

	// User type
	USER_TYPE_USER = 1

	// ✅ Success Codes
	StatusSuccess     = 3300 // General success
	StatusUserCreated = 3301 // User created successfully
	StatusUserFetched = 3302 // User fetched successfully
	StatusUpdated     = 3303 // Record updated successfully
	StatusDeleted     = 3304 // Record deleted successfully

	// ❌ Error Codes
	StatusBadRequest      = 3400 // Bad request
	StatusUnauthorized    = 3401 // Unauthorized access
	StatusForbidden       = 3403 // Forbidden
	StatusNotFound        = 3404 // Resource not found
	StatusConflict        = 3409 // Conflict (duplicate entry, etc.)
	StatusServerError     = 3500 // Internal server error
	StatusDatabaseError   = 3501 // Database error
	StatusValidationError = 3502 // Validation error
)
