package exc

type UnauthorizedError struct {
	Message string
}

func (unauthorizedError UnauthorizedError) Error() string {
	return unauthorizedError.Message
}
