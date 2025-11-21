package parser

import (
	"embed"
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

type lexError struct {
	Line, Column int
	error
}

type lexErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []lexError
}

func (l *lexErrorListener) SyntaxError(_ antlr.Recognizer, _ any, line, column int, msg string, _ antlr.RecognitionException) {
	l.Errors = append(l.Errors, lexError{
		line,
		column,
		errors.New(msg),
	})
}

func lexSource(input string) ([]antlr.Token, []lexError) {
	is := antlr.NewInputStream(input)
	lexer := NewMiniCLexer(is)
	lexer.RemoveErrorListeners()
	errorListener := &lexErrorListener{}
	lexer.AddErrorListener(errorListener)

	tokens := lexer.GetAllTokens()
	return tokens, errorListener.Errors
}

//go:embed testdata/lexer/success/*.mc
var testLexerSuccessFs embed.FS

func TestSuccessLexing(t *testing.T) {
	testCases, err := loadTestFiles(testLexerSuccessFs, ".mc")
	require.NoError(t, err)

	lexer := NewMiniCLexer(nil)
	tokenTypes := lexer.GetSymbolicNames()

	for name, source := range testCases {
		t.Run(name, func(t *testing.T) {
			tokens, lexErrors := lexSource(source)
			require.Empty(t, lexErrors)

			for _, token := range tokens {
				var tokenName string
				switch token.GetTokenType() {
				case MiniCLexerIDENT:
					tokenName = token.GetText()
				case MiniCLexerDECIMAL_LIT, MiniCLexerOCTAL_LIT, MiniCLexerHEX_LIT:
					tokenName = fmt.Sprintf("(%s)", token.GetText())
				default:
					tokenName = fmt.Sprintf("[%s]", tokenTypes[token.GetTokenType()])
				}
				fmt.Printf("%3d:%d\t%s\n", token.GetLine(), token.GetColumn(), tokenName)
			}
		})
	}
}
