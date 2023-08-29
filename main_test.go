package html2json

import (
	"strings"
	"testing"
)

func TestParseHTML(t *testing.T) {
	html := "<html><head></head><body><h1>Hello World</h1></body></html>"
	jsonString, err := ParseHTML(html)
	if err != nil {
		t.Error(err)
	}

	expected := `{"tag":"html","attrs":{},"children":[{"tag":"head","attrs":{},"children":null},{"tag":"body","attrs":{},"children":[{"tag":"h1","attrs":{},"children":[{"tag":"text","attrs":{"text":"Hello World"},"children":[{"tag":"","attrs":null,"children":null}]}]}]}]}`
	if jsonString != expected {
		t.Errorf("got %s, want %s", jsonString, expected)
	}
}

func TestParseHTML_StripsIDAttr(t *testing.T) {
	html := "<html><head></head><body><h1 id='fancy-heading'>Hello World</h1></body></html>"
	jsonString, err := ParseHTML(html)
	if err != nil {
		t.Error(err)
	}

	expected := `{"tag":"html","attrs":{},"children":[{"tag":"head","attrs":{},"children":null},{"tag":"body","attrs":{},"children":[{"tag":"h1","attrs":{},"children":[{"tag":"text","attrs":{"text":"Hello World"},"children":[{"tag":"","attrs":null,"children":null}]}]}]}]}`
	if strings.Contains(jsonString, "fancy-heading") {
		t.Errorf("got %s, want %s", jsonString, expected)
	}
}

func TestParseHTML_StripsClassAttr(t *testing.T) {
	html := "<html><head></head><body><h1>Hello World</h1><p class='description'>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p></body></html>"
	jsonString, err := ParseHTML(html)
	if err != nil {
		t.Error(err)
	}

	expected := `{"tag":"html","attrs":{},"children":[{"tag":"head","attrs":{},"children":null},{"tag":"body","attrs":{},"children":[{"tag":"h1","attrs":{},"children":[{"tag":"text","attrs":{"text":"Hello World"},"children":[{"tag":"","attrs":null,"children":null}]}]},{"tag":"p","attrs":{},"children":[{"tag":"text","attrs":{"text":"Lorem ipsum dolor sit amet, consectetur adipiscing elit."},"children":[{"tag":"","attrs":null,"children":null}]}]}]}]}`
	if strings.Contains(jsonString, "description") {
		t.Errorf("got %s, want %s", jsonString, expected)
	}
}

func TestParseHTML_StripsStyleAttr(t *testing.T) {
	html := "<html><head></head><body><div style='background: red;'><h1>Hello World</h1></div></body></html>"
	jsonString, err := ParseHTML(html)
	if err != nil {
		t.Error(err)
	}

	expected := `{"tag":"html","attrs":{},"children":[{"tag":"head","attrs":{},"children":null},{"tag":"body","attrs":{},"children":[{"tag":"div","attrs":{},"children":[{"tag":"h1","attrs":{},"children":[{"tag":"text","attrs":{"text":"Hello World"},"children":[{"tag":"","attrs":null,"children":null}]}]}]}]}]}`
	if strings.Contains(jsonString, "style") {
		t.Errorf("got %s, want %s", jsonString, expected)
	}
}
