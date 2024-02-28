package controller

// error messages
const (
	errInternalServerMsg = "internal server error"
)

// messages
const (
	successMsg = "success"
)

type message struct {
	Message string `json:"message"`
}

func newMessage(msg string) message {
	return message{Message: msg}
}

type token struct {
	Token string `json:"token"`
}

func newToken(accessToken string) token {
	return token{Token: accessToken}
}
