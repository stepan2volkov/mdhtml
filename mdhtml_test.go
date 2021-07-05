package mdhtml

import "testing"

func TestHeadingTrunslate(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{"Heading 1", []string{"# Heading 1"}, "<h1>Heading 1</h1>"},
		{"Heading 2", []string{"## Heading 2"}, "<h2>Heading 2</h2>"},
		{"Heading 3", []string{"### Heading 3"}, "<h3>Heading 3</h3>"},
		{"Heading 4", []string{"#### Heading 4"}, "<h4>Heading 4</h4>"},
		{"Heading 5", []string{"##### Heading 5"}, "<h5>Heading 5</h5>"},
		{"Heading 6", []string{"###### Heading 6"}, "<h6>Heading 6</h6>"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Trunslate(tc.input)
			if err != nil {
				t.Fatalf("error raised: %v", err)
			}
			if got != tc.want {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})

	}
}
