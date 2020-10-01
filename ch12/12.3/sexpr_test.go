package sexpr

import (
	"testing"
)

func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)
}

func TestBool(t *testing.T) {
	tests := []struct {
		v    bool
		want string
	}{
		{true, "t"},
		{false, "nil"},
	}
	for _, test := range tests {
		data, err := Marshal(test.v)
		if err != nil {
			t.Errorf("Marshal(%t): %s", test.v, err)
		}
		if string(data) != test.want {
			t.Errorf("Marshal(%t) got %s, wanted %s", test.v, data, test.want)
		}
	}
}

func TestFloat32(t *testing.T) {
	tests := []struct {
		v    float32
		want string
	}{
		{1.8e9, "1.8e+09"},
		{2.0, "2"},
		{0, "0"},
	}
	for _, test := range tests {
		data, err := Marshal(test.v)
		if err != nil {
			t.Errorf("Marshal(%f): %s", test.v, err)
		}
		if string(data) != test.want {
			t.Errorf("Marshal(%f) got %s, wanted %s", test.v, data, test.want)
		}
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		v    float64
		want string
	}{
		{1.8e9, "1.8e+09"},
		{2.0, "2"},
		{0, "0"},
	}
	for _, test := range tests {
		data, err := Marshal(test.v)
		if err != nil {
			t.Errorf("Marshal(%g): %s", test.v, err)
		}
		if string(data) != test.want {
			t.Errorf("Marshal(%g) got %s, wanted %s", test.v, data, test.want)
		}
	}
}

func TestComplex64(t *testing.T) {
	tests := []struct {
		v    complex64
		want string
	}{
		{1 + 2i, "#C(1 2)"},
		{3 - 4i, "#C(3 -4)"},
		{-1.8 + -3.6i, "#C(-1.8 -3.6)"},
	}
	for _, test := range tests {
		data, err := Marshal(test.v)
		if err != nil {
			t.Errorf("Marshal(%g): %s", test.v, err)
		}
		if string(data) != test.want {
			t.Errorf("Marshal(%g) got %s, wanted %s", test.v, data, test.want)
		}
	}
}

func TestComplex128(t *testing.T) {
	tests := []struct {
		v    complex128
		want string
	}{
		{1 + 2i, "#C(1 2)"},
		{3 - 4i, "#C(3 -4)"},
		{-1.8 + -3.6i, "#C(-1.8 -3.6)"},
	}
	for _, test := range tests {
		data, err := Marshal(test.v)
		if err != nil {
			t.Errorf("Marshal(%g): %s", test.v, err)
		}
		if string(data) != test.want {
			t.Errorf("Marshal(%g) got %s, wanted %s", test.v, data, test.want)
		}
	}
}

func TestInterface(t *testing.T) {
	type Interface interface{}
	type Wrapper struct {
		i Interface
	}
	w := Wrapper{3}
	data, err := Marshal(w)
	if err != nil {
		t.Errorf("Marshal(%v): %s", w, err)
	}
	want := `((i ("sexpr.Interface" 3)))`
	if string(data) != want {
		t.Errorf("Marshal(%v) got %s, wanted %s", w, data, want)
	}
}
