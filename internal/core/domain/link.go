package domain

import (
	"fmt"
	"net/http"
	"regexp"
	"sync"
	"time"
)

var (
	domain_regexp = regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)
	timeout       = time.Duration(10 * time.Second)
)
var wg sync.WaitGroup

type Link struct {
	ExternalLink     int `json:"external_links_count"`
	InternalLink     int `json:"internal_links_count"`
	InAccessibleLink int `json:"inaccessible_links_count"`
}

type InAccessible struct {
	Link       string
	Status     bool
	StatusCode int
}

func (node *HtmlNode) CheckLinks() (*Link, error) {
	var internalLink, externalLink, inAccessibleLink []string
	links := node.NodeParse("a", "href")

	for _, link := range links {
		if domain_regexp.MatchString(link) {
			externalLink = append(externalLink, link)
		} else {
			//fixed internal link
			website_url := node.Url
			internalLink = append(internalLink, website_url+link)
		}
	}
	all := append(externalLink, internalLink...)
	c := isAccessiable(all)
	for m := range c {
		fmt.Printf("Link %s -> status %s \n",
			m.Link, http.StatusText(m.StatusCode))
		if !m.Status {
			inAccessibleLink = append(inAccessibleLink,
				m.Link)
		}
	}
	return &Link{
		ExternalLink:     len(externalLink),
		InternalLink:     len(internalLink),
		InAccessibleLink: len(inAccessibleLink),
	}, nil
}

func isAccessiable(links []string) chan *InAccessible {
	c := make(chan *InAccessible)
	for _, link := range links {
		wg.Add(1)
		go visitLink(link, c)
	}
	go func() {
		wg.Wait()
		close(c)
	}()
	return c
}

func visitLink(link string, c chan *InAccessible) {
	client := &http.Client{
		Timeout: timeout,
	}
	defer wg.Done()
	resp, err := client.Head(link)
	if err == nil && resp.StatusCode == http.StatusOK {
		c <- &InAccessible{
			Link:       link,
			Status:     true,
			StatusCode: resp.StatusCode,
		}
	} else {
		var sCode int
		if err != nil {
			sCode = http.StatusNotFound
		} else {
			sCode = resp.StatusCode
		}
		c <- &InAccessible{
			Link:       link,
			Status:     false,
			StatusCode: sCode,
		}
	}
}
