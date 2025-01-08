package parser

const (
	EOF                   = "EOF"
	FLAG                  = "FLAG"
	POSITIONAL_PARAM      = "POSITIONAL_PARAM"
	POSITIONAL_PARAM_LIST = "POSITIONAL_PARAM_LIST"
	OPTIONAL_PARAM        = "OPTIONAL_PARAM"
	FLAG_PARAM            = "FLAG_PARAM"
)

type Token struct {
	Type    string
	Literal string
	Pos     int
}

func newToken(tokenType, literal string, pos int) *Token {
	return &Token{
		Type:    tokenType,
		Literal: literal,
		Pos:     pos,
	}
}
