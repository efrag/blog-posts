package main

import (
	"bufio"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

type Structor struct {
	pDir        string // package directory
	getTemplate *template.Template
	setTemplate *template.Template
}

type parsedFile struct {
	PackageName string
	Imports     []string
	StructName  string
	Fields      []parsedField
}

type parsedField struct {
	SName string
	FName string
	FType string
	CName string
}

func newParsedFile(name string) parsedFile {
	return parsedFile{
		PackageName: name,
		Imports:     make([]string, 0),
		Fields:      make([]parsedField, 0),
	}
}

func (s *Structor) Generate() error {
	pFiles, err := s.parse()
	if err != nil {
		return err
	}

	// for each struct let's create the accessors
	for _, pFile := range pFiles {
		w, err := s.getWriter(pFile.StructName)
		if err != nil {
			return err
		}
		w.WriteString("// DO NOT EDIT: file has been automatically generated\n")
		w.WriteString("package " + pFile.PackageName + "\n\n")

		for _, i := range pFile.Imports {
			w.WriteString("import " + i + "\n")
		}

		for _, f := range pFile.Fields {
			err := s.getTemplate.Execute(w, f)
			if err != nil {
				return err
			}

			err = s.setTemplate.Execute(w, f)
			if err != nil {
				return err
			}
		}

		w.Flush()
	}

	return nil
}

func (s *Structor) getWriter(stName string) (*bufio.Writer, error) {
	fn := structAccessorFileName(stName)
	f, err := os.Create(path.Join(s.pDir, fn))
	if err != nil {
		return nil, err
	}

	return bufio.NewWriter(f), nil
}

func (s *Structor) parse() ([]parsedFile, error) {
	pFiles := make([]parsedFile, 0)

	files, err := ioutil.ReadDir(s.pDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if strings.HasSuffix(file.Name(), "_accessors.go") {
			continue
		}
		fn := path.Join(s.pDir, file.Name())

		src, err := ioutil.ReadFile(fn)
		if err != nil {
			return nil, err
		}

		// Create the AST by parsing src.
		fset := token.NewFileSet() // positions are relative to fset
		f, err := parser.ParseFile(fset, fn, src, 0)
		if err != nil {
			return nil, err
		}

		// Inspect the AST and print all identifiers and literals.
		pFile := newParsedFile(f.Name.Name)
		ast.Inspect(f, func(n ast.Node) bool {
			switch t := n.(type) {
			case *ast.TypeSpec:
				e, ok := t.Type.(*ast.StructType)
				if ok {
					pFile.StructName = t.Name.Name
					for _, f := range e.Fields.List {
						pFile.Fields = append(pFile.Fields, parsedField{
							SName: t.Name.Name,
							FName: f.Names[0].Name,
							CName: strings.Title(f.Names[0].Name),
							FType: string(src[f.Type.Pos()-1 : f.Type.End()-1]),
						})
					}
				}
			case *ast.ImportSpec:
				pFile.Imports = append(pFile.Imports, t.Path.Value)
			}
			return true
		})
		pFiles = append(pFiles, pFile)
	}

	return pFiles, nil
}

func structAccessorFileName(name string) string {
	sn := toSnakeCase(name)
	return sn + "_accessors.go"
}

func toSnakeCase(str string) string {
	// https://gist.github.com/stoewer/fbe273b711e6a06315d19552dd4d33e6
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}

func NewStructor(pDirectory string) *Structor {
	return &Structor{
		pDir: pDirectory,
		getTemplate: template.Must(template.New("getter").Parse(`
func (t *{{.SName}}) Get{{ .CName }}() {{ .FType }} {
	return t.{{ .FName }}
}
`)),
		setTemplate: template.Must(template.New("setter").Parse(`
func (t *{{.SName}}) Set{{ .CName }}(f {{ .FType }}) {
	t.{{ .FName}} = f
}
`)),
	}
}
