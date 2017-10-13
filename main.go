package main

import (
	"os"
	"text/template"

	"github.com/urfave/cli"
)

const templ string = `
package java

import "github.com/juntaki/jnigo"

var jvm *jnigo.JVM

func init() {
	jvm = jnigo.CreateJVM()
}

type {{.Name}} struct{
	jclass *jnigo.JClass
}

func New{{.Name}}(args ...interface{}) (*{{.Name}}, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil{
		return nil, err
	}
	jclass, err := jvm.NewJClass("{{.FQCN}}", convertedArgs)
	if err != nil{
		return nil, err
	}
	return &{{.Name}}{
		jclass: jclass,
	}, nil
}
{{range .Variable}}
func (c *{{$.Name}}) Get{{.Symbol}}() (jnigo.JObject, error) {
	return c.jclass.GetField("{{.Symbol}}", "{{.Signature}}")
}
{{end}}

{{range .Variable}}
func (c *{{$.Name}}) Set{{.Symbol}}(arg interface{}) (error) {
	convertedArg = jvm.Convert(arg)
	return c.jclass.SetField("{{.Symbol}}", convertedArg)
}
{{end}}

{{range .Function}}
func (c *{{$.Name}}) {{.Symbol}}(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := {{.SignatureMap}}
	return c.jclass.CallFunction("{{.Symbol}}", sigMap[sigArgs], convertedArgs)
}
{{end}}
{{range .StaticFunction}}
func {{$.Name}}{{.Symbol}}(args ...interface{}) (jnigo.JObject, error) {
	convertedArgs, err := jvm.ConvertAll(args)
	if err != nil {
		return nil, err
	}
	sigArgs := ""
	for _, arg := range convertedArgs {
		sigArgs += arg.Signature()
	}
	sigMap := {{.SignatureMap}}
	return jvm.CallStaticFunction("{{$.FQCN}}", "{{.Symbol}}", sigMap[sigArgs], convertedArgs)
}
{{end}}
`

func main() {
	var classpath string
	var classfile string

	app := cli.NewApp()
	app.Name = "javago"
	app.Usage = "Java binding generator for Go"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "classpath",
			Value:       "*:.",
			Usage:       "Specifies the path the javap command uses to look up classes.",
			Destination: &classpath,
		},
		cli.StringFlag{
			Name:        "classfile",
			Value:       "",
			Usage:       "classfile",
			Destination: &classfile,
		},
	}

	app.Action = func(c *cli.Context) error {
		if classfile == "" {
			cli.ShowCommandHelpAndExit(c, "javago", 1)
		}

		clazz := runJavap(classpath, classfile)

		tmpl := template.Must(template.New("class").Parse(templ))

		path := "java"
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := os.Mkdir(path, 0777)
			if err != nil {
				panic(err)
			}
		}
		file, err := os.Create("java/" + classfile + ".go")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		err = tmpl.Execute(file, clazz)
		if err != nil {
			panic(err)
		}
		return nil
	}

	app.Run(os.Args)
}
