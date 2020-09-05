package main

import (
	"fmt"
	"log"
	"sort"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"}, // loop

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"linear algebra":        {"calculus"}, // loop
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	order, valid := topoSort(prereqs)
	if !valid {
		log.Fatalln("Loop detected")
	}

	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, bool) {
	var order []string
	seen := make(map[string]bool)

	loopChecker := make(map[string]bool)

	var visitAll func(items []string) bool
	visitAll = func(items []string) bool {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				loopChecker[item] = true
				if !visitAll(m[item]) {
					return false
				}
				loopChecker[item] = false
				order = append(order, item)
			} else if loopChecker[item] {
				return false
			}
		}
		return true
	}

	var courses []string
	for key := range m {
		courses = append(courses, key)
	}
	sort.Strings(courses)

	if found := visitAll(courses); found {
		return order, true
	}
	return nil, false
}
