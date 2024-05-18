package utilities

type MaybeError struct {
	Name string
	Err  error
}

func NewMaybeError(name string, err error) *MaybeError {
	return &MaybeError{
		Name: name,
		Err:  err,
	}
}
