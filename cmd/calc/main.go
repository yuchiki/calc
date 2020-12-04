package main

import (
	"flag"
	"fmt"
	"strconv"

	parsec "github.com/prataprc/goparsec"
)

func main() {
	flag.Parse()
	expression := flag.Arg(0)
	scanner := parsec.NewScanner([]byte(expression)).TrackLineno()

	fmt.Println(expression)

	ast := parsec.NewAST("expr", 1000)

	parser := makeParser(ast)
	node, scanner := ast.Parsewith(parser, scanner)

	_ = node

	ast.Prettyprint()

	answer := eval(node)
	fmt.Println(answer)
}

func makeParser(ast *parsec.AST) parsec.Parser {
	plus := parsec.Atom("+", "PLUS")
	parser := ast.And("plusExpr", nil, parsec.Int(), plus, parsec.Int())
	return parser
}

func eval(node parsec.Queryable) int {
	i1, _ := strconv.Atoi(node.GetChildren()[0].GetValue())
	i3, _ := strconv.Atoi(node.GetChildren()[2].GetValue())

	answer := i1 + i3
	return answer
}
