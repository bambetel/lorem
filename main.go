package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("hello")
	var nPar, nSent, nWords, nChars int
	flag.IntVar(&nPar, "p", 1, "n paragraphs")
	flag.IntVar(&nSent, "s", 0, "n sentences")
	flag.IntVar(&nWords, "w", 0, "n words")
	flag.IntVar(&nChars, "c", 0, "n characters")
	flag.Parse()

	// sent := readSentences("latin.txt")
	sss := readWords("latin.txt")
	www := GetUnique(sss)
	fmt.Printf("GET WORDS\n%v\n", www)
	res := lorem(sss, nPar, nSent, nWords, nChars)
	fmt.Println(res)
}

func readWords(path string) [][]string {
	fmt.Println("readWords()")
	res := make([][]string, 0)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(SplitAt([]byte{'\n', '.', '!', '?'}))
	i := 0
	for scanner.Scan() {
		if line := strings.TrimSpace(scanner.Text()); len(line) > 0 {
			// sentences := strings.FieldsFunc(line, splitSent)
			// fmt.Printf("%d: %v@\n", i, line)
			res = append(res, GetWords(line))
			i++
		}
	}
	fmt.Printf("readWords returning %d sentences\n", len(res))
	return res
}

func GetUnique(sent [][]string) map[string]int {
	words := make(map[string]int)
	for i := 0; i < len(sent); i++ {
		for j := 0; j < len(sent[i]); j++ {
			words[sent[i][j]] = words[sent[i][j]] + 1

		}
	}
	return words
}

func SplitAt(sep []byte) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		dataLen := len(data)
		if atEOF && dataLen == 0 {
			return 0, nil, nil
		}
		i := 0
	strcspn:
		for j := 0; j < dataLen; j++ {
			for _, s := range sep {
				if data[j] == s {
					i = j
					break strcspn
				}
			}
		}
		if i >= 0 {
			return i + 1, data[0:i], nil
		}
		if atEOF {
			return dataLen, data, nil
		}
		return 0, nil, nil
	}
}

//	func readSentences(path string) []string {
//		file, err := os.Open(path)
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer file.Close()
//
//		scanner := bufio.NewScanner(file)
//		n := 0
//		sentences := make([]string, 0)
//		words := make(map[string]int)
//
//		for scanner.Scan() {
//			line := strings.TrimSpace(scanner.Text())
//			if len(line) > 0 {
//				add := strings.FieldsFunc(line, splitSent)
//				for i := range add {
//					add[i] = strings.TrimSpace(add[i])
//					if len(add[i]) > 0 {
//						// fmt.Printf("%4d: %s\n", n, add[i])
//						sentences = append(sentences, add[i])
//						ww := GetWords(add[i])
//						for _, w := range ww {
//							_, ok := words[w]
//							if ok {
//								words[w]++
//							} else {
//								words[w] = 1
//							}
//						}
//						n++
//					}
//				}
//			}
//		}
//		// fmt.Printf("%q\n", sentences)
//		fmt.Printf("WORDS: %v\n", words)
//		// error checking - where?
//		if err := scanner.Err(); err != nil {
//			log.Fatal(err)
//		}
//		return sentences
//	}
func splitSent(c rune) bool {
	return (c == '\n' || c == '.')
}
