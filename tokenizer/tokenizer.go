package tokenizer

import (
	"fmt"
	"strings"
)

// A Token contains the type and the value of a part of code
type Token struct {
	Type  string
	Value string
}

// IF ( condition ) { ... } ENDIF=> condition false on saute après le endif condition true on change rien

const (
	LeftParentheses  = "CHOUQUETTE"
	RightParentheses = "CLAFOUTIS"
	Quotes           = "PARISBREST"
	SemiColon        = "BAGUETTE"
	If			 	 = "IF"
	EndIf			 = "ENDFOR"

	ParenthesesTag = "PARENTHESES"
	QuotesTag      = "QUOTES"
	SemiColonTag   = "SEMICOLON"
	NameTag        = "NAME"
	Number         = "NUMBER"
	IfTag		   = "IF"
	EndIfTag	   = "ENDIF"
	OperatorTag	   = "OPERATOR"
)

var (
	i          int
	character  string
	characters []string
	tokens     []Token
)

// Tokenize takes the given code, tokenize it and returns an array of Tokens
func Tokenize(code string) []Token {
	code += " "
	characters = strings.Split(code, "")

	for i = 0; i < len(characters); i++ {
		character = characters[i]

		// If the character is a letter, get all the next letters and create a token
		if IsLetter(character) {
			tokens = append(tokens, AppendStringValue())
		}

		// If the character is a numbers, get all the next numbers and create a token
		if IsNumber(character) {
			tokens = append(tokens, AppendNumberValue())
		}

		if IsSpecialChar(character){
			tokens = append(tokens, AppendSpecialCharValue())
		}
	}
	fmt.Println(tokens)
	return tokens
}

// AppendStringValue browses the next letters to add a token
func AppendStringValue() Token {
	var value string

	// Iterate through all the next letters
	for IsLetter(character) {
		value += character
		i++
		character = characters[i]
	}

	// If the value is a parentheses append the parentheses token
	if IsParentheses(value) {
		return Token{
			Type:  ParenthesesTag,
			Value: value,
		}
	}

	// If the value is a quote append the quote token
	if value == Quotes {
		return Token{
			Type:  QuotesTag,
			Value: value,
		}
	}

	// If the value is a semicolon append the semicolon token
	if value == SemiColon {
		return Token{
			Type:  SemiColonTag,
			Value: value,
		}
	}

	if value == If {
		return Token{
			Type: IfTag,
			Value: value,
		}
	}

	if value == EndIf {
		return Token{
			Type: EndIfTag,
			Value: value,
		}
	}

	return Token{
		Type:  NameTag,
		Value: value,
	}
}

// AppendNumberValue browses the next numbers to add a token
func AppendNumberValue() Token {
	var value string

	for IsNumber(character) {
		value += character
		i++
		character = characters[i]
	}

	return Token{
		Type:  Number,
		Value: value,
	}
}

func AppendSpecialCharValue() Token {
	var value string

	for IsSpecialChar(character) {
		value += character
		i++
		character = characters[i]
	}
	return Token{
		Type: OperatorTag,
		Value: value,
	}
}
