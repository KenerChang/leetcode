package basiccalculator

import (
	"strconv"
)

const (
	LOWEST int = iota
	SUMMIN     // sum and minus
	PREFIX     // -x
	PAREN      // paren
)

var (
	digits map[byte]bool = map[byte]bool{
		'0': true,
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
	}
)

func calculate(s string) int {
	// Use Pratt's parsing algorithm

	// tokenize
	tokens := tokenize(s)

	// parse
	parser := &PrattParser{
		tokens: tokens,
	}

	return parser.Parse()
}

func tokenize(s string) []string {
	results := []string{}
	i := 0
	for i < len(s) {
		switch s[i] {
		case ' ':
			// skip
			i++
		case '(', ')', '+', '-':
			results = append(results, string(s[i]))
			i++
		default:
			// parse number
			num := []byte{s[i]}
			i++
			for i < len(s) && digits[s[i]] {
				num = append(num, s[i])
				i++
			}

			results = append(results, string(num))
		}
	}
	return results
}

// Simplified Pratt Parser
type PrattParser struct {
	tokens []string
	cur    int
}

func (p *PrattParser) Parse() int {
	return p.parseExpression(LOWEST)
}

func (p *PrattParser) parseExpression(precedence int) int {
	// parse prefix
	var left int
	switch p.curToken() {
	case "(":
		p.cur++
		left = p.parseExpression(LOWEST)
		p.cur++
	case "-":
		p.cur++
		left = -p.parseExpression(PREFIX)
	case "+":
		p.cur++
		left += p.parseExpression(PREFIX)
	case ")":
		p.cur++
		// return left
	default:
		left, _ = strconv.Atoi(p.curToken())
		p.cur++
	}

	// fmt.Printf("cur: %s, precedence:%d, curPrecedence: %d\n", p.curToken(), precedence, p.curPrecedence())

	for p.cur < len(p.tokens) && precedence < p.curPrecedence() {
		// parse infi
		// fmt.Printf("left: %d, p.curToken: %s\n", left, p.curToken())
		// fmt.Printf("cur: %s, precedence:%d, curPrecedence: %d\n", p.curToken(), precedence, p.curPrecedence())
		switch p.curToken() {
		case "+":
			p.cur++
			right := p.parseExpression(SUMMIN)
			// fmt.Printf("left %d + right %d = %d \n", left, right, left+right)
			left += right
			// left += p.parseExpression(SUMMIN)
		case "-":
			p.cur++
			right := p.parseExpression(SUMMIN)
			// fmt.Printf("left %d - right %d = %d \n", left, right, left-right)
			// fmt.Printf("cur: %s, precedence:%d, curPrecedence: %d\n", p.curToken(), precedence, p.curPrecedence())

			left -= right
			// left -= p.parseExpression(SUMMIN)
		case ")":
			return left
		}
	}

	return left
}

func (p *PrattParser) curToken() string {
	return p.tokens[p.cur]
}

func (p *PrattParser) curPrecedence() int {
	if p.cur >= len(p.tokens) {
		return LOWEST
	}

	switch p.tokens[p.cur] {
	case "+", "-":
		return SUMMIN
	case "(", ")":
		return PAREN
	default:
		return LOWEST
	}
}
