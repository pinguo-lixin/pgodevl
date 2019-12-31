package cmd

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var (
	controllerPath string
	outputComments bool
)

// listurlCmd represents the listurl command
var listurlCmd = &cobra.Command{
	Use:   "listurl",
	Short: "List urls",
	Long: `Analyze all controllers (structs under "controller" directory and with "Controller" suffix)
and it's methods with "Action" prefix to generate urls`,
	Run: func(cmd *cobra.Command, args []string) {
		listURL(controllerPath, outputComments)
	},
}

func init() {
	rootCmd.AddCommand(listurlCmd)

	listurlCmd.Flags().BoolVarP(&outputComments, "comments", "c", true, "output comments")
	listurlCmd.Flags().StringVarP(&controllerPath, "p", "", "pkg/controller", "specify controller path")
}

var urlFormatRegexp = regexp.MustCompile("\\B([A-Z])")

func urlFormat(s string) string {
	r := urlFormatRegexp.ReplaceAllString(s, "-$1")
	return strings.ToLower(r)
}

func listURL(controllerPath string, withComments bool) {
	urls, urlComments := make([]string, 0), make([]string, 0)

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
							u := fmt.Sprintf("%s%s/%s", urlPathes, a, b)
							urls = append(urls, u)
							if withComments {
								urlComments = append(urlComments, methodComment(typ))
							}
						}
					}
					return true
				})
			}
		}
		return nil
	})

	for i, u := range urls {
		if withComments {
			fmt.Printf("%-30s%s\n", u, urlComments[i])
		} else {
			fmt.Printf("%s\n", u)
		}
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

func methodComment(n *ast.FuncDecl) string {
	if n.Doc.Text() == "" {
		return ""
	}
	c := strings.Split(n.Doc.Text(), "\n")
	if len(c) < 1 {
		return ""
	}
	return strings.TrimLeft(c[0], n.Name.Name)
}
