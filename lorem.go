package main

import (
	"math/rand"
	"strings"
)

func lorem(sent [][]string, nPar, nSent, nWords, nBytes int) string {
	b := strings.Builder{}

	var outS, outW, outB int
	s := 0
	writeSentence := func(maxWords int) int {
		end := min(len(sent[s]), maxWords)
		str := strings.Join(sent[s][:end], " ")
		b.WriteString(str)
		outW += len(strings.Split(str, " ")) // placeholder
		outB += len(str)
		outS++
		b.WriteString(". ")
		s = (s + 1) % len(sent)
		return end
	}
	getSpP := func() int {
		min := 3
		max := 9
		return rand.Intn(max-min) + min
	}
	if nPar > 1 {
		// paragraph mode - generate n paragraphs
		// random n of random TODO non-repeating sentences
		for i := 0; i < nPar; i++ {
			rnd := getSpP()
			for j := 0; j < rnd; j++ {
				writeSentence(1000)
			}
			b.WriteString("\n\n")
		}
	} else if nSent > 0 {
		// sentence mode - n random sentences
		for i := 0; i < nSent; i++ {
			writeSentence(1000)
		}
	} else if nWords > 0 {
		// word mode - add sentences until word count is met
		// limit last sentence length - TODO - smarter (?), sentence permutations by length (?)
		for outW < nWords {
			n := writeSentence(nWords - outW)
			if n < 1 {
				break // what?
			}
		}
	} else {
		// byte mode - TODO
		for outB < nBytes {
			n := writeSentence(1000)
			if n < 1 {
				break // what?
			}
		}
		// naiive cut string
		return b.String()[:nBytes-1] + "."
	}
	return b.String()
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
