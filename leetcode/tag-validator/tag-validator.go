package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

//
// Lexical analyzer
//

type TokenType string

const (
	TokenTypeEOF      TokenType = "EOF"
	TokenTypeTagOpen  TokenType = "TagOpen"
	TokenTypeTagClose TokenType = "TagClose"
	TokenTypeText     TokenType = "Text"
)

type Token struct {
	Type  TokenType
	Value string
	Pos   int
}

type Tokenizer struct {
	input []rune
	pos   int
}

func newTokenizer(input string) *Tokenizer {
	var runes []rune
	for _, r := range input {
		runes = append(runes, r)
	}
	return &Tokenizer{input: runes}
}

func (tok *Tokenizer) next() (Token, error) {
	// Remember the starting position for diagnostic purposes.
	tokenBegin := tok.pos
	makeToken := func(tt TokenType, value string) Token {
		return Token{Type: tt, Value: value, Pos: tokenBegin}
	}
	fail := func(format string, args ...interface{}) error {
		return fmt.Errorf("offset "+strconv.Itoa(tokenBegin)+": "+format, args...)
	}
	// Check for end-of-input.
	if tok.pos >= len(tok.input) {
		return makeToken(TokenTypeEOF, ""), nil
	}
	// Anything that does _not_ begin with '<' is a plain-text block.
	if tok.input[tok.pos] != '<' {
		tok.pos++
		for tok.pos < len(tok.input) && tok.input[tok.pos] != '<' {
			tok.pos++
		}
		text := string(tok.input[tokenBegin:tok.pos])
		return makeToken(TokenTypeText, text), nil
	}
	// Try to match a CDATA block.
	if n := tok.matchHere("<![CDATA["); n > 0 {
		tok.pos += n
		textBegin := tok.pos
		for tok.pos < len(tok.input) {
			if n := tok.matchHere("]]>"); n > 0 {
				text := string(tok.input[textBegin:tok.pos])
				tok.pos += n
				return makeToken(TokenTypeText, text), nil
			}
			tok.pos++
		}
		return Token{}, fail("end-of-input looking for matching ]]>")
	}
	// Has to be an opening tag or a closing tag.
	tokenType := TokenTypeTagOpen
	if tok.matchHere("</") != 0 {
		tokenType = TokenTypeTagClose
		tok.pos += 2
	} else {
		tok.pos += 1
	}
	tagBegin := tok.pos
	for ; tok.pos < len(tok.input); tok.pos++ {
		if tok.input[tok.pos] == '>' {
			tagName := string(tok.input[tagBegin:tok.pos])
			tok.pos++
			if !isValidTagName(tagName) {
				return Token{}, fail("invalid tag name %q", tagName)
			}
			return makeToken(tokenType, tagName), nil
		}
	}
	return Token{}, fail("end-of-input looking for matching '>'")
}

// If s can be found at the current position, returns len(s), else 0.
func (tok *Tokenizer) matchHere(s string) int {
	n := len(s)
	if tok.pos+n > len(tok.input) {
		return 0
	}
	for i := 0; i < n; i++ {
		if tok.input[tok.pos+i] != rune(s[i]) {
			return 0
		}
	}
	return n
}

// Reports whether s satifies the rules for a tag name.
func isValidTagName(s string) bool {
	if len(s) < 1 || len(s) > 9 {
		return false
	}
	for _, r := range s {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}

//
// Recursive descent parser
//

func validate(input string) error {
	// Weird rule: the entire input has to be contained in a tag.
	// To simplify the necessary validation, collect all tokens in a slice.
	var tokens []Token
	pos := 0
	{
		stream := newTokenizer(input)
		for {
			tok, err := stream.next()
			if err != nil {
				return err
			}
			if tok.Type == TokenTypeEOF {
				break
			}
			tokens = append(tokens, tok)
		}
	}

	if tokens[0].Type != TokenTypeTagOpen {
		return errors.New("input must be wrapped in a tag")
	}

	// Create a simple stack of tags.
	var (
		stack []Token
		push  = func(t Token) { stack = append(stack, t) }
		pop   = func() { stack = stack[:len(stack)-1] }
		top   = func() Token {
			if len(stack) == 0 {
				return Token{}
			}
			return stack[len(stack)-1]
		}
	)

	for pos := 0; pos < len(tokens); pos++ {
		// Verify that tags are balanced.
		switch tok.Type {
		case TokenTypeTagOpen:
			push(tok)
		case TokenTypeTagClose:
			cur := top()
			if cur.Type != TokenTypeTagOpen || cur.Value != tok.Value {
				return fmt.Errorf("unbalanced tags: </%s> at offset %d matches <%s> at offset %d",
					tok.Value, tok.Pos, cur.Value, cur.Pos)
			}
			pop()
			if len(stack) == 0 && pos != len(tokens)-1 {
				return errors.New("input must be wrapped in a tag") // weird rule (see above)
			}
		}
	}
	if len(stack) != 0 {
		tok := top()
		return fmt.Errorf("unclosed tag <%s> at offset %d", tok.Value, tok.Pos)
	}
	return nil
}

func main() {
	// const input = "<DIV>Hello, <B>world</B></DIV>! <![CDATA[This is a <div></div>]]> Goodbye, <I>world!</I>"
	const input = "<DIV>This is the first line <![CDATA[<div>]]></DIV>"
	fmt.Println(input)
	tokenizer := newTokenizer(input)
	for {
		tok, err := tokenizer.next()
		if err != nil {
			log.Fatal(err)
		}
		if tok.Type == TokenTypeEOF {
			break
		}
		fmt.Printf("pos=%3d type=%v value=%q\n", tok.Pos, tok.Type, tok.Value)
	}
	if err := validate(input); err != nil {
		log.Fatalf("invalid input: %+v", err)
	}
	fmt.Println("ok -- input is valid")
}
