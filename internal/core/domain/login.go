package domain

func (t *HtmlNode) CheckLoginForm() bool {
	var login bool

	//input type submit and password
	form := t.NodeParse("input", "type")
	for _, t := range form {
		if t == "password" || t == "submit" {
			login = true
			break
		}
	}
	return login
}
