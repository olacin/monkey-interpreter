# monkey-interpreter

This repository contains source code for the "Monkey" programming language interpreter I built while reading [Thorsten Ball](https://github.com/mrnugget)'s book [_Write an Interpreter in Go_](https://interpreterbook.com/).

## REPL

A basic REPL is available to try the language:
```
go run ./cmd/repl
Monkey REPL 0.0.1
>>> let add = fn(x, y) {x + y}
>>> let x = add(3, 5)
>>> x
8
```