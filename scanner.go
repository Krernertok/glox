package glox

import "fmt"

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
