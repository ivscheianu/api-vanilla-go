package error

type PersistenceError struct {
	message string
	cause   error
}

func (this *PersistenceError) Error() string {
	return this.cause.Error()
}

func (this *PersistenceError) GetMessage() string {
	return this.message
}

func (this *PersistenceError) GetCause() error {
	return this.cause
}

func CreatePersistenceError(message string, cause error) *PersistenceError {
	return &PersistenceError{message, cause}
}
