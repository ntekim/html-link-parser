package api

import (
	"golang.org/x/net/html"
	"io"
	"strings"
)

// Link represent a link <a href=""  /> in an html document
type Link struct {
	Href string
	Text string
}

// Parse takes a html document or a url and returns the slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := LinkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLinks(node))
	}

	return links, nil
}

func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type == html.CommentNode {
		return ""
	}

	if n.Type != html.ElementNode {
		return ""
	}

	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += getText(c) + " "
	}

	return strings.Join(strings.Fields(ret), " ")
}

func buildLinks(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	ret.Text = getText(n)

	return ret
}

func LinkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, LinkNodes(c)...)
	}

	return ret
}

//func dfs(n *html.Node, padding string) {
//	msg := n.Data
//	if n.Type == html.ElementNode {
//		msg = "<" + msg + ">"
//	}
//
//	fmt.Println(padding, msg)
//
//	for c := n.FirstChild; c != nil; c = c.NextSibling {
//		dfs(c, padding+"  ")
//	}
//}
