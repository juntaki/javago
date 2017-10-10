package main

import (
	"bufio"
	"os/exec"
	"regexp"
	"strings"
)

func runJavap(classpath, importpath string) Clazz {
	cmd := exec.Command("javap", "-public", "-s", "-classpath", classpath, importpath)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic("wa")
		//return nil, err
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
		Name:    importpathArray[len(importpathArray)-1],
		Members: make([]Member, 0),
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

		clazz.Members = append(clazz.Members, Member{
			Symbol:    symbol,
			Signature: signature,
			Static:    static,
			Function:  function,
		})
	}
	cmd.Wait()
	return clazz
}

type Clazz struct {
	Name    string
	Members []Member
}

type Member struct {
	Symbol    string
	Signature string
	Static    bool
	Function  bool
}
