package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var funcSignagure = regexp.MustCompile(`\((.*)\).*`)

func runJavap(classpath, importpath string) Clazz {
	cmd := exec.Command("javap", "-public", "-s", "-classpath", classpath, importpath)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic("Failed to exec javap")
	}

	re := regexp.MustCompile(`\s(\S*?)(|\(.*?\))\s*(|throws.*?);`)

	// // []["symbol", "signature"]
	// variables := make([][]string, 0)
	// staticVariables := make([][]string, 0)
	// functions := make([][]string, 0)
	// staticFunctions := make([][]string, 0)

	// Compiled from ... : ignore
	// ..... <ClassName> (implements xxx) {
	//  .... <Variable name>;
	//  descriptor: <Signature>
	// (empty line): ignore
	// ..... <Method name>(.....);
	// descriptor: <Signature>
	// } : ignore
	cmd.Start()
	scanner := bufio.NewScanner(stdout)

	importpathArray := strings.Split(importpath, ".")
	clazz := Clazz{
		Name:           importpathArray[len(importpathArray)-1],
		FQCN:           strings.Replace(importpath, ".", "/", -1),
		StaticFunction: make(map[string]StaticFunction, 0),
		Function:       make(map[string]Function, 0),
		StaticVariable: make([]StaticVariable, 0),
		Variable:       make([]Variable, 0),
	}
scan:
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "Compiled") || text == "" || text == "}" {
			continue scan
		}
		array := strings.Split(text, " ")

		if array[0] != "" {
			continue scan
		}

		match := re.FindStringSubmatch(text)
		symbol := match[1]

		// is function?
		function := (match[2] != "")

		// is static?
		static := false
		for _, word := range array {
			if word == "static" {
				static = true
			}
			if word == "abstract" {
				scanner.Scan()
				continue scan
			}
		}

		// Get next line
		scanner.Scan()
		next := strings.Split(scanner.Text(), " ")

		signature := next[len(next)-1]

		if function {
			argSig := funcSignagure.FindStringSubmatch(signature)[1]
			if static {
				if staticFunction, ok := clazz.StaticFunction[symbol]; ok {
					staticFunction.Signature[argSig] = signature
				} else {
					clazz.StaticFunction[symbol] = StaticFunction{
						Symbol:    symbol,
						Signature: map[string]string{argSig: signature},
					}
				}
			} else {
				if function, ok := clazz.Function[symbol]; ok {
					function.Signature[argSig] = signature
				} else {
					clazz.Function[symbol] = Function{
						Symbol:    symbol,
						Signature: map[string]string{argSig: signature},
					}
				}
			}
		} else {
			if static {
				clazz.StaticVariable = append(clazz.StaticVariable, StaticVariable{
					Symbol:    symbol,
					Signature: signature,
				})
			} else {
				clazz.Variable = append(clazz.Variable, Variable{
					Symbol:    symbol,
					Signature: signature,
				})
			}
		}

	}
	cmd.Wait()
	return clazz
}

type Clazz struct {
	Name           string
	FQCN           string
	StaticFunction map[string]StaticFunction
	Function       map[string]Function
	StaticVariable []StaticVariable
	Variable       []Variable
}

type StaticFunction struct {
	Symbol    string
	Signature map[string]string // [Arg]Full
}

func (f StaticFunction) SignatureMap() string {
	return fmt.Sprintf("%#v", f.Signature)
}

type Function struct {
	Symbol    string
	Signature map[string]string // [Arg]Full
}

func (f Function) SignatureMap() string {
	return fmt.Sprintf("%#v", f.Signature)
}

type StaticVariable struct {
	Symbol    string
	Signature string
}

type Variable struct {
	Symbol    string
	Signature string
}
