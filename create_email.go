package kukumailgo

import (
	"errors"
	"strings"
)

func (s *Session) CreateDesignateEmail(name, domain string) (err error) {
	_, err = s.createEmail(s.designateEndpoint(domain, name))
	return err
}

func (s *Session) CreateRandomEmail() (string, error) {
	return s.createEmail(s.normalEndpoint())
}

func (s *Session) createEmail(url string) (string, error) {
	r, err := s.get(url, nil)
	if err != nil {
		return "", err
	}
	body, err := readBody(r)
	if err != nil {
		return "", err
	}

	bodyString := string(body)

	if !strings.HasPrefix(bodyString, "OK:") {
		return "", errors.New("unknown error: " + bodyString)
	}

	return strings.TrimPrefix(bodyString, "OK: "), nil
}
