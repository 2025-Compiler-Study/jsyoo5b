package parser_test

import (
	"embed"
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/require"
	"mini-c/internal/parser"
	"mini-c/internal/testHelper"
	"testing"
)

var tokenTypes []string

func init() {
	lexer := parser.NewMiniCLexer(nil)
	tokenTypes = lexer.GetSymbolicNames()
}

//go:embed testdata/lexer/success/*.mc
var testLexerSuccessFs embed.FS

func TestSuccessLexing(t *testing.T) {
	testCases, err := testHelper.LoadTestFile(testLexerSuccessFs, ".mc")
	require.NoError(t, err)

	for name, source := range testCases {
		t.Run(name, func(t *testing.T) {
			tokens, lexErrors := lexSource(source)
			require.Empty(t, lexErrors)

			for _, tok := range tokens {
				var tokenName string
				switch tok.GetTokenType() {
				case parser.MiniCLexerIDENT:
					tokenName = tok.GetText()
				case parser.MiniCLexerDECIMAL_LIT,
					parser.MiniCLexerOCTAL_LIT,
					parser.MiniCLexerHEX_LIT:
					tokenName = fmt.Sprintf("(%s)", tok.GetText())
				default:
					tokenName = fmt.Sprintf("<%s>", tokenTypes[tok.GetTokenType()])
				}
				fmt.Printf("%3d:%d\t%s\n", tok.GetLine(), tok.GetColumn(), tokenName)
			}
		})
	}
}

//go:embed testdata/lexer/fail/*.mc
var testLexerFailFs embed.FS

func TestFailLexing(t *testing.T) {
	testCases, err := testHelper.LoadTestFile(testLexerFailFs, ".mc")
	require.NoError(t, err)

	for name, source := range testCases {
		t.Run(name, func(t *testing.T) {
			_, lexErrors := lexSource(source)
			require.NotEmpty(t, lexErrors)

			for _, e := range lexErrors {
				fmt.Printf("%3d:%d\t%s\n", e.Line, e.Column, e.error)
			}
		})
	}
}

func lexSource(input string) ([]antlr.Token, []lexError) {
	is := antlr.NewInputStream(input)
	lexer := parser.NewMiniCLexer(is)
	lexer.RemoveErrorListeners()
	errorListener := &lexErrorListener{}
	lexer.AddErrorListener(errorListener)

	tokens := lexer.GetAllTokens()
	return tokens, errorListener.Errors
}

type lexError struct {
	Line   int
	Column int
	error
}

type lexErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []lexError
}

func (l *lexErrorListener) SyntaxError(
	_ antlr.Recognizer,
	_ any,
	line, column int,
	msg string,
	_ antlr.RecognitionException,
) {
	l.Errors = append(l.Errors, lexError{
		Line:   line,
		Column: column,
		error:  errors.New(msg),
	})
}
