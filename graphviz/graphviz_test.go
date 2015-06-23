package graphviz

import (
	"testing"
)

func TestDrawSVGSuccess(t *testing.T) {
	if !Installed() {
		t.Skip("graphviz not installed")
	}
	svg, err := DrawSVG("digraph {}")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf(svg)
}

func TestDrawSVGError(t *testing.T) {
	if !Installed() {
		t.Skip("graphviz not installed")
	}
	_, err := DrawSVG("digraph }")
	if err == nil {
		t.Fatalf("expected error")
	}
	t.Log(err)
}
