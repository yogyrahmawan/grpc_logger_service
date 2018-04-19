package domain

// ApplicationError handle error in application level
type ApplicationError struct {
	At      string
	Message string
	Details string
}

// NewApplicationError create new application error
func NewApplicationError(at, message, details string) *ApplicationError {
	a := new(ApplicationError)
	a.At = at
	a.Message = message
	a.Details = details

	return a
}

// Error returning error string
func (a *ApplicationError) Error() string {
	return a.At + ": " + a.Message + ", " + a.Details
}
