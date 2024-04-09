package constants

const (
	ValidatedEmailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	ValidatePANCardNumber = `^[A-Z]{5}[0-9]{4}[A-Z]{1}$`
	ValidatePassword      = `^[a-zA-Z0-9!@#$%^&*()_+}{":;'?/><.,~]{8,}$`
)
