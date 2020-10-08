package params

import (
	"testing"
)

func TestPack(t *testing.T) {
	s := struct {
		Class string `http:"c"`
		Year  int    `http:"y"`
	}{"Math", 1980}
	p, err := Pack(&s)
	if err != nil {
		t.Errorf("Pack(%#v): %s", s, err)
	}
	want := "c=Math&y=1980"
	got := p.RawQuery
	if got != want {
		t.Errorf("Pack(%#v): got %q, want %q", s, got, want)
	}
}
