package demo

import (
	"errors"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ASTDemo() {
	fset := token.NewFileSet()
	// 这里取绝对路径，方便打印出来的语法树可以转跳到编辑器
	path, _ := filepath.Abs("./demo/ast_demo.go")
	f, err := parser.ParseFile(fset, path, nil, parser.AllErrors)
	if err != nil {
		log.Println(err)
		return
	}
	// 打印语法树
	//ast.Print(fset, f)
	fmt.Println(f.Imports)
	for _, item := range f.Imports {
		fmt.Println(item.Name)
	}
}

func getImportPkg(pkg string) (string, error) {
	p, err := build.Import(pkg, "", build.FindOnly)
	if err != nil {
		return "", err
	}
	return p.Dir, err
}

func parseDir(dir, pkgName string) (*ast.Package, error) {
	pkgMap, err := parser.ParseDir(
		token.NewFileSet(),
		dir,
		func(info os.FileInfo) bool {
			// skip go-test
			return !strings.Contains(info.Name(), "_test.go")
		},
		parser.Mode(0), // no comment
	)
	if err != nil {
		return nil, err
	}

	pkg, ok := pkgMap[pkgName]
	if !ok {
		err := errors.New("not found")
		return nil, err
	}

	return pkg, nil
}
