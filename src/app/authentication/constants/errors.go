package constants

import "errors"

var (
	ErrInValidation            = errors.New("validation : ")
	ErrGeneric                 = errors.New("error")
	ErrInvalidName             = errors.New("name is invalid")
	ErrInvalidPhoneNumber      = errors.New("phone number is invalid")
	ErrInvalidEmail            = errors.New("email is invalid")
	ErrInvalidPanCard          = errors.New("invalid PanCard")
	ErrInvalidPassword         = errors.New("password is weak")
	ErrConnectingDB            = errors.New("error while connecting db")
	ErrExecutingQuery          = errors.New("error executing query")
	ErrOpenDatabaseConnection  = errors.New("error opening database connection")
	ErrReadConfigFile          = errors.New("error reading config file")
	ErrDecodeConfig            = errors.New("unable to decode config into struct")
	ErrDatabasePing            = errors.New("error pinging database")
	ErrCheckingEmail           = errors.New("error checking email existence")
	ErrCheckingPanCard         = errors.New("error checking pancard existence")
	ErrEmailExsits             = errors.New("error email already exsits")
	ErrPanCardlExsits          = errors.New("error pancard already exsits")
	ErrInsertingCustomerRecord = errors.New("error inserting customer record")
	ErrCustomerNotFound        = errors.New("customer not found")
	ErrQueryInDB               = errors.New("error querying database")
	ErrEmail                   = errors.New("invalid email")
	ErrPassword                = errors.New("invalid password")
	ErrPasswordNotExists       = errors.New("password does not exist")
	ErrEmailNotExists          = errors.New("email does not exist")
	ErrSignIn                  = errors.New("unable to signin")
	ErrOtpGeneration           = errors.New("failed to generate OTP")
	ErrOtpVerification         = errors.New("failed to verify OTP")
)
