package evaluator

import (
	"strings"

	"github.com/olacin/monkey-interpreter/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", arg.Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				if len(arg.Value) > 0 {
					return &object.String{Value: string(arg.Value[0])}
				}
				return NULL
			case *object.Array:
				if len(arg.Elements) > 0 {
					return arg.Elements[0]
				}
				return NULL
			default:
				return newError("argument to `first` not supported, got %s", arg.Type())
			}
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				if len(arg.Value) > 0 {
					chr := arg.Value[len(arg.Value)-1]
					return &object.String{Value: string(chr)}
				}
				return NULL
			case *object.Array:
				if len(arg.Elements) > 0 {
					return arg.Elements[len(arg.Elements)-1]
				}
				return NULL
			default:
				return newError("argument to `last` not supported, got %s", arg.Type())
			}
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				if len(arg.Value) > 0 {
					stringCopy := strings.Clone(arg.Value[1:])
					return &object.String{Value: stringCopy}
				}
				return NULL
			case *object.Array:
				if len(arg.Elements) > 0 {
					length := len(arg.Elements)
					newElements := make([]object.Object, length-1)
					copy(newElements, arg.Elements[1:length])
					return &object.Array{Elements: newElements}
				}
				return NULL
			default:
				return newError("argument to `rest` not supported, got %s", arg.Type())
			}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) > 0 {
					length := len(arg.Elements)
					newElements := make([]object.Object, length+1)
					copy(newElements, arg.Elements)
					newElements[length] = args[1]
					return &object.Array{Elements: newElements}
				}
				return NULL
			default:
				return newError("argument to `push` not supported, got %s", arg.Type())
			}
		},
	},
}
