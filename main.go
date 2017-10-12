package main

import (
	"os"
	"text/template"

	"github.com/urfave/cli"
)

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
		clazz := runJavap(classpath, classfile)
		tmpl := template.Must(template.ParseFiles("class.tmpl"))

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
