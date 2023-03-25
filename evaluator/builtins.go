package evaluator

import (
	"monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len":   object.GenBuiltinByName("len"),
	"puts":  object.GenBuiltinByName("puts"),
	"first": object.GenBuiltinByName("puts"),
	"last":  object.GenBuiltinByName("puts"),
	"rest":  object.GenBuiltinByName("puts"),
	"push":  object.GenBuiltinByName("puts"),
}
