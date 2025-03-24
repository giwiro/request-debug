package exc

type InternalError struct {
	Message string
}

func (i InternalError) Error() string {
	return i.Message
}
