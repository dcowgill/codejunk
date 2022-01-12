package d18

import (
	"adventofcode2021/lib"
	"fmt"
	"strconv"
	"unicode"
)

func Run() {
	lib.Run(18, part1, part2)
}

func part1() int64 {
	return int64(magnitude(lib.Foldl(add, lib.Map(parseNum, realInput))))
}

func part2() int64 {
	addends := lib.Map(parseNum, realInput)
	max := 0
	for i := range addends {
		for j := range addends {
			if i != j {
				v := magnitude(add(addends[i].clone(), addends[j].clone()))
				max = lib.Max(max, v)
			}
		}
	}
	return int64(max)
}

func magnitude(n *Node) int {
	switch {
	case n == nil:
		return 0
	case n.isNumber():
		return n.Value
	}
	return 3*magnitude(n.L) + 2*magnitude(n.R)
}

func add(lhs, rhs *Node) *Node {
	return reduce(&Node{L: lhs, R: rhs})
}

//====================
// Reducing
//====================

func reduce(n *Node) *Node {
	for explode(n) || split(n) {
		// continue
	}
	return n
}

func explode(root *Node) bool {
	_, _, ok := explodeRec(root, 0)
	return ok
}

func explodeRec(n *Node, depth int) (int, int, bool) {
	if n == nil {
		return 0, 0, false
	} else if n.isPair() && depth >= 4 {
		lv, rv := n.L.Value, n.R.Value
		n.L, n.R, n.Value = nil, nil, 0
		return lv, rv, true
	} else if lv, rv, ok := explodeRec(n.L, depth+1); ok {
		return lv, addLeft(n.R, rv), true
	} else if lv, rv, ok := explodeRec(n.R, depth+1); ok {
		return addRight(n.L, lv), rv, true
	}
	return 0, 0, false
}

func addLeft(n *Node, v int) int {
	switch {
	case n == nil || v < 0:
		return v
	case n.isNumber():
		n.Value += v
		return -1
	case addLeft(n.L, v) < 0:
		return -1
	case addLeft(n.R, v) < 0:
		return -1
	}
	return v
}

func addRight(n *Node, v int) int {
	switch {
	case n == nil || v < 0:
		return v
	case n.isNumber():
		n.Value += v
		return -1
	case addRight(n.R, v) < 0:
		return -1
	case addRight(n.L, v) < 0:
		return -1
	}
	return v
}

func split(n *Node) bool {
	switch {
	case n == nil:
		return false
	case n.isNumber() && n.Value > 9:
		n.L = &Node{Value: n.Value / 2}
		n.R = &Node{Value: n.Value/2 + (n.Value % 2)}
		n.Value = 0
		return true
	case split(n.L):
		return true
	case split(n.R):
		return true
	}
	return false
}

//====================
// Parsing
//====================

type Node struct {
	L, R  *Node
	Value int
}

func (n *Node) isNumber() bool { return n != nil && n.L == nil && n.R == nil }
func (n *Node) isPair() bool   { return n != nil && n.L.isNumber() && n.R.isNumber() }

func (n *Node) String() string {
	if n.isNumber() {
		return fmt.Sprintf("%d", n.Value)
	}
	return fmt.Sprintf("[%s,%s]", n.L, n.R)
}

func (n *Node) clone() *Node {
	if n == nil {
		return nil
	}
	return &Node{
		L:     n.L.clone(),
		R:     n.R.clone(),
		Value: n.Value,
	}
}

func parseNum(s string) *Node {
	return parse(tokenize(s))
}

func parse(ls *Lexemes) *Node {
	var node Node
	switch tok := ls.peek(); tok {
	case BracketOpen:
		ls.consume(BracketOpen)
		node.L = parse(ls)
		ls.consume(Comma)
		node.R = parse(ls)
		ls.consume(BracketClose)
	case Number:
		node.Value = ls.consume(Number).Value
	default:
		panic(fmt.Sprintf("unexpected token %v", tok))
	}
	return &node
}

//====================
// Lexing
//====================

type Token int

const (
	Number       Token = 1
	BracketOpen  Token = 2
	BracketClose Token = 3
	Comma        Token = 4
)

type Lexeme struct {
	Token Token
	Value int
}

type Lexemes struct {
	lexemes []Lexeme
	pos     int
}

func (ls *Lexemes) consume(tok Token) Lexeme {
	l := ls.lexemes[ls.pos]
	if l.Token != tok {
		panic(fmt.Sprintf("consume(%v): next lexeme is %+v", tok, l))
	}
	ls.pos++
	return l
}

func (ls Lexemes) peek() Token {
	return ls.lexemes[ls.pos].Token
}

func tokenize(s string) *Lexemes {
	var lexemes []Lexeme
	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch {
		case ch == '[':
			lexemes = append(lexemes, Lexeme{Token: BracketOpen})
		case ch == ']':
			lexemes = append(lexemes, Lexeme{Token: BracketClose})
		case ch == ',':
			lexemes = append(lexemes, Lexeme{Token: Comma})
		case '0' <= ch && ch <= '9':
			n, len := tokenizeNumber(s[i:])
			lexemes = append(lexemes, Lexeme{Number, n})
			i += len - 1
		case unicode.IsSpace(rune(ch)):
			continue
		default:
			panic(fmt.Sprintf("unexpected character '%c' at offset %d in %q", ch, i, s))
		}
	}
	return &Lexemes{lexemes: lexemes}
}

func tokenizeNumber(s string) (val, len int) {
	i := 0
	for i < len(s) && '0' <= s[i] && s[i] <= '9' {
		i++
	}
	n, err := strconv.ParseInt(s[:i], 10, 64)
	if err != nil {
		panic(err)
	}
	return int(n), i
}
