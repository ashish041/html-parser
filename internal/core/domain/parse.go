package domain

import "golang.org/x/net/html"

type HtmlParse interface {
	NodeParse(tag string, att string) []string
	TokenParse(tag string) (data []string)
	CheckLinks() (*Link, error)
	CheckLoginForm() bool
}

type HtmlNode struct {
	Tree  *html.Node
	Token *html.Tokenizer
	Url   string
}

func NewParse(url *Url) (HtmlParse, error) {
	node, err := html.Parse(url.ResponseBody)
	if err != nil {
		return nil, err
	}
	return &HtmlNode{
		Tree: node,
		Url:  url.Url,
	}, nil
}

func (node *HtmlNode) NodeParse(tag string, att string) []string {
	var vals []string
	var f func(*html.Node)

	f = func(n *html.Node) {
		switch n.Type {
		case html.TextNode:
			if n.Parent.Data == tag && att == "" {
				vals = append(vals, n.Data)
			}
		case html.ElementNode:
			if n.Data == tag {
				for _, a := range n.Attr {
					if a.Key == att {
						vals = append(vals, a.Val)
						break
					}
				}
			}
		case html.DoctypeNode:
			if tag == "DOCTYPE" {
				v := htmlVersion(n.Attr)
				vals = append(vals, v)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(node.Tree)

	return vals
}

func (node *HtmlNode) TokenParse(tag string) (data []string) {
	var vals []string
	var isTag bool
	for {
		tt := node.Token.Next()
		switch {
		case tt == html.ErrorToken:
			return vals
		case tt == html.StartTagToken:
			t := node.Token.Token()
			isTag = t.Data == tag
		case tt == html.TextToken:
			t := node.Token.Token()
			if isTag {
				vals = append(vals, t.Data)
			}
			isTag = false
		}
	}
}
