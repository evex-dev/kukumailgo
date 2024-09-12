package kukumailgo

import (
	"io"
	"net/http"
)

func (s *Session) get(url string, body io.Reader, header ...http.Header) (*http.Response, error) {
	return s.req("GET", url, body, header...)
}

func (s *Session) post(url string, body io.Reader, header ...http.Header) (*http.Response, error) {
	return s.req("POST", url, body, header...)
}

func (s *Session) req(method, url string, body io.Reader, header ...http.Header) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if len(header) > 0 {
		req.Header = header[0]
	}
	req.AddCookie(&http.Cookie{
		Name:  "cookie_csrf_token",
		Value: s.CsrfToken,
	})
	req.AddCookie(&http.Cookie{
		Name:  "cookie_sessionhash",
		Value: s.SessionHash,
	})
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func readBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
