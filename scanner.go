package glox

import "fmt"

type TokenType int

const (
	// single-character tokens
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// one or two character tokens
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// literals
	IDENTIFIER
	STRING
	NUMBER

	// keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

type Token struct {
	tokenType TokenType
	lexeme    String
	literal   any
	line      int
}

func (t Token) String() string {
	return fmt.Fprint("%s %s %s", t.tokenType, t.lexeme, t.literal)
}

type scanError struct {
	line    int
	where   string
	message string
}

func (se scanError) Error() string {
	return fmt.Sprintf("[line %d] Error %s: %s", se.line, se.where, se.message)
}

func Run(source string) error {
	tokens, err := scanTokens(source)
	if err != nil {
		return err
	}

	for _, token := range tokens {
		fmt.Println(token)
	}
	return nil
}

func scanTokens(source string) ([]string, error) {
	var tokens []string
	return tokens, nil
}
