package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var mapPackage = map[string]string{}

type pair struct {
	left  string
	right string
}

var graph []pair

func split(text string, count int) int {
	ss := strings.Split(text, " ")
	for _, s := range ss {
		if _, ok := mapPackage[s]; !ok {
			count++
			mapPackage[s] = fmt.Sprintf("P%d", count)
		}
	}
	p := pair{left: mapPackage[ss[0]], right: mapPackage[ss[1]]}
	graph = append(graph, p)
	return count
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for scanner.Scan() {
		count = split(scanner.Text(), count)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fmt.Println("@startuml")
	fmt.Println("hide circle")
	fmt.Println("hide empty members")
	for k, v := range mapPackage {
		fmt.Printf("class \"%v\" as %v\n", strings.Replace(k, "@", "\\n", -1), v)
	}
	for _, p := range graph {
		fmt.Printf("\"%v\" ..> \"%v\"\n", p.left, p.right)
	}
	fmt.Println("@enduml")

	return
}
