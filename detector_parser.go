package strings2

import (
	"fmt"
	"strings"
	"unicode"
)

type tokenType int

const (
	tokenEOF tokenType = iota
	tokenIdent
	tokenOpenParen
	tokenCloseParen
	tokenComma
)

type token struct {
	typ tokenType
	val string
}

func tokenizeDetectorExpr(s string) []token {
	var tokens []token
	var curr strings.Builder
	for _, r := range s {
		switch r {
		case '(', ')', ',':
			if curr.Len() > 0 {
				tokens = append(tokens, token{tokenIdent, strings.TrimSpace(curr.String())})
				curr.Reset()
			}
			var typ tokenType
			if r == '(' {
				typ = tokenOpenParen
			} else if r == ')' {
				typ = tokenCloseParen
			} else {
				typ = tokenComma
			}
			tokens = append(tokens, token{typ, string(r)})
		case ' ', '\t', '\n', '\r':
			if curr.Len() > 0 {
				tokens = append(tokens, token{tokenIdent, strings.TrimSpace(curr.String())})
				curr.Reset()
			}
		default:
			curr.WriteRune(r)
		}
	}
	if curr.Len() > 0 {
		tokens = append(tokens, token{tokenIdent, strings.TrimSpace(curr.String())})
	}
	return tokens
}

func parseDetectorExpr(tokens []token, pos *int) (DelimiterDetector, error) {
	if *pos >= len(tokens) {
		return nil, fmt.Errorf("unexpected EOF")
	}
	t := tokens[*pos]
	*pos++

	if t.typ != tokenIdent {
		return nil, fmt.Errorf("expected identifier, got %q", t.val)
	}

	ident := strings.TrimSpace(strings.ToLower(t.val))

	if *pos < len(tokens) && tokens[*pos].typ == tokenOpenParen {
		*pos++
		var args []DelimiterDetector
		for *pos < len(tokens) && tokens[*pos].typ != tokenCloseParen {
			arg, err := parseDetectorExpr(tokens, pos)
			if err != nil {
				return nil, err
			}
			args = append(args, arg)
			if *pos < len(tokens) && tokens[*pos].typ == tokenComma {
				*pos++
			} else {
				break
			}
		}
		if *pos >= len(tokens) || tokens[*pos].typ != tokenCloseParen {
			return nil, fmt.Errorf("expected ')', got EOF/other")
		}
		*pos++

		switch ident {
		case "any", "union":
			return func(subs []SubPart, index int) int {
				for _, arg := range args {
					if l := arg(subs, index); l > 0 {
						return l
					}
				}
				return 0
			}, nil
		case "not":
			if len(args) != 1 {
				return nil, fmt.Errorf("not() expects exactly one argument")
			}
			return func(subs []SubPart, index int) int {
				if args[0](subs, index) == 0 {
					return 1
				}
				return 0
			}, nil
		default:
			return nil, fmt.Errorf("unknown function: %s", ident)
		}
	} else {
		switch ident {
		case "numeric":
			return func(subs []SubPart, index int) int {
				if unicode.IsDigit(subs[index].Rune()) {
					return 1
				}
				return 0
			}, nil
		case "whitespace":
			return func(subs []SubPart, index int) int {
				if unicode.IsSpace(subs[index].Rune()) {
					return 1
				}
				return 0
			}, nil
		case "nonalphanumeric":
			return func(subs []SubPart, index int) int {
				r := subs[index].Rune()
				if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
					return 1
				}
				return 0
			}, nil
		case "tab":
			return func(subs []SubPart, index int) int {
				if subs[index].Rune() == '\t' {
					return 1
				}
				return 0
			}, nil
		default:
			return nil, fmt.Errorf("unknown identifier: %s", ident)
		}
	}
}

// ParseDelimiterDetector parses a string expression into a DelimiterDetector function.
func ParseDelimiterDetector(expr string) (DelimiterDetector, error) {
	tokens := tokenizeDetectorExpr(expr)
	if len(tokens) == 0 {
		return nil, nil
	}
	pos := 0
	det, err := parseDetectorExpr(tokens, &pos)
	if err != nil {
		return nil, err
	}
	if pos < len(tokens) {
		return nil, fmt.Errorf("unexpected token at end: %q", tokens[pos].val)
	}
	return det, nil
}

// ParserDelimiterDetector is a typed option for DelimiterDetector configuration.
type ParserDelimiterDetector struct {
    Detector DelimiterDetector
}

func (b ParserDelimiterDetector) Apply(p *ParserConfig) {
	p.DelimiterDetector = b.Detector
}

// WithDelimiterDetector enables or disables treating non-alphanumeric characters as delimiters.
func WithDelimiterDetector(detector DelimiterDetector) ParserOption {
	return ParserDelimiterDetector{Detector: detector}
}
