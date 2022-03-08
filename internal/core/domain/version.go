package domain

import (
	"strings"

	"golang.org/x/net/html"
)

const (
	HTML_1   = "HTML 1.0"
	HTML_2   = "HTML 2.0"
	HTML_3   = "HTML 3.2"
	HTML_4   = "HTML 4.0"
	HTML_4_1 = "HTML 4.01"
	XHTML_1  = "XHTML 1.0"
	XHTML_2  = "XHTML 1.1"
	XHTML_3  = "XHTML 2.0"
	HTML5    = "HTML 5.0"
)

func htmlVersion(att []html.Attribute) string {
	for _, a := range att {
		if strings.Contains(a.Val, XHTML_1) {
			return XHTML_1
		} else if strings.Contains(a.Val, XHTML_2) {
			return XHTML_2
		} else if strings.Contains(a.Val, XHTML_3) {
			return XHTML_3
		} else if strings.Contains(a.Val, HTML_4_1) {
			return HTML_4_1
		} else if strings.Contains(a.Val, HTML_1) {
			return HTML_1
		} else if strings.Contains(a.Val, HTML_2) {
			return HTML_2
		} else if strings.Contains(a.Val, HTML_3) {
			return HTML_3
		} else if strings.Contains(a.Val, HTML_4) {
			return HTML_4
		}
	}
	return HTML5
}
