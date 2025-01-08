package parser

import "fmt"

type ParseError struct {
	Pos     int
	Message string
	Token   Token
}

func (e ParseError) Error() string {
	return fmt.Sprintf("%s at position %d", e.Message, e.Pos)
}

func New(message string, token Token) *ParseError {
	return &ParseError{
		Pos:     token.Pos,
		Message: message,
		Token:   token,
	}
}

func Parse(input string) (*StatementNode, error) {
	p := newParser(input)
	flag, err := p.parseFlag()
	if err != nil {
		return nil, err
	}
	params, err := p.parseParams()
	if err != nil {
		return nil, err
	}
	return &StatementNode{
		Flag:   flag,
		Params: params,
	}, nil
}

type parser struct {
	scanner scanner
}

func newParser(input string) parser {
	p := parser{
		scanner: newScanner(input),
	}
	return p
}

func (p *parser) parseFlag() (*FlagNode, error) {
	token, err := p.nextToken()
	if err != nil {
		return nil, err
	}
	if token.Type != FLAG {
		return nil, New("Expected flag, got something else", *token)
	}
	return &FlagNode{
		Literal: token.Literal,
	}, nil
}

func (p *parser) parseParams() ([]ParamNode, error) {
	params := []ParamNode{}

	for {
		token, err := p.nextToken()
		if err != nil {
			return nil, err
		}
		if token.Type == EOF {
			break
		}

		switch token.Type {
		case POSITIONAL_PARAM:
			if err := assertPositionalBeforeOptionalParam(params, *token); err != nil {
				return nil, err
			}
			params = append(params, &PositionalParamNode{Literal: token.Literal})
		case OPTIONAL_PARAM:
			params = append(params, &OptionalParamNode{Literal: token.Literal})
		case FLAG_PARAM:
			params = append(params, &FlagParamNode{Literal: token.Literal})
		case POSITIONAL_PARAM_LIST:
			if err := assertPositionalBeforeOptionalParam(params, *token); err != nil {
				return nil, err
			}
			params = append(params, &PositionalParamListNode{Literal: token.Literal})
		}
	}

	return params, nil
}

func (p *parser) nextToken() (*Token, error) {
	return p.scanner.NextToken()
}

func assertPositionalBeforeOptionalParam(params []ParamNode, token Token) error {
	if len(params) > 0 {
		if _, ok := params[len(params)-1].(*OptionalParamNode); ok {
			return New("Positional parameter cannot be after optional parameter", token)
		}
	}
	return nil
}
