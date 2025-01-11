package parser

import "strings"

type Node interface {
}

type StatementNode struct {
	Flag   *FlagNode
	Params *ParamNode
}

type FlagNode struct {
	Literal string
	Fname   string
}

func newFlagNode(literal string) FlagNode {
	return FlagNode{
		Literal: literal,
		Fname:   literal[1:],
	}
}

type ParamNode struct {
	FlagParams          []FlagParamNode
	PositionalParams    []PositionalParamNode
	OptionalParams      []OptionalParamNode
	PositionalParamList *PositionalParamListNode
}

// PositionalParamNode
type PositionalParamNode struct {
	Literal string
	Pname   string
}

func newPositionalParamNode(literal string) PositionalParamNode {
	return PositionalParamNode{
		Literal: literal,
		Pname:   literal[1 : len(literal)-1],
	}
}

// PositionalParamListNode
type PositionalParamListNode struct {
	Literal string
	Pname   string
}

func newPositionalParamListNode(literal string) PositionalParamListNode {
	return PositionalParamListNode{
		Literal: literal,
		Pname:   literal[1 : len(literal)-4],
	}
}

// OptionalParamNode
type OptionalParamNode struct {
	Literal string
	Pname   string
}

func newOptionalParamNode(literal string) OptionalParamNode {
	return OptionalParamNode{
		Literal: literal,
		Pname:   literal[1 : len(literal)-1],
	}
}

// FlagParamNode
type FlagParamNode struct {
	Flag    *FlagNode
	Literal string
	Pname   string
}

func newFlagParamNode(literal string) FlagParamNode {
	i := strings.Index(literal, " ")
	var flag FlagNode
	var pname string
	if i == -1 {
		flag = newFlagNode(literal)
	} else {
		flag = newFlagNode(literal[:i])
		pname = literal[len(flag.Literal)+1 : len(literal)-1]
	}
	return FlagParamNode{
		Flag:    &flag,
		Literal: literal,
		Pname:   pname,
	}
}
