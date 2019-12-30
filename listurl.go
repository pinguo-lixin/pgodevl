package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var urlFormatRegexp = regexp.MustCompile("\\B([A-Z])")

func urlFormat(s string) string {
	r := urlFormatRegexp.ReplaceAllString(s, "-$1")
	return strings.ToLower(r)
}

func listURL(controllerPath string) {
	var urls []string

	filepath.Walk(controllerPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			fSet := token.NewFileSet()
			pkgs, err := parser.ParseDir(fSet, path, nil, parser.ParseComments)
			if err != nil {
				return err
			}

			for _, pkg := range pkgs {
				var urlPathes string
				for fName := range pkg.Files {
					path := fName[len(controllerPath):]
					urlPathes = strings.ReplaceAll(filepath.Dir(path), "\\", "/")
					if urlPathes != "/" {
						urlPathes += "/"
					}
					break
				}

				ast.Inspect(pkg, func(node ast.Node) bool {
					if typ, ok := node.(*ast.FuncDecl); ok {
						structName, funcName := receiver(typ)
						if strings.HasSuffix(structName, "Controller") && strings.HasPrefix(funcName, "Action") {
							a, b := urlFormat(structName[0:len(structName)-10]), urlFormat(funcName[6:])
							urls = append(urls, fmt.Sprintf("%s%s/%s", urlPathes, a, b))
						}
					}
					return true
				})
			}
		}
		return nil
	})

	for _, u := range urls {
		fmt.Println(u)
	}
}

func receiver(f *ast.FuncDecl) (receiver, funcName string) {
	if f.Recv == nil || f.Recv.List == nil {
		return
	}

	r, ok := f.Recv.List[0].Type.(*ast.StarExpr)
	if !ok {
		return
	}

	return r.X.(*ast.Ident).Name, f.Name.Name
}
