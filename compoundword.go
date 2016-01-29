package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

var filename string
var useTrie bool
var useTime bool

func init() {
	flag.StringVar(&filename, "f", "word.list", "file to load the list from")
	flag.BoolVar(&useTrie, "t", false, "use the trie implementation instead of the regular one")
	flag.BoolVar(&useTime, "time", false, "set to output the time it takes to find the word(s)")
}

type lengthList []string

func (l lengthList) Len() int {
	return len(l)
}

func (l lengthList) Less(i, j int) bool {
	return len(l[i]) > len(l[j])
}

func (l lengthList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func FindWord(list []string) []string {
	sList := lengthList(list)
	//Sort the list from largest to smallest so we can stop once we've found a matching word and checked the rest of the words of the same size.
	//This also has the advantage that we can start checking words after the word we're on because
	//a word cannot contain a word that is bigger than it is.
	sort.Sort(sList)
	results := make([]string, 0, 1)
	for i := 0; i < len(sList)-1; i++ {
		//if we have at least one result and the word we're checking is smaller we're done
		if len(results) > 0 && len(sList[i]) < len(results[0]) {
			return results
		}
		if isCompound(sList[i], sList[i+1:]) {
			results = append(results, sList[i])
		}
	}
	return results
}

func isCompound(word string, list []string) bool {
	//Make sure that "words" with no characters don't match.
	if len(word) == 0 {
		return false
	}
	return checkList(word, list, word)
}

func checkList(word string, list []string, original string) bool {
	//Anchor the function when you run out of characters.
	if len(word) == 0 {
		return true
	}
	for i := 0; i < len(list); i++ {
		//Make sure we're not matching the word against another copy of itself.
		if list[i] != original {
			index := strings.Index(word, list[i])
			if index != -1 {
				//If the word is found within the word we're checking, take it out and check the two parts that are left.
				//It only passes in the word matched and the rest of the list because we've already checked everything before that.
				if checkList(word[:index], list[i:], word) && checkList(word[index+len(list[i]):], list[i:], word) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	flag.Parse()
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error loading file use -f filename to select a file containing the list: ", err)
		panic(err)
	}
	list := make([]string, 0, 100000)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		list = append(list, scan.Text())
	}
	if !useTrie {
		startTime := time.Now()
		results := FindWord(list)
		fmt.Println(results)
		timeTaken := time.Since(startTime)
		if useTime {
			fmt.Println("In: ", timeTaken)
		}
	} else {
		startTime := time.Now()
		results := FindWord2(list)
		fmt.Println(results)
		timeTaken := time.Since(startTime)
		if useTime {
			fmt.Println("In: ", timeTaken)
		}
	}
}

//FindWord2 is an alternative method using a trie.
func FindWord2(list []string) []string {
	m := NewMatcher()
	for i := 0; i < len(list); i++ {
		m.Insert(list[i])
	}
	sort.Sort(lengthList(list))
	results := make([]string, 0, 10)
	for i := 0; i < len(list); i++ {
		if len(results) > 0 && (len(list[i]) < len(results[0])) {
			return results
		}
		if m.Match(list[i]) {
			results = append(results, list[i])
		}
	}
	return nil
}

//ListTree is a simple implementation of a Trie for matching words on the list.
type ListTree struct {
	tree map[string]*node
}

func NewMatcher() *ListTree {
	t := new(ListTree)
	t.tree = make(map[string]*node)
	return t
}

func (lt *ListTree) Match(word string) bool {
	if len(word) == 0 {
		return false
	}
	return nextMatch(word, lt.tree, 0)
}
func nextMatch(word string, m map[string]*node, matches int) bool {
	current := m
	if len(word) == 0 {
		if matches > 1 {
			return true
		} else {
			return false
		}
	}
	for i := 0; i < len(word); i++ {
		match, ok := current[string(word[i])]
		if ok {
			if match.finished > 0 {
				if nextMatch(word[i+1:], m, matches+1) {
					return true
				}
			}
			if match.rest == nil {
				return false
			}
			current = match.rest
		} else {
			return false
		}
	}
	return false
}

type node struct {
	finished int
	rest     map[string]*node
}

func (lt *ListTree) Insert(word string) {
	current := new(node)
	current.rest = lt.tree
	for i := 0; i < len(word); i++ {
		match, ok := current.rest[string(word[i])]
		if ok {
			current = match
		} else {
			n := new(node)
			n.rest = make(map[string]*node)
			n.finished = 0
			current.rest[string(word[i])] = n
			current = n
		}
	}
	current.finished++
}
