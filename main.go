package html2json

import (
	"encoding/json"
	"strings"

	"golang.org/x/net/html"
)

type HTMLElement struct {
	Tag      string            `json:"tag"`
	Attrs    map[string]string `json:"attrs"`
	Children []*HTMLElement    `json:"children"`
}

func parseNode(node *html.Node) *HTMLElement {
	if node.Type == html.DocumentNode {
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			return parseNode(child)
		}
		return nil
	} else if node.Type == html.ElementNode {
		element := &HTMLElement{
			Tag:   node.Data,
			Attrs: make(map[string]string),
		}
		for _, attr := range node.Attr {
			if attr.Key == "id" || attr.Key == "class" || attr.Key == "style" {
				continue
			}
			element.Attrs[attr.Key] = attr.Val
		}
		for child := node.FirstChild; child != nil; child = child.NextSibling {
			element.Children = append(element.Children, parseNode(child))
		}
		return element
	} else if node.Type == html.TextNode {
		return &HTMLElement{
			Tag:      "text",
			Attrs:    map[string]string{"text": node.Data},
			Children: []*HTMLElement{{Tag: "", Attrs: nil, Children: nil}},
		}
	} else {
		return nil
	}
}

func ParseHTML(htmlString string) (string, error) {
	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		return "", err
	}
	element := parseNode(doc)

	jsonBytes, err := json.Marshal(element)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
