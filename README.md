[![Go Report Card](https://goreportcard.com/badge/github.com/emt7x/gocalc)](https://goreportcard.com/badge/github.com/emt7x/gocalc)

# GoCalc

GoCalc is a terminal based calculator that is formed from the use of an AST (Abstract Syntax Tree). It uses a `lexer` -> `parser` -> `ast` -> `evaluation` process, and follows the rules of BODMAS which can be seen when a calculation is processed.

# BODMAS

BODMAS is a mathematical acronym for Brackets, Division, Multiplication, Addition then Subtraction. It is a logical mathematical guideline followed by all calculators, and I have emulated this process inside of GoCalc, allowing for you to use the rules of BODMAS inside to get a true result.

# Syntax

For an addition expression, you would just input: `1 + 1` to return the value of `2`. You can also have a more advanced expression such as `(4 + 2) / 2` to return the value of `3`.

The supported operators for GoCalc are;

- Add `+`
- Subtract `-`
- Multiply `*`
- Divide `/`
- Modulus `%`

Examples of plausable expressions;

- `((1 + 1) / 3) * 5` -> `3.333333333333333`
- `1 * 2 / 3` -> `0.6666666666666666`
- `1 - 3 + 5 * 2` -> `8`
