package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"os"
	"strconv"
)

const WithoutFor = 0

func goCounter(counter int, forIterCounter int) int {
	if forIterCounter == 0 {
		counter++
	} else {
		counter--
		counter += forIterCounter
	}
	return counter
}

func goFinder(currentNode ast.Node, counter int, forIterCounter int) int {
	ast.Inspect(currentNode, func(n ast.Node) bool {
		if _, ok := n.(*ast.GoStmt); ok {
			counter = goCounter(counter, forIterCounter)
		}
		return true
	})
	return counter
}

func goInForFinder(currentNode ast.Node, counter int) int {
	ast.Inspect(currentNode, func(n ast.Node) bool {
		forStatement, ok := n.(*ast.ForStmt)
		if ok {
			iterCountField := forStatement.Cond.(*ast.BinaryExpr).Y
			iterCount, err := strconv.Atoi(types.ExprString(iterCountField)) // I haven't find another way to get count of iterations in for loop
			if err != nil {
				fmt.Printf("cannot convert 'for condition expression operand' to int: %v", err)
				iterCount = 0
			}
			counter = goFinder(forStatement, counter, iterCount)
		}
		return true
	})
	return counter
}

func countGoroutinesInFunc(filePath string, functionName string) (runCounter int) {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, filePath, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("cannot parse file %s", filePath)
	}

	ast.Inspect(file, func(n ast.Node) bool {
		funcDeclaration, ok := n.(*ast.FuncDecl)
		if ok && funcDeclaration.Name.String() == functionName {
			runCounter = goFinder(funcDeclaration, runCounter, WithoutFor)
			runCounter = goInForFinder(funcDeclaration, runCounter)
		}
		return true
	})
	return
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: ./main [file_path] [function_name]")
		return
	}
	filePath, functionName := args[0], args[1]
	fmt.Printf("I've count %d runs of goroutines in %s function of %s file.",
		countGoroutinesInFunc(filePath, functionName),
		functionName,
		filePath,
	)
}
