package he_ast

import (
	"helang-go/helang/core"
)

type BasicAST interface {
	Evaluate(map[string]*core.U8) (*core.U8, error)
}
