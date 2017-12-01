package sparkling

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

// TestNew verifies that the returned instance is of the proper type
func TestSparkling_New(t *testing.T) {
	sp := New(os.Stdout)
	if reflect.TypeOf(sp).String() != "*sparkling.Sparkling" {
		t.Error("New returned incorrect type")
	}
}

// TestSparkling_Render verifies that sparkling renders as expected
func TestSparkling_Render(t *testing.T) {
	var buf bytes.Buffer
	sp := New(&buf)
	sp.AddSeries([]float64{0, 30, 55, 80, 33, 150}, "Awesome")
	sp.Render()

	want := "Awesome ▁▂▃▄▂█\n"

	s := buf.String()
	got := s[len(s)-len(want):]

	if got != want {
		t.Errorf("sparkling.Render() = %s, want: %s", got, want)
	}
}

// TestSparkling_RenderFloats verifies that sparkling renders as expected
func TestSparkling_RenderFloats(t *testing.T) {
	var buf bytes.Buffer
	sp := New(&buf)
	sp.AddSeries([]float64{0.3, 1.9, 4.5, 2.3, 1.9, 0.8}, "")
	sp.Render()

	want := "▁▂█▄▂▁\n"

	s := buf.String()
	got := s[len(s)-len(want):]

	if got != want {
		t.Errorf("sparkling.Render() = %s, want: %s", got, want)
	}
}

// TestNewSeries verifies that the returned instance is of the proper type
func TestSeries_NewSeries(t *testing.T) {
	s := NewSeries([]float64{1, 2, 3}, "Test")
	if reflect.TypeOf(s).String() != "*sparkling.Series" {
		t.Error("New returned incorrect type")
	}
}
