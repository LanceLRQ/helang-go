package he_ast

import (
	"helang-go/helang/core"
)

type BasicAST interface {
	evaluate(map[string]*core.U8) (*core.U8, error)
}
