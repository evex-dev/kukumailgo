package kukumailgo

import "net/http"

type Session struct {
	SessionHash string
	client      *http.Client
	CsrfToken   string
}

type Mail struct {
	Num string
	Key string
}
