package constants

import "errors"

var (
	ErrInvalidEmail            = errors.New("invalid email format")
	ErrInvalidPanCardFormat    = errors.New("invalid PanCard format")
	ErrInvalidPasswordFormat   = errors.New("invalid Password format")
	ErrConnectingDB            = errors.New("error while connecting db")
	ErrExecutingQuery          = errors.New("error executing query")
	ErrCompileRegex            = errors.New("error compiling regular expression")
	ErrOpenDatabaseConnection  = errors.New("error opening database connection")
	ErrReadConfigFile          = errors.New("error reading config file")
	ErrDecodeConfig            = errors.New("unable to decode config into struct")
	ErrDatabasePing            = errors.New("error pinging database")
	ErrCheckingEmail           = errors.New("error checking email existence")
	ErrCheckingPanCard         = errors.New("error checking pancard existence")
	ErrEmailExsits             = errors.New("err email already exsits")
	ErrPanCardlExsits          = errors.New("err pancard already exsits")
	ErrInsertingCustomerRecord = errors.New("error inserting customer record")
)
