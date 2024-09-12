package kukumailgo

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
	"strings"
)

func (s *Session) GetAllReceivedEmail() ([]Mail, error) {
	return s.getEmails(s.getReceiveEndpoint())
}

func (s *Session) SearchReceivedEmail(content string) ([]Mail, error) {
	return s.getEmails(s.searchReceiveEndpoint(content))
}

func (s *Session) getEmails(url string) ([]Mail, error) {
	r, err := s.get(url, nil)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		return nil, err
	}

	content := doc.Find("script").Text()

	re := regexp.MustCompile(`openMailData\('(\d+)', '(\w+)'`)

	matches := re.FindAllStringSubmatch(content, -1)

	var emails []Mail

	for _, match := range matches {
		emails = append(emails, Mail{
			Num: match[1],
			Key: match[2],
		})
	}

	return emails, nil
}

func (s *Session) GetEmailContent(mail Mail) ([]byte, error) {
	r, err := s.post(s.getEmailContentEndpoint(), strings.NewReader(fmt.Sprintf("num=%s&key=%s&noscroll=1", mail.Num, mail.Key)), http.Header{"Content-Type": {"application/x-www-form-urlencoded"}})
	if err != nil {
		return nil, err
	}
	return readBody(r)
}
