package parser

import "testing"

func TestFlagNode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  FlagNode
	}{
		{
			"alphabet character flag",
			"-a",
			FlagNode{Literal: "-a", Fname: "a"},
		},
		{
			"numeric character flag",
			"-0",
			FlagNode{Literal: "-0", Fname: "0"},
		},
		{
			"multi-character flag",
			"-add",
			FlagNode{Literal: "-add", Fname: "add"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := newFlagNode(tt.input)
			if n.Literal != tt.want.Literal {
				t.Errorf("Expected %s, got %s", tt.want.Literal, n.Literal)
			}
			if n.Fname != tt.want.Fname {
				t.Errorf("Expected %s, got %s", tt.want.Fname, n.Fname)
			}
		})
	}
}

func TestPositionalParamNode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  PositionalParamNode
	}{
		{
			"param",
			"<title>",
			PositionalParamNode{Literal: "<title>", Pname: "title"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := newPositionalParamNode(tt.input)
			if n.Literal != tt.want.Literal {
				t.Errorf("Expected %s, got %s", tt.want.Literal, n.Literal)
			}
			if n.Pname != tt.want.Pname {
				t.Errorf("Expected %s, got %s", tt.want.Pname, n.Pname)
			}
		})
	}
}

func TestPositionalParamListNode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  PositionalParamListNode
	}{
		{
			"param list",
			"<title>...",
			PositionalParamListNode{Literal: "<title>...", Pname: "title"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := newPositionalParamListNode(tt.input)
			if n.Literal != tt.want.Literal {
				t.Errorf("Expected %s, got %s", tt.want.Literal, n.Literal)
			}
			if n.Pname != tt.want.Pname {
				t.Errorf("Expected %s, got %s", tt.want.Pname, n.Pname)
			}
		})
	}
}

func TestOptionalParamNode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  OptionalParamNode
	}{
		{
			"param",
			"[description]",
			OptionalParamNode{Literal: "[description]", Pname: "description"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := newOptionalParamNode(tt.input)
			if n.Literal != tt.want.Literal {
				t.Errorf("Expected %s, got %s", tt.want.Literal, n.Literal)
			}
			if n.Pname != tt.want.Pname {
				t.Errorf("Expected %s, got %s", tt.want.Pname, n.Pname)
			}
		})
	}
}

func TestFlagParamNode(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  FlagParamNode
	}{
		{
			"Param with pname",
			"[-t title]",
			FlagParamNode{Literal: "[-t title]", Pname: "title"},
		},
		{
			"Param without pname",
			"[-t]",
			FlagParamNode{Literal: "[-t]", Pname: ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := newFlagParamNode(tt.input)
			if n.Literal != tt.want.Literal {
				t.Errorf("Expected %s, got %s", tt.want.Literal, n.Literal)
			}
			if n.Pname != tt.want.Pname {
				t.Errorf("Expected %s, got %s", tt.want.Pname, n.Pname)
			}
		})
	}
}
