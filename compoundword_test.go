package main

import (
	"testing"
)

var wordlists = []struct {
	list     []string
	expected []string
}{
	{
		[]string{
			"aaaccc",
			"aacc",
			"aaa",
			"cc",
			"c",
		},
		[]string{
			"aaaccc",
		},
	},
	{
		[]string{
			"ball",
			"base",
			"sit",
			"superlongWord",
			"baseball",
		},
		[]string{
			"baseball",
		},
	},
	{
		[]string{
			"baseball",
			"baseball",
			"homerun",
			"home",
			"run",
			"longerwordthathaswordsnotinthelist",
		},
		[]string{
			"homerun",
		},
	},
	{
		[]string{
			"baseball",
			"ball",
			"baseball",
			"homerun",
			"base",
			"home",
			"run",
			"longerwordthathaswordsnotinthelist",
		},
		[]string{
			"baseball",
			"baseball",
		},
	},
	{
		[]string{
			"home",
			"baseball",
			"homework",
			"ball",
			"base",
			"work",
			"otherlongerword",
		},
		[]string{
			"homework",
			"baseball",
		},
	},
	{
		[]string{
			"home",
			"basebaseball",
			"homework",
			"ball",
			"base",
			"work",
			"otherlongerword",
		},
		[]string{
			"basebaseball",
		},
	},
}

func TestFindWord(t *testing.T) {
	for i := 0; i < len(wordlists); i++ {
		result := FindWord(wordlists[i].list)
		if !checkResult(result, wordlists[i].expected) {
			t.Error("Result and expected don't match", "\n\nResult:   ", result, "\nExpected: ", wordlists[i].expected, "\nList: ", i)
		}
	}
}

func TestFindWord2(t *testing.T) {
	for i := 0; i < len(wordlists); i++ {
		result := FindWord2(wordlists[i].list)
		if !checkResult(result, wordlists[i].expected) {
			t.Error("Result and expected don't match", "\n\nResult:   ", result, "\nExpected: ", wordlists[i].expected, "\nList: ", i)
		}
	}
}

func checkResult(result, expected []string) bool {
	if len(result) != len(expected) {
		return false
	}
	for i := 0; i < len(result); i++ {
		found := false
		for y := 0; y < len(expected); y++ {
			if expected[y] == result[i] {
				found = true
			}
		}
		if false == found {
			return false
		}
	}
	return true
}
