package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="..." /a>) in an HTML document.
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and will return a
// slice of links, parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	dfs := &dfs{}
	dfs.process(doc)

	return dfs.links, nil
}

type dfs struct {
	links []Link
	text  string
}

func (dfs *dfs) newLink(n *html.Node) Link {
	var l Link
	for _, v := range n.Attr {
		if v.Key == "href" {
			l.Href = v.Val
			break
		}
	}
	l.Text = strings.Join(strings.Fields(dfs.text), " ")
	dfs.text = ""
	return l
}

func (dfs *dfs) process(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		dfs.processText(n.FirstChild)
		dfs.links = append(dfs.links, dfs.newLink(n))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs.process(c)
	}
}

func (dfs *dfs) processText(n *html.Node) {
	if n.Type == html.TextNode {
		dfs.text += n.Data
	}
	c := n.FirstChild
	if c == nil {
		c = n.NextSibling
	}
	for ; c != nil; c = c.NextSibling {
		dfs.processText(c)
	}
}
