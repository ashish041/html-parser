package domain_test

import (
	"strings"
	"testing"

	"github.com/ashish041/html-parser/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func Test_checkLinks(t *testing.T) {
	s := `<p>Links:</p><ul>
		<li><a href="https://google.com">Foo</a>
		<li><a href="/bar/baz">BarBaz</a>
	</ul>`

	node, err := html.Parse(strings.NewReader(s))
	assert.Equal(t, err, nil)

	p := &domain.HtmlNode{
		Tree: node,
	}
	link, err := p.CheckLinks()
	assert.Equal(t, err, nil)

	assert.Equal(t, link.InternalLink, 1)
	assert.Equal(t, link.ExternalLink, 1)
	assert.Equal(t, link.InAccessibleLink, 1)
}

func Test_checkLoginForm(t *testing.T) {
	s := `<html>
	<head>
		<title></title>
	</head>
		<body>
			<form action="/login" method="post">
				Username:<input type="text" name="username">
				Password:<input type="password" name="password">
				<input type="submit" value="Login">
			</form>
		</body>
	</html>`

	node, err := html.Parse(strings.NewReader(s))
	assert.Equal(t, err, nil)

	p := &domain.HtmlNode{
		Tree: node,
	}
	exist := p.CheckLoginForm()

	assert.True(t, exist)
}

func TestBoard_checkHeadings(t *testing.T) {
	s := `<html>
		<head>
			<title>Title</title>
		</head>
		<body>
			<h1>Heading 1</h1>
			<h2>Heading 2</h2>
			<h1>Heading 1</h1>
			<h4>Heading 4</h4>
			<h5>Heading 5</h5>
			<h6>Heading 6</h6>
			<h5>Heading 5</h5>
		</body>
	</html>`

	node, err := html.Parse(strings.NewReader(s))
	assert.Equal(t, err, nil)

	p := &domain.HtmlNode{
		Tree: node,
	}
	h1 := len(p.NodeParse("h1", ""))
	h2 := len(p.NodeParse("h2", ""))
	h3 := len(p.NodeParse("h3", ""))
	h4 := len(p.NodeParse("h4", ""))
	h5 := len(p.NodeParse("h5", ""))
	h6 := len(p.NodeParse("h6", ""))

	assert.Equal(t, h1, 2)
	assert.Equal(t, h2, 1)
	assert.Equal(t, h3, 0)
	assert.Equal(t, h4, 1)
	assert.Equal(t, h5, 2)
	assert.Equal(t, h6, 1)
}

func Test_htmlTitle(t *testing.T) {
	s := `<!DOCTYPE html>
	<html>
		<title>The title Attribute</title>
		<body>
			<h2 title="I'm a header">The title Attribute</h2>
			<p title="I'm a tooltip">Mouse over this paragraph,
			to display the title attribute as a tooltip.</p>
		</body>
	</html>`

	node, err := html.Parse(strings.NewReader(s))
	assert.Equal(t, err, nil)

	p := &domain.HtmlNode{
		Tree: node,
	}
	title := strings.Join(p.NodeParse("title", ""), "")

	assert.Equal(t, title, "The title Attribute")
}

func Test_checkhtmlVersion(t *testing.T) {
	s := `<!DOCTYPE html>
	<html>
		<title>World Health Organization </title>
		<body>
			<h2 title="I'm a header">The title Attribute</h2>
			<p>This is a paragraph.</p>
		</body>
	</html>`

	node, err := html.Parse(strings.NewReader(s))
	assert.Equal(t, err, nil)

	p := &domain.HtmlNode{
		Tree: node,
	}
	version := strings.Join(p.NodeParse(
		"DOCTYPE", ""), "")

	assert.Equal(t, version, "HTML 5.0")
}

func Test_checkhtmlVersion2(t *testing.T) {
	s := `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
	"http://www.w3.org/TR/html4/loose.dtd">
	<html>
		<title>World Health Organization </title>
		<body>
			<h2 title="I'm a header">The title Attribute</h2>
			<p>This is a paragraph.</p>
		</body>
	</html>`

	node, err := html.Parse(strings.NewReader(s))
	assert.Equal(t, err, nil)

	p := &domain.HtmlNode{
		Tree: node,
	}
	version := strings.Join(p.NodeParse(
		"DOCTYPE", ""), "")

	assert.Equal(t, version, "HTML 4.01")
}
