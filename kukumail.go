package kukumailgo

import (
	"net/http"
	"time"
)

func New(sessionHash, csrfToken string) *Session {
	s := &Session{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		SessionHash: sessionHash,
		CsrfToken:   csrfToken,
	}

	s.post("https://m.kuku.lu", nil)

	return s
}
