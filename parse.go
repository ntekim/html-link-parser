package api

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
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
	for _, node := range nodes {
		fmt.Println(node)
	}

	//dfs(doc, "")
	return nil, nil
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
