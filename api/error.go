package api

/** Модель, сообщения об ошибке. */
type Error struct {
	Message string    `json:"message"`
}

func NewError(msg string) *Error {
	return &Error{Message: msg}
}
