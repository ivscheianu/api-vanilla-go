package error

type ServiceError struct {
	message        string
	cause          error
	httpStatusCode int
}

func (this *ServiceError) GetMessage() string {
	return this.message
}

func (this *ServiceError) GetCause() error {
	return this.cause
}

func (this *ServiceError) GetStatusCode() int {
	return this.httpStatusCode
}

func CreateServiceError(message string, cause error, httpStatusCode int) *ServiceError {
	return &ServiceError{message, cause, httpStatusCode}
}
