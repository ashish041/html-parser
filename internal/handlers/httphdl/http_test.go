package httphdl_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/ashish041/html-parser/internal/core/domain"
	"github.com/ashish041/html-parser/internal/handlers/httphdl"
	"github.com/stretchr/testify/assert"
)

func Test_GetGeneralInformations(t *testing.T) {
	s := `<!DOCTYPE html>
	<html>
		<head>
			<title>Href Attribute Example</title>
		</head>
		<body>
			<h1>Href Attribute Example</h1>
			<h3>Href Attribute Example</h3>
			<h5>Href Attribute Example</h5>
			<p>
			<a href="https://www.freecodecamp.org/contribute/">
			The freeCodeCamp Contribution Page</a>
			shows you how and where you can contribute to freeCodeCamp's community and growth.
			<a href="/freecodecamp/contribute/">
			The freeCodeCamp Contribution Page 2</a>
			</p>
			<form>
				<label for="username">Username:</label>
				<input type="text" name="username" id="username" />
				<label for="password">Password:</label>
				<input type="password" name="password" id="password" />
				<input type="submit" value="Submit" />
			</form>
		</body>
	</html>`

	r := ioutil.NopCloser(strings.NewReader(s))
	url := &domain.Url{
		ResponseBody: r,
	}

	d := domain.DomainLogic{}
	hdl := httphdl.NewHTTPHandler(d)

	info, err := hdl.Domain.Get(url)
	assert.Equal(t, err, nil)

	if assert.NotNil(t, info) {
		assert.Equal(t, info.HtmlVersion, "HTML 5.0")
		assert.Equal(t, info.PageTitle, "Href Attribute Example")
		assert.Equal(t, info.Headings.H1, 1)
		assert.Equal(t, info.Headings.H3, 1)
		assert.Equal(t, info.Headings.H5, 1)
		assert.Equal(t, info.Link.ExternalLink, 1)
		assert.Equal(t, info.Link.InternalLink, 1)
		assert.Equal(t, info.Link.InAccessibleLink, 2)
		assert.Equal(t, info.LoginForm, true)
	}
}
