package compilestudy

import (
	"go/ast"
	"go/parser"
	"go/token"
)

//抽象语法树解析
func main() {
	/*
		前端：构建一个结构化的源码文件，预处理
			a:词法和语法分析  cmd/compile/internal/syctax
			（1）将源码字节流输入词法解析器，输出为token序列--即go语言定义的关键字，去除空格逗号等无效字符
			（2）将token序列输入语法分析器按照规定好的文法--LALR，转换为抽象语法树
			（3）再次过程中的语法错误会终止分析过程；
			b:类型检查typecheck与转换AST cmd/complile/internal/gc
			（1）深度优先遍历AST--多叉树 的节点进行类型检查
			（2）改写关键字、展开语法糖 比如go、make、map
		后端：优化代码，并转化成可执行文件
			a.转换为SSA，生成中间代码 cmd/compile/internal/gc.compileFunctions
			  遍历并替换AST上某些关键字，将其转换为某个具体的运行函数
			  cmd/complile/internal/gc 和 cmd/complile/internal/ssa将所有文件中的函数抽象出来
			  加入一个队列，用协程并发生成中间代码；

			b.生成机器代码 cmd/complile/internal/ssa和cmd/internal/obj
			  cmd/complile/internal 链接不同的指令集生成对应平台的机器码

			  SSA代码后的产物
			  pgen.go Compile --> pp.Flush() 后面是汇编的执行 调用汇编器生成机器码
			  plan9 汇编器


		查看中间代码的生成过程：
		go build -gcflags -S main.go
		查看编译器的优化过程：
		GOSSAFUNC=main go build main.go
	*/
	src := `
	package main
	import "fmt"
	func main() {
		i := 10
		go func(){
			k := 2
		}()
		fmt.Println("%d".i)
	}`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}
	ast.Print(fset, f)
}
