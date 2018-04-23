package domain

// StoreError handle error in application level
type StoreError struct {
	At      string
	Message string
	Details string
}

// NewStoreError create new application error
func NewStoreError(at, message, details string) *StoreError {
	a := new(StoreError)
	a.At = at
	a.Message = message
	a.Details = details

	return a
}

// Error returning error string
func (a *StoreError) Error() string {
	return a.At + ": " + a.Message + ", " + a.Details
}
