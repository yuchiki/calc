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

	plus := parsec.Atom("+", "PLUS")

	ast := parsec.NewAST("expr", 1000)
	parser := ast.And("plusExpr", nil, parsec.Int(), plus, parsec.Int())

	node, scanner := ast.Parsewith(parser, scanner)

	_ = node

	ast.Prettyprint()

	i1, _ := strconv.Atoi(node.GetChildren()[0].GetValue())
	i3, _ := strconv.Atoi(node.GetChildren()[2].GetValue())

	answer := i1 + i3
	fmt.Println(answer)
}
