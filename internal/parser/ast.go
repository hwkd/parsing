package parser

type Node interface {
}

type StatementNode struct {
	Flag   *FlagNode
	Params []ParamNode
}

type FlagNode struct {
	Literal string
}

type ParamNode interface {
	GetLiteral() string
	GetName() string
}

type PositionalParamNode struct {
	Literal string
}

func (p *PositionalParamNode) GetLiteral() string {
	return p.Literal
}

func (p *PositionalParamNode) GetName() string {
	return p.Literal[1 : len(p.Literal)-1]
}

type PositionalParamListNode struct {
	Literal string
}

func (p *PositionalParamListNode) GetLiteral() string {
	return p.Literal
}

func (p *PositionalParamListNode) GetName() string {
	return p.Literal[1 : len(p.Literal)-1]
}

type OptionalParamNode struct {
	Literal string
}

func (o *OptionalParamNode) GetLiteral() string {
	return o.Literal
}

func (o *OptionalParamNode) GetName() string {
	return o.Literal[1 : len(o.Literal)-1]
}

type FlagParamNode struct {
	Flag    *FlagNode
	Literal string
}

func (f *FlagParamNode) GetLiteral() string {
	return f.Literal
}

func (f *FlagParamNode) GetName() string {
	return f.Literal[1 : len(f.Literal)-1]
}
