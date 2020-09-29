package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}

	depends := dependency(importPath(os.Args[1:]))
	sort.StringSlice(depends).Sort()

	for _, pkg := range depends {
		fmt.Println(pkg)
	}
}

// 获取指定包的导入路径
func importPath(patterns []string) []string {
	args := []string{"list", "-f={{.ImportPath}}"}
	for _, pkg := range patterns {
		args = append(args, pkg)
	}
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		commandExitError("search importPath", err)
	}
	return strings.Fields(string(out))
}

// 遍历 GOPATH 中所有包，若某个包的导入依赖中包含指定的包就追加到返回列表
func dependency(importPaths []string) []string {

	args := []string{"list", `-f={{.ImportPath}} {{join .Deps " "}}`, "..."}
	out, err := exec.Command("go", args...).Output()
	if err != nil {
		commandExitError("traverse importPath", err)
	}

	contains := make(map[string]bool)
	for _, pkg := range importPaths {
		contains[pkg] = true
	}

	var depends []string
	s := bufio.NewScanner(bytes.NewReader(out))
	for s.Scan() {
		fields := strings.Fields(s.Text())
		pkg := fields[0]
		deps := fields[1:]
		for _, dep := range deps {
			if contains[dep] {
				depends = append(depends, pkg)
				break
			}
		}
	}
	return depends
}

func commandExitError(context string, err error) {
	exitError, ok := err.(*exec.ExitError)
	if !ok {
		log.Fatalf("%s: %s", context, err)
	}
	log.Printf("%s: %s", context, err)
	os.Stderr.Write(exitError.Stderr)
	os.Exit(1)
}
