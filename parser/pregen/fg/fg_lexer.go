// Code generated from parser/FG.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 29, 201,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	25, 5, 25, 154, 10, 25, 3, 26, 3, 26, 3, 27, 3, 27, 5, 27, 160, 10, 27,
	3, 27, 3, 27, 3, 27, 7, 27, 165, 10, 27, 12, 27, 14, 27, 168, 11, 27, 3,
	28, 6, 28, 171, 10, 28, 13, 28, 14, 28, 172, 3, 28, 3, 28, 3, 29, 3, 29,
	3, 29, 3, 29, 7, 29, 181, 10, 29, 12, 29, 14, 29, 184, 11, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 7, 30, 195, 10, 30,
	12, 30, 14, 30, 198, 11, 30, 3, 30, 3, 30, 3, 182, 2, 31, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 49, 2, 51, 2, 53, 26, 55, 27, 57, 28, 59, 29, 3, 2, 5, 4, 2,
	67, 92, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 12, 12, 15, 15, 2,
	205, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25,
	3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2,
	33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2,
	2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2,
	2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2,
	2, 2, 3, 61, 3, 2, 2, 2, 5, 63, 3, 2, 2, 2, 7, 65, 3, 2, 2, 2, 9, 67, 3,
	2, 2, 2, 11, 69, 3, 2, 2, 2, 13, 71, 3, 2, 2, 2, 15, 73, 3, 2, 2, 2, 17,
	75, 3, 2, 2, 2, 19, 77, 3, 2, 2, 2, 21, 79, 3, 2, 2, 2, 23, 81, 3, 2, 2,
	2, 25, 83, 3, 2, 2, 2, 27, 85, 3, 2, 2, 2, 29, 87, 3, 2, 2, 2, 31, 92,
	3, 2, 2, 2, 33, 102, 3, 2, 2, 2, 35, 107, 3, 2, 2, 2, 37, 115, 3, 2, 2,
	2, 39, 122, 3, 2, 2, 2, 41, 129, 3, 2, 2, 2, 43, 134, 3, 2, 2, 2, 45, 141,
	3, 2, 2, 2, 47, 145, 3, 2, 2, 2, 49, 153, 3, 2, 2, 2, 51, 155, 3, 2, 2,
	2, 53, 159, 3, 2, 2, 2, 55, 170, 3, 2, 2, 2, 57, 176, 3, 2, 2, 2, 59, 190,
	3, 2, 2, 2, 61, 62, 7, 61, 2, 2, 62, 4, 3, 2, 2, 2, 63, 64, 7, 36, 2, 2,
	64, 6, 3, 2, 2, 2, 65, 66, 7, 42, 2, 2, 66, 8, 3, 2, 2, 2, 67, 68, 7, 43,
	2, 2, 68, 10, 3, 2, 2, 2, 69, 70, 7, 125, 2, 2, 70, 12, 3, 2, 2, 2, 71,
	72, 7, 97, 2, 2, 72, 14, 3, 2, 2, 2, 73, 74, 7, 63, 2, 2, 74, 16, 3, 2,
	2, 2, 75, 76, 7, 48, 2, 2, 76, 18, 3, 2, 2, 2, 77, 78, 7, 39, 2, 2, 78,
	20, 3, 2, 2, 2, 79, 80, 7, 37, 2, 2, 80, 22, 3, 2, 2, 2, 81, 82, 7, 120,
	2, 2, 82, 24, 3, 2, 2, 2, 83, 84, 7, 46, 2, 2, 84, 26, 3, 2, 2, 2, 85,
	86, 7, 127, 2, 2, 86, 28, 3, 2, 2, 2, 87, 88, 7, 104, 2, 2, 88, 89, 7,
	119, 2, 2, 89, 90, 7, 112, 2, 2, 90, 91, 7, 101, 2, 2, 91, 30, 3, 2, 2,
	2, 92, 93, 7, 107, 2, 2, 93, 94, 7, 112, 2, 2, 94, 95, 7, 118, 2, 2, 95,
	96, 7, 103, 2, 2, 96, 97, 7, 116, 2, 2, 97, 98, 7, 104, 2, 2, 98, 99, 7,
	99, 2, 2, 99, 100, 7, 101, 2, 2, 100, 101, 7, 103, 2, 2, 101, 32, 3, 2,
	2, 2, 102, 103, 7, 111, 2, 2, 103, 104, 7, 99, 2, 2, 104, 105, 7, 107,
	2, 2, 105, 106, 7, 112, 2, 2, 106, 34, 3, 2, 2, 2, 107, 108, 7, 114, 2,
	2, 108, 109, 7, 99, 2, 2, 109, 110, 7, 101, 2, 2, 110, 111, 7, 109, 2,
	2, 111, 112, 7, 99, 2, 2, 112, 113, 7, 105, 2, 2, 113, 114, 7, 103, 2,
	2, 114, 36, 3, 2, 2, 2, 115, 116, 7, 116, 2, 2, 116, 117, 7, 103, 2, 2,
	117, 118, 7, 118, 2, 2, 118, 119, 7, 119, 2, 2, 119, 120, 7, 116, 2, 2,
	120, 121, 7, 112, 2, 2, 121, 38, 3, 2, 2, 2, 122, 123, 7, 117, 2, 2, 123,
	124, 7, 118, 2, 2, 124, 125, 7, 116, 2, 2, 125, 126, 7, 119, 2, 2, 126,
	127, 7, 101, 2, 2, 127, 128, 7, 118, 2, 2, 128, 40, 3, 2, 2, 2, 129, 130,
	7, 118, 2, 2, 130, 131, 7, 123, 2, 2, 131, 132, 7, 114, 2, 2, 132, 133,
	7, 103, 2, 2, 133, 42, 3, 2, 2, 2, 134, 135, 7, 107, 2, 2, 135, 136, 7,
	111, 2, 2, 136, 137, 7, 114, 2, 2, 137, 138, 7, 113, 2, 2, 138, 139, 7,
	116, 2, 2, 139, 140, 7, 118, 2, 2, 140, 44, 3, 2, 2, 2, 141, 142, 7, 104,
	2, 2, 142, 143, 7, 111, 2, 2, 143, 144, 7, 118, 2, 2, 144, 46, 3, 2, 2,
	2, 145, 146, 7, 82, 2, 2, 146, 147, 7, 116, 2, 2, 147, 148, 7, 107, 2,
	2, 148, 149, 7, 112, 2, 2, 149, 150, 7, 118, 2, 2, 150, 151, 7, 104, 2,
	2, 151, 48, 3, 2, 2, 2, 152, 154, 9, 2, 2, 2, 153, 152, 3, 2, 2, 2, 154,
	50, 3, 2, 2, 2, 155, 156, 4, 50, 59, 2, 156, 52, 3, 2, 2, 2, 157, 160,
	5, 49, 25, 2, 158, 160, 7, 97, 2, 2, 159, 157, 3, 2, 2, 2, 159, 158, 3,
	2, 2, 2, 160, 166, 3, 2, 2, 2, 161, 165, 5, 49, 25, 2, 162, 165, 7, 97,
	2, 2, 163, 165, 5, 51, 26, 2, 164, 161, 3, 2, 2, 2, 164, 162, 3, 2, 2,
	2, 164, 163, 3, 2, 2, 2, 165, 168, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 166,
	167, 3, 2, 2, 2, 167, 54, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 169, 171, 9,
	3, 2, 2, 170, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 170, 3, 2, 2,
	2, 172, 173, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 175, 8, 28, 2, 2, 175,
	56, 3, 2, 2, 2, 176, 177, 7, 49, 2, 2, 177, 178, 7, 44, 2, 2, 178, 182,
	3, 2, 2, 2, 179, 181, 11, 2, 2, 2, 180, 179, 3, 2, 2, 2, 181, 184, 3, 2,
	2, 2, 182, 183, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 183, 185, 3, 2, 2, 2,
	184, 182, 3, 2, 2, 2, 185, 186, 7, 44, 2, 2, 186, 187, 7, 49, 2, 2, 187,
	188, 3, 2, 2, 2, 188, 189, 8, 29, 3, 2, 189, 58, 3, 2, 2, 2, 190, 191,
	7, 49, 2, 2, 191, 192, 7, 49, 2, 2, 192, 196, 3, 2, 2, 2, 193, 195, 10,
	4, 2, 2, 194, 193, 3, 2, 2, 2, 195, 198, 3, 2, 2, 2, 196, 194, 3, 2, 2,
	2, 196, 197, 3, 2, 2, 2, 197, 199, 3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 199,
	200, 8, 30, 3, 2, 200, 60, 3, 2, 2, 2, 10, 2, 153, 159, 164, 166, 172,
	182, 196, 4, 8, 2, 2, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "'\"'", "'('", "')'", "'{'", "'_'", "'='", "'.'", "'%'", "'#'",
	"'v'", "','", "'}'", "'func'", "'interface'", "'main'", "'package'", "'return'",
	"'struct'", "'type'", "'import'", "'fmt'", "'Printf'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "FUNC", "INTERFACE",
	"MAIN", "PACKAGE", "RETURN", "STRUCT", "TYPE", "IMPORT", "FMT", "PRINTF",
	"NAME", "WHITESPACE", "COMMENT", "LINE_COMMENT",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "FUNC", "INTERFACE", "MAIN", "PACKAGE",
	"RETURN", "STRUCT", "TYPE", "IMPORT", "FMT", "PRINTF", "LETTER", "DIGIT",
	"NAME", "WHITESPACE", "COMMENT", "LINE_COMMENT",
}

type FGLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewFGLexer(input antlr.CharStream) *FGLexer {

	l := new(FGLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "FG.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FGLexer tokens.
const (
	FGLexerT__0         = 1
	FGLexerT__1         = 2
	FGLexerT__2         = 3
	FGLexerT__3         = 4
	FGLexerT__4         = 5
	FGLexerT__5         = 6
	FGLexerT__6         = 7
	FGLexerT__7         = 8
	FGLexerT__8         = 9
	FGLexerT__9         = 10
	FGLexerT__10        = 11
	FGLexerT__11        = 12
	FGLexerT__12        = 13
	FGLexerFUNC         = 14
	FGLexerINTERFACE    = 15
	FGLexerMAIN         = 16
	FGLexerPACKAGE      = 17
	FGLexerRETURN       = 18
	FGLexerSTRUCT       = 19
	FGLexerTYPE         = 20
	FGLexerIMPORT       = 21
	FGLexerFMT          = 22
	FGLexerPRINTF       = 23
	FGLexerNAME         = 24
	FGLexerWHITESPACE   = 25
	FGLexerCOMMENT      = 26
	FGLexerLINE_COMMENT = 27
)
