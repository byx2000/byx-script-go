package parser

import (
	. "byx-script-go/src/common"
	"errors"
	"fmt"
)

type ParseResult struct {
	Result any
	Remain Input
}

type ParseFunc []func(Input) (ParseResult, error)

type Parser struct {
	parse ParseFunc
}

func Fail() Parser {
	return Parser{ParseFunc{func(input Input) (ParseResult, error) {
		return ParseResult{}, errors.New("no error message")
	}}}
}

func Any() Parser {
	return Parser{ParseFunc{func(input Input) (ParseResult, error) {
		if input.End() {
			return ParseResult{}, errors.New("unexpected end of file")
		}
		c := input.Current()
		return ParseResult{c, input.Next()}, nil
	}}}
}

func Chs(chs ...rune) Parser {
	set := make(map[rune]bool)
	for _, c := range chs {
		set[c] = true
	}
	return Parser{ParseFunc{func(input Input) (ParseResult, error) {
		if input.End() {
			return ParseResult{}, errors.New("unexpected end of file")
		}
		c := input.Current()
		_, exist := set[c]
		if !exist {
			msg := fmt.Sprintf("at row %d, col %d: expected character in %v but met %c", input.Row(), input.Col(), chs, c)
			return ParseResult{}, errors.New(msg)
		}
		return ParseResult{c, input.Next()}, nil
	}}}
}

func Ch(c rune) Parser {
	return Parser{ParseFunc{func(input Input) (ParseResult, error) {
		if input.End() {
			return ParseResult{}, errors.New("unexpected end of file")
		}
		ch := input.Current()
		if c != ch {
			msg := fmt.Sprintf("at row %d, col %d: expected %c", input.Row(), input.Col(), c)
			return ParseResult{}, errors.New(msg)
		}
		return ParseResult{c, input.Next()}, nil
	}}}
}

func Not(c rune) Parser {
	return Parser{ParseFunc{func(input Input) (ParseResult, error) {
		if input.End() {
			return ParseResult{}, errors.New("unexpected end of file")
		}
		ch := input.Current()
		if c == ch {
			msg := fmt.Sprintf("at row %d, col %d: unexpected character %c", input.Row(), input.Col(), ch)
			return ParseResult{}, errors.New(msg)
		}
		return ParseResult{ch, input.Next()}, nil
	}}}
}

func Range(c1 rune, c2 rune) Parser {
	return Parser{ParseFunc{func(input Input) (ParseResult, error) {
		if input.End() {
			return ParseResult{}, errors.New("unexpected end of file")
		}
		c := input.Current()
		if (c-c1)*(c-c2) > 0 {
			msg := fmt.Sprintf("index at %d: expected character in range (%c, %c) but met %c", input.index, c1, c2, c)
			return ParseResult{}, errors.New(msg)
		}
		return ParseResult{c, input.Next()}, nil
	}}}
}

func String(s string) Parser {
	return Parser{ParseFunc{func(input Input) (ParseResult, error) {
		i := input
		for _, c := range s {
			if i.End() {
				return ParseResult{}, errors.New("unexpected end of file")
			}
			if i.Current() != c {
				msg := fmt.Sprintf("at row %d, col %d: expected %s", input.row, input.col, s)
				return ParseResult{}, errors.New(msg)
			}
			i = i.Next()
		}
		return ParseResult{s, i}, nil
	}}}
}

func Map(p Parser, mapper func(any) any) Parser {
	return Parser{ParseFunc{func(input Input) (ParseResult, error) {
		r, err := p.parse[0](input)
		if err != nil {
			return ParseResult{}, err
		}
		return ParseResult{mapper(r.Result), r.Remain}, nil
	}}}
}

func And(lhs Parser, rhs Parser) Parser {
	return Parser{ParseFunc{func(input Input) (ParseResult, error) {
		r1, err := lhs.parse[0](input)
		if err != nil {
			return ParseResult{}, err
		}
		r2, err := rhs.parse[0](r1.Remain)
		if err != nil {
			return ParseResult{}, err
		}
		return ParseResult{Pair{r1.Result, r2.Result}, r2.Remain}, nil
	}}}
}

func Seq(parsers ...Parser) Parser {
	return Parser{ParseFunc{
		func(input Input) (ParseResult, error) {
			rs := make([]any, 0)
			for _, p := range parsers {
				r, err := p.parse[0](input)
				if err != nil {
					return ParseResult{}, err
				}
				rs = append(rs, r.Result)
				input = r.Remain
			}
			return ParseResult{rs, input}, nil
		},
	}}
}

