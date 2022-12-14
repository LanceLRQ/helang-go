package core

import "errors"

var HeLangException = errors.New("HeLangException")

var HeLangCompilerException = errors.New("HeLangCompilerException")

var BadTokenException = errors.New("BadTokenException")

var BadStatementException = errors.New("BadStatementException")

var CyberNameException = errors.New("CyberNameException")

var CyberArithmeticException = errors.New("CyberArithmeticException")

var CyberGrammarException = errors.New("CyberLexerException")

var CyberNotSupportedException = errors.New("CyberNotSupportedException")

var CyberNetworkException = errors.New("CyberNetworkException")