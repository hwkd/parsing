package parser

import "testing"

func TestScanner(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Token
	}{
		{
			"single flag",
			"-l",
			[]Token{
				{FLAG, "-l", 0},
				{EOF, "", 2},
			},
		},
		{
			"flag with single positional parameter",
			"-sp <id>",
			[]Token{
				{FLAG, "-sp", 0},
				{POSITIONAL_PARAM, "<id>", 4},
				{EOF, "", 8},
			},
		},
		{
			"flag with flag params",
			"-x [-y] [-z]",
			[]Token{
				{FLAG, "-x", 0},
				{FLAG_PARAM, "[-y]", 3},
				{FLAG_PARAM, "[-z]", 8},
				{EOF, "", 12},
			},
		},
		{
			"flag with positional and optional param",
			"-a <title> [description]",
			[]Token{
				{FLAG, "-a", 0},
				{POSITIONAL_PARAM, "<title>", 3},
				{OPTIONAL_PARAM, "[description]", 11},
				{EOF, "", 24},
			},
		},
		{
			"flag with flag params",
			"-u [-t title] [-d description]",
			[]Token{
				{FLAG, "-u", 0},
				{FLAG_PARAM, "[-t title]", 3},
				{FLAG_PARAM, "[-d description]", 14},
				{EOF, "", 30},
			},
		},
		{
			"flag with positional parameter list",
			"-d <id>...",
			[]Token{
				{FLAG, "-d", 0},
				{POSITIONAL_PARAM_LIST, "<id>...", 3},
				{EOF, "", 10},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newScanner(tt.input)
			for i := 0; ; i++ {
				tok, err := s.NextToken()
				if err != nil {
					t.Fatalf("Expected nil, got '%s' while parsing for '%s'", err, tt.input)
				}
				want := tt.want[i]
				if tok.Type != want.Type {
					t.Fatalf("Expected '%s', got '%s'", want.Type, tok.Type)
				}
				if tok.Literal != want.Literal {
					t.Fatalf("Expected '%s', got '%s'", want.Literal, tok.Literal)
				}
				if tok.Pos != want.Pos {
					t.Fatalf("Expected '%d', got '%d'", want.Pos, tok.Pos)
				}
				if tok.Type == EOF {
					break
				}
			}
		})
	}
}
