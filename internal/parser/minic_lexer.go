// Code generated from MiniCLexer.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MiniCLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var MiniCLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func miniclexerLexerInit() {
	staticData := &MiniCLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'const'", "'int'", "'void'", "'if'", "'else'", "'while'", "'return'",
		"", "'+'", "'-'", "'*'", "'/'", "'%'", "'++'", "'--'", "'='", "'+='",
		"'-='", "'*='", "'/='", "'%='", "'!'", "'&&'", "'||'", "'=='", "'!='",
		"'<'", "'<='", "'>'", "'>='", "'('", "')'", "'{'", "'}'", "'['", "']'",
		"','", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "CONST", "INT", "VOID", "IF", "ELSE", "WHILE", "RETURN", "IDENT",
		"PLUS", "MINUS", "MUL", "DIV", "MOD", "INC", "DEC", "ASSIGN", "PLUS_ASSIGN",
		"MINUS_ASSIGN", "MUL_ASSIGN", "DIV_ASSIGN", "MOD_ASSIGN", "LOGICAL_NOT",
		"LOGICAL_AND", "LOGIC_OR", "EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS",
		"GREATER", "GREATER_OR_EQUALS", "L_PAREN", "R_PAREN", "L_CURLY", "R_CURLY",
		"L_BRACKET", "R_BRACKET", "COMMA", "SEMI", "DECIMAL_LIT", "OCTAL_LIT",
		"HEX_LIT", "WS", "C89_COMMENT", "CPP_COMMENT",
	}
	staticData.RuleNames = []string{
		"CONST", "INT", "VOID", "IF", "ELSE", "WHILE", "RETURN", "IDENT", "PLUS",
		"MINUS", "MUL", "DIV", "MOD", "INC", "DEC", "ASSIGN", "PLUS_ASSIGN",
		"MINUS_ASSIGN", "MUL_ASSIGN", "DIV_ASSIGN", "MOD_ASSIGN", "LOGICAL_NOT",
		"LOGICAL_AND", "LOGIC_OR", "EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS",
		"GREATER", "GREATER_OR_EQUALS", "L_PAREN", "R_PAREN", "L_CURLY", "R_CURLY",
		"L_BRACKET", "R_BRACKET", "COMMA", "SEMI", "DECIMAL_LIT", "OCTAL_LIT",
		"HEX_LIT", "DEC_DIGIT", "OCT_DIGIT", "HEX_DIGIT", "ALPHABET", "WS",
		"C89_COMMENT", "CPP_COMMENT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 44, 284, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 137, 8, 7, 1, 7, 1, 7, 1, 7, 5,
		7, 142, 8, 7, 10, 7, 12, 7, 145, 9, 7, 3, 7, 147, 8, 7, 1, 8, 1, 8, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13,
		1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21,
		1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 29, 1, 30, 1, 30, 1, 31, 1, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1,
		34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38,
		5, 38, 225, 8, 38, 10, 38, 12, 38, 228, 9, 38, 3, 38, 230, 8, 38, 1, 39,
		1, 39, 4, 39, 234, 8, 39, 11, 39, 12, 39, 235, 1, 40, 1, 40, 1, 40, 4,
		40, 241, 8, 40, 11, 40, 12, 40, 242, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43,
		1, 43, 1, 44, 1, 44, 1, 45, 4, 45, 254, 8, 45, 11, 45, 12, 45, 255, 1,
		45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 5, 46, 264, 8, 46, 10, 46, 12, 46,
		267, 9, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 1, 47, 1,
		47, 5, 47, 278, 8, 47, 10, 47, 12, 47, 281, 9, 47, 1, 47, 1, 47, 1, 265,
		0, 48, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28,
		57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37,
		75, 38, 77, 39, 79, 40, 81, 41, 83, 0, 85, 0, 87, 0, 89, 0, 91, 42, 93,
		43, 95, 44, 1, 0, 8, 1, 0, 49, 57, 2, 0, 88, 88, 120, 120, 1, 0, 48, 57,
		1, 0, 48, 55, 3, 0, 48, 57, 65, 70, 97, 102, 2, 0, 65, 90, 97, 122, 3,
		0, 9, 10, 13, 13, 32, 32, 2, 0, 10, 10, 13, 13, 291, 0, 1, 1, 0, 0, 0,
		0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
		0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0,
		0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0,
		0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1,
		0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41,
		1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0,
		49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0,
		0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0,
		0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0,
		0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1,
		0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95,
		1, 0, 0, 0, 1, 97, 1, 0, 0, 0, 3, 103, 1, 0, 0, 0, 5, 107, 1, 0, 0, 0,
		7, 112, 1, 0, 0, 0, 9, 115, 1, 0, 0, 0, 11, 120, 1, 0, 0, 0, 13, 126, 1,
		0, 0, 0, 15, 146, 1, 0, 0, 0, 17, 148, 1, 0, 0, 0, 19, 150, 1, 0, 0, 0,
		21, 152, 1, 0, 0, 0, 23, 154, 1, 0, 0, 0, 25, 156, 1, 0, 0, 0, 27, 158,
		1, 0, 0, 0, 29, 161, 1, 0, 0, 0, 31, 164, 1, 0, 0, 0, 33, 166, 1, 0, 0,
		0, 35, 169, 1, 0, 0, 0, 37, 172, 1, 0, 0, 0, 39, 175, 1, 0, 0, 0, 41, 178,
		1, 0, 0, 0, 43, 181, 1, 0, 0, 0, 45, 183, 1, 0, 0, 0, 47, 186, 1, 0, 0,
		0, 49, 189, 1, 0, 0, 0, 51, 192, 1, 0, 0, 0, 53, 195, 1, 0, 0, 0, 55, 197,
		1, 0, 0, 0, 57, 200, 1, 0, 0, 0, 59, 202, 1, 0, 0, 0, 61, 205, 1, 0, 0,
		0, 63, 207, 1, 0, 0, 0, 65, 209, 1, 0, 0, 0, 67, 211, 1, 0, 0, 0, 69, 213,
		1, 0, 0, 0, 71, 215, 1, 0, 0, 0, 73, 217, 1, 0, 0, 0, 75, 219, 1, 0, 0,
		0, 77, 229, 1, 0, 0, 0, 79, 231, 1, 0, 0, 0, 81, 237, 1, 0, 0, 0, 83, 244,
		1, 0, 0, 0, 85, 246, 1, 0, 0, 0, 87, 248, 1, 0, 0, 0, 89, 250, 1, 0, 0,
		0, 91, 253, 1, 0, 0, 0, 93, 259, 1, 0, 0, 0, 95, 273, 1, 0, 0, 0, 97, 98,
		5, 99, 0, 0, 98, 99, 5, 111, 0, 0, 99, 100, 5, 110, 0, 0, 100, 101, 5,
		115, 0, 0, 101, 102, 5, 116, 0, 0, 102, 2, 1, 0, 0, 0, 103, 104, 5, 105,
		0, 0, 104, 105, 5, 110, 0, 0, 105, 106, 5, 116, 0, 0, 106, 4, 1, 0, 0,
		0, 107, 108, 5, 118, 0, 0, 108, 109, 5, 111, 0, 0, 109, 110, 5, 105, 0,
		0, 110, 111, 5, 100, 0, 0, 111, 6, 1, 0, 0, 0, 112, 113, 5, 105, 0, 0,
		113, 114, 5, 102, 0, 0, 114, 8, 1, 0, 0, 0, 115, 116, 5, 101, 0, 0, 116,
		117, 5, 108, 0, 0, 117, 118, 5, 115, 0, 0, 118, 119, 5, 101, 0, 0, 119,
		10, 1, 0, 0, 0, 120, 121, 5, 119, 0, 0, 121, 122, 5, 104, 0, 0, 122, 123,
		5, 105, 0, 0, 123, 124, 5, 108, 0, 0, 124, 125, 5, 101, 0, 0, 125, 12,
		1, 0, 0, 0, 126, 127, 5, 114, 0, 0, 127, 128, 5, 101, 0, 0, 128, 129, 5,
		116, 0, 0, 129, 130, 5, 117, 0, 0, 130, 131, 5, 114, 0, 0, 131, 132, 5,
		110, 0, 0, 132, 14, 1, 0, 0, 0, 133, 147, 3, 89, 44, 0, 134, 137, 5, 95,
		0, 0, 135, 137, 3, 89, 44, 0, 136, 134, 1, 0, 0, 0, 136, 135, 1, 0, 0,
		0, 137, 143, 1, 0, 0, 0, 138, 142, 5, 95, 0, 0, 139, 142, 3, 89, 44, 0,
		140, 142, 3, 83, 41, 0, 141, 138, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0, 141,
		140, 1, 0, 0, 0, 142, 145, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 144,
		1, 0, 0, 0, 144, 147, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 146, 133, 1, 0,
		0, 0, 146, 136, 1, 0, 0, 0, 147, 16, 1, 0, 0, 0, 148, 149, 5, 43, 0, 0,
		149, 18, 1, 0, 0, 0, 150, 151, 5, 45, 0, 0, 151, 20, 1, 0, 0, 0, 152, 153,
		5, 42, 0, 0, 153, 22, 1, 0, 0, 0, 154, 155, 5, 47, 0, 0, 155, 24, 1, 0,
		0, 0, 156, 157, 5, 37, 0, 0, 157, 26, 1, 0, 0, 0, 158, 159, 5, 43, 0, 0,
		159, 160, 5, 43, 0, 0, 160, 28, 1, 0, 0, 0, 161, 162, 5, 45, 0, 0, 162,
		163, 5, 45, 0, 0, 163, 30, 1, 0, 0, 0, 164, 165, 5, 61, 0, 0, 165, 32,
		1, 0, 0, 0, 166, 167, 5, 43, 0, 0, 167, 168, 5, 61, 0, 0, 168, 34, 1, 0,
		0, 0, 169, 170, 5, 45, 0, 0, 170, 171, 5, 61, 0, 0, 171, 36, 1, 0, 0, 0,
		172, 173, 5, 42, 0, 0, 173, 174, 5, 61, 0, 0, 174, 38, 1, 0, 0, 0, 175,
		176, 5, 47, 0, 0, 176, 177, 5, 61, 0, 0, 177, 40, 1, 0, 0, 0, 178, 179,
		5, 37, 0, 0, 179, 180, 5, 61, 0, 0, 180, 42, 1, 0, 0, 0, 181, 182, 5, 33,
		0, 0, 182, 44, 1, 0, 0, 0, 183, 184, 5, 38, 0, 0, 184, 185, 5, 38, 0, 0,
		185, 46, 1, 0, 0, 0, 186, 187, 5, 124, 0, 0, 187, 188, 5, 124, 0, 0, 188,
		48, 1, 0, 0, 0, 189, 190, 5, 61, 0, 0, 190, 191, 5, 61, 0, 0, 191, 50,
		1, 0, 0, 0, 192, 193, 5, 33, 0, 0, 193, 194, 5, 61, 0, 0, 194, 52, 1, 0,
		0, 0, 195, 196, 5, 60, 0, 0, 196, 54, 1, 0, 0, 0, 197, 198, 5, 60, 0, 0,
		198, 199, 5, 61, 0, 0, 199, 56, 1, 0, 0, 0, 200, 201, 5, 62, 0, 0, 201,
		58, 1, 0, 0, 0, 202, 203, 5, 62, 0, 0, 203, 204, 5, 61, 0, 0, 204, 60,
		1, 0, 0, 0, 205, 206, 5, 40, 0, 0, 206, 62, 1, 0, 0, 0, 207, 208, 5, 41,
		0, 0, 208, 64, 1, 0, 0, 0, 209, 210, 5, 123, 0, 0, 210, 66, 1, 0, 0, 0,
		211, 212, 5, 125, 0, 0, 212, 68, 1, 0, 0, 0, 213, 214, 5, 91, 0, 0, 214,
		70, 1, 0, 0, 0, 215, 216, 5, 93, 0, 0, 216, 72, 1, 0, 0, 0, 217, 218, 5,
		44, 0, 0, 218, 74, 1, 0, 0, 0, 219, 220, 5, 59, 0, 0, 220, 76, 1, 0, 0,
		0, 221, 230, 5, 48, 0, 0, 222, 226, 7, 0, 0, 0, 223, 225, 3, 83, 41, 0,
		224, 223, 1, 0, 0, 0, 225, 228, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226,
		227, 1, 0, 0, 0, 227, 230, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 229, 221,
		1, 0, 0, 0, 229, 222, 1, 0, 0, 0, 230, 78, 1, 0, 0, 0, 231, 233, 5, 48,
		0, 0, 232, 234, 3, 85, 42, 0, 233, 232, 1, 0, 0, 0, 234, 235, 1, 0, 0,
		0, 235, 233, 1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 80, 1, 0, 0, 0, 237,
		238, 5, 48, 0, 0, 238, 240, 7, 1, 0, 0, 239, 241, 3, 87, 43, 0, 240, 239,
		1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242, 240, 1, 0, 0, 0, 242, 243, 1, 0,
		0, 0, 243, 82, 1, 0, 0, 0, 244, 245, 7, 2, 0, 0, 245, 84, 1, 0, 0, 0, 246,
		247, 7, 3, 0, 0, 247, 86, 1, 0, 0, 0, 248, 249, 7, 4, 0, 0, 249, 88, 1,
		0, 0, 0, 250, 251, 7, 5, 0, 0, 251, 90, 1, 0, 0, 0, 252, 254, 7, 6, 0,
		0, 253, 252, 1, 0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 253, 1, 0, 0, 0, 255,
		256, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 258, 6, 45, 0, 0, 258, 92,
		1, 0, 0, 0, 259, 260, 5, 47, 0, 0, 260, 261, 5, 42, 0, 0, 261, 265, 1,
		0, 0, 0, 262, 264, 9, 0, 0, 0, 263, 262, 1, 0, 0, 0, 264, 267, 1, 0, 0,
		0, 265, 266, 1, 0, 0, 0, 265, 263, 1, 0, 0, 0, 266, 268, 1, 0, 0, 0, 267,
		265, 1, 0, 0, 0, 268, 269, 5, 42, 0, 0, 269, 270, 5, 47, 0, 0, 270, 271,
		1, 0, 0, 0, 271, 272, 6, 46, 0, 0, 272, 94, 1, 0, 0, 0, 273, 274, 5, 47,
		0, 0, 274, 275, 5, 47, 0, 0, 275, 279, 1, 0, 0, 0, 276, 278, 8, 7, 0, 0,
		277, 276, 1, 0, 0, 0, 278, 281, 1, 0, 0, 0, 279, 277, 1, 0, 0, 0, 279,
		280, 1, 0, 0, 0, 280, 282, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0, 282, 283,
		6, 47, 0, 0, 283, 96, 1, 0, 0, 0, 12, 0, 136, 141, 143, 146, 226, 229,
		235, 242, 255, 265, 279, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MiniCLexerInit initializes any static state used to implement MiniCLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMiniCLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MiniCLexerInit() {
	staticData := &MiniCLexerLexerStaticData
	staticData.once.Do(miniclexerLexerInit)
}

// NewMiniCLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMiniCLexer(input antlr.CharStream) *MiniCLexer {
	MiniCLexerInit()
	l := new(MiniCLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &MiniCLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "MiniCLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MiniCLexer tokens.
const (
	MiniCLexerCONST             = 1
	MiniCLexerINT               = 2
	MiniCLexerVOID              = 3
	MiniCLexerIF                = 4
	MiniCLexerELSE              = 5
	MiniCLexerWHILE             = 6
	MiniCLexerRETURN            = 7
	MiniCLexerIDENT             = 8
	MiniCLexerPLUS              = 9
	MiniCLexerMINUS             = 10
	MiniCLexerMUL               = 11
	MiniCLexerDIV               = 12
	MiniCLexerMOD               = 13
	MiniCLexerINC               = 14
	MiniCLexerDEC               = 15
	MiniCLexerASSIGN            = 16
	MiniCLexerPLUS_ASSIGN       = 17
	MiniCLexerMINUS_ASSIGN      = 18
	MiniCLexerMUL_ASSIGN        = 19
	MiniCLexerDIV_ASSIGN        = 20
	MiniCLexerMOD_ASSIGN        = 21
	MiniCLexerLOGICAL_NOT       = 22
	MiniCLexerLOGICAL_AND       = 23
	MiniCLexerLOGIC_OR          = 24
	MiniCLexerEQUALS            = 25
	MiniCLexerNOT_EQUALS        = 26
	MiniCLexerLESS              = 27
	MiniCLexerLESS_OR_EQUALS    = 28
	MiniCLexerGREATER           = 29
	MiniCLexerGREATER_OR_EQUALS = 30
	MiniCLexerL_PAREN           = 31
	MiniCLexerR_PAREN           = 32
	MiniCLexerL_CURLY           = 33
	MiniCLexerR_CURLY           = 34
	MiniCLexerL_BRACKET         = 35
	MiniCLexerR_BRACKET         = 36
	MiniCLexerCOMMA             = 37
	MiniCLexerSEMI              = 38
	MiniCLexerDECIMAL_LIT       = 39
	MiniCLexerOCTAL_LIT         = 40
	MiniCLexerHEX_LIT           = 41
	MiniCLexerWS                = 42
	MiniCLexerC89_COMMENT       = 43
	MiniCLexerCPP_COMMENT       = 44
)
