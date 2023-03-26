package evaluator

import (
	"monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len":   object.GenBuiltinByName("len"),
	"puts":  object.GenBuiltinByName("puts"),
	"first": object.GenBuiltinByName("first"),
	"last":  object.GenBuiltinByName("last"),
	"rest":  object.GenBuiltinByName("rest"),
	"push":  object.GenBuiltinByName("push"),
}
