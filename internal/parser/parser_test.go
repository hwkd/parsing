package parser

import (
	"testing"
)

func TestLegalSyntax(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  *StatementNode
	}{
		{
			"single flag",
			"-l",
			&StatementNode{
				Flag: &FlagNode{
					Literal: "-l",
				},
				Params: nil,
			},
		},
		{
			"flag with single positional parameter",
			"-sp <id>",
			&StatementNode{
				Flag: &FlagNode{
					Literal: "-sp",
				},
				Params: &ParamNode{
					PositionalParams: []PositionalParamNode{
						{Literal: "<id>"},
					},
				},
			},
		},
		{
			"flag with positional and optional param",
			"-a <title> [description]",
			&StatementNode{
				Flag: &FlagNode{
					Literal: "-a",
				},
				Params: &ParamNode{
					PositionalParams: []PositionalParamNode{
						{Literal: "<title>"},
					},
					OptionalParams: []OptionalParamNode{
						{Literal: "[description]"},
					},
				},
			},
		},
		{
			"flag with flag params",
			"-u [-t title] [-d description]",
			&StatementNode{
				Flag: &FlagNode{
					Literal: "-u",
				},
				Params: &ParamNode{
					FlagParams: []FlagParamNode{
						{Literal: "[-t title]"},
						{Literal: "[-d description]"},
					},
				},
			},
		},
		{
			"flag with positional parameter list",
			"-d <id>...",
			&StatementNode{
				Flag: &FlagNode{
					Literal: "-d",
				},
				Params: &ParamNode{
					PositionalParamList: &PositionalParamListNode{
						Literal: "<id>...",
					},
				},
			},
		},
		{
			"flag with flag parameter and positional parameter list",
			"-x [-t title] [-d description] <id>...",
			&StatementNode{
				Flag: &FlagNode{
					Literal: "-x",
				},
				Params: &ParamNode{
					FlagParams: []FlagParamNode{
						{Literal: "[-t title]"},
						{Literal: "[-d description]"},
					},
					PositionalParamList: &PositionalParamListNode{
						Literal: "<id>...",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.input)
			if err != nil {
				t.Fatalf("Expected nil, got '%v'", err)
			}
			if got.Flag.Literal != tt.want.Flag.Literal {
				t.Errorf("Expected '%s', got '%s'", tt.want.Flag.Literal, got.Flag.Literal)
			}
			if tt.want.Params != nil {
				a := tt.want.Params.PositionalParamList
				b := got.Params.PositionalParamList

				if a != nil && b == nil || a == nil && b != nil {
					t.Fatalf("Expected '%v', got '%v'", a, b)
				}

				if a != nil && b != nil {
					if a.Literal != b.Literal {
						t.Errorf("Expected '%s', got '%s'", a.Literal, b.Literal)
					}
				}

				if tt.want.Params.PositionalParams != nil {
					a := tt.want.Params.PositionalParams
					b := got.Params.PositionalParams
					if len(a) != len(b) {
						t.Fatalf("Expected '%d' positional params, got '%d'", len(a), len(b))
					}
					for i, param := range b {
						if param.Literal != a[i].Literal {
							t.Errorf("Expected '%s', got '%s'", a[i].Literal, param.Literal)
						}
					}
				}

				if tt.want.Params.OptionalParams != nil {
					a := tt.want.Params.OptionalParams
					b := got.Params.OptionalParams
					if len(a) != len(b) {
						t.Fatalf("Expected '%d' optional params, got '%d'", len(a), len(b))
					}
					for i, param := range b {
						if param.Literal != a[i].Literal {
							t.Errorf("Expected '%s', got '%s'", a[i].Literal, param.Literal)
						}
					}
				}

				if tt.want.Params.FlagParams != nil {
					a := tt.want.Params.FlagParams
					b := got.Params.FlagParams
					if len(a) != len(b) {
						t.Fatalf("Expected '%d' flag params, got '%d'", len(a), len(b))
					}
					for i, param := range b {
						if param.Literal != a[i].Literal {
							t.Errorf("Expected '%s', got '%s'", a[i].Literal, param.Literal)
						}
					}
				}
			}
		})
	}
}

func TestIllegalSyntax(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			"only positional parameter with no flag",
			"<id>",
		},
		{
			"positinal parameter after optional parameter",
			"-a [description] <id>",
		},
		{
			"positional parameter list after optional parameter",
			"-b [description] <id>...",
		},
		{
			"flag parameter after optional parameter",
			"-c [description] [-t title]",
		},
		{
			"flag parameter after positional parameter",
			"-d <id> [-t title]",
		},
		{
			"flag parameter after positional parameter list",
			"-e <id>... [-t title]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse(tt.input)
			if err == nil {
				t.Fatalf("Expected error, got nil")
			}
		})
	}
}
