package main

import (
	"strings"
	"testing"
)

var sent [][]string

func TestMain(t *testing.T) {
	sent = readWords("latin.txt")
}

func TestLoremPara(t *testing.T) {
	np := 4
	res := lorem(sent, np, 0, 0, 0)
	if n := countPars(res); n != np {
		t.Fatalf("paragraph number not expected! %d vs %d\n", n, np)
	}
}

func TestLoremWords(t *testing.T) {
	nw := 123
	res := lorem(sent, 0, 0, nw, 0)
	if n := WordCount(res); n != nw {
		t.Fatalf("word number not expected! %d vs %d\n", n, nw)
	}
}

func TestLoremBytes(t *testing.T) {
	nb := 1116
	res := lorem(sent, 0, 0, 0, nb)
	if n := len(res); n != nb {
		t.Fatalf("byte length not expected! %d vs %d\n", n, nb)
	}
}

func countPars(str string) int {
	splitFunc := func(c rune) bool {
		return (c == '\n')
	}
	pars := strings.FieldsFunc(str, splitFunc)
	return len(pars)
}
