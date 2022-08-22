package core

import "errors"

var HeLangException = errors.New("HeLangException")

var BadTokenException = errors.New("BadTokenException")

var BadStatementException = errors.New("BadStatementException")

var CyberNameException = errors.New("CyberNameException")

var CyberSubtractionException = errors.New("CyberSubtractionException")

var CyberGrammarException = errors.New("CyberLexerException")

var CyberNotSupportedException = errors.New("CyberNotSupportedException")

var CyberNetworkException = errors.New("CyberNetworkException")