func Or(lhs Parser, rhs Parser) Parser {
	return Parser{ParseFunc{func(input Input) (ParseResult, error) {
		r, err := lhs.parse[0](input)
		if err == nil {
			return r, nil
		}
		r, err = rhs.parse[0](input)
		if err != nil {
			return ParseResult{}, err
		}
		return r, nil
	}}}
}

func OneOf(p1 Parser, p2 Parser, parsers ...Parser) Parser {
	p := Or(p1, p2)
	for _, pp := range parsers {
		p = Or(p, pp)
	}
	return p
}

func SkipFirst(p1 Parser, p2 Parser) Parser {
	return p1.And(p2).Map(func(p any) any {
		return p.(Pair).Second
	})
}

func SkipSecond(p1 Parser, p2 Parser) Parser {
	return p1.And(p2).Map(func(p any) any {
		return p.(Pair).First
	})
}

type SkipWrapper struct {
	lhs Parser
	And func(Parser) Parser
}

func Skip(lhs Parser) SkipWrapper {
	return SkipWrapper{lhs, func(rhs Parser) Parser {
		return SkipFirst(lhs, rhs)
	}}
}

func Many(p Parser) Parser {
	return Parser{ParseFunc{func(input Input) (ParseResult, error) {
		rs := make([]any, 0)
		for {
			r, err := p.parse[0](input)
			if err != nil {
				break
			}
			rs = append(rs, r.Result)
			input = r.Remain
		}
		return ParseResult{rs, input}, nil
	}}}
}

func Many1(p Parser) Parser {
	return p.And(p.Many()).Map(func(p any) any {
		pair := p.(Pair)
		rs := make([]any, 0)
		rs = append(rs, pair.First)
		rs = append(rs, pair.Second.([]any)...)
		return rs
	})
}

func Optional(p Parser, defaultValue any) Parser {
	return Parser{ParseFunc{func(input Input) (ParseResult, error) {
		r, err := p.parse[0](input)
		if err != nil {
			return ParseResult{defaultValue, input}, nil
		}
		return r, nil
	}}}
}

func Lazy(factory func() Parser) Parser {
	return Parser{ParseFunc{func(input Input) (ParseResult, error) {
		return factory().parse[0](input)
	}}}
}

func Peek(probe Parser, success Parser, failed Parser) Parser {
	return Parser{ParseFunc{func(input Input) (ParseResult, error) {
		_, err := probe.parse[0](input)
		if err != nil {
			return failed.parse[0](input)
		}
		return success.parse[0](input)
	}}}
}

func SeparateBy(delimiter Parser, p Parser) Parser {
	return p.And(Skip(delimiter).And(p).Many()).Map(func(p any) any {
		var result []any
		result = append(result, p.(Pair).First)
		for _, e := range p.(Pair).Second.([]any) {
			result = append(result, e)
		}
		return result
	})
}

func Fatal(p Parser) Parser {
	return Parser{ParseFunc{func(input Input) (ParseResult, error) {
		r, e := p.parse[0](input)
		if e != nil {
			panic(e)
		}
		return r, nil
	}}}
}

func NewParser() Parser {
	return Parser{ParseFunc{nil}}
}

func (p Parser) ParseToEnd(s string) (any, error) {
	r, err := p.parse[0](CreateInput(s))
	if err != nil {
		return nil, err
	}
	remain := r.Remain
	if !remain.End() {
		return nil, errors.New(fmt.Sprintf("at row %d, col %d: end of file not reached", remain.Row(), remain.Col()))
	}
	return r.Result, nil
}

func (p Parser) Set(parser Parser) {
	p.parse[0] = parser.parse[0]
}

func (p Parser) And(rhs Parser) Parser {
	return And(p, rhs)
}

func (p Parser) Or(rhs Parser) Parser {
	return Or(p, rhs)
}

func (p Parser) Many() Parser {
	return Many(p)
}

func (p Parser) Many1() Parser {
	return Many1(p)
}

func (p Parser) Map(mapper func(any) any) Parser {
	return Map(p, mapper)
}

func (p Parser) Skip(rhs Parser) Parser {
	return SkipSecond(p, rhs)
}

func (p Parser) SurroundBy(parser Parser) Parser {
	return Seq(parser, p, parser).Map(func(rs any) any {
		return rs.([]any)[1]
	})
}

func (p Parser) ManyUntil(until Parser) Parser {
	return Peek(until, Fail(), p).Many()
}

func (p Parser) Optional(defaultValue any) Parser {
	return Optional(p, defaultValue)
}

func (p Parser) Fatal() Parser {
	return Fatal(p)
}
