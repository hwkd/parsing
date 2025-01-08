package parser

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
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
				Params: []ParamNode{
					&PositionalParamNode{
						Literal: "<id>",
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
				Params: []ParamNode{
					&PositionalParamNode{
						Literal: "<title>",
					},
					&OptionalParamNode{
						Literal: "[description]",
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
				Params: []ParamNode{
					&FlagParamNode{
						Literal: "[-t title]",
					},
					&FlagParamNode{
						Literal: "[-d description]",
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
				Params: []ParamNode{
					&PositionalParamListNode{
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
				Params: []ParamNode{
					&FlagParamNode{
						Literal: "[-t title]",
					},
					&FlagParamNode{
						Literal: "[-d description]",
					},
					&PositionalParamListNode{
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
			if len(got.Params) != len(tt.want.Params) {
				t.Fatalf("Expected '%d' params, got '%d'", len(tt.want.Params), len(got.Params))
			}
			fmt.Println(got.Params)
			for i, param := range got.Params {
				if param.GetLiteral() != tt.want.Params[i].GetLiteral() {
					t.Errorf("Expected '%s', got '%s'", tt.want.Params[i].GetLiteral(), param.GetLiteral())
				}
			}
		})
	}
}

func TestParserError(t *testing.T) {
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
			"[description] <id>",
		},
		{
			"positional parameter list after optional parameter",
			"[description] <id>...",
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
