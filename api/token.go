package api

/** Моедль представления, токена. */
type Token struct {
	Token string `json:"token"`
}

func NewToken(toke string) *Token {
	return &Token{Token: toke}
}
