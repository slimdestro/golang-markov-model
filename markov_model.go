// This function takes a string and returns a Markov chain of the string
package main

import (
	"fmt"
	"strings"
)

func MarkovChain(s string) map[string][]string {
	chain := make(map[string][]string)
	words := strings.Split(s, " ")
	for i := 0; i < len(words)-1; i++ {
		curr := words[i]
		next := words[i+1]
		if _, ok := chain[curr]; ok {
			chain[curr] = append(chain[curr], next)
		} else {
			chain[curr] = []string{next}
		}
	}

	return chain
}

func main() {
	fmt.Println(MarkovChain("This is a test string"))
}


// you can place below code inside _test.go file if youre using this repo
func TestMarkovChain(t *testing.T) {
	s := "This is a test string"
	expected := map[string][]string{
		"This":  []string{"is"},
		"is":    []string{"a"},
		"a":     []string{"test"},
		"test":  []string{"string"},
		"string": []string{},
	}

	actual := MarkovChain(s)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}