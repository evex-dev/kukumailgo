package kukumailgo

func (s *Session) designateEndpoint(domain, name string) string {
	return "https://m.kuku.lu/index.php?action=addMailAddrByManual&by_system=1&csrf_token_check=" + s.CsrfToken + "&newdomain=" + domain + "&newuser=" + name
}

func (s *Session) normalEndpoint() string {
	return "https://m.kuku.lu/index.php?action=addMailAddrByAuto&by_system=1"
}

func (s *Session) expirationEndpoint() string {
	return "https://m.kuku.lu/index.php?action=addMailAddrByOnetime&by_system=1&csrf_token_check=" + s.CsrfToken
}

func (s *Session) searchReceiveEndpoint(content string) string {
	return "https://m.kuku.lu/recv._ajax.php?&q=" + content + "&csrf_token_check=" + s.CsrfToken
}

func (s *Session) getReceiveEndpoint() string {
	return "https://m.kuku.lu/recv._ajax.php?&csrf_token_check=" + s.CsrfToken
}

func (s *Session) getEmailContentEndpoint() string {
	return "https://m.kuku.lu/smphone.app.recv.view.php"
}
