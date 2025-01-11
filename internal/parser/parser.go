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

func NewParserError(message string, token Token) *ParseError {
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
		return nil, NewParserError("Expected flag, got something else", *token)
	}
	return &FlagNode{
		Literal: token.Literal,
	}, nil
}

func (p *parser) parseParams() (*ParamNode, error) {
	params := &ParamNode{}

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
			err := assertNoPriorOptionalParam(
				*params,
				*token,
				"Positional parameter cannot be after an optional parameter",
			)
			if err != nil {
				return nil, err
			}
			params.PositionalParams = append(
				params.PositionalParams,
				newPositionalParamNode(token.Literal),
			)
		case OPTIONAL_PARAM:
			err := assertNoPriorPositionalParamList(
				*params,
				*token,
				"Optional parameter cannot be after a positional parameter list",
			)
			if err != nil {
				return nil, err
			}
			node := OptionalParamNode{Literal: token.Literal}
			params.OptionalParams = append(params.OptionalParams, node)
		case FLAG_PARAM:
			err := assertNoPriorParams(
				*params,
				*token,
				"Flag parameter cannot be after a positional parameter",
			)
			if err != nil {
				return nil, err
			}
			node := FlagParamNode{Literal: token.Literal}
			params.FlagParams = append(params.FlagParams, node)
		case POSITIONAL_PARAM_LIST:
			err := assertNoPriorOptionalParam(
				*params,
				*token,
				"Positional parameter list cannot be after an optional parameter",
			)
			if err != nil {
				return nil, err
			}
			node := &PositionalParamListNode{Literal: token.Literal}
			params.PositionalParamList = node
		}
	}

	return params, nil
}

func (p *parser) nextToken() (*Token, error) {
	return p.scanner.NextToken()
}

func assertNoPriorPositionalParams(params ParamNode, token Token, msg string) error {
	if len(params.PositionalParams) > 0 || len(params.PositionalParamList.Literal) > 0 {
		return NewParserError(msg, token)
	}
	return nil
}

func assertNoPriorPositionalParamList(params ParamNode, token Token, msg string) error {
	if params.PositionalParamList != nil {
		return NewParserError(msg, token)
	}
	return nil
}

func assertNoPriorOptionalParam(params ParamNode, token Token, msg string) error {
	if len(params.OptionalParams) > 0 {
		return NewParserError(msg, token)
	}
	return nil
}

func assertNoPriorParams(params ParamNode, token Token, msg string) error {
	if len(params.PositionalParams) > 0 || len(params.OptionalParams) > 0 || params.PositionalParamList != nil {
		return NewParserError(msg, token)
	}
	return nil
}
