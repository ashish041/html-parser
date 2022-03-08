package domain

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Url struct {
	ResponseBody io.ReadCloser
	Url          string
}

type Headings struct {
	H1 int `json:"h1_headings_count"`
	H2 int `json:"h2_headings_count"`
	H3 int `json:"h3_headings_count"`
	H4 int `json:"h4_headings_count"`
	H5 int `json:"h5_headings_count"`
	H6 int `json:"h6_headings_count"`
}

type Informations struct {
	HtmlVersion string    `json:"html_version"`
	PageTitle   string    `json:"page_title"`
	Headings    *Headings `json:"headings"`
	Link        *Link     `json:"links"`
	LoginForm   bool      `json:"login_form"`
}

type DomainLogic struct {
}

func (d DomainLogic) New(url string) (*Url, error) {
	client := &http.Client{
		Timeout: timeout,
	}
	if !strings.HasPrefix(url, "https://") &&
		!strings.HasPrefix(url, "http://") {
		url = fmt.Sprintf("http://%s", url)
	}
	response, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return &Url{
		ResponseBody: response.Body,
		Url:          url,
	}, nil
}

func (d DomainLogic) Get(url *Url) (*Informations, error) {
	p, err := NewParse(url)
	if err != nil {
		return nil, err
	}
	linkCount, err := p.CheckLinks()
	if err != nil {
		return nil, err
	}
	return &Informations{
		HtmlVersion: strings.Join(p.NodeParse(
			"DOCTYPE", ""), ""),
		PageTitle: strings.Join(p.NodeParse(
			"title", ""), ""),
		Headings: &Headings{
			H1: len(p.NodeParse("h1", "")),
			H2: len(p.NodeParse("h2", "")),
			H3: len(p.NodeParse("h3", "")),
			H4: len(p.NodeParse("h4", "")),
			H5: len(p.NodeParse("h5", "")),
			H6: len(p.NodeParse("h6", "")),
		},
		Link:      linkCount,
		LoginForm: p.CheckLoginForm(),
	}, nil
}
