package main

import (
	"Go_exe/chapter7/exe7_13_14/eval"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入一个表达式：（例如：3 + x * 5）")
	stdin.Scan()
	exprStr := stdin.Text()
	if exprStr == "" {
		fmt.Println("表达式为空")
		os.Exit(0)
	}
	expr, err := eval.Parse(exprStr)
	if err != nil {
		fmt.Printf("Expression %s parses error: %v", exprStr, err)
		os.Exit(1)
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		fmt.Printf("Expression %s check error: %v", exprStr, err)
		os.Exit(2)
	}
	if len(vars) > 0 {
		env := eval.Env{}
		for v := range vars {
			fmt.Printf("请输入变量%s的值：", v)
			stdin.Scan()
			f, err := strconv.ParseFloat(stdin.Text(), 64)
			if err != nil {
				fmt.Printf("ParseFloat error: %v", err)
				os.Exit(3)
			}
			env[v] = f
		}
		got := fmt.Sprintf("%.6g", expr.Eval(env))
		fmt.Printf("%s=%s\n", exprStr, got)
	} else {
		got := fmt.Sprintf("%.6g", expr.Eval(nil))
		fmt.Printf("%s=%s\n", exprStr, got)
	}

}
