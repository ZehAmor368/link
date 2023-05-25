package link_test

import (
	link "link/pkg"
	"os"
	"strings"
	"testing"
)

func TestExtract(t *testing.T) {
	testCases := []struct {
		desc   string
		path   string
		expect []link.Link
	}{
		{
			desc:   "test with ex1.html",
			path:   "../ex1.html",
			expect: []link.Link{{Href: "/other-page", Text: "A link to another page"}},
		},
		{
			desc: "test with ex2.html",
			path: "../ex2.html",
			expect: []link.Link{
				{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"},
				{Href: "https://github.com/gophercises", Text: "Gophercises is on Github!"},
			},
		},
		{
			desc: "test with ex3.html",
			path: "../ex3.html",
			expect: []link.Link{
				{Href: "#", Text: "Login"},
				{Href: "/lost", Text: "Lost? Need help?"},
				{Href: "https://twitter.com/marcusolsson", Text: "@marcusolsson"},
			},
		},
		{
			desc:   "test with ex4.html",
			path:   "../ex4.html",
			expect: []link.Link{{Href: "/dog-cat", Text: "dog cat"}},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			f, err := os.Open(tC.path)
			if err != nil {
				t.Error("open file", err)
			}
			links, err := link.Parse(f)
			if err != nil {
				t.Error("extracting links", err)
			}
			if len(links) != len(tC.expect) {
				t.Error("expected", len(tC.expect), "links, got", len(links))
			}
			for i, link := range links {
				if strings.Compare(link.Href, tC.expect[i].Href) != 0 {
					t.Errorf("expected Href of link #%d to be equal to: %q, got: %q\n", i, tC.expect[i].Href, link.Href)
				}
				if strings.Compare(link.Text, tC.expect[i].Text) != 0 {
					t.Errorf("expected Text of link #%d to be equal to: %q, got: %q\n", i, tC.expect[i].Text, link.Text)
				}
			}
		})
	}
}
