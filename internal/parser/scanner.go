package parser

import "fmt"

type ScanError struct {
	Pos      int
	Expected string
	Ch       byte
}

func newScanError(pos int, expected string, ch byte) *ScanError {
	return &ScanError{pos, expected, ch}
}

func (e ScanError) Error() string {
	return fmt.Sprintf(
		"Expected %s, got '%c' at position %d",
		e.Expected,
		e.Ch,
		e.Pos,
	)
}

type scanner struct {
	input   string
	ch      byte
	pos     int
	readPos int
}

func newScanner(input string) scanner {
	s := scanner{input: input}
	s.read()
	return s
}

func (s *scanner) NextToken() (*Token, error) {
	s.eatSpaces()

	switch s.ch {
	case '-':
		pos := s.pos
		s.read()
		if err := s.assertAlphaNumeric(); err != nil {
			return nil, err
		}
		chars := s.readCharacters()
		return newToken(FLAG, fmt.Sprintf("-%s", chars), pos), nil

	case '<':
		pos := s.pos
		s.read()
		if err := s.assertAlphabet(); err != nil {
			return nil, err
		}
		pname := s.readParamName()
		if s.ch != '>' {
			return nil, newScanError(s.pos, "'>'", s.ch)
		}
		s.read()
		if s.ch == '.' {
			s.read()
			if err := s.assertDot(); err != nil {
				return nil, err
			}
			s.read()
			if err := s.assertDot(); err != nil {
				return nil, err
			}
			s.read()
			return newToken(POSITIONAL_PARAM_LIST, fmt.Sprintf("<%s>...", pname), pos), nil
		}
		return newToken(POSITIONAL_PARAM, fmt.Sprintf("<%s>", pname), pos), nil

	case '[':
		pos := s.pos
		s.read()
		s.eatSpaces()
		if s.ch == '-' {
			s.read()
			err := s.assertAlphaNumeric()
			if err != nil {
				return nil, err
			}
			flag := s.readCharacters()
			s.eatSpaces()
			if isAlphabet(s.ch) {
				pname := s.readParamName()
				s.eatSpaces()
				err = s.assertRightSquareBracket()
				if err != nil {
					return nil, err
				}
				s.read()
				return newToken(FLAG_PARAM, fmt.Sprintf("[-%s %s]", flag, pname), pos), nil
			} else {
				err = s.assertRightSquareBracket()
				if err != nil {
					return nil, err
				}
				s.read()
				return newToken(FLAG_PARAM, fmt.Sprintf("[-%s]", flag), pos), nil
			}
		}
		err := s.assertAlphabet()
		if err != nil {
			return nil, err
		}
		pname := s.readParamName()
		s.eatSpaces()
		err = s.assertRightSquareBracket()
		if err != nil {
			return nil, err
		}
		s.read()
		return newToken(OPTIONAL_PARAM, fmt.Sprintf("[%s]", pname), pos), nil

	case 0:
		return newToken(EOF, "", s.pos), nil

	default:
		return nil, newScanError(s.pos, "'-', '[', or '<'", s.ch)
	}
}

func (s *scanner) read() {
	if s.readPos >= len(s.input) {
		s.ch = 0
	} else {
		s.ch = s.input[s.readPos]
	}
	s.pos = s.readPos
	s.readPos++
}

func (s *scanner) readCharacters() string {
	pos := s.pos
	for isAlphanNumeric(s.ch) {
		s.read()
	}
	return s.input[pos:s.pos]
}

func (s *scanner) readParamName() string {
	pos := s.pos
	for isAlphabet(s.ch) {
		s.read()
	}
	return s.input[pos:s.pos]
}

func (s *scanner) eatSpaces() {
	for s.ch == ' ' {
		s.read()
	}
}

func (s *scanner) assertAlphaNumeric() *ScanError {
	if !isAlphanNumeric(s.ch) {
		return newScanError(s.pos, "[0-9a-zA-Z]", s.ch)
	}
	return nil
}

func (s *scanner) assertAlphabet() *ScanError {
	if !isAlphabet(s.ch) {
		return newScanError(s.pos, "[a-zA-Z]", s.ch)
	}
	return nil
}

func (s *scanner) assertRightSquareBracket() *ScanError {
	if s.ch != ']' {
		return newScanError(s.pos, "']'", s.ch)
	}
	return nil
}

func (s *scanner) assertDot() *ScanError {
	if s.ch != '.' {
		return newScanError(s.pos, "'.'", s.ch)
	}
	return nil
}

func isAlphanNumeric(ch byte) bool {
	return isDigit(ch) || isAlphabet(ch)
}

func isDigit(ch byte) bool {
	return (ch >= '0' && ch <= '9')
}

func isAlphabet(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